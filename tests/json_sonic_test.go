package tests

import (
	"testing"

	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/webrpc/webrpc/tests/client"
)

// jsonCfg mirrors the exact config the -json=sonic generator emits.
var sonicCfg = sonic.Config{
	NoNullSliceOrMap:        true,
	NoValidateJSONMarshaler: true,
	NoValidateJSONSkip:      true,
}.Froze()

// The sonic wire format renders nil required containers as []/{} so
// JavaScript clients never receive null for a list or map.
func TestSonicOptFieldsSerialization(t *testing.T) {
	b, err := sonicCfg.Marshal(client.OptFields{})
	require.NoError(t, err)

	var m map[string]any
	require.NoError(t, sonicCfg.Unmarshal(b, &m))
	assert.Equal(t, []any{}, m["reqList"])
	assert.Equal(t, map[string]any{}, m["reqMap"])
	for _, key := range []string{"optBool", "optInt", "optString", "optList", "optMap"} {
		assert.NotContains(t, m, key)
	}

	b, err = sonicCfg.Marshal(client.OptFields{OptList: []string{}, OptMap: map[string]uint32{}})
	require.NoError(t, err)
	require.NoError(t, sonicCfg.Unmarshal(b, &m))
	assert.Equal(t, []any{}, m["optList"])
	assert.Equal(t, map[string]any{}, m["optMap"])
}
