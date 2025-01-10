package lexer

import (
	"testing"
)

func TestParse1(t *testing.T) {
	source := "age >= 45"
	l := Parse(source)
	l.Dump()
}

func TestParse2(t *testing.T) {
	source := "int age = 45"
	l := Parse(source)
	l.Dump()
}

func TestParse3(t *testing.T) {
	source := "intA = 10"
	l := Parse(source)
	l.Dump()
}

func TestParse4(t *testing.T) {
	source := "2 + 3 * 5"
	l := Parse(source)
	l.Dump()
}
