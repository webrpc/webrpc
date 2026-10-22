package ridl

import (
	"fmt"
	"strings"
)

// expectErrorList parses a bar-separated list of error names
// Example: ErrorName1 | ErrorName2 | ErrorName3
func (p *parser) expectErrorList() ([]*TokenNode, error) {
	errorNames := []*TokenNode{}

	// Expect at least one error name
	matches, err := p.match(tokenWhitespace, tokenWord)
	if err != nil {
		return nil, err
	}
	names, err := splitErrorNames(matches[1])
	if err != nil {
		p.pos--
		return nil, err
	}
	errorNames = append(errorNames, names...)

	// Parse additional bar-separated error names
	for {
		matches, err := p.match(tokenWhitespace, tokenBar, tokenWhitespace, tokenWord)
		if err != nil {
			break // No more error names
		}
		names, err := splitErrorNames(matches[3])
		if err != nil {
			p.pos--
			return nil, err
		}
		errorNames = append(errorNames, names...)
	}

	return errorNames, nil
}

// splitErrorNames splits one word token on '|': a list without spaces around
// the bar, e.g. `+ error A|B`, lexes as a single word.
func splitErrorNames(tok *token) ([]*TokenNode, error) {
	if !strings.Contains(tok.val, "|") {
		return []*TokenNode{newTokenNode(tok)}, nil
	}
	names := []*TokenNode{}
	for _, val := range strings.Split(tok.val, "|") {
		if val == "" {
			return nil, fmt.Errorf("empty error name in '%s'", tok.val)
		}
		names = append(names, newTokenNode(&token{
			tt:   tok.tt,
			val:  val,
			pos:  tok.pos,
			line: tok.line,
			col:  tok.col,
		}))
	}
	return names, nil
}
