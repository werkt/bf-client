package view

import (
  "context"
  "fmt"

  reapi "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
  bfpb "github.com/bazelbuild/bazel-buildfarm/build/buildfarm/v1test"
  ui "github.com/gizak/termui/v3"
  "github.com/gizak/termui/v3/widgets"
  "github.com/golang/protobuf/ptypes"
  "github.com/hashicorp/golang-lru/v2"
  "github.com/werkt/bf-client/client"
  "google.golang.org/genproto/googleapis/longrunning"
)

type operation struct {
  target string
  mnemonic string
  build string
}

type operationList struct {
  a *client.App
  ops []string
  opcache *lru.Cache[string, operation]
  mode int
  selected int
  v View
  field int
  queues []*client.Queue
}

func NewOperationList(a *client.App, mode int, v View) *operationList {
  opcache, _ := lru.New[string, operation](60)
  var queues []*client.Queue
  if mode != 3 {
    c := bfpb.NewOperationQueueClient(a.Conn)
    status, err := c.Status(context.Background(), &bfpb.BackplaneStatusRequest {
      InstanceName: "shard",
    })
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
    a: a,
    mode: mode,
    selected: 0,
    opcache: opcache,
    field: 0,
    v: v,
    queues: queues,
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

func (v operationList) fetchDispatched() []*client.Operation {
  var ops []*client.Operation
  var nextCursor, cursor uint64
  for nextCursor, cursor = 1, 0; len(ops) < 20 && nextCursor != 0; cursor = nextCursor {
    var opsPage []string
    var err error
    opsPage, nextCursor, err = v.a.Client.HScan(context.Background(), "DispatchedOperations", cursor, "*", 20).Result()
    if err != nil {
      panic(err)
    }
    for i, op := range opsPage {
      if i % 2 == 0 {
        ops = append(ops, &client.Operation { Name: op })
      }
    }
  }
  c := longrunning.NewOperationsClient(v.a.Conn)
  for _, op := range ops {
    if _, ok := v.a.Metadatas[op.Name]; !ok {
      v.a.Fetches++
      o, err := c.GetOperation(context.Background(), &longrunning.GetOperationRequest {
        Name: op.Name,
      })
      if err != nil {
        continue
      }
      op.Metadata = getRequestMetadata(o)
    }
    if v.a.Fetches > 10 {
      break
    }
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
  case 3: return v.fetchDispatched()
  default:
    return make([]*client.Operation, 0)
  }
}

func (v operationList) length() int64 {
  client := v.a.Client
  switch v.selected {
  case 1, 2: return v.queuesLength()
  case 3:
    length := client.HLen(context.Background(), "DispatchedOperations").Val()
    return length
  default:
    return 0
  }
}

func (v *operationList) createOperationView() View {
  if len(v.ops) > 0 {
    return NewOperation(v.a, v.ops[v.selected], v)
  }
  return v
}

func (v *operationList) Handle(e ui.Event) View {
  switch e.ID {
  case "<Escape>", "q", "<C-c>":
    ui.Clear()
    return v.v
  case "j", "<Down>":
    v.selected++
    v.selected %= 20
  case "k", "<Up>":
    v.selected += 19
    v.selected %= 20
  case "h", "<Left>":
    v.field += 3
    v.field %= 4
  case "l", "<Right>":
    v.field++
    v.field %= 4
  case "<Enter>":
    ui.Clear()
    return v.createOperationView()
  }
  return v
}

func (v *operationList) Update() {
  v.a.Fetches++
  ops := v.fetch()

  v.ops = make([]string, 0)
  for _, op := range ops {
    v.ops = append(v.ops, op.Name)
    if _, ok := v.a.Ops[op.Name]; !ok {
      m := op.Metadata
      if m != nil {
        v.opcache.Add(op.Name, operation {
          build: m.CorrelatedInvocationsId,
          target: m.TargetId,
          mnemonic: m.ActionMnemonic,
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

func (v operationList) Render() []ui.Drawable {
  ops := widgets.NewList()
  fields := [...]string{"name", "target", "build", "mnemonic"}
  ops.Title = fmt.Sprintf("Operations (%s)", fields[v.field])
  ops.Rows = make([]string, 20)
  for i := 0; i < 20 && i < len(v.ops); i++ {
    name := v.ops[i]
    var row string
    if v.field == 0 {
      row = name
    } else {
      op, _ := v.opcache.Get(name)
      if v.field == 1 {
        row = op.target
      } else if v.field == 2 {
        row = op.build
      } else if v.field == 3 {
        row = op.mnemonic
      }
    }
    ops.Rows[i] = row
  }
  ops.SelectedRowStyle = ui.NewStyle(ui.ColorBlack, ui.ColorWhite)
  ops.SelectedRow = v.selected
  ops.WrapText = false
  ops.SetRect(0, 0, 80, 30)

  return []ui.Drawable { ops }
}

func getRequestMetadata(o *longrunning.Operation) *reapi.RequestMetadata {
  m := o.Metadata
  em := &reapi.ExecuteOperationMetadata{}
  qm := &bfpb.QueuedOperationMetadata{}
  xm := &bfpb.ExecutingOperationMetadata{}
  cm := &bfpb.CompletedOperationMetadata{}
  if ptypes.Is(m, em) {
    return nil
  } else if ptypes.Is(m, qm) {
    if err := ptypes.UnmarshalAny(m, qm); err != nil {
      return nil
    } else {
      return qm.RequestMetadata
    }
  } else if ptypes.Is(m, xm) {
    if err := ptypes.UnmarshalAny(m, xm); err != nil {
      return nil
    } else {
      return xm.RequestMetadata
    }
  } else if ptypes.Is(m, cm) {
    if err := ptypes.UnmarshalAny(m, cm); err != nil {
      return nil
    } else {
      return cm.RequestMetadata
    }
  } else {
    return nil
  }
}
