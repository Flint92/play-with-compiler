package script

import (
	"fmt"
	"github.com/flint92/play-with-compiler/craft/ast"
	"github.com/flint92/play-with-compiler/craft/evl"
	"github.com/flint92/play-with-compiler/craft/lexer"
	"github.com/flint92/play-with-compiler/craft/tokens"
)

func Parse(script string) (error, *ast.Node) {
	reader := lexer.Parse(script)
	return prog(reader)
}

func prog(reader *lexer.Lexer) (error, *ast.Node) {
	node := ast.NewNode(ast.Program, "pwd")

	for reader.Peek() != tokens.EOF {
		err, child := intDeclare(reader)
		if err != nil {
			return err, nil
		}

		if child == nil {
			err, child = expressionStmt(reader)
			if err != nil {
				return err, nil
			}
		}

		if child == nil {
			err, child = assignmentSmt(reader)
			if err != nil {
				return err, nil
			}
		}

		if child == nil {
			return fmt.Errorf("unknown statement"), nil
		}

		node.AddChild(child)
	}

	return nil, node
}

func assignmentSmt(reader *lexer.Lexer) (error, *ast.Node) {
	token := reader.Peek()
	if token.Kind == tokens.Identifier {
		token = reader.Read()
		node := ast.NewNode(ast.AssignmentStmt, token.Text)
		token = reader.Peek()
		if token.Kind == tokens.Assignment {
			reader.Read()
			err, child := evl.Additive(reader)
			if err != nil {
				return err, nil
			}
			if child == nil {
				return fmt.Errorf("expect an expression after ="), nil
			}
			node.AddChild(child)

			token = reader.Peek()
			if token.Kind == tokens.SemiColon {
				reader.Read()
			} else {
				return fmt.Errorf("expect a semicolon at the end of line"), nil
			}
		} else {
			reader.Unread()
			return nil, nil
		}
	}
	return nil, nil
}

func expressionStmt(reader *lexer.Lexer) (error, *ast.Node) {
	pos := reader.Position()
	err, node := evl.Additive(reader)
	if err != nil {
		return err, nil
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

	return nil, node
}

func intDeclare(reader *lexer.Lexer) (error, *ast.Node) {
	token := reader.Peek()
	if token.Kind == tokens.Int {
		token = reader.Read()

		if reader.Peek().Kind == tokens.Identifier {
			token = reader.Read()
			node := ast.NewNode(ast.IntDeclaration, token.Text)
			token = reader.Peek()
			if token.Kind == tokens.Assignment {
				reader.Read()
				err, child := evl.Additive(reader)
				if err != nil {
					return err, nil
				}

				if child == nil {
					return fmt.Errorf("expect an expression after ="), nil
				}
				node.AddChild(child)

				token = reader.Peek()
				if token.Kind == tokens.SemiColon {
					reader.Read()
					return nil, node
				} else {
					return fmt.Errorf("expect a semicolon at the end of line"), nil
				}
			}

			if token.Kind == tokens.SemiColon {
				reader.Read()
				return nil, node
			} else {
				return fmt.Errorf("expect a semicolon at the end of line"), nil
			}
		} else {
			return fmt.Errorf("expect an identifier after int"), nil
		}
	}
	return nil, nil
}
