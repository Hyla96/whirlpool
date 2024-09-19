package lexer

import (
	"whirlpool/src/token"
)

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

func (l *Lexer) NextToken() *token.Token {
	l.consumeWhitespace()

	var tok *token.Token

	tok = l.getDoubleCharacter()
	if tok != nil {
		return tok
	}

	tok = l.getSingleCharacter()
	if tok != nil {
		return tok
	}

	if isLetter(l.ch) {
		return l.getIdentifier()
	}

	if isNumber(l.ch) {
		return l.getNumber()
	}

	return &token.Token{
		Type:    token.ILLEGAL,
		Literal: "Token not valid",
	}
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

func (l *Lexer) getSingleCharacter() *token.Token {
	var tok *token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '+':
		tok = newToken(token.SUM, l.ch)
	case '-':
		tok = newToken(token.SUBTRACT, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '!':
		tok = newToken(token.BANG, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	case 0:
		tok = newToken(token.EOF, l.ch)
	}

	if tok != nil {
		l.readChar()
	}

	return tok
}
func (l *Lexer) getDoubleCharacter() *token.Token {
	//TODO: It should be optimized by including this check in the singleCharacter
	// First you check input[x] and if it matches with doubleCharacter[0] you also check x+1
	// This is easier to read tho' :)
	if len(l.input)-l.position < 2 {
		return nil
	}

	var tok *token.Token

	literal := string(l.input[l.position])
	literal += string(l.input[l.position+1])

	switch literal {
	case "==":
		tok = &token.Token{Type: token.EQUAL, Literal: literal}
	case "!=":
		tok = &token.Token{Type: token.NOT_EQUAL, Literal: literal}
	case ">=":
		tok = &token.Token{Type: token.GTE, Literal: literal}
	case "<=":
		tok = &token.Token{Type: token.LTE, Literal: literal}
	case "->":
		tok = &token.Token{Type: token.FLOW, Literal: literal}
	}

	if tok != nil {
		l.readChar()
		l.readChar()
	}

	return tok
}

func (l *Lexer) getIdentifier() *token.Token {
	init := l.position

	for l.position < len(l.input) && (isLetter(l.ch) || isNumber(l.ch)) {
		l.readChar()
	}

	return &token.Token{
		Type:    token.LookupIdent(l.input[init:l.position]),
		Literal: l.input[init:l.position],
	}
}
func (l *Lexer) getNumber() *token.Token {
	init := l.position

	for l.position < len(l.input) && isNumber(l.ch) {
		l.readChar()
	}

	return &token.Token{
		Type:    token.INT,
		Literal: l.input[init:l.position],
	}
}

func (l *Lexer) consumeWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func isNumber(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
func newToken(t token.TokenType, ch byte) *token.Token {
	return &token.Token{
		Type:    t,
		Literal: string(ch),
	}
}
