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
			node, err := Parse(scriptText)
			if err != nil {
				_, _ = fmt.Fprintf(out, "%s\n>>", err.Error())
				scriptText = ""
				continue
			}

			result, err := s.eval(node, "")
			if err != nil {
				_, _ = fmt.Fprintf(out, "%s\n>>", err.Error())
			} else {
				_, _ = fmt.Fprintf(out, "%d\n>>", result)
			}
			scriptText = ""
		}

	}
}

func (s *Script) eval(node *ast.Node, indent string) (int64, error) {
	result := int64(0)

	switch node.Type() {
	case ast.Program:
		for _, child := range node.Children() {
			ro, err := s.eval(child, indent+"\t")
			if err != nil {
				return 0, err
			} else {
				result = ro
			}
		}
	case ast.Additive:
		child1 := node.Children()[0]
		val1, err := s.eval(child1, indent+"\t")
		if err != nil {
			return 0, err
		}

		child2 := node.Children()[1]
		val2, err := s.eval(child2, indent+"\t")
		if err != nil {
			return 0, err
		}

		if "+" == node.Text() {
			result = val1 + val2
		} else {
			result = val1 - val2
		}
	case ast.Multiplicative:
		child1 := node.Children()[0]
		val1, err := s.eval(child1, indent+"\t")
		if err != nil {
			return 0, err
		}

		child2 := node.Children()[1]
		val2, err := s.eval(child2, indent+"\t")
		if err != nil {
			return 0, err
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
			return 0, err
		}
		result = i
	case ast.Identifier:
		v, ok := s.variables[node.Text()]
		if !ok {
			return 0, fmt.Errorf("unknown variable: %s", node.Text())
		}
		result = v
	case ast.AssignmentStmt:
		_, ok := s.variables[node.Text()]
		if !ok {
			return 0, fmt.Errorf("unknown variable: %s", node.Text())
		}
		if len(node.Children()) > 0 {
			child := node.Children()[0]
			val, err := s.eval(child, indent+"\t")
			if err != nil {
				return 0, err
			}
			s.variables[node.Text()] = val
			result = val
		}
	case ast.IntDeclaration:
		if len(node.Children()) > 0 {
			child := node.Children()[0]
			val, err := s.eval(child, indent+"\t")
			if err != nil {
				return 0, err
			}
			s.variables[node.Text()] = val
			result = val
		}
	default:
		return 0, fmt.Errorf("unexpected node type: %s\n", ast.GetNodeTypeString(node.Type()))
	}

	return result, nil
}
