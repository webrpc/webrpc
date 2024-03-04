package ridl

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"runtime"
	"strings"
)

const (
	wordError   = "error"
	wordEnum    = "enum"
	wordImport  = "import"
	wordMap     = "map"
	wordStruct  = "struct"
	wordName    = "name"
	wordProxy   = "proxy"
	wordService = "service"
	wordStream  = "stream"
	wordVersion = "version"
	wordWebRPC  = "webrpc"
)

var (
	errUnexpectedToken = errors.New(`unexpected token`)
	errUnexpectedEOF   = errors.New(`unexpected EOF`)
	errUnexpectedEOL   = errors.New(`unexpected EOL`)
)

var eofToken = &token{tt: tokenEOF}

type parserState func(*parser) parserState

type parser struct {
	tokens []token
	length int
	pos    int

	words    chan interface{}
	comments map[int][]string

	root RootNode
}

func newParser(src []byte) (*parser, error) {
	tokens, comments, err := tokenize(src)
	if err != nil {
		return nil, err
	}

	p := &parser{
		words:    make(chan interface{}),
		tokens:   tokens,
		length:   len(tokens),
		comments: comments,
	}
	return p, nil
}

func (p *parser) emit(message interface{}) {
	p.words <- message
}

func (p *parser) collect() error {
	for tok := range p.words {
		switch word := tok.(type) {

		case error:
			return word

		case Node:
			p.root.Push(word)

		case tokenType:
			if tok == tokenEOF {
				return nil
			}

		default:
			log.Fatalf("unexpected: %v (%T)", word, word)
		}
	}

	return nil
}

func (p *parser) run() error {
	go func() {
		state := parserDefaultState
		for state != nil {
			fname := runtime.FuncForPC(reflect.ValueOf(state).Pointer()).Name()
			_ = fname
			//log.Printf("state: %v", fname)
			state = state(p)
		}
	}()

	err := p.collect()
	if err != nil {
		return err
	}

	return nil
}

func (p *parser) continueUntilEOL() error {
	words := []string{}

	for {
		tok := p.cursor()

		switch tok.tt {
		case tokenNewLine, tokenEOF:
			return nil
		case tokenWord:
			words = append(words, tok.String())
		}

		p.next()
	}

	panic("unreachable")
}

func (p *parser) expectOptionalCommentOrEOL() error {
	for {
		tok := p.cursor()

		switch tok.tt {

		case tokenWhitespace:
			p.next()
			continue

		case tokenEOF, tokenNewLine:
			// end of line
			return nil

		case tokenHash:
			return p.continueUntilEOL()

		default:
			// another kind of token, stop
			return errUnexpectedToken
		}

		panic("unreachable")
	}

	return nil
}

func (p *parser) stateError(err error) parserState {
	cur := p.cursor()
	err = fmt.Errorf("%d:%d: error near %q: %v", cur.line, cur.col, cur.val, err)
	p.emit(err)
	return nil
}

func (p *parser) rewind(n int) bool {
	if p.pos-n < 0 {
		return false
	}
	p.pos = p.pos - n
	return true
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
		return eofToken
	}
	return &p.tokens[p.pos]
}

func (p *parser) expectStringValue() (*token, error) {
	tok := p.cursor()
	p.next()

	if tok.tt != tokenQuote {
		return nil, errors.New("strings must start with a quote")
	}

	tokens := []*token{}

loop:
	for {
		tok := p.cursor()
		p.next()

		switch tok.tt {

		case tokenNewLine:
			return nil, errUnexpectedEOL

		case tokenEOF:
			return nil, errUnexpectedEOF

		case tokenBackslash:
			// push backlash
			tokens = append(tokens, tok)

			// push whatever else comes next
			tok = p.cursor()
			p.next()

		case tokenQuote:
			// end of string
			break loop
		}

		tokens = append(tokens, tok)
	}

	return composedValue(tokens)
}

func (p *parser) expectLiteralValue() (*token, error) {
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

		case tokenWord:
			tokens = append(tokens, tok)
			p.next()

		default:
			break loop
		}
	}

	return composedValue(tokens)
}

func (p *parser) match(tokenTypes ...tokenType) ([]*token, error) {
	matcher := func(p *parser) ([]*token, error) {
		tokens := make([]*token, 0, len(tokenTypes))
		matches := 0

		for _, expecting := range tokenTypes {
			tok := p.cursor()

			if tok.tt == expecting {
				tokens = append(tokens, tok)
				p.next()
				matches++
				continue
			}

			if expecting == tokenEOL {
				err := p.expectOptionalCommentOrEOL()
				if err != nil {
					return nil, err
				}
				tokens = append(tokens, &token{
					tt:   tokenEOL,
					line: tok.line,
					col:  tok.col,
				})
				break
			}

			if expecting != tokenWhitespace {
				p.rewind(matches)
				return nil, fmt.Errorf(`expecting %v, got %v`, expecting, tok)
			}

			tokens = append(tokens, &token{
				tt:   tokenWhitespace,
				val:  " ",
				line: tok.line,
				col:  tok.col,
			})

		}
		return tokens, nil
	}

	return matcher(p)
}

