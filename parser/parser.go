package parser

import (
    "github.com/gipx/token"
    "github.com/gipx/ast"
    "github.com/gipx/lexer"
)

type Parser struct {

  l *lexer.Lexer

  currToken token.Token
  peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {

  p := &Parser{l: l}

  // Read two tokens, so curToken and peekToken are both set
  p.NextToken()
  p.NextToken()

  return p
}

func (p *Parser) NextToken() {

  p.currToken = p.peekToken
  p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
  return nil 
}
