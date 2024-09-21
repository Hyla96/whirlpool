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

type Identifier struct {
	Token *token.Token
	Value string
}

func (bs *Identifier) expressionNode() {}
func (bs *Identifier) TokenLiteral() string {
	return bs.Token.Literal
}
