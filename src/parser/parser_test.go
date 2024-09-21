package parser

import (
	"testing"
	"whirlpool/src/ast"
	"whirlpool/src/lexer"
)

func TestBuoyStatement1(t *testing.T) {
	input := `
	buoy x = 5;
	buoy y = 10;
	buoy z = 1001271;
	`
	program := getProgram(input, t)
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"z"},
	}
	for i, tt := range tests {
		statement := program.Statements[i]
		if !testBuoyStatement(t, statement, tt.expectedIdentifier) {
			return
		}
	}
}

func getProgram(input string, t *testing.T) *ast.Program {
	lex := lexer.New(input)
	pars := New(lex)

	program := pars.ParseProgram()
	checkParserError(t, pars)
	if program == nil {
		t.Fatal("ParseProgram returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("ParseProgram returned %q statements instead of 3", len(program.Statements))
	}

	return program
}

func testBuoyStatement(t *testing.T, statement ast.Statement, expected string) bool {
	if statement.TokenLiteral() != "buoy" {
		t.Errorf("TokenLiteral is not buoy. got=%q", statement.TokenLiteral())
		return false
	}

	s, ok := statement.(*ast.BuoyStatement)

	if !ok {
		t.Errorf("Statement is not buoy type. got=%T", statement)
		return false
	}

	if s.Name.Value != expected {
		t.Errorf("Statement name is not %s. got=%s", expected, s.Name.Value)
		return false
	}

	if s.Name.TokenLiteral() != expected {
		t.Errorf("Statement Name's TokenLiteral() is not %s. got=%s", expected, s.Name.TokenLiteral())
		return false
	}

	return true
}

func checkParserError(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))

	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