func (p *parser) expectMetadataKey() (*token, error) {
	matches, err := p.match(tokenWord)
	if err != nil {
		return nil, err
	}

	tokens := []*token{matches[0]}

loop:
	for {
		tok := p.cursor()

		switch tok.tt {
		case tokenQuestionMark, tokenWhitespace, tokenEqual:
			break loop
		case tokenWord:
			tokens = append(tokens, tok)
		case tokenDot:
			matches, err := p.match(tokenDot, tokenWord)
			if err != nil {
				return nil, err
			}

			tokens = append(tokens, matches...)
			continue
		default:
			return nil, errUnexpectedToken
		}

		p.next()
	}

	return composedValue(tokens)
}

func (p *parser) expectMetadataValue() (*token, error) {
	tok := p.cursor()

	if tok.tt == tokenQuote {
		var err error
		tok, err = p.expectStringValue()
		if err != nil {
			return nil, err
		}
		return tok, nil
	}

	tokens := []*token{}
loop:
	for {
		tok := p.cursor()

		switch tok.tt {
		case tokenWhitespace, tokenNewLine, tokenEOF:
			break loop
		}

		tokens = append(tokens, tok)
		p.next()
	}

	return composedValue(tokens)
}

func parserStateEOL(p *parser) parserState {
	err := p.expectOptionalCommentOrEOL()
	if err != nil {
		return p.stateError(err)
	}

	return parserStateContinue
}

func parserStateComment(p *parser) parserState {
	if err := p.continueUntilEOL(); err != nil {
		return p.stateError(err)
	}
	return parserDefaultState
}

func parserStateDeclaration(p *parser) parserState {
	word := p.cursor()
	if word.tt != tokenWord {
		return p.stateError(errUnexpectedToken)
	}

	switch word.val {
	case wordWebRPC, wordName, wordVersion:
		// <word> = <value> # optional comment
		return parserStateDefinition
	case wordImport:
		// import
		//   - <value> [<# comment>]
		return parserStateImport
	case wordError:
		// error <code> <name> <message> [HTTP <status code>]
		return parserStateError
	case wordEnum:
		// enum <name>: <type>
		//   - <name>[<space>=<space><value>][<#comment>]
		return parserStateEnum
	case "message":
		// Deprecated in v0.9.0.
		return p.stateError(fmt.Errorf("keyword \"message\" was renamed to \"struct\", see https://github.com/webrpc/webrpc/blob/master/CHANGELOG.md#ridl-v090-migration-guide"))
	case wordStruct:
		// struct <name>
		//   - <name>: <type>
		//     + <tag.name> = <VALUE>
		return parserStateStruct
	case wordService:
		// service <name>
		//   - <name>([arguments]) [=> ([arguments])]
		return parserStateService
	default:
		return p.stateError(errUnexpectedToken)
	}

	return parserDefaultState
}

func parserDefaultState(p *parser) parserState {
	tok := p.cursor()

	switch tok.tt {

	case tokenWhitespace:
		return parserStateSpace

	case tokenNewLine:
		return parserStateNewLine

	case tokenWord:
		return parserStateDeclaration

	case tokenHash:
		return parserStateComment

	case tokenEOF:
		return parserStateEOF

	default:
		return p.stateError(errUnexpectedToken)
	}

	return nil
}

func parserStateUnexpectedEOF(p *parser) parserState {
	return p.stateError(errUnexpectedEOF)
}

func parserStateEOF(p *parser) parserState {
	close(p.words)
	return nil
}

func parserStateSpace(p *parser) parserState {
	p.next()

	for p.cursor().tt == tokenWhitespace {
		// Keep skipping newlines.
		p.next()
	}

	return parserDefaultState
}

func parserStateNewLine(p *parser) parserState {
	p.next()

	for p.cursor().tt == tokenNewLine {
		// Keep skipping newlines.
		p.next()
	}

	return parserDefaultState
}

func parserStateContinue(p *parser) parserState {
	if !p.next() {
		return parserStateEOF
	}
	return parserDefaultState
}

func composedValue(tokens []*token) (*token, error) {
	if len(tokens) < 1 {
		return nil, errors.New("expecting token")
	}

	baseToken := tokens[0]

	values := []string{}
	for _, t := range tokens {
		values = append(values, t.val)
	}

	stringValue, err := unescapeString(strings.Join(values, ""))
	if err != nil {
		return nil, err
	}

	return &token{
		tt:   tokenComposed,
		val:  stringValue,
		line: baseToken.line,
		col:  baseToken.col,
	}, nil
}

func parseComments(comments map[int][]string, currentLine int) []*CommentNode {
	found := false
	iteration := 0
	commentNodes := []*CommentNode{}

	for ; currentLine >= 0; currentLine-- {
		val, ok := comments[currentLine]
		if ok {
			found = true
			commentNodes = append(commentNodes, &CommentNode{
				value: strings.Join(val, " "),
			})

			delete(comments, currentLine)
		}

		if !ok && iteration > 1 {
			break
		}

		if iteration > 1 && !found {
			break
		}

		iteration++
	}

	if len(commentNodes) > 0 {
		// slices.Reverse is introduced with go 1.22
		for i, j := 0, len(commentNodes)-1; i < j; i, j = i+1, j-1 {
			commentNodes[i], commentNodes[j] = commentNodes[j], commentNodes[i]
		}
	}

	return commentNodes
}
