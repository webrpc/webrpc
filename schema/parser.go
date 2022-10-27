package schema

import (
	"encoding/json"
	"regexp"
	"unicode"
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

	err = schema.Validate()
	if err != nil {
		return schema, err
	}

	return schema, nil
}

var NameWhitelistRexp = regexp.MustCompile(`^[a-zA-Z]+[a-zA-Z0-9_]*$`)

func IsValidArgName(s string) bool {
	if !NameWhitelistRexp.MatchString(s) {
		return false
	}
	return true
}

func IsStartsWithUpper(s string) bool {
	if len(s) == 0 {
		return false
	}
	if !unicode.IsUpper(rune(s[0])) {
		return false
	}
	return true
}
