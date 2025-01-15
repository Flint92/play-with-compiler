// Code generated from PlayScript.g4 by ANTLR 4.13.2. DO NOT EDIT.

package play // PlayScript
import "github.com/antlr4-go/antlr/v4"

// BasePlayScriptListener is a complete listener for a parse tree produced by PlayScriptParser.
type BasePlayScriptListener struct{}

var _ PlayScriptListener = &BasePlayScriptListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePlayScriptListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePlayScriptListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePlayScriptListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePlayScriptListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterBracedExpression is called when production bracedExpression is entered.
func (s *BasePlayScriptListener) EnterBracedExpression(ctx *BracedExpressionContext) {}

// ExitBracedExpression is called when production bracedExpression is exited.
func (s *BasePlayScriptListener) ExitBracedExpression(ctx *BracedExpressionContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePlayScriptListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePlayScriptListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BasePlayScriptListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BasePlayScriptListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BasePlayScriptListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BasePlayScriptListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BasePlayScriptListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BasePlayScriptListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BasePlayScriptListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BasePlayScriptListener) ExitLiteral(ctx *LiteralContext) {}

// EnterIntegerLiteral is called when production integerLiteral is entered.
func (s *BasePlayScriptListener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production integerLiteral is exited.
func (s *BasePlayScriptListener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {}

// EnterFloatLiteral is called when production floatLiteral is entered.
func (s *BasePlayScriptListener) EnterFloatLiteral(ctx *FloatLiteralContext) {}

// ExitFloatLiteral is called when production floatLiteral is exited.
func (s *BasePlayScriptListener) ExitFloatLiteral(ctx *FloatLiteralContext) {}
