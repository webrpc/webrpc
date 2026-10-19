package ridl

import (
	"fmt"
	"strconv"
	"strings"

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
			annStr := "+ @" + name
			if ann.value != nil {
				annStr += ":" + quoteIfNeeded(ann.value.String())
			}
			if name == "success" {
				// '+ @success' below a method is route-scoped; the annotation must be renamed there.
				p.warn(fmt.Sprintf("%s:%d: warning: deprecated annotation syntax: move '%s' below the '- %s' method line and rename it (e.g. '+ @successful:ok'): '+ @success' is reserved for route success codes", p.file, ann.AnnotationType().tok.line, annStr, methodName.val))
				continue
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
			mn.errors = errorNames
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
			anns, err := parseAnnotations(p)
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
		if next := p.peek(); next.tt == tokenWord && next.val == "success" {
			return parseRouteSuccess(p, s)
		}
		return parseMethodAnnotations(p, s)
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
	s.lastMetaRoute = true
	return nil
}

// quoteIfNeeded renders an annotation value as valid RIDL in deprecation
// warnings: bare only when the lexer reads it as a single word token (a word
// beginning followed by no word-break runes), quoted otherwise.
func quoteIfNeeded(val string) string {
	runes := []rune(val)
	bare := len(runes) > 0 && isWord(runes[0])
	for i := 1; bare && i < len(runes); i++ {
		bare = !isWordBreak(runes[i])
	}
	if !bare {
		return strconv.Quote(val)
	}
	return val
}

// parseMethodAnnotations parses `+ @name[:value]` annotations (one or more
// groups per line) and attaches them to the method defined just above it.
func parseMethodAnnotations(p *parser, s *ServiceNode) error {
	s.lastMetaRoute = false
	if len(s.methods) == 0 {
		return fmt.Errorf("annotation must be declared under a service method")
	}
	mn := s.methods[len(s.methods)-1]

	anns := []*AnnotationNode{}
	seen := map[string]struct{}{}
	for _, existing := range mn.annotations {
		seen[existing.AnnotationType().String()] = struct{}{}
	}
	lastName := ""
	for p.cursor().tt == tokenAt {
		more, err := parseAnnotations(p)
		if err != nil {
			return err
		}
		if len(more) == 0 {
			return fmt.Errorf("expecting annotation name after '@', e.g. '+ @internal'")
		}
		for _, ann := range more {
			key := ann.AnnotationType().String()
			lastName = key
			if _, ok := seen[key]; ok {
				return fmt.Errorf("duplicate annotation type: %s", key)
			}
			seen[key] = struct{}{}
		}
		anns = append(anns, more...)
		for p.cursor().tt == tokenWhitespace {
			p.next()
		}
	}
	if err := p.expectOptionalCommentOrEOL(); err != nil {
		if tok := p.cursor(); tok.tt == tokenWord && lastName != "" {
			return fmt.Errorf("expecting ':' before the annotation value, e.g. '+ @%s:%s'", lastName, tok.val)
		}
		return fmt.Errorf("expecting comment or end of line after annotation, got %q", p.cursor().val)
	}
	mn.annotations = append(mn.annotations, anns...)
	return nil
}

// parseRouteSuccess parses a `+ @success <code>` meta line and attaches the
// status code to the route line directly above it.
func parseRouteSuccess(p *parser, s *ServiceNode) error {
	if _, err := p.match(tokenAt); err != nil {
		return err
	}
	nameTok := p.cursor()
	if nameTok.tt != tokenWord || nameTok.val != "success" {
		return fmt.Errorf("expecting '@success <code>' after '+', e.g. '+ @success 201'")
	}
	p.next()

	if len(s.methods) == 0 {
		return fmt.Errorf("@success must be declared under a service method")
	}
	mn := s.methods[len(s.methods)-1]
	if len(mn.routes) == 0 || !s.lastMetaRoute {
		return fmt.Errorf("@success must directly follow a route line of method '%s', e.g. '+ POST /webhooks' then '+ @success 201'", mn.name.String())
	}
	route := mn.routes[len(mn.routes)-1]
	if len(route.success) > 0 {
		return fmt.Errorf("route '%s %s' of method '%s' already declares @success", route.verb.String(), route.path.String(), mn.name.String())
	}

	if _, err := p.match(tokenWhitespace); err != nil {
		return fmt.Errorf("expecting status code after '@success': %w", err)
	}
	codeTok := p.cursor()
	if codeTok.tt != tokenWord {
		return fmt.Errorf("expecting '@success' status code to be a number but got '%s'", codeTok.val)
	}
	// A pipe-separated set of allowed success codes, e.g. `+ @success 200|201`;
	// the first is the default the handler can override at runtime.
	codes := []*TokenNode{}
	seen := map[int]bool{}
	for _, val := range strings.Split(codeTok.val, "|") {
		if expectNumber(codeTok, val) != nil {
			return fmt.Errorf("expecting '@success' status code to be a number but got '%s'", codeTok.val)
		}
		code, err := strconv.Atoi(val)
		if err != nil {
			return fmt.Errorf("expecting '@success' status code to be a number but got '%s'", codeTok.val)
		}
		if code < 200 || code > 299 {
			return fmt.Errorf("invalid @success status code '%s' of method '%s': must be a 2xx code (200-299)", val, mn.name.String())
		}
		if seen[code] {
			return fmt.Errorf("duplicate @success status code '%s' of method '%s'", val, mn.name.String())
		}
		seen[code] = true
		// A single output may still be an empty struct, which only the schema
		// validation (always run next) can resolve.
		if (code == 204 || code == 205) && len(mn.outputs.arguments) > 1 {
			return fmt.Errorf("invalid @success status code '%s' of method '%s': 204/205 require an empty response (no outputs or an empty struct)", val, mn.name.String())
		}
		codes = append(codes, newTokenNode(&token{tt: codeTok.tt, val: val, pos: codeTok.pos, line: codeTok.line, col: codeTok.col}))
	}
	p.next()

	if err := p.expectOptionalCommentOrEOL(); err != nil {
		return err
	}

	route.success = codes
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

func parseAnnotations(p *parser) ([]*AnnotationNode, error) {
	annotations := []*AnnotationNode{}
	annotationsMap := map[string]struct{}{}

	// @acl:admin
	// @auth:"cookies,authorization,query" @internal
	matcher := []tokenType{tokenAt, tokenWord}

	for {
		annotationMatches, err := p.match(matcher...)
		if err != nil {
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

		if p.cursor().tt != tokenColon {
			annotations = append(annotations, annotation)
			continue
		}

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

	return annotations, nil
}
