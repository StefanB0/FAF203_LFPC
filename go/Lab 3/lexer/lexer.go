package lexer

import (
	"strings"
)

type Lexer struct {
	input 		string
	ch 			byte
	pointer 	int
	peekpointer int
	line 		int
	tok			Token
}

func New(input string) *Lexer {
	l := &Lexer{input: input,
				pointer: 0,
				peekpointer: 0,
				line: 1,}
	l.ch = input[0]
	return l
}

func (l *Lexer) getWord() string {
	for isLetter(l.input[l.peekpointer]) || isDigit(l.input[l.peekpointer]) {
		l.peekpointer++
	}
	word:= l.input[l.pointer : l.peekpointer]
	l.peekpointer--
	l.pointer = l.peekpointer
	return word
}

func (l *Lexer) getstring() {
	l.tok.Type = STRING
	l.tok.Line = l.line

	for l.input[l.peekpointer] != '"'{
		l.peekpointer++
	}

	l.tok.Value = l.input[l.pointer : l.peekpointer]

	l.pointer = l.peekpointer
}

func (l *Lexer) getnumber() {
	for (l.input[l.peekpointer] <= '9' && l.input[l.peekpointer] >= '0') ||
			l.input[l.peekpointer] == '.'{
		l.peekpointer++
	}


	l.tok.Value = l.input[l.pointer : l.peekpointer]
	l.tok.Line = l.line
	l.pointer = l.peekpointer

	nr:= l.tok.Value

	if strings.Contains(nr, ".") {
		l.tok.Type = FLOAT

	} else {
		l.tok.Type = INTEGER
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) advance() {
	l.pointer++
	l.peekpointer++
	if l.pointer >= len(l.input) {
		l.ch = 0
	} else {
	l.ch = l.input[l.pointer]
	}
}

func (l *Lexer) peek() byte{
	return l.input[l.pointer+1]
}

func (l *Lexer) skipWhitespaces() {
	if l.pointer >= len(l.input) {
		return
	}

	for l.input[l.pointer] == ' '|| l.input[l.pointer] == '\t' || l.input[l.pointer] == '\n' || l.input[l.pointer] == '\r' {
		if l.input[l.pointer] == '\n'{
			l.line++
		}

		l.advance()
		if l.pointer >= len(l.input) {
			return
		}
	}

}

func (l *Lexer) NextToken() Token {

	l.skipWhitespaces()

	switch l.ch {
	case '=':
		if l.peek() == '=' {
			l.tok = Token{EQUAL, "==", l.line}
			l.advance()
		} else {
			l.tok = Token{ASSIGN, "=", l.line}
		}
	case '+':
		l.tok = Token{PLUS, "+", l.line}
	case '-':
		l.tok = Token{MINUS, "-", l.line}
	case '!':
		if l.peek() == '=' {
			l.tok = Token{NOT_EQUAL, "!=", l.line}
			l.advance()
		} else {
			l.tok = Token{BANG, "!", l.line}
		}
	case '"':
		l.advance()
		l.getstring()
	case '/':
		l.tok = Token{SLASH, "/", l.line}
	case '*':
		l.tok = Token{ASTERISK, "*", l.line}
	case '<':
		if l.peek() == '=' {
			l.tok = Token{LESS_EQ_THAN, "<=", l.line}
			l.advance()
			} else {
			l.tok = Token{LESS_THAN, "<", l.line}
		}
	case '>':
		if l.peek() == '=' {
			l.tok = Token{LESS_EQ_THAN, "<=", l.line}
			l.advance()
			} else {
			l.tok = Token{LESS_THAN, "<", l.line}
		}
	case ';':
		l.tok = Token{SEMICOLON, ";", l.line}
	case ',':
		l.tok = Token{COMMA, ",", l.line}
	case '{':
		l.tok = Token{LBRACE, "{", l.line}
	case '}':
		l.tok = Token{RBRACE, "}", l.line}
	case '(':
		l.tok = Token{LPAR, "(", l.line}
	case ')':
		l.tok = Token{RPAR, ")", l.line}
	case 0:
		l.tok = Token{EOF, "", l.line}
	default:
		if isLetter(l.ch){
			word:= l.getWord()
			if tokentype, ok := keywords[word]; ok{
				l.tok = Token{tokentype, word, l.line}
			} else {
				l.tok = Token{IDENTIFIER, word, l.line}
			}
		} else if isDigit(l.ch) {
			l.getnumber()
		} else {
			l.tok = Token{ILLEGAL, string(l.ch), l.line}
		}
	}

	l.advance()

	return l.tok
}
