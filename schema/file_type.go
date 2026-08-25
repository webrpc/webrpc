package schema

import (
	"fmt"
)

// validateFileTypeUsage enforces the placement rules of the `file` core type.
//
// `file` and `[]file` are allowed only as:
//   - the type of a method input argument, or
//   - the type of a top-level field of a struct that is used as a method input.
//
// A method may return at most one `file` output (a download method). Any
// remaining outputs are metadata: they are JSON-encoded into the
// Webrpc-Response header instead of the body, which carries the raw file.
// Everywhere else — nested structs, maps, aliases, list-of-list, streaming
// methods, or file-carrying structs used outside of a method input — the
// file type is forbidden.
func (s *WebRPCSchema) validateFileTypeUsage() error {
	// Struct types carrying top-level file fields; only valid as method inputs.
	fileStructs := map[string]bool{}
	for _, t := range s.Types {
		if t.Kind != TypeKind_Struct {
			continue
		}
		for _, field := range t.Fields {
			if isFileArgType(field.Type) {
				fileStructs[t.Name] = true
				break
			}
		}
	}

	for _, t := range s.Types {
		switch t.Kind {
		case TypeKind_Alias:
			if s.varTypeHasFile(t.Type, nil) {
				return fmt.Errorf("schema error: alias '%s' cannot be declared over the file type", t.Name)
			}
		case TypeKind_Struct:
			for _, field := range t.Fields {
				if isFileArgType(field.Type) {
					// Direct file / []file field; whether the struct itself is
					// used correctly is validated at the method level below.
					continue
				}
				if s.varTypeHasFile(field.Type, nil) {
					return fmt.Errorf("schema error: field '%s' of struct '%s' cannot carry the file type: file and []file are only allowed as method input arguments or as top-level fields of a method input struct", field.Name, t.Name)
				}
			}
		}
	}

	for _, svc := range s.Services {
		for _, m := range svc.Methods {
			usesFile := false

			for _, in := range m.Inputs {
				switch {
				case isFileArgType(in.Type):
					usesFile = true
				case in.Type.Type == T_Struct && fileStructs[in.Type.Struct.Name]:
					usesFile = true
				case s.varTypeHasFile(in.Type, nil):
					return fmt.Errorf("schema error: input '%s' of method '%s' in service '%s' has invalid file placement: file and []file are only allowed as a direct input argument or as a top-level field of an input struct", in.Name, m.Name, svc.Name)
				}
			}

			fileOutputs := 0
			for _, out := range m.Outputs {
				if out.Type.Type == T_File {
					usesFile = true
					fileOutputs++
					if fileOutputs > 1 {
						return fmt.Errorf("schema error: method '%s' in service '%s' declares more than one file output: a method may return at most one file", m.Name, svc.Name)
					}
					continue
				}
				if s.varTypeHasFile(out.Type, nil) {
					return fmt.Errorf("schema error: output '%s' of method '%s' in service '%s' has invalid file placement: only a single 'file' output is allowed", out.Name, m.Name, svc.Name)
				}
			}

			if usesFile && (m.StreamInput || m.StreamOutput) {
				return fmt.Errorf("schema error: method '%s' in service '%s' cannot use the file type: streaming methods do not support file inputs or outputs", m.Name, svc.Name)
			}
		}
	}

	return nil
}

// isFileArgType reports whether the type is exactly `file` or `[]file`, the
// only two shapes allowed as input arguments and input struct fields.
func isFileArgType(vt *VarType) bool {
	if vt == nil {
		return false
	}
	if vt.Type == T_File {
		return true
	}
	return vt.Type == T_List && vt.List != nil && vt.List.Elem != nil && vt.List.Elem.Type == T_File
}

// varTypeHasFile reports whether the type transitively references the file
// type anywhere — through lists, maps, struct fields, or aliases.
func (s *WebRPCSchema) varTypeHasFile(vt *VarType, visited map[*Type]bool) bool {
	if vt == nil {
		return false
	}
	if vt.Type == T_File {
		return true
	}
	if visited == nil {
		visited = map[*Type]bool{}
	}

	// An alias' underlying VarType carries only its expr; parse it on the fly.
	if vt.Type == T_Unknown && vt.Expr != "" {
		parsed := &VarType{}
		if err := ParseVarTypeExpr(s, vt.Expr, parsed); err != nil {
			return false
		}
		vt = parsed
		if vt.Type == T_File {
			return true
		}
	}

	switch vt.Type {
	case T_List:
		if vt.List != nil {
			return s.varTypeHasFile(vt.List.Elem, visited)
		}
	case T_Map:
		if vt.Map != nil {
			return s.varTypeHasFile(vt.Map.Key, visited) || s.varTypeHasFile(vt.Map.Value, visited)
		}
	case T_Struct:
		if vt.Struct != nil && vt.Struct.Type != nil && !visited[vt.Struct.Type] {
			visited[vt.Struct.Type] = true
			for _, field := range vt.Struct.Type.Fields {
				if s.varTypeHasFile(field.Type, visited) {
					return true
				}
			}
		}
	case T_Alias:
		if vt.Alias != nil && vt.Alias.Type != nil && !visited[vt.Alias.Type] {
			visited[vt.Alias.Type] = true
			return s.varTypeHasFile(vt.Alias.Type.Type, visited)
		}
	}

	return false
}
