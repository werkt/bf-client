package view

import (
  "container/list"
  "context"
  "fmt"
  "regexp"
  "time"
  bfpb "github.com/bazelbuild/bazel-buildfarm/build/buildfarm/v1test"
  reapi "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
  ui "github.com/gizak/termui/v3"
  "github.com/gizak/termui/v3/widgets"
  "github.com/golang/protobuf/proto"
  "github.com/golang/protobuf/ptypes"
  "github.com/werkt/bf-client/client"
  "google.golang.org/genproto/googleapis/longrunning"
)

type operationView struct {
  a *client.App
  name string
  op *longrunning.Operation
  err error
  v View
  selection int // enum of correlated, invocation, action
  selectableFields int
  selectionActions []func(*operationView) View
}

func NewOperation(a *client.App, name string, v View) *operationView {
  return &operationView {
    a: a,
    name: name,
    op: &longrunning.Operation{},
    v: v,
    selection: 0,
    selectableFields: 0,
  }
}

func (v *operationView) Handle(e ui.Event) View {
  switch e.ID {
  case "<Escape>", "q", "<C-c>":
    ui.Clear()
    return v.v
  case "j", "<Down>":
    v.selection++
    if v.selectableFields > 0 {
      v.selection %= v.selectableFields
    } else {
      v.selection = 0
    }
  case "k", "<Up>":
    v.selection--
    if v.selectableFields > 0 {
      v.selection %= v.selectableFields
    } else {
      v.selection = 0
    }
  case "<Enter>":
    if v.selectableFields > 0 {
      return v.selectionActions[v.selection](v)
    }
  }
  return v
}

func (v *operationView) fetch() (*longrunning.Operation, error) {
  ops := longrunning.NewOperationsClient(v.a.Conn)

  return ops.GetOperation(context.Background(), &longrunning.GetOperationRequest {
    Name: v.name,
  })
}

func (v *operationView) Update() {
  if v.err != nil || !v.op.Done {
    v.a.Fetches++
    v.op, v.err = v.fetch()
  }
}

func (v *operationView) Render() []ui.Drawable {
  op := widgets.NewParagraph()
  op.Title = v.name
  op.WrapText = true
  if v.err != nil {
    op.Text = string(v.err.Error())
    v.selectableFields = 0
  } else {
    v.renderOperation(op)
  }
  op.SetRect(0, 0, 120, 60)

  return []ui.Drawable { op }
}

