package lexer

import (
	"testing"

	"github.com/haapjari/books/writing-an-interpreter-in-go/interpreter/pkg/token"

	"github.com/stretchr/testify/assert"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		assert.Equal(t, tt.expectedType, tok.Type, "tests[%d] - tokentype wrong", i)
		assert.Equal(t, tt.expectedLiteral, tok.Literal, "tests[%d] - literal wrong", i)
	}
}
