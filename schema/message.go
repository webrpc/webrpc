package schema

type Message struct {
	Name   VarName        `json:"name"`
	Type   MessageType    `json:"type"`
	Fields []MessageField `json:"fields"`
}

type MessageType string // TODO: must be `enum` or `struct`

type MessageField struct {
	Name     VarName           `json:"name"`
	Type     DataType          `json:"type"`
	Optional bool              `json:"optional"`
	Value    string            `json:"value"` // used by enums
	Tags     []MessageFieldTag `json:"tags"`
}

type MessageFieldTag map[string]interface{}
