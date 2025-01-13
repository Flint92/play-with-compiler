package play

import "github.com/antlr4-go/antlr/v4"

func Eval(script string) int64 {
	// Create the lexer
	lexer := NewPlayScriptLexer(antlr.NewInputStream(script))

	// Create a buffer of tokens between the lexer and parser
	stream := antlr.NewCommonTokenStream(lexer, 0)

	// Create the parser
	parser := NewPlayScriptParser(stream)

	// Parse the input
	tree := parser.Expression()

	// Create the visitor
	visitor := NewPlayScriptVisitor()

	// Visit the parse tree
	return visitor.Visit(tree).(int64)
}
