package ridl

func parserStateImportMembers(p *parser, in *ImportNode) parserState {

	tok := p.cursor()

	switch tok.tt {
	case tokenNewLine, tokenWhitespace:
		// ignore whitespace
		p.next()

	case tokenDash:

		// - <value> [<# comment>][EOL]
		if _, err := p.match(tokenDash, tokenWhitespace); err != nil {
			return p.stateError(err)
		}

		member, err := p.expectLiteralValue()
		if err != nil {
			return p.stateError(err)
		}

		in.members = append(in.members, newTokenNode(member))

	case tokenHash:
		p.continueUntilEOL()

	default:
		p.emit(in)

		return parserDefaultState

	}

	return parserStateImportMembers(p, in)
}

func parserStateImportValue(p *parser) parserState {
	tok := p.cursor()

	switch tok.tt {

	case tokenNewLine, tokenWhitespace:
		// ignore whitespace
		p.next()

	case tokenDash:

		// - <value> [<# comment>][EOL]
		if _, err := p.match(tokenDash, tokenWhitespace); err != nil {
			return p.stateError(err)
		}

		path, err := p.expectPath()
		if err != nil {
			return p.stateError(err)
		}

		p.emit(&ImportNode{
			path:    newTokenNode(path),
			members: []*TokenNode{},
		})

		return parserStateImportValue

	case tokenHash:
		p.continueUntilEOL()

	default:
		return parserDefaultState

	}

	return parserStateImportValue
}

func parserStateImport(p *parser) parserState {

	{
		// import[<space>][<# comment>]<EOL>
		matches, err := p.match(tokenWord, tokenEOL)
		if err == nil {
			if err = expectWord(matches[0], wordImport); err != nil {
				return p.stateError(err)
			}
			return parserStateImportValue
		}
	}

	{
		// import<space><import path><# comment>]<EOL>
		_, err := p.match(tokenWhitespace)
		if err != nil {
			return p.stateError(err)
		}

		path, err := p.expectPath()
		if err != nil {
			return p.stateError(err)
		}

		importNode := &ImportNode{
			path: newTokenNode(path),
		}

		return parserStateImportMembers(p, importNode)
	}

	return parserDefaultState

}

func (p *parser) expectPath() (*token, error) {
	tokens := []*token{}

loop:
	for {
		// TODO: should we allow enclosing literal values with quotes?
		tok := p.cursor()

		switch tok.tt {

		// literal value enclosed by string
		case tokenQuote:
			var err error
			tok, err = p.expectStringValue()
			if err != nil {
				return nil, err
			}
			return tok, nil

		case tokenWord, tokenDot, tokenSlash:
			tokens = append(tokens, tok)
			p.next()

		default:
			break loop
		}
	}

	return composedValue(tokens)
}

func (p *parser) expectMetadataValue() (*token, error) {
	tokens := []*token{}

loop:
	for {
		// TODO: should we allow enclosing literal values with quotes?
		tok := p.cursor()

		switch tok.tt {

		case tokenWord, tokenDash, tokenDot, tokenComma:
			tokens = append(tokens, tok)
			p.next()

		default:
			break loop
		}
	}

	return composedValue(tokens)
}
