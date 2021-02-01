package ridl

// type <name>: <type>
//  + <tag.name> = <VALUE>
//  + <tag.name2> = <VALUE>

func parserStateTypeDefinition(tn *TypeNode) parserState {
	return func(p *parser) parserState {
		tok := p.cursor()

		switch tok.tt {

		case tokenPlusSign:
			return parserStateTypeFieldMetaDefinition(tn)

		case tokenNewLine, tokenWhitespace:
			// case tokenWhitespace:
			p.next()

		case tokenHash:
			p.continueUntilEOL()

		default:
			p.emit(tn)
			return parserDefaultState

		}

		return parserStateTypeDefinition(tn)
	}
}

func parserStateTypeFieldMetaDefinition(tn *TypeNode) parserState {
	return func(p *parser) parserState {
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

		tn.extra.meta = append(tn.extra.meta, &DefinitionNode{
			leftNode:  newTokenNode(left),
			rightNode: newTokenNode(right),
		})

		return parserStateTypeFieldMeta(tn)
	}
}

func parserStateTypeFieldMeta(tn *TypeNode) parserState {
	return func(p *parser) parserState {

		tok := p.cursor()

		switch tok.tt {

		case tokenNewLine, tokenWhitespace:
			p.next()

		case tokenHash:
			p.continueUntilEOL()

		case tokenPlusSign:
			return parserStateTypeFieldMetaDefinition(tn)

		default:
			p.emit(tn)
			return parserDefaultState

		}

		return parserStateTypeFieldMeta(tn)
	}
}

func parserStateType(p *parser) parserState {
	// type <name>: <type>[<# comment>]
	matches, err := p.match(tokenWord, tokenWhitespace, tokenWord, tokenColon, tokenWhitespace, tokenWord, tokenEOL)
	if err != nil {
		return p.stateError(err)
	}

	if matches[0].val != wordType {
		return p.stateError(errUnexpectedToken)
	}

	return parserStateTypeDefinition(&TypeNode{
		name:     newTokenNode(matches[2]),
		typeType: newTokenNode(matches[5]),
		extra:    &DefinitionNode{},
	})
}
