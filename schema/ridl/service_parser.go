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
	currentPosition := p.pos

	annotationMatchers := [][]tokenType{
		{tokenWhitespace, tokenAt, tokenWord, tokenColon, tokenWord},
		{tokenWhitespace, tokenAt, tokenWord},
	}

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
			newline++
			lineStartPosition := p.pos
			p.next()

			matches := 0

		loop:
			for {
				for i, matcher := range annotationMatchers {
					annotationMatches, err := p.match(matcher...)
					// when we don't match annotation format
					if err != nil {
						if i == len(annotationMatchers)-1 {
							// if parsing of the first combination on new line does not succeed
							// we don't continue and stop annotation parsing
							if matches == 0 {
								break outerloop
							}

							break loop
						}

						continue
					}

					annotationNode := &AnnotationNode{
						annotationType: newTokenNode(annotationMatches[2]),
					}
					if len(annotationMatches) == 5 {
						annotationNode.value = newTokenNode(annotationMatches[4])
					}

					annotations = append(annotations, annotationNode)

					matches++
					break
				}
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
