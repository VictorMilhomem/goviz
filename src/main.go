package main

import (
	"fmt"
	"log"

	"github.com/VictorMilhomem/goviz/src/lexer"
)

func main() {
	source := "digraph {\n A -> B  }"
	lex := lexer.NewLexer(source)
	err := lex.Tokenizer()
	if err != nil {
		log.Fatal(err)
	}

	for i := range lex.Tokens {
		fmt.Printf("(Kind: %s, Text: %s)\n", lex.Tokens[i].GetKindToText(), lex.Tokens[i].GetToken())
	}
}
