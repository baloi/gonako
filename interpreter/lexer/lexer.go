package lexer

import "interpreter/token"

type Lexer struct {
	input string
	position int
	readPosition int
	ch byte
}

func New(input string) *Lexer {
	l := &Lexer {input: input,}
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	return tok
}
