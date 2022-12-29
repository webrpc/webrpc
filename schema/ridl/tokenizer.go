package ridl

func tokenize(src []byte) ([]token, error) {
	lx := newLexer(string(src))

	tokens := []token{}
	for tok := range lx.tokens {
		if tok.tt == tokenEOF {
			break
		}
		tokens = append(tokens, tok)
	}

	return tokens, nil
}
