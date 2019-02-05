package schema

import "fmt"

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
	s := string(b)
	fmt.Println("woohoo", s)

	// start, cannot be len 0, or error

	// first, lets get datatype

	// second, if complex type, parse those (ie. map/list)

	// note: for structs, what do we do..? cuz its a ref
	// to another type.., we can have a var that is a Message
	// we can check list of Messages..
	//
	// .. we cant do that at unmarshal step, would need to
	// be post-processed to get the ref from the MessageName

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
