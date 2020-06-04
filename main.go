package main

import (
  "context"
  "fmt"
  "log"
  "strings"
  "time"
  "container/list"
  "os"
  "sort"
  "sync"

  ui "github.com/gizak/termui"
  "github.com/gizak/termui/widgets"
  redis "github.com/redis/go-redis/v9"

  reapi "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
  "github.com/golang/protobuf/proto"
  "github.com/golang/protobuf/ptypes"
  "github.com/golang/protobuf/jsonpb"
  "google.golang.org/grpc"
  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"
  "google.golang.org/genproto/googleapis/longrunning"
  bfpb "github.com/bazelbuild/bazel-buildfarm/build/buildfarm/v1test"

  tm "github.com/nsf/termbox-go"
)

type component interface {
  open()
  close()
  handle(ui.Event)
  update() component
  render()
  done() bool
}

type view interface {
  handle(ui.Event) view
  update()
  render() []ui.Drawable
}

type app struct {
  redisHost string
  reapiHost string
  _done bool
  client *redis.ClusterClient
  conn *grpc.ClientConn
  workerConns map[string]*grpc.ClientConn
  ops map[string]*longrunning.Operation
  metadatas map[string]*reapi.RequestMetadata
  invocations map[string][]string
  fetches uint
}

type baseComponent struct {
  a *app
  v view
}

type queueView struct {
  a *app
  selected int
  s stats
}

type operationListView struct {
  a *app
  fetch func() []string
  length func() int64
  ops []string
  selected int
  v view
}

type operationView struct {
  a *app
  name string
  op *longrunning.Operation
  err error
  v view
}

type profileResult struct {
  profile *bfpb.WorkerProfileMessage
  stale int
}

type stats struct {
  workers []string
  status bfpb.BackplaneStatus
  profiles map[string]*profileResult
  last time.Time
  prequeueData list.List
  prequeueSum float64
  queueData list.List
  queueSum float64
  dispatchedData list.List
  dispatchedSum float64
  ticks float64
  mutex *sync.Mutex
}

func connect(host string) *grpc.ClientConn {
  var opts []grpc.DialOption = []grpc.DialOption{grpc.WithInsecure()}
  conn, err := grpc.Dial(host, opts...)
  if err != nil {
    panic(err)
  }
  return conn
}

func (a *app) getWorkerConn(worker string) *grpc.ClientConn {
  if a.workerConns[worker] == nil {
    // var addr [4]byte
    // var port int
    // fmt.Sscanf(worker, "%d-%d-%d-%d:%d", &addr[0], &addr[1], &addr[2], &addr[3], &port)
    // host := fmt.Sprintf("%d.%d.%d.%d:%d", addr[0], addr[1], addr[2], addr[3], port)
    a.workerConns[worker] = connect(worker)
  }
  return a.workerConns[worker]
}

func (c *baseComponent) open() {
  c.a.client = redis.NewClusterClient(&redis.ClusterOptions{
      Addrs: []string{c.a.redisHost},
      Password: "",
  })
  c.a.conn = connect(c.a.reapiHost)
}

func (c *baseComponent) close() {
  c.a.conn.Close()
}

func (c *baseComponent) handle(e ui.Event) {
  c.v = c.v.handle(e)
}

