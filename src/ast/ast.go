package ast

import "whirlpool/src/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

type BuoyStatement struct {
	Token *token.Token
	Name  *Identifier
	Value Expression
}

func (bs *BuoyStatement) statementNode() {}
func (bs *BuoyStatement) TokenLiteral() string {
	return bs.Token.Literal
}

type Identifier struct {
	Token *token.Token
	Value string
}

func (bs *Identifier) expressionNode() {}
func (bs *Identifier) TokenLiteral() string {
	return bs.Token.Literal
}
