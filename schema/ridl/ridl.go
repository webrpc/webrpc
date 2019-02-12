package ridl

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/webrpc/webrpc/schema"
)

type Tree struct {
}

func tokenize(input string) ([]token, error) {
	lx := newLexer(string(input))

	tokens := []token{}
	for {
		tok := <-lx.tokens
		if tok.tt == tokenSpace {
			continue
		}
		if tok.tt == tokenEOF {
			break
		}
		tokens = append(tokens, tok)
	}

	return tokens, nil
}

func Parse(input string) (*schema.WebRPCSchema, error) {
	p, err := newParser(input)
	if err != nil {
		return nil, err
	}
	if err = p.run(); err != nil {
		return nil, err
	}

	if p.tree.definitions["service"] == nil {
		return nil, errors.New(`missing "service" declaration`)
	}
	if p.tree.definitions["version"] == nil {
		return nil, errors.New(`missing "version" declaration`)
	}

	s := &schema.WebRPCSchema{
		Schema: "webrpc/v0.1.0",
		App:    fmt.Sprintf("%s/%s", p.tree.definitions["service"].value(), p.tree.definitions["version"].value()),
	}

	if len(p.tree.imports) > 0 {
		s.Imports = []string{}
		for _, tok := range p.tree.imports {
			s.Imports = append(s.Imports, tok.val)
		}
	}

	if len(p.tree.enums) > 0 {
		if s.Messages == nil {
			s.Messages = []*schema.Message{}
		}
		for _, enum := range p.tree.enums {
			fields := []*schema.MessageField{}

			for i, value := range enum.values {
				field := &schema.MessageField{
					Name: schema.VarName(value.left.val),
				}
				if value.right != nil {
					field.Value = value.right.val
				} else {
					field.Value = strconv.Itoa(i)
				}
				fields = append(fields, field)
			}

			s.Messages = append(s.Messages, &schema.Message{
				Name:   schema.VarName(enum.name.val),
				Type:   schema.MessageType("enum"),
				Fields: fields,
			})
		}
	}

	return s, nil
}
