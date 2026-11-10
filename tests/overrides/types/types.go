// Package types holds the Go types behind the go.field.type overrides in rest_overrides.ridl.
package types

import "strings"

// Label has no methods, so REST values decode as JSON strings.
type Label string

// Tag decodes with UnmarshalText, which lowercases the raw value.
type Tag string

func (t Tag) MarshalText() ([]byte, error) { return []byte(t), nil }

func (t *Tag) UnmarshalText(b []byte) error {
	*t = Tag(strings.ToLower(string(b)))
	return nil
}
