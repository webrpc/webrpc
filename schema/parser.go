package schema

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type Parser interface {
	Parse(schema *WebRPCSchema) error
}

func ParseSchema(data []byte) (*WebRPCSchema, error) {
	var schema *WebRPCSchema
	err := json.Unmarshal(data, &schema)
	if err != nil {
		return nil, err
	}

	err = schema.Parse(nil)
	if err != nil {
		return schema, err
	}

	return schema, nil
}

func ParseSchemaFile(fpath string) (*WebRPCSchema, error) {
	return nil, errors.New("unimplemented")
}