func (v *queueView) createOperationListView() view {
  client := v.a.client
  fetch := func() []string {
    switch v.selected {
    case 1:
      ops := make([]string, 0)
      for _, entry := range client.LRange(context.Background(), "{Execution}:PreQueuedOperations", 0, 20).Val() {
        ee := &bfpb.ExecuteEntry{}
        err := jsonpb.Unmarshal(strings.NewReader(entry), ee)
        if err != nil {
          ops = append(ops, err.Error())
        } else {
          ops = append(ops, ee.OperationName)
        }
      }
      return ops
    case 2:
      ops := make([]string, 0)
      for _, entry := range client.LRange(context.Background(), "{Execution}:QueuedOperations", 0, 20).Val() {
        qe := &bfpb.QueueEntry{}
        err := jsonpb.Unmarshal(strings.NewReader(entry), qe)
        if err != nil {
          ops = append(ops, err.Error())
        } else {
          ops = append(ops, qe.ExecuteEntry.OperationName)
        }
      }
      return ops
    case 3:
      ops := make([]string, 0)
      var nextCursor, cursor uint64
      for nextCursor, cursor = 1, 0; len(ops) < 20 && nextCursor != 0; cursor = nextCursor {
        var opsPage []string
        opsPage, nextCursor = client.HScan(context.Background(), "DispatchedOperations", cursor, "", 20).Val()
        for i, op := range opsPage {
          if i % 2 == 0 {
            ops = append(ops, op)
          }
        }
      }
      return ops
    default:
      return make([]string, 0)
    }
  }
  length := func() int64 {
    switch v.selected {
    case 1:
      return client.LLen(context.Background(), "{Execution}:PreQueuedOperations").Val()
    case 2:
      return client.LLen(context.Background(), "{Execution}:QueuedOperations").Val()
    case 3:
      length := client.HLen(context.Background(), "DispatchedOperations").Val()
      return length
    default:
      return 0
    }
  }
  return &operationListView {
    a: v.a,
    fetch: fetch,
    length: length,
    selected: 0,
    v: v,
  }
}

func (v *operationView) handle(e ui.Event) view {
  switch e.ID {
  case "<Escape>":
    ui.Clear()
    return v.v
  }
  return v
}

func (v *operationView) fetch() (*longrunning.Operation, error) {
  ops := longrunning.NewOperationsClient(v.a.conn)

  return ops.GetOperation(context.Background(), &longrunning.GetOperationRequest {
    Name: v.name,
  })
}

func (v *operationView) update() {
  if v.err != nil || !v.op.Done {
    v.a.fetches++
    v.op, v.err = v.fetch()
  }
}

func (v operationListView) createOperationView() view {
  return &operationView {
    a: v.a,
    name: v.ops[v.selected],
    op: &longrunning.Operation{},
    v: &v,
  }
}

