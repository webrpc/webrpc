package ridl

import (
	"fmt"
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
				p.emit(s)
				return parserDefaultState
			}

		default:
			p.emit(s)
			return parserDefaultState

		}

		return parserStateServiceMethod(s)
	}
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

	if len(s.methods) == 0 {
		return fmt.Errorf("route must be declared under a service method")
	}
	if len(s.methodAnnotations) > 0 {
		return fmt.Errorf("route must directly follow its method: move the annotations above the next method")
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

		annotationsMap[annotation.annotationType.String()] = struct{}{}

		if p.cursor().tt != tokenColon {
			annotations = append(annotations, annotation)
			continue
		}

		p.next()

		// @acl:admin
		if p.cursor() != eofToken && p.cursor().tt == tokenWord {
			annotation.value = newTokenNode(p.cursor())
		}

		// @auth:"cookies,authorization,query"
		if p.cursor() != eofToken && p.cursor().tt == tokenQuote {
			annotationValue, err := p.expectStringValue()
			if err != nil {
				return nil, fmt.Errorf("parse string value: %w", err)
			}

			annotation.value = newTokenNode(annotationValue)
		}

		p.next()

		annotations = append(annotations, annotation)
	}

	return annotations, nil
}
