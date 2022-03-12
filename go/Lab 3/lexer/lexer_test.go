package lexer

import (
	"testing"

	"lab3/token"
)

func TestNextToken(t *testing.T) {
	input := `function appendlazy(x){
		return x = "lazy" + x;
	}

	Program{

	let filename = "file.txt";

	rename(filename, appendlazy(filename));
	}
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.FUNCTION, "function"},
		{token.IDENT, "appendlazy"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.STRING, "lazy"},
		{token.PLUS, "+"},
		{token.IDENT, "x"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.PROGRAM, "Program"},
		{token.LBRACE, "{"},
		{token.LET, "let"},
		{token.IDENT, "filename"},
		{token.ASSIGN, "="},
		{token.STRING, "file.txt"},
		{token.SEMICOLON, ";"},
		{token.RENAME, "rename"},
		{token.LPAREN, "("},
		{token.IDENT, "filename"},
		{token.COMMA, ","},
		{token.IDENT, "appendlazy"},
		{token.LPAREN, "("},
		{token.IDENT, "filename"},
		{token.RPAREN, ")"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
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
