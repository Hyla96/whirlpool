package ast

import (
	"bytes"
	"whirlpool/src/token"
)

type BuoyStatement struct {
	Token *token.Token
	Name  *Identifier
	Value Expression
}

func (bs *BuoyStatement) statementNode() {}
func (bs *BuoyStatement) TokenLiteral() string {
	return bs.Token.Literal
}

func (bs *BuoyStatement) String() string {
	var out bytes.Buffer

	out.WriteString(bs.TokenLiteral() + " ")
	out.WriteString(bs.Name.String())
	out.WriteString(" = ")

	if bs.Value != nil {
		out.WriteString(bs.Value.String())
	}

	out.WriteString(";")

	return out.String()
}
