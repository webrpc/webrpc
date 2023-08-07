package ridl

// alias <name>: <type>
//  + <tag.name> = <VALUE>
//  + <tag.name2> = <VALUE>

func parserStateAliasDefinition(tn *AliasNode) parserState {
	return func(p *parser) parserState {
		tok := p.cursor()

		switch tok.tt {

		case tokenPlusSign:
			return parserStateAliasFieldMetaDefinition(tn)

		case tokenNewLine, tokenWhitespace:
			// case tokenWhitespace:
			p.next()

		case tokenHash:
			p.continueUntilEOL()

		default:
			p.emit(tn)
			return parserDefaultState

		}

		return parserStateAliasDefinition(tn)
	}
}

func parserStateAliasFieldMetaDefinition(tn *AliasNode) parserState {
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

		return parserStateAliasFieldMeta(tn)
	}
}

func parserStateAliasFieldMeta(tn *AliasNode) parserState {
	return func(p *parser) parserState {

		tok := p.cursor()

		switch tok.tt {

		case tokenNewLine, tokenWhitespace:
			p.next()

		case tokenHash:
			p.continueUntilEOL()

		case tokenPlusSign:
			return parserStateAliasFieldMetaDefinition(tn)

		default:
			p.emit(tn)
			return parserDefaultState

		}

		return parserStateAliasFieldMeta(tn)
	}
}

func parserStateType(p *parser) parserState {
	// alias <name>: <type>[<# comment>]
	matches, err := p.match(tokenWord, tokenWhitespace, tokenWord, tokenColon, tokenWhitespace, tokenWord, tokenEOL)
	if err != nil {
		return p.stateError(err)
	}

	if matches[0].val != wordAlias {
		return p.stateError(errUnexpectedToken)
	}

	return parserStateAliasDefinition(&AliasNode{
		name:      newTokenNode(matches[2]),
		aliasType: newTokenNode(matches[5]),
		extra:     &DefinitionNode{},
	})
}
