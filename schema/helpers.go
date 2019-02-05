package schema

import (
	"encoding/json"

	"github.com/pkg/errors"
)

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
