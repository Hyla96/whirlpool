package parser

import (
	"fmt"
	"whirlpool/src/ast"
	"whirlpool/src/lexer"
	"whirlpool/src/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  *token.Token
	nextToken *token.Token

	errors []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.readToken()
	p.readToken()

	return p
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

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.BUOY:
		return p.parseBuoyStatement()
	default:
		return nil
	}

}
func (p *Parser) parseBuoyStatement() *ast.BuoyStatement {
	stmt := &ast.BuoyStatement{
		Token: p.curToken,
	}

	if !p.expectNext(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{
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
