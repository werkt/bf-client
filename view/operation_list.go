package view

import (
  "context"
  "fmt"
  "sync"
  "time"

  bfpb "github.com/buildfarm/buildfarm/build/buildfarm/v1test"
  ui "github.com/gizak/termui/v3"
  "github.com/hashicorp/golang-lru/v2"
  "github.com/werkt/bf-client/client"
  "github.com/golang/protobuf/ptypes"
  "google.golang.org/genproto/googleapis/longrunning"
)

type operation struct {
  target string
  mnemonic string
  build string
  workerStart time.Time
}

type operationList struct {
  list *client.List
  Filter string
  a *client.App
  ops []string
  opcache *lru.Cache[string, operation]
  mode int
  v View
  field int
  reversed bool
  queues []*client.Queue
  fetchToken string
}

func NewOperationList(a *client.App, mode int, v View) *operationList {
  opcache, _ := lru.New[string, operation](1024)
  var queues []*client.Queue
  if mode != 3 {
    c := bfpb.NewOperationQueueClient(a.Conn)
    start := time.Now()
    status, err := c.Status(context.Background(), &bfpb.BackplaneStatusRequest {
      InstanceName: "shard",
    })
    a.LastReapiLatency = time.Since(start)
    if err != nil {
      panic(err)
    }
    var names []string
    if mode == 1 {
      names = append(names, status.Prequeue.Name)
    } else {
      for _, queue := range status.OperationQueue.Provisions {
        names = append(names, queue.Name)
      }
    }
    for _, name := range names {
      queues = append(queues, client.NewQueue(context.Background(), a.Client, name))
    }
  }
  return &operationList {
    list: client.NewList(),
    a: a,
    mode: mode,
    opcache: opcache,
    field: 0,
    reversed: false,
    v: v,
    queues: queues,
    fetchToken: "",
  }
}

func (v operationList) fetchQueues(max int64, cb func(string) (*client.Operation, error)) []*client.Operation {
  var ops []*client.Operation
  for _, queue := range v.queues {
    ops = append(ops, queue.Slice(context.Background(), v.a.Client, 0, max, cb)...)
    if int64(len(ops)) >= max {
      break
    }
  }
  return ops
}

func (v operationList) fetchFiltered(filter string) []*client.Operation {
  var ops []*client.Operation
  c := longrunning.NewOperationsClient(v.a.Conn)
  for nextPageToken := "initial"; nextPageToken != ""; v.fetchToken = nextPageToken {
    r, err := c.ListOperations(context.Background(), &longrunning.ListOperationsRequest {
      Name: fmt.Sprintf("%s/executions", v.a.Instance),
      Filter: filter,
      PageSize: 100,
      PageToken: v.fetchToken,
    })
    if err != nil {
      panic(err)
    }
    opsPage := r.Operations
    for _, op := range opsPage {
      ops = append(ops, &client.Operation{ Name: op.Name })
    }
    nextPageToken = r.NextPageToken
  }
  return ops
}

func (v operationList) queuesLength() int64 {
  var sum int64 = 0
  for _, queue := range v.queues {
    l, err := queue.Length(context.Background(), v.a.Client)
    if err != nil {
      panic(err)
    }
    sum += int64(l)
  }
  return sum
}

func (v operationList) fetch() []*client.Operation {
  switch v.mode {
  case 1: return v.fetchQueues(20, client.ParsePrequeueName)
  case 2: return v.fetchQueues(20, client.ParseQueueName)
  case 3: return v.fetchFiltered("status=dispatched")
  case 4: return v.fetchFiltered(v.Filter)
  default:
    return make([]*client.Operation, 0)
  }
}

func (v *operationList) createOperationView() View {
  if len(v.ops) > 0 {
    return NewOperation(v.a, v.list.Rows[v.list.SelectedRow].(*stageEx).name, v)
  }
  return v
}

