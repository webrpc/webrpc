package ridl

func parserStateEnumExplicitValue(en *EnumNode, dn *DefinitionNode) parserState {
	return func(p *parser) parserState {
		// [= <value>]
		_, err := p.match(tokenWhitespace, tokenEqual, tokenWhitespace)

		if err == nil {
			explicitValue, err := p.expectLiteralValue()
			if err != nil {
				return p.stateError(err)
			}
			dn.rightNode = newTokenNode(explicitValue)
			en.comment = parseComments(p.comments, explicitValue.line)
		}

		en.values = append(en.values, dn)

		return parserStateEnumDefinition(en)
	}
}

func parserStateEnumDefinition(et *EnumNode) parserState {
	return func(p *parser) parserState {
		tok := p.cursor()

		switch tok.tt {

		case tokenDash:
			// - <name> [= <value>] [<# comment>][EOL]
			matches, err := p.match(tokenDash, tokenWhitespace, tokenWord)

			if err != nil {
				return p.stateError(err)
			}

			// Collect additional tokenDot and tokenWord tokens to support
			// dot-separated names like "v1.5" (lexed as word "v1", dot ".", word "5").
			nameTokens := []*token{matches[2]}
			for {
				tok := p.cursor()
				if tok.tt == tokenDot || tok.tt == tokenWord {
					nameTokens = append(nameTokens, tok)
					p.next()
				} else {
					break
				}
			}
			nameTok, err := composedValue(nameTokens)
			if err != nil {
				return p.stateError(err)
			}

			return parserStateEnumExplicitValue(et, &DefinitionNode{
				leftNode: newTokenNode(nameTok),
				comment:  parseComments(p.comments, matches[0].line),
			})

		case tokenNewLine, tokenWhitespace:
			p.next()

		case tokenHash:
			p.continueUntilEOL()

		default:
			p.emit(et)
			return parserDefaultState

		}

		return parserStateEnumDefinition(et)
	}
}

func parserStateEnum(p *parser) parserState {
	// enum <name>: <type>[<# comment>]
	matches, err := p.match(tokenWord, tokenWhitespace, tokenWord, tokenColon, tokenWhitespace, tokenWord, tokenEOL)
	if err != nil {
		return p.stateError(err)
	}

	if matches[0].val != wordEnum {
		return p.stateError(errUnexpectedToken)
	}

	return parserStateEnumDefinition(&EnumNode{
		name:     newTokenNode(matches[2]),
		enumType: newTokenNode(matches[5]),
		values:   []*DefinitionNode{},
		comment:  parseComments(p.comments, matches[0].line),
		line:     matches[0].line,
	})
}
