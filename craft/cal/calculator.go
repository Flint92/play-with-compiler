package cal

import (
	"fmt"
	"github.com/flint92/play-with-compiler/craft/ast"
	"github.com/flint92/play-with-compiler/craft/evl"
	"github.com/flint92/play-with-compiler/craft/lexer"
	"github.com/flint92/play-with-compiler/craft/tokens"
	"strconv"
)

type SimpleCalculator struct {
	reader tokens.TokenReader
	root   *ast.Node
}

func Eval(script string) (int64, error) {
	reader := lexer.Parse(script)

	cal := &SimpleCalculator{
		reader: reader,
	}

	root, err := cal.parse()
	if err != nil {
		return 0, err
	}

	cal.root = root

	return eval(root, "")
}

func (cal *SimpleCalculator) parse() (*ast.Node, error) {
	child, err := evl.Additive(cal.reader)
	if err != nil {
		return nil, err
	}

	root := ast.NewNode(ast.Program, "Calculator")
	if child != nil {
		root.AddChild(child)
	}
	return root, nil
}

func eval(node *ast.Node, indent string) (int64, error) {
	result := int64(0)

	fmt.Printf("%s Calculating: %s\n", indent, ast.GetNodeTypeString(node.Type()))

	switch node.Type() {
	case ast.Program:
		for _, child := range node.Children() {
			ro, err := eval(child, indent+"\t")
			if err != nil {
				return 0, err
			} else {
				result = ro
			}
		}
	case ast.Additive:
		child1 := node.Children()[0]
		val1, err := eval(child1, indent+"\t")
		if err != nil {
			return 0, err
		}

		child2 := node.Children()[1]
		val2, err := eval(child2, indent+"\t")
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
		val1, err := eval(child1, indent+"\t")
		if err != nil {
			return 0, err
		}

		child2 := node.Children()[1]
		val2, err := eval(child2, indent+"\t")
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
	default:
		return 0, fmt.Errorf("unexpected node type: %s\n", ast.GetNodeTypeString(node.Type()))
	}

	fmt.Printf("%sResult: %d\n", indent, result)

	return result, nil
}
