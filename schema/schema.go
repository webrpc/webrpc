package schema

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// the webrpc proto schema -- would be nice to think of a new name for this
// so its more distinct / identifiable

// but essentially, it takes the proto.json of a webrpc api schema, and will
// be used by `gen` pkg to code-generate a bunch of stuff

type WebRPCSchema struct {
	Schema string `json:"schema"`
	App    string `json:"app"`

	Messages []Message `json:"messages"`
	Services []Service `json:"services"`
}

// TODO: how do we do Enums ..?

type Message struct {
	Name   string         `json:"name"`
	Type   MessageType    `json:"type"`
	Fields []MessageField `json:"fields"`
}

type MessageType string // TODO: must be `enum` or `struct`

type MessageField struct {
	Name     string            `json:"name"`
	Type     DataType          `json:"type"`
	Optional bool              `json:"optional"`
	Tags     []MessageFieldTag `json:"tags"`
}

type MessageFieldTag map[string]interface{}

type Service struct {
	Name    string   `json:"name"`
	Methods []Method `json:"methods"`
}

type Method struct {
	Name    string         `json:"name"`
	Inputs  []MethodInput  `json:"inputs"`
	Outputs []MethodOutput `json:"outputs`
}

type MethodInput struct {
	Name string   `json:"name"`
	Type DataType `json:"type"`
}

type MethodOutput struct {
	Name string   `json:"name"`
	Type DataType `json:"type"`
}

// VarName used for Message.Name, MessageField.Name, Service.Name
type VarName string

func (s *WebRPCSchema) Validate() error {
	return nil
}

// TODO: have each object
// should we pass the parent.. how does it know the context?

type Validator interface {
	Validate() error
}

func ParseSchema(data []byte) (*WebRPCSchema, error) {
	var schema *WebRPCSchema
	err := json.Unmarshal(data, &schema)
	if err != nil {
		return nil, err
	}

	err = schema.Validate()
	if err != nil {
		return schema, err
	}

	return schema, nil
}

func ParseSchemaFile(fpath string) (*WebRPCSchema, error) {
	return nil, errors.New("unimplemented")
}
