package schema

import "strings"

// FilePart describes one file-carrying slot of a method's multipart upload
// request. Each part maps to a `file` or `[]file` method input argument, or
// to a top-level `file` / `[]file` field of a struct input argument.
//
// This type is used as a template helper for adding file upload support in
// various codegen targets.
type FilePart struct {
	PartName   string // multipart part name: the argument name, or the struct field's json name
	ArgIndex   int    // index of the method input argument carrying the file
	ArgName    string // name of that input argument
	FieldName  string // Go field name inside the input struct; "" when the argument itself is the file
	StructName string // input struct type name; "" when the argument itself is the file
	Repeated   bool   // true for []file
}

// MethodHasFileUpload returns true if the method takes a `file` or `[]file`
// input argument, or a struct input argument with top-level file fields, and
// therefore uses a multipart/form-data request body.
func MethodHasFileUpload(m *Method) bool {
	return len(MethodFileParts(m)) > 0
}

// MethodHasFileDownload returns true if the method returns a `file` as its
// single output, and therefore responds with a raw file body.
func MethodHasFileDownload(m *Method) bool {
	if m == nil || len(m.Outputs) != 1 || m.Outputs[0] == nil {
		return false
	}
	return m.Outputs[0].Type != nil && m.Outputs[0].Type.Type == T_File
}

// MethodFileParts returns the multipart file parts of a method's upload
// request, in the order the client is expected to send them (after the
// leading "json" part).
func MethodFileParts(m *Method) []*FilePart {
	if m == nil || m.StreamInput || m.StreamOutput {
		return nil
	}

	var parts []*FilePart
	for i, in := range m.Inputs {
		if in == nil || in.Type == nil {
			continue
		}
		if repeated, ok := fileArgShape(in.Type); ok {
			parts = append(parts, &FilePart{
				PartName: in.Name,
				ArgIndex: i,
				ArgName:  in.Name,
				Repeated: repeated,
			})
			continue
		}
		if in.Type.Type == T_Struct && in.Type.Struct != nil && in.Type.Struct.Type != nil {
			for _, field := range in.Type.Struct.Type.Fields {
				repeated, ok := fileArgShape(field.Type)
				if !ok {
					continue
				}
				parts = append(parts, &FilePart{
					PartName:   fieldJSONName(field),
					ArgIndex:   i,
					ArgName:    in.Name,
					FieldName:  fieldGoName(field),
					StructName: in.Type.Struct.Name,
					Repeated:   repeated,
				})
			}
		}
	}
	return parts
}

// fileArgShape reports whether the type is exactly `file` (false, true) or
// `[]file` (true, true).
func fileArgShape(vt *VarType) (repeated bool, ok bool) {
	if vt == nil {
		return false, false
	}
	if vt.Type == T_File {
		return false, true
	}
	if vt.Type == T_List && vt.List != nil && vt.List.Elem != nil && vt.List.Elem.Type == T_File {
		return true, true
	}
	return false, false
}

// fieldJSONName returns the field's name on the wire, honoring the `json`
// and `go.tag.json` field meta, mirroring the Go generator's field template.
func fieldJSONName(field *TypeField) string {
	name := field.Name
	for _, meta := range field.Meta {
		if v, ok := meta["json"].(string); ok {
			name = v
		}
		if v, ok := meta["go.tag.json"].(string); ok {
			name = v
		}
	}
	if i := strings.Index(name, ","); i >= 0 {
		name = name[:i]
	}
	return name
}

// fieldGoName returns the field's generated Go name, honoring the
// `go.field.name` field meta, mirroring the Go generator's field template.
func fieldGoName(field *TypeField) string {
	name := field.Name
	if name != "" {
		name = strings.ToUpper(name[:1]) + name[1:]
	}
	for _, meta := range field.Meta {
		if v, ok := meta["go.field.name"].(string); ok {
			name = v
		}
	}
	return name
}
