package schema

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
)

const (
	VERSION = "v1"
)

type Import struct {
	Path    string   `json:"path"`
	Members []string `json:"members"`
}

// schema of webrpc json file, and validations
type WebRPCSchema struct {
	WebRPCVersion string    `json:"webrpc"`
	Name          string    `json:"name"`
	SchemaVersion string    `json:"version"`
	Imports       []*Import `json:"imports"`

	Messages []*Message `json:"messages"`
	Services []*Service `json:"services"`
}

// Parse parses the schema through the AST, intended to be called
// after the json has been unmarshalled
func (s *WebRPCSchema) Parse(schema *WebRPCSchema) error {
	if schema == nil {
		schema = s // we can take nil here, as this is the schema
	}

	if schema.WebRPCVersion != VERSION {
		return errors.Errorf("webrpc schema version, '%s' is invalid, try '%s'", schema.WebRPCVersion, VERSION)
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

func (s *WebRPCSchema) SchemaHash() (string, error) {
	// TODO: lets later make this even more deterministic in face of re-ordering
	// definitions within the ridl file
	jsonString, err := s.ToJSON(false)
	if err != nil {
		return "", err
	}

	h := sha1.New()
	h.Write([]byte(jsonString))
	return hex.EncodeToString(h.Sum(nil)), nil
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

func (s *WebRPCSchema) GetMessageByName(name string) *Message {
	name = strings.ToLower(name)
	for _, message := range s.Messages {
		if strings.ToLower(string(message.Name)) == name {
			return message
		}
	}
	return nil
}

func (s *WebRPCSchema) GetServiceByName(name string) *Service {
	name = strings.ToLower(name)
	for _, service := range s.Services {
		if strings.ToLower(string(service.Name)) == name {
			return service
		}
	}
	return nil
}

func (s *WebRPCSchema) HasFieldType(fieldType string) (bool, error) {
	fieldType = strings.ToLower(fieldType)
	_, ok := DataTypeFromString[fieldType]
	if !ok {
		return false, errors.Errorf("webrpc: invalid data type '%s'", fieldType)
	}

	for _, m := range s.Messages {
		for _, f := range m.Fields {
			if DataTypeToString[f.Type.Type] == fieldType {
				return true, nil
			}
		}
	}

	for _, s := range s.Services {
		for _, m := range s.Methods {
			for _, i := range m.Inputs {
				if DataTypeToString[i.Type.Type] == fieldType {
					return true, nil
				}
			}
			for _, o := range m.Outputs {
				if DataTypeToString[o.Type.Type] == fieldType {
					return true, nil
				}
			}
		}
	}

	return false, nil
}
