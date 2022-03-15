package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENTIFIER" // add, foobar, x, y, ...
	INT   = "INTEGER"   // 1343456

	// Operators
	ASSIGN   = "ASSIGN" // =
	PLUS     = "PLUS" // +
	MINUS    = "MINUS" // -
	BANG     = "BANG" // !
	ASTERISK = "ASTERISK" // *
	SLASH    = "SLASH" // /

	LT = "LT" // <
	GT = "GT" // >

	EQ     = "EQ" // ==
	NOT_EQ = "NOT_EQ" // !=

	// Delimiters
	COMMA     = "COMMA" // ,
	SEMICOLON = "SEMICOLON" // ;

	LPAREN = "LPAREN" // (
	RPAREN = "RPAREN" // )
	LBRACE = "LBRACE" // {
	RBRACE = "RBRACE" // }

	// String
	STRING 	 = "STRING"

	// Keywords
	PROGRAM  = "PROGRAM"
	FUNCTION = "FUNCTION"
	RENAME   = "RENAME"
	MOVE 	 = "MOVE"
	READ 	 = "READ"
	WEIGHT 	 = "WEIGHT"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"Program": 	PROGRAM,
	"function": FUNCTION,
	"rename": 	RENAME,
	"move": 	MOVE,
	"read": 	READ,
	"weight": 	WEIGHT,
	"let":    	LET,
	"true":   	TRUE,
	"false":  	FALSE,
	"if":     	IF,
	"else":   	ELSE,
	"return": 	RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
