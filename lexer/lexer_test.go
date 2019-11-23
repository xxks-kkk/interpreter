package lexer

import (
	"testing"

	"github.com/xxks-kkk/interpreter/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral token.TokenLiteral
	}{
		{token.ASSIGN, token.TokenLiteral("=")},
		{token.PLUS, token.TokenLiteral("+")},
		{token.LPAREN, token.TokenLiteral("(")},
		{token.RPAREN, token.TokenLiteral(")")},
		{token.LBRACE, token.TokenLiteral("{")},
		{token.RBRACE, token.TokenLiteral("}")},
		{token.COMMA, token.TokenLiteral(",")},
		{token.SEMICOLON, token.TokenLiteral(";")},
		{token.EOF, token.TokenLiteral("")},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
