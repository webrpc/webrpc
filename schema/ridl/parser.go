package ridl

import (
	"errors"
	"fmt"
	"log"
)

var (
	errUnexpectedEOF = errors.New(`unexpected EOF`)
)

type definition struct {
	left  *token
	right *token
}

func (d *definition) value() string {
	return d.right.val
}

type tree struct {
	definitions map[string]*definition

	imports []*token
}

var invalidToken = &token{tt: tokenInvalid}

type parseState func(*parser) parseState

type parser struct {
	tokens []token
	length int
	pos    int

	lastErr error

	tree tree
}

func newParser(input string) (*parser, error) {
	tokens, err := tokenize(input)
	if err != nil {
		return nil, err
	}
	p := &parser{
		tokens: tokens,
		length: len(tokens),
		tree: tree{
			definitions: make(map[string]*definition),
			imports:     []*token{},
		},
	}
	return p, nil
}

func (p *parser) run() error {
	for state := parseStateStart; state != nil; {
		state = state(p)
	}
	log.Printf("p: %#v", p)
	if p.lastErr != nil {
		return p.lastErr
	}
	return nil
}

func (p *parser) continueUntilEOL() error {
	for {
		tok := p.cursor()
		log.Printf("drop: %v", tok)
		if tok.tt == tokenNewLine {
			// EOL
			return nil
		}

		if !p.next() {
			// EOF
			return nil
		}
	}

	panic("reached")
}

func (p *parser) expectCommentOrNewLine() error {
	for {

		tok := p.cursor()
		log.Printf("tok: %q", tok.val)

		switch tok.tt {
		case tokenNewLine:
			// got new line, continue
		case tokenHash:
			if !p.next() {
				return nil
			}
			return p.continueUntilEOL()
		default:
			// another kind of token, stop
			return nil
		}
		if !p.next() {
			return nil
		}
	}
	return nil
}

func (p *parser) expectNext(tt tokenType) error {
	if !p.next() {
		return errors.New(`unexpected EOF`)
	}

	tok := p.cursor()
	if tok.tt != tt {
		return fmt.Errorf(`expecting %v, got %v`, tt, tok)
	}
	return nil
}

func (p *parser) stateError(err error) parseState {
	p.lastErr = fmt.Errorf("error: %v, near: %v \n", err, p.cursor())
	return nil
}

func (p *parser) next() bool {
	if p.pos >= p.length {
		return false
	}
	p.pos = p.pos + 1
	return true
}

func (p *parser) cursor() *token {
	if p.pos >= p.length {
		return invalidToken
	}
	return &p.tokens[p.pos]
}

func parseStateNewLine(p *parser) parseState {
	for p.next(); p.cursor().tt == tokenNewLine; {
	}
	return parseStateStart
}

func parseStateContinue(p *parser) parseState {
	if !p.next() {
		return parseStateEOF
	}
	return parseStateStart
}

func parseStateDefinition(word *token) parseState {
	return func(p *parser) parseState {
		if err := p.expectNext(tokenEqual); err != nil {
			return p.stateError(err)
		}
		if err := p.expectNext(tokenWord); err != nil {
			return p.stateError(err)
		}

		if p.tree.definitions[word.val] != nil {
			return p.stateError(fmt.Errorf(`can't redefine "%v", previous definition found at %v`, word.val, p.tree.definitions[word.val].left))
		}
		p.tree.definitions[word.val] = &definition{
			left:  word,
			right: p.cursor(),
		}
		log.Printf("definition")
		return parseStateEOL
	}
}

func parseStateImport(p *parser) parseState {
	if !p.next() {
		return parseStateUnexpectedEOF
	}

	for {
		if err := p.expectCommentOrNewLine(); err != nil {
			return p.stateError(err)
		}

		tok := p.cursor()
		log.Printf("inspect: %v", tok)
		switch tok.tt {
		case tokenNewLine:
			if !p.next() {
				return nil
			}
		case tokenMinusSign:
			log.Printf("got minus")
			if err := p.expectNext(tokenWord); err != nil {
				return p.stateError(err)
			}
			p.tree.imports = append(p.tree.imports, p.cursor())
			if !p.next() {
				return nil
			}
		default:
			return parseStateContinue
		}
	}

	log.Fatalf("got: %v", p.cursor())
	return nil
}

func parseStateEOL(p *parser) parseState {
	if err := p.expectCommentOrNewLine(); err != nil {
		return p.stateError(err)
	}
	return parseStateContinue
}

func parseStateComment(p *parser) parseState {
	if err := p.continueUntilEOL(); err != nil {
		return p.stateError(err)
	}
	return parseStateContinue
}

func parseStateLine(p *parser) parseState {
	// word is the beginning of a line
	word := p.cursor()
	log.Printf("word: %v", word)

	switch word.val {
	case "ridl", "service", "version":
		return parseStateDefinition(word)
	case "import":
		return parseStateImport
	default:
		return p.stateError(errors.New("unknown definition"))
	}

	return parseStateStart
}

func parseStateStart(p *parser) parseState {
	tok := p.cursor()
	switch tok.tt {
	case tokenNewLine:
		return parseStateNewLine
	case tokenWord:
		return parseStateLine
	case tokenHash:
		return parseStateComment
	case tokenInvalid: // abrupt EOF
		return nil
	}
	log.Fatalf("parse state: %v", tok)
	return nil
}

func parseStateUnexpectedEOF(p *parser) parseState {
	return p.stateError(errUnexpectedEOF)
}

func parseStateEOF(p *parser) parseState {
	return nil
}
