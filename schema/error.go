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
	if s.Code == 0 {
		return fmt.Errorf("schema error: error code cannot be 0")
	}
	s.Name = strings.TrimSpace(s.Name)
	if s.Name == "" {
		return fmt.Errorf("schema error: name cannot be empty")
	}
	n := strings.Fields(s.Name)
	if len(n) > 1 {
		return fmt.Errorf("schema error: name must be a single word")
	}
	if !IsStartsWithUpper(s.Name) {
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
	for _, serr := range schema.Errors {
		nName := strings.ToLower(serr.Name)
		if _, ok := nameList[nName]; ok {
			return fmt.Errorf("schema error: detected duplicate error name of '%s'", serr.Name)
		}
		if _, ok := codeList[serr.Code]; ok {
			return fmt.Errorf("schema error: detected duplicate error code of '%d'", serr.Code)
		}
		nameList[nName] = struct{}{}
		codeList[serr.Code] = struct{}{}
	}

	return nil
}
