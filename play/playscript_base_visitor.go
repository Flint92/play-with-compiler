// Code generated from PlayScript.g4 by ANTLR 4.13.2. DO NOT EDIT.

package play // PlayScript
import "github.com/antlr4-go/antlr/v4"

type BasePlayScriptVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePlayScriptVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitPrimary(ctx *PrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlayScriptVisitor) VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}
