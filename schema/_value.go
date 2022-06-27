package schema

import (
	"fmt"
	"strings"
)

// in the future..

type Value struct {
	Name  string      `json:"name"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`

	// Schema *WebRPCSchema `json:"-"` // denormalize/back-reference
}

func (s *Value) Parse(schema *WebRPCSchema) error {
	if s.Name == "" {
		return fmt.Errorf("schema error: name cannot be empty")
	}
	if s.Value == nil {
		return fmt.Errorf("schema error: value cannot be empty")
	}

	// check for duplicate names
	nameList := map[string]struct{}{}
	for _, sval := range schema.Values {
		nName := strings.ToLower(sval.Name)
		if _, ok := nameList[nName]; ok {
			return fmt.Errorf("schema error: detected duplicate value name of '%s'", sval.Name)
		}
		nameList[nName] = struct{}{}
	}

	return nil
}
