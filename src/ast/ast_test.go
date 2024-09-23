package ast

import (
	"github.com/Hyla96/whirlpool/src/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&BuoyStatement{
				Token: &token.Token{
					Type:    token.BUOY,
					Literal: "buoy",
				},
				Name: &IdentifierExpression{
					Token: &token.Token{
						Type:    token.IDENT,
						Literal: "something",
					},
					Value: "something",
				},
				Value: &IdentifierExpression{
					Token: &token.Token{
						Type:    token.INT,
						Literal: "10",
					},
					Value: "10",
				},
			},
		},
	}

	if program.String() != "buoy something = 10;" {
		t.Errorf("program.String() is wrong, got=%q", program.String())
	}
}
