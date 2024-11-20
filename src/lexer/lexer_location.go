package lexer

type LexerLoc struct {
	line   int
	column int
}

func NewLexerLoc() *LexerLoc {
	return &LexerLoc{
		line:   1,
		column: 0,
	}
}
