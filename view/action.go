package view

import (
  reapi "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
  bfpb "github.com/buildfarm/buildfarm/build/buildfarm/v1test"
  ui "github.com/gizak/termui/v3"
  "github.com/gizak/termui/v3/widgets"
  "github.com/golang/protobuf/proto"
  "github.com/werkt/bf-client/client"
)

type actionView struct {
  a *client.App
  d bfpb.Digest
  action *reapi.Action
  err error
  v View
  selection int
}

func NewAction(a *client.App, d bfpb.Digest, v View) View {
  return &actionView {
    a: a,
    d: d,
    v: v,
  }
}

func (v *actionView) Handle(e ui.Event) View {
  switch e.ID {
  case "j", "<Down>", "k", "<Up>":
    // same movement up or down for 2 elements
    v.selection++
    v.selection %= 2
  case "<Escape>", "q", "<C-c>":
    ui.Clear()
    return v.v
  case "<Enter>":
    switch v.selection {
    case 0:
      return NewCommand(v.a, client.ToDigest(*v.action.CommandDigest, v.d.DigestFunction), v)
    case 1:
      return NewInput(v.a, client.ToDigest(*v.action.InputRootDigest, v.d.DigestFunction), v)
    }
  }
  return v
}

func (v *actionView) Update() {
  a := &reapi.Action{}
  err := client.Expect(v.a.Conn, v.d, a)
  if err != nil {
    v.err = err
  } else {
    v.action = a
  }
}

func (v actionView) Render() []ui.Drawable {
  p := widgets.NewParagraph()
  p.Title = client.DigestString(v.d)
  p.WrapText = true
  p.SetRect(0, 0, 120, 60)
  if v.err != nil {
    p.Text = string(v.err.Error())
  } else {
    p.Text = renderAction(v.action, v.d.DigestFunction, v.selection)
  }
  return []ui.Drawable { p }
}

func renderAction(a *reapi.Action, df reapi.DigestFunction_Value, selection int) string {
  text := "command: " + renderREDigest(*a.CommandDigest, df, selection == 0) + "\n"
  text += "input_root: " + renderREDigest(*a.InputRootDigest, df, selection == 1) + "\n"
  // text += "salt: " + proto.MarshalTextString(a.Salt) + "\n"
  if len(a.Platform.Properties) > 0 {
    text += "platform: " + proto.MarshalTextString(a.Platform)
  }
  return text
}
