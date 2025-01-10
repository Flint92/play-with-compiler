package tokens

import "fmt"

type TokenKind int

const (
	Plus TokenKind = iota
	Minus
	Star
	Slash
	Percent

	GE
	GT
	EQ
	LE
	LT

	SemiColon
	OpenParen
	CloseParen

	Assignment

	If
	Else

	Int

	Identifier

	IntLiteral
	StringLiteral

	Eof
)

func GetTokenKindString(kind TokenKind) string {
	switch kind {
	case Plus:
		return "Plus"
	case Minus:
		return "Minus"
	case Star:
		return "Star"
	case Slash:
		return "Slash"
	case Percent:
		return "Percent"
	case GE:
		return "GE"
	case GT:
		return "GT"
	case EQ:
		return "EQ"
	case LE:
		return "LE"
	case LT:
		return "LT"
	case SemiColon:
		return "SemiColon"
	case OpenParen:
		return "OpenParen"
	case CloseParen:
		return "CloseParen"
	case Assignment:
		return "Assignment"
	case If:
		return "If"
	case Else:
		return "Else"
	case Int:
		return "Int"
	case Identifier:
		return "Identifier"
	case IntLiteral:
		return "Literal(int)"
	case StringLiteral:
		return "Literal(string)"
	case Eof:
		return "EOF"
	default:
		return fmt.Sprintf("unknown token kind: %d\n", kind)
	}
}
