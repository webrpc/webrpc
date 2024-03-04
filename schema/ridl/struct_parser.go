package ridl

func parserStateStructFieldMetaDefinition(mn *StructNode) parserState {
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

		return parserStateStructFieldMeta(mn)
	}
}

func parserStateStructFieldMeta(mn *StructNode) parserState {
	return func(p *parser) parserState {

		tok := p.cursor()

		switch tok.tt {

		case tokenNewLine, tokenWhitespace:
			p.next()

		case tokenHash:
			p.continueUntilEOL()

		case tokenPlusSign:
			return parserStateStructFieldMetaDefinition(mn)

		default:
			return parserStateStructField(mn)

		}

		return parserStateStructFieldMeta(mn)
	}
}

func parserStateStructFieldDefinition(mn *StructNode) parserState {
	return func(p *parser) parserState {
		// - <name>: <type> [<# comment>][EOL]
		matches, err := p.match(tokenDash, tokenWhitespace, tokenWord)
		if err != nil {
			return p.stateError(err)
		}

		field := &DefinitionNode{
			leftNode: newTokenNode(matches[2]),
			comments: parseComments(p.comments, matches[0].line),
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

		return parserStateStructFieldMeta(mn)
	}
}

func parserStateStructField(mn *StructNode) parserState {
	return func(p *parser) parserState {
		tok := p.cursor()

		switch tok.tt {

		case tokenNewLine, tokenWhitespace:
			p.next()

		case tokenHash:
			p.continueUntilEOL()

		case tokenDash:
			return parserStateStructFieldDefinition(mn)

		default:
			p.emit(mn)
			return parserDefaultState

		}

		return parserStateStructField(mn)
	}
}

func parserStateStruct(p *parser) parserState {
	// struct <name>
	matches, err := p.match(tokenWord, tokenWhitespace, tokenWord)
	if err != nil {
		return p.stateError(err)
	}

	if matches[0].val != wordStruct {
		return p.stateError(errUnexpectedToken)
	}

	return parserStateStructField(&StructNode{
		name:     newTokenNode(matches[2]),
		fields:   []*DefinitionNode{},
		comments: parseComments(p.comments, matches[0].line),
	})
}
