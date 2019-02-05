package schema

import (
	"bytes"
	"strings"

	"github.com/pkg/errors"
)

type VarType struct {
	expr string
	Type DataType

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

type VarListType struct {
	Elem *VarType
}

type VarMapType struct {
	Key   DataType // TODO: in fact, key can only be reduced set.. number or string, create VarMapKeyType as subset of DataType
	Value *VarType
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
		// create sub-type object for list element
		vt.List = &VarListType{}

		// shift expr, and keep parsing
		expr = strings.TrimPrefix(expr, DataTypeToString[T_List])
		err := parseVarTypeExpr(schema, expr, vt.List.Elem)
		if err != nil {
			return err
		}

	case T_Map:
		// parse map expr
		key, value, err := parseMapExpr(expr)
		if err != nil {
			return err
		}

		keyDataType, ok := DataTypeFromString[key]
		if !ok {
			return errors.Errorf("parse error: invalid map key type %s for expr %s", key, expr)
		}

		// create sub-type object for map
		vt.Map = &VarMapType{Key: keyDataType}

		// shift expr and keep parsing
		expr = value
		err = parseVarTypeExpr(schema, expr, vt.Map.Value)

	case T_Invalid:
		panic("TODO")
		// TODO: check schema.Messages list to ensure the name matches..
		// or return an error..
		// setup a ref.. etc.

	default:
		// basic type, we're done here
	}

	return nil
}

func parseMapExpr(expr string) (string, string, error) {
	mapKeyword := DataTypeToString[T_Map]

	if !strings.HasPrefix(expr, mapKeyword) {
		return "", "", errors.Errorf("parse error: invalud map expr for %s", expr)
	}

	expr = expr[len(mapKeyword):]

	if expr[0:1] != "<" {
		return "", "", errors.Errorf("parse error: invalid map syntax for %s", expr)
	}
	if expr[len(expr)-1:] != ">" {
		return "", "", errors.Errorf("parse error: invalid map syntax for %s", expr)
	}
	expr = expr[1 : len(expr)-2]

	p := strings.Index(",", expr)
	if p < 0 {
		return "", "", errors.Errorf("parse error: invalid map syntax for %s", expr)
	}

	key := expr[0 : p-1]
	value := expr[p:]

	return key, value, nil
}

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
