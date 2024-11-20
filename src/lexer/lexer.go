package lexer

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
