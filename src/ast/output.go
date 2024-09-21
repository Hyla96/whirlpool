package ast

import "whirlpool/src/token"

type OutputStatement struct {
	Token       *token.Token
	ReturnValue Expression
}

func (bs *OutputStatement) statementNode() {}
func (bs *OutputStatement) TokenLiteral() string {
	return bs.Token.Literal
}
