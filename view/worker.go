package view

import (
  "context"
  "errors"
  "fmt"

  reapi "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
  bfpb "github.com/buildfarm/buildfarm/build/buildfarm/v1test"
  ui "github.com/gizak/termui/v3"
  "github.com/gizak/termui/v3/widgets"
  "github.com/werkt/bf-client/client"
  "google.golang.org/genproto/googleapis/longrunning"
  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"
  "github.com/dustin/go-humanize"
)

type worker struct {
  a *client.App
  v View
  w string
  profile *bfpb.WorkerProfileMessage
  title *widgets.Paragraph
  match *client.List
  inputFetch *client.List
  execute *client.List
  reportResult *client.List
}

func NewStageList() *client.List {
  list := client.NewList()
  list.SelectedRow = -1
  list.SelectedRowStyle = ui.NewStyle(ui.ColorBlack, ui.ColorWhite)
  return list
}

func NewWorker(a *client.App, w string, v View) *worker {
  title := widgets.NewParagraph()
  match := NewStageList()
  inputFetch := NewStageList()
  execute := NewStageList()
  reportResult := NewStageList()
  execute.SelectedRow = 0
  return &worker {
    a: a,
    v: v,
    w: w,
    title: title,
    match: match,
    inputFetch: inputFetch,
    execute: execute,
    reportResult: reportResult,
  }
}

func (v *worker) currentOperationName() string {
  lists := []*client.List { v.match, v.inputFetch, v.execute, v.reportResult }
  for _, list := range lists {
    if list.SelectedRow != -1 && list.SelectedRow < len(list.Rows) {
      return list.Rows[list.SelectedRow].String()
    }
  }
  return ""
}

func (v *worker) cancelOperation() {
  name := v.currentOperationName()
  if len(name) != 0 {
    ops := longrunning.NewOperationsClient(v.a.Conn)
    _, err := ops.CancelOperation(context.Background(), &longrunning.CancelOperationRequest {
      Name: name,
    })
    if err != nil {
      st, ok := status.FromError(err)
      if !ok || st.Code() != codes.Unknown {
        panic(err)
      }
      // buildfarm spits out an unknown for already-done
    }
  }
}

func (v *worker) selectedList() *client.List {
  lists := []*client.List { v.match, v.inputFetch, v.execute, v.reportResult }
  for _, list := range lists {
    if list.SelectedRow != -1 {
      return list
    }
  }
  return nil
}

func (v *worker) Handle(e ui.Event) View {
  switch e.ID {
  case "<Escape>", "q", "<C-c>":
    ui.Clear()
    return v.v
  case "X":
    v.cancelOperation()
  case "<Enter>":
    return NewOperation(v.a, v.currentOperationName(), v)
  case "j", "<Down>":
    v.selectedList().ScrollDown()
  case "k", "<Up>":
    v.selectedList().ScrollUp()
  case "<Tab>":
    if v.match.SelectedRow != -1 {
      v.match.SelectedRow = -1
      v.inputFetch.SelectedRow = 0
    } else if v.inputFetch.SelectedRow != -1 {
      v.inputFetch.SelectedRow = -1
      v.execute.SelectedRow = 0
    } else if v.execute.SelectedRow != -1 {
      v.execute.SelectedRow = -1
      v.reportResult.SelectedRow = 0
    } else {
      v.reportResult.SelectedRow = -1
      v.match.SelectedRow = 0
    }
  }
  return v
}

type nodeType string

func (n nodeType) String() string {
  return string(n)
}

func selectedTitle(s bool, t string) string {
  if s {
    return ">" + t + "<"
  }
  return t
}

func populateRows(l *client.List, r []string) {
  rows := make([]fmt.Stringer, len(r))
  for i, row := range r {
    rows[i] = nodeType(row)
  }
  l.Rows = rows
}