func (v *operationView) renderOperation(op *widgets.Paragraph) {
  m := v.op.Metadata
  em := &reapi.ExecuteOperationMetadata{}
  qm := &bfpb.QueuedOperationMetadata{}
  xm := &bfpb.ExecutingOperationMetadata{}
  cm := &bfpb.CompletedOperationMetadata{}
  if ptypes.Is(m, em) {
    if err := ptypes.UnmarshalAny(m, em); err != nil {
      op.Text = err.Error()
      v.selectableFields = 0
    } else {
      m, ok := v.a.Metadatas[v.name]
      op.Text = ""
      if ok {
        op.Text += renderRequestMetadata(m, v.selection)
        v.selectableFields = 3
      } else {
        v.selectableFields = 1
      }
      op.Text += renderExecuteOperationMetadata(em, v.selection)
      v.selectionActions = make([]func(*operationView) View, v.selectableFields)
      v.selectionActions[0] = func(ov *operationView) View { return NewAction(v.a, em.ActionDigest, ov) }
    }
  } else if ptypes.Is(m, qm) {
    if err := ptypes.UnmarshalAny(m, qm); err != nil {
      op.Text = err.Error()
      v.selectableFields = 0
    } else {
      op.Text = renderRequestMetadata(qm.RequestMetadata, v.selection)
      op.Text += renderExecuteOperationMetadata(qm.ExecuteOperationMetadata, v.selection)
      op.Text += fmt.Sprintf("queued operation: %s\n", renderDigest(qm.QueuedOperationDigest, v.selection == 3))
      v.selectableFields = 4
      v.selectionActions = make([]func(*operationView) View, v.selectableFields)
      v.selectionActions[0] = createCorrelatedView
      v.selectionActions[1] = createInvocationView
      v.selectionActions[2] = func(ov *operationView) View { return NewAction(v.a, qm.ExecuteOperationMetadata.ActionDigest, ov) }
      v.selectionActions[3] = func (ov *operationView) View {
        return createQueuedOperationView(ov, qm.QueuedOperationDigest)
      }
    }
  } else if ptypes.Is(m, xm) {
    if err := ptypes.UnmarshalAny(m, xm); err != nil {
      op.Text = err.Error()
      v.selectableFields = 0
    } else {
      op.Text = renderRequestMetadata(xm.RequestMetadata, v.selection)
      op.Text += renderExecuteOperationMetadata(xm.ExecuteOperationMetadata, v.selection)
      t := time.Unix(xm.StartedAt / 1000, (xm.StartedAt % 1000) * 1000000)
      op.Text += fmt.Sprintf("started at: %s, running for %s\n", formatTime(t), time.Now().Sub(t).String())
      op.Text += fmt.Sprintf("executing on: %s\n", boldIfSelected(xm.ExecutingOn, v.selection == 3))
      v.selectableFields = 4
      v.selectionActions = make([]func(*operationView) View, v.selectableFields)
      v.selectionActions[0] = createCorrelatedView
      v.selectionActions[1] = createInvocationView
      v.selectionActions[2] = func(ov *operationView) View { return NewAction(v.a, xm.ExecuteOperationMetadata.ActionDigest, ov); }
      v.selectionActions[3] = func (ov *operationView) View {
        return createWorkerView(ov, xm.ExecutingOn)
      }
    }
  } else if ptypes.Is(m, cm) {
    if err := ptypes.UnmarshalAny(m, cm); err != nil {
      op.Text = err.Error()
      v.selectableFields = 0
    } else {
      op.Text = renderRequestMetadata(cm.RequestMetadata, v.selection)
      op.Text += renderExecuteOperationMetadata(cm.ExecuteOperationMetadata, v.selection)
      v.selectableFields = 3
      v.selectionActions = make([]func(*operationView) View, v.selectableFields)
      v.selectionActions[0] = createCorrelatedView
      v.selectionActions[1] = createInvocationView
      v.selectionActions[2] = func(ov *operationView) View { return NewAction(v.a, cm.ExecuteOperationMetadata.ActionDigest, ov); }
    }
  } else {
    op.Text = proto.MarshalTextString(v.op)
    v.selectableFields = 0
  }

  switch r := v.op.Result.(type) {
  case *longrunning.Operation_Error:
    op.Text += "error: " + proto.MarshalTextString(r.Error)
  case *longrunning.Operation_Response:
    er := &reapi.ExecuteResponse{}
    if ptypes.Is(r.Response, er) {
      if err := ptypes.UnmarshalAny(r.Response, er); err != nil {
        op.Text += err.Error()
        v.selectableFields = 0
      } else {
        ex := renderExecuteResponse(er, v.selection)
        op.Text += ex.text
        selectableFields := v.selectableFields
        v.selectableFields = ex.fields
        selectionActions := v.selectionActions
        v.selectionActions = make([]func(*operationView) View, v.selectableFields)
        for i := 0; i < selectableFields; i++ {
          v.selectionActions[i] = selectionActions[i]
        }
        for e, i := ex.actions.Front(), selectableFields; e != nil; e = e.Next() {
          v.selectionActions[i] = e.Value.(func (v *operationView) View)
          i++
        }
      }
    }
  }
}

type executeText struct {
  text string
  fields int
  actions *list.List
}

func renderExecuteResponse(er *reapi.ExecuteResponse, selection int) executeText {
  ex := renderActionResult(er.Result, selection)
  text := ex.text
  text += proto.MarshalTextString(er.Status)
  text += fmt.Sprintf("served from cache: %v\n", er.CachedResult)
  // server logs...
  if len(er.Message) > 0 {
    text += "message: " + er.Message + "\n"
  }
  return executeText {
    text: text,
    fields: ex.fields,
    actions: ex.actions,
  }
}

