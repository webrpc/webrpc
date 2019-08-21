package ridl

import (
	"fmt"
	"strconv"

	"github.com/webrpc/webrpc/schema"
)

var (
	schemaMessageTypeEnum   = schema.MessageType("enum")
	schemaMessageTypeStruct = schema.MessageType("struct")
)

func buildArgumentsList(s *schema.WebRPCSchema, args []*ArgumentNode) ([]*schema.MethodArgument, error) {
	output := []*schema.MethodArgument{}

	for _, arg := range args {

		var varType schema.VarType
		err := schema.ParseVarTypeExpr(s, arg.TypeName().String(), &varType)
		if err != nil {
			return nil, fmt.Errorf("unknown data type: %v", arg.TypeName().String())
		}

		methodArgument := &schema.MethodArgument{
			Name:     schema.VarName(arg.Name().String()),
			Type:     &varType,
			Stream:   arg.Stream(),
			Optional: arg.Optional(),
		}

		output = append(output, methodArgument)
	}

	return output, nil
}

func Parse(input string) (*schema.WebRPCSchema, error) {
	p, err := newParser(input)
	if err != nil {
		return nil, err
	}

	if err = p.run(); err != nil {
		return nil, err
	}

	s := &schema.WebRPCSchema{
		Imports:  []string{},
		Messages: []*schema.Message{},
		Services: []*schema.Service{},
	}

	// main definitions
	for _, line := range p.root.Definitions() {
		key, value := line.Left().String(), line.Right().String()

		switch key {
		case wordWebRPC:
			if s.WebRPCVersion != "" {
				return nil, fmt.Errorf(`webrpc was previously declared`)
			}
			s.WebRPCVersion = value
		case wordName:
			if s.Name != "" {
				return nil, fmt.Errorf(`name was previously declared`)
			}
			s.Name = value
		case wordVersion:
			if s.SchemaVersion != "" {
				return nil, fmt.Errorf(`version was previously declared`)
			}
			s.SchemaVersion = value
		default:
			return nil, fmt.Errorf("unknown definition %q", key)
		}
	}

	// imports
	for _, line := range p.root.Imports() {
		s.Imports = append(s.Imports, line.Path())
	}

	// pushing enums (1st pass)
	for _, line := range p.root.Enums() {
		s.Messages = append(s.Messages, &schema.Message{
			Name:   schema.VarName(line.Name().String()),
			Type:   schemaMessageTypeEnum,
			Fields: []*schema.MessageField{},
		})
	}

	// pushing messages (1st pass)
	for _, line := range p.root.Messages() {
		s.Messages = append(s.Messages, &schema.Message{
			Name: schema.VarName(line.Name().String()),
			Type: schemaMessageTypeStruct,
		})
	}

	// enum fields
	for _, line := range p.root.Enums() {
		name := schema.VarName(line.Name().String())
		enumDef := s.GetMessageByName(string(name))

		if enumDef == nil {
			return nil, fmt.Errorf("unexpected error, could not find definition for: %v", name)
		}

		var enumType schema.VarType
		err := schema.ParseVarTypeExpr(s, line.TypeName().String(), &enumType)
		if err != nil {
			return nil, fmt.Errorf("unknown data type: %v", line.TypeName())
		}

		for i, def := range line.Values() {
			key, val := def.Left().String(), def.Right().String()

			if val == "" {
				val = strconv.Itoa(i)
			}

			enumDef.Fields = append(enumDef.Fields, &schema.MessageField{
				Name:  schema.VarName(key),
				Type:  &enumType,
				Value: val,
			})
		}

		enumDef.EnumType = &enumType
	}

	// message fields
	for _, line := range p.root.Messages() {
		name := schema.VarName(line.Name().String())
		messageDef := s.GetMessageByName(string(name))

		if messageDef == nil {
			return nil, fmt.Errorf("unexpected error, could not find definition for: %v", name)
		}

		for _, def := range line.Fields() {
			fieldName, fieldType := def.Left().String(), def.Right().String()

			var varType schema.VarType
			err := schema.ParseVarTypeExpr(s, fieldType, &varType)
			if err != nil {
				return nil, fmt.Errorf("unknown data type: %v", fieldType)
			}

			field := &schema.MessageField{
				Name:     schema.VarName(fieldName),
				Optional: def.Optional(),
				Type:     &varType,
			}
			for _, meta := range def.Meta() {
				key, val := meta.Left().String(), meta.Right().String()
				field.Meta = append(field.Meta, schema.MessageFieldMeta{
					key: val,
				})
			}
			messageDef.Fields = append(messageDef.Fields, field)
		}
	}

	// Services
	for _, service := range p.root.Services() {
		methods := []*schema.Method{}

		for _, method := range service.Methods() {

			inputs, err := buildArgumentsList(s, method.Inputs())
			if err != nil {
				return nil, err
			}

			outputs, err := buildArgumentsList(s, method.Outputs())
			if err != nil {
				return nil, err
			}

			// push method
			methods = append(methods, &schema.Method{
				Name:         schema.VarName(method.Name().String()),
				StreamInput:  method.StreamInput(),
				StreamOutput: method.StreamOutput(),
				Inputs:       inputs,
				Outputs:      outputs,
			})
		}

		// push service
		s.Services = append(s.Services, &schema.Service{
			Name:    schema.VarName(service.Name().String()),
			Methods: methods,
		})
	}

	// run through schema validator, last step to ensure all is good.
	err = s.Parse(nil)
	if err != nil {
		return nil, err
	}

	return s, nil
}
