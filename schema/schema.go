package schema

import (
	"bytes"
	"encoding/json"

	"github.com/pkg/errors"
)

const (
	VERSION = "v1"
)

// schema of webrpc json file, and validations
type WebRPCSchema struct {
	Schema  string   `json:"webrpc"`
	Name    string   `json:"name"`
	Version string   `json:"version"`
	Imports []string `json:"imports"`

	Messages []*Message `json:"messages"`
	Services []*Service `json:"services"`
}

// Parse parses the schema through the AST, intended to be called
// after the json has been unmarshalled
func (s *WebRPCSchema) Parse(schema *WebRPCSchema) error {
	if schema == nil {
		schema = s // we can take nil here, as this is the schema
	}

	if schema.Schema != VERSION {
		return errors.Errorf("webrpc schema version, '%s' is invalid, try '%s'", schema.Schema, VERSION)
	}

	for _, msg := range s.Messages {
		err := msg.Parse(s)
		if err != nil {
			return err
		}
	}
	for _, svc := range s.Services {
		err := svc.Parse(s)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *WebRPCSchema) ToJSON(optIndent ...bool) (string, error) {
	indent := false
	if len(optIndent) > 0 {
		indent = optIndent[0]
	}

	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	if indent {
		enc.SetIndent("", " ")
	}

	err := enc.Encode(s)
	if err != nil {
		return "", err
	}

	return string(buf.Bytes()), nil
}
