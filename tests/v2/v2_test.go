//go:build go1.27 || goexperiment.jsonv2

package v2

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/webrpc/webrpc/tests/optional"
	"github.com/webrpc/webrpc/tests/optional/types"
)

func TestOptionalPointerValues(t *testing.T) {
	t.Run("nil pointers are omitted", func(t *testing.T) {
		b, err := jsonCfg.Marshal(optional.OptionalValues{})
		require.NoError(t, err)
		assert.JSONEq(t, `{"required":"","bare":null}`, string(b))
	})

	t.Run("present zero values survive a round trip", func(t *testing.T) {
		zero := types.ZeroValue("")
		pointerZero := types.PointerZero("")
		stamp := time.Time{}
		empty := []string{}
		v := optional.OptionalValues{
			Inferred:        &zero,
			Custom:          &zero,
			PointerReceiver: &pointerZero,
			Timestamp:       &stamp,
			Object:          &optional.Object{},
			PointerList:     &empty,
			ListOverride:    empty,
			MapOverride:     map[string]string{},
			Explicit:        &zero,
			Renamed:         &zero,
			Hidden:          &zero,
		}
		b, err := jsonCfg.Marshal(v)
		require.NoError(t, err)
		assert.JSONEq(t, `{
			"required":"","inferred":"","custom":"","pointerReceiver":"",
			"timestamp":"0001-01-01T00:00:00Z","object":{"value":""},
			"pointerList":[],"listOverride":[],"mapOverride":{},
			"renamedValue":"","bare":null
		}`, string(b))
	})
}

func TestNilContainers(t *testing.T) {
	b, err := jsonCfg.Marshal(OptFields{})
	require.NoError(t, err)
	assert.Contains(t, string(b), `"reqList":[]`)
	assert.Contains(t, string(b), `"reqMap":{}`)
}

func TestCaseInsensitiveNames(t *testing.T) {
	var out OptFields
	require.NoError(t, jsonCfg.Unmarshal([]byte(`{"REQSTRING":"x"}`), &out))
	assert.Equal(t, "x", out.ReqString)
}

func TestMapKeysSorted(t *testing.T) {
	m := map[string]int{"delta": 4, "alpha": 1, "charlie": 3, "bravo": 2, "echo": 5}
	first, err := jsonCfg.Marshal(m)
	require.NoError(t, err)
	second, err := jsonCfg.Marshal(m)
	require.NoError(t, err)
	assert.Equal(t, string(first), string(second))
	assert.Equal(t, `{"alpha":1,"bravo":2,"charlie":3,"delta":4,"echo":5}`, string(first))
}
