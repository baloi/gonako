package lexer

import (
	"testing"
	"interpreter/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){}`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	} {
			{token.ASSIGN, "="},
			{token.PLUS, "+"},
	}


	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests [%d] - tokenType wrong. expected=%q got=%q",
				i, tt.expectedType, tok.Type)
		}
	}
}
