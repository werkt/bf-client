package view

import (
  "fmt"
  "github.com/gammazero/deque"
  reapi "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
  bfpb "github.com/buildfarm/buildfarm/build/buildfarm/v1test"
  ui "github.com/gizak/termui/v3"
  "github.com/gizak/termui/v3/widgets"
  "github.com/werkt/bf-client/client"
  "google.golang.org/grpc"
)

type nodeValue struct {
  name string
  size int
}

func (nv nodeValue) String() string {
  if nv.size == 0 {
    return nv.name
  }
  // maybe have N / size if opened
  return fmt.Sprintf("%s (%d)", nv.name, nv.size)
}

type inputView struct {
  a *client.App
  d bfpb.Digest
  i map[string]*reapi.Directory
  err error
  nodes []*client.TreeNode
  t *client.Tree
  v View
}

func NewInput(a *client.App, d bfpb.Digest, v View) View {
  return &inputView {
    a: a,
    d: d,
    v: v,
  }
}

func (v *inputView) Update() {
  if v.i != nil {
    v.t.Title = "Directory: " + client.DigestString(v.d) + fmt.Sprintf(" (%d/%d)", v.t.SelectedRow, v.t.Size())
    return
  }

  v.i = make(map[string]*reapi.Directory)
  client.FetchTree(v.d, v.i, v.a.Conn)

  t := client.NewTree()
  t.Focused = true
  t.Title = "Directory: " + client.DigestString(v.d)
  t.WrapText = false
  t.SelectedRowStyle = ui.NewStyle(ui.ColorBlack, ui.ColorWhite)
  root := client.DigestString(v.d)
  sizes := make(map[string]int)
  v.nodes = createInputNodes(v.i[root], root, v.d.DigestFunction, v.i, sizes)
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

func fetchTreeRecursive(d reapi.Digest, df reapi.DigestFunction_Value, i map[string]*reapi.Directory, conn *grpc.ClientConn) {
  var q deque.Deque[reapi.Digest]
  q.PushFront(d)
  for q.Len() != 0 {
    d := q.PopBack()
    digest := client.ToDigest(d, df)
    i[client.DigestString(digest)] = nil
    fetchDirectory(digest, &q, i, conn)
  }
}

func fetchDirectory(d bfpb.Digest, q *deque.Deque[reapi.Digest], i map[string]*reapi.Directory, conn *grpc.ClientConn) {
  dir := &reapi.Directory{}
  if err := client.Expect(conn, d, dir); err != nil {
    return
  }
  i[client.DigestString(d)] = dir
  for _, cd := range dir.Directories {
    cs := client.DigestString(client.ToDigest(*cd.Digest, d.DigestFunction))
    _, exists := i[cs]
    if !exists {
      if cd.Digest.SizeBytes == 0 {
        i[cs] = &reapi.Directory{}
      } else {
        q.PushFront(*cd.Digest)
      }
    }
  }
}

func createInputNodes(d *reapi.Directory, dd string, df reapi.DigestFunction_Value, i map[string]*reapi.Directory, sizes map[string]int) []*client.TreeNode {
  nodes := []*client.TreeNode{}
  size := 0
  for _, n := range d.Directories {
    child := client.DigestString(client.ToDigest(*n.Digest, df))
    childNodes := createInputNodes(i[child], child, df, i, sizes)
    childSize := sizes[child]
    nodes = append(nodes, &client.TreeNode{
      Value: nodeValue{name: n.Name, size: childSize},
      Nodes: childNodes,
    })
    size += childSize
  }
  size += len(d.Files)
  for _, n := range d.Files {
    nodes = append(nodes, &client.TreeNode{
      Value: nodeValue{name: n.Name},
    })
  }
  sizes[dd] = size
  return nodes
}

func expandNodeAll(n *client.TreeNode, e bool) {
  if len(n.Nodes) > 0 {
    n.Expanded = e
    for _, cn := range n.Nodes {
      expandNodeAll(cn, e)
    }
  }
}
