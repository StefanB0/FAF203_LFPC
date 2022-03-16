package lexer

import (
	"testing"
	"os"
)

func TestLexer(t *testing.T) {

	input, _ := os.ReadFile("output/example.txt")
	l := New(string(input))

	tests := []struct {
		expectedType Tokentype
		expectedValue string
		expectedLine int
	}{
		{FUNCTION, "function", 1},
		{IDENTIFIER, "appendlazy", 1},
		{LPAR, "(", 1},
		{IDENTIFIER, "x", 1},
		{RPAR, ")", 1},
		{LBRACE, "{", 1},
		{RETURN, "return", 2},
		{IDENTIFIER, "x", 2},
		{ASSIGN, "=", 2},
		{STRING, "lazy", 2},
		{PLUS, "+", 2},
		{IDENTIFIER, "x", 2},
		{SEMICOLON, ";", 2},
		{RBRACE, "}", 3},
		{PROGRAM, "Program", 5},
		{LBRACE, "{", 5},
		{LET, "let", 7},
		{IDENTIFIER, "filename", 7},
		{ASSIGN, "=", 7},
		{STRING, "file.txt", 7},
		{SEMICOLON, ";", 7},
		{RENAME, "rename", 9},
		{LPAR, "(", 9},
		{IDENTIFIER, "filename", 9},
		{COMMA, ",", 9},
		{IDENTIFIER, "appendlazy", 9},
		{LPAR, "(", 9},
		{IDENTIFIER, "filename", 9},
		{RPAR, ")", 9},
		{RPAR, ")", 9},
		{SEMICOLON, ";", 9},
		{RBRACE, "}", 10},
		{EOF, "", 11},
	}

	for i, test := range tests {
		tk:= l.NextToken()

		if tk.Type != test.expectedType {
			t.Fatalf("Test %d failed. Expected %q but got %q", i+1, test.expectedType, tk.Type)
		}

		if tk.Value != test.expectedValue {
			t.Fatalf("Test %d failed. Expected %q but got %q", i+1, test.expectedValue, tk.Value)
		}

		if tk.Line != test.expectedLine {
			t.Fatalf("Test %d failed. Expected %q but got %q", i+1, test.expectedLine, tk.Line)
		}
	}
}
