// Code generated from PlayScript.g4 by ANTLR 4.13.2. DO NOT EDIT.

package play // PlayScript
import "github.com/antlr4-go/antlr/v4"

type BasePlayScriptVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePlayScriptVisitor) VisitBracedExpression(ctx *BracedExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitPrimary(ctx *PrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitFloatLiteral(ctx *FloatLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}
