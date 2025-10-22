package schema

import "strings"

// SchemaBigIntFieldsByType inspects the schema and returns a map of struct type names to a list of fields
// that are of type bigint (case insensitive), including nested structs and lists of bigints.
//
// This function is used as a template helper for adding bigint support in various codegen targets.
func SchemaBigIntFieldsByType(s *WebRPCSchema) map[string][]any {
	out := make(map[string][]any)
	// cache for recursive detection to avoid cycles
	visited := make(map[string]bool)
	memo := make(map[string]bool)

	var structHasBigIntRecursive func(t *Type) bool
	structHasBigIntRecursive = func(t *Type) bool {
		if t == nil || t.Kind != TypeKind_Struct {
			return false
		}
		name := t.Name
		if v, ok := memo[name]; ok {
			return v
		}
		if visited[name] { // cycle guard
			return false
		}
		visited[name] = true
		defer func() { visited[name] = false }()

		for _, f := range t.Fields {
			if f.Type == nil {
				continue
			}
			// direct scalar bigint
			if strings.EqualFold(f.Type.Expr, "bigint") {
				memo[name] = true
				return true
			}
			// list of bigint
			if f.Type.Type == T_List && f.Type.List != nil && strings.EqualFold(f.Type.List.Elem.Expr, "bigint") {
				memo[name] = true
				return true
			}
			// nested struct (struct types are parsed as T_Struct)
			if f.Type.Type == T_Struct && f.Type.Struct != nil && structHasBigIntRecursive(f.Type.Struct.Type) {
				memo[name] = true
				return true
			}
			// list of nested struct
			if f.Type.Type == T_List && f.Type.List != nil && f.Type.List.Elem.Type == T_Struct && f.Type.List.Elem.Struct != nil && structHasBigIntRecursive(f.Type.List.Elem.Struct.Type) {
				memo[name] = true
				return true
			}
		}
		memo[name] = false
		return false
	}

	for _, t := range s.Types {
		if t.Kind != TypeKind_Struct {
			continue
		}
		var list []any
		for _, f := range t.Fields {
			if f.Type == nil {
				continue
			}
			// scalar bigint
			if strings.EqualFold(f.Type.Expr, "bigint") {
				list = append(list, string(f.Name))
				continue
			}
			// list of bigint
			if f.Type.Type == T_List && f.Type.List != nil && strings.EqualFold(f.Type.List.Elem.Expr, "bigint") {
				list = append(list, string(f.Name)+"[]")
				continue
			}
			// nested struct containing (directly or transitively) bigint
			if f.Type.Type == T_Struct && f.Type.Struct != nil && structHasBigIntRecursive(f.Type.Struct.Type) {
				list = append(list, [2]string{string(f.Name), f.Type.Struct.Name})
				continue
			}
			// list of nested struct (emit field name, and nested type suffixed with [] to indicate array)
			if f.Type.Type == T_List && f.Type.List != nil && f.Type.List.Elem.Type == T_Struct && f.Type.List.Elem.Struct != nil && structHasBigIntRecursive(f.Type.List.Elem.Struct.Type) {
				list = append(list, [2]string{string(f.Name), f.Type.List.Elem.Struct.Name + "[]"})
				continue
			}
		}
		if len(list) > 0 {
			out[t.Name] = list
		}
	}
	return out
}
