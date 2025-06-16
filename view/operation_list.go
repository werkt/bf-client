package view

import (
  "cmp"
  "context"
  "fmt"
  "maps"
  "slices"
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
  Name string
  a *client.App
  ops []string
  opcache *lru.Cache[string, operation]
  mode int
  v View
  field int
  reversed bool
  queues []*client.Queue
  fetchToken string
  grouped bool
}

func NewOperationList(a *client.App, mode int, v View) *operationList {
  opcache, _ := lru.New[string, operation](10000)
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
    Name: "executions",
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

func (v operationList) fetchFiltered(filter string, name string) []*client.Operation {
  var ops []*client.Operation
  c := longrunning.NewOperationsClient(v.a.Conn)
  for nextPageToken := "initial"; nextPageToken != ""; v.fetchToken = nextPageToken {
    req := &longrunning.ListOperationsRequest {
      Name: fmt.Sprintf("%s/%s", v.a.Instance, name),
      Filter: filter,
      PageSize: 100,
      PageToken: v.fetchToken,
    }
    r, err := c.ListOperations(context.Background(), req)
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
  case 3: return v.fetchFiltered("status=dispatched", v.Name)
  case 4: return v.fetchFiltered(v.Filter, v.Name)
  default:
    return make([]*client.Operation, 0)
  }
}

func (v *operationList) createOperationView() View {
  if len(v.ops) > 0 {
    return NewDocument(v.a, v.list.Rows[v.list.SelectedRow].(*stageEx).name, v)
  }
  return v
}

func (v *operationList) selectedName() string {
  return v.list.Rows[v.list.SelectedRow].(*stageEx).name
}

func (v *operationList) Handle(e ui.Event) View {
  switch e.ID {
  case "<Escape>", "q", "<C-c>":
    ui.Clear()
    return v.v
  case "G":
    // group the operations in the list by the current visible field and count them
    v.grouped = !v.grouped
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
    if v.Name == "executions" {
      return v.createOperationView()
    }
    if v.Name == "toolInvocations" {
      olv := NewOperationList(v.a, 4, v)
      olv.Filter = "toolInvocationId=" + v.selectedName()
      return olv
    }
    return v
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

func opStringer(name string, op *operation, field func () int) *stageEx {
  row := &stageEx{
    field: field,
    name: name,
  }
  if op != nil {
    row.fence, row.target, row.mnemonic, row.build = op.workerStart, op.target, op.mnemonic, op.build
  } else {
    row.target, row.mnemonic, row.build = "unknown", "unknown", "unknown"
  }
  return row
}

func (v operationList) renderItemized() []fmt.Stringer {
  rows := make([]fmt.Stringer, len(v.ops))
  for i := 0; i < len(v.ops); i++ {
    name := v.ops[i]
    op, ok := v.opcache.Get(name)
    o := &op
    if !ok {
      o = nil
    }
    rows[i] = opStringer(name, o, func () int { return v.field })
  }
  return rows
}

type groupResult struct {
  name string
  count int
}

func (r groupResult) String() string {
  return fmt.Sprintf("%s: %d", r.name, r.count)
}

func (v operationList) renderGrouped() []fmt.Stringer {
  buckets := make(map[string]*groupResult)
  for _, name := range v.ops {
    op, ok := v.opcache.Get(name)
    o := &op
    if !ok {
      o = nil
    }
    val := opStringer(name, o, func () int { return v.field }).label()
    r, ok := buckets[val]
    if !ok {
      r = &groupResult {
        name: val,
        count: 0,
      }
      buckets[val] = r
    }
    r.count++
  }

  results := slices.SortedStableFunc(maps.Values(buckets), func (a, b *groupResult) int {
    c := cmp.Compare(b.count, a.count)
    if v.reversed {
      c = -c
    }
    return c
  })
  rows := make([]fmt.Stringer, len(results))
  for i, r := range results {
    rows[i] = r
  }
  return rows
}

func fieldName(grouped bool, field int) string {
  fields := [...]string{"name", "target", "mnemonic", "build"}
  name := ""
  if grouped {
    name = "Grouped by "
  }
  name += fields[field]
  return name
}

func (v operationList) renderTitle() string {
  return fmt.Sprintf("%s Operations (%s) %d", v.modeTitle(), fieldName(v.grouped, v.field), len(v.ops))
}

func (v operationList) Render() []ui.Drawable {
  v.list.Title = v.renderTitle()

  var rows []fmt.Stringer
  if v.grouped {
    rows = v.renderGrouped()
  } else {
    rows = v.renderItemized()
    fence := func(e1, e2 fmt.Stringer) bool {
      stageEx1, stageEx2 := e1.(*stageEx), e2.(*stageEx)
      return (stageEx1.fence.Compare(stageEx2.fence) < 0) != v.reversed
    }
    byExec(fence).Sort(rows)
  }
  v.list.Rows = rows
  v.list.SelectedRowStyle = ui.NewStyle(ui.ColorBlack, ui.ColorWhite)
  v.list.WrapText = false
  v.list.SetRect(0, 0, 160, 30)

  return []ui.Drawable { v.list }
}
