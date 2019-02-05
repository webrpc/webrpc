package schema

type Message struct {
	Name   *VarName        `json:"name"`
	Type   *MessageType    `json:"type"`
	Fields []*MessageField `json:"fields"`
}

// NOTE: for now just skip the enum stuff during generation..
type MessageType string // TODO: must be `enum` or `struct`

type MessageField struct {
	Name *VarName `json:"name"`
	Type *VarType `json:"type"`

	Optional bool   `json:"optional"`
	Value    string `json:"value"` // used by enums

	// Tags store extra metadata for plugins
	Tags []MessageFieldTag `json:"tags"`
}

type MessageFieldTag map[string]interface{}
