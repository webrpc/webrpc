package schema

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
)

const (
	SCHEMA_VERSION = "v1"
)

// schema of webrpc json file, and validations
type WebRPCSchema struct {
	WebrpcVersion string `json:"webrpc"`
	SchemaName    string `json:"name"`
	SchemaVersion string `json:"version"`

	Types    []*Type    `json:"types"`
	Services []*Service `json:"services"`

	// Deprecated. Renamed to Types. Keep this field for now, so we can
	// error out & advise users to migrate to v0.9.0+ schema format.
	Deprecated_Messages []interface{} `json:"messages"`
}

// Validate validates the schema through the AST, intended to be called after
// the json has been unmarshalled
func (s *WebRPCSchema) Validate() error {
	if s.WebrpcVersion != SCHEMA_VERSION {
		return fmt.Errorf("webrpc schema version, '%s' is invalid, try '%s'", s.WebrpcVersion, SCHEMA_VERSION)
	}

	for _, msg := range s.Types {
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

	if len(s.Deprecated_Messages) > 0 {
		return fmt.Errorf(" field \"messages\" was renamed to \"types\", see https://github.com/webrpc/webrpc/tree/master/CHANGELOG.md#JSON+schema+v0.9.0+migration+guide")
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

func (s *WebRPCSchema) GetTypeByName(name string) *Type {
	name = strings.ToLower(name)
	for _, message := range s.Types {
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
	_, ok := CoreTypeFromString[fieldType]
	if !ok {
		return false, fmt.Errorf("webrpc: invalid data type '%s'", fieldType)
	}

	for _, m := range s.Types {
		for _, f := range m.Fields {
			if CoreTypeToString[f.Type.Type] == fieldType {
				return true, nil
			}
		}
	}

	for _, s := range s.Services {
		for _, m := range s.Methods {
			for _, i := range m.Inputs {
				if CoreTypeToString[i.Type.Type] == fieldType {
					return true, nil
				}
			}
			for _, o := range m.Outputs {
				if CoreTypeToString[o.Type.Type] == fieldType {
					return true, nil
				}
			}
		}
	}

	return false, nil
}
