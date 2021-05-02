package schema

import (
	"fmt"
	"strings"
)

type Type struct {
	Kind      string       `json:"kind"`
	Name      VarName      `json:"name"`
	Type      *VarType     `json:"type,omitempty"`
	Fields    []*TypeField `json:"fields,omitempty"`
	TypeExtra `json:",omitempty"`
}

type TypeField struct {
	Name      VarName  `json:"name"`
	Type      *VarType `json:"type,omitempty"`
	TypeExtra `json:",omitempty"`
}

type TypeExtra struct {
	Optional bool   `json:"optional,omitempty"` // used by structs
	Value    string `json:"value,omitempty"`    // used by enums

	// Meta store extra metadata on a field for plugins
	Meta []TypeFieldMeta `json:"meta,omitempty"`
}

type TypeFieldMeta map[string]interface{}

func (m *Type) Parse(schema *WebRPCSchema) error {
	// Type name
	typName := string(m.Name)
	if typName == "" {
		return fmt.Errorf("schema error: type name cannot be empty")
	}

	// Ensure we don't have dupe types (w/ normalization)
	name := strings.ToLower(typName)
	for _, msg := range schema.Types {
		if msg != m && name == strings.ToLower(string(msg.Name)) {
			return fmt.Errorf("schema error: duplicate type detected, '%s'", typName)
		}
	}

	// Ensure we have a valid kind
	if m.Kind != "type" && m.Kind != "enum" && m.Kind != "struct" {
		return fmt.Errorf("schema error: type must be one of 'type', 'enum' or 'struct' for '%s'", typName)
	}

	// NOTE: so far, lets allow messages with no fields.. so just empty object, why, I dunno, but gRPC allows it
	// Ensure we have some fields
	// if len(m.Fields) == 0 {
	// 	return fmt.Errorf("schema error: message type must contain at least one field for '%s'", msgName)
	// }

	// Verify field names and ensure we don't have any duplicate field names
	fieldList := map[string]string{}
	for _, field := range m.Fields {
		if string(field.Name) == "" {
			return fmt.Errorf("schema error: detected empty field name in type '%s", typName)
		}

		fieldName := string(field.Name)
		nFieldName := strings.ToLower(fieldName)

		// Verify name format
		if !IsValidArgName(fieldName) {
			return fmt.Errorf("schema error: invalid field name of '%s' in type '%s'", fieldName, typName)
		}

		// Ensure no dupes
		if _, ok := fieldList[nFieldName]; ok {
			return fmt.Errorf("schema error: detected duplicate field name of '%s' in type '%s'", fieldName, typName)
		}
		fieldList[nFieldName] = fieldName
	}

	// For type only
	if m.Kind == "type" {
		// ensure valid type expression
		err := m.Type.Parse(schema)
		if err != nil {
			return err
		}
	}

	// For enums only, ensure all field types are the same
	if m.Kind == "enum" {
		// ensure enum fields have value key set
		for _, field := range m.Fields {
			if field.Value == "" {
				return fmt.Errorf("schema error: enum '%s' with field '%s' is missing value", m.Name, field.Name)
			}
			if field.Type != nil {
				return fmt.Errorf("schema error: enum '%s' with field '%s', must omit 'type'", m.Name, field.Name)
			}
		}

		// ensure enum type is one of the allowed types.. aka integer
		fieldType := m.Type
		if !isValidVarType(fieldType.String(), VarIntegerCoreTypes) {
			return fmt.Errorf("schema error: enum '%s' field '%s' is invalid. must be an integer type.", m.Name, fieldType.String())
		}
	}

	// For structs only
	if m.Kind == "struct" {
		for _, field := range m.Fields {
			// Parse+validate type fields
			err := field.Type.Parse(schema)
			if err != nil {
				return err
			}

			// Ensure value isn't set
			if field.Value != "" {
				return fmt.Errorf("schema error: struct '%s' with field '%s' cannot contain value field - please remove it", m.Name, field.Name)
			}
		}

	}

	return nil
}
