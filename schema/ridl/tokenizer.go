package ridl

func tokenize(src []byte) ([]token, map[int][]string, error) {
	lx := newLexer(string(src))

	tokens := []token{}
	comments := make(map[int][]string)
	parseComment := false
	for tok := range lx.tokens {
		if tok.tt == tokenEOF {
			break
		}

		if parseComment {
			if tok.tt != tokenNewLine && tok.tt != tokenWhitespace {
				comments[tok.line] = append(comments[tok.line], tok.String())
			}
		}

		if tok.tt == tokenHash {
			parseComment = true
			comments[tok.line] = []string{}
		}

		if tok.tt == tokenNewLine {
			parseComment = false
		}

		tokens = append(tokens, tok)
	}

	return tokens, comments, nil
}
