package ridl

import (
	"fmt"
)

// error <code> <name> <message> [HTTP <status code>]
func parserStateError(p *parser) parserState {
	// error <code> <name>
	matches, err := p.match(tokenWord, tokenWhitespace, tokenWord, tokenWhitespace, tokenWord, tokenWhitespace)
	if err != nil {
		return p.stateError(err)
	}

	if matches[0].val != wordError {
		return p.stateError(errUnexpectedToken)
	}

	if err := expectNumber(matches[2], matches[2].val); err != nil {
		return p.stateError(fmt.Errorf("expecting error code to be a number but got '%s'", matches[2].val))
	}

	return parserStateErrorMessage(&ErrorNode{
		code: newTokenNode(matches[2]),
		name: newTokenNode(matches[4]),
	})
}

// <message> [HTTP <status code>]
func parserStateErrorMessage(et *ErrorNode) parserState {
	return func(p *parser) parserState {
		var err error
		tok := p.cursor()

		// message, err := p.expectLiteralValue()
		// if err != nil {
		// 	return p.stateError(err)
		// }

		switch tok.tt {

		// "<message>"
		case tokenQuote:
			tok, err = p.expectStringValue()
			if err != nil {
				return p.stateError(err)
			}
			et.message = newTokenNode(tok)

		// <message>
		case tokenWord:
			if tok.val == "HTTP" {
				p.rewind(1)
				return parserStateErrorExplicitStatusCode(et)
			}
			et.message = newTokenNode(tok)
			p.next()

		default:
			return p.stateError(fmt.Errorf("expected <message> but got %v", tok))
		}

		if et.httpStatus == nil {
			return parserStateErrorExplicitStatusCode(et)
		}

		if err := p.expectOptionalCommentOrEOL(); err != nil {
			return p.stateError(err)
		}

		p.emit(et)
		return parserDefaultState

	}
}

// [HTTP <status code>]
func parserStateErrorExplicitStatusCode(et *ErrorNode) parserState {
	return func(p *parser) parserState {
		// Try to match HTTP
		matches, err := p.match(tokenWhitespace, tokenWord)
		if err != nil {
			if err := p.expectOptionalCommentOrEOL(); err != nil {
				return p.stateError(err)
			}

			p.emit(et)
			return parserDefaultState
		}

		if err := expectWord(matches[1], "HTTP"); err != nil {
			p.rewind(1)
			return p.stateError(fmt.Errorf("expecting optional 'HTTP <status code>' but got '%s'", matches[1].val))
		}

		// Match <status code>
		matches, err = p.match(tokenWhitespace, tokenWord)
		if err != nil {
			return p.stateError(fmt.Errorf("expecting '<status code>': %w", err))
		}

		if err := expectNumber(matches[1], matches[1].val); err != nil {
			p.rewind(1)
			return p.stateError(fmt.Errorf("expecting HTTP '<status code>' to be a number but got '%s'", matches[1].val))
		}

		et.httpStatus = newTokenNode(matches[1])

		if et.message == nil {
			if _, err = p.match(tokenWhitespace); err != nil {
				return p.stateError(fmt.Errorf("expecting 'whitespace <message>': %w", err))
			}
			return parserStateErrorMessage(et)
		}

		if err := p.expectOptionalCommentOrEOL(); err != nil {
			return p.stateError(err)
		}

		p.emit(et)
		return parserDefaultState
	}
}
