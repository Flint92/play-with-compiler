// Code generated from PlayScript.g4 by ANTLR 4.13.2. DO NOT EDIT.

package play // PlayScript
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type PlayScriptParser struct {
	*antlr.BaseParser
}

var PlayScriptParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func playscriptParserInit() {
	staticData := &PlayScriptParserStaticData
	staticData.LiteralNames = []string{
		"", "'abstract'", "'assert'", "'boolean'", "'break'", "'byte'", "'case'",
		"'catch'", "'char'", "'class'", "'const'", "'continue'", "'default'",
		"'do'", "'double'", "'else'", "'enum'", "'extends'", "'final'", "'finally'",
		"'float'", "'for'", "'if'", "'goto'", "'implements'", "'import'", "'instanceof'",
		"'int'", "'interface'", "'long'", "'native'", "'new'", "'package'",
		"'private'", "'protected'", "'public'", "'return'", "'short'", "'static'",
		"'strictfp'", "'super'", "'switch'", "'synchronized'", "'this'", "'throw'",
		"'throws'", "'transient'", "'try'", "'void'", "'volatile'", "'while'",
		"'function'", "'string'", "", "", "", "", "", "", "", "", "", "'null'",
		"'('", "')'", "'{'", "'}'", "'['", "']'", "';'", "','", "'.'", "'='",
		"'>'", "'<'", "'!'", "'~'", "'?'", "':'", "'=='", "'<='", "'>='", "'!='",
		"'&&'", "'||'", "'++'", "'--'", "'+'", "'-'", "'*'", "'/'", "'&'", "'|'",
		"'^'", "'%'", "'+='", "'-='", "'*='", "'/='", "'&='", "'|='", "'^='",
		"'%='", "'<<='", "'>>='", "'>>>='", "'->'", "'::'", "'@'", "'...'",
	}
	staticData.SymbolicNames = []string{
		"", "ABSTRACT", "ASSERT", "BOOLEAN", "BREAK", "BYTE", "CASE", "CATCH",
		"CHAR", "CLASS", "CONST", "CONTINUE", "DEFAULT", "DO", "DOUBLE", "ELSE",
		"ENUM", "EXTENDS", "FINAL", "FINALLY", "FLOAT", "FOR", "IF", "GOTO",
		"IMPLEMENTS", "IMPORT", "INSTANCEOF", "INT", "INTERFACE", "LONG", "NATIVE",
		"NEW", "PACKAGE", "PRIVATE", "PROTECTED", "PUBLIC", "RETURN", "SHORT",
		"STATIC", "STRICTFP", "SUPER", "SWITCH", "SYNCHRONIZED", "THIS", "THROW",
		"THROWS", "TRANSIENT", "TRY", "VOID", "VOLATILE", "WHILE", "FUNCTION",
		"STRING", "DECIMAL_LITERAL", "HEX_LITERAL", "OCT_LITERAL", "BINARY_LITERAL",
		"FLOAT_LITERAL", "HEX_FLOAT_LITERAL", "BOOL_LITERAL", "CHAR_LITERAL",
		"STRING_LITERAL", "NULL_LITERAL", "LPAREN", "RPAREN", "LBRACE", "RBRACE",
		"LBRACK", "RBRACK", "SEMI", "COMMA", "DOT", "ASSIGN", "GT", "LT", "BANG",
		"TILDE", "QUESTION", "COLON", "EQUAL", "LE", "GE", "NOTEQUAL", "AND",
		"OR", "INC", "DEC", "ADD", "SUB", "MUL", "DIV", "BITAND", "BITOR", "CARET",
		"MOD", "ADD_ASSIGN", "SUB_ASSIGN", "MUL_ASSIGN", "DIV_ASSIGN", "AND_ASSIGN",
		"OR_ASSIGN", "XOR_ASSIGN", "MOD_ASSIGN", "LSHIFT_ASSIGN", "RSHIFT_ASSIGN",
		"URSHIFT_ASSIGN", "ARROW", "COLONCOLON", "AT", "ELLIPSIS", "WS", "COMMENT",
		"LINE_COMMENT", "IDENTIFIER",
	}
	staticData.RuleNames = []string{
		"bracedExpression", "expression", "primary", "expressionList", "functionCall",
		"literal", "integerLiteral", "floatLiteral",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 113, 85, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 3, 1, 25, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1,
		45, 8, 1, 10, 1, 12, 1, 48, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3,
		2, 56, 8, 2, 1, 3, 1, 3, 1, 3, 5, 3, 61, 8, 3, 10, 3, 12, 3, 64, 9, 3,
		1, 4, 1, 4, 1, 4, 3, 4, 69, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 3, 5, 79, 8, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 0, 1, 2, 8, 0, 2,
		4, 6, 8, 10, 12, 14, 0, 6, 2, 0, 89, 90, 94, 94, 1, 0, 87, 88, 2, 0, 73,
		74, 80, 81, 2, 0, 79, 79, 82, 82, 1, 0, 53, 56, 1, 0, 57, 58, 92, 0, 16,
		1, 0, 0, 0, 2, 24, 1, 0, 0, 0, 4, 55, 1, 0, 0, 0, 6, 57, 1, 0, 0, 0, 8,
		65, 1, 0, 0, 0, 10, 78, 1, 0, 0, 0, 12, 80, 1, 0, 0, 0, 14, 82, 1, 0, 0,
		0, 16, 17, 5, 65, 0, 0, 17, 18, 5, 72, 0, 0, 18, 19, 3, 2, 1, 0, 19, 20,
		5, 66, 0, 0, 20, 1, 1, 0, 0, 0, 21, 22, 6, 1, -1, 0, 22, 25, 3, 4, 2, 0,
		23, 25, 3, 8, 4, 0, 24, 21, 1, 0, 0, 0, 24, 23, 1, 0, 0, 0, 25, 46, 1,
		0, 0, 0, 26, 27, 10, 6, 0, 0, 27, 28, 7, 0, 0, 0, 28, 45, 3, 2, 1, 7, 29,
		30, 10, 5, 0, 0, 30, 31, 7, 1, 0, 0, 31, 45, 3, 2, 1, 6, 32, 33, 10, 4,
		0, 0, 33, 34, 7, 2, 0, 0, 34, 45, 3, 2, 1, 5, 35, 36, 10, 3, 0, 0, 36,
		37, 7, 3, 0, 0, 37, 45, 3, 2, 1, 4, 38, 39, 10, 2, 0, 0, 39, 40, 5, 83,
		0, 0, 40, 45, 3, 2, 1, 3, 41, 42, 10, 1, 0, 0, 42, 43, 5, 84, 0, 0, 43,
		45, 3, 2, 1, 2, 44, 26, 1, 0, 0, 0, 44, 29, 1, 0, 0, 0, 44, 32, 1, 0, 0,
		0, 44, 35, 1, 0, 0, 0, 44, 38, 1, 0, 0, 0, 44, 41, 1, 0, 0, 0, 45, 48,
		1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 3, 1, 0, 0, 0,
		48, 46, 1, 0, 0, 0, 49, 50, 5, 63, 0, 0, 50, 51, 3, 2, 1, 0, 51, 52, 5,
		64, 0, 0, 52, 56, 1, 0, 0, 0, 53, 56, 3, 10, 5, 0, 54, 56, 5, 113, 0, 0,
		55, 49, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 55, 54, 1, 0, 0, 0, 56, 5, 1, 0,
		0, 0, 57, 62, 3, 2, 1, 0, 58, 59, 5, 70, 0, 0, 59, 61, 3, 2, 1, 0, 60,
		58, 1, 0, 0, 0, 61, 64, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0,
		0, 63, 7, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 65, 66, 5, 113, 0, 0, 66, 68,
		5, 63, 0, 0, 67, 69, 3, 6, 3, 0, 68, 67, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0,
		69, 70, 1, 0, 0, 0, 70, 71, 5, 64, 0, 0, 71, 9, 1, 0, 0, 0, 72, 79, 3,
		12, 6, 0, 73, 79, 3, 14, 7, 0, 74, 79, 5, 60, 0, 0, 75, 79, 5, 61, 0, 0,
		76, 79, 5, 59, 0, 0, 77, 79, 5, 62, 0, 0, 78, 72, 1, 0, 0, 0, 78, 73, 1,
		0, 0, 0, 78, 74, 1, 0, 0, 0, 78, 75, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 78,
		77, 1, 0, 0, 0, 79, 11, 1, 0, 0, 0, 80, 81, 7, 4, 0, 0, 81, 13, 1, 0, 0,
		0, 82, 83, 7, 5, 0, 0, 83, 15, 1, 0, 0, 0, 7, 24, 44, 46, 55, 62, 68, 78,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// PlayScriptParserInit initializes any static state used to implement PlayScriptParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPlayScriptParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PlayScriptParserInit() {
	staticData := &PlayScriptParserStaticData
	staticData.once.Do(playscriptParserInit)
}

// NewPlayScriptParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPlayScriptParser(input antlr.TokenStream) *PlayScriptParser {
	PlayScriptParserInit()
	this := new(PlayScriptParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &PlayScriptParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "PlayScript.g4"

	return this
}

// PlayScriptParser tokens.
const (
	PlayScriptParserEOF               = antlr.TokenEOF
	PlayScriptParserABSTRACT          = 1
	PlayScriptParserASSERT            = 2
	PlayScriptParserBOOLEAN           = 3
	PlayScriptParserBREAK             = 4
	PlayScriptParserBYTE              = 5
	PlayScriptParserCASE              = 6
	PlayScriptParserCATCH             = 7
	PlayScriptParserCHAR              = 8
	PlayScriptParserCLASS             = 9
	PlayScriptParserCONST             = 10
	PlayScriptParserCONTINUE          = 11
	PlayScriptParserDEFAULT           = 12
	PlayScriptParserDO                = 13
	PlayScriptParserDOUBLE            = 14
	PlayScriptParserELSE              = 15
	PlayScriptParserENUM              = 16
	PlayScriptParserEXTENDS           = 17
	PlayScriptParserFINAL             = 18
	PlayScriptParserFINALLY           = 19
	PlayScriptParserFLOAT             = 20
	PlayScriptParserFOR               = 21
	PlayScriptParserIF                = 22
	PlayScriptParserGOTO              = 23
	PlayScriptParserIMPLEMENTS        = 24
	PlayScriptParserIMPORT            = 25
	PlayScriptParserINSTANCEOF        = 26
	PlayScriptParserINT               = 27
	PlayScriptParserINTERFACE         = 28
	PlayScriptParserLONG              = 29
	PlayScriptParserNATIVE            = 30
	PlayScriptParserNEW               = 31
	PlayScriptParserPACKAGE           = 32
	PlayScriptParserPRIVATE           = 33
	PlayScriptParserPROTECTED         = 34
	PlayScriptParserPUBLIC            = 35
	PlayScriptParserRETURN            = 36
	PlayScriptParserSHORT             = 37
	PlayScriptParserSTATIC            = 38
	PlayScriptParserSTRICTFP          = 39
	PlayScriptParserSUPER             = 40
	PlayScriptParserSWITCH            = 41
	PlayScriptParserSYNCHRONIZED      = 42
	PlayScriptParserTHIS              = 43
	PlayScriptParserTHROW             = 44
	PlayScriptParserTHROWS            = 45
	PlayScriptParserTRANSIENT         = 46
	PlayScriptParserTRY               = 47
	PlayScriptParserVOID              = 48
	PlayScriptParserVOLATILE          = 49
	PlayScriptParserWHILE             = 50
	PlayScriptParserFUNCTION          = 51
	PlayScriptParserSTRING            = 52
	PlayScriptParserDECIMAL_LITERAL   = 53
	PlayScriptParserHEX_LITERAL       = 54
	PlayScriptParserOCT_LITERAL       = 55
	PlayScriptParserBINARY_LITERAL    = 56
	PlayScriptParserFLOAT_LITERAL     = 57
	PlayScriptParserHEX_FLOAT_LITERAL = 58
	PlayScriptParserBOOL_LITERAL      = 59
	PlayScriptParserCHAR_LITERAL      = 60
	PlayScriptParserSTRING_LITERAL    = 61
	PlayScriptParserNULL_LITERAL      = 62
	PlayScriptParserLPAREN            = 63
	PlayScriptParserRPAREN            = 64
	PlayScriptParserLBRACE            = 65
	PlayScriptParserRBRACE            = 66
	PlayScriptParserLBRACK            = 67
	PlayScriptParserRBRACK            = 68
	PlayScriptParserSEMI              = 69
	PlayScriptParserCOMMA             = 70
	PlayScriptParserDOT               = 71
	PlayScriptParserASSIGN            = 72
	PlayScriptParserGT                = 73
	PlayScriptParserLT                = 74
	PlayScriptParserBANG              = 75
	PlayScriptParserTILDE             = 76
	PlayScriptParserQUESTION          = 77
	PlayScriptParserCOLON             = 78
	PlayScriptParserEQUAL             = 79
	PlayScriptParserLE                = 80
	PlayScriptParserGE                = 81
	PlayScriptParserNOTEQUAL          = 82
	PlayScriptParserAND               = 83
	PlayScriptParserOR                = 84
	PlayScriptParserINC               = 85
	PlayScriptParserDEC               = 86
	PlayScriptParserADD               = 87
	PlayScriptParserSUB               = 88
	PlayScriptParserMUL               = 89
	PlayScriptParserDIV               = 90
	PlayScriptParserBITAND            = 91
	PlayScriptParserBITOR             = 92
	PlayScriptParserCARET             = 93
	PlayScriptParserMOD               = 94
	PlayScriptParserADD_ASSIGN        = 95
	PlayScriptParserSUB_ASSIGN        = 96
	PlayScriptParserMUL_ASSIGN        = 97
	PlayScriptParserDIV_ASSIGN        = 98
	PlayScriptParserAND_ASSIGN        = 99
	PlayScriptParserOR_ASSIGN         = 100
	PlayScriptParserXOR_ASSIGN        = 101
	PlayScriptParserMOD_ASSIGN        = 102
	PlayScriptParserLSHIFT_ASSIGN     = 103
	PlayScriptParserRSHIFT_ASSIGN     = 104
	PlayScriptParserURSHIFT_ASSIGN    = 105
	PlayScriptParserARROW             = 106
	PlayScriptParserCOLONCOLON        = 107
	PlayScriptParserAT                = 108
	PlayScriptParserELLIPSIS          = 109
	PlayScriptParserWS                = 110
	PlayScriptParserCOMMENT           = 111
	PlayScriptParserLINE_COMMENT      = 112
	PlayScriptParserIDENTIFIER        = 113
)

// PlayScriptParser rules.
const (
	PlayScriptParserRULE_bracedExpression = 0
	PlayScriptParserRULE_expression       = 1
	PlayScriptParserRULE_primary          = 2
	PlayScriptParserRULE_expressionList   = 3
	PlayScriptParserRULE_functionCall     = 4
	PlayScriptParserRULE_literal          = 5
	PlayScriptParserRULE_integerLiteral   = 6
	PlayScriptParserRULE_floatLiteral     = 7
)

// IBracedExpressionContext is an interface to support dynamic dispatch.
type IBracedExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext
	RBRACE() antlr.TerminalNode

	// IsBracedExpressionContext differentiates from other interfaces.
	IsBracedExpressionContext()
}

type BracedExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBracedExpressionContext() *BracedExpressionContext {
	var p = new(BracedExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlayScriptParserRULE_bracedExpression
	return p
}

func InitEmptyBracedExpressionContext(p *BracedExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlayScriptParserRULE_bracedExpression
}

func (*BracedExpressionContext) IsBracedExpressionContext() {}

func NewBracedExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BracedExpressionContext {
	var p = new(BracedExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_bracedExpression

	return p
}

func (s *BracedExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *BracedExpressionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLBRACE, 0)
}

func (s *BracedExpressionContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserASSIGN, 0)
}

func (s *BracedExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BracedExpressionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRBRACE, 0)
}

func (s *BracedExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracedExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BracedExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterBracedExpression(s)
	}
}

func (s *BracedExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitBracedExpression(s)
	}
}

func (s *BracedExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitBracedExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) BracedExpression() (localctx IBracedExpressionContext) {
	localctx = NewBracedExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PlayScriptParserRULE_bracedExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.Match(PlayScriptParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(17)
		p.Match(PlayScriptParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(18)
		p.expression(0)
	}
	{
		p.SetState(19)
		p.Match(PlayScriptParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetBop returns the bop token.
	GetBop() antlr.Token

	// SetBop sets the bop token.
	SetBop(antlr.Token)

	// Getter signatures
	Primary() IPrimaryContext
	FunctionCall() IFunctionCallContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode
	MOD() antlr.TerminalNode
	ADD() antlr.TerminalNode
	SUB() antlr.TerminalNode
	LE() antlr.TerminalNode
	GE() antlr.TerminalNode
	GT() antlr.TerminalNode
	LT() antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	NOTEQUAL() antlr.TerminalNode
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	bop    antlr.Token
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlayScriptParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlayScriptParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetBop() antlr.Token { return s.bop }

func (s *ExpressionContext) SetBop(v antlr.Token) { s.bop = v }

func (s *ExpressionContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *ExpressionContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) MUL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserMUL, 0)
}

func (s *ExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserDIV, 0)
}

func (s *ExpressionContext) MOD() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserMOD, 0)
}

