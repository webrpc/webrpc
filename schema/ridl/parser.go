package ridl

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

var (
	errUnexpectedEOF = errors.New(`unexpected EOF`)
)

type methodArgument struct {
	left   *token
	right  *token
	stream bool
}

type method struct {
	name    *token
	inputs  []*methodArgument
	outputs []*methodArgument
}

type service struct {
	name    *token
	methods []*method
}

type definition struct {
	left     *token
	right    *token
	optional bool

	meta []*definition
}

type enum struct {
	name     *token
	enumType *token
	values   []*definition
}

type message struct {
	name   *token
	fields []*definition
}

func (d *definition) value() string {
	return d.right.val
}

type tree struct {
	definitions map[string]*definition

	imports  []*token
	enums    []*enum
	messages []*message
	services []*service
}

var (
	eofToken = &token{tt: tokenEOF}
)

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
			enums:       []*enum{},
		},
	}
	return p, nil
}

func (p *parser) run() error {
	for state := parseStateStart; state != nil; {
		state = state(p)
	}
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
		log.Printf("discard?: %q", tok.val)

		switch tok.tt {
		case tokenNewLine:
			// got new line, continue
		case tokenExtra:
			// unknown char
		case tokenEOF:
			// got EOF
			return nil
		case tokenHash:
			if !p.next() {
				return nil
			}
			return p.continueUntilEOL()
		default:
			// another kind of token, stop
			log.Printf("discard? -> nope")
			return nil
		}
		if !p.next() {
			return nil
		}
	}
	return nil
}

func (p *parser) expectDelimiter(baseToken *token, delimiters ...tokenType) (*token, error) {
	if baseToken == nil {
		baseToken = &token{}
	}

	values := []string{baseToken.val}

loop:
	for {
		if !p.next() {
			break loop
		}

		tok := p.cursor()
		for i := range delimiters {
			if tok.tt == delimiters[i] {
				if err := p.back(); err != nil {
					return nil, err
				}
				break loop
			}
		}

		values = append(values, tok.val)
	}

	return &token{
		tt:   tokenComposed,
		val:  strings.Join(values, ""),
		line: baseToken.line,
		col:  baseToken.col,
	}, nil
}