func (v *operationList) Handle(e ui.Event) View {
  switch e.ID {
  case "<Escape>", "q", "<C-c>":
    ui.Clear()
    return v.v
  case "j", "<Down>":
    v.list.ScrollDown()
  case "k", "<Up>":
    v.list.ScrollUp()
  case "h", "<Left>":
    v.field += 3
    v.field %= 4
  case "l", "<Right>":
    v.field++
    v.field %= 4
  case "<Enter>":
    ui.Clear()
    return v.createOperationView()
  case ">", "<":
    v.reversed = !v.reversed
  }
  return v
}

func (v *operationList) Update() {
  v.a.Fetches++
  ops := v.fetch()

  var wg sync.WaitGroup
  for _, op := range ops {
    if v.opcache.Contains(op.Name) {
      continue
    }
    if o, ok := v.a.Ops[op.Name]; !ok || o == nil {
      m := op.Metadata
      if m == nil {
        v.a.Fetches++
        wg.Add(1)
        go getExecution(v.a, op.Name, v.a.Conn, &wg)
      }
    }
  }
  wg.Wait()
  v.ops = make([]string, 0)
  for _, op := range ops {
    if v.opcache.Contains(op.Name) {
      v.ops = append(v.ops, op.Name)
    } else if o, ok := v.a.Ops[op.Name]; ok && o != nil {
      v.ops = append(v.ops, op.Name)
      m := client.RequestMetadata(o)
      if m != nil {
        eam, err := client.ExecutedActionMetadata(o)
        var workerStart time.Time
        if err == nil {
          workerStart, _ = ptypes.Timestamp(eam.WorkerStartTimestamp)
        }
        v.opcache.Add(op.Name, operation {
          build: m.CorrelatedInvocationsId,
          target: m.TargetId,
          mnemonic: m.ActionMnemonic,
          workerStart: workerStart,
        })
        v.a.Ops[op.Name] = nil
        v.a.Metadatas[op.Name] = op.Metadata
        var opInvocations []string
        var ok bool
        if opInvocations, ok = v.a.Invocations[m.ToolInvocationId]; !ok {
          opInvocations = make([]string, 1)
          v.a.Invocations[m.ToolInvocationId] = opInvocations
        }
        v.a.Invocations[m.ToolInvocationId] = append(opInvocations, op.Name)
      }
    }
  }
}

func (v operationList) modeTitle() string {
  switch v.mode {
  case 1: return "Prequeue"
  case 2: return "Queue"
  case 3: return "Dispatched"
  case 4: return fmt.Sprintf("Filter(%s)", v.Filter)
  default: return "Unknown"
  }
}

func (v operationList) Render() []ui.Drawable {
  fields := [...]string{"name", "target", "build", "mnemonic"}
  v.list.Title = fmt.Sprintf("%s Operations (%s) %d", v.modeTitle(), fields[v.field], len(v.ops))
  rows := make([]fmt.Stringer, len(v.ops))
  for i := 0; i < len(v.ops); i++ {
    name := v.ops[i]
    op, ok := v.opcache.Get(name)
    if !ok {
      panic("opcache didn't contain: " + name)
    }
    rows[i] = &stageEx{
      field: func () int { return v.field },
      name: name,
      fence: op.workerStart,
      target: op.target,
      mnemonic: op.mnemonic,
      build: op.build,
    }
  }
  v.list.Rows = rows
  fence := func(e1, e2 fmt.Stringer) bool {
    stageEx1, stageEx2 := e1.(*stageEx), e2.(*stageEx)
    return (stageEx1.fence.Compare(stageEx2.fence) < 0) != v.reversed
  }
  byExec(fence).Sort(v.list.Rows)
  v.list.SelectedRowStyle = ui.NewStyle(ui.ColorBlack, ui.ColorWhite)
  v.list.WrapText = false
  v.list.SetRect(0, 0, 80, 30)

  return []ui.Drawable { v.list }
}
