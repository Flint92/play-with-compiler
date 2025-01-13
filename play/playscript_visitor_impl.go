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
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitClassDeclaration(ctx *ClassDeclarationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitClassBody(ctx *ClassBodyContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitClassBodyDeclaration(ctx *ClassBodyDeclarationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitMemberDeclaration(ctx *MemberDeclarationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitFunctionBody(ctx *FunctionBodyContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitTypeTypeOrVoid(ctx *TypeTypeOrVoidContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitQualifiedNameList(ctx *QualifiedNameListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitFormalParameters(ctx *FormalParametersContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitFormalParameterList(ctx *FormalParameterListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitFormalParameter(ctx *FormalParameterContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitLastFormalParameter(ctx *LastFormalParameterContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitVariableModifier(ctx *VariableModifierContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitQualifiedName(ctx *QualifiedNameContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitFieldDeclaration(ctx *FieldDeclarationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitConstructorDeclaration(ctx *ConstructorDeclarationContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitVariableDeclarators(ctx *VariableDeclaratorsContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitVariableDeclarator(ctx *VariableDeclaratorContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitVariableDeclaratorId(ctx *VariableDeclaratorIdContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitVariableInitializer(ctx *VariableInitializerContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitArrayInitializer(ctx *ArrayInitializerContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitTypeArgument(ctx *TypeArgumentContext) interface{} {
	//TODO implement me
	panic("implement me")
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

func (s *ScriptVisitorImpl) VisitProg(ctx *ProgContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitBlock(ctx *BlockContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitBlockStatements(ctx *BlockStatementsContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitBlockStatement(ctx *BlockStatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitStatement(ctx *StatementContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitSwitchLabel(ctx *SwitchLabelContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitForControl(ctx *ForControlContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitForInit(ctx *ForInitContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitEnhancedForControl(ctx *EnhancedForControlContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitParExpression(ctx *ParExpressionContext) interface{} {
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

func (s *ScriptVisitorImpl) VisitTypeList(ctx *TypeListContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitTypeType(ctx *TypeTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitFunctionType(ctx *FunctionTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitCreator(ctx *CreatorContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitSuperSuffix(ctx *SuperSuffixContext) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ScriptVisitorImpl) VisitArguments(ctx *ArgumentsContext) interface{} {
	//TODO implement me
	panic("implement me")
}
