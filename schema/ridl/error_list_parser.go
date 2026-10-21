package ridl

import (
	"errors"
	"fmt"
	"strings"
)

// emptyErrorNameError reports an empty segment in a bar-separated error list.
// offset locates the segment inside the word token, so the parse error points
// at that segment rather than the start of the whole list.
type emptyErrorNameError struct {
	word   string
	offset int
}

func (e *emptyErrorNameError) Error() string {
	return fmt.Sprintf("empty error name in '%s'", e.word)
}

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
		p.pointAtErrorName(matches[1], p.pos-1, err)
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
			p.pointAtErrorName(matches[3], p.pos-1, err)
			return nil, err
		}
		errorNames = append(errorNames, names...)
	}

	return errorNames, nil
}

// pointAtErrorName moves the cursor to the empty segment that splitErrorNames
// rejected, so the parse error reports its column instead of the line break.
func (p *parser) pointAtErrorName(tok *token, idx int, err error) {
	var empty *emptyErrorNameError
	if !errors.As(err, &empty) {
		return
	}
	if idx < 0 || idx >= p.length {
		return
	}
	p.pos = idx
	p.tokens[idx].col = tok.col - len(tok.val) + empty.offset
}

// splitErrorNames splits one word token on '|': a list without spaces around
// the bar, e.g. `errors A|B`, lexes as a single word. Each name carries the
// pos and col of its own segment.
func splitErrorNames(tok *token) ([]*TokenNode, error) {
	if !strings.Contains(tok.val, "|") {
		return []*TokenNode{newTokenNode(tok)}, nil
	}
	names := []*TokenNode{}
	offset := 0
	for i, val := range strings.Split(tok.val, "|") {
		if i > 0 {
			offset++ // the separating bar
		}
		if val == "" {
			return nil, &emptyErrorNameError{word: tok.val, offset: offset}
		}
		end := offset + len(val)
		names = append(names, newTokenNode(&token{
			tt:   tok.tt,
			val:  val,
			pos:  tok.pos - len(tok.val) + end,
			line: tok.line,
			col:  tok.col - len(tok.val) + end,
		}))
		offset = end
	}
	return names, nil
}
