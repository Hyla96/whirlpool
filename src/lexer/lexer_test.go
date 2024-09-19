package lexer

import (
	"testing"
	"whirlpool/src/token"
)

func TestNextToken1(t *testing.T) {
	input := `=><`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.GREATER_THAN, ">"},
		{token.LESS_THAN, "<"},
	}

	testInput(t, input, tests)
}

func TestNextToken2(t *testing.T) {
	input := `siphon num > 10 -> pipe`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.SIPHON, "siphon"},
		{token.IDENT, "num"},
		{token.GREATER_THAN, ">"},
		{token.INT, "10"},
		{token.FLOW_OPERATOR, "->"},
		{token.IDENT, "pipe"},
	}

	testInput(t, input, tests)
}

func TestNextToken3(t *testing.T) {
	input := `siphon num>10->pipe`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.SIPHON, "siphon"},
		{token.IDENT, "num"},
		{token.GREATER_THAN, ">"},
		{token.INT, "10"},
		{token.FLOW_OPERATOR, "->"},
		{token.IDENT, "pipe"},
	}

	testInput(t, input, tests)
}
func TestNextToken4(t *testing.T) {
	input := `siphon num2>10->pipe`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.SIPHON, "siphon"},
		{token.IDENT, "num2"},
		{token.GREATER_THAN, ">"},
		{token.INT, "10"},
		{token.FLOW_OPERATOR, "->"},
		{token.IDENT, "pipe"},
	}

	testInput(t, input, tests)
}

func testInput(t *testing.T, input string, tests []struct {
	expectedType    token.TokenType
	expectedLiteral string
}) {
	l := New(input)

	for _, value := range tests {
		tok := l.NextToken()

		if tok.Type != value.expectedType {
			t.Fatalf("Token type wrong. Received %q and expected %q - %q", tok.Type, value.expectedType, tok.Literal)
		}

		if tok.Literal != value.expectedLiteral {
			t.Fatalf("Token literal wrong. Received %q and expected %q - %q", tok.Literal, value.expectedLiteral, tok.Type)
		}

		t.Logf("Successfully read %q as %q", tok.Literal, tok.Type)
	}
}
