package evl

import (
	"fmt"
	"github.com/flint92/play-with-compiler/craft/ast"
	"github.com/flint92/play-with-compiler/craft/tokens"
)

func Additive(reader tokens.TokenReader) (error, *ast.Node) {
	err, child1 := Multiplicative(reader)
	if err != nil {
		return err, nil
	}

	node := child1
	if child1 != nil {
		for {
			token := reader.Peek()
			if token.Kind == tokens.Plus || token.Kind == tokens.Minus {
				reader.Read()
				err, child2 := Multiplicative(reader)
				if err != nil {
					return err, nil
				}

				if child2 == nil {
					return fmt.Errorf("expect an multiplicative after an operator"), nil
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

	return nil, node
}

func Multiplicative(reader tokens.TokenReader) (error, *ast.Node) {
	err, child1 := Primary(reader)
	if err != nil {
		return err, nil
	}

	node := child1
	if child1 != nil {
		for {
			token := reader.Peek()
			if token.Kind == tokens.Star || token.Kind == tokens.Slash || token.Kind == tokens.Percent {
				reader.Read()
				err, child2 := Primary(reader)
				if err != nil {
					return err, nil
				}

				if child2 == nil {
					return fmt.Errorf("expect an primary after an operator"), nil
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

	return nil, node
}

func Primary(reader tokens.TokenReader) (error, *ast.Node) {
	token := reader.Peek()

	switch token.Kind {
	case tokens.IntLiteral:
		reader.Read()
		return nil, ast.NewNode(ast.IntLiteral, token.Text)
	case tokens.Identifier:
		reader.Read()
		return nil, ast.NewNode(ast.Identifier, token.Text)
	case tokens.OpenParen:
		reader.Read()
		err, node := Additive(reader)
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
