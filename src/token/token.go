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
}

const (
	// Special
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Keywords
	SIPHON  TokenType = "SIPHON"
	FLICKER TokenType = "FLICKER"
	CYCLONE TokenType = "CYCLONE"

	// Operators
	ASSIGN        TokenType = "="
	GREATER_THAN  TokenType = ">"
	LESS_THAN     TokenType = "<"
	FLOW_OPERATOR TokenType = "->"

	// Identifiers
	IDENT TokenType = "IDENT"

	// Literals
	INT TokenType = "INT"
)
