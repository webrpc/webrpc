package ridl

import "strings"

// tokenize
// Returns:
//   - []token: A slice containing all the tokens parsed from the source.
//   - map[int]string: A map where the key is the line number and the value is the comment present in that line.
//
// This function extracts tokens from it. The comments are associated with the line
// numbers where they occur, providing a convenient map for comment retrieval
func tokenize(src []byte) ([]token, map[int]string, error) {
	lx := newLexer(string(src))
	tokens := []token{}
	lineComments := make(map[int]string)

	commentLine := false
	commentTokens := []string{}

	for tok := range lx.tokens {
		if tok.tt == tokenEOF {
			break
		}

		if commentLine {
			if tok.tt != tokenNewLine && tok.tt != tokenWhitespace {
				commentTokens = append(commentTokens, tok.String())
			}
		}

		if tok.tt == tokenHash {
			commentLine = true
		}

		if tok.tt == tokenNewLine {
			if commentLine {
				lineComments[tok.line] = strings.Join(commentTokens, " ")
				commentTokens = []string{}
			}
			commentLine = false
		}

		tokens = append(tokens, tok)
	}

	return tokens, lineComments, nil
}
