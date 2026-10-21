package ridl

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/webrpc/webrpc/schema"
)

func parseStateServiceMethodDefinition(sn *ServiceNode) parserState {
	return func(p *parser) parserState {
		var streamInput, proxy bool

		defer func() {
			// clear annotation buffer
			sn.methodAnnotations = []*AnnotationNode{}
		}()

		// check for annotation duplicates
		annotations := make(map[string]struct{})
		for _, ann := range sn.methodAnnotations {
			if _, ok := annotations[ann.AnnotationType().String()]; ok {
				return p.stateError(fmt.Errorf("duplicate annotation type: %v", ann.AnnotationType()))
			}

			annotations[ann.AnnotationType().String()] = struct{}{}
		}
		// - <name>([arguments]) [=> [([ return values ])]]
		matches, err := p.match(tokenDash, tokenWhitespace, tokenWord)
		if err != nil {
			return p.stateError(err)
		}

		methodName := matches[2]

		if matches[2].val == "stream" {
			streamInput = true

			matches, err := p.match(tokenWhitespace, tokenWord)
			if err != nil {
				return p.stateError(err)
			}

			methodName = matches[1]
		}

		if matches[2].val == wordProxy {
			proxy = true

			matches, err := p.match(tokenWhitespace, tokenWord)
			if err != nil {
				return p.stateError(err)
			}

			methodName = matches[1]
		}

		commentLine := matches[0].line
		// we have to start parsing comments from the line of last annotation
		if len(sn.methodAnnotations) > 0 {
			commentLine = sn.methodAnnotations[len(sn.methodAnnotations)-1].AnnotationType().tok.line
		}

		// Emit deprecation warnings for above-method annotations now that we
		// know the method they belong to.
		for _, ann := range sn.methodAnnotations {
			name := ann.AnnotationType().String()
			if name == wordSuccess {
				// '@success' is not an annotation anymore, so the generic hint
				// would name a form that is itself a hard error.
				p.warn(fmt.Sprintf("%s:%d: warning: deprecated annotation syntax: '@success' is no longer supported: move it below the '- %s' method line as '+ success <codes>' (above-method annotations are deprecated and will be removed in a future release)", p.file, ann.AnnotationType().tok.line, methodName.val))
				continue
			}
			annStr := "+ @" + name
			if ann.value != nil {
				// Quote unless the bare to-EOL form would re-read the value exactly.
				v := ann.value.String()
				if !bareAnnotationValueReproducible(v) {
					v = strconv.Quote(v)
				}
				annStr += " = " + v
			}
			p.warn(fmt.Sprintf("%s:%d: warning: deprecated annotation syntax: move '%s' below the '- %s' method line (above-method annotations are deprecated and will be removed in a future release)", p.file, ann.AnnotationType().tok.line, annStr, methodName.val))
		}

		mn := &MethodNode{
			name:    newTokenNode(methodName),
			proxy:   proxy,
			comment: parseComments(p.comments, commentLine),
			inputs: argumentList{
				stream:    streamInput,
				arguments: []*ArgumentNode{},
			},
			outputs: argumentList{
				arguments: []*ArgumentNode{},
			},
			annotations: sn.methodAnnotations,
		}

		if proxy {
			sn.methods = append(sn.methods, mn)

			return parserStateServiceMethod(sn)
		}

		inputArguments, err := p.expectArgumentList()
		if err != nil {
			return p.stateError(err)
		}
		mn.inputs.arguments = inputArguments

		matches, err = p.match(tokenWhitespace, tokenRocket, tokenWhitespace)
		if err == nil {

			// is stream?
			matches, err := p.match(tokenWord, tokenWhitespace)
			if err == nil {
				if matches[0].val == wordStream {
					mn.outputs.stream = true
				} else {
					return p.stateError(errUnexpectedToken)
				}
			}

			// => ()
			outputArguments, err := p.expectArgumentList()
			if err != nil {
				return p.stateError(err)
			}

			mn.outputs.arguments = outputArguments
		}

		// Check for optional errors clause
		pos := p.pos
		matches, err = p.match(tokenWhitespace, tokenWord)
		if err == nil && matches[1].val == wordErrors {
			// Parse bar-separated list of error names
			errorNames, err := p.expectErrorList()
			if err != nil {
				return p.stateError(err)
			}
			if err := addMethodErrors(mn, errorNames); err != nil {
				return p.stateError(err)
			}
		} else {
			p.pos = pos // not an errors clause, e.g. a route line on the next line
		}

		sn.methods = append(sn.methods, mn)

		return parserStateServiceMethod(sn)
	}
}

