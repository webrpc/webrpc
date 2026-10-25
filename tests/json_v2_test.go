// A plain go1.27 tag, not `go1.27 || goexperiment.jsonv2`: the vet stdversion
// pass derives the file's language version from the build tag and only honors
// a tag that cannot be satisfied without a go version. The generated v2 output
// keeps the `||` form, since stdversion skips generated files.
//go:build go1.27

package tests

import (
	"testing"

	json "encoding/json/v2"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/webrpc/webrpc/tests/client"
)

// The encoding/json/v2 wire format renders nil required containers as []/{}
// and still omits unset optional fields, matching -json=sonic.
func TestV2OptFieldsSerialization(t *testing.T) {
	b, err := json.Marshal(client.OptFields{})
	require.NoError(t, err)

	var m map[string]any
	require.NoError(t, json.Unmarshal(b, &m))
	assert.Equal(t, []any{}, m["reqList"])
	assert.Equal(t, map[string]any{}, m["reqMap"])
	assert.Equal(t, false, m["reqBool"])
	assert.Equal(t, float64(0), m["reqInt"])
	assert.Equal(t, "", m["reqString"])
	for _, key := range []string{"optBool", "optInt", "optString", "optList", "optMap"} {
		assert.NotContains(t, m, key)
	}
}

// Set optional fields are always present, even at their zero value.
func TestV2OptionalFieldsSerialization(t *testing.T) {
	optBool := false
	optInt := int64(0)
	optString := ""
	b, err := json.Marshal(client.OptFields{
		OptBool:   &optBool,
		OptInt:    &optInt,
		OptString: &optString,
		OptList:   []string{},
		OptMap:    map[string]uint32{},
	})
	require.NoError(t, err)

	var m map[string]any
	require.NoError(t, json.Unmarshal(b, &m))
	assert.Equal(t, false, m["optBool"])
	assert.Equal(t, float64(0), m["optInt"])
	assert.Equal(t, "", m["optString"])
	assert.Equal(t, []any{}, m["optList"])
	assert.Equal(t, map[string]any{}, m["optMap"])
}

// Required zero scalars and a zero timestamp must never be dropped.
func TestV2ZeroValuesSerialization(t *testing.T) {
	b, err := json.Marshal(timestampFields{})
	require.NoError(t, err)

	var m map[string]any
	require.NoError(t, json.Unmarshal(b, &m))
	assert.Equal(t, "0001-01-01T00:00:00Z", m["ts"])
	assert.NotContains(t, m, "optTs")
}

// A nil required struct pointer is null under every encoder.
func TestV2NilStructPointer(t *testing.T) {
	b, err := json.Marshal(client.CreateItemRequest{})
	require.NoError(t, err)

	var m map[string]any
	require.NoError(t, json.Unmarshal(b, &m))
	assert.Equal(t, nil, m["item"])
}
