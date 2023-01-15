package scanner

import (
	"glox/token"
	"glox/util"
	"strconv"
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
	case '/':
		if s.match('/') {
			for !s.isAtEnd() && s.peek() != '\n' {
				s.advance()
			}
		} else {
			s.addToken(token.SLASH, nil)
		}
	case ' ':
	case '\r':
	case '\t':
	case '\n':
		s.line++
	case '"':
		s.readString()
	default:
		if isDigit(c) {
			s.readNumber()
		} else if isAlpha(c) {
            s.readIdentifier()
		} else {
            util.Error(uint(s.line), "Unexpected character: ["+string(c)+"]")
		}
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

func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return 0
	}

	return rune(s.source[s.current])
}

func (s *Scanner) peekNext() rune {
	if (s.current + 1) >= uint(len(s.source)) {
		return 0
	}

	return rune(s.source[s.current+1])
}

func (s *Scanner) readString() {
	for !s.isAtEnd() && s.peek() != '"' {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		util.Error(s.line, "Unterminated string.")
		return
	}

	// The closing ".
	s.advance()

	// Trim the surrounding quotes
	value := s.source[s.start+1 : s.current-1]
	s.addToken(token.STRING, value)
}

func (s *Scanner) readNumber() {
	for isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && isDigit(s.peekNext()) {
		// Consume '.'
		s.advance()

		for isDigit(s.peek()) {
			s.advance()
		}
	}

	value, err := strconv.ParseFloat(s.source[s.start:s.current], 64)
	if err != nil {
		util.Error(s.line, "Error while parsing Number.")
		return
	}

	s.addToken(token.NUMBER, value)
}

func (s *Scanner) readIdentifier() {
    for isAlphaNumeric(s.peek()) {
        s.advance()
    }

    text := s.source[s.start: s.current]
    if val, ok := keywords[text]; ok {
        s.addToken(val, text)
    } else {
        s.addToken(token.IDENTIFIER, text)
    }
}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') || 
            (c >= 'A' && c <= 'Z') || c == '_'
}

func isAlphaNumeric(c rune) bool {
    return isAlpha(c) || isDigit(c)
}

func choice(condition bool, v1 interface{}, v2 interface{}) interface{} {
	if condition {
		return v1
	}
	return v2
}

var keywords map[string]token.TokenType = map[string]token.TokenType {
    "and": token.AND,
    "class": token.CLASS,
    "else": token.ELSE,
    "false": token.FALSE,
    "for": token.FOR,
    "fun": token.FUN,
    "if": token.IF,
    "nil": token.NIL,
    "or": token.OR,
    "print": token.PRINT,
    "return": token.RETURN,
    "super": token.SUPER,
    "this": token.THIS,
    "true": token.TRUE,
    "where": token.WHILE,
    "while": token.WHILE,
}

