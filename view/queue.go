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

type nodeValue string

func (nv nodeValue) String() string {
  return string(nv)
}

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

type Queue struct {
  a *client.App
  selected int
  s stats
  h int
  meter *client.List
}

func NewQueue(a *client.App, selected int) *Queue {
  _, h := ui.TerminalDimensions()
  meter := client.NewList()
  meter.SelectedRow = -1
  return &Queue {
    a: a,
    s: stats {
      profiles: make(map[string]*profileResult),
      last: time.Now(),
      mutex: &sync.Mutex{},
    },
    selected: 3,
    meter: meter,
    h: h,
  }
}

func (v *Queue) Handle(e ui.Event) View {
  switch e.ID {
  case "<Escape>", "q", "<C-c>":
    v.a.Done = true
  case "j", "<Down>":
    if v.meter.SelectedRow == -1 {
      v.selected++
      v.selected %= 4
    } else {
      v.meter.ScrollDown()
    }
  case "k", "<Up>":
    if v.meter.SelectedRow == -1 {
      v.selected += 3
      v.selected %= 4
    } else {
      v.meter.ScrollUp()
    }
  case "l", "h", "<Right>", "<Left>":
    if v.selected == 0 {
      if v.meter.SelectedRow == -1 {
        v.meter.SelectedRow = 0
      } else {
        v.meter.SelectedRow = -1
      }
    }
  case "<Enter>":
    if v.selected == 0 {
      // get the worker out of the list
      return NewWorker(v.a, v.meter.Rows[v.meter.SelectedRow].(Worker).w, v)
    } else if v.meter.SelectedRow == -1 {
      ui.Clear()
      return NewOperationList(v.a, v.selected, v)
    }
  }
  return v
}

func (v *Queue) Update() {
  s := &v.s
  c := bfpb.NewOperationQueueClient(v.a.Conn)
  var st *bfpb.BackplaneStatus
  if v.selected == 0 {
    if v.selected == 0 && s.workers != nil {
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
    for _, provision := range s.status.OperationQueue.Provisions {
      for _, of := range provision.InternalSizes {
        s.queueSum += float64(of)
      }
    }
    s.dispatchedSum += float64(s.status.DispatchedSize)
  }
  s.ticks++
}

func (v Queue) Render() []ui.Drawable {
  s := v.s
  p := widgets.NewParagraph()
  p.Text = v.a.RedisHost + "\n" + v.a.ReapiHost + "\n" + formatTime(s.last)
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
    info = renderWorkersInfo(&s, v.meter, v.h)
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
func renderWorkersInfo(s *stats, meter *client.List, h int) ui.Drawable {
  plen := len(s.workers)

  meter.SelectedRowStyle = ui.NewStyle(ui.ColorBlack, ui.ColorWhite)
  meter.SetRect(19, 4, 180, h - 6)
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
