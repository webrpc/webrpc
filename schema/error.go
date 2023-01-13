package schema

import (
	"fmt"
	"strings"
)

type Error struct {
	Code       uint64 `json:"code"`
	Name       string `json:"name"`
	Message    string `json:"message"`
	HTTPStatus uint64 `json:"httpStatus,omitempty"`

	// Schema *WebRPCSchema `json:"-"` // denormalize/back-reference
}

func (s *Error) Parse(schema *WebRPCSchema) error {
	s.Name = strings.TrimSpace(s.Name)
	if s.Name == "" {
		return fmt.Errorf("schema error: name cannot be empty")
	}
	if s.Code == 0 {
		return fmt.Errorf("schema error: error code cannot be 0")
	}
	if s.Code < 100 {
		return fmt.Errorf("schema error: error code must 100 or greater, '%s'", s.Name)
	}
	n := strings.Fields(s.Name)
	if len(n) > 1 {
		return fmt.Errorf("schema error: name must be a single word")
	}
	if !startsWithUpper(s.Name) {
		return fmt.Errorf("schema error: error name must start with upper case for '%s'", s.Name)
	}
	if s.Message == "" {
		return fmt.Errorf("schema error: message cannot be empty")
	}
	if s.HTTPStatus > 0 && s.HTTPStatus < 400 || s.HTTPStatus >= 600 {
		return fmt.Errorf("schema error: httpStatus number is invalid for error type '%s'", s.Name)
	}

	// check for duplicate codes or names
	nameList := map[string]struct{}{}
	codeList := map[uint64]struct{}{}
	for _, e := range schema.Errors {
		name := strings.ToLower(e.Name)
		if _, ok := nameList[name]; ok {
			return fmt.Errorf("schema error: detected duplicate error name of '%s'", e.Name)
		}
		if _, ok := codeList[e.Code]; ok {
			return fmt.Errorf("schema error: detected duplicate error code of '%d'", e.Code)
		}
		nameList[name] = struct{}{}
		codeList[e.Code] = struct{}{}
	}

	return nil
}
