package ast

import (
	"bytes"
	"whirlpool/src/token"
)

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
