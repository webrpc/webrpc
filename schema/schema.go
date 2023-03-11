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
	Errors   []*Error   `json:"errors"`
	Services []*Service `json:"services"`

	// Deprecated. Renamed to Types. Keep this field for now, so we can
	// error out & advise users to migrate to v0.9.0+ schema format.
	Deprecated_Messages []interface{} `json:"messages,omitempty"`
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

	for _, e := range s.Errors {
		err := e.Parse(s)
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
		return fmt.Errorf(" field \"messages\" was renamed to \"types\", see https://github.com/webrpc/webrpc/blob/master/CHANGELOG.md#json-schema-v090-migration-guide")
	}

	return nil
}

func (s *WebRPCSchema) SchemaHash() (string, error) {
	// TODO: lets later make this even more deterministic in face of re-ordering
	// definitions within the ridl file

	jsonString, err := s.ToJSON()
	if err != nil {
		return "", err
	}

	h := sha1.New()
	h.Write([]byte(jsonString))
	return hex.EncodeToString(h.Sum(nil)), nil
}

func (s *WebRPCSchema) ToJSON() (string, error) {
	var buf bytes.Buffer

	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", " ")

	if err := enc.Encode(s); err != nil {
		return "", err
	}

	return buf.String(), nil
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