func (v *operationListView) handle(e ui.Event) view {
  switch e.ID {
  case "<Escape>":
    ui.Clear()
    return v.v
  case "j":
    v.selected++
    v.selected %= 20
  case "k":
    v.selected += 19
    v.selected %= 20
  case "<Enter>":
    ui.Clear()
    return v.createOperationView()
  }
  return v
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

func (v *operationListView) update() {
  if int64(len(v.ops)) == v.length() {
    return;
  }
  v.a.fetches++
  v.ops = v.fetch()
  c := longrunning.NewOperationsClient(v.a.conn)

  for _, op := range v.ops {
    if _, ok := v.a.ops[op]; !ok {
      v.a.fetches++
      o, err := c.GetOperation(context.Background(), &longrunning.GetOperationRequest {
        Name: op,
      })
      if err != nil {
        continue
      }
      m := getRequestMetadata(o)
      v.a.ops[op] = o
      if m != nil {
        v.a.metadatas[op] = m
        var opInvocations []string
        var ok bool
        if opInvocations, ok = v.a.invocations[m.ToolInvocationId]; !ok {
          opInvocations = make([]string, 1)
          v.a.invocations[m.ToolInvocationId] = opInvocations
        }
        v.a.invocations[m.ToolInvocationId] = append(opInvocations, op)
      }
    }
    if v.a.fetches > 10 {
      break
    }
  }
}

func (v *queueView) handle(e ui.Event) view {
  switch e.ID {
  case "<Escape>", "q", "<C-c>":
    v.a._done = true
  case "j":
    v.selected++
    v.selected %= 4
  case "k":
    v.selected += 3
    v.selected %= 4
  case "<Enter>":
    ui.Clear()
    return v.createOperationListView()
  }
  return v
}

func (c *baseComponent) update() component {
  c.v.update()
  return c
}

func fetchProfile(v *queueView, worker string, conn *grpc.ClientConn, wg *sync.WaitGroup) {
  defer wg.Done()

  workerProfile := bfpb.NewWorkerProfileClient(conn)
  clientDeadline := time.Now().Add(time.Millisecond * 30)
  ctx, _ := context.WithDeadline(context.Background(), clientDeadline)
  profile, err := workerProfile.GetWorkerProfile(ctx, &bfpb.WorkerProfileRequest {})
  if err == nil {
    v.s.mutex.Lock()
    v.s.profiles[worker] = &profileResult {profile: profile, stale: 0}
    v.s.mutex.Unlock()
  } else {
    st, ok := status.FromError(err)
    if !ok || st.Code() != codes.DeadlineExceeded {
      panic(err)
    } else {
      v.s.mutex.Lock()
      result := v.s.profiles[worker]
      if result == nil {
        v.s.profiles[worker] = &profileResult {profile: &bfpb.WorkerProfileMessage{}, stale: 1}
      } else {
        result.stale++
      }
      v.s.mutex.Unlock()
    }
  }
}

func (v *queueView) update() {
  s := &v.s
  c := bfpb.NewOperationQueueClient(v.a.conn)
  var status *bfpb.BackplaneStatus
  if v.selected == 0 {
    if v.selected == 0 && s.workers != nil {
      var wg sync.WaitGroup
      for _, worker := range s.workers {
        wg.Add(1)
        go fetchProfile(v, worker, v.a.getWorkerConn(worker), &wg)
      }
      wg.Wait()
    }
  }
  status, err := c.Status(context.Background(), &bfpb.BackplaneStatusRequest {
    InstanceName: "shard",
  })
  if err == nil {
    s.status = *status
    s.workers = status.ActiveWorkers;
  } else {
    panic(err)
  }
  now := time.Now()
  if s.last.Add(time.Second / 10).Before(now) {
    if s.prequeueData.Len() > 60 {
      s.prequeueData.Remove(s.prequeueData.Back())
      s.queueData.Remove(s.queueData.Back())
      s.dispatchedData.Remove(s.dispatchedData.Back())
    }
    if s.ticks > 0 {
      s.prequeueData.PushFront(s.prequeueSum / s.ticks)
      s.queueData.PushFront(s.queueSum / s.ticks)
      s.dispatchedData.PushFront(s.dispatchedSum / s.ticks)
    } else {
      s.prequeueData.PushFront(float64(0));
      s.queueData.PushFront(float64(0));
      s.dispatchedData.PushFront(float64(0));
    }
    s.prequeueSum = 0
    s.queueSum = 0
    s.dispatchedSum = 0
    s.ticks = 0
    s.last = now
  }
  if status != nil {
    s.prequeueSum += float64(s.status.Prequeue.Size)
    for _, provision := range s.status.OperationQueue.Provisions {
      for _, of := range provision.InternalSizes {
        s.queueSum += float64(of)
      }
    }
    s.dispatchedSum += float64(s.status.DispatchedSize)
  }
  s.ticks++
}

func selectStat(selected bool, msg string) string {
  if selected {
    return "[" + msg + "](fg:black,bg:white)"
  }
  return msg
}

func (c baseComponent) render() {
  w := c.v.render()

  f := widgets.NewParagraph()
  f.Text = fmt.Sprintf("Fetches: %d", c.a.fetches)
  f.SetRect(0, 40, 20, 43)

  ui.Render(append(w, f)...)
}

func (v operationListView) render() []ui.Drawable {
  ops := widgets.NewList()
  ops.Title = "Operations"
  ops.Rows = make([]string, 20)
  for i := 0; i < 20 && i < len(v.ops); i++ {
    ops.Rows[i] = v.ops[i]
  }
  ops.SelectedRowStyle = ui.NewStyle(ui.ColorBlack, ui.ColorWhite)
  ops.SelectedRow = v.selected
  ops.WrapText = false
  ops.SetRect(0, 0, 80, 30)

  return []ui.Drawable { ops }
}

func renderExecuteOperationMetadata(em *reapi.ExecuteOperationMetadata) string {
  stage := &reapi.ExecuteOperationMetadata {
    Stage: em.Stage,
  }
  text := proto.MarshalTextString(stage) + "\n"
  text += fmt.Sprintf("action: %s\n", renderDigest(em.ActionDigest))
  return text
}

func renderDigest(d *reapi.Digest) string {
  if d == nil {
    return "nil digest"
  }
  return fmt.Sprintf("%s/%d", d.Hash, d.SizeBytes)
}

func renderInline(l int) string {
  if l > 0 {
    return " inline"
  }
  return ""
}

func renderExecutable(e bool) string {
  if e {
    return "*"
  }
  return ""
}

func formatTime(t time.Time) string {
  return t.Format("Mon Jan 2 15:04:05 -0700 MST 2006")
}

func renderExecutionMetadata(em *reapi.ExecutedActionMetadata) string {
  text := fmt.Sprintf("worker: %s\n", em.Worker)
  var qt, wst, wct, ifst, ifct, est, ect, oust, ouct time.Time
  var err error
  if qt, err = ptypes.Timestamp(em.QueuedTimestamp); err != nil {
    text += err.Error() + "\n"
    return text
  }
  if wst, err = ptypes.Timestamp(em.WorkerStartTimestamp); err != nil {
    text += err.Error() + "\n"
    return text
  }
  if wct, err = ptypes.Timestamp(em.WorkerCompletedTimestamp); err != nil {
    text += err.Error() + "\n"
    return text
  }
  if ifst, err = ptypes.Timestamp(em.InputFetchStartTimestamp); err != nil {
    text += err.Error() + "\n"
    return text
  }
  if ifct, err = ptypes.Timestamp(em.InputFetchCompletedTimestamp); err != nil {
    text += err.Error() + "\n"
    return text
  }
  if est, err = ptypes.Timestamp(em.ExecutionStartTimestamp); err != nil {
    text += err.Error() + "\n"
    return text
  }
  if ect, err = ptypes.Timestamp(em.ExecutionCompletedTimestamp); err != nil {
    text += err.Error() + "\n"
    return text
  }
  if oust, err = ptypes.Timestamp(em.OutputUploadStartTimestamp); err != nil {
    text += err.Error() + "\n"
    return text
  }
  if ouct, err = ptypes.Timestamp(em.OutputUploadCompletedTimestamp); err != nil {
    text += err.Error() + "\n"
    return text
  }
  text += fmt.Sprintf("queued at: %s\n", formatTime(qt))
  qstall := wst.Sub(qt)
  text += fmt.Sprintf("worker start: %s stalled\n", qstall)
  ifstall := ifst.Sub(wst)
  text += fmt.Sprintf("input fetch start: %s elapsed, %s stalled\n", ifst.Sub(qt).String(), ifstall.String())
  text += fmt.Sprintf("input fetch complete: %s elapsed, took %s\n", ifct.Sub(qt).String(), ifct.Sub(ifst).String())
  estall := est.Sub(ifct)
  text += fmt.Sprintf("execute start: %s elapsed, %s stalled\n", est.Sub(qt).String(), estall.String())
  text += fmt.Sprintf("execute complete: %s elapsed, took %s\n", ect.Sub(qt).String(), ect.Sub(est).String())
  oustall := oust.Sub(ect)
  text += fmt.Sprintf("output upload start: %s elapsed, %s stalled\n", oust.Sub(qt).String(), oustall.String())
  text += fmt.Sprintf("output upload complete: %s elapsed, took %s\n", ouct.Sub(qt).String(), ouct.Sub(oust).String())
  text += fmt.Sprintf(
      "worker complete: %s elapsed, %s total stalled\n",
      wct.Sub(qt).String(),
      (wct.Sub(ouct) + oustall + estall + ifstall + qstall).String())
  return text
}

func renderActionResult(ar *reapi.ActionResult) string {
  text := fmt.Sprintf("exit code: %d\n", ar.ExitCode)
  if len(ar.StderrRaw) > 0 || (ar.StdoutDigest != nil && ar.StdoutDigest.SizeBytes > 0) {
    text += fmt.Sprintf("stdout: %s%s\n", renderDigest(ar.StdoutDigest), renderInline(len(ar.StdoutRaw)))
  }
  if len(ar.StderrRaw) > 0 || (ar.StderrDigest != nil && ar.StderrDigest.SizeBytes > 0) {
    text += fmt.Sprintf("stderr: %s%s\n", renderDigest(ar.StderrDigest), renderInline(len(ar.StderrRaw)))
  }
  for _, of := range ar.OutputFiles {
    text += fmt.Sprintf("file: %s%s (%s)%s\n", renderExecutable(of.IsExecutable), of.Path, renderDigest(of.Digest), renderInline(len(of.Contents)))
  }
  for _, ofs := range ar.OutputFileSymlinks {
    text += fmt.Sprintf("symlink: %s -> %s\n", ofs.Path, ofs.Target)
  }
  for _, od := range ar.OutputDirectories {
    text += fmt.Sprintf("directory: %s (%s)\n", od.Path, renderDigest(od.TreeDigest))
  }
  text += renderExecutionMetadata(ar.ExecutionMetadata)
  return text
}

func renderExecuteResponse(er *reapi.ExecuteResponse) string {
  text := renderActionResult(er.Result)
  text += proto.MarshalTextString(er.Status)
  text += fmt.Sprintf("served from cache: %v\n", er.CachedResult)
  // server logs...
  if len(er.Message) > 0 {
    text += "message: " + er.Message + "\n"
  }
  return text
}

func (v operationView) render() []ui.Drawable {
  op := widgets.NewParagraph()
  op.Title = v.name
  op.WrapText = true
  if v.err != nil {
    op.Text = string(v.err.Error())
  } else {
    m := v.op.Metadata
    em := &reapi.ExecuteOperationMetadata{}
    qm := &bfpb.QueuedOperationMetadata{}
    xm := &bfpb.ExecutingOperationMetadata{}
    cm := &bfpb.CompletedOperationMetadata{}
    if ptypes.Is(m, em) {
      if err := ptypes.UnmarshalAny(m, em); err != nil {
        op.Text = err.Error()
      } else {
        op.Text = renderExecuteOperationMetadata(em)
      }
    } else if ptypes.Is(m, qm) {
      if err := ptypes.UnmarshalAny(m, qm); err != nil {
        op.Text = err.Error()
      } else {
        op.Text = "request metadata: " + proto.MarshalTextString(qm.RequestMetadata)
        op.Text += renderExecuteOperationMetadata(qm.ExecuteOperationMetadata)
        op.Text += fmt.Sprintf("queued operation: %s\n", renderDigest(qm.QueuedOperationDigest))
      }
    } else if ptypes.Is(m, xm) {
      if err := ptypes.UnmarshalAny(m, xm); err != nil {
        op.Text = err.Error()
      } else {
        op.Text = "request metadata: " + proto.MarshalTextString(xm.RequestMetadata)
        op.Text += renderExecuteOperationMetadata(xm.ExecuteOperationMetadata)
        t := time.Unix(xm.StartedAt / 1000, (xm.StartedAt % 1000) * 1000000)
        op.Text += fmt.Sprintf("started at: %s, running for %s\n", formatTime(t), time.Now().Sub(t).String())
        op.Text += fmt.Sprintf("executing on: %s\n", xm.ExecutingOn)
      }
    } else if ptypes.Is(m, cm) {
      if err := ptypes.UnmarshalAny(m, cm); err != nil {
        op.Text = err.Error()
      } else {
        op.Text = "request metadata: " + proto.MarshalTextString(cm.RequestMetadata)
        op.Text += renderExecuteOperationMetadata(cm.ExecuteOperationMetadata)
      }
    } else {
      op.Text = proto.MarshalTextString(v.op)
    }
    switch r := v.op.Result.(type) {
    case *longrunning.Operation_Error:
      op.Text += "error: " + proto.MarshalTextString(r.Error)
    case *longrunning.Operation_Response:
      er := &reapi.ExecuteResponse{}
      if ptypes.Is(r.Response, er) {
        if err := ptypes.UnmarshalAny(r.Response, er); err != nil {
          op.Text += err.Error()
        } else {
          op.Text += renderExecuteResponse(er)
        }
      }
    }
  }
  op.SetRect(0, 0, 120, 60)

  return []ui.Drawable { op }
}

type By func(w1, w2 *string) bool

func (by By) Sort(workers []string) {
  ws := &workerSorter{
    workers: workers,
    by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
  }
  sort.Sort(ws)
}

type workerSorter struct {
  workers []string
  by      func(w1, w2 *string) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *workerSorter) Len() int {
  return len(s.workers)
}

// Swap is part of sort.Interface.
func (s *workerSorter) Swap(i, j int) {
  s.workers[i], s.workers[j] = s.workers[j], s.workers[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *workerSorter) Less(i, j int) bool {
  return s.by(&s.workers[i], &s.workers[j])
}

func (v queueView) render() []ui.Drawable {
  s := v.s
  p := widgets.NewParagraph()
  p.Text = v.a.redisHost + "\n" + v.a.reapiHost + "\n" + formatTime(s.last)
  p.SetRect(0, 0, 80, 5)

  stats := widgets.NewParagraph()
  stats.Text = selectStat(v.selected == 0, fmt.Sprintf("Workers: %v\n", len(s.workers)))
  stats.Text += selectStat(v.selected == 1, fmt.Sprintf("Prequeue: %v\n", s.status.Prequeue.Size))
  queueSize := int64(0)
  for _, provision := range s.status.OperationQueue.Provisions {
    for _, of := range provision.InternalSizes {
      queueSize += of
    }
  }
  stats.Text += selectStat(v.selected == 2, fmt.Sprintf("Queue: %v\n", queueSize))
  stats.Text += selectStat(v.selected == 3, fmt.Sprintf("Dispatched: %v\n", s.status.DispatchedSize))
  stats.SetRect(0, 4, 20, 10)

  var info ui.Drawable
  if v.selected == 0 {
    plen := len(s.workers)

    meter := widgets.NewList()
    meter.SetRect(19, 4, 180, plen + 6)
    meter.Title = "Workers";

    wl := 0
    for _, worker := range s.workers {
      twl := len(worker)
      if twl > wl {
        wl = twl
      }
    }
    exec_used := func(worker *string) int {
      profile := s.profiles[*worker].profile
      for _, stage := range profile.Stages {
        if stage.Name == "ExecuteActionStage" {
          return int(stage.SlotsUsed)
        }
      }
      return 0;
    }
    exec_avail := func(worker *string) int {
      profile := s.profiles[*worker].profile
      for _, stage := range profile.Stages {
        if stage.Name == "ExecuteActionStage" {
          return int(stage.SlotsConfigured)
        }
      }
      return 0;
    }
    exec := func(w1, w2 *string) bool {
      if exec_used(w1) == exec_used(w2) {
        if exec_avail(w1) == exec_avail(w2) {
          return *w1 < *w2
        }
        return exec_avail(w1) < exec_avail(w2)
      }
      return exec_used(w1) < exec_used(w2)
    }
    By(exec).Sort(s.workers)

    n := 0
    rows := make([]string, plen)
    for _, worker := range s.workers {
      result := s.profiles[worker]
      profile := result.profile
      var input_fetch_used, input_fetch_slots int
      var execute_action_used, execute_action_slots int
      for _, stage := range profile.Stages {
        if stage.Name == "InputFetchStage" {
          input_fetch_used = int(stage.SlotsUsed)
          input_fetch_slots = int(stage.SlotsConfigured)
        } else if stage.Name == "ExecuteActionStage" {
          execute_action_used = int(stage.SlotsUsed)
          execute_action_slots = int(stage.SlotsConfigured)
        }
      }
      row := strings.Repeat(" ", wl - len(worker))
      row += worker + ": [" + strings.Repeat(" ", input_fetch_slots - input_fetch_used)
      row += strings.Repeat("#", input_fetch_used) + "]("
      if input_fetch_used == input_fetch_slots {
        row += "fg:black,mod:dim,bg:blue"
      } else {
        row += "fg:blue"
      }
      row += ")[" + strings.Repeat("#", execute_action_used) + "]("
      if execute_action_used == execute_action_slots {
        row += "fg:black,mod:dim,bg:red"
      } else {
        row += "fg:red"
      }
      row += fmt.Sprintf(") %d/%d", execute_action_used, execute_action_slots)
      if result.stale > 0 {
        row += " stale"
      }
      rows[n] = row
      n++
    }
    meter.Rows = rows

    info = meter
  } else {
    plot := widgets.NewPlot()
    plot.Data = make([][]float64, 1)
    var container list.List
    switch v.selected {
    case 1:
      plot.Title = "Prequeue"
      plot.LineColors[0] = ui.ColorRed
      container = s.prequeueData
    case 2:
      plot.Title = "Queue"
      plot.LineColors[0] = ui.ColorYellow
      container = s.queueData
    case 3:
      plot.Title = "Dispatched"
      plot.LineColors[0] = ui.ColorCyan
      container = s.dispatchedData
    }
    plot.Data[0] = make([]float64, 60)
    n := 0
    for e := container.Front(); n < 60 && n < container.Len(); n++ {
      plot.Data[0][59 - n] = e.Value.(float64)
      e = e.Next()
    }
    for ; n < 60; n++ {
      plot.Data[0][59 - n] = float64(0)
    }
    plot.SetRect(19, 4, 90, 30)
    plot.AxesColor = ui.ColorWhite
    plot.Marker = widgets.MarkerBraille
    plot.PlotType = widgets.ScatterPlot

    info = plot
  }

  return []ui.Drawable { p, stats, info }
}

func (c baseComponent) done() bool {
  return c.a._done
}

func main() {
  if err := ui.Init(); err != nil {
    log.Fatalf("failed to initialize termui: %v", err)
  }
  tm.SetInputMode(tm.InputEsc)
  defer ui.Close()

  a := &app {
    redisHost: os.Args[1] + ":6379",
    reapiHost: os.Args[2],
    _done: false,
    ops: make(map[string]*longrunning.Operation),
    metadatas: make(map[string]*reapi.RequestMetadata),
    invocations: make(map[string][]string),
    workerConns: make(map[string]*grpc.ClientConn),
  }
  var c component = &baseComponent {
    a: a,
    v: &queueView {
      a: a,
      s: stats {
        profiles: make(map[string]*profileResult),
        last: time.Now(),
        mutex: &sync.Mutex{},
      },
      selected: 3,
    },
  }

  c.open()

  uiEvents := ui.PollEvents()
  ticker := time.NewTicker(time.Millisecond / 60).C
  for !c.done() {
    select {
    case e := <-uiEvents:
      c.handle(e)
    case <-ticker:
      a.fetches = 0
      c = c.update()
      c.render()
    }
  }

  c.close()
}
