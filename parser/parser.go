package parser

import (
    "github.com/gipx/token"
    "github.com/gipx/ast"
    "github.com/gipx/lexer"
)

type Parser struct {

  l *lexer.Lexer

  curToken token.Token
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

  p.curToken = p.peekToken
  p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
  program := &ast.Program{}

  program.Statements = []ast.Statement{}

  for p.curToken.Type != token.EOF {
    stmt := p.parseStatement()
    if stmt != nil {
      program.Statements = append(program.Statements, stmt)
    }
    p.NextToken()
  }

  return program
}

func (p *Parser) ParseStatement() ast.Statement {
  switch p.curToken.Type {
  case token.LET:
    return p.parseLetStatement()
  default:
    return nil
  }
}
