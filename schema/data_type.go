package schema

import (
	"bytes"
	"encoding/json"
)

type DataType int

const (
	T_Invalid DataType = iota

	T_Null
	T_Any
	T_Byte
	T_Bool

	T_UInt8
	T_Uint16
	T_Uint32
	T_Uint64

	T_Int8
	T_Int16
	T_Int32
	T_Int64

	T_Float32
	T_Float64

	T_String

	T_List
	T_Map

	T_Struct // aka a message
	// rename to.. T_Message ?
)

var DataTypeToString = map[DataType]string{
	T_Null: "null",
	T_Any:  "any",
	T_Byte: "byte",
	T_Bool: "bool",

	T_UInt8:  "uint8",
	T_Uint16: "uint16",
	T_Uint32: "uint32",
	T_Uint64: "uint64",

	T_Int8:  "int8",
	T_Int16: "int16",
	T_Int32: "int32",
	T_Int64: "int64",

	T_Float32: "float32",
	T_Float64: "float64",

	T_String: "string",

	T_Map:  "map",
	T_List: "[]",
}

var DataTypeFromString = map[string]DataType{
	"null": T_Null,
	"any":  T_Any,
	"byte": T_Byte,
	"bool": T_Bool,

	"uint8":  T_UInt8,
	"uint16": T_Uint16,
	"uint32": T_Uint32,
	"uint64": T_Uint64,

	"int8":  T_Int8,
	"int16": T_Int16,
	"int32": T_Int32,
	"int64": T_Int64,

	"float32": T_Float32,
	"float64": T_Float64,

	"string": T_String,

	"map": T_Map,
	"[]":  T_List,
}

func (t DataType) String() string {
	return DataTypeToString[t]
}

func (t DataType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(DataTypeToString[t])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (t *DataType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*t = DataTypeFromString[j]
	return nil
}

func (t *DataType) Validate() error {
	// return errors.Errorf("invalid data DataType %s", t)
	return nil
}
