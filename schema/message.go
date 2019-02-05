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

func (m *Message) Parse(schema *WebRPCSchema) error {
	// Ensure we don't have dupe message types (w/ normalization)
	name := strings.ToLower(m.Name.String())
	for _, msg := range schema.Messages {
		if msg != m && name == strings.ToLower(msg.Name.String()) {
			return errors.Errorf("parse error: duplicate message type detected, '%s'", m.Name)
		}
	}

	// Ensure we have a message type
	if m.Type == nil || (string(*m.Type) != "enum" && string(*m.Type) != "struct") {
		return errors.Errorf("parse error: message type must be 'enum' or 'struct' for '%s'", m.Name.String())
	}

	// Ensure we have some fields
	if len(m.Fields) == 0 {
		return errors.Errorf("parse error: message type must contain at least one field for '%s'", m.Name.String())
	}

	// Parse message fields
	for _, field := range m.Fields {
		err := field.Type.Parse(schema)
		if err != nil {
			return err
		}
	}

	return nil
}

// NOTE: for now just skip the enum stuff during generation..
type MessageType string // TODO: must be `enum` or `struct`

type MessageField struct {
	Name *VarName `json:"name"`
	Type *VarType `json:"type"`

	Optional bool   `json:"optional"`
	Value    string `json:"value"` // used by enums

	// Tags store extra metadata for plugins
	Tags []MessageFieldTag `json:"tags"`
}

type MessageFieldTag map[string]interface{}
