package ridl

// expectErrorList parses a bar-separated list of error names
// Example: ErrorName1 | ErrorName2 | ErrorName3
func (p *parser) expectErrorList() ([]*TokenNode, error) {
	errorNames := []*TokenNode{}

	// Expect at least one error name
	matches, err := p.match(tokenWhitespace, tokenWord)
	if err != nil {
		return nil, err
	}

	errorNames = append(errorNames, newTokenNode(matches[1]))

	// Parse additional bar-separated error names
	for {
		matches, err := p.match(tokenWhitespace, tokenBar, tokenWhitespace, tokenWord)
		if err != nil {
			break // No more error names
		}
		errorNames = append(errorNames, newTokenNode(matches[3]))
	}

	return errorNames, nil
}
