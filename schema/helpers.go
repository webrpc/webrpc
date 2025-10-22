package schema

import "strings"

// SchemaBigIntFieldsByType inspects the schema and returns a map of struct type names to a list of fields
// that are of type bigint (case insensitive), including nested structs and lists of bigints.
//
// This function is used as a template helper for adding bigint support in various codegen targets.
//
// NOTE: we also add support for "compat" mode / flag, used for old typescript generation that uses
// the names Args and Return for suffixes instead of Request and Resposne suffixes for implicit method types.
func SchemaBigIntFieldsByType(s *WebRPCSchema, optCompatMode ...bool) map[string][]any {
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

	// helper to collect bigint-related fields for a given struct type
	collectStruct := func(t *Type) []any {
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
			// list of nested struct containing bigint
			if f.Type.Type == T_List && f.Type.List != nil && f.Type.List.Elem.Type == T_Struct && f.Type.List.Elem.Struct != nil && structHasBigIntRecursive(f.Type.List.Elem.Struct.Type) {
				list = append(list, [2]string{string(f.Name), f.Type.List.Elem.Struct.Name + "[]"})
				continue
			}
		}
		return list
	}

	for _, t := range s.Types {
		if t.Kind != TypeKind_Struct {
			continue
		}
		list := collectStruct(t)
		if len(list) > 0 {
			out[t.Name] = list
		}
	}

	// Synthetic implicit req/resp types for non-succinct method inputs/outputs.
	// For a method Foo with Succinct=false (or unspecified), we have implicit FooResponse
	// comprised of its output arguments. If any output argument (directly or nested)
	// contains bigint, include FooResponse in the result. For nested structs, we emit
	// [fieldName, StructName] similar to real struct handling.
	// Likewise, implicit FooRequest (or FooArgs in compat mode) is synthesized from inputs.
	for _, svc := range s.Services {
		for _, m := range svc.Methods {
			if m == nil || m.Succinct { // skip succinct methods, they don't get FooResponse types
				continue
			}

			// Build synthetic request type from inputs
			requestName := m.Name
			if len(optCompatMode) > 0 && optCompatMode[0] {
				requestName += "Args"
			} else {
				requestName += "Request"
			}
			var reqList []any
			for _, arg := range m.Inputs {
				if arg == nil || arg.Type == nil {
					continue
				}
				vt := arg.Type
				if strings.EqualFold(vt.Expr, "bigint") { // scalar bigint
					reqList = append(reqList, string(arg.Name))
					continue
				}
				if vt.Type == T_List && vt.List != nil && strings.EqualFold(vt.List.Elem.Expr, "bigint") { // list of bigint
					reqList = append(reqList, string(arg.Name)+"[]")
					continue
				}
				if vt.Type == T_Struct && vt.Struct != nil && structHasBigIntRecursive(vt.Struct.Type) { // nested struct
					reqList = append(reqList, [2]string{string(arg.Name), vt.Struct.Name})
					continue
				}
				if vt.Type == T_List && vt.List != nil && vt.List.Elem.Type == T_Struct && vt.List.Elem.Struct != nil && structHasBigIntRecursive(vt.List.Elem.Struct.Type) { // list of nested struct
					reqList = append(reqList, [2]string{string(arg.Name), vt.List.Elem.Struct.Name + "[]"})
					continue
				}
			}
			if len(reqList) > 0 {
				out[requestName] = reqList
			}

			responseName := m.Name
			if len(optCompatMode) > 0 && optCompatMode[0] {
				responseName += "Return"
			} else {
				responseName += "Response"
			}
			var respList []any
			for _, arg := range m.Outputs {
				if arg == nil || arg.Type == nil {
					continue
				}
				vt := arg.Type
				// direct scalar bigint
				if strings.EqualFold(vt.Expr, "bigint") {
					respList = append(respList, string(arg.Name))
					continue
				}
				// list of bigint
				if vt.Type == T_List && vt.List != nil && strings.EqualFold(vt.List.Elem.Expr, "bigint") {
					respList = append(respList, string(arg.Name)+"[]")
					continue
				}
				// nested struct
				if vt.Type == T_Struct && vt.Struct != nil && structHasBigIntRecursive(vt.Struct.Type) {
					respList = append(respList, [2]string{string(arg.Name), vt.Struct.Name})
					continue
				}
				// list of nested struct
				if vt.Type == T_List && vt.List != nil && vt.List.Elem.Type == T_Struct && vt.List.Elem.Struct != nil && structHasBigIntRecursive(vt.List.Elem.Struct.Type) {
					respList = append(respList, [2]string{string(arg.Name), vt.List.Elem.Struct.Name + "[]"})
					continue
				}
			}
			if len(respList) > 0 {
				out[responseName] = respList
			}
		}
	}
	return out
}