func (s *ExpressionContext) ADD() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserADD, 0)
}

func (s *ExpressionContext) SUB() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserSUB, 0)
}

func (s *ExpressionContext) LE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLE, 0)
}

func (s *ExpressionContext) GE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserGE, 0)
}

func (s *ExpressionContext) GT() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserGT, 0)
}

func (s *ExpressionContext) LT() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLT, 0)
}

func (s *ExpressionContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserEQUAL, 0)
}

func (s *ExpressionContext) NOTEQUAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserNOTEQUAL, 0)
}

func (s *ExpressionContext) AND() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserAND, 0)
}

func (s *ExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserOR, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *PlayScriptParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, PlayScriptParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(22)
			p.Primary()
		}

	case 2:
		{
			p.SetState(23)
			p.FunctionCall()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(44)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(26)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(27)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64((_la-89)) & ^0x3f) == 0 && ((int64(1)<<(_la-89))&35) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(28)
					p.expression(7)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(29)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(30)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == PlayScriptParserADD || _la == PlayScriptParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(31)
					p.expression(6)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(32)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(33)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64((_la-73)) & ^0x3f) == 0 && ((int64(1)<<(_la-73))&387) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(34)
					p.expression(5)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(35)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(36)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == PlayScriptParserEQUAL || _la == PlayScriptParserNOTEQUAL) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(37)
					p.expression(4)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(38)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(39)

					var _m = p.Match(PlayScriptParserAND)

					localctx.(*ExpressionContext).bop = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(40)
					p.expression(3)
				}

			case 6:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(41)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(42)

					var _m = p.Match(PlayScriptParserOR)

					localctx.(*ExpressionContext).bop = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(43)
					p.expression(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode
	Literal() ILiteralContext
	IDENTIFIER() antlr.TerminalNode

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlayScriptParserRULE_primary
	return p
}

