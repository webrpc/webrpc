package schema

import (
	"bytes"
	"encoding/json"
)

type CoreType int

const (
	T_Unknown CoreType = iota

	T_Null
	T_Any
	T_Byte
	T_Bool

	T_Uint
	T_Uint8
	T_Uint16
	T_Uint32
	T_Uint64

	T_Int
	T_Int8
	T_Int16
	T_Int32
	T_Int64

	T_Float32
	T_Float64

	T_String

	T_Timestamp

	T_List
	T_Map

	T_Primitive
	T_Struct // aka, a reference to our own webrpc struct type
	T_Alias  // aka, a reference to our own webrpc alias type
)

var DataTypeToString = map[CoreType]string{
	T_Null: "null",
	T_Any:  "any",
	T_Byte: "byte",
	T_Bool: "bool",

	T_Uint:   "uint",
	T_Uint8:  "uint8",
	T_Uint16: "uint16",
	T_Uint32: "uint32",
	T_Uint64: "uint64",

	T_Int:   "int",
	T_Int8:  "int8",
	T_Int16: "int16",
	T_Int32: "int32",
	T_Int64: "int64",

	T_Float32: "float32",
	T_Float64: "float64",

	T_String: "string",

	T_Timestamp: "timestamp",

	T_Map:  "map",
	T_List: "[]",
}

var CoreTypeFromString = map[string]CoreType{
	"null": T_Null,
	"any":  T_Any,
	"byte": T_Byte,
	"bool": T_Bool,

	"uint":   T_Uint,
	"uint8":  T_Uint8,
	"uint16": T_Uint16,
	"uint32": T_Uint32,
	"uint64": T_Uint64,

	"int":   T_Int,
	"int8":  T_Int8,
	"int16": T_Int16,
	"int32": T_Int32,
	"int64": T_Int64,

	"float32": T_Float32,
	"float64": T_Float64,

	"string": T_String,

	"timestamp": T_Timestamp,

	"map": T_Map,
	"[]":  T_List,
}

func (t CoreType) String() string {
	return DataTypeToString[t]
}

func (t CoreType) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBufferString(`"`)
	buf.WriteString(DataTypeToString[t])
	buf.WriteString(`"`)
	return buf.Bytes(), nil
}

func (t *CoreType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*t = CoreTypeFromString[j]
	return nil
}
