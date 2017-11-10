package lexer


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
