package lexer

import (
	"github.com/pykmi/coconut/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	// skip whitespaces, tabs and newlines
	lexer.skipWhiteSpace()

	switch lexer.ch {
	case '+':
		tok = newToken(token.PLUS, lexer.ch)
	case '-':
		tok = newToken(token.MINUS, lexer.ch)
	case '/':
		tok = newToken(token.SLASH, lexer.ch)
	case '*':
		tok = newToken(token.ASTERISK, lexer.ch)
	case '<':
		tok = newToken(token.LT, lexer.ch)
	case '>':
		tok = newToken(token.GT, lexer.ch)
	case ';':
		tok = newToken(token.SEMICOLON, lexer.ch)
	case '(':
		tok = newToken(token.LPAREN, lexer.ch)
	case ')':
		tok = newToken(token.RPAREN, lexer.ch)
	case ',':
		tok = newToken(token.COMMA, lexer.ch)
	case '{':
		tok = newToken(token.LBRACE, lexer.ch)
	case '}':
		tok = newToken(token.RBRACE, lexer.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	case '=':
		if lexer.peekChar() == '=' {
			ch := lexer.ch
			lexer.readChar()
			tok = token.Token{
				Type:    token.EQ,
				Literal: string(ch) + string(lexer.ch),
			}
		} else {
			tok = newToken(token.ASSIGN, lexer.ch)
		}
	case '!':
		if lexer.peekChar() == '=' {
			ch := lexer.ch
			lexer.readChar()
			tok = token.Token{
				Type:    token.NOT_EQ,
				Literal: string(ch) + string(lexer.ch),
			}
		} else {
			tok = newToken(token.BANG, lexer.ch)
		}
	default:
		if isLetter(lexer.ch) {
			literal := lexer.readIdentifier()
			return token.Token{
				Literal: literal,
				Type:    token.LookupIdent(literal),
			}
		} else if isDigit(lexer.ch) {
			return token.Token{
				Type:    token.INT,
				Literal: lexer.readNumber(),
			}
		}
		tok = newToken(token.ILLEGAL, lexer.ch)
	}

	lexer.readChar()
	return tok
}

func (lexer *Lexer) peekChar() byte {
	if lexer.readPosition >= len(lexer.input) {
		return 0
	}

	return lexer.input[lexer.readPosition]
}

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.ch = 0
	} else {
		lexer.ch = lexer.input[lexer.readPosition]
	}

	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) readIdentifier() string {
	position := lexer.position

	for isLetter(lexer.ch) {
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) skipWhiteSpace() {
	for lexer.ch == ' ' || lexer.ch == '\t' || lexer.ch == '\n' || lexer.ch == '\r' {
		lexer.readChar()
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (lexer *Lexer) readNumber() string {
	position := lexer.position
	for isDigit(lexer.ch) {
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}
