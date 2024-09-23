package ast

import "whirlpool/src/token"

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

type Identifier struct {
	Token *token.Token
	Value string
}

type IntegerLiteral struct {
	Token *token.Token
	Value int64
}

func (bs *Identifier) expressionNode() {}
func (bs *Identifier) TokenLiteral() string {
	return bs.Token.Literal
}
func (bs *Identifier) String() string {
	return bs.Value
}
func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}
func (il *IntegerLiteral) String() string {
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
