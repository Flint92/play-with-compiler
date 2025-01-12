package ast

import "fmt"

type NodeType int

const (
	Program NodeType = iota

	IntDeclaration
	ExpressionStmt
	AssignmentStmt

	Primary
	Multiplicative
	Additive

	Identifier
	IntLiteral
)

func GetNodeTypeString(nType NodeType) string {
	switch nType {
	case Program:
		return "Program"
	case IntDeclaration:
		return "IntDeclaration"
	case ExpressionStmt:
		return "ExpressionStmt"
	case AssignmentStmt:
		return "AssignmentStmt"
	case Primary:
		return "Primary"
	case Multiplicative:
		return "Multiplicative"
	case Additive:
		return "Additive"
	case Identifier:
		return "Identifier"
	case IntLiteral:
		return "IntLiteral"
	default:
		return fmt.Sprintf("unknown node type %d\n", nType)
	}
}
