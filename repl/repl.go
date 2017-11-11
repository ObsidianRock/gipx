package repl

import (
  "bufio"
  "fmt"
  "io"
  "github.com/gipx/lexer"
  "github.com/gipx/token"
)

const PROMPT = ">>"


func Start(in io.Reader, out io.Writer){

  scanner := bufio.NewScanner(in)

  for {
    fmt.Printf(PROMPT)
    scanned := scanner.Scan()

    if !scanned{
      return
    }

    line := scanner.Text()
    l := lexer.New(line)

    for tok := l.NextToken(); tok.Typer != token.EOF; tok = l.NextToken() {
        fmt.Printf("%+v\n", tok)
    }
  }
}