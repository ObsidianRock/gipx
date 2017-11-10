package lexer


import (
    "github.com/gipx/token"
)

type Lexer struct {

  input        string
  position     int     // current Positon in input
  readPosition int    // current reading position
  ch           byte  // current char under examination
}


func New(input string) *Lexer {

  lx := &Lexer{input: input}
  lx.readChar()
  return lx
}

func (l *Lexer) readChar() {

  if l.readPosition >= len(l.input){
    l.ch = 0
  } else {
    l.ch = l.input[l.readPosition]
  }

  l.position = l.readPosition
  l.readPosition++
}

func (l *Lexer) NextToken() token.Token {

  var tok token.Token

  switch l.ch {
  case '=':
    tok = newToken(token.ASSIGN, l.ch)
  case ';':
    tok = newToken(token.SEMICOLON, l.ch)
  case '(':
    tok = newToken(token.LPARAN, l.ch)
  case ')':
    tok = newToken(token.RPARAN, l.ch)
  case '{':
    tok = newToken(token.LBRACE, l.ch)
  case '}':
    tok = newToken(token.RBRACE, l.ch)
  case ',':
    tok = newToken(token.COMMA, l.ch)
  case '+':
    tok = newToken(token.PLUS, l.ch)
  case '0':
    tok.Literal = ""
    tok.Type = token.EOF
  default:
    if isLetter(l.ch){
      tok.Literal = l.readIdentifier()
      return tok
    } else{
      tok = newToken(token.ILLEGAL, l.ch)
    }

  }

  l.readChar()
  return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {

  return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
  position = l.position

  for isLetter(l.ch) {
    l.readChar()
  }
  return l.input[position: l.position]
}

func isLetter(ch byte) bool {

  return 'a' <= ch && <= 'z' || 'A' <= ch && <= 'Z' || ch == '_'
}
