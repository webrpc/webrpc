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
		if tok.tt == tokenEOF {
			break
		}
		if tok.tt == tokenNewLine || tok.tt == tokenSpace {
			continue
		}
		log.Printf("tok: %#v", tok)
	}
	return nil, nil
}
