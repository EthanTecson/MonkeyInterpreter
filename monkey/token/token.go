type TokenType string

type Token struct {
	Type TokenType
	Literal String
}

const (
	ILLEGAL = "ILLEGAL"	
	EOF = "EOF"

	// Identifiers
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT = "INT" // 67676767

	// Operators
	ASSIGN = "="
	PLUS = "+"

	// Delimiters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)