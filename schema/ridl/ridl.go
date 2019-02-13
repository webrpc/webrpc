package ridl

import (
	"errors"
	"fmt"
	"log"
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

	if p.tree.definitions["name"] == nil {
		return nil, errors.New(`missing "name" declaration`)
	}
	if p.tree.definitions["version"] == nil {
		return nil, errors.New(`missing "version" declaration`)
	}

	s := &schema.WebRPCSchema{
		Schema:  "webrpc/v0.1.0", // TODO: lets add schema.VERSION constant to track, or top-level webrpc.VERSION
		Name:    p.tree.definitions["name"].value(),
		Version: p.tree.definitions["version"].value(),
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
			varType, err := schema.NewVarTypeFromString(enum.enumType.val)
			if err != nil {
				return nil, fmt.Errorf("unknown data type: %v", enum.enumType)
			}

			for i := range enum.values {
				value := enum.values[i]
				log.Printf("value: %#v -> %#v", value.left, value.right)
				field := &schema.MessageField{
					Name: schema.VarName(value.left.val),
					Type: varType,
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

	if len(p.tree.messages) > 0 {
		if s.Messages == nil {
			s.Messages = []*schema.Message{}
		}
		for _, message := range p.tree.messages {
			fields := []*schema.MessageField{}

			for i := range message.fields {
				value := message.fields[i]
				varType := &schema.VarType{}
				err := schema.ParseVarTypeExpr(s, value.right.val, varType)
				if err != nil {
					return nil, fmt.Errorf("unknown data type: %v", value.right.val)
				}
				field := &schema.MessageField{
					Name: schema.VarName(value.left.val),
					Type: varType,
				}
				for _, meta := range value.meta {
					log.Printf("meta: %v", meta)
					field.Meta = append(field.Meta, schema.MessageFieldMeta{
						meta.left.val: meta.right.val,
					})
				}
				fields = append(fields, field)
			}

			s.Messages = append(s.Messages, &schema.Message{
				Name:   schema.VarName(message.name.val),
				Type:   schema.MessageType("message"),
				Fields: fields,
			})
		}
	}

	if len(p.tree.services) > 0 {
		if s.Services == nil {
			s.Services = []*schema.Service{}
		}
		for _, service := range p.tree.services {
			methods := []*schema.Method{}

			for i := range service.methods {
				value := service.methods[i]

				method := &schema.Method{
					Name:    schema.VarName(value.name.val),
					Inputs:  []*schema.MethodArgument{},
					Outputs: []*schema.MethodArgument{},
				}
				for _, arg := range value.inputs {
					varType, err := schema.NewVarTypeFromString(arg.right.val)
					if err != nil {
						return nil, fmt.Errorf("unknown data type: %v", arg.right.val)
					}
					method.Inputs = append(method.Inputs, &schema.MethodArgument{
						Name:   schema.VarName(arg.left.val),
						Type:   varType,
						Stream: arg.stream,
					})
				}
				/*
					for _, meta := range value.meta {
						log.Printf("meta: %v", meta)
						field.Meta = append(field.Meta, schema.MessageFieldMeta{
							meta.left.val: meta.right.val,
						})
					}
				*/
				methods = append(methods, method)
			}

			s.Services = append(s.Services, &schema.Service{
				Name:    schema.VarName(service.name.val),
				Methods: methods,
			})
		}
	}

	return s, nil
}
