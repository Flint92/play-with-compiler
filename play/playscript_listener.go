// Code generated from PlayScript.g4 by ANTLR 4.13.2. DO NOT EDIT.

package play // PlayScript
import "github.com/antlr4-go/antlr/v4"

// PlayScriptListener is a complete listener for a parse tree produced by PlayScriptParser.
type PlayScriptListener interface {
	antlr.ParseTreeListener

	// EnterBracedExpression is called when entering the bracedExpression production.
	EnterBracedExpression(c *BracedExpressionContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterIntegerLiteral is called when entering the integerLiteral production.
	EnterIntegerLiteral(c *IntegerLiteralContext)

	// EnterFloatLiteral is called when entering the floatLiteral production.
	EnterFloatLiteral(c *FloatLiteralContext)

	// ExitBracedExpression is called when exiting the bracedExpression production.
	ExitBracedExpression(c *BracedExpressionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitIntegerLiteral is called when exiting the integerLiteral production.
	ExitIntegerLiteral(c *IntegerLiteralContext)

	// ExitFloatLiteral is called when exiting the floatLiteral production.
	ExitFloatLiteral(c *FloatLiteralContext)
}