func parserStateServiceMethod(s *ServiceNode) parserState {
	return func(p *parser) parserState {
		tok := p.cursor()

		switch tok.tt {
		case tokenNewLine, tokenWhitespace:
			p.next()

		case tokenAt:
			anns, err := parseAnnotations(p, false)
			if err != nil {
				return p.stateError(err)
			}
			s.methodAnnotations = append(s.methodAnnotations, anns...)

		case tokenHash:
			err := p.continueUntilEOL()
			if err != nil {
				return p.stateError(err)
			}

		case tokenDash:
			state := parseStateServiceMethodDefinition(s)
			return state

		case tokenPlusSign:
			// + POST /users/{userId}, a REST route for the preceding method
			if err := parseMethodRoute(p, s); err != nil {
				return p.stateError(err)
			}

		case tokenWord:
			switch {
			case tok.val == wordPath:
				return p.stateError(fmt.Errorf("service-level 'path' is no longer supported: declare absolute routes under methods, e.g. '+ GET /users/{userId}'"))

			case schema.IsRouteVerb(tok.val):
				return p.stateError(fmt.Errorf("route '%s' must start with '+': write '+ %s /path' under the method", tok.val, tok.val))

			default:
				// any other word ends the service block
				if err := pendingAnnotationsError(s); err != nil {
					return p.stateError(err)
				}
				p.emit(s)
				return parserDefaultState
			}

		default:
			if err := pendingAnnotationsError(s); err != nil {
				return p.stateError(err)
			}
			p.emit(s)
			return parserDefaultState

		}

		return parserStateServiceMethod(s)
	}
}

// pendingAnnotationsError reports annotations buffered after the last method;
// they would otherwise be silently dropped when the service block ends.
func pendingAnnotationsError(s *ServiceNode) error {
	if len(s.methodAnnotations) == 0 {
		return nil
	}
	ann := s.methodAnnotations[0]
	return fmt.Errorf("annotation '@%s' must be placed directly above a '- Method' line", ann.AnnotationType().String())
}