func InitEmptyPrimaryContext(p *PrimaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlayScriptParserRULE_primary
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLPAREN, 0)
}

func (s *PrimaryContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrimaryContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRPAREN, 0)
}

func (s *PrimaryContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *PrimaryContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserIDENTIFIER, 0)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterPrimary(s)
	}
}

func (s *PrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitPrimary(s)
	}
}

func (s *PrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PlayScriptParserRULE_primary)
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PlayScriptParserLPAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(49)
			p.Match(PlayScriptParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.expression(0)
		}
		{
			p.SetState(51)
			p.Match(PlayScriptParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PlayScriptParserDECIMAL_LITERAL, PlayScriptParserHEX_LITERAL, PlayScriptParserOCT_LITERAL, PlayScriptParserBINARY_LITERAL, PlayScriptParserFLOAT_LITERAL, PlayScriptParserHEX_FLOAT_LITERAL, PlayScriptParserBOOL_LITERAL, PlayScriptParserCHAR_LITERAL, PlayScriptParserSTRING_LITERAL, PlayScriptParserNULL_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(53)
			p.Literal()
		}

	case PlayScriptParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(54)
			p.Match(PlayScriptParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlayScriptParserRULE_expressionList
	return p
}

func InitEmptyExpressionListContext(p *ExpressionListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlayScriptParserRULE_expressionList
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserCOMMA)
}

