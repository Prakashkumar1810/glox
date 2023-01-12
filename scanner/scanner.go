package scanner

import (
	"glox/token"
	"glox/util"
)

type Scanner struct {
	source  string
	tokens  []token.Token
	line    uint
	start   uint
	current uint
}

func NewScanner(source string) Scanner {
	return Scanner{
		source:  source,
		tokens:  []token.Token{},
		line:    1,
		start:   0,
		current: 0,
	}
}

func (s *Scanner) ScanTokens() []token.Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, token.Token{
		TokenType: token.EOF,
		Lexeme:    "",
		Literal:   nil,
		Line:      s.line,
	})
	return s.tokens
}

func (s *Scanner) scanToken() {
	c := s.advance()

	switch c {
	case '(':
		s.addToken(token.LEFT_PAREN, nil)
	case ')':
		s.addToken(token.RIGHT_PAREN, nil)
	case '{':
		s.addToken(token.LEFT_BRACE, nil)
	case '}':
		s.addToken(token.RIGHT_BRACE, nil)
	case ',':
		s.addToken(token.COMMA, nil)
	case '.':
		s.addToken(token.DOT, nil)
	case '-':
		s.addToken(token.MINUS, nil)
	case '+':
		s.addToken(token.PLUS, nil)
	case ';':
		s.addToken(token.SEMICOLON, nil)
	case '*':
		s.addToken(token.STAR, nil)
    case '!':
        s.addToken(choice(s.match('='), token.BANG_EQUAL, token.BANG).(token.TokenType), nil)
    case '=':
        s.addToken(choice(s.match('='), token.EQUAL_EQUAL, token.EQUAL).(token.TokenType), nil)
    case '<':
        s.addToken(choice(s.match('='), token.LESS_EQUAL, token.LESS).(token.TokenType), nil)
    case '>':
        s.addToken(choice(s.match('='), token.GREATER_EQUAL, token.GREATER).(token.TokenType), nil)
	default:
		util.Error(uint(s.line), "Unexpected character.")
	}
}

func (s *Scanner) addToken(tokenType token.TokenType, literal interface{}) {
	text := s.source[s.start:s.current]

	s.tokens = append(s.tokens, token.Token{
		TokenType: tokenType,
		Lexeme:    text,
		Literal:   literal,
		Line:      s.line,
	})
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= uint(len(s.source))
}

func (s *Scanner) advance() rune {
	c := s.source[s.current]
	s.current++
	return rune(c)
}

func (s *Scanner) match(expected rune) bool {
	if s.isAtEnd() || (expected != rune(s.source[s.current])) {
		return false
	}

	s.current++
	return true
}

func choice(condition bool, v1 interface{}, v2 interface{}) interface{} {
    if condition {
        return v1
    }
    return v2
}
