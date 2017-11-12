package ast

import (
  "github.com/gipx/token"
  "testing"
)


func TestString (t *testing.T) {

  program := &Program{
    Statements: []Statement{
      &LetStatement{
        Token: token.Token{Type: token.LET, Literal: "let"},
        Name: &Identifier{
          Token: token.Token{Type: token.IDENT, Literal: "va1"},
          Value: "var1",
        },
        Value: &Identifier{
          Token: token.Token{Type: token.IDENT, Literal: "var2"},
          Value: "var1",
        },
      },
    },
  }

  if program.String() != "let var1 = var2;" {
    t.Errorf("program.String() wrong. got=%q", program.String())
  }


}
