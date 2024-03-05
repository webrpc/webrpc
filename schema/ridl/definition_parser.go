package ridl

func parserStateDefinitionValue(dn *DefinitionNode) parserState {
	return func(p *parser) parserState {
		value, err := p.expectLiteralValue()
		if err != nil {
			return p.stateError(err)
		}

		dn.rightNode = newTokenNode(value)

		p.emit(dn)

		return parserStateEOL
	}
}

func parserStateDefinition(p *parser) parserState {
	// <word><space>=<space>[<# comment>]<EOL>
	tokens, err := p.match(tokenWord, tokenWhitespace, tokenEqual, tokenWhitespace)
	if err != nil {
		return p.stateError(err)
	}

	return parserStateDefinitionValue(&DefinitionNode{
		leftNode: newTokenNode(tokens[0]),
		comment:  parseComments(p.comments, tokens[0].line),
	})
}
