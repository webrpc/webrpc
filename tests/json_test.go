package tests

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/webrpc/webrpc/tests/client"
)

// JSON tags on generated structs: required fields always serialize, optional
// (`?`) fields with `,omitzero` are omitted when unset.
func TestOptionalFieldsJSONSerialization(t *testing.T) {
	t.Run("unset optional fields are omitted from JSON", func(t *testing.T) {
		v := client.OptFields{
			ReqBool:   true,
			ReqInt:    1,
			ReqString: "x",
			ReqList:   []string{"a"},
			ReqMap:    map[string]uint32{"a": 1},
		}
		b, err := json.Marshal(v)
		require.NoError(t, err)

		var m map[string]any
		require.NoError(t, json.Unmarshal(b, &m))
		assert.Equal(t, true, m["reqBool"])
		assert.Equal(t, float64(1), m["reqInt"])
		assert.Equal(t, "x", m["reqString"])
		assert.Equal(t, []any{"a"}, m["reqList"])
		assert.Equal(t, map[string]any{"a": float64(1)}, m["reqMap"])
		for _, key := range []string{"optBool", "optInt", "optString", "optList", "optMap"} {
			assert.NotContains(t, m, key)
		}
	})

	t.Run("optional fields set to a zero value are present in JSON", func(t *testing.T) {
		optBool := false
		optInt := int64(0)
		optString := ""
		v := client.OptFields{
			ReqBool:   false,
			ReqInt:    0,
			ReqString: "",
			ReqList:   []string{},
			ReqMap:    map[string]uint32{},
			OptBool:   &optBool,
			OptInt:    &optInt,
			OptString: &optString,
			OptList:   []string{},
			OptMap:    map[string]uint32{},
		}
		b, err := json.Marshal(v)
		require.NoError(t, err)

		var m map[string]any
		require.NoError(t, json.Unmarshal(b, &m))
		assert.Contains(t, m, "optBool")
		assert.Equal(t, false, m["optBool"])
		assert.Contains(t, m, "optInt")
		assert.Equal(t, float64(0), m["optInt"])
		assert.Contains(t, m, "optString")
		assert.Equal(t, "", m["optString"])
		assert.Contains(t, m, "optList")
		assert.Equal(t, []any{}, m["optList"])
		assert.Contains(t, m, "optMap")
		assert.Equal(t, map[string]any{}, m["optMap"])

		// Required zero values must never be dropped.
		assert.Equal(t, false, m["reqBool"])
		assert.Equal(t, float64(0), m["reqInt"])
		assert.Equal(t, "", m["reqString"])
		assert.Equal(t, []any{}, m["reqList"])
		assert.Equal(t, map[string]any{}, m["reqMap"])
	})
}
