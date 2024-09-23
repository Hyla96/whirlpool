package ast

import (
	"bytes"
	"github.com/Hyla96/whirlpool/token"
)

type InfixExpression struct {
	Token    *token.Token
	Operator string
	Left     Expression
	Right    Expression
}

func (pe *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Left.String())
	out.WriteString(" ")
	out.WriteString(pe.Operator)
	out.WriteString(" ")
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}
func (pe *InfixExpression) expressionNode() {}
func (pe *InfixExpression) TokenLiteral() string {
	return pe.Token.Literal
}
