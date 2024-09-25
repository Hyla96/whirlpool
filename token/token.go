package token

type TokenType string
type Token struct {
	Type    TokenType
	Literal string
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

var keywords = map[string]TokenType{
	"siphon":  SIPHON,
	"flicker": FLICKER,
	"buoy":    BUOY,
	"cyclone": CYCLONE,
	"output":  OUTPUT,
	"true":    TRUE,
	"false":   FALSE,
}

const (
	// Special
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Keywords
	SIPHON    TokenType = "SIPHON"
	FLICKER   TokenType = "FLICKER"
	BUOY      TokenType = "BUOY"
	CYCLONE   TokenType = "CYCLONE"
	TRUE      TokenType = "TRUE"
	FALSE     TokenType = "FALSE"
	OUTPUT    TokenType = "OUTPUT"
	SEMICOLON TokenType = ";"
	COLON     TokenType = ";"
	LP        TokenType = "("
	RP        TokenType = ")"
	ASSIGN    TokenType = "="

	// Binary Operators
	SUM      TokenType = "+"
	SUBTRACT TokenType = "-"
	ASTERISK TokenType = "*"
	SLASH    TokenType = "/"

	// Prefix Operators
	BANG TokenType = "!"

	// Comparison Operators
	GT        TokenType = ">"
	LT        TokenType = "<"
	EQUAL     TokenType = "=="
	NOT_EQUAL TokenType = "!="
	GTE       TokenType = ">="
	LTE       TokenType = "<="

	// Other operators
	FLOW TokenType = "->"
	// Identifiers
	IDENT TokenType = "IDENT"

	// Literals
	INT TokenType = "INT"
)
