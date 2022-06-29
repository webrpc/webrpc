package ridl

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

func parserStateErrorDefinition(et *ErrorNode) parserState {
	return func(p *parser) parserState {
		tok := p.cursor()

		switch tok.tt {

		case tokenNewLine, tokenWhitespace:
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
	matches, err := p.match(tokenWord, tokenWhitespace, tokenWord, tokenWhitespace, tokenWord, tokenWhitespace, tokenWord, tokenEOL)
	if err != nil {
		return p.stateError(err)
	}

	spew.Dump(matches)

	if matches[0].val != wordError {
		return p.stateError(errUnexpectedToken)
	}

	if err := expectNumber(matches[2], matches[2].val); err != nil {
		return p.stateError(fmt.Errorf("expecting error code to be a number but got '%s'", matches[2].val))
	}

	return parserStateErrorDefinition(&ErrorNode{
		code:    newTokenNode(matches[2]),
		name:    newTokenNode(matches[4]),
		message: newTokenNode(matches[6]),
		// TODO: rest.. etc..
	})
}
