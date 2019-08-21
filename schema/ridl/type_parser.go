package ridl

import (
	"errors"
)

func (p *parser) expectMapDefinition() (*token, error) {

	tok := p.cursor()
	if tok.val != "map" {
		return nil, errors.New(`expecting "map"`)
	}
	p.next()

	tokens := []*token{tok}

loop:
	for {
		tok := p.cursor()

		switch tok.tt {

		case tokenWhitespace:
			p.next()
			continue loop

		case tokenCloseAngleBracket:
			p.next()
			break loop

		case tokenOpenAngleBracket:
			types, err := p.expectTypeList()
			if err != nil {
				return nil, err
			}

			tokens = append(tokens, types)
			continue loop

		default:
			return nil, errUnexpectedToken(tok)
		}
	}

	return composedValue(tokens)
}

func (p *parser) expectArgumentList() ([]*ArgumentNode, error) {
	values := []*ArgumentNode{}

	matches, err := p.match(tokenWord, tokenWhitespace)
	if err == nil {
		if matches[0].val != "stream" {
			return nil, errUnexpectedToken(matches[0])
		}
	}

	tok, err := p.match(tokenOpenParen)
	if err != nil {
		return nil, err
	}

	tokens := []*token{tok[0]}

loop:
	for {
		tok := p.cursor()

		switch tok.tt {

		case tokenWhitespace:
			p.next()
			continue

		case tokenCloseParen:
			tokens = append(tokens, tok)
			p.next()
			break loop

		case tokenComma:
			tokens = append(tokens, tok)
			p.next()

		case tokenWord:
			var argument []*token
			var name *token

			optional := false

			matches, err := p.match(tokenWord, tokenQuestionMark, tokenColon, tokenWhitespace)
			if err == nil {
				argument = []*token{matches[0], matches[1], matches[2]}
				name = matches[0]
				optional = true
			} else {
				matches, err = p.match(tokenWord, tokenColon, tokenWhitespace)
				if err != nil {
					return nil, err
				}
				argument = []*token{matches[0], matches[1]}
				name = matches[0]
			}

			varType, err := p.expectType()
			if err != nil {
				return nil, err
			}

			values = append(values, &ArgumentNode{
				name:         newTokenNode(name),
				argumentType: newTokenNode(varType),
				optional:     optional,
			})

			tokens = append(tokens, append(argument, varType)...)

		default:
			return nil, errUnexpectedToken(tok)

		}
	}

	//return composedValue(tokens)

	return values, nil
}

func (p *parser) expectTypeList() (*token, error) {

	tok, err := p.match(tokenOpenAngleBracket)
	if err != nil {
		return nil, err
	}

	tokens := []*token{tok[0]}

loop:
	for {
		tok := p.cursor()

		switch tok.tt {

		case tokenWhitespace:
			p.next()
			continue

		case tokenCloseAngleBracket:
			tokens = append(tokens, tok)
			break loop

		case tokenComma:
			tokens = append(tokens, tok)
			p.next()

		case tokenWord, tokenOpenBracket:
			tok, err := p.expectType()
			if err != nil {
				return nil, err
			}

			tokens = append(tokens, tok)

		default:
			return nil, errUnexpectedToken(tok)
		}
	}

	return composedValue(tokens)
}

func (p *parser) expectType() (*token, error) {
	tokens := []*token{}

loop:
	for {
		tok := p.cursor()

		switch tok.tt {

		case tokenNewLine, tokenWhitespace, tokenComma, tokenCloseAngleBracket, tokenCloseParen, tokenHash:
			break loop

		case tokenEOF:
			return nil, errUnexpectedEOF

		case tokenOpenBracket:
			brackets, err := p.match(tokenOpenBracket, tokenCloseBracket)
			if err != nil {
				return nil, err
			}

			typeName, err := p.expectType()
			if err != nil {
				return nil, err
			}

			tokens = append(tokens, append(brackets, typeName)...)
			continue loop

		case tokenWord:

			switch tok.val {

			case "map":
				var err error

				tok, err = p.expectMapDefinition()
				if err != nil {
					return nil, err
				}

				tokens = append(tokens, tok)
				continue loop
			}

			tokens = append(tokens, tok)
			p.next()
			continue

		default:
			return nil, errUnexpectedToken(tok)
		}

		p.next()
	}

	return composedValue(tokens)
}
