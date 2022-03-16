package lexer

type Tokentype string

type Token struct {
	Type Tokentype
	Value string
	Line int
}

const (
	EOF     = "EOF"

	// errors
	ILLEGAL = "ILLEGAL"
	ILL_FLOAT = "ILLEGAL_FLOAT"
	ILL_INTEGER = "ILLEGAL INTEGER"

	// Identifiers

	IDENTIFIER = "IDENTIFIER"
	INTEGER = "INTEGER"
	FLOAT = "FLOAT"
	STRING = "STRING"

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

	// Operators
	ASSIGN   = "=" // =
	PLUS     = "+" // +
	MINUS    = "-" // -
	BANG     = "!" // !
	ASTERISK = "*" // *
	SLASH    = "/" // /

	EQUAL = "EQUAL"
	NOT_EQUAL = "NOT_EQUAL"

	LESS_THAN = "<"
	GREATER_THAN = ">"
	LESS_EQ_THAN = "<="
	GREATER_EQ_THAN = ">="

	// Delimiters
	COMMA = ","
	SEMICOLON = ";"

	LPAR = "("
	RPAR = ")"
	LBRACE = "{"
	RBRACE = "}"
)

var keywords = map[string]Tokentype{
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

// func GetKeyword(id string) Tokentype {
// 	if tok, ok := keywords[id]; ok {
// 		return tok
// 	}
// 	return IDENTIFIER
// }