func (s *ExpressionListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCOMMA, i)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterExpressionList(s)
	}
}

func (s *ExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitExpressionList(s)
	}
}

func (s *ExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PlayScriptParserRULE_expressionList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.expression(0)
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PlayScriptParserCOMMA {
		{
			p.SetState(58)
			p.Match(PlayScriptParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(59)
			p.expression(0)
		}

		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	ExpressionList() IExpressionListContext

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlayScriptParserRULE_functionCall
	return p
}

func InitEmptyFunctionCallContext(p *FunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlayScriptParserRULE_functionCall
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserIDENTIFIER, 0)
}

func (s *FunctionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLPAREN, 0)
}

func (s *FunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRPAREN, 0)
}

func (s *FunctionCallContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PlayScriptParserRULE_functionCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.Match(PlayScriptParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(66)
		p.Match(PlayScriptParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-53)) & ^0x3f) == 0 && ((int64(1)<<(_la-53))&1152921504606849023) != 0 {
		{
			p.SetState(67)
			p.ExpressionList()
		}

	}
	{
		p.SetState(70)
		p.Match(PlayScriptParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IntegerLiteral() IIntegerLiteralContext
	FloatLiteral() IFloatLiteralContext
	CHAR_LITERAL() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode
	BOOL_LITERAL() antlr.TerminalNode
	NULL_LITERAL() antlr.TerminalNode

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlayScriptParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlayScriptParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) IntegerLiteral() IIntegerLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerLiteralContext)
}

