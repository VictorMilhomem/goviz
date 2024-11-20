package lexer

import "regexp"

type TokenKind uint8

type Token struct {
	kind TokenKind
	text string
}

const (
	EOF TokenKind = iota
	ID
	INT
	FLOAT

	OPEN_BRACE
	CLOSE_BRACE
	OPEN_BRACKET
	CLOSE_BRACKET

	// PUNCTS
	SEMICOLON
	COMMA
	COLON

	MINUS

	// KEYWORD
	STRICT
	GRAPH
	SUBGRAPH
)

var KEYWORDS = map[string]TokenKind{
	"strict":   STRICT,
	"graph":    GRAPH,
	"subgraph": SUBGRAPH,
}

var PATTERNS = map[TokenKind]*regexp.Regexp{
	INT:   regexp.MustCompile(`\d+`),
	FLOAT: regexp.MustCompile(`\d+\.\d+`),
	ID:    regexp.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]*`),
}

func NewToken(kind TokenKind, text string) Token {
	return Token{
		kind: kind,
		text: text,
	}
}
