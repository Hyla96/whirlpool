package token

type TokenType string
type Token struct {
	Type    TokenType
	Literal string
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
