package schema

import (
	"strings"

	"github.com/pkg/errors"
)

type Message struct {
	Name   *VarName        `json:"name"`
	Type   *MessageType    `json:"type"`
	Fields []*MessageField `json:"fields"`
}

type MessageType string // "enum" | "struct"

type MessageField struct {
	Name *VarName `json:"name"`
	Type *VarType `json:"type"`

	Optional bool   `json:"optional"`
	Value    string `json:"value"` // used by enums

	// Tags store extra metadata for plugins
	Tags []MessageFieldTag `json:"tags"`
}

type MessageFieldTag map[string]interface{}

func (m *Message) Parse(schema *WebRPCSchema) error {
	// Message name
	if m.Name == nil || string(*m.Name) == "" {
		return errors.Errorf("schema error: message name cannot be empty")
	}
	msgName := string(*m.Name)

	// Ensure we don't have dupe message types (w/ normalization)
	name := strings.ToLower(msgName)
	for _, msg := range schema.Messages {
		if msg != m && name == strings.ToLower(msg.Name.String()) {
			return errors.Errorf("schema error: duplicate message type detected, '%s'", msgName)
		}
	}

	// Ensure we have a message type
	if m.Type == nil || (string(*m.Type) != "enum" && string(*m.Type) != "struct") {
		return errors.Errorf("schema error: message type must be 'enum' or 'struct' for '%s'", msgName)
	}

	// Ensure we have some fields
	if len(m.Fields) == 0 {
		return errors.Errorf("schema error: message type must contain at least one field for '%s'", msgName)
	}

	// Verify field names and ensure we don't have any duplicate field names
	fieldList := map[string]string{}
	for _, field := range m.Fields {
		if field.Name == nil || string(*field.Name) == "" {
			return errors.Errorf("schema error: detected empty field name in message '%s", msgName)
		}

		fieldName := string(*field.Name)
		nFieldName := strings.ToLower(fieldName)

		if _, ok := fieldList[nFieldName]; ok {
			return errors.Errorf("schema error: detected duplicate field name of '%s' in message '%s'", fieldName, msgName)
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

	return nil
}
