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
}

// Validate validates the schema through the AST, intended to be called after
// the json has been unmarshalled
func (s *WebRPCSchema) Validate() error {
	if s.WebrpcVersion != SCHEMA_VERSION {
		return fmt.Errorf("webrpc schema version, '%s' is invalid, try '%s'", s.WebrpcVersion, SCHEMA_VERSION)
	}

	for i := range s.Types {
		err := s.Types[i].Parse(s)
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
		if strings.ToLower(message.Name) == name {
			return message
		}
	}
	return nil
}

func (s *WebRPCSchema) GetServiceByName(name string) *Service {
	name = strings.ToLower(name)
	for _, service := range s.Services {
		if strings.ToLower(service.Name) == name {
			return service
		}
	}
	return nil
}

func MatchServices(s *WebRPCSchema, services []string) *WebRPCSchema {
	if s == nil {
		return nil
	}

	if len(services) == 0 {
		return s
	}

	serviceMap := map[string]struct{}{}
	for _, srv := range services {
		serviceMap[srv] = struct{}{}
	}

	var matchedServices []*Service
	for _, srv := range s.Services {
		if _, ok := serviceMap[srv.Name]; ok {
			matchedServices = append(matchedServices, srv)
		}
	}

	s.Services = matchedServices

	return s
}

func MethodTreeShake(s *WebRPCSchema) *WebRPCSchema {
	if s == nil {
		return nil
	}

	// Collect type names referenced by service methods (including nested references).
	used := map[string]bool{}

	// Helper: recursively walk VarType and mark used type names, including nested struct fields and map/list elements.
	var walkVarType func(vt *VarType)
	walkVarType = func(vt *VarType) {
		if vt == nil {
			return
		}

		// List element
		if vt.List != nil {
			walkVarType(vt.List.Elem)
		}

		// Map key/value
		if vt.Map != nil {
			walkVarType(vt.Map.Key)
			walkVarType(vt.Map.Value)
		}

		// Struct type: mark and traverse its fields
		if vt.Struct != nil && vt.Struct.Type != nil {
			t := vt.Struct.Type
			if t.Name != "" {
				used[t.Name] = true
			}
			// Traverse struct fields to include transitive type usage
			for _, f := range t.Fields {
				walkVarType(f.Type)
			}
		}

		// Enum type: mark usage
		if vt.Enum != nil && vt.Enum.Type != nil {
			t := vt.Enum.Type
			if t.Name != "" {
				used[t.Name] = true
			}
		}
	}

	// Traverse all methods of all services to collect referenced types.
	for _, srv := range s.Services {
		for _, m := range srv.Methods {
			for _, in := range m.Inputs {
				walkVarType(in.Type)
			}
			for _, out := range m.Outputs {
				walkVarType(out.Type)
			}
		}
	}

	// Filter types: keep only those marked as used.
	var kept []*Type
	for _, t := range s.Types {
		if used[t.Name] {
			kept = append(kept, t)
		}
	}
	s.Types = kept

	return s
}

func MatchMethodsWithAnnotations(s *WebRPCSchema, annotations map[string]string) *WebRPCSchema {
	if s == nil {
		return nil
	}

	for i, srv := range s.Services {
		var methods []*Method
		for _, method := range srv.Methods {
			match := false
			for name, value := range annotations {
				ann, ok := method.Annotations[name]
				if !ok {
					continue
				}

				trimmedAnnVal := strings.Trim(value, " ")
				if trimmedAnnVal == "" {
					// match annotation types only
					if name == ann.AnnotationType {
						match = true
						break
					}
				} else {
					// match @key:value
					if name == ann.AnnotationType && value == ann.Value {
						match = true
						break
					}
				}
			}

			if match {
				methods = append(methods, method)
			}
		}

		s.Services[i].Methods = methods
	}

	return s
}

func IgnoreMethodsWithAnnotations(s *WebRPCSchema, annotations map[string]string) *WebRPCSchema {
	if s == nil {
		return nil
	}

	for i, srv := range s.Services {
		var methods []*Method
		for _, method := range srv.Methods {
			ignore := false
			for name, value := range annotations {
				ann, ok := method.Annotations[name]
				if !ok {
					continue
				}

				trimmedAnnVal := strings.Trim(value, " ")
				if trimmedAnnVal == "" {
					// match annotation types only
					if name == ann.AnnotationType {
						ignore = true
						break
					}
				} else {
					// match @key:value
					if name == ann.AnnotationType && value == ann.Value {
						ignore = true
						break
					}
				}
			}

			if !ignore {
				methods = append(methods, method)
			}
		}

		s.Services[i].Methods = methods
	}

	return s
}
