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
	"cyclone": CYCLONE,
	"true":    TRUE,
	"false":   FALSE,
}

const (
	// Special
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Keywords
	SIPHON  TokenType = "SIPHON"
	FLICKER TokenType = "FLICKER"
	CYCLONE TokenType = "CYCLONE"
	TRUE    TokenType = "TRUE"
	FALSE   TokenType = "FALSE"

	// Single Character Operators
	ASSIGN    TokenType = "="
	GT        TokenType = ">"
	LT        TokenType = "<"
	SUM       TokenType = "+"
	SUBTRACT  TokenType = "-"
	ASTERISK  TokenType = "*"
	SLASH     TokenType = "/"
	BANG      TokenType = "!"
	SEMICOLON TokenType = ";"
	COLON     TokenType = ";"

	// Double Character Operators
	FLOW      TokenType = "->"
	EQUAL     TokenType = "=="
	NOT_EQUAL TokenType = "!="
	GTE       TokenType = ">="
	LTE       TokenType = "<="

	// Identifiers
	IDENT TokenType = "IDENT"

	// Literals
	INT TokenType = "INT"
)
