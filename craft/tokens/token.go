package tokens

var EOF = Token{Kind: Eof, Text: ""}

type Token struct {
	Kind TokenKind
	Text string
}

type TokenReader interface {
	Read() Token
	Peek() Token
	Unread()
	Position() int
}
