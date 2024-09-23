package ast

import (
	"bytes"
	"github.com/Hyla96/whirlpool/src/token"
)

type Node interface {
	TokenLiteral() string
	String() string
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
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type IdentifierExpression struct {
	Token *token.Token
	Value string
}

func (bs *IdentifierExpression) expressionNode() {}
func (bs *IdentifierExpression) TokenLiteral() string {
	return bs.Token.Literal
}
func (bs *IdentifierExpression) String() string {
	return bs.Value
}

type IntegerLiteralExpression struct {
	Token *token.Token
	Value int64
}

func (il *IntegerLiteralExpression) expressionNode() {}
func (il *IntegerLiteralExpression) TokenLiteral() string {
	return il.Token.Literal
}
func (il *IntegerLiteralExpression) String() string {
	return il.TokenLiteral()
}

type ExpressionStatement struct {
	Token      *token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
