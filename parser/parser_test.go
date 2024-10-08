package parser

import (
	"fmt"
	"github.com/Hyla96/whirlpool/ast"
	"github.com/Hyla96/whirlpool/lexer"
	"testing"
)

func TestOutputStatement1(t *testing.T) {
	input := `
	output 5;
	output 10;
	output z;
	`
	program := getProgram(input, t)
	if len(program.Statements) != 3 {
		t.Fatalf("ParseProgram returned %q statements instead of 3", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		outputStmt, ok := stmt.(*ast.OutputStatement)

		if !ok {
			t.Errorf("Statement is not output type. got=%T", stmt)
		}

		if outputStmt.TokenLiteral() != "output" {
			t.Errorf("literal not output, got %q", stmt.TokenLiteral())
		}
	}
}

func TestBooleanExpression(t *testing.T) {
	input := `
	true;
	false;
	`
	program := getProgram(input, t)

	if len(program.Statements) != 2 {
		t.Fatalf("ParseProgram returned %d statements instead of 2", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		boolean, ok := stmt.(*ast.ExpressionStatement)

		if !ok {
			t.Errorf("Statement is not ExpressionStatement")
		}

		_, ok = boolean.Expression.(*ast.Boolean)

		if !ok {
			t.Errorf("Expression is not boolean")
		}
	}
}
func TestIntegerLiteralExpression(t *testing.T) {
	input := `
	5;
	`
	program := getProgram(input, t)

	if len(program.Statements) != 1 {
		t.Fatalf("ParseProgram returned %d statements instead of 1", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Error("Statement is not ExpressionStatement")
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)

	if !ok {
		t.Error("Expression is not int literal")
	}

	if literal.Value != 5 {
		t.Errorf("Literal is not 5, got=%q", literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Errorf("TokenLiteral is not 5, got=%q", literal.TokenLiteral())
	}
}
func TestIdentifierExpression(t *testing.T) {
	input := `
	num;
	`
	program := getProgram(input, t)

	if len(program.Statements) != 1 {
		t.Fatalf("ParseProgram returned %d statements instead of 1", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Error("Statement is not ExpressionStatement")
	}

	iden, ok := stmt.Expression.(*ast.IdentifierExpression)

	if !ok {
		t.Error("Expression is not identifier")
	}

	if iden.Value != "num" {
		t.Errorf("IdentifierExpression is not num, got=%q", iden.Value)
	}
}

func TestPrefixExpressions(t *testing.T) {
	prefixTest := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}

	for _, tt := range prefixTest {
		program := getProgram(tt.input, t)

		if len(program.Statements) != 1 {
			t.Fatalf("Statements length not 1, got %d", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("Expected to get ExpressionStatement, but got %T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("Expected to get PrefixExpression, but got %T", stmt)
		}

		if exp.Operator != tt.operator {
			t.Fatalf("Expected to get %q, but got %q", tt.operator, exp.Operator)
		}

		if !testLiteralExpression(t, exp.Right, tt.integerValue) {
			return
		}
	}
}

func TestInfixExpressions(t *testing.T) {

	infixTest := []struct {
		input    string
		left     interface{}
		operator string
		right    interface{}
	}{
		{"5 + 5", 5, "+", 5},
		{"5 - 5", 5, "-", 5},
		{"5 * 5", 5, "*", 5},
		{"5 / 5", 5, "/", 5},
		{"5 > 5", 5, ">", 5},
		{"5 < 5", 5, "<", 5},
		{"5 >= 5", 5, ">=", 5},
		{"5 <= 5", 5, "<=", 5},
		{"5 == 5", 5, "==", 5},
		{"5 != 5", 5, "!=", 5},
		{"true == true", true, "==", true},
		{"true!=false", true, "!=", false},
		{"false != true", false, "!=", true},
		{"false == false", false, "==", false},
	}

	for _, tt := range infixTest {
		program := getProgram(tt.input, t)

		if len(program.Statements) != 1 {
			t.Fatalf("Statements length not 1, got %d", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("Expected to get ExpressionStatement, but got %T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.InfixExpression)

		if !ok {
			t.Fatalf("Expected to get InfixExpression, but got %T", stmt)
		}

		if !testLiteralExpression(t, exp.Left, tt.left) {

		}
		if exp.Operator != tt.operator {
			t.Fatalf("Expected to get %q, but got %q", tt.operator, exp.Operator)
		}

		if !testLiteralExpression(t, exp.Right, tt.right) {
			return
		}
	}
}

func TestBuoyStatement1(t *testing.T) {
	input := `
	buoy x = 5;
	buoy y = 10;
	buoy z = 1001271;
	`
	program := getProgram(input, t)

	if len(program.Statements) != 3 {
		t.Fatalf("ParseProgram returned %q statements instead of 3", len(program.Statements))
	}
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

func TestPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"5 + 5", "(5 + 5)"},
		{"-a * b", "((-a) * b)"},
		{"!-a", "(!(-a))"},
		{"a + b + c", "((a + b) + c)"},
		{"a + b * c", "(a + (b * c))"},
		{"(a + b) * c", "((a + b) * c)"},
		{"a + b / c", "(a + (b / c))"},
		{"3 + 4; -5 * 5", "(3 + 4)((-5) * 5)"},
		{"3 + 4 == -5 * 5", "((3 + 4) == ((-5) * 5))"},
		{"false", "false"},
		{"true", "true"},
		{"3 > 5 == false", "((3 > 5) == false)"},
	}

	for _, tt := range tests {
		program := getProgram(tt.input, t)
		actual := program.String()
		if actual != tt.expected {
			t.Errorf("expected %q, got %q", tt.expected, actual)
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

	return program
}

func testBuoyStatement(t *testing.T, stmt ast.Statement, expected string) bool {
	if stmt.TokenLiteral() != "buoy" {
		t.Errorf("TokenLiteral is not buoy. got=%q", stmt.TokenLiteral())
		return false
	}

	s, ok := stmt.(*ast.BuoyStatement)

	if !ok {
		t.Errorf("Statement is not buoy type. got=%T", stmt)
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

func testLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, exp, int64(v))
	case int64:
		return testIntegerLiteral(t, exp, v)
	case string:
		return testIdentifier(t, exp, v)
	case bool:
		return testBoolean(t, exp, v)
	}
	t.Errorf("type of exp not handled. got=%T", exp)
	return false
}
func testIdentifier(t *testing.T, il ast.Expression, value string) bool {
	v, ok := il.(*ast.IdentifierExpression)
	if !ok {
		t.Errorf("expected Identifier Expression, but got %T", v)
		return false
	}

	if v.Value != value {
		t.Errorf("expected Value %s, got %s", value, v.Value)
		return false
	}

	if v.TokenLiteral() != fmt.Sprintf("%s", value) {
		t.Errorf("expected TokenLiteral to return %s, got %s", value, v.TokenLiteral())
		return false
	}

	return true
}
func testBoolean(t *testing.T, il ast.Expression, value bool) bool {
	v, ok := il.(*ast.Boolean)
	if !ok {
		t.Errorf("expected Boolean, but got %T", v)
		return false
	}

	if v.Value != value {
		t.Errorf("expected Value %t, got %t", value, v.Value)
		return false
	}

	if v.TokenLiteral() != fmt.Sprintf("%t", value) {
		t.Errorf("expected TokenLiteral to return %t, got %s", value, v.TokenLiteral())
		return false
	}

	return true
}
func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("expected IntegerLiteral, but got %T", integ)
		return false
	}

	if integ.Value != value {
		t.Errorf("expected Value %d, got %d", value, integ.Value)
		return false
	}

	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("expected TokenLiteral to return %d, got %s", value, integ.TokenLiteral())
		return false
	}

	return true
}
