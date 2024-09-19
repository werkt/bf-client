package view

import (
  "github.com/gammazero/deque"
  reapi "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
  ui "github.com/gizak/termui/v3"
  "github.com/gizak/termui/v3/widgets"
  "github.com/werkt/bf-client/client"
  "google.golang.org/grpc"
)

type nodeValue string

func (nv nodeValue) String() string {
  return string(nv)
}

type inputView struct {
  a *client.App
  d *reapi.Digest
  i map[string]*reapi.Directory
  err error
  nodes []*widgets.TreeNode
  t *widgets.Tree
  v View
}

func NewInput(a *client.App, d *reapi.Digest, v View) View {
  return &inputView {
    a: a,
    d: d,
    v: v,
  }
}

func (v *inputView) Update() {
  if v.i != nil {
    return
  }

  v.i = make(map[string]*reapi.Directory)
  client.FetchTree(v.d, v.i, v.a.Conn)

  t := widgets.NewTree()
  t.Title = "Directory: " + client.DigestString(v.d)
  t.WrapText = false
  t.SelectedRowStyle = ui.NewStyle(ui.ColorBlack, ui.ColorWhite)
  v.nodes = createInputNodes(v.i[client.DigestString(v.d)], v.i)
  t.SetNodes(v.nodes)
  // setting this on every frame seems to jank it up
  w, h := ui.TerminalDimensions()
  t.SetRect(0, 0, w, h)
  v.t = t
}

func (v *inputView) Handle(e ui.Event) View {
  switch e.ID {
  case "<Escape>", "q", "<C-c>":
    ui.Clear()
    return v.v
  case "j", "<Down>":
    v.t.ScrollDown()
  case "k", "<Up>":
    v.t.ScrollUp()
  case "<Resize>":
    w, h := ui.TerminalDimensions()
    v.t.SetRect(0, 0, w, h)
  case "l", "<Right>":
    v.t.Expand()
  case "L", "<S-Right>":
    n := v.t.SelectedNode()
    for _, cn := range n.Nodes {
      expandNodeAll(cn, true)
    }
    // hack to get prepareNodes
    v.t.Expand()
  case "h", "<Left>":
    v.t.Collapse()
  case "H", "<S-Left>":
    n := v.t.SelectedNode()
    for _, cn := range n.Nodes {
      expandNodeAll(cn, false)
    }
    // hack to get prepareNodes
    v.t.Collapse()
  case "<Enter>":
    v.t.ToggleExpand()
  case "E":
    v.t.ExpandAll()
  case "C":
    v.t.CollapseAll()
  }
  return v
}

func (v inputView) Render() []ui.Drawable {
  var r ui.Drawable
  if v.err != nil {
    p := widgets.NewParagraph()
    p.Text = string(v.err.Error())
    r = p
  } else {
    r = v.t
  }
  return []ui.Drawable { r }
}

func fetchTreeRecursive(d *reapi.Digest, i map[string]*reapi.Directory, conn *grpc.ClientConn) {
  var q deque.Deque[*reapi.Digest]
  q.PushFront(d)
  for q.Len() != 0 {
    d := q.PopBack()
    i[client.DigestString(d)] = nil
    fetchDirectory(d, &q, i, conn)
  }
}

func fetchDirectory(d *reapi.Digest, q *deque.Deque[*reapi.Digest], i map[string]*reapi.Directory, conn *grpc.ClientConn) {
  dir := &reapi.Directory{}
  if err := client.Expect(conn, d, dir); err != nil {
    return
  }
  i[client.DigestString(d)] = dir
  for _, cd := range dir.Directories {
    cs := client.DigestString(cd.Digest)
    _, exists := i[cs]
    if !exists {
      if cd.Digest.SizeBytes == 0 {
        i[cs] = &reapi.Directory{}
      } else {
        q.PushFront(cd.Digest)
      }
    }
  }
}

func createInputNodes(d *reapi.Directory, i map[string]*reapi.Directory) []*widgets.TreeNode {
  nodes := []*widgets.TreeNode{}
  for _, n := range d.Directories {
    nodes = append(nodes, &widgets.TreeNode{
      Value: nodeValue(n.Name),
      Nodes: createInputNodes(i[client.DigestString(n.Digest)], i),
    })
  }
  for _, n := range d.Files {
    nodes = append(nodes, &widgets.TreeNode{
      Value: nodeValue(n.Name),
    })
  }
  return nodes
}

func expandNodeAll(n *widgets.TreeNode, e bool) {
  if len(n.Nodes) > 0 {
    n.Expanded = e
    for _, cn := range n.Nodes {
      expandNodeAll(cn, e)
    }
  }
}