func (s *LiteralContext) FloatLiteral() IFloatLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloatLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloatLiteralContext)
}

func (s *LiteralContext) CHAR_LITERAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCHAR_LITERAL, 0)
}

func (s *LiteralContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserSTRING_LITERAL, 0)
}

func (s *LiteralContext) BOOL_LITERAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserBOOL_LITERAL, 0)
}

func (s *LiteralContext) NULL_LITERAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserNULL_LITERAL, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PlayScriptParserRULE_literal)
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PlayScriptParserDECIMAL_LITERAL, PlayScriptParserHEX_LITERAL, PlayScriptParserOCT_LITERAL, PlayScriptParserBINARY_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(72)
			p.IntegerLiteral()
		}

	case PlayScriptParserFLOAT_LITERAL, PlayScriptParserHEX_FLOAT_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(73)
			p.FloatLiteral()
		}

	case PlayScriptParserCHAR_LITERAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(74)
			p.Match(PlayScriptParserCHAR_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PlayScriptParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(75)
			p.Match(PlayScriptParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PlayScriptParserBOOL_LITERAL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(76)
			p.Match(PlayScriptParserBOOL_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PlayScriptParserNULL_LITERAL:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(77)
			p.Match(PlayScriptParserNULL_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIntegerLiteralContext is an interface to support dynamic dispatch.
type IIntegerLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DECIMAL_LITERAL() antlr.TerminalNode
	HEX_LITERAL() antlr.TerminalNode
	OCT_LITERAL() antlr.TerminalNode
	BINARY_LITERAL() antlr.TerminalNode

	// IsIntegerLiteralContext differentiates from other interfaces.
	IsIntegerLiteralContext()
}

type IntegerLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerLiteralContext() *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlayScriptParserRULE_integerLiteral
	return p
}

func InitEmptyIntegerLiteralContext(p *IntegerLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlayScriptParserRULE_integerLiteral
}

func (*IntegerLiteralContext) IsIntegerLiteralContext() {}

func NewIntegerLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_integerLiteral

	return p
}

func (s *IntegerLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerLiteralContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserDECIMAL_LITERAL, 0)
}

func (s *IntegerLiteralContext) HEX_LITERAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserHEX_LITERAL, 0)
}

func (s *IntegerLiteralContext) OCT_LITERAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserOCT_LITERAL, 0)
}

func (s *IntegerLiteralContext) BINARY_LITERAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserBINARY_LITERAL, 0)
}

func (s *IntegerLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterIntegerLiteral(s)
	}
}

func (s *IntegerLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitIntegerLiteral(s)
	}
}

func (s *IntegerLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitIntegerLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) IntegerLiteral() (localctx IIntegerLiteralContext) {
	localctx = NewIntegerLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PlayScriptParserRULE_integerLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135107988821114880) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFloatLiteralContext is an interface to support dynamic dispatch.
type IFloatLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FLOAT_LITERAL() antlr.TerminalNode
	HEX_FLOAT_LITERAL() antlr.TerminalNode

	// IsFloatLiteralContext differentiates from other interfaces.
	IsFloatLiteralContext()
}

type FloatLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatLiteralContext() *FloatLiteralContext {
	var p = new(FloatLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlayScriptParserRULE_floatLiteral
	return p
}

func InitEmptyFloatLiteralContext(p *FloatLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlayScriptParserRULE_floatLiteral
}

func (*FloatLiteralContext) IsFloatLiteralContext() {}

func NewFloatLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatLiteralContext {
	var p = new(FloatLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_floatLiteral

	return p
}

func (s *FloatLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatLiteralContext) FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserFLOAT_LITERAL, 0)
}

func (s *FloatLiteralContext) HEX_FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserHEX_FLOAT_LITERAL, 0)
}

func (s *FloatLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitFloatLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) FloatLiteral() (localctx IFloatLiteralContext) {
	localctx = NewFloatLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PlayScriptParserRULE_floatLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PlayScriptParserFLOAT_LITERAL || _la == PlayScriptParserHEX_FLOAT_LITERAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *PlayScriptParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PlayScriptParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
