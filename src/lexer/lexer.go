package lexer

import "fmt"

type LexerError struct {
	errorString string
}

func (l *LexerError) Error() string {
	return l.errorString
}

func NewLexerError(msg string) error {
	return &LexerError{
		msg,
	}
}

type Lexer struct {
	source   string
	Tokens   []Token
	location *LexerLoc
}

func NewLexer(source string) *Lexer {
	return &Lexer{
		source:   source,
		Tokens:   make([]Token, 0),
		location: NewLexerLoc(),
	}
}

func (l *Lexer) advance_n(n int) {
	l.location.column += n
}

func (l *Lexer) advance_line() {
	l.advance_n(1)
	l.location.line += 1
}

func (l *Lexer) push_token(token Token) {
	l.Tokens = append(l.Tokens, token)
}

func (l *Lexer) at_eof() bool {
	return l.location.column >= len(l.source)
}

func (l *Lexer) at() int {
	return l.location.column
}

func (l *Lexer) lookahead(n int) rune {
	return rune(l.source[l.location.column+n])
}

func (l *Lexer) remainder() string {
	return l.source[l.location.column:]
}

func (l *Lexer) lookup_keyword(id string) TokenKind {
	kind, ok := KEYWORDS[id]

	if !ok {
		return ID
	}

	return kind
}

func (l *Lexer) is_whitespace(char rune) bool {
	return char == ' ' || char == '\t'
}

func (l *Lexer) Tokenizer() error {
	for !l.at_eof() {
		current := rune(l.source[l.at()])
		if l.is_whitespace(current) {
			l.advance_n(1)
			continue
		}

		switch current {
		case '\n':
			l.advance_line()
			continue
		case '[':
			l.push_token(NewToken(OPEN_BRACKET, string(current)))
			l.advance_n(1)
			continue
		case ']':
			l.push_token(NewToken(CLOSE_BRACKET, string(current)))
			l.advance_n(1)
			continue
		case '{':
			l.push_token(NewToken(OPEN_BRACE, string(current)))
			l.advance_n(1)
			continue
		case '}':
			l.push_token(NewToken(CLOSE_BRACE, string(current)))
			l.advance_n(1)
			continue
		case ';':
			l.push_token(NewToken(SEMICOLON, string(current)))
			l.advance_n(1)
		case ',':
			l.push_token(NewToken(COMMA, string(current)))
			l.advance_n(1)
			continue
		case ':':
			l.push_token(NewToken(COLON, string(current)))
			l.advance_n(1)
			continue
		case '-':
			look := l.lookahead(1)
			if look == '>' {
				l.push_token(NewToken(ARROW, string(current)+string(look)))
				l.advance_n(2)

			} else {
				l.push_token(NewToken(MINUS, string(current)))
				l.advance_n(1)
			}

			continue
		default:
			matched := false
			for kind, pattern := range PATTERNS {
				loc := pattern.FindStringIndex(l.remainder())

				if loc != nil && loc[0] == 0 {
					matchedText := l.remainder()[loc[0]:loc[1]]
					if kind == ID {
						kind = l.lookup_keyword(matchedText)
					}
					l.push_token(NewToken(kind, matchedText))
					matched = true
					l.advance_n(len(matchedText))
					break
				}
			}

			if !matched {
				// make a better error handling here
				return NewLexerError(
					fmt.Sprintf("Unknown Token %s at\nLine: %d\nPosition: %d", string(current), l.location.line, l.location.column),
				)
			}
		}

	}
	l.push_token(NewToken(EOF, ""))
	return nil
}
