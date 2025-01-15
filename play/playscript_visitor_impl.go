package play

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"strconv"
)

type ScriptVisitorImpl struct {
}

func NewPlayScriptVisitor() *ScriptVisitorImpl {
	return &ScriptVisitorImpl{}

}

func (s *ScriptVisitorImpl) Visit(tree antlr.ParseTree) interface{} {
	switch t := tree.(type) {
	case *ExpressionContext:
		return s.VisitExpression(t)
	case antlr.RuleNode:
		return s.VisitChildren(t)
	default:
		panic(fmt.Sprintf("Not implemented yet %T", t))
	}
}

func (s *ScriptVisitorImpl) VisitChildren(node antlr.RuleNode) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitTerminal(node antlr.TerminalNode) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitErrorNode(node antlr.ErrorNode) interface{} {
	panic(node.GetPayload())
}

func (s *ScriptVisitorImpl) VisitLiteral(ctx *LiteralContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitFloatLiteral(ctx *FloatLiteralContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitExpression(ctx *ExpressionContext) interface{} {
	if ctx.GetChildCount() == 1 {
		// Single number
		value, _ := strconv.ParseInt(ctx.GetText(), 10, 64)
		return value
	} else if ctx.GetChildCount() == 3 {
		// Binary operation
		left := s.Visit(ctx.GetChild(0).(antlr.ParseTree)).(int64)
		right := s.Visit(ctx.GetChild(2).(antlr.ParseTree)).(int64)
		operator := ctx.GetChild(1).GetPayload().(*antlr.CommonToken).GetText()

		switch operator {
		case "+":
			return left + right
		case "-":
			return left - right
		case "*":
			return left * right
		case "/":
			return left / right
		case "%":
			return left % right
		}
	}
	return 0
}

func (s *ScriptVisitorImpl) VisitPrimary(ctx *PrimaryContext) interface{} {
	//TODO implement me
	panic("implement me")
}
