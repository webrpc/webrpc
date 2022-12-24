package ridl

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/webrpc/webrpc/schema"
)

var (
	schemaMessageTypeEnum   = schema.MessageType("enum")
	schemaMessageTypeStruct = schema.MessageType("struct")
)

type Parser struct {
	parent  *Parser
	imports map[string]struct{}

	reader *schema.Reader
}

func NewParser(r *schema.Reader) *Parser {
	return &Parser{
		reader: r,
		imports: map[string]struct{}{
			// this file imports itself
			r.File: struct{}{},
		},
	}
}

func (p *Parser) Parse() (*schema.WebRPCSchema, error) {
	s, err := p.parse()
	if err != nil {
		return nil, err
	}

	// run through schema validator, last step to ensure all is good.
	err = s.Validate()
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (p *Parser) importRIDLFile(path string) (*schema.WebRPCSchema, error) {
	if mockImport {
		return &schema.WebRPCSchema{}, nil
	}

	for node := p; node != nil; node = node.parent {
		if _, imported := node.imports[path]; imported {
			return nil, fmt.Errorf("circular import %q in file %q", filepath.Base(path), p.reader.File)
		}
		node.imports[path] = struct{}{}
	}

	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	m := NewParser(schema.NewReader(fp, path))
	m.parent = p
	return m.Parse()
}

func (p *Parser) parse() (*schema.WebRPCSchema, error) {
	q, err := newParser(p.reader)
	if err != nil {
		return nil, err
	}

	if err = q.run(); err != nil {
		return nil, err
	}

	s := &schema.WebRPCSchema{
		Messages: []*schema.Message{},
		Services: []*schema.Service{},
	}

	// main definitions
	for _, line := range q.root.Definitions() {
		key, value := line.Left().String(), line.Right().String()

		switch key {
		case wordWebRPC:
			if s.WebrpcVersion != "" {
				return nil, fmt.Errorf(`webrpc was previously declared`)
			}
			s.WebrpcVersion = value
		case wordName:
			if s.SchemaName != "" {
				return nil, fmt.Errorf(`name was previously declared`)
			}
			s.SchemaName = value
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
	for _, line := range q.root.Imports() {
		importPath := filepath.Join(filepath.Dir(p.reader.File), line.Path().String())

		imported, err := p.importRIDLFile(importPath)
		if err != nil {
			return nil, p.trace(err, line.Path())
		}

		members := []string{}
		for _, member := range line.Members() {
			members = append(members, member.String())
		}

		for i := range imported.Messages {
			if isImportAllowed(string(imported.Messages[i].Name), members) {
				s.Messages = append(s.Messages, imported.Messages[i])
			}
		}
		for i := range imported.Services {
			if isImportAllowed(string(imported.Services[i].Name), members) {
				s.Services = append(s.Services, imported.Services[i])
			}
		}
	}

	// pushing enums (1st pass)
	for _, line := range q.root.Enums() {
		s.Messages = append(s.Messages, &schema.Message{
			Name:   schema.VarName(line.Name().String()),
			Type:   schemaMessageTypeEnum,
			Fields: []*schema.MessageField{},
		})
	}

	// pushing messages (1st pass)
	for _, line := range q.root.Messages() {
		s.Messages = append(s.Messages, &schema.Message{
			Name: schema.VarName(line.Name().String()),
			Type: schemaMessageTypeStruct,
		})
	}

	// pushing services (1st pass)
	for _, service := range q.root.Services() {
		// push service
		s.Services = append(s.Services, &schema.Service{
			Name: schema.VarName(service.Name().String()),
		})
	}

	// enum fields
	for _, line := range q.root.Enums() {
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
	for _, line := range q.root.Messages() {
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
	for _, service := range q.root.Services() {
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

		serviceDef := s.GetServiceByName(service.Name().String())
		serviceDef.Methods = methods
	}

	return s, nil
}

func (p *Parser) trace(err error, tok *TokenNode) error {
	return fmt.Errorf(
		"%v\nnear string %q\n\tfrom %v:%d:%d",
		err,
		tok.tok.val,
		p.reader.File,
		tok.tok.line,
		tok.tok.col,
	)
}

func isImportAllowed(name string, whitelist []string) bool {
	if len(whitelist) < 1 {
		return true
	}
	for i := range whitelist {
		if name == whitelist[i] {
			return true
		}
	}
	return false
}

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
			Optional: arg.Optional(),
		}

		output = append(output, methodArgument)
	}

	return output, nil
}
