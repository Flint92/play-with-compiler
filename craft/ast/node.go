package ast

type Node struct {
	parent   *Node
	children []*Node
	nType    NodeType
	text     string
}

func NewNode(nType NodeType, text string) *Node {
	return &Node{
		nType:    nType,
		text:     text,
		children: make([]*Node, 0),
	}
}

func (s *Node) Parent() *Node {
	return s.parent
}

func (s *Node) Children() []*Node {
	return s.children[:]
}

func (s *Node) Type() NodeType {
	return s.nType
}

func (s *Node) Text() string {
	return s.text
}

func (s *Node) AddChild(child *Node) {
	s.children = append(s.children, child)
	child.parent = s
}
