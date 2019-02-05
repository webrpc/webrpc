package schema

import (
	"bytes"
	"strings"

	"github.com/davecgh/go-spew/spew"
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

	// switch t.Type {
	// case T_Invalid:
	// 	return "<invalid>"

	// case T_List:
	// 	// TODO: support [][]string
	// 	// TODO: support []<message> type..? yes
	// 	return "list, todo.."

	// case T_Map:
	// 	// TODO: support map<string,map<string,uint32>>
	// 	// TODO: support map<string,User>
	// 	return "map, todo.."

	// case T_Struct:
	// 	// TODO: ..
	// 	return "struct, todo.."

	// default:
	// 	// basic type
	// 	return t.Type.String()
	// }
}

func (t VarType) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBufferString(`"`)
	buf.WriteString(t.Type.String())
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
	s := t.expr

	// parse data type from string
	dataType, ok := DataTypeFromString[s]

	if !ok {
		// test for complex datatype
		if strings.HasPrefix(s, DataTypeToString[T_List]) {
			dataType = T_List
		} else if strings.HasPrefix(s, DataTypeToString[T_Map]) {
			dataType = T_Map
		}
	}

	// Set core data type
	t.Type = dataType

	// For complex types, keep parsing
	switch t.Type {
	case T_List:
		vt := &VarType{}
		err := parseListTypeString(s, vt)
		if err != nil {
			return err
		}
		// TODO: set t to tt..
		_ = vt
		spew.Dump(vt)
		return nil

	case T_Map:
		tt, err := parseMapTypeString(s)
		if err != nil {
			return err
		}
		// TODO: set t to tt..
		_ = tt
		spew.Dump(tt)
		return nil

	case T_Invalid:

		// TODO: for Message/Struct type, we can put a stub(?),
		// but need to do pass at the post-processing stage
		// if it maps to a messageRef

		// let's assume its a message ref then, until post-processor verifies

		// TODO: ..........
		return nil

	default:
		// basic type, we're done here
		return nil

	}
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

func parseListTypeString(s string, vt *VarType) error {
	if !strings.HasPrefix(s, DataTypeToString[T_List]) {
		return nil
	}

	// TODO: support [][]string

	// okay...... so, we need to unmarshal in one step..
	// and post-process in another, its clear now..
	// so we have all messages

	// TODO: support []<message>
	return nil
}

func parseMapTypeString(s string) (*VarType, error) {
	// TODO: support map<string,map<string,uint32>>
	// TODO: support map<string,User>
	return nil, nil
}

func parseStructTypeString(s string) (*VarType, error) {
	return nil, nil
}
