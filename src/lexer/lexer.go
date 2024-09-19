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
	case ' ':
		l.readChar()
		return l.NextToken()
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '>':
		tok = newToken(token.GREATER_THAN, l.ch)
	case '<':
		tok = newToken(token.LESS_THAN, l.ch)
	case 0:
		tok = newToken(token.EOF, l.ch)
	default:
		if isLetter(l.ch) {
			tok = l.getIdentifier()
		} else if isNumber(l.ch) {
			tok = l.getNumber()
		} else if isFlowOperator(l.input, l.position) {
			tok = l.getFlowOperator()
		} else {
			tok = token.Token{
				Type:    token.ILLEGAL,
				Literal: "Token not valid",
			}
		}
	}

	l.readChar()
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

func (l *Lexer) getIdentifier() token.Token {
	literal := ""
	var tok token.Token

	for isLetter(l.ch) || isNumber(l.ch) {
		literal += string(l.ch)
		l.readChar()
	}

	tok.Literal = literal

	switch literal {
	case "siphon":
		tok.Type = token.SIPHON
	case "flicker":
		tok.Type = token.FLICKER
	case "cyclone":
		tok.Type = token.CYCLONE
	default:
		tok.Type = token.IDENT
	}

	return tok
}
func (l *Lexer) getFlowOperator() token.Token {
	// Skip one
	l.readChar()
	return token.Token{
		Type:    token.FLOW_OPERATOR,
		Literal: "->",
	}
}
func (l *Lexer) getNumber() token.Token {
	literal := ""

	for isNumber(l.ch) {
		literal += string(l.ch)
		l.readChar()
	}

	return token.Token{
		Type:    token.INT,
		Literal: literal,
	}
}
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func isFlowOperator(input string, position int) bool {
	literal := string(input[position])
	literal += string(input[position+1])
	return literal == "->"
}
func isNumber(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
func newToken(t token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    t,
		Literal: string(ch),
	}
}
