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
	DIRECT
	UNDIRECT
	EQUALS

	// PUNCTS
	SEMICOLON
	COMMA
	COLON

	MINUS

	// KEYWORD
	STRICT
	GRAPH
	DIGRAPH
	SUBGRAPH
)

var KEYWORDS = map[string]TokenKind{
	"strict":   STRICT,
	"graph":    GRAPH,
	"digraph":  DIGRAPH,
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

func (t *Token) GetToken() string {
	return t.text
}

func (t *Token) GetKindToText() string {
	switch t.kind {
	case ID:
		return "ID"
	case INT:
		return "INT"
	case FLOAT:
		return "FLOAT"
	case OPEN_BRACE:
		return "OPEN_BRACE"
	case CLOSE_BRACE:
		return "CLOSE_BRACE"
	case OPEN_BRACKET:
		return "OPEN_BRACKET"
	case CLOSE_BRACKET:
		return "CLOSE_BRACKET"
	case DIRECT:
		return "DIRECT"
	case UNDIRECT:
		return "UNDIRECT"
	case SEMICOLON:
		return "SEMICOLON"
	case COMMA:
		return "COMMA"
	case COLON:
		return "COLON"
	case MINUS:
		return "MINUS"
	case EQUALS:
		return "EQUALS"
	case STRICT:
	case GRAPH:
	case SUBGRAPH:
	case DIGRAPH:
		return "KEYWORD"
	}
	return "EOF"
}
