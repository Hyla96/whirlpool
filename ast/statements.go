package ast

import (
	"bytes"
	"github.com/Hyla96/whirlpool/token"
)

type BuoyStatement struct {
	Token *token.Token
	Name  *IdentifierExpression
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

type OutputStatement struct {
	Token       *token.Token
	ReturnValue Expression
}

func (os *OutputStatement) statementNode() {}
func (os *OutputStatement) TokenLiteral() string {
	return os.Token.Literal
}

func (os *OutputStatement) String() string {
	var out bytes.Buffer

	out.WriteString(os.TokenLiteral() + " ")

	if os.ReturnValue != nil {
		out.WriteString(os.ReturnValue.String())
	}

	out.WriteString(";")
	return out.String()
}
