package script

import (
	"fmt"
	"github.com/flint92/play-with-compiler/craft/ast"
	"github.com/flint92/play-with-compiler/craft/evl"
	"github.com/flint92/play-with-compiler/craft/lexer"
	"github.com/flint92/play-with-compiler/craft/tokens"
)

func Parse(script string) (*ast.Node, error) {
	reader := lexer.Parse(script)
	return prog(reader)
}

func prog(reader *lexer.Lexer) (*ast.Node, error) {
	node := ast.NewNode(ast.Program, "pwd")

	for reader.Peek() != tokens.EOF {
		child, err := intDeclare(reader)
		if err != nil {
			return nil, err
		}

		if child == nil {
			child, err = expressionStmt(reader)
			if err != nil {
				return nil, err
			}
		}

		if child == nil {
			child, err = assignmentSmt(reader)
			if err != nil {
				return nil, err
			}
		}

		if child == nil {
			return nil, fmt.Errorf("unknown statement")
		}

		node.AddChild(child)
	}

	return node, nil
}

func assignmentSmt(reader *lexer.Lexer) (*ast.Node, error) {
	token := reader.Peek()
	if token.Kind == tokens.Identifier {
		token = reader.Read()
		node := ast.NewNode(ast.AssignmentStmt, token.Text)
		token = reader.Peek()
		if token.Kind == tokens.Assignment {
			reader.Read()
			child, err := evl.Additive(reader)
			if err != nil {
				return nil, err
			}
			if child == nil {
				return nil, fmt.Errorf("expect an expression after =")
			}
			node.AddChild(child)

			token = reader.Peek()
			if token.Kind == tokens.SemiColon {
				reader.Read()
			} else {
				return nil, fmt.Errorf("expect a semicolon at the end of line")
			}
		} else {
			reader.Unread()
			return nil, nil
		}
	}
	return nil, nil
}

func expressionStmt(reader *lexer.Lexer) (*ast.Node, error) {
	pos := reader.Position()
	node, err := evl.Additive(reader)
	if err != nil {
		return nil, err
	}

	if node != nil {
		token := reader.Peek()
		if token.Kind == tokens.SemiColon {
			reader.Read()
		} else {
			reader.SetPosition(pos)
			return nil, nil
		}
	}

	return node, nil
}

func intDeclare(reader *lexer.Lexer) (*ast.Node, error) {
	token := reader.Peek()
	if token.Kind == tokens.Int {
		token = reader.Read()

		if reader.Peek().Kind == tokens.Identifier {
			token = reader.Read()
			node := ast.NewNode(ast.IntDeclaration, token.Text)
			token = reader.Peek()
			if token.Kind == tokens.Assignment {
				reader.Read()
				child, err := evl.Additive(reader)
				if err != nil {
					return nil, err
				}

				if child == nil {
					return nil, fmt.Errorf("expect an expression after =")
				}
				node.AddChild(child)

				token = reader.Peek()
				if token.Kind == tokens.SemiColon {
					reader.Read()
					return node, nil
				} else {
					return nil, fmt.Errorf("expect a semicolon at the end of line")
				}
			}

			if token.Kind == tokens.SemiColon {
				reader.Read()
				return node, nil
			} else {
				return nil, fmt.Errorf("expect a semicolon at the end of line")
			}
		} else {
			return nil, fmt.Errorf("expect an identifier after int")
		}
	}
	return nil, nil
}
