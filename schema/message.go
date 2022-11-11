package schema

import (
	"fmt"
	"strings"
)

type Message struct {
	Name   VarName         `json:"name"`
	Type   MessageType     `json:"type"`
	Fields []*MessageField `json:"fields"`

	// EnumType determined for enum types during parsing time
	EnumType *VarType `json:"-"`
}

type MessageType string // "enum" | "struct"

type MessageField struct {
	Name VarName  `json:"name"`
	Type *VarType `json:"type"`

	Optional bool   `json:"optional"`
	Value    string `json:"value"` // used by enums

	// Meta store extra metadata on a field for plugins
	Meta []MessageFieldMeta `json:"meta"`
}

type MessageFieldMeta map[string]interface{}

func (m *Message) Parse(schema *WebRPCSchema) error {
	// Message name
	msgName := string(m.Name)
	if msgName == "" {
		return fmt.Errorf("schema error: message name cannot be empty")
	}

	// Ensure we don't have dupe message types (w/ normalization)
	name := strings.ToLower(msgName)
	for _, msg := range schema.Messages {
		if msg != m && name == strings.ToLower(string(msg.Name)) {
			return fmt.Errorf("schema error: duplicate message type detected, '%s'", msgName)
		}
	}

	// Ensure we have a message type
	if string(m.Type) != "enum" && string(m.Type) != "struct" {
		return fmt.Errorf("schema error: message type must be 'enum' or 'struct' for '%s'", msgName)
	}

	// NOTE: so far, lets allow messages with no fields.. so just empty object, why, I dunno, but gRPC allows it
	// Ensure we have some fields
	// if len(m.Fields) == 0 {
	// 	return errors.Errorf("schema error: message type must contain at least one field for '%s'", msgName)
	// }

	// Verify field names and ensure we don't have any duplicate field names
	fieldList := map[string]string{}
	for _, field := range m.Fields {
		if string(field.Name) == "" {
			return fmt.Errorf("schema error: detected empty field name in message '%s", msgName)
		}

		fieldName := string(field.Name)
		nFieldName := strings.ToLower(fieldName)

		// Verify name format
		if !IsValidArgName(fieldName) {
			return fmt.Errorf("schema error: invalid field name of '%s' in message '%s'", fieldName, msgName)
		}

		// Ensure no dupes
		if _, ok := fieldList[nFieldName]; ok {
			return fmt.Errorf("schema error: detected duplicate field name of '%s' in message '%s'", fieldName, msgName)
		}
		fieldList[nFieldName] = fieldName
	}

	// Parse+validate message fields
	for _, field := range m.Fields {
		err := field.Type.Parse(schema)
		if err != nil {
			return err
		}
	}

	// For enums only, ensure all field types are the same
	if m.Type == "enum" {
		// ensure enum fields have value key set, and are all of the same type
		fieldTypes := map[string]struct{}{}
		for _, field := range m.Fields {
			fieldType := field.Type.Expr
			fieldTypes[fieldType] = struct{}{}
			if field.Value == "" {
				return fmt.Errorf("schema error: enum message '%s' with field '%s' is missing value", m.Name, field.Name)
			}
		}
		if len(fieldTypes) > 1 {
			return fmt.Errorf("schema error: enum message '%s' must all have the same field type", m.Name)
		}

		// ensure enum type is one of the allowed types.. aka integer
		fieldType := m.Fields[0].Type
		if !isValidVarType(fieldType.String(), VarIntegerDataTypes) {
			return fmt.Errorf("schema error: enum message '%s' field '%s' is invalid. must be an integer type.", m.Name, fieldType.String())
		}
		m.EnumType = fieldType
	}

	// For structs only
	if m.Type == "struct" {
		for _, field := range m.Fields {
			if field.Value != "" {
				return fmt.Errorf("schema error: struct message '%s' with field '%s' cannot contain value field - please remove it", m.Name, field.Name)
			}
		}

	}

	return nil
}
