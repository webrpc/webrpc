package schema

import (
	"strings"

	"github.com/pkg/errors"
)

type VarType struct {
	Type DataType

	Map    *VarMapType
	List   *VarListType
	Struct *VarStructType
}

func (t VarType) String() string {
	return ""
}

func (t VarType) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (t *VarType) UnmarshalJSON(b []byte) error {
	if len(b) <= 2 {
		return errors.Errorf("type cannot be empty")
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

	// if dataType == T_Invalid {
	// 	return errors.Errorf("unknown type %s", s)
	// }

	// Set core data type
	t.Type = dataType

	// For complex types, keep parsing
	switch t.Type {
	case T_List:
		// TODO: support [][]string
		// TODO: support []<message> type..? yes

	case T_Map:
		// TODO: support map<string,map<string,uint32>>
		// TODO: support map<string,User>

	case T_Invalid:
		// let's assume its a message ref then, until post-processor verifies

	default:
		// basic type, we're done here
		return nil

	}

	// TODO: for Message/Struct type, we can put a stub(?),
	// but need to do pass at the post-processing stage
	// if it maps to a messageRef

	return nil
}

type VarMapType struct {
	Key   DataType
	Value VarType
}

type VarListType struct {
	Elem VarType
}

type VarStructType struct {
	// Fields []VarType
	Name    string
	Message *Message
}

// func testListPrefix(s string) bool {
// 	if strings.HasPrefix(s, DataTypeToString[T_List])
// }
