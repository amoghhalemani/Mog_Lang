package lexer

type Token struct {
	Type  string
	Value string
}

const (
	KEYWORD    = "KEYWORD"
	IDENTIFIER = "IDENTIFIER"
	NUMBER     = "NUMBER"
	EQUAL      = "EQUAL"
	LPAREN     = "LPAREN"
	RPAREN     = "RPAREN"
	EOF        = "EOF"
	ILLEGAL    = "ILLEGAL"
)

type Lexer struct {
	source string
	pos    int
}

func New(source string) *Lexer {
	return &Lexer{source: source, pos: 0}
}

func (l *Lexer) NextToken() Token {
	for l.pos < len(l.source) && (l.source[l.pos] == ' ' || l.source[l.pos] == '\n' || l.source[l.pos] == '\t' || l.source[l.pos] == '\r') {
		l.pos++
	}
	if l.pos >= len(l.source) {
		return Token{Type: EOF, Value: ""}
	}
	ch := l.source[l.pos]
	if ch == '=' {
		l.pos++
		return Token{Type: EQUAL, Value: "="}
	}
	if ch == '(' {
		l.pos++
		return Token{Type: LPAREN, Value: "("}
	}
	if ch == ')' {
		l.pos++
		return Token{Type: RPAREN, Value: ")"}
	}
	if ch >= '0' && ch <= '9' {
		start := l.pos
		for l.pos < len(l.source) && l.source[l.pos] >= '0' && l.source[l.pos] <= '9' {
			l.pos++
		}
		return Token{Type: NUMBER, Value: l.source[start:l.pos]}
	}
	if ch >= 'a' && ch <= 'z' {
		start := l.pos
		for l.pos < len(l.source) && l.source[l.pos] >= 'a' && l.source[l.pos] <= 'z' {
			l.pos++
		}
		word := l.source[start:l.pos]
		if word == "let" || word == "print" {
			return Token{Type: KEYWORD, Value: word}
		}
		return Token{Type: IDENTIFIER, Value: word}
	}
	l.pos++
	return Token{Type: ILLEGAL, Value: string(ch)}
}
