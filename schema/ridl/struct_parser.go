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
			leftNode:    newTokenNode(matches[2]),
			comment:     parseComments(p.comments, matches[0].line),
			annotations: mn.fieldAnnotations,
		}
		mn.fieldAnnotations = nil

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

		case tokenAt:
			anns, err := parseAnnotations(p)
			if err != nil {
				return p.stateError(err)
			}
			mn.fieldAnnotations = append(mn.fieldAnnotations, anns...)
			if err := checkDuplicateAnnotations(mn.fieldAnnotations); err != nil {
				return p.stateError(err)
			}

		case tokenDash:
			return parserStateStructFieldDefinition(mn)

		default:
			// Any buffered @tag lines here aren't for a field of this struct
			// (no "- name: type" followed them) — they belong to whatever
			// top-level definition comes next.
			if len(mn.fieldAnnotations) > 0 {
				p.pendingAnnotations = append(p.pendingAnnotations, mn.fieldAnnotations...)
				mn.fieldAnnotations = nil
			}
			p.emit(mn)
			return parserDefaultState

		}

		return parserStateStructField(mn)
	}
}

func parserStateStruct(p *parser) parserState {
	annotations := p.pendingAnnotations
	p.pendingAnnotations = nil

	// struct <name>
	matches, err := p.match(tokenWord, tokenWhitespace, tokenWord)
	if err != nil {
		return p.stateError(err)
	}

	if matches[0].val != wordStruct {
		return p.stateError(errUnexpectedToken)
	}

	return parserStateStructField(&StructNode{
		name:        newTokenNode(matches[2]),
		fields:      []*DefinitionNode{},
		comment:     parseComments(p.comments, matches[0].line),
		line:        matches[0].line,
		annotations: annotations,
	})
}
