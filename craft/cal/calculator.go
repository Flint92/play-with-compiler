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

func (cal *SimpleCalculator) parse() (error, *ast.Node) {
	err, child := evl.Additive(cal.reader)
	if err != nil {
		return err, nil
	}

	root := ast.NewNode(ast.Program, "Calculator")
	if child != nil {
		root.AddChild(child)
	}
	return nil, root
}

func eval(node *ast.Node, indent string) (error, int64) {
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
