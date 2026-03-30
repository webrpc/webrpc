package ridl

func parserStateTypeAliasMetaDefinition(ta *TypeAliasNode) parserState {
	return func(p *parser) parserState {
		// + <tag.name> = value
		_, err := p.match(tokenPlusSign, tokenWhitespace)
		if err != nil {
			return p.stateError(err)
		}

		left, err := p.expectMetadataKey()
		if err != nil {
			return p.stateError(err)
		}

		_, err = p.match(tokenWhitespace, tokenEqual, tokenWhitespace)
		if err != nil {
			return p.stateError(err)
		}

		right, err := p.expectMetadataValue()
		if err != nil {
			return p.stateError(err)
		}

		ta.meta = append(ta.meta, &DefinitionNode{
			leftNode:  newTokenNode(left),
			rightNode: newTokenNode(right),
		})

		return parserStateTypeAliasMeta(ta)
	}
}

func parserStateTypeAliasMeta(ta *TypeAliasNode) parserState {
	return func(p *parser) parserState {
		tok := p.cursor()

		switch tok.tt {
		case tokenNewLine, tokenWhitespace:
			p.next()

		case tokenHash:
			p.continueUntilEOL()

		case tokenPlusSign:
			return parserStateTypeAliasMetaDefinition(ta)

		default:
			p.emit(ta)
			return parserDefaultState
		}

		return parserStateTypeAliasMeta(ta)
	}
}

func parserStateTypeAlias(p *parser) parserState {
	// type <name>: <type>[<# comment>]
	matches, err := p.match(tokenWord, tokenWhitespace, tokenWord, tokenColon, tokenWhitespace)
	if err != nil {
		return p.stateError(err)
	}

	if matches[0].val != wordType {
		return p.stateError(errUnexpectedToken)
	}

	typeToken, err := p.expectType()
	if err != nil {
		return p.stateError(err)
	}

	ta := &TypeAliasNode{
		name:      newTokenNode(matches[2]),
		aliasType: newTokenNode(typeToken),
		comment:   parseComments(p.comments, matches[0].line),
		line:      matches[0].line,
	}

	return parserStateTypeAliasMeta(ta)
}
