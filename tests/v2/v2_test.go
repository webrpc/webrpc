// This file avoids importing encoding/json/v2 directly: its build tag also
// holds under Go 1.25/1.26 with GOEXPERIMENT=jsonv2, where the vet stdversion
// pass rejects v2 std symbols at the module's language version. It asserts the
// wire format the generated v2 helpers produce, and parses it with the v1
// encoding/json decoder.
//go:build go1.27 || goexperiment.jsonv2

package v2

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// v2WireFields mirrors a generated struct with a required nil struct pointer
// and a required timestamp.
type v2WireFields struct {
	RPtr *Item     `json:"rPtr"`
	Ts   time.Time `json:"ts"`
}

func countTrailingNewlines(s string) int {
	return len(s) - len(strings.TrimRight(s, "\n"))
}

// v2 renders nil required containers as []/{} instead of null.
func TestV2OptFieldsWire(t *testing.T) {
	var buf bytes.Buffer
	require.NoError(t, jsonEncode(&buf, OptFields{}))
	assert.True(t, strings.HasSuffix(buf.String(), "\n"))

	var m map[string]any
	require.NoError(t, json.Unmarshal(buf.Bytes(), &m))
	assert.Equal(t, []any{}, m["reqList"])
	assert.Equal(t, map[string]any{}, m["reqMap"])
	assert.Equal(t, false, m["reqBool"])
	assert.Equal(t, float64(0), m["reqInt"])
	assert.Equal(t, "", m["reqString"])
	for _, key := range []string{"optBool", "optInt", "optString", "optList", "optMap"} {
		assert.NotContains(t, m, key)
	}
}

// Set optional fields are present, even at their zero value.
func TestV2OptionalFieldsWire(t *testing.T) {
	optBool := false
	optInt := int64(0)
	optString := ""
	var buf bytes.Buffer
	require.NoError(t, jsonEncode(&buf, OptFields{
		OptBool:   &optBool,
		OptInt:    &optInt,
		OptString: &optString,
		OptList:   []string{},
		OptMap:    map[string]uint32{},
	}))

	var m map[string]any
	require.NoError(t, json.Unmarshal(buf.Bytes(), &m))
	assert.Equal(t, false, m["optBool"])
	assert.Equal(t, float64(0), m["optInt"])
	assert.Equal(t, "", m["optString"])
	assert.Equal(t, []any{}, m["optList"])
	assert.Equal(t, map[string]any{}, m["optMap"])
}

// A nil required struct pointer stays null, and a zero timestamp is never
// dropped; both fields must appear in the encoded object.
func TestV2RequiredPointerAndTimestampWire(t *testing.T) {
	var buf bytes.Buffer
	require.NoError(t, jsonEncode(&buf, v2WireFields{}))
	out := buf.String()

	assert.Contains(t, out, `"rPtr":null`)
	assert.Contains(t, out, `"ts":"0001-01-01T00:00:00Z"`)
	assert.Equal(t, 1, countTrailingNewlines(out))
}

// json.Encoder.Encode appends one newline; the v2 helper must match exactly.
func TestV2EncodeTrailingNewline(t *testing.T) {
	var buf bytes.Buffer
	require.NoError(t, jsonEncode(&buf, OptFields{}))
	assert.Equal(t, 1, countTrailingNewlines(buf.String()))
}

// v2 matches field names case-insensitively like v1, so differently-cased
// members still fill the generated struct.
func TestV2DecodeCaseInsensitive(t *testing.T) {
	var out OptFields
	require.NoError(t, jsonUnmarshal([]byte(`{"REQLIST":["a"],"ReqString":"x"}`), &out))
	assert.Equal(t, []string{"a"}, out.ReqList)
	assert.Equal(t, "x", out.ReqString)
}

func TestV2JSONHelpers(t *testing.T) {
	var buf bytes.Buffer
	require.NoError(t, jsonEncode(&buf, OptFields{ReqList: []string{"a"}}))
	assert.True(t, strings.HasSuffix(buf.String(), "\n"))

	var out OptFields
	require.NoError(t, jsonDecode(bytes.NewReader(buf.Bytes()), &out))
	assert.Equal(t, []string{"a"}, out.ReqList)

	assert.ErrorIs(t, jsonDecode(strings.NewReader(""), &out), io.EOF)
	assert.ErrorIs(t, jsonDecode(strings.NewReader("  \n"), &out), io.EOF)
	assert.Error(t, jsonDecode(strings.NewReader(`{"reqList":["a"]} trailing`), &out))
}
