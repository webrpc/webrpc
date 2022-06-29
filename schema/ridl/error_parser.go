package ridl

import (
	"fmt"
)

func parserStateErrorDefinition(et *ErrorNode) parserState {
	return func(p *parser) parserState {
		var err error
		tok := p.cursor()

		switch tok.tt {

		case tokenQuote:
			tok, err = p.expectStringValue()
			if err != nil {
				return p.stateError(err)
			}
			et.message = newTokenNode(tok)

		case tokenMinusSign:
			matches, err := p.match(tokenDash, tokenDash, tokenWhitespace, tokenWord)
			if err != nil {
				return p.stateError(err)
			}
			et.httpStatus = newTokenNode(matches[3])

		case tokenWhitespace, tokenNewLine:
			p.next()

		case tokenHash:
			p.continueUntilEOL()

		default:
			p.emit(et)
			return parserDefaultState

		}

		return parserStateErrorDefinition(et)
	}
}

func parserStateError(p *parser) parserState {
	// error <code> <name> <message> [-- <http status code>]
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

	return parserStateErrorDefinition(&ErrorNode{
		code: newTokenNode(matches[2]),
		name: newTokenNode(matches[4]),
	})
}
