package schema

import (
	"bytes"
	"strings"

	"github.com/pkg/errors"
)

type VarType struct {
	expr string
	Type DataType

	Map    *VarMapType
	List   *VarListType
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
		return errors.Errorf("json error: type cannot be empty")
	}
	s := string(b) // string value will be wrapped in quotes

	// validate its a string value
	if s[0:1] != "\"" {
		return errors.Errorf("json error: string value is expected")
	}
	if s[len(s)-1:] != "\"" {
		return errors.Errorf("json error: string value is expected")
	}

	// trim string quotes from the json string
	s = s[1:]
	s = s[:len(s)-1]

	// set the expr from value
	t.expr = s

	return nil
}

func (t *VarType) Parse(schema *WebRPCSchema) error {
	if t.expr == "" {
		return errors.Errorf("parse error: type expr cannot be empty")
	}
	return parseVarTypeExpr(schema, t.expr, t)
}

type VarMapType struct {
	Key   DataType
	Value VarType
}

type VarListType struct {
	Elem VarType
}

type VarStructType struct {
	// Fields []VarType // redundant, its already in Message with more data..
	Name    string
	Message *Message
}

func parseVarTypeExpr(schema *WebRPCSchema, expr string, vt *VarType) error {
	if expr == "" {
		return nil
	}

	// parse data type from string
	dataType, ok := DataTypeFromString[expr]

	if !ok {
		// test for complex datatype
		if strings.HasPrefix(expr, DataTypeToString[T_List]) {
			dataType = T_List
		} else if strings.HasPrefix(expr, DataTypeToString[T_Map]) {
			dataType = T_Map
		}
	}

	// Set core data type
	vt.Type = dataType

	// For complex types, keep parsing
	switch vt.Type {
	case T_List:
		// vt := &VarType{}
		// err := parseListTypeString(s, vt)
		// if err != nil {
		// 	return err
		// }
		// // TODO: set vt to t..
		// spew.Dump(vt)

	case T_Map:
		// vt := &VarType{}
		// err := parseMapTypeString(s, vt)
		// if err != nil {
		// 	return err
		// }
		// // TODO: set vt to t..
		// spew.Dump(vt)

	case T_Invalid:

		// TODO: check schema.Messages list to ensure the name matches..
		// or return an error..
		// setup a ref.. etc.

	default:
		// basic type, we're done here
	}

	return nil
}

// func parseListTypeString(expr string, vt *VarType) error {
// 	// TODO: support [][]string
// 	// TODO: support []<message>
// 	return nil
// }

// func parseMapTypeString(expr string, vt *VarType) error {
// 	// TODO: support map<string,map<string,uint32>>
// 	// TODO: support map<string,User>
// 	return nil
// }

// func parseStructTypeString(expr string, vt *VarType) error {
// 	return nil
// }

func buildVarTypeExpr(vt *VarType) string {
	switch vt.Type {
	case T_Invalid:
		return "<invalid>"

	case T_List:
		// TODO: support [][]string
		// TODO: support []<message> type..? yes
		return "list, todo.."

	case T_Map:
		// TODO: support map<string,map<string,uint32>>
		// TODO: support map<string,User>
		return "map, todo.."

	case T_Struct:
		// TODO: ..
		return "struct, todo.."

	default:
		// basic type
		return vt.Type.String()
	}
}
