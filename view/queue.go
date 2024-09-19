package view

import (
  "container/list"
  "context"
  "fmt"
  "strings"
  "sort"
  "sync"
  "time"
  "github.com/gizak/termui/v3/widgets"
  ui "github.com/gizak/termui/v3"
  bfpb "github.com/bazelbuild/bazel-buildfarm/build/buildfarm/v1test"
  "github.com/werkt/bf-client/client"
  "google.golang.org/grpc"
  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"
)

type profileResult struct {
  profile *bfpb.WorkerProfileMessage
  stale int
  message string
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

type numValue struct {
  fmt string
  value int
  mode int
  parent *numValue
}

func (nv numValue) String() string {
  return fmt.Sprintf(nv.fmt, nv.value)
}

type Queue struct {
  a *client.App
  focused bool
  s stats
  h int
  meter *client.List
  stats *client.Tree
  workers numValue
  prequeueNode client.TreeNode
  prequeue numValue
  queueNode client.TreeNode
  queue numValue
  dispatched numValue
}

func statNode(nv *numValue) *client.TreeNode {
  return &client.TreeNode{ Value: nv, }
}

func NewQueue(a *client.App, selected int) *Queue {
  _, h := ui.TerminalDimensions()
  meter := client.NewList()
  meter.SelectedRow = -1
  q := &Queue {
    a: a,
    s: stats {
      profiles: make(map[string]*profileResult),
      last: time.Now(),
      mutex: &sync.Mutex{},
    },
    meter: meter,
    h: h,
    stats: client.NewTree(),
    workers: numValue{ fmt: "Workers: %v", },
    prequeue: numValue{ fmt: "Prequeue: %v", mode: 1 },
    queue: numValue{ fmt: "Queue: %v", mode: 2 },
    dispatched: numValue{ fmt: "Dispatched: %v", mode: 3 },
  }
  q.stats.Focused = true
  q.stats.SelectedRow = selected
  q.stats.SelectedRowStyle = ui.NewStyle(ui.ColorBlack, ui.ColorWhite)
  q.queueNode.Value = &q.queue
  q.prequeueNode.Value = &q.prequeue
  q.stats.SetNodes([]*client.TreeNode{
    statNode(&q.workers),
    &q.prequeueNode,
    &q.queueNode,
    statNode(&q.dispatched),
  })
  return q
}

func (v *Queue) Handle(e ui.Event) View {
  switch e.ID {
  case "<Escape>", "q", "<C-c>":
    v.a.Done = true
  case "j", "<Down>":
    if v.stats.Focused {
      v.stats.ScrollDown()
    }
    if v.meter.SelectedRow != -1 {
      v.meter.ScrollDown()
    }
  case "k", "<Up>":
    if v.stats.Focused {
      v.stats.ScrollUp()
    } else {
      v.meter.ScrollUp()
    }
  case "h", "<Left>":
    if (v.stats.Focused) {
      v.stats.Collapse()
      ui.Clear()
    } else {
      v.stats.Focused = true
    }
    if v.stats.SelectedRow == 0 {
      if v.meter.SelectedRow == -1 {
        v.meter.SelectedRow = 0
      } else {
        v.meter.SelectedRow = -1
      }
    }
  case "l", "<Right>":
    if v.stats.SelectedRow != 0 {
      if (v.stats.Focused) {
        v.stats.Expand()
        ui.Clear()
      }
    } else {
      v.stats.Focused = v.meter.SelectedRow != -1
      if v.meter.SelectedRow == -1 {
        v.meter.SelectedRow = 0
      } else {
        v.meter.SelectedRow = -1
      }
    }
  case "<Enter>":
    if v.meter.SelectedRow >= 0 {
      // get the worker out of the list
      return NewWorker(v.a, v.meter.Rows[v.meter.SelectedRow].(Worker).w, v)
    } else {
      ui.Clear()
      return NewOperationList(v.a, v.stats.SelectedNode().Value.(*numValue).mode, v)
    }
  /*
  case "S":
    return NewServerTest(v.a, v)
  */
  }
  return v
}

func updateProvisionNode(node *client.TreeNode, provision *bfpb.QueueStatus) {
  size := int64(0)
  n := len(provision.InternalSizes)
  if len(node.Nodes) != n {
    node.Nodes = make([]*client.TreeNode, n)
    for i := 0; i < n; i++ {
      node.Nodes[i] = statNode(&numValue{ fmt: "%v", parent: node.Value.(*numValue) })
    }
  }
  for i, of := range provision.InternalSizes {
    nodeNumValue := node.Nodes[i].Value.(*numValue)
    nodeNumValue.value = int(of)
    size += of
  }
  nodeNumValue := node.Value.(*numValue)
  nodeNumValue.value = int(size)
}

func (v *Queue) updateProvisionNodes(provisions []*bfpb.QueueStatus) {
  if len(v.queueNode.Nodes) != len(provisions) {
    for _, provision := range provisions {
      node := &client.TreeNode{
        Value: &numValue{ fmt: provision.Name + ": %v", parent: &v.queue },
        Nodes: make([]*client.TreeNode, 0),
      }
      v.queueNode.Nodes = append(v.queueNode.Nodes, node)
    }
  }
  for i, provision := range provisions {
    updateProvisionNode(v.queueNode.Nodes[i], provision)
  }
}

func (v *Queue) Update() {
  s := &v.s
  c := bfpb.NewOperationQueueClient(v.a.Conn)
  var st *bfpb.BackplaneStatus
  if v.stats.SelectedRow == 0 {
    if s.workers != nil {
      var wg sync.WaitGroup
      for _, worker := range s.workers {
        wg.Add(1)
        go fetchProfile(v, worker, v.a.GetWorkerConn(worker, v.a.CA), &wg)
      }
      wg.Wait()
    }
  }
  st, err := c.Status(context.Background(), &bfpb.BackplaneStatusRequest {
    InstanceName: "shard",
  })
  if err == nil {
    s.status = *st
    s.workers = st.ActiveWorkers;
  } else {
    st, ok := status.FromError(err)
    if !ok || (st.Code() != codes.Unknown && st.Code() != codes.Unavailable) {
      panic(err)
    }
    return
  }
  v.workers.value = len(s.workers)
  v.prequeue.value = int(s.status.Prequeue.Size)
  v.queue.value = int(s.status.OperationQueue.Size)
  v.dispatched.value = int(s.status.DispatchedSize)
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
  if st != nil {
    s.prequeueSum += float64(s.status.Prequeue.Size)
    s.queueSum = float64(s.status.OperationQueue.Size)
    s.dispatchedSum += float64(s.status.DispatchedSize)
  }
  updateProvisionNode(&v.prequeueNode, s.status.Prequeue)
  v.updateProvisionNodes(s.status.OperationQueue.Provisions)
  s.ticks++
}

type dims struct {
  width int
  height int
}

func nodeDimensions(node *client.TreeNode, level int) dims {
  d := dims{ len(node.Value.String()) + level * 2, 1 }
  if node.Expanded {
    for _, n := range node.Nodes {
      nd := nodeDimensions(n, level+1)
      if nd.width > d.width {
        d.width = nd.width
      }
      d.height += nd.height
    }
  }
  return d
}

func treeDimensions(t *client.Tree) dims {
  d := dims{ 0, 0 }
  t.Walk(func(n *client.TreeNode) int {
    nd := nodeDimensions(n, 0)
    if nd.width > d.width {
      d.width = nd.width
    }
    d.height += nd.height
    return -1
  })
  return d
}

func (v Queue) Render() []ui.Drawable {
  s := v.s
  p := widgets.NewParagraph()
  p.Text = v.a.RedisHost + "\n" + v.a.ReapiHost + "\n" + formatTime(s.last)
  p.SetRect(0, 0, 80, 5)

  d := treeDimensions(v.stats)
  d.width += 4
  d.height += 6
  v.stats.SetRect(0, 4, d.width, d.height)

  var info ui.Drawable
  if v.stats.SelectedRow == 0 {
    info = renderWorkersInfo(&s, v.meter, d.width, v.h)
  } else {
    plot := widgets.NewPlot()
    plot.Data = make([][]float64, 1)
    var chartValue *numValue
    for chartValue, _ = v.stats.SelectedNode().Value.(*numValue);
        chartValue.parent != nil;
        chartValue = chartValue.parent { }
    mode := chartValue.mode
    var container list.List
    switch mode {
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
    plot.SetRect(d.width, 4, d.width + 71, 30)
    plot.AxesColor = ui.ColorWhite
    plot.Marker = widgets.MarkerBraille
    plot.PlotType = widgets.ScatterPlot

    info = plot
  }

  return []ui.Drawable{ p, v.stats, info }
}

func fetchProfile(v *Queue, worker string, conn *grpc.ClientConn, wg *sync.WaitGroup) {
  defer wg.Done()

  workerProfile := bfpb.NewWorkerProfileClient(conn)
  clientDeadline := time.Now().Add(time.Millisecond * 30)
  ctx, _ := context.WithDeadline(context.Background(), clientDeadline)
  profile, err := workerProfile.GetWorkerProfile(ctx, &bfpb.WorkerProfileRequest {})
  if err == nil {
    v.s.mutex.Lock()
    v.s.profiles[worker] = &profileResult {profile: profile, stale: 0, message: ""}
    v.s.mutex.Unlock()
  } else {
    st, ok := status.FromError(err)
    v.s.mutex.Lock()
    if !ok || st.Code() != codes.DeadlineExceeded {
      result := v.s.profiles[worker]
      if result == nil {
        v.s.profiles[worker] = &profileResult {profile: &bfpb.WorkerProfileMessage{}, stale: 1, message: st.String()}
      } else {
        result.stale++
        result.message = st.String()
      }
    } else {
      result := v.s.profiles[worker]
      if result == nil {
        v.s.profiles[worker] = &profileResult {profile: &bfpb.WorkerProfileMessage{}, stale: 1, message: ""}
      } else {
        result.stale++
      }
    }
    v.s.mutex.Unlock()
  }
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
  return !s.by(&s.workers[i], &s.workers[j])
}

type Worker struct {
  w string
  row string
}

func (w Worker) String() string {
  return w.row
}

// List needs work on draw, flip for only background, etc
func renderWorkersInfo(s *stats, meter *client.List, x int, h int) ui.Drawable {
  plen := len(s.workers)

  meter.SelectedRowStyle = ui.NewStyle(ui.ColorBlack, ui.ColorWhite)
  meter.SetRect(x, 4, x + 161, h - 6)
  meter.Title = "Workers";

  wl := 0
  for _, worker := range s.workers {
    twl := len(worker)
    if twl > wl {
      wl = twl
    }
  }
  exec_used := func(worker *string) int {
    profileResult := s.profiles[*worker]
    if profileResult == nil {
      return -1
    }
    profile := profileResult.profile
    for _, stage := range profile.Stages {
      if stage.Name == "ExecuteActionStage" {
        return int(stage.SlotsUsed)
      }
    }
    return 0;
  }
  exec_avail := func(worker *string) int {
    profileResult := s.profiles[*worker]
    if profileResult == nil {
      return -1
    }
    profile := profileResult.profile
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
  rows := make([]fmt.Stringer, plen)
  for _, w := range s.workers {
    rows[n] = renderWorkerRow(s.profiles[w], w, wl)
    n++
  }
  meter.Rows = rows

  return meter
}

func renderWorkerRow(r *profileResult, w string, wl int) Worker {
  var profile *bfpb.WorkerProfileMessage
  if r == nil {
    r = &profileResult {profile: &bfpb.WorkerProfileMessage{}, stale: 1, message: "uninitialized"}
    profile = r.profile
  } else {
    profile = r.profile
  }
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
  row := strings.Repeat(" ", wl - len(w))
  row += w + ": [" + strings.Repeat(" ", input_fetch_slots - input_fetch_used)
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
  if r.stale > 0 {
    row += " stale"
    if len(r.message) > 0 {
      row += ", " + r.message
    }
  }
  return Worker {
    w: w,
    row: row,
  }
}

func formatTime(t time.Time) string {
  return t.Format("Mon Jan 2 15:04:05 -0700 MST 2006")
}

func selectStat(selected bool, msg string) string {
  if selected {
    return "[" + msg + "](fg:black,bg:white)"
  }
  return msg
}