func opMatchesStage(op *longrunning.Operation, stage *bfpb.StageInformation) (bool, error) {
  em, err := client.ExecuteOperationMetadata(op)
  if err != nil {
    return false, err
  }
  switch stage.Name {
  case "MatchStage", "InputFetchStage": 
    return em.Stage == reapi.ExecutionStage_QUEUED, nil
  case "ExecuteActionStage":
    return em.Stage == reapi.ExecutionStage_EXECUTING, nil
  case "ReportResultStage":
    return em.Stage == reapi.ExecutionStage_COMPLETED, nil
  }
  return false, errors.New("Unknown stage: " + stage.Name)
}

func (v worker) Render() []ui.Drawable {
  v.title.Text = fmt.Sprintf(
      "%s CAS Count: %d Size: %s (%d%%) Unref: %d%%",
      v.w, v.profile.CasEntryCount, humanize.Bytes(uint64(v.profile.CasSize)),
      int((float64(v.profile.CasSize) / float64(v.profile.CasMaxSize)) * 100),
      int((float64(v.profile.CasUnreferencedEntryCount) / float64(v.profile.CasEntryCount)) * 100))
  v.title.Border = false
  v.title.SetRect(0, -1, 80, 2)
  v.match.Title = selectedTitle(v.match.SelectedRow != -1, "Match")

  v.reportResult.Title = selectedTitle(v.reportResult.SelectedRow != -1, "ReportResult")

  // need some expander logic

  c := longrunning.NewOperationsClient(v.a.Conn)
  for _, stage := range v.profile.Stages {
    // for all operations in stages
    for _, name := range stage.OperationNames {
      // if operation not in cache, fetch it
      fetch := false
      if op, ok := v.a.Ops[name]; !ok {
        fetch = true
      // if stage is not the operation current stage, fetch it
      } else if op != nil {
        match, err := opMatchesStage(op, stage)
        if err != nil {
          panic(err)
        }
        fetch = !match
      }
      if fetch {
        o, _ := c.GetOperation(context.Background(), &longrunning.GetOperationRequest {
          Name: name,
        })
        v.a.Ops[name] = o
      }
    }

    switch stage.Name {
    case "MatchStage":
      populateRows(v.match, stage.OperationNames)
    case "InputFetchStage":
      v.inputFetch.Title = fmt.Sprintf(selectedTitle(v.inputFetch.SelectedRow != -1, "InputFetch") + " %d/%d", stage.SlotsUsed, stage.SlotsConfigured)
      populateRows(v.inputFetch, stage.OperationNames)
    case "ExecuteActionStage":
      v.execute.Title = fmt.Sprintf(selectedTitle(v.execute.SelectedRow != -1, "Execute") + " %d/%d", stage.SlotsUsed, stage.SlotsConfigured)
      // TODO get some time spent doing this
      populateRows(v.execute, stage.OperationNames)
    case "ReportResultStage":
      v.reportResult.Title = fmt.Sprintf(selectedTitle(v.reportResult.SelectedRow != -1, "Report Result") + " %d/%d", stage.SlotsUsed, stage.SlotsConfigured)
      populateRows(v.reportResult, stage.OperationNames)
    }
  }

  row := 1
  v.match.SetRect(0, row, 80, row + 3)
  row += 3
  inputFetchRows := len(v.inputFetch.Rows)
  inputFetchHeight := Min(inputFetchRows, 5)
  v.inputFetch.SetRect(0, row, 80, row + 2 + inputFetchHeight)
  row += 2 + inputFetchHeight
  executeRows := len(v.execute.Rows)
  executeHeight := Min(executeRows, 5)
  v.execute.SetRect(0, row, 80, row + 2 + executeHeight)
  row += 2 + executeHeight
  reportResultRows := len(v.reportResult.Rows)
  reportResultHeight := Min(reportResultRows, 5)
  v.reportResult.SetRect(0, row, 80, row + 2 + reportResultHeight)

  return []ui.Drawable { v.title, v.match, v.inputFetch, v.execute, v.reportResult }
}

func (v *worker) Update() {
  conn := v.a.GetWorkerConn(v.w, v.a.CA)
  workerProfile := bfpb.NewWorkerProfileClient(conn)
  profile, err := workerProfile.GetWorkerProfile(context.Background(), &bfpb.WorkerProfileRequest {})
  if err == nil {
    v.profile = profile
  }
}
