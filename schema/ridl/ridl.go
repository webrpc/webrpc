package ridl

import (
	"fmt"
	"io"
	"io/fs"
	"path"
	"strconv"
	"strings"

	"github.com/webrpc/webrpc/schema"
)

var (
	schemaTypeKindEnum   = "enum"
	schemaTypeKindStruct = "struct"
)

type Parser struct {
	parent  *Parser
	imports map[string]struct{}

	reader io.Reader
	path   string
	fsys   fs.FS
}

func NewParser(fsys fs.FS, path string) *Parser {
	return &Parser{
		fsys: fsys,
		path: path,
		imports: map[string]struct{}{
			// this file imports itself
			path: {},
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

func (p *Parser) importRIDLFile(filename string) (*schema.WebRPCSchema, error) {
	for node := p; node != nil; node = node.parent {
		if _, imported := node.imports[filename]; imported {
			return nil, fmt.Errorf("circular import %q in file %q", path.Base(filename), p.path)
		}
		node.imports[filename] = struct{}{}
	}

	m := NewParser(p.fsys, filename)
	m.parent = p
	return m.Parse()
}

func (p *Parser) parse() (*schema.WebRPCSchema, error) {
	if !fs.ValidPath(p.path) {
		return nil, fmt.Errorf("invalid fs.FS path %q, see https://pkg.go.dev/io/fs#ValidPath", p.path)
	}

	src, err := fs.ReadFile(p.fsys, p.path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	q, err := newParser(src)
	if err != nil {
		return nil, err
	}

	if err = q.run(); err != nil {
		return nil, err
	}

	s := &schema.WebRPCSchema{
		Types:    []*schema.Type{},
		Errors:   []*schema.Error{},
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
		importPath := path.Join(path.Dir(p.path), line.Path().String())

		imported, err := p.importRIDLFile(importPath)
		if err != nil {
			return nil, p.trace(err, line.Path())
		}

		members := []string{}
		for _, member := range line.Members() {
			members = append(members, member.String())
		}

		for i := range imported.Types {
			if isImportAllowed(imported.Types[i].Name, members) {
				s.Types = append(s.Types, imported.Types[i])
			}
		}
		for i := range imported.Services {
			if isImportAllowed(imported.Services[i].Name, members) {
				s.Services = append(s.Services, imported.Services[i])
			}
		}
	}

	// pushing enums (1st pass)
	for _, line := range q.root.Enums() {
		s.Types = append(s.Types, &schema.Type{
			Kind:   schemaTypeKindEnum,
			Name:   line.Name().String(),
			Fields: []*schema.TypeField{},
		})
	}

	// pushing types (1st pass)
	for _, line := range q.root.Structs() {
		s.Types = append(s.Types, &schema.Type{
			Kind:     schemaTypeKindStruct,
			Name:     line.Name().String(),
			Comments: parseComment(line.Comment()),
		})
	}

	// pushing services (1st pass)
	for _, service := range q.root.Services() {
		srv := &schema.Service{
			Name:     service.Name().String(),
			Comments: parseComment(service.Comment()),
		}

		s.Services = append(s.Services, srv)
	}

	// enum fields
	for _, line := range q.root.Enums() {
		name := line.Name().String()
		enumDef := s.GetTypeByName(name)

		if enumDef == nil {
			return nil, fmt.Errorf("unexpected error, could not find definition for: %v", name)
		}

		var enumType schema.VarType
		err := schema.ParseVarTypeExpr(s, line.TypeName().String(), &enumType)
		if err != nil {
			return nil, fmt.Errorf("enum %q: unknown type: %v", name, line.TypeName())
		}
		enumDef.Type = &enumType

		for i, def := range line.Values() {
			key, val := def.Left().String(), def.Right().String()
			if val == "" {
				val = strconv.Itoa(i)
			}

			elems := &schema.TypeField{
				Name: key,
				TypeExtra: schema.TypeExtra{
					Value: val,
				},
				Comments: parseComment(def.Comment()),
			}

			enumDef.Fields = append(enumDef.Fields, elems)
		}
	}

	// error types
	for _, line := range q.root.Errors() {
		var errorType schema.Error
		code, _ := strconv.ParseInt(line.code.String(), 10, 32)
		errorType.Code = int(code)
		errorType.Name = line.name.String()
		errorType.Message = line.message.String()
		if line.httpStatus != nil {
			httpStatus, _ := strconv.ParseInt(line.httpStatus.String(), 10, 32)
			errorType.HTTPStatus = int(httpStatus)
		} else {
			errorType.HTTPStatus = 400 // Default HTTP status code
		}
		s.Errors = append(s.Errors, &errorType)
	}

	// struct fields
	for _, line := range q.root.Structs() {
		name := line.Name().String()
		structDef := s.GetTypeByName(string(name))

		if structDef == nil {
			return nil, fmt.Errorf("unexpected error, could not find definition for: %v", name)
		}

		for _, def := range line.Fields() {
			fieldName, fieldType := def.Left().String(), def.Right().String()

			var varType schema.VarType
			err := schema.ParseVarTypeExpr(s, fieldType, &varType)
			if err != nil {
				return nil, fmt.Errorf("struct %q: unknown type of field %q: %v", name, fieldName, fieldType)
			}

			field := &schema.TypeField{
				Name: fieldName,
				Type: &varType,
				TypeExtra: schema.TypeExtra{
					Optional: def.Optional(),
				},
				Comments: parseComment(def.Comment()),
			}
			for _, meta := range def.Meta() {
				key, val := meta.Left().String(), meta.Right().String()
				field.Meta = append(field.Meta, schema.TypeFieldMeta{
					key: val,
				})
			}
			structDef.Fields = append(structDef.Fields, field)
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

			// push m
			m := &schema.Method{
				Name:         method.Name().String(),
				StreamInput:  method.StreamInput(),
				StreamOutput: method.StreamOutput(),
				Inputs:       inputs,
				Outputs:      outputs,
				Comments:     parseComment(method.Comment()),
				Annotations:  buildAnnotations(method),
			}

			methods = append(methods, m)
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
		p.path,
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

	// succint form
	if len(args) == 1 && args[0].inlineStruct != nil {
		node := args[0].inlineStruct
		structName := node.tok.val

		typ := s.GetTypeByName(structName)
		if typ.Kind != "struct" {
			return nil, fmt.Errorf("expecting struct type for inline definition of '%s'", structName)
		}

		for _, arg := range typ.Fields {
			methodArgument := &schema.MethodArgument{
				Name:      arg.Name,
				Type:      arg.Type,
				Optional:  arg.Optional,
				TypeExtra: arg.TypeExtra,
			}
			output = append(output, methodArgument)
		}

		return output, nil
	}

	// normal form
	for _, arg := range args {
		var varType schema.VarType
		err := schema.ParseVarTypeExpr(s, arg.TypeName().String(), &varType)
		if err != nil {
			return nil, fmt.Errorf("parsing argument %v: %w", arg.TypeName(), err)
		}

		methodArgument := &schema.MethodArgument{
			Name:     arg.Name().String(),
			Type:     &varType,
			Optional: arg.Optional(),
		}
		output = append(output, methodArgument)
	}

	return output, nil
}

func parseComment(comment string) []string {
	if comment == "" {
		return []string{}
	}

	return strings.Split(comment, "\n")
}

func buildAnnotations(method *MethodNode) schema.Annotations {
	annotations := make(map[string]*schema.Annotation)

	for _, a := range method.Annotations() {
		an := &schema.Annotation{
			AnnotationType: a.AnnotationType().String(),
		}
		if a.Value() != nil {
			an.Value = a.Value().String()
		}

		annotations[a.AnnotationType().String()] = an
	}

	return schema.Annotations(annotations)
}
