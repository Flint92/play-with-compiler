package cal

import (
	"fmt"
	"github.com/flint92/play-with-compiler/craft/ast"
	"github.com/flint92/play-with-compiler/craft/lexer"
	"github.com/flint92/play-with-compiler/craft/tokens"
	"strconv"
)

type SimpleCalculator struct {
	reader tokens.TokenReader
	root   ast.Node
}

func Eval(script string) (error, int64) {
	reader := lexer.Parse(script)

	cal := &SimpleCalculator{
		reader: reader,
	}

	err, root := cal.parse()
	if err != nil {
		return err, 0
	}

	cal.root = root

	return eval(root, "")
}

func (cal *SimpleCalculator) parse() (error, ast.Node) {
	err, child := additive(cal.reader)
	if err != nil {
		return err, nil
	}

	root := newSimpleNode(ast.Program, "Calculator")
	if child != nil {
		root.addChild(child)
	}
	return nil, root
}

func eval(node ast.Node, indent string) (error, int64) {
	result := int64(0)

	fmt.Printf("%s Calculating: %s\n", indent, ast.GetNodeTypeString(node.Type()))

	switch node.Type() {
	case ast.Program:
		for _, child := range node.Children() {
			err, ro := eval(child, indent+"\t")
			if err != nil {
				return err, 0
			} else {
				result = ro
			}
		}
	case ast.Additive:
		child1 := node.Children()[0]
		err, val1 := eval(child1, indent+"\t")
		if err != nil {
			return err, 0
		}

		child2 := node.Children()[1]
		err, val2 := eval(child2, indent+"\t")
		if err != nil {
			return err, 0
		}

		if "+" == node.Text() {
			result = val1 + val2
		} else {
			result = val1 - val2
		}
	case ast.Multiplicative:
		child1 := node.Children()[0]
		err, val1 := eval(child1, indent+"\t")
		if err != nil {
			return err, 0
		}

		child2 := node.Children()[1]
		err, val2 := eval(child2, indent+"\t")
		if err != nil {
			return err, 0
		}

		if "*" == node.Text() {
			result = val1 * val2
		} else if "/" == node.Text() {
			result = val1 / val2
		} else {
			result = val1 % val2
		}
	case ast.IntLiteral:
		i, err := strconv.ParseInt(node.Text(), 10, 64)
		if err != nil {
			return err, 0
		}
		result = i
	default:
		return fmt.Errorf("unexpected node type: %s\n", ast.GetNodeTypeString(node.Type())), 0
	}

	fmt.Printf("%sResult: %d\n", indent, result)

	return nil, result
}

func additive(reader tokens.TokenReader) (error, *simpleNode) {
	err, child1 := multiplicative(reader)
	if err != nil {
		return err, nil
	}

	node := child1
	if child1 != nil {
		for {
			token := reader.Peek()
			if token.Kind == tokens.Plus || token.Kind == tokens.Minus {
				reader.Read()
				err, child2 := multiplicative(reader)
				if err != nil {
					return err, nil
				}

				if child2 == nil {
					return fmt.Errorf("expect an multiplicative after an operator"), nil
				}

				node = newSimpleNode(ast.Additive, token.Text)
				node.addChild(child1)
				node.addChild(child2)

				child1 = node
			} else {
				break
			}
		}
	}

	return nil, node
}

func multiplicative(reader tokens.TokenReader) (error, *simpleNode) {
	err, child1 := primary(reader)
	if err != nil {
		return err, nil
	}

	node := child1
	if child1 != nil {
		for {
			token := reader.Peek()
			if token.Kind == tokens.Star || token.Kind == tokens.Slash || token.Kind == tokens.Percent {
				reader.Read()
				err, child2 := primary(reader)
				if err != nil {
					return err, nil
				}

				if child2 == nil {
					return fmt.Errorf("expect an primary after an operator"), nil
				}

				node = newSimpleNode(ast.Multiplicative, token.Text)
				node.addChild(child1)
				node.addChild(child2)

				child1 = node
			} else {
				break
			}
		}
	}

	return nil, node
}

func primary(reader tokens.TokenReader) (error, *simpleNode) {
	token := reader.Peek()

	switch token.Kind {
	case tokens.IntLiteral:
		reader.Read()
		return nil, newSimpleNode(ast.IntLiteral, token.Text)
	case tokens.Identifier:
		reader.Read()
		return nil, newSimpleNode(ast.Identifier, token.Text)
	case tokens.OpenParen:
		reader.Read()
		err, node := additive(reader)
		if err != nil {
			return err, nil
		}

		if node == nil {
			return fmt.Errorf("expect an additive expression inside parenthesis"), nil
		}

		token = reader.Peek()
		if token.Kind != tokens.CloseParen {
			return fmt.Errorf("expecting close parenthesis"), nil
		}

		reader.Read()

		return nil, node
	default:
		return fmt.Errorf("unexpected token kind %s\n", tokens.GetTokenKindString(token.Kind)), nil
	}
}

type simpleNode struct {
	parent   *simpleNode
	children []*simpleNode
	nType    ast.NodeType
	text     string
}

func newSimpleNode(nType ast.NodeType, text string) *simpleNode {
	return &simpleNode{
		nType:    nType,
		text:     text,
		children: make([]*simpleNode, 0),
	}
}

func (s *simpleNode) Parent() ast.Node {
	return s.parent
}

func (s *simpleNode) Children() []ast.Node {
	if len(s.children) == 0 {
		return []ast.Node{}
	}
	children := make([]ast.Node, 0)
	for _, child := range s.children {
		children = append(children, child)
	}
	return children
}

func (s *simpleNode) Type() ast.NodeType {
	return s.nType
}

func (s *simpleNode) Text() string {
	return s.text
}

func (s *simpleNode) addChild(child *simpleNode) {
	s.children = append(s.children, child)
}
