package ast

type Node interface {
	Parent() Node
	Children() []Node
	Type() NodeType
	Text() string
}
