package schema

import (
	"fmt"
	"strings"
	"unicode"
)

const (
	TypeKind_Struct = "struct"
	TypeKind_Enum   = "enum"
)

type Type struct {
	Kind      string       `json:"kind"`
	Name      string       `json:"name"`
	Type      *VarType     `json:"type,omitempty"`
	Fields    []*TypeField `json:"fields,omitempty"`
	TypeExtra `json:",omitempty"`
}

type TypeField struct {
	Name      string   `json:"name"`
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

func (t *Type) Parse(schema *WebRPCSchema) error {
	// Type name
	typName := string(t.Name)
	if typName == "" {
		return fmt.Errorf("schema error: type name cannot be empty")
	}
	if !startsWithUpper(typName) {
		return fmt.Errorf("schema error: type name must start with upper case for '%s'", typName)
	}

	// Ensure we don't have dupe types (w/ normalization)
	name := strings.ToLower(typName)
	for _, msg := range schema.Types {
		if msg != t && name == strings.ToLower(string(msg.Name)) {
			return fmt.Errorf("schema error: duplicate type detected, '%s'", typName)
		}
	}

	// Ensure we have a valid kind
	if t.Kind != TypeKind_Enum && t.Kind != TypeKind_Struct {
		return fmt.Errorf("schema error: type must be one of 'enum', or 'struct' for '%s'", typName)
	}

	// Verify field names and ensure we don't have any duplicate field names
	// NOTE: Allow structs with no fields.
	fieldList := map[string]string{}
	jsonFieldList := map[string]string{}
	for _, field := range t.Fields {
		if string(field.Name) == "" {
			return fmt.Errorf("schema error: detected empty field name in type '%s", typName)
		}

		fieldName := string(field.Name)

		// Verify name format
		if !IsValidArgName(fieldName) {
			return fmt.Errorf("schema error: invalid field name of '%s' in type '%s'", fieldName, typName)
		}

		// Ensure no dupes
		fieldNameLower := strings.ToLower(fieldName)
		if _, ok := fieldList[fieldNameLower]; ok {
			return fmt.Errorf("schema error: detected duplicate field name of '%s' in type '%s'", fieldName, typName)
		}
		fieldList[fieldNameLower] = fieldName

		// Verify json meta format, as it overrides field name in JSON
		// and in JavaScript/TypeScript generated code.
		jsonFieldName := fieldName
		for _, meta := range field.TypeExtra.Meta {
			if jsonMeta, ok := meta["json"]; ok {
				jsonMetaString, ok := jsonMeta.(string)
				if !ok {
					return fmt.Errorf("schema error: invalid json type '%T' in field '%s' in type '%s': must be string", jsonMeta, fieldName, typName)
				}

				if jsonMetaString == "-" {
					// Skip the special `json = -` value, which makes the field ignored in generated clients.
					continue
				}

				if !IsValidArgName(jsonMetaString) {
					return fmt.Errorf("schema error: invalid json name '%s' in field '%s' in type '%s'", jsonMetaString, fieldName, typName)
				}

				jsonFieldName = jsonMetaString
			}
		}

		// Ensure no dupes
		jsonMetaStringLower := strings.ToLower(jsonFieldName)
		if _, ok := jsonFieldList[jsonMetaStringLower]; ok {
			return fmt.Errorf("schema error: detected duplicate json name '%s' in field '%s' in type '%s'", jsonFieldName, fieldName, typName)
		}
		jsonFieldList[jsonMetaStringLower] = fieldName
	}

	// For enums only, ensure all field types are the same
	if t.Kind == TypeKind_Enum {
		// ensure enum fields have value key set
		for _, field := range t.Fields {
			if field.Value == "" {
				return fmt.Errorf("schema error: enum '%s' with field '%s' is missing value", t.Name, field.Name)
			}
			if field.Type != nil {
				return fmt.Errorf("schema error: enum '%s' with field '%s', must omit 'type'", t.Name, field.Name)
			}
		}

		// ensure enum type is one of the allowed types.. aka integer
		fieldType := t.Type
		if !isValidVarType(fieldType.String(), VarIntegerCoreTypes) {
			return fmt.Errorf("schema error: enum '%s' field '%s' is invalid. must be an integer type.", t.Name, fieldType.String())
		}
	}

	// For structs only
	if t.Kind == TypeKind_Struct {
		for _, field := range t.Fields {
			// Parse+validate type fields
			err := field.Type.Parse(schema)
			if err != nil {
				return err
			}

			// Ensure value isn't set
			if field.Value != "" {
				return fmt.Errorf("schema error: struct '%s' with field '%s' cannot contain value field - please remove it", t.Name, field.Name)
			}
		}

	}

	return nil
}

func startsWithUpper(s string) bool {
	if len(s) == 0 {
		return false
	}
	if !unicode.IsUpper(rune(s[0])) {
		return false
	}
	return true
}
