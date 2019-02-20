package ridl

import (
	"errors"
	"fmt"
	"strings"
)

var (
	errUnexpectedEOF = errors.New(`unexpected EOF`)
)

type methodArgument struct {
	left     *token
	right    *token
	stream   bool
	optional bool
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
			return nil
		}
		if !p.next() {
			return nil
		}
	}
	return nil
}

func (p *parser) expectTypeExpr(baseToken *token) (*token, error) {
	invalidErr := errors.New("invalid data type")

	if baseToken == nil {
		baseToken = &token{}
	}

	tokens := []*token{}
	mapopen := []int{}
	mapclose := []int{}

	inmap := false

	for {
		if !p.next() {
			break
		}

		// start of a map
		if p.cursor().tt == tokenWord && p.cursor().val == "map" {

			tokens = append(tokens, p.cursor())
			if err := p.expectNext(tokenOpenAngleBracket); err != nil {
				return nil, invalidErr
			}
			tokens = append(tokens, p.cursor())

			mapopen = append(mapopen, len(tokens))
			inmap = true

			continue
		}

		// primitive
		if p.cursor().tt == tokenWord {
			tokens = append(tokens, p.cursor())
			if inmap {
				if err := p.expectNext(tokenComma); err != nil {
					return nil, invalidErr
				}
				tokens = append(tokens, p.cursor())
				inmap = false
			}
			continue
		}

		// list
		if p.cursor().tt == tokenOpenBracket {
			tokens = append(tokens, p.cursor())
			if err := p.expectNext(tokenCloseBracket); err != nil {
				return nil, invalidErr
			}
			tokens = append(tokens, p.cursor())
			if inmap {
				if err := p.expectNext(tokenComma); err != nil {
					return nil, invalidErr
				}
				tokens = append(tokens, p.cursor())
				inmap = false
			}
			continue
		}

		// closing a map
		if p.cursor().tt == tokenCloseAngleBracket {
			if len(mapclose) >= len(mapopen) {
				return nil, invalidErr
			}
			mapclose = append(mapclose, len(tokens))
			tokens = append(tokens, p.cursor())
			continue
		}

		// end expr
		if len(mapopen) > 0 {
			if len(mapopen) == len(mapclose) {
				break
			}
		} else if p.cursor().tt == tokenComma || p.cursor().tt == tokenCloseParen {
			break
		}

		return nil, invalidErr
	}

	values := []string{}
	for _, t := range tokens {
		values = append(values, t.val)
	}

	return &token{
		tt:   tokenComposed,
		val:  strings.Join(values, ""),
		line: baseToken.line,
		col:  baseToken.col,
	}, nil
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

	var err error
	var methodDefinition *method
	var composedToken *token

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
			var methodArg *methodArgument

			// INPUTS
			if err := p.expectNext(tokenOpenParen); err != nil {
				return p.stateError(err)
			}
			if !p.next() {
				return parseStateUnexpectedEOF
			}
			if p.cursor().tt == tokenCloseParen {
				goto outputs
			}

		inputargs:
			if p.cursor().tt != tokenWord {
				return p.stateError(errors.New("expecting word"))
			}
			methodArg = &methodArgument{}
			methodArg.left = p.cursor()

			if !p.next() {
				return parseStateUnexpectedEOF
			}
			if p.cursor().tt == tokenQuestionMark {
				methodArg.optional = true
				if !p.next() {
					return parseStateUnexpectedEOF
				}
			}

			if p.cursor().tt != tokenColon {
				return p.stateError(errors.New("expecting colon"))
			}

			// type definition
			if !p.next() {
				return parseStateUnexpectedEOF
			}
			if p.cursor().val == "stream" {
				methodArg.stream = true
			} else {
				p.back()
			}

			composedToken, err = p.expectTypeExpr(p.cursor())
			if err != nil {
				return p.stateError(err)
			}

			methodArg.right = composedToken
			methodDefinition.inputs = append(methodDefinition.inputs, methodArg)

			if p.cursor().tt == tokenComma {
				// next input arg
				if !p.next() {
					return parseStateUnexpectedEOF
				}
				goto inputargs
			}
			if p.cursor().tt != tokenCloseParen {
				return p.stateError(errors.New("expecting close paren"))
			}

			// OUTPUTS
		outputs:
			// empty outputs?
			if !p.next() {
				return parseStateUnexpectedEOF
			}

			if p.cursor().tt == tokenNewLine || p.cursor().tt == tokenHash {
				goto methoddef // we done with this line
			}

			// output args provided
			if p.cursor().tt != tokenRocket {
				return p.stateError(errors.New("expecting rocket"))
			}
			if err := p.expectNext(tokenOpenParen); err != nil {
				return p.stateError(err)
			}

			if !p.next() {
				return parseStateUnexpectedEOF
			}

		outputargs:
			if p.cursor().tt != tokenWord {
				return p.stateError(errors.New("expecting word"))
			}
			methodArg = &methodArgument{}
			methodArg.left = p.cursor()

			if !p.next() {
				return parseStateUnexpectedEOF
			}
			if p.cursor().tt == tokenQuestionMark {
				methodArg.optional = true
				if !p.next() {
					return parseStateUnexpectedEOF
				}
			}

			if p.cursor().tt != tokenColon {
				return p.stateError(errors.New("expecting colon"))
			}

			// type definition
			if !p.next() {
				return parseStateUnexpectedEOF
			}
			if p.cursor().val == "stream" {
				methodArg.stream = true
			} else {
				p.back()
			}

			composedToken, err = p.expectTypeExpr(p.cursor())
			if err != nil {
				return p.stateError(err)
			}

			methodArg.right = composedToken

			methodDefinition.outputs = append(methodDefinition.outputs, methodArg)

			if p.cursor().tt == tokenComma {
				// next output arg
				if !p.next() {
					return parseStateUnexpectedEOF
				}
				goto outputargs
			}
			if p.cursor().tt != tokenCloseParen {
				return p.stateError(errors.New("expecting close paren"))
			}

		methoddef:
			methods = append(methods, methodDefinition)
			if !p.next() {
				goto loop
			}
		default:
			break loop
		}
	}

	p.tree.services = append(p.tree.services, &service{
		name:    serviceName,
		methods: methods,
	})

	return parseStateStart
}

func parseStateMessage(p *parser) parseState {
	if err := p.expectNext(tokenWord); err != nil {
		return p.stateError(err)
	}
	messageName := p.cursor()

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

			if !p.next() {
				return parseStateUnexpectedEOF
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
			if err := p.expectNext(tokenWord); err != nil {
				return p.stateError(err)
			}
			fieldDefinition = &definition{left: p.cursor()}

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

	return parseStateStart
}

func parseStateEnum(p *parser) parseState {
	// Enum definition
	if err := p.expectNext(tokenWord); err != nil {
		return p.stateError(err)
	}
	enumName := p.cursor()

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
				return p.stateError(err)
			}
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
			if !p.next() {
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

	switch word.val {
	case "webrpc", "name", "version":
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
	switch tok.tt {
	case tokenNewLine:
		return parseStateNewLine
	case tokenWord:
		return parseStateLine
	case tokenHash:
		return parseStateComment
	case tokenEOF: // abrupt EOF
		return nil
	default:
	}
	return nil
}

func parseStateUnexpectedEOF(p *parser) parseState {
	return p.stateError(errUnexpectedEOF)
}

func parseStateEOF(p *parser) parseState {
	return nil
}
