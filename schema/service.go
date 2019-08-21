package schema

import (
	"strings"

	"github.com/pkg/errors"
)

type Service struct {
	Name    VarName   `json:"name"`
	Methods []*Method `json:"methods"`

	Schema *WebRPCSchema `json:"-"` // denormalize/back-reference
}

type Method struct {
	Name VarName `json:"name"`

	StreamInput  bool `json:"-"` // TODO: actual implementation
	StreamOutput bool `json:"-"` // TODO: actual implementation
	Proxy        bool `json:"-"` // TODO: actual implementation

	Inputs  []*MethodArgument `json:"inputs"`
	Outputs []*MethodArgument `json:"outputs"`

	Service *Service `json:"-"` // denormalize/back-reference
}

type MethodArgument struct {
	Name     VarName  `json:"name"`
	Type     *VarType `json:"type"`
	Optional bool     `json:"optional"`
	Stream   bool     `json:"stream"` // TOOD(future)

	InputArg  bool `json:"-"` // denormalize/back-reference
	OutputArg bool `json:"-"` // denormalize/back-reference
}

func (s *Service) Parse(schema *WebRPCSchema) error {
	s.Schema = schema // back-ref

	// Service name
	serviceName := string(s.Name)
	if string(s.Name) == "" {
		return errors.Errorf("schema error: service name cannot be empty")
	}

	// Ensure we don't have dupe service names (w/ normalization)
	name := strings.ToLower(string(s.Name))
	for _, svc := range schema.Services {
		if svc != s && name == strings.ToLower(string(svc.Name)) {
			return errors.Errorf("schema error: duplicate service name detected in service '%s'", serviceName)
		}
	}

	// Ensure methods is defined
	if len(s.Methods) == 0 {
		return errors.Errorf("schema error: methods cannot be empty for service '%s'", serviceName)
	}

	// Verify method names and ensure we don't have any duplicate method names
	methodList := map[string]string{}
	for _, method := range s.Methods {
		if string(method.Name) == "" {
			return errors.Errorf("schema error: detected empty method name in service '%s", serviceName)
		}

		methodName := string(method.Name)
		nMethodName := strings.ToLower(methodName)

		if _, ok := methodList[nMethodName]; ok {
			return errors.Errorf("schema error: detected duplicate method name of '%s' in service '%s'", methodName, serviceName)
		}
		methodList[nMethodName] = methodName
	}

	// Parse+validate methods
	for _, method := range s.Methods {
		err := method.Parse(schema, s)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Method) Parse(schema *WebRPCSchema, service *Service) error {
	if service == nil {
		return errors.Errorf("parse error, service arg cannot be nil")
	}
	m.Service = service // back-ref
	serviceName := string(service.Name)

	// Parse+validate inputs
	for _, input := range m.Inputs {
		input.InputArg = true // back-ref
		if input.Name == "" {
			return errors.Errorf("schema error: detected empty input argument name for method '%s' in service '%s'", m.Name, serviceName)
		}
		err := input.Type.Parse(schema)
		if err != nil {
			return err
		}
	}

	// Parse+validate outputs
	for _, output := range m.Outputs {
		output.OutputArg = true // back-ref
		if output.Name == "" {
			return errors.Errorf("schema error: detected empty output name for method '%s' in service '%s'", m.Name, serviceName)
		}
		err := output.Type.Parse(schema)
		if err != nil {
			return err
		}
	}

	// Note, we allow zero inputs and zero outputs

	return nil
}
