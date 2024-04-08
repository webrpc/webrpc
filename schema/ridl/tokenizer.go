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

	parsingComment := false
	commentTokens := []string{}

	for tok := range lx.tokens {
		if tok.tt == tokenEOF {
			break
		}

		// start parsing comment tokens until new line
		if tok.tt == tokenHash {
			parsingComment = true
			tokens = append(tokens, tok)
			continue
		}

		if parsingComment && tok.tt != tokenNewLine && tok.tt != tokenWhitespace {
			commentTokens = parseCommentToken(tok, tokens, commentTokens)
		}

		if tok.tt == tokenNewLine && parsingComment {
			lineComments[tok.line] = strings.Join(commentTokens, " ")
			commentTokens = []string{}
			parsingComment = false
		}

		tokens = append(tokens, tok)
	}

	return tokens, lineComments, nil
}

func parseCommentToken(curToken token, tokens []token, commentTokens []string) []string {
	tokenLen := len(tokens)

	if tokenLen == 0 {
		commentTokens = append(commentTokens, curToken.String())
		return commentTokens
	}

	// previous token was whitespace that means new word
	if tokens[tokenLen-1].tt == tokenWhitespace {
		commentTokens = append(commentTokens, curToken.String())
		return commentTokens
	}

	commentLen := len(commentTokens)
	if commentLen > 0 {
		// previous token was not whitespace that could be slash or any character
		// need to append current char to previous comment token
		// for instance https://www.google.com has `:`, `/`, `/` ....
		commentTokens[commentLen-1] += curToken.String()
	} else {
		commentTokens = append(commentTokens, curToken.String())
	}
	return commentTokens
}
