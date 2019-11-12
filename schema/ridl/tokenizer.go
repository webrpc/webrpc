package ridl

import (
	"io"
	"io/ioutil"
)

func tokenize(r io.Reader) ([]token, error) {
	in, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	lx := newLexer(string(in))

	tokens := []token{}
	for tok := range lx.tokens {
		if tok.tt == tokenEOF {
			break
		}
		tokens = append(tokens, tok)
	}

	return tokens, nil
}
