package ridl

func tokenize(input string) ([]token, error) {
	lx := newLexer(string(input))

	tokens := []token{}
	for tok := range lx.tokens {
		if tok.tt == tokenEOF {
			break
		}
		tokens = append(tokens, tok)
	}

	return tokens, nil
}
