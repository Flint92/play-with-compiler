package evl

import (
	"fmt"
	"github.com/flint92/play-with-compiler/craft/ast"
	"github.com/flint92/play-with-compiler/craft/tokens"
)

func Additive(reader tokens.TokenReader) (*ast.Node, error) {
	child1, err := Multiplicative(reader)
	if err != nil {
		return nil, err
	}

	node := child1
	if child1 != nil {
		for {
			token := reader.Peek()
			if token.Kind == tokens.Plus || token.Kind == tokens.Minus {
				reader.Read()
				child2, err := Multiplicative(reader)
				if err != nil {
					return nil, err
				}

				if child2 == nil {
					return nil, fmt.Errorf("expect an multiplicative after an operator")
				}

				node = ast.NewNode(ast.Additive, token.Text)
				node.AddChild(child1)
				node.AddChild(child2)

				child1 = node
			} else {
				break
			}
		}
	}

	return node, nil
}

func Multiplicative(reader tokens.TokenReader) (*ast.Node, error) {
	child1, err := Primary(reader)
	if err != nil {
		return nil, err
	}

	node := child1
	if child1 != nil {
		for {
			token := reader.Peek()
			if token.Kind == tokens.Star || token.Kind == tokens.Slash || token.Kind == tokens.Percent {
				reader.Read()
				child2, err := Primary(reader)
				if err != nil {
					return nil, err
				}

				if child2 == nil {
					return nil, fmt.Errorf("expect an primary after an operator")
				}

				node = ast.NewNode(ast.Multiplicative, token.Text)
				node.AddChild(child1)
				node.AddChild(child2)

				child1 = node
			} else {
				break
			}
		}
	}

	return node, nil
}

func Primary(reader tokens.TokenReader) (*ast.Node, error) {
	token := reader.Peek()

	switch token.Kind {
	case tokens.IntLiteral:
		reader.Read()
		return ast.NewNode(ast.IntLiteral, token.Text), nil
	case tokens.Identifier:
		reader.Read()
		return ast.NewNode(ast.Identifier, token.Text), nil
	case tokens.OpenParen:
		reader.Read()
		node, err := Additive(reader)
		if err != nil {
			return nil, err
		}

		if node == nil {
			return nil, fmt.Errorf("expect an additive expression inside parenthesis")
		}

		token = reader.Peek()
		if token.Kind != tokens.CloseParen {
			return nil, fmt.Errorf("expecting close parenthesis")
		}

		reader.Read()

		return node, nil
	default:
		return nil, fmt.Errorf("unexpected token kind %s\n", tokens.GetTokenKindString(token.Kind))
	}
}
