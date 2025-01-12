package lexer

import (
	"fmt"
	"github.com/flint92/play-with-compiler/craft/tokens"
	"unicode"
)

type DFAState int

const (
	Initial DFAState = iota

	Int
	IdInt1
	IdInt2
	IdInt3

	Id
	GT
	GE

	Assignment

	Plus
	Minus
	Star
	Slash
	Percent

	SemiColon
	OpenParen
	CloseParen

	IntLiteral
)

type Lexer struct {
	source string
	pos    int
	tokens []tokens.Token
}

func Parse(source string) *Lexer {
	l := newLexer(source)
	l.run()
	return l
}

func newLexer(source string) *Lexer {
	return &Lexer{
		source: source,
		pos:    0,
		tokens: make([]tokens.Token, 0),
	}
}

func (l *Lexer) Read() tokens.Token {
	tok := l.Peek()
	if tok == tokens.EOF {
		return tokens.EOF
	}

	l.pos++
	return tok
}

func (l *Lexer) Peek() tokens.Token {
	if l.pos < len(l.tokens) {
		return l.tokens[l.pos]
	}
	return tokens.EOF
}

func (l *Lexer) Unread() {
	if l.pos > 0 {
		l.pos--
	}
}

func (l *Lexer) Position() int {
	return l.pos
}

func (l *Lexer) SetPosition(pos int) {
	if pos > 0 && pos < len(l.tokens) {
		l.pos = pos
	}
}

func (l *Lexer) Dump() {
	fmt.Printf("%-20s\t%-20s\n", "kind", "text")
	for _, token := range l.tokens {
		fmt.Printf("%-20s\t%-20s\n", tokens.GetTokenKindString(token.Kind), token.Text)
	}
}

func (l *Lexer) run() {
	textBuf := make([]rune, 0)
	state := Initial
	token := tokens.Token{}

	flushToken := func() {
		if len(textBuf) > 0 {
			token.Text = string(textBuf)
			l.tokens = append(l.tokens, token)
			token = tokens.Token{}
			textBuf = textBuf[0:0]
		}
	}

	initToken := func(ch rune) DFAState {
		flushToken()

		var newState DFAState
		if unicode.IsLetter(ch) {
			if ch == 'i' {
				newState = IdInt1
			} else {
				newState = Id
			}
			token.Kind = tokens.Identifier
			textBuf = append(textBuf, ch)
		} else if unicode.IsDigit(ch) {
			newState = IntLiteral
			token.Kind = tokens.IntLiteral
			textBuf = append(textBuf, ch)
		} else if ch == '>' {
			newState = GT
			token.Kind = tokens.GT
			textBuf = append(textBuf, ch)
		} else if ch == '+' {
			newState = Plus
			token.Kind = tokens.Plus
			textBuf = append(textBuf, ch)
		} else if ch == '-' {
			newState = Minus
			token.Kind = tokens.Minus
			textBuf = append(textBuf, ch)
		} else if ch == '*' {
			newState = Star
			token.Kind = tokens.Star
			textBuf = append(textBuf, ch)
		} else if ch == '/' {
			newState = Slash
			token.Kind = tokens.Slash
			textBuf = append(textBuf, ch)
		} else if ch == '%' {
			newState = Percent
			token.Kind = tokens.Percent
			textBuf = append(textBuf, ch)
		} else if ch == ';' {
			newState = SemiColon
			token.Kind = tokens.SemiColon
			textBuf = append(textBuf, ch)
		} else if ch == '(' {
			newState = OpenParen
			token.Kind = tokens.OpenParen
			textBuf = append(textBuf, ch)
		} else if ch == ')' {
			newState = CloseParen
			token.Kind = tokens.CloseParen
			textBuf = append(textBuf, ch)
		} else if ch == '=' {
			newState = Assignment
			token.Kind = tokens.Assignment
			textBuf = append(textBuf, ch)
		} else {
			newState = Initial // skip all unknown patterns
		}
		return newState
	}

	for _, ch := range []rune(l.source) {
		switch state {
		case Initial:
			state = initToken(ch)
		case Id:
			if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
				textBuf = append(textBuf, ch)
			} else {
				state = initToken(ch)
			}
		case IdInt1:
			if ch == 'n' {
				state = IdInt2
				textBuf = append(textBuf, ch)
			} else if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
				state = Id
				textBuf = append(textBuf, ch)
			} else {
				state = initToken(ch)
			}
		case IdInt2:
			if ch == 't' {
				state = IdInt3
				textBuf = append(textBuf, ch)
			} else if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
				state = Id
				textBuf = append(textBuf, ch)
			} else {
				state = initToken(ch)
			}
		case IdInt3:
			if unicode.IsSpace(ch) {
				state = Int
			} else {
				state = Id
				textBuf = append(textBuf, ch)
			}
		case Int:
			token.Kind = tokens.Int
			state = initToken(ch)
		case GT:
			if ch == '=' {
				token.Kind = tokens.GE
				state = GE
				textBuf = append(textBuf, ch)
			} else {
				state = initToken(ch)
			}
		case GE:
			state = initToken(ch)
		case Assignment:
			state = initToken(ch)
		case Plus:
			state = initToken(ch)
		case Minus:
			state = initToken(ch)
		case Star:
			state = initToken(ch)
		case Slash:
			state = initToken(ch)
		case Percent:
			state = initToken(ch)
		case SemiColon:
			state = initToken(ch)
		case OpenParen:
			state = initToken(ch)
		case CloseParen:
			state = initToken(ch)
		case IntLiteral:
			if unicode.IsDigit(ch) {
				textBuf = append(textBuf, ch)
			} else {
				state = initToken(ch)
			}
		default:
			state = initToken(ch)
		}
	}

	flushToken()
}
