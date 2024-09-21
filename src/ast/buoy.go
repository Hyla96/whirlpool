package ast

import "whirlpool/src/token"

type BuoyStatement struct {
	Token *token.Token
	Name  *Identifier
	Value Expression
}

func (bs *BuoyStatement) statementNode() {}
func (bs *BuoyStatement) TokenLiteral() string {
	return bs.Token.Literal
}
