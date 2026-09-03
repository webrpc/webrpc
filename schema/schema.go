package schema

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
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
	BasePath      string `json:"basepath"`

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

	if len(s.Services) > 0 && s.SchemaVersion == "" {
		return fmt.Errorf("schema error: version is required when services are defined")
	}

	if s.BasePath != "" {
		if strings.ContainsAny(s.BasePath, " \t\r\n") {
			return fmt.Errorf("schema error: basepath must not contain whitespace")
		}
		if strings.ContainsAny(s.BasePath, "?#") {
			return fmt.Errorf("schema error: basepath must not include query or fragment")
		}
		if strings.Contains(s.BasePath, "://") {
			return fmt.Errorf("schema error: basepath must be a path, not a URL")
		}
		pathToParse := s.BasePath
		if !strings.HasPrefix(pathToParse, "/") {
			pathToParse = "/" + pathToParse
		}
		parsed, err := url.ParseRequestURI(pathToParse)
		if err != nil {
			return fmt.Errorf("schema error: basepath is invalid: %w", err)
		}
		if parsed.Scheme != "" || parsed.Host != "" || parsed.Opaque != "" || parsed.User != nil {
			return fmt.Errorf("schema error: basepath must be a path, not a URL")
		}
		if !strings.HasPrefix(s.BasePath, "/") {
			s.BasePath = "/" + s.BasePath
		}
		if !strings.HasSuffix(s.BasePath, "/") {
			s.BasePath += "/"
		}
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

	if err := s.validateFileTypeUsage(); err != nil {
		return err
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

// MethodTreeShake removes unused types from the schema based on the methods which
// are *not used* in the output. We look up services which are excluded, and then
// omit types which are referenced by those methods only, and not other methods
// from other services.
func MethodTreeShake(s *WebRPCSchema, services []string) *WebRPCSchema {
	if s == nil {
		return nil
	}

	// Build set of included services from input list
	include := map[string]struct{}{}
	for _, name := range services {
		include[name] = struct{}{}
	}

	// Collect type names referenced by included vs excluded service methods
	usedIncluded := map[string]bool{}
	usedExcluded := map[string]bool{}

	// Helper: recursively walk VarType and mark used type names, including nested struct fields and map/list elements.
	var walkVarType func(vt *VarType, mark map[string]bool)
	walkVarType = func(vt *VarType, mark map[string]bool) {
		if vt == nil {
			return
		}

		// List element
		if vt.List != nil {
			walkVarType(vt.List.Elem, mark)
		}

		// Map key/value
		if vt.Map != nil {
			walkVarType(vt.Map.Key, mark)
			walkVarType(vt.Map.Value, mark)
		}

		// Struct type: mark and traverse its fields
		if vt.Struct != nil && vt.Struct.Type != nil {
			t := vt.Struct.Type
			if t.Name != "" {
				mark[t.Name] = true
			}
			// Traverse struct fields to include transitive type usage
			for _, f := range t.Fields {
				walkVarType(f.Type, mark)
			}
		}

		// Enum type: mark usage
		if vt.Enum != nil && vt.Enum.Type != nil {
			t := vt.Enum.Type
			if t.Name != "" {
				mark[t.Name] = true
			}
		}
	}

	// Traverse all methods of services; split by included/excluded.
	for _, srv := range s.Services {
		_, isIncluded := include[srv.Name]
		mark := usedExcluded
		if isIncluded {
			mark = usedIncluded
		}
		for _, m := range srv.Methods {
			for _, in := range m.Inputs {
				walkVarType(in.Type, mark)
			}
			for _, out := range m.Outputs {
				walkVarType(out.Type, mark)
			}
		}
	}

	// Keep only types that are transitively referenced by included service methods.
	var kept []*Type
	for _, t := range s.Types {
		if usedIncluded[t.Name] {
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

// annotationsMatch reports whether ann contains at least one annotation
// matching the given filter, where filter is a map of annotation name to
// value (empty value means match by annotation type only, regardless of value).
func annotationsMatch(ann Annotations, filter map[string]string) bool {
	for name, value := range filter {
		a, ok := ann[name]
		if !ok {
			continue
		}

		trimmedValue := strings.Trim(value, " ")
		if trimmedValue == "" {
			// match annotation types only
			if name == a.AnnotationType {
				return true
			}
		} else {
			// match @key:value
			if name == a.AnnotationType && value == a.Value {
				return true
			}
		}
	}

	return false
}

// referencingTypeOrMethodNames returns the names of types (aliases, struct/enum
// fields) and service methods that reference the type named targetName,
// excluding excludeType itself.
func referencingTypeOrMethodNames(s *WebRPCSchema, targetName string, excludeType *Type) []string {
	var checkVarType func(vt *VarType) bool
	checkVarType = func(vt *VarType) bool {
		if vt == nil {
			return false
		}
		if vt.List != nil && checkVarType(vt.List.Elem) {
			return true
		}
		if vt.Map != nil && (checkVarType(vt.Map.Key) || checkVarType(vt.Map.Value)) {
			return true
		}
		if vt.Struct != nil && vt.Struct.Type != nil && vt.Struct.Type.Name == targetName {
			return true
		}
		if vt.Enum != nil && vt.Enum.Type != nil && vt.Enum.Type.Name == targetName {
			return true
		}
		return false
	}

	referrers := map[string]struct{}{}

	for _, t := range s.Types {
		if t == excludeType {
			continue
		}
		if t.Type != nil && checkVarType(t.Type) {
			referrers[t.Name] = struct{}{}
		}
		for _, f := range t.Fields {
			if checkVarType(f.Type) {
				referrers[t.Name] = struct{}{}
			}
		}
	}

	for _, srv := range s.Services {
		for _, m := range srv.Methods {
			for _, in := range m.Inputs {
				if checkVarType(in.Type) {
					referrers[srv.Name+"."+m.Name+"()"] = struct{}{}
				}
			}
			for _, out := range m.Outputs {
				if checkVarType(out.Type) {
					referrers[srv.Name+"."+m.Name+"()"] = struct{}{}
				}
			}
		}
	}

	names := make([]string, 0, len(referrers))
	for name := range referrers {
		names = append(names, name)
	}
	sort.Strings(names)

	return names
}

// filterTypesWithAnnotations keeps types in s.Types for which keep(t) returns true.
// Removing a type that is still referenced elsewhere in the schema is a hard error.
func filterTypesWithAnnotations(s *WebRPCSchema, keep func(t *Type) bool) (*WebRPCSchema, error) {
	if s == nil {
		return nil, nil
	}

	var kept []*Type
	for _, t := range s.Types {
		if keep(t) {
			kept = append(kept, t)
			continue
		}

		if refs := referencingTypeOrMethodNames(s, t.Name, t); len(refs) > 0 {
			return nil, fmt.Errorf("schema error: cannot ignore/exclude type '%s' via annotation, it is still referenced by: %s", t.Name, strings.Join(refs, ", "))
		}
	}
	s.Types = kept

	return s, nil
}

// MatchTypesWithAnnotations keeps only the types matching one of the given annotations.
// It is a hard error to exclude a type that is still referenced elsewhere in the schema.
func MatchTypesWithAnnotations(s *WebRPCSchema, annotations map[string]string) (*WebRPCSchema, error) {
	return filterTypesWithAnnotations(s, func(t *Type) bool {
		return annotationsMatch(t.Annotations, annotations)
	})
}

// IgnoreTypesWithAnnotations removes types matching one of the given annotations.
// It is a hard error to ignore a type that is still referenced elsewhere in the schema.
func IgnoreTypesWithAnnotations(s *WebRPCSchema, annotations map[string]string) (*WebRPCSchema, error) {
	return filterTypesWithAnnotations(s, func(t *Type) bool {
		return !annotationsMatch(t.Annotations, annotations)
	})
}

// MatchTypeFieldsWithAnnotations keeps, on every type, only the fields matching
// one of the given annotations.
func MatchTypeFieldsWithAnnotations(s *WebRPCSchema, annotations map[string]string) *WebRPCSchema {
	if s == nil {
		return nil
	}

	for _, t := range s.Types {
		var fields []*TypeField
		for _, f := range t.Fields {
			if annotationsMatch(f.Annotations, annotations) {
				fields = append(fields, f)
			}
		}
		t.Fields = fields
	}

	return s
}

// IgnoreTypeFieldsWithAnnotations removes, from every type, the fields matching
// one of the given annotations.
func IgnoreTypeFieldsWithAnnotations(s *WebRPCSchema, annotations map[string]string) *WebRPCSchema {
	if s == nil {
		return nil
	}

	for _, t := range s.Types {
		var fields []*TypeField
		for _, f := range t.Fields {
			if !annotationsMatch(f.Annotations, annotations) {
				fields = append(fields, f)
			}
		}
		t.Fields = fields
	}

	return s
}
