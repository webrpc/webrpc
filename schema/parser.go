package schema

import (
	"encoding/json"
	"regexp"
)

type Parser interface {
	Parse(schema *WebRPCSchema) error
}

func ParseSchemaJSON(jsondata []byte) (*WebRPCSchema, error) {
	var schema *WebRPCSchema
	err := json.Unmarshal(jsondata, &schema)
	if err != nil {
		return nil, err
	}

	err = schema.Parse(nil)
	if err != nil {
		return schema, err
	}

	return schema, nil
}

var NameWhitelistRexp = regexp.MustCompile(`(?i)^[a-zA-Z0-9]+(\w*[a-zA-Z0-9]+)?$`)

func IsValidArgName(s string) bool {
	// TODO: fix this regexp
	if !NameWhitelistRexp.MatchString(s) {
		return false
	}
	return true
}
