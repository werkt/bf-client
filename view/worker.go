package view

import (
  "context"
  "fmt"

  bfpb "github.com/buildfarm/buildfarm/build/buildfarm/v1test"
  ui "github.com/gizak/termui/v3"
  "github.com/werkt/bf-client/client"
  "google.golang.org/genproto/googleapis/longrunning"
  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"
)

type worker struct {
  a *client.App
  v View
  w string
  profile *bfpb.WorkerProfileMessage
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
  match := NewStageList()
  inputFetch := NewStageList()
  execute := NewStageList()
  reportResult := NewStageList()
  execute.SelectedRow = 0
  return &worker {
    a: a,
    v: v,
    w: w,
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

func (v worker) Render() []ui.Drawable {
  v.match.Title = selectedTitle(v.match.SelectedRow != -1, "Match")
  v.match.SetRect(0, 0, 80, 3)

  v.reportResult.Title = selectedTitle(v.reportResult.SelectedRow != -1, "ReportResult")

  for _, stage := range v.profile.Stages {
    switch stage.Name {
    case "MatchStage":
      populateRows(v.match, stage.OperationNames)
    case "InputFetchStage":
      v.inputFetch.Title = fmt.Sprintf(selectedTitle(v.inputFetch.SelectedRow != -1, "InputFetch") + " %d/%d", stage.SlotsUsed, stage.SlotsConfigured)
      populateRows(v.inputFetch, stage.OperationNames)
      v.inputFetch.SetRect(0, 3, 80, 5 + len(v.inputFetch.Rows))
    case "ExecuteActionStage":
      v.execute.Title = fmt.Sprintf(selectedTitle(v.execute.SelectedRow != -1, "Execute") + " %d/%d", stage.SlotsUsed, stage.SlotsConfigured)
      // TODO get some time spent doing this
      populateRows(v.execute, stage.OperationNames)
    case "ReportResultStage":
      populateRows(v.reportResult, stage.OperationNames)
    }
  }
  i := 5 + len(v.inputFetch.Rows)
  v.execute.SetRect(0, i, 80, i + 2 + len(v.execute.Rows))
  i += 2 + len(v.execute.Rows)
  v.reportResult.SetRect(0, i, 80, i + 3)

  return []ui.Drawable { v.match, v.inputFetch, v.execute, v.reportResult }
}

func (v *worker) Update() {
  conn := v.a.GetWorkerConn(v.w, v.a.CA)
  workerProfile := bfpb.NewWorkerProfileClient(conn)
  profile, err := workerProfile.GetWorkerProfile(context.Background(), &bfpb.WorkerProfileRequest {})
  if err == nil {
    v.profile = profile
  }
}
