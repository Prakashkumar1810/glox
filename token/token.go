package token

import "fmt"

type Token struct {
    TokenType TokenType
    Lexeme string
    Literal interface{}
    Line uint
}

func (t Token) String() string {
    return string(t.TokenType) + " " + t.Lexeme + " " + fmt.Sprint(t.Literal)
}
