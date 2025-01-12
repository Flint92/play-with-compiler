package script

import (
	"bufio"
	"fmt"
	"github.com/flint92/play-with-compiler/craft/ast"
	"io"
	"strconv"
	"strings"
)

type Script struct {
	variables map[string]int64
}

func NewScript() *Script {
	return &Script{
		variables: make(map[string]int64),
	}
}

func (s *Script) Run(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	_, _ = fmt.Fprint(out, "\n>>")
	scriptText := ""
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if line == "exit();" {
			_, _ = fmt.Fprintf(out, "%s\n", "Bye!")
			break
		}

		scriptText += line + "\n"
		if strings.HasSuffix(line, ";") {
			err, node := Parse(scriptText)
			if err != nil {
				_, _ = fmt.Fprintf(out, "%s\n>>", err.Error())
				scriptText = ""
				continue
			}

			err, result := s.eval(node, "")
			if err != nil {
				_, _ = fmt.Fprintf(out, "%s\n>>", err.Error())
			} else {
				_, _ = fmt.Fprintf(out, "%d\n>>", result)
			}
			scriptText = ""
		}

	}
}

func (s *Script) eval(node *ast.Node, indent string) (error, int64) {
	result := int64(0)

	switch node.Type() {
	case ast.Program:
		for _, child := range node.Children() {
			err, ro := s.eval(child, indent+"\t")
			if err != nil {
				return err, 0
			} else {
				result = ro
			}
		}
	case ast.Additive:
		child1 := node.Children()[0]
		err, val1 := s.eval(child1, indent+"\t")
		if err != nil {
			return err, 0
		}

		child2 := node.Children()[1]
		err, val2 := s.eval(child2, indent+"\t")
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
		err, val1 := s.eval(child1, indent+"\t")
		if err != nil {
			return err, 0
		}

		child2 := node.Children()[1]
		err, val2 := s.eval(child2, indent+"\t")
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
	case ast.Identifier:
		v, ok := s.variables[node.Text()]
		if !ok {
			return fmt.Errorf("unknown variable: %s", node.Text()), 0
		}
		result = v
	case ast.AssignmentStmt:
		_, ok := s.variables[node.Text()]
		if !ok {
			return fmt.Errorf("unknown variable: %s", node.Text()), 0
		}
		if len(node.Children()) > 0 {
			child := node.Children()[0]
			err, val := s.eval(child, indent+"\t")
			if err != nil {
				return err, 0
			}
			s.variables[node.Text()] = val
			result = val
		}
	case ast.IntDeclaration:
		if len(node.Children()) > 0 {
			child := node.Children()[0]
			err, val := s.eval(child, indent+"\t")
			if err != nil {
				return err, 0
			}
			s.variables[node.Text()] = val
			result = val
		}
	default:
		return fmt.Errorf("unexpected node type: %s\n", ast.GetNodeTypeString(node.Type())), 0
	}

	return nil, result
}
