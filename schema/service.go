package schema

import (
	"strings"

	"github.com/pkg/errors"
)

type Service struct {
	Name    VarName   `json:"name"`
	Methods []*Method `json:"methods"`
}

type Method struct {
	Name    VarName         `json:"name"`
	Inputs  []*MethodInput  `json:"inputs"`
	Outputs []*MethodOutput `json:"outputs`
}

type MethodInput struct {
	Name   VarName        `json:"name"`
	Type   MethodArgument `json:"type"`   // TODO: VarType? Argument? ..perhaps ServiceArgument if we only allow struct..
	Stream bool           `json:"stream"` // TOOD(future)
}

type MethodOutput struct {
	Name   VarName        `json:"name"`
	Type   MethodArgument `json:"type"`   // TODO: same as input above..
	Stream bool           `json:"stream"` // TOOD(future)
}

type MethodArgument string

func (s *Service) Parse(schema *WebRPCSchema) error {
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
		err := method.Parse(schema)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Method) Parse(schema *WebRPCSchema) error {
	return nil
}
