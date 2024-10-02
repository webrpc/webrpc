package ridl

import "fmt"

func parseStateServiceMethodDefinition(sn *ServiceNode) parserState {
	return func(p *parser) parserState {
		var streamInput, proxy bool

		annotations, err := parseAnnotations(p)
		if err != nil {
			return p.stateError(err)
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
		if len(annotations) > 0 {
			commentLine = annotations[len(annotations)-1].AnnotationType().tok.line
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
			annotations: annotations,
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
			p.continueUntilEOL()

		case tokenHash:
			p.continueUntilEOL()

		case tokenDash:
			return parseStateServiceMethodDefinition(s)

		default:
			p.emit(s)
			return parserDefaultState

		}

		return parserStateServiceMethod(s)
	}
}

func parserStateService(p *parser) parserState {
	matches, err := p.match(tokenWord, tokenWhitespace, tokenWord, tokenEOL)
	if err != nil {
		return p.stateError(err)
	}

	if matches[0].val != "service" {
		return p.stateError(errUnexpectedToken)
	}

	return parserStateServiceMethod(&ServiceNode{
		name:    newTokenNode(matches[2]),
		methods: []*MethodNode{},
		comment: parseComments(p.comments, matches[0].line),
	})
}

func parseAnnotations(p *parser) ([]*AnnotationNode, error) {
	annotations := []*AnnotationNode{}
	annotationsMap := map[string]struct{}{}

	currentPosition := p.pos

	// @acl:admin
	// @auth:"cookies,authorization,query" @internal
	matcher := []tokenType{tokenWhitespace, tokenAt, tokenWord}
	newline := 0

outerloop:
	for {
		prev := p.prev()
		if !prev {
			break
		}

		if p.cursor().tt == tokenNewLine {
			newline++
		}

		if p.cursor().tt == tokenNewLine && newline > 1 {
			lineStartPosition := p.pos
			p.next()

			matches := 0
			for {
				annotationMatches, err := p.match(matcher...)
				// when we don't match annotation format
				if err != nil {
					if matches == 0 {
						break outerloop
					}

					break
				}

				matches++

				if _, ok := annotationsMap[annotationMatches[2].String()]; ok {
					return nil, fmt.Errorf("duplicate annotation detected: %s", annotationMatches[2].String())
				}

				annotation := &AnnotationNode{
					annotationType: newTokenNode(annotationMatches[2]),
				}

				annotationsMap[annotationMatches[2].String()] = struct{}{}

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

			err := p.goTo(lineStartPosition)
			if err != nil {
				return annotations, fmt.Errorf("goto: %w", err)
			}
		}
	}

	// rewind to position before we started parsing annotations
	err := p.goTo(currentPosition)
	if err != nil {
		return annotations, fmt.Errorf("goto: %w", err)
	}

	return annotations, nil
}
