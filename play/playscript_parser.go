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
		"expression", "primary", "integerLiteral",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 113, 30, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 16, 8, 0, 10, 0, 12, 0, 19, 9, 0,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 26, 8, 1, 1, 2, 1, 2, 1, 2, 0, 1, 0,
		3, 0, 2, 4, 0, 3, 2, 0, 89, 90, 94, 94, 1, 0, 87, 88, 1, 0, 53, 56, 29,
		0, 6, 1, 0, 0, 0, 2, 25, 1, 0, 0, 0, 4, 27, 1, 0, 0, 0, 6, 7, 6, 0, -1,
		0, 7, 8, 3, 2, 1, 0, 8, 17, 1, 0, 0, 0, 9, 10, 10, 2, 0, 0, 10, 11, 7,
		0, 0, 0, 11, 16, 3, 0, 0, 3, 12, 13, 10, 1, 0, 0, 13, 14, 7, 1, 0, 0, 14,
		16, 3, 0, 0, 2, 15, 9, 1, 0, 0, 0, 15, 12, 1, 0, 0, 0, 16, 19, 1, 0, 0,
		0, 17, 15, 1, 0, 0, 0, 17, 18, 1, 0, 0, 0, 18, 1, 1, 0, 0, 0, 19, 17, 1,
		0, 0, 0, 20, 21, 5, 63, 0, 0, 21, 22, 3, 0, 0, 0, 22, 23, 5, 64, 0, 0,
		23, 26, 1, 0, 0, 0, 24, 26, 3, 4, 2, 0, 25, 20, 1, 0, 0, 0, 25, 24, 1,
		0, 0, 0, 26, 3, 1, 0, 0, 0, 27, 28, 7, 2, 0, 0, 28, 5, 1, 0, 0, 0, 3, 15,
		17, 25,
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
	PlayScriptParserRULE_expression     = 0
	PlayScriptParserRULE_primary        = 1
	PlayScriptParserRULE_integerLiteral = 2
)

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
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode
	MOD() antlr.TerminalNode
	ADD() antlr.TerminalNode
	SUB() antlr.TerminalNode

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
	_startState := 0
	p.EnterRecursionRule(localctx, 0, PlayScriptParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(7)
		p.Primary()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(15)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(9)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(10)

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
					p.SetState(11)
					p.expression(3)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(12)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(13)

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
					p.SetState(14)
					p.expression(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
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
	IntegerLiteral() IIntegerLiteralContext

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

func (s *PrimaryContext) IntegerLiteral() IIntegerLiteralContext {
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
	p.EnterRule(localctx, 2, PlayScriptParserRULE_primary)
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PlayScriptParserLPAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(20)
			p.Match(PlayScriptParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(21)
			p.expression(0)
		}
		{
			p.SetState(22)
			p.Match(PlayScriptParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PlayScriptParserDECIMAL_LITERAL, PlayScriptParserHEX_LITERAL, PlayScriptParserOCT_LITERAL, PlayScriptParserBINARY_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(24)
			p.IntegerLiteral()
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
	p.EnterRule(localctx, 4, PlayScriptParserRULE_integerLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(27)
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

func (p *PlayScriptParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
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
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