// parseMethodRoute parses a route meta line, e.g. `+ GET /users/{userId}`,
// and appends it to the method defined just above it.
func parseMethodRoute(p *parser, s *ServiceNode) error {
	if _, err := p.match(tokenPlusSign); err != nil {
		return err
	}
	if p.cursor().tt == tokenWhitespace {
		p.next()
	}

	if p.cursor().tt == tokenAt {
		// A '+ @...' line is a method-level meta line, so a pending method
		// annotation above it can no longer attach to the method below.
		if err := pendingAnnotationsError(s); err != nil {
			return err
		}
		if next := p.peek(); next.tt == tokenWord && next.val == wordSuccess {
			return fmt.Errorf("'+ @success' is no longer supported: write '+ success <codes>' without '@', e.g. '+ success 201'")
		}
		return parseMethodAnnotations(p, s)
	}

	tok := p.cursor()
	if tok.tt == tokenWord && (tok.val == wordSuccess || tok.val == wordError) {
		if err := pendingAnnotationsError(s); err != nil {
			return err
		}
		if tok.val == wordSuccess {
			return parseRouteSuccess(p, s)
		}
		return parseMethodErrors(p, s)
	}

	if len(s.methods) == 0 {
		return fmt.Errorf("route must be declared under a service method")
	}
	if len(s.methodAnnotations) > 0 {
		return fmt.Errorf("route must directly follow its method: declare annotations as '+ @name' under a method, not above it")
	}
	mn := s.methods[len(s.methods)-1]

	verbTok := p.cursor()
	if verbTok.tt == tokenNewLine || verbTok.tt == tokenEOF {
		return fmt.Errorf("expecting route verb after '+'")
	}
	if verbTok.tt != tokenWord || !schema.IsRouteVerb(verbTok.val) {
		return fmt.Errorf("expecting route '+ VERB /path' with VERB one of %s, got %q", strings.Join(schema.RouteVerbs, ", "), verbTok.val)
	}
	p.next()

	if _, err := p.match(tokenWhitespace); err != nil {
		return fmt.Errorf("expecting route path after '%s': %w", verbTok.val, err)
	}

	pathToken, err := p.expectRoutePath()
	if err != nil {
		return fmt.Errorf("expecting route path after '%s': %w", verbTok.val, err)
	}
	if err := schema.CheckRoutePathSyntax(pathToken.val); err != nil {
		return fmt.Errorf("invalid route of method '%s': %w", mn.name.String(), err)
	}
	if err := p.expectOptionalCommentOrEOL(); err != nil {
		return err
	}

	mn.routes = append(mn.routes, &RouteNode{verb: newTokenNode(verbTok), path: newTokenNode(pathToken)})
	return nil
}

