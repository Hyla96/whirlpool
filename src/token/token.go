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

	// Single Character Operators
	ASSIGN       TokenType = "="
	GREATER_THAN TokenType = ">"
	LESS_THAN    TokenType = "<"
	SUM          TokenType = "+"
	SUBTRACT     TokenType = "-"
	MULTIPLY     TokenType = "*"
	DIVIDE       TokenType = "/"
	NOT          TokenType = "!"
	// Double Character Operators
	FLOW                  TokenType = "->"
	EQUAL                 TokenType = "=="
	NOT_EQUAL             TokenType = "!="
	GREATER_THAN_OR_EQUAL TokenType = ">="
	LESS_THAN_OR_EQUAL    TokenType = "<="

	// Identifiers
	IDENT TokenType = "IDENT"

	// Literals
	INT TokenType = "INT"
)
