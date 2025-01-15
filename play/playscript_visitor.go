// Code generated from PlayScript.g4 by ANTLR 4.13.2. DO NOT EDIT.

package play // PlayScript
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by PlayScriptParser.
type PlayScriptVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PlayScriptParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by PlayScriptParser#integerLiteral.
	VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{}
}
