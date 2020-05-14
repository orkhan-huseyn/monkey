package token

// Type denotes type of token
type Type string

// Token denotes a token object
type Token struct {
	Type    Type
	Literal string
}

const (
	// ILLEGAL is illegal identifier
	ILLEGAL = "ILLEGAL"
	// EOF denotes end of file
	EOF = "EOF"

	// IDENT represents an identifier
	IDENT = "IDENT"
	// INT represents an integer
	INT = "INT"

	// ASSIGN represents assignment operator
	ASSIGN = "="
	// PLUS represents addition operator
	PLUS = "+"
	// MINUS represents subtraction operator
	MINUS = "-"
	// BANG represents negation operator
	BANG = "!"
	// ASTERISK represents multiplication operator
	ASTERISK = "*"
	// SLASH represents division operator
	SLASH = "/"

	// LT - less than operator
	LT = "<"
	// GT represents greater than operator
	GT = ">"
	// EQ represents equals operators
	EQ = "=="
	// NOTEQ represents not equals operator
	NOTEQ = "!="

	// COMMA represents comma :)
	COMMA = ","
	// SEMICOLON represents, well, semicolon
	SEMICOLON = ";"

	// LPAREN represents left parenthesis
	LPAREN = "("
	// RPAREN represents right parenthesis
	RPAREN = ")"
	// LBRACE represents left curly brace
	LBRACE = "{"
	// RBRACE represents right curly brace
	RBRACE = "}"

	// FUNCTION represents a fn keyword
	FUNCTION = "FUNCTION"
	// LET denotes let keyword
	LET = "LET"
	// TRUE represents true keyrord
	TRUE = "TRUE"
	// FALSE represents false keyword
	FALSE = "FALSE"
	// IF represents if keyword
	IF = "IF"
	// ELSE represents else keyword
	ELSE = "ELSE"
	// RETURN represents return keyword
	RETURN = "RETURN"
)

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent finds correct token
// for program keyword
func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
