package ridl

func parserStateMessageFieldMetaDefinition(mn *MessageNode) parserState {
	return func(p *parser) parserState {
		// add to latest field
		field := mn.fields[len(mn.fields)-1]

		// + <tag.name> = value
		_, err := p.match(tokenPlusSign, tokenWhitespace)
		if err != nil {
			return p.stateError(err)
		}

		// tag
		left, err := p.expectMetadataKey()
		if err != nil {
			return p.stateError(err)
		}

		// =
		_, err = p.match(tokenWhitespace, tokenEqual, tokenWhitespace)
		if err != nil {
			return p.stateError(err)
		}

		// - or value
		right, err := p.expectMetadataValue()
		if err != nil {
			return p.stateError(err)
		}

		field.meta = append(field.meta, &DefinitionNode{
			leftNode:  newTokenNode(left),
			rightNode: newTokenNode(right),
		})

		return parserStateMessageFieldMeta(mn)
	}
}

func parserStateMessageFieldMeta(mn *MessageNode) parserState {
	return func(p *parser) parserState {

		tok := p.cursor()

		switch tok.tt {

		case tokenNewLine, tokenWhitespace:
			p.next()

		case tokenHash:
			p.continueUntilEOL()

		case tokenPlusSign:
			return parserStateMessageFieldMetaDefinition(mn)

		default:
			return parserStateMessageField(mn)

		}

		return parserStateMessageFieldMeta(mn)
	}
}

func parserStateMessageFieldDefinition(mn *MessageNode) parserState {
	return func(p *parser) parserState {
		// - <name>: <type> [<# comment>][EOL]
		matches, err := p.match(tokenDash, tokenWhitespace, tokenWord)
		if err != nil {
			return p.stateError(err)
		}

		field := &DefinitionNode{
			leftNode: newTokenNode(matches[2]),
		}

		// ?
		_, err = p.match(tokenQuestionMark)
		if err == nil {
			// got question mark
			field.optional = true
		}

		// :
		_, err = p.match(tokenColon, tokenWhitespace)
		if err != nil {
			return p.stateError(err)
		}

		typeToken, err := p.expectType()
		if err != nil {
			return p.stateError(err)
		}
		field.rightNode = newTokenNode(typeToken)

		mn.fields = append(mn.fields, field)

		return parserStateMessageFieldMeta(mn)
	}
}

func parserStateMessageField(mn *MessageNode) parserState {
	return func(p *parser) parserState {
		tok := p.cursor()

		switch tok.tt {

		case tokenNewLine, tokenWhitespace:
			p.next()

		case tokenHash:
			p.continueUntilEOL()

		case tokenDash:
			return parserStateMessageFieldDefinition(mn)

		default:
			p.emit(mn)
			return parserDefaultState

		}

		return parserStateMessageField(mn)
	}
}

func parserStateMessage(p *parser) parserState {
	// message <name>
	matches, err := p.match(tokenWord, tokenWhitespace, tokenWord)
	if err != nil {
		return p.stateError(err)
	}

	if matches[0].val != wordMessage {
		return p.stateError(errUnexpectedToken)
	}

	return parserStateMessageField(&MessageNode{
		name:   newTokenNode(matches[2]),
		fields: []*DefinitionNode{},
	})
}