func (p *parser) expectComposedToken(baseToken *token, tt ...tokenType) (*token, error) {
	if baseToken == nil {
		baseToken = &token{}
	}

	values := make([]string, 0, len(tt)+1)
	values = append(values, baseToken.val)

	for i := 0; ; {
		if !p.next() {
			return nil, errors.New(`unexpected EOF`)
		}

		tok := p.cursor()
		if tok.tt != tt[i] {
			return nil, fmt.Errorf("unexpected token %v, expecting %v", tok, tt[i])
		}
		values = append(values, tok.val)
		i = i + 1

		if i >= len(tt) {
			return &token{
				tt:   tokenComposed,
				val:  strings.Join(values, ""),
				line: baseToken.line,
				col:  baseToken.col,
			}, nil
		}
	}
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

func (p *parser) back() error {
	if p.pos <= 0 {
		return errors.New("can't go back")
	}
	p.pos = p.pos - 1
	return nil
}

func (p *parser) next() bool {
	if p.pos >= p.length {
		log.Printf("EOF")
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

func parseStateService(p *parser) parseState {
	if err := p.expectNext(tokenWord); err != nil {
		return p.stateError(err)
	}
	serviceName := p.cursor()

	methods := []*method{}

	if !p.next() {
		return parseStateUnexpectedEOF
	}

	var methodDefinition *method

loop:
	for {
		if err := p.expectCommentOrNewLine(); err != nil {
			return p.stateError(err)
		}

		tok := p.cursor()
		switch tok.tt {
		case tokenNewLine:
			if !p.next() {
				return nil
			}
		case tokenMinusSign:
			if err := p.expectNext(tokenWord); err != nil {
				return p.stateError(err)
			}
			methodDefinition = &method{
				name:    p.cursor(),
				inputs:  []*methodArgument{},
				outputs: []*methodArgument{},
			}
			log.Printf("methodDefinition: %v", methodDefinition)

			if err := p.expectNext(tokenOpenParen); err != nil {
				return p.stateError(err)
			}

			if err := p.expectNext(tokenWord); err == nil {
				methodArgument := &methodArgument{}
				if p.cursor().val == "stream" {
					methodArgument.stream = true
					if err := p.expectNext(tokenWord); err != nil {
						return p.stateError(err)
					}
				}
				methodArgument.right = p.cursor()
				// check for close parens
				methodDefinition.inputs = append(methodDefinition.inputs, methodArgument)
				if !p.next() {
					return parseStateUnexpectedEOF
				}
			}
			if p.cursor().tt != tokenCloseParen {
				return p.stateError(errors.New("expecting close paren"))
			}

			if err := p.expectNext(tokenColon); err == nil {
				if err := p.expectNext(tokenWord); err != nil {
					//return p.stateError(err)
				}
				methodArgument := &methodArgument{}
				if p.cursor().val == "stream" {
					methodArgument.stream = true
					if err := p.expectNext(tokenWord); err != nil {
						return p.stateError(err)
					}
				}
				methodArgument.right = p.cursor()
				methodDefinition.outputs = append(methodDefinition.outputs, methodArgument)
			}

			methods = append(methods, methodDefinition)
			if !p.next() {
				break loop
			}
		default:
			break loop
		}
	}

	p.tree.services = append(p.tree.services, &service{
		name:    serviceName,
		methods: methods,
	})

	log.Printf("service name %v", serviceName)

	return parseStateStart
}

func parseStateMessage(p *parser) parseState {
	if err := p.expectNext(tokenWord); err != nil {
		return p.stateError(err)
	}
	messageName := p.cursor()

	log.Printf("messageName: %v", messageName)

	fields := []*definition{}

	if !p.next() {
		return parseStateUnexpectedEOF
	}

	var fieldDefinition *definition

loop:
	for {
		if err := p.expectCommentOrNewLine(); err != nil {
			return p.stateError(err)
		}

		tok := p.cursor()
		switch tok.tt {
		case tokenNewLine:
			if !p.next() {
				return nil
			}
		case tokenPlusSign:
			if fieldDefinition == nil {
				return p.stateError(errors.New("unexpected meta property"))
			}

			if err := p.expectNext(tokenWord); err != nil {
				return p.stateError(err)
			}
			metaName := p.cursor()

			if err := p.expectNext(tokenEqual); err != nil {
				return p.stateError(err)
			}

			if err := p.expectNext(tokenWord); err != nil {
				return p.stateError(err)
			}

			composedToken, err := p.expectDelimiter(p.cursor(), tokenSpace, tokenHash, tokenNewLine)
			if err != nil {
				return p.stateError(err)
			}
			//var metaValue *token
			metaValue := composedToken

			fieldDefinition.meta = append(fieldDefinition.meta, &definition{
				left:  metaName,
				right: metaValue,
			})

			if !p.next() {
				break loop
			}
		case tokenMinusSign:
			log.Printf("got minus sign")
			if err := p.expectNext(tokenWord); err != nil {
				return p.stateError(err)
			}
			fieldDefinition = &definition{left: p.cursor()}
			log.Printf("fieldDefinition: %v", fieldDefinition)

			if err := p.expectNext(tokenColon); err != nil {
				if p.cursor().tt != tokenQuestionMark {
					return p.stateError(err)
				}
				fieldDefinition.optional = true
				if err := p.expectNext(tokenColon); err != nil {
					return p.stateError(err)
				}
			}

			// TODO: new state machine
			if err := p.expectNext(tokenWord); err != nil {
				// if not a word, could be []
				if p.cursor().tt != tokenOpenBracket {
					return p.stateError(err)
				}
				if err := p.expectNext(tokenCloseBracket); err != nil {
					return p.stateError(err)
				}
				if err := p.expectNext(tokenWord); err != nil {
					return p.stateError(err)
				}
				// TODO: compose token properly
				fieldDefinition.right = &token{val: "[]" + p.cursor().val}
			} else {
				if p.cursor().val == "map" {
					// TODO: check grammar
					composedToken, err := p.expectDelimiter(
						p.cursor(),
						tokenSpace,
						tokenHash,
						tokenNewLine,
					)
					if err != nil {
						return p.stateError(err)
					}
					fieldDefinition.right = composedToken
				} else {
					fieldDefinition.right = p.cursor()
				}
			}

			fields = append(fields, fieldDefinition)
			if !p.next() {
				break loop
			}
		default:
			break loop
		}
	}

	p.tree.messages = append(p.tree.messages, &message{
		name:   messageName,
		fields: fields,
	})

	log.Printf("message name %v", messageName)

	return parseStateStart
}

func parseStateEnum(p *parser) parseState {
	// Enum definition
	if err := p.expectNext(tokenWord); err != nil {
		return p.stateError(err)
	}
	enumName := p.cursor()

	log.Printf("ENUM cursor: %v", enumName)

	if err := p.expectNext(tokenColon); err != nil {
		return p.stateError(err)
	}

	if err := p.expectNext(tokenWord); err != nil {
		return p.stateError(err)
	}
	enumType := p.cursor()

	values := []*definition{}

	if !p.next() {
		return parseStateUnexpectedEOF
	}

	// Parse enum values
loop:
	for {

		if err := p.expectCommentOrNewLine(); err != nil {
			return p.stateError(err)
		}

		tok := p.cursor()
		switch tok.tt {
		case tokenNewLine:
			if !p.next() {
				break loop
			}
		case tokenMinusSign:
			log.Printf("got minus")
			if err := p.expectNext(tokenWord); err != nil {
				return p.stateError(err)
			}
			enumDefinition := &definition{left: p.cursor()}

			if err := p.expectNext(tokenEqual); err == nil {
				// has "=" sign
				if err := p.expectNext(tokenWord); err != nil {
					return p.stateError(err)
				}
				enumDefinition.right = p.cursor()
			}
			if err := p.continueUntilEOL(); err != nil {
				log.Printf("UNTIL EOL")
				return p.stateError(err)
			}
			log.Printf("enumDefinition: %#v", enumDefinition.left)
			values = append(values, enumDefinition)
			if !p.next() {
				break loop
			}
		default:
			break loop
		}
	}

	p.tree.enums = append(p.tree.enums, &enum{
		name:     enumName,
		enumType: enumType,
		values:   values,
	})

	log.Printf("enumName: %v, enumType: %v, values: %v", enumName, enumType, values)

	return parseStateStart
}

func parseStateImport(p *parser) parseState {
	if !p.next() {
		return parseStateUnexpectedEOF
	}

loop:
	for {
		if err := p.expectCommentOrNewLine(); err != nil {
			return p.stateError(err)
		}

		tok := p.cursor()
		switch tok.tt {
		case tokenNewLine:
			if !p.next() {
				return nil
			}
		case tokenMinusSign:
			if err := p.expectNext(tokenWord); err != nil {
				return p.stateError(err)
			}

			p.tree.imports = append(p.tree.imports, p.cursor())
			log.Printf("pushed import")
			if !p.next() {
				log.Printf("NEXT")
				break loop
			}
		default:
			break loop
		}
	}

	return parseStateStart
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
	case "ridl", "name", "version":
		return parseStateDefinition(word)
	case "import":
		return parseStateImport
	case "enum":
		return parseStateEnum
	case "message":
		return parseStateMessage
	case "service":
		return parseStateService
	default:
		return p.stateError(fmt.Errorf("unknown definition near: %v", word))
	}

	return parseStateStart
}

func parseStateStart(p *parser) parseState {
	tok := p.cursor()
	log.Printf("tok: %v", tok)
	switch tok.tt {
	case tokenNewLine:
		return parseStateNewLine
	case tokenWord:
		return parseStateLine
	case tokenHash:
		return parseStateComment
	case tokenEOF: // abrupt EOF
		log.Printf("got EOF")
		return nil
	default:
		log.Printf("GOT STATE: %v", tok)
	}
	return nil
}

func parseStateUnexpectedEOF(p *parser) parseState {
	return p.stateError(errUnexpectedEOF)
}

func parseStateEOF(p *parser) parseState {
	return nil
}
