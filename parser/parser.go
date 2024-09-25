package parser

import (
	"fmt"
	"github.com/Hyla96/whirlpool/ast"
	"github.com/Hyla96/whirlpool/lexer"
	"github.com/Hyla96/whirlpool/token"
	"strconv"
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(expression ast.Expression) ast.Expression
)

type precedence int

const (
	_ precedence = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

var precedences = map[token.TokenType]precedence{
	token.EQUAL:     EQUALS,
	token.NOT_EQUAL: EQUALS,
	token.LT:        LESSGREATER,
	token.LTE:       LESSGREATER,
	token.GT:        LESSGREATER,
	token.GTE:       LESSGREATER,
	token.SUM:       SUM,
	token.SUBTRACT:  SUM,
	token.ASTERISK:  PRODUCT,
	token.SLASH:     PRODUCT,
}

type Parser struct {
	errors []string

	l         *lexer.Lexer
	curToken  *token.Token
	nextToken *token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:              l,
		errors:         []string{},
		prefixParseFns: make(map[token.TokenType]prefixParseFn),
		infixParseFns:  make(map[token.TokenType]infixParseFn),
	}

	p.registerPrefixParsingFn(token.IDENT, p.parseIdentifierExpression)

	p.registerPrefixParsingFn(token.INT, p.parseIntegerLiteralExpression)

	p.registerPrefixParsingFn(token.TRUE, p.parseBooleanExpression)
	p.registerPrefixParsingFn(token.FALSE, p.parseBooleanExpression)

	p.registerPrefixParsingFn(token.BANG, p.parsePrefixExpression)
	p.registerPrefixParsingFn(token.SUBTRACT, p.parsePrefixExpression)

	p.registerPrefixParsingFn(token.LP, p.parseGroupedExpression)

	p.registerInfixParsingFn(token.SUM, p.parseInfixExpression)
	p.registerInfixParsingFn(token.SUBTRACT, p.parseInfixExpression)
	p.registerInfixParsingFn(token.ASTERISK, p.parseInfixExpression)
	p.registerInfixParsingFn(token.SLASH, p.parseInfixExpression)
	p.registerInfixParsingFn(token.GT, p.parseInfixExpression)
	p.registerInfixParsingFn(token.GTE, p.parseInfixExpression)
	p.registerInfixParsingFn(token.LT, p.parseInfixExpression)
	p.registerInfixParsingFn(token.LTE, p.parseInfixExpression)
	p.registerInfixParsingFn(token.EQUAL, p.parseInfixExpression)
	p.registerInfixParsingFn(token.NOT_EQUAL, p.parseInfixExpression)

	p.readToken()
	p.readToken()

	return p
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	exp := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.readToken()
	right := p.parseExpression(PREFIX)
	exp.Right = right
	return exp
}

func (p *Parser) parseInfixExpression(expression ast.Expression) ast.Expression {
	exp := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     expression,
	}

	prec := p.currentPrecedence()
	p.readToken()
	exp.Right = p.parseExpression(prec)

	return exp
}

func (p *Parser) parseIdentifierExpression() ast.Expression {
	return &ast.IdentifierExpression{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.readToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectNext(token.RP) {
		return nil
	}

	return exp
}

func (p *Parser) parseBooleanExpression() ast.Expression {
	return &ast.Boolean{
		Token: p.curToken,
		Value: p.curToken.Type == token.TRUE,
	}
}
func (p *Parser) parseIntegerLiteralExpression() ast.Expression {
	literal := &ast.IntegerLiteral{
		Token: p.curToken,
	}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	literal.Value = value
	return literal
}

func (p *Parser) readToken() {
	p.curToken = p.nextToken
	p.nextToken = p.l.NextToken()
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) nextTokenError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.nextToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.readToken()
	}

	return program
}

func (p *Parser) noPrefixParseFnError(t *token.Token) {
	msg := fmt.Sprintf("no prefix parse function for %s found `%s`", t.Type, t.Literal)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.BUOY:
		return p.parseBuoyStatement()
	case token.OUTPUT:
		return p.parseOutputStatement()
	default:
		return p.parseExpressionStatement()
	}

}
func (p *Parser) parseOutputStatement() *ast.OutputStatement {
	stmt := &ast.OutputStatement{
		Token: p.curToken,
	}

	for !p.currentTokenIs(token.SEMICOLON) {
		p.readToken()
	}

	return stmt
}
func (p *Parser) parseBuoyStatement() *ast.BuoyStatement {
	stmt := &ast.BuoyStatement{
		Token: p.curToken,
	}

	if !p.expectNext(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.IdentifierExpression{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	if !p.expectNext(token.ASSIGN) {
		return nil
	}

	//TODO: here evaluate expression

	for !p.expectNext(token.SEMICOLON) {
		p.readToken()
	}

	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{
		Token: p.curToken,
	}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.nextTokenIs(token.SEMICOLON) { // this makes semicolon optional
		p.readToken()
	}

	return stmt
}
func (p *Parser) parseExpression(t precedence) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]

	if prefix == nil {
		p.noPrefixParseFnError(p.curToken)
		return nil
	}

	leftExp := prefix()

	for !p.nextTokenIs(token.SEMICOLON) && t < p.nextPrecedence() {
		infix := p.infixParseFns[p.nextToken.Type]
		if infix == nil {
			return leftExp
		}

		p.readToken()

		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) nextPrecedence() precedence {
	if pre, ok := precedences[p.nextToken.Type]; ok {
		return pre
	}
	return LOWEST
}

func (p *Parser) currentPrecedence() precedence {
	if pre, ok := precedences[p.curToken.Type]; ok {
		return pre
	}
	return LOWEST
}

func (p *Parser) currentTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) nextTokenIs(t token.TokenType) bool {
	return p.nextToken.Type == t
}
func (p *Parser) expectNext(t token.TokenType) bool {

	if p.nextTokenIs(t) {
		p.readToken()
		return true
	}

	p.nextTokenError(t)
	return false
}

func (p *Parser) registerPrefixParsingFn(tt token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tt] = fn
}

func (p *Parser) registerInfixParsingFn(tt token.TokenType, fn infixParseFn) {
	p.infixParseFns[tt] = fn
}
