package lexer

import "whirlpool/src/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}

	l.readChar()

	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '>':
		tok = newToken(token.GREATER_THAN, l.ch)
	case '<':
		tok = newToken(token.LESS_THAN, l.ch)
	case 0:
		tok = newToken(token.EOF, l.ch)
	}

	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func newToken(t token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    t,
		Literal: string(ch),
	}
}
