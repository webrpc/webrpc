package schema

import (
	"encoding/json"
	"regexp"

	"github.com/pkg/errors"
)

type Parser interface {
	Parse(schema *WebRPCSchema) error
}

var NameWhitelistRexp = regexp.MustCompile("(?i)[a-z0-9_]+")

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

func IsValidArgName(s string) bool {
	// TODO: fix this regexp
	if !NameWhitelistRexp.MatchString(s) {
		return false
	}
	return true
}