// bareAnnotationValueReproducible reports whether the below-method `=` syntax
// would re-read v exactly from a bare value. A bare value ends at the first
// whitespace and a '#' starts a comment, so anything else needs quotes.
func bareAnnotationValueReproducible(v string) bool {
	if v == "" || strings.HasPrefix(v, `"`) || strings.Contains(v, "#") {
		return false
	}
	for _, r := range v {
		if unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

// parseMethodAnnotations parses a single `+ @name[ = value]` annotation (one
// annotation per line) and attaches it to the method defined just above it.
func parseMethodAnnotations(p *parser, s *ServiceNode) error {
	if len(s.methods) == 0 {
		return fmt.Errorf("annotation must be declared under a service method")
	}
	mn := s.methods[len(s.methods)-1]

	anns, err := parseAnnotations(p, true)
	if err != nil {
		return err
	}
	if len(anns) == 0 {
		return fmt.Errorf("expecting annotation name after '@', e.g. '+ @internal'")
	}
	if len(anns) > 1 {
		return fmt.Errorf("unexpected '@': one annotation per line")
	}
	ann := anns[0]
	for _, existing := range mn.annotations {
		if existing.AnnotationType().String() == ann.AnnotationType().String() {
			return fmt.Errorf("duplicate annotation type: %s", ann.AnnotationType().String())
		}
	}
	if err := p.expectOptionalCommentOrEOL(); err != nil {
		if tok := p.cursor(); tok.tt == tokenWord {
			return fmt.Errorf("expecting '=' before the annotation value, e.g. '+ @%s = %s'", ann.AnnotationType().String(), tok.val)
		}
		return fmt.Errorf("expecting comment or end of line after annotation, got %q", p.cursor().val)
	}
	mn.annotations = append(mn.annotations, ann)
	return nil
}

// parseRouteSuccess parses a `+ success <code>` option line and records the
// method's success codes; ridl.go applies them to every route of the method.
func parseRouteSuccess(p *parser, s *ServiceNode) error {
	nameTok := p.cursor()
	if nameTok.tt != tokenWord || nameTok.val != wordSuccess {
		return fmt.Errorf("expecting 'success <code>' after '+', e.g. '+ success 201'")
	}
	p.next()

	if len(s.methods) == 0 {
		return fmt.Errorf("'+ success' must be declared under a service method")
	}
	mn := s.methods[len(s.methods)-1]
	if mn.successCodes != nil {
		return fmt.Errorf("method '%s' already declares '+ success'", mn.name.String())
	}

	if _, err := p.match(tokenWhitespace); err != nil {
		return fmt.Errorf("expecting status code after '+ success': %w", err)
	}
	codeTok := p.cursor()
	if codeTok.tt != tokenWord {
		return fmt.Errorf("expecting '+ success' status code to be a number but got '%s'", codeTok.val)
	}
	// A pipe-separated set of allowed success codes, e.g. `+ success 200|201`;
	// the first is the default the handler can override at runtime.
	codes := []*TokenNode{}
	seen := map[int]bool{}
	for _, val := range strings.Split(codeTok.val, "|") {
		if expectNumber(codeTok, val) != nil {
			return fmt.Errorf("expecting '+ success' status code to be a number but got '%s'", codeTok.val)
		}
		code, err := strconv.Atoi(val)
		if err != nil {
			return fmt.Errorf("expecting '+ success' status code to be a number but got '%s'", codeTok.val)
		}
		if code < 200 || code > 299 {
			return fmt.Errorf("invalid '+ success' status code '%s' of method '%s': must be a 2xx code (200-299)", val, mn.name.String())
		}
		if seen[code] {
			return fmt.Errorf("duplicate '+ success' status code '%s' of method '%s'", val, mn.name.String())
		}
		seen[code] = true
		// A single output may still be an empty struct, which only the schema
		// validation (always run next) can resolve.
		if (code == 204 || code == 205) && len(mn.outputs.arguments) > 1 {
			return fmt.Errorf("invalid '+ success' status code '%s' of method '%s': 204/205 require an empty response (no outputs or an empty struct)", val, mn.name.String())
		}
		codes = append(codes, newTokenNode(&token{tt: codeTok.tt, val: val, pos: codeTok.pos, line: codeTok.line, col: codeTok.col}))
	}
	p.next()

	if err := p.expectOptionalCommentOrEOL(); err != nil {
		return err
	}

	mn.successCodes = codes
	return nil
}

// parseMethodErrors parses a `+ error ErrOne|ErrTwo` option line and appends the
// names to the method's error list; ridl.go maps them onto schema.Method.Errors.
func parseMethodErrors(p *parser, s *ServiceNode) error {
	nameTok := p.cursor()
	if nameTok.tt != tokenWord || nameTok.val != wordError {
		return fmt.Errorf("expecting 'error <Name>' after '+', e.g. '+ error NotFound'")
	}
	p.next()

	if len(s.methods) == 0 {
		return fmt.Errorf("'+ error' must be declared under a service method")
	}
	mn := s.methods[len(s.methods)-1]

	errorNames, err := p.expectErrorList()
	if err != nil {
		// The empty-segment message is already exact; wrapping it would hide
		// the name it was rejected from.
		var empty *emptyErrorNameError
		if errors.As(err, &empty) {
			return err
		}
		return fmt.Errorf("expecting error name after '+ error', e.g. '+ error NotFound': %w", err)
	}
	if err := p.expectOptionalCommentOrEOL(); err != nil {
		return err
	}

	return addMethodErrors(mn, errorNames)
}

// addMethodErrors appends names to the method, rejecting a name declared twice
// within the list, across '+ error' lines, or against the 'errors A|B' clause,
// so both forms behave identically.
func addMethodErrors(mn *MethodNode, names []*TokenNode) error {
	for _, name := range names {
		for _, declared := range mn.errors {
			if declared.String() == name.String() {
				return fmt.Errorf("method '%s' declares error '%s' twice", mn.name.String(), name.String())
			}
		}
		mn.errors = append(mn.errors, name)
	}
	return nil
}

func parserStateService(p *parser) parserState {
	matches, err := p.match(tokenWord, tokenWhitespace)
	if err != nil {
		return p.stateError(err)
	}

	if matches[0].val != "service" {
		return p.stateError(errUnexpectedToken)
	}

	serviceName, err := p.expectLiteralValue()
	if err != nil {
		return p.stateError(err)
	}
	if err := p.expectOptionalCommentOrEOL(); err != nil {
		return p.stateError(err)
	}

	return parserStateServiceMethod(&ServiceNode{
		name:    newTokenNode(serviceName),
		methods: []*MethodNode{},
		comment: parseComments(p.comments, matches[0].line),
	})
}

// parseAnnotations parses annotation groups. When toEOL is false (the
// deprecated above-method form) it reads `@name[:value]` pairs and loops over
// several of them on one line; when true it parses exactly one below-method
// `+ @name[ = value]` annotation, with struct-tag value rules.
func parseAnnotations(p *parser, toEOL bool) ([]*AnnotationNode, error) {
	annotations := []*AnnotationNode{}
	annotationsMap := map[string]struct{}{}

	// @acl:admin
	// @auth:"cookies,authorization,query" @internal
	matcher := []tokenType{tokenAt, tokenWord}

	for {
		annotationMatches, err := p.match(matcher...)
		if err != nil {
			if toEOL && len(annotations) == 0 {
				return nil, fmt.Errorf("expecting annotation name after '@', e.g. '+ @internal'")
			}
			break
		}

		annotation := &AnnotationNode{
			annotationType: newTokenNode(annotationMatches[1]),
		}

		if _, ok := annotationsMap[annotation.annotationType.String()]; ok {
			return nil, fmt.Errorf("duplicate annotation type: %s", annotation.annotationType.String())
		}
		// annotationsMap is reset per call; the caller checks duplicates across lines.

		annotationsMap[annotation.annotationType.String()] = struct{}{}

		if toEOL {
			for p.cursor().tt == tokenWhitespace {
				p.next()
			}
			// + @name = value — a quoted value keeps its spaces and escapes, a
			// bare value ends at the first whitespace or '#' and keeps
			// backslashes literal.
			if p.cursor().tt == tokenEqual {
				p.next()
				for p.cursor().tt == tokenWhitespace {
					p.next()
				}
				switch p.cursor().tt {
				case tokenHash, tokenNewLine, tokenEOF:
					return nil, fmt.Errorf("expecting value after '=' in annotation '@%s'", annotation.annotationType.String())
				}
				annotationValue, err := p.expectAnnotationValue()
				if err != nil {
					return nil, fmt.Errorf("expecting value after '=' in annotation '@%s': %w", annotation.annotationType.String(), err)
				}
				annotation.value = newTokenNode(annotationValue)
				for p.cursor().tt == tokenWhitespace {
					p.next()
				}
				if tok := p.cursor(); tok.tt != tokenAt && tok != eofToken && tok.tt != tokenNewLine && tok.tt != tokenHash {
					return nil, fmt.Errorf("unexpected %q after the value of annotation '@%s': quote values containing spaces, e.g. '+ @%s = \"two words\"'", tok.val, annotation.annotationType.String(), annotation.annotationType.String())
				}
			} else if p.cursor().tt == tokenColon {
				return nil, fmt.Errorf("expecting '=' before the annotation value, e.g. '+ @%s = value'", annotation.annotationType.String())
			}

			annotations = append(annotations, annotation)

			for p.cursor().tt == tokenWhitespace {
				p.next()
			}
			if p.cursor().tt == tokenAt {
				return nil, fmt.Errorf("unexpected '@': one annotation per line")
			}
			break
		}

		if p.cursor().tt != tokenColon {
			annotations = append(annotations, annotation)
		} else {
			p.next()
			switch tok := p.cursor(); tok.tt {
			// @acl:admin
			case tokenWord:
				annotation.value = newTokenNode(tok)
				p.next()

			// @auth:"cookies,authorization,query"
			case tokenQuote:
				// expectStringValue advances the cursor past the closing quote.
				annotationValue, err := p.expectStringValue()
				if err != nil {
					return nil, fmt.Errorf("parse string value: %w", err)
				}

				annotation.value = newTokenNode(annotationValue)

			default:
				return nil, fmt.Errorf("expecting value after ':' in annotation '@%s'", annotation.annotationType.String())
			}

			annotations = append(annotations, annotation)
		}
	}

	return annotations, nil
}
