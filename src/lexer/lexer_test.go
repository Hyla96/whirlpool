package lexer

import (
	"testing"
	"whirlpool/src/token"
)

func TestNextToken(t *testing.T) {
	input := `siphon num > 10 -> pipe`

	tests := []struct {
		expetedType     token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
	}

	l := New(input)

	for index, value := range tests {
		tok := l.NextToken()

		if tok.Type != value.expetedType {
			t.Fatalf("Token type wrong. Received %q and expected %q", tok.Type, value.expetedType)
		}

		if tok.Literal != value.expectedLiteral {
			t.Fatalf("Token literal wrong. Received %q and expected %q", tok.Literal, value.expectedLiteral)
		}
	}
}
