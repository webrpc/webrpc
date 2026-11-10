package tests

//go:generate webrpc-gen -schema=./schema/optional_values.ridl -target=golang -pkg=optional -out=./optional/optional.gen.go

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/webrpc/webrpc/tests/optional"
	"github.com/webrpc/webrpc/tests/optional/types"
)

func TestOptionalPointerValues(t *testing.T) {
	t.Run("nil pointers are omitted", func(t *testing.T) {
		b, err := json.Marshal(optional.OptionalValues{})
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
		b, err := json.Marshal(v)
		require.NoError(t, err)
		assert.JSONEq(t, `{
			"required":"","inferred":"","custom":"","pointerReceiver":"",
			"timestamp":"0001-01-01T00:00:00Z","object":{"value":""},
			"pointerList":[],"listOverride":[],"mapOverride":{},
			"renamedValue":"","bare":null
		}`, string(b))

		var decoded optional.OptionalValues
		require.NoError(t, json.Unmarshal(b, &decoded))
		require.NotNil(t, decoded.Inferred)
		require.NotNil(t, decoded.Custom)
		require.NotNil(t, decoded.PointerReceiver)
		require.NotNil(t, decoded.Timestamp)
		assert.Equal(t, zero, *decoded.Inferred)
		assert.Equal(t, zero, *decoded.Custom)
		assert.Equal(t, pointerZero, *decoded.PointerReceiver)
		assert.True(t, decoded.Timestamp.IsZero())
	})

	t.Run("nonzero values and explicit tags are preserved", func(t *testing.T) {
		value := types.ZeroValue("set")
		b, err := json.Marshal(optional.OptionalValues{
			Inferred: &value, Custom: &value, ValueOverride: value,
			Explicit: &value, Renamed: &value, Hidden: &value, BareTag: &value,
		})
		require.NoError(t, err)
		assert.JSONEq(t, `{
			"required":"","inferred":"set","custom":"set","valueOverride":"set",
			"explicit":"set","renamedValue":"set","bare":"set"
		}`, string(b))
	})
}
