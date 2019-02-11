package ridl

import (
	"log"
)

type Tree struct {
}

type Token struct {
}

func Tokenize(input string) (*Tree, error) {
	lx := Lex(string(input))
	for {
		tok := <-lx.tokens
		if tok.tt == tokenSpace {
			continue
		}
		if tok.tt == tokenEOF {
			break
		}
		if tok.tt == tokenNewLine {
			log.Printf("newline")
			continue
		}
		log.Printf("%v", tok)
	}
	return nil, nil
}
