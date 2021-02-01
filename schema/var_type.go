package schema

import (
	"bytes"
	"fmt"
	"strings"
)

type VarType struct {
	expr string
	Type CoreType

	List   *VarListType
	Map    *VarMapType
	Struct *VarStructType
}

func (t *VarType) String() string {
	return t.expr
}

func (t *VarType) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBufferString(`"`)
	buf.WriteString(t.String())
	buf.WriteString(`"`)
	return buf.Bytes(), nil
}

func (t *VarType) UnmarshalJSON(b []byte) error {
	if len(b) <= 2 {
		return fmt.Errorf("json error: type cannot be empty")
	}
	s := string(b) // string value will be wrapped in quotes

	// validate its a string value
	if s[0:1] != "\"" {
		return fmt.Errorf("json error: string value is expected")
	}
	if s[len(s)-1:] != "\"" {
		return fmt.Errorf("json error: string value is expected")
	}

	// trim string quotes from the json string
	s = s[1:]
	s = s[:len(s)-1]

	// set the expr from value
	t.expr = s

	return nil
}

func (t *VarType) Parse(schema *WebRPCSchema) error {
	if t == nil || t.expr == "" {
		return fmt.Errorf("schema error: type expr cannot be empty")
	}
	err := ParseVarTypeExpr(schema, t.expr, t)
	if err != nil {
		return err
	}
	t.expr = buildVarTypeExpr(t, "")
	return nil
}

type VarListType struct {
	Elem *VarType
}

type VarMapType struct {
	Key   CoreType // see, VarMapKeyCoreTypes -- only T_String or T_XintX supported
	Value *VarType
}

type VarStructType struct {
	Name string
	Type *Type
}

func ParseVarTypeExpr(schema *WebRPCSchema, expr string, vt *VarType) error {
	if expr == "" {
		return nil
	}
	vt.expr = expr

	// parse data type from string
	dataType, ok := CoreTypeFromString[expr]

	if !ok {
		// test for complex datatype
		if isListExpr(expr) {
			dataType = T_List
		} else if isMapExpr(expr) {
			dataType = T_Map
		}
	}

	// Set core data type
	vt.Type = dataType

	// For complex types, keep parsing
	switch vt.Type {
	case T_List:
		// create sub-type object for list element
		vt.List = &VarListType{Elem: &VarType{}}

		// shift expr, and keep parsing
		expr = strings.TrimPrefix(expr, DataTypeToString[T_List])
		err := ParseVarTypeExpr(schema, expr, vt.List.Elem)
		if err != nil {
			return err
		}

	case T_Map:
		// parse map expr
		key, value, err := parseMapExpr(expr)
		if err != nil {
			return err
		}

		keyDataType, ok := CoreTypeFromString[key]
		if !ok {
			return fmt.Errorf("schema error: invalid map key type '%s' for expr '%s'", key, expr)
		}

		// create sub-type object for map
		vt.Map = &VarMapType{Key: keyDataType, Value: &VarType{}}

		// shift expr and keep parsing
		expr = value
		err = ParseVarTypeExpr(schema, expr, vt.Map.Value)
		if err != nil {
			return err
		}

	case T_Unknown:
		structExpr := expr
		typ, ok := getType(schema, structExpr)
		if !ok || typ == nil {
			return fmt.Errorf("schema error: invalid struct type '%s'", structExpr)
		}

		vt.Type = T_Struct
		vt.Struct = &VarStructType{Name: structExpr, Type: typ}

	default:
		// basic type, we're done here
	}

	return nil
}

func parseMapExpr(expr string) (string, string, error) {
	if !isMapExpr(expr) {
		return "", "", fmt.Errorf("schema error: invalid map expr for '%s'", expr)
	}

	mapKeyword := DataTypeToString[T_Map]
	expr = expr[len(mapKeyword):]

	if expr[0:1] != "<" {
		return "", "", fmt.Errorf("schema error: invalid map syntax for '%s'", expr)
	}
	if expr[len(expr)-1:] != ">" {
		return "", "", fmt.Errorf("schema error: invalid map syntax for '%s'", expr)
	}
	expr = expr[1 : len(expr)-1]

	p := strings.Index(expr, ",")
	if p < 0 {
		return "", "", fmt.Errorf("schema error: invalid map syntax for '%s'", expr)
	}

	key := expr[0:p]
	value := expr[p+1:]

	if !isValidVarKeyType(key) {
		return "", "", fmt.Errorf("schema error: invalid map key '%s' for '%s'", key, expr)
	}

	return key, value, nil
}

func buildVarTypeExpr(vt *VarType, expr string) string {
	switch vt.Type {
	case T_Unknown:
		return "<unknown>"

	case T_List:
		expr += "[]" + buildVarTypeExpr(vt.List.Elem, expr)
		return expr

	case T_Map:
		expr += fmt.Sprintf("map<%s,%s>", vt.Map.Key, buildVarTypeExpr(vt.Map.Value, ""))
		return expr

	case T_Struct:
		expr += vt.Struct.Name
		return expr

	default:
		// basic type
		expr += vt.Type.String()
		return expr
	}
}

func isListExpr(expr string) bool {
	listTest := DataTypeToString[T_List]
	return strings.HasPrefix(expr, listTest)
}

func isMapExpr(expr string) bool {
	mapTest := DataTypeToString[T_Map] + "<"
	return strings.HasPrefix(expr, mapTest)
}

func getType(schema *WebRPCSchema, structExpr string) (*Type, bool) {
	for _, typ := range schema.Types {
		if structExpr == string(typ.Name) {
			return typ, true
		}
	}
	return nil, false
}

var VarKeyCoreTypes = []CoreType{
	T_String, T_Uint, T_Uint8, T_Uint16, T_Uint32, T_Uint64, T_Int, T_Int8, T_Int16, T_Int32, T_Int64,
}

var VarIntegerCoreTypes = []CoreType{
	T_Uint, T_Uint8, T_Uint16, T_Uint32, T_Uint64, T_Int, T_Int8, T_Int16, T_Int32, T_Int64,
}

func isValidVarKeyType(s string) bool {
	return isValidVarType(s, VarKeyCoreTypes)
}

func isValidVarType(s string, allowedList []CoreType) bool {
	dt, ok := CoreTypeFromString[s]
	if !ok {
		return false
	}
	for _, t := range allowedList {
		if dt == t {
			return true
		}
	}
	return false
}
