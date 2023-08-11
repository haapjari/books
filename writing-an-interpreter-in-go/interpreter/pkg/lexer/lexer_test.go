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
		{token.EOF, ""},
		{token.EOF, ""},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		assert.Equal(t, tt.expectedType, tok.Type, "tests[%d] - tokentype wrong", i)
		assert.Equal(t, tt.expectedLiteral, tok.Literal, "tests[%d] - literal wrong", i)
	}
}
func TestLexer_NextToken(t *testing.T) {
	input := `
	
	let five = 5;
	let ten = 10;
	
	let add = fn(x, y) {
	x + y;
	};

	let result = add(five, ten);
	!=/*5;
	5 < 10 > 5;

	`

	// TODO: Fix this test.
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},

		// 	x + y;
		// 	};
		//
		// 	let result = add(five, ten);
		// 	!=/*5;
		// 	5 < 10 > 5;

	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		assert.Equal(t, tt.expectedType, tok.Type, "tests[%d] - tokentype wrong", i)
		assert.Equal(t, tt.expectedLiteral, tok.Literal, "tests[%d] - literal wrong", i)
	}
}
