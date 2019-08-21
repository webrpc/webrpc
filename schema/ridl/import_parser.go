package ridl

func (p *parser) expectImportPath() error {
	// <value>
	value, err := p.expectPath()
	if err != nil {
		return err
	}

	p.emit(&ImportNode{path: newTokenNode(value)})

	return nil

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

		if err := p.expectImportPath(); err != nil {
			return p.stateError(err)
		}

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
			if err = expectWord(matches[0], "import"); err != nil {
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

		if err := p.expectImportPath(); err != nil {
			return p.stateError(err)
		}
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