func renderRequestMetadata(rm *reapi.RequestMetadata, selection int) string {
  metadataSelections := [...]string {"tool_invocation_id", "correlated_invocations_id"}
  text := "request metadata: " + proto.MarshalTextString(rm)
  if selection < 2 {
    field := metadataSelections[selection]
    match := field + ": (.*)"
    re := regexp.MustCompile(match)
    text = re.ReplaceAllString(text, field + ": [$1](mod:bold)")
  }
  return text
}

func renderActionResult(ar *reapi.ActionResult, selection int) executeText {
  text := fmt.Sprintf("exit code: %d\n", ar.ExitCode)
  base := 3
  actions := list.New()
  if len(ar.StderrRaw) > 0 || (ar.StdoutDigest != nil && ar.StdoutDigest.SizeBytes > 0) {
    text += fmt.Sprintf("stdout: %s%s\n", renderDigest(ar.StdoutDigest, selection == 3), renderInline(len(ar.StdoutRaw)))
    base++
    actions.PushBack(func (v *operationView) View { return contentView(v, ar.StdoutDigest, ar.StdoutRaw, "stdout") })
  }
  if len(ar.StderrRaw) > 0 || (ar.StderrDigest != nil && ar.StderrDigest.SizeBytes > 0) {
    text += fmt.Sprintf("stderr: %s%s\n", renderDigest(ar.StderrDigest, selection == 4), renderInline(len(ar.StderrRaw)))
    base++
    actions.PushBack(func (v *operationView) View { return contentView(v, ar.StderrDigest, ar.StderrRaw, "stderr") })
  }
  for i, of := range ar.OutputFiles {
    path := renderPath(of.Path, of.IsExecutable)
    text += fmt.Sprintf("file: %s (%s)%s\n", path, renderDigest(of.Digest, selection == base + i), renderInline(len(of.Contents)))
    actions.PushBack(func (v *operationView) View { return contentView(v, of.Digest, nil, path) })
  }
  base += len(ar.OutputFiles)
  for _, ofs := range ar.OutputFileSymlinks {
    text += fmt.Sprintf("symlink: %s -> %s\n", ofs.Path, ofs.Target)
  }
  for i, od := range ar.OutputDirectories {
    text += fmt.Sprintf("directory: %s (%s)\n", od.Path, renderDigest(od.TreeDigest, selection == base + i))
    actions.PushBack(func (v *operationView) View { return outputDirectoryView(v, od.TreeDigest, od.Path) })
  }
  fields := base + len(ar.OutputDirectories)
  text += renderExecutionMetadata(ar.ExecutionMetadata)
  return executeText {
    text: text,
    fields: fields,
    actions: actions,
  }
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

func renderExecuteOperationMetadata(em *reapi.ExecuteOperationMetadata, selection int) string {
  stage := &reapi.ExecuteOperationMetadata {
    Stage: em.Stage,
  }
  text := proto.MarshalTextString(stage) + "\n"
  return text + fmt.Sprintf("action: %s\n", renderDigest(em.ActionDigest, selection == 2))
}

func renderInline(l int) string {
  if l > 0 {
    return " inline"
  }
  return ""
}

func renderPath(p string, e bool) string {
  if e {
    return "*" + p
  }
  return p
}

func boldIfSelected(s string, selected bool) string {
  if selected {
    return "[" + s + "](mod:bold)"
  }
  return s
}

// deserves its own file probably
func renderDigest(d *reapi.Digest, selected bool) string {
  return boldIfSelected(client.DigestString(d), selected)
}

func createCorrelatedView(v *operationView) View {
  return v
}

func createInvocationView(v *operationView) View {
  return v
}

func createQueuedOperationView(v *operationView, d *reapi.Digest) View {
  return v
}

func outputDirectoryView(v *operationView, d *reapi.Digest, t string) View {
  return v
}

func createWorkerView(v *operationView, w string) View {
  return v
}

func contentView(v *operationView, d *reapi.Digest, c []byte, t string) View {
  return v
}
