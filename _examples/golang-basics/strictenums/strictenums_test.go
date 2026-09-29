package strictenums

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStrictEnums_KnownNamesRoundTrip(t *testing.T) {
	in := Project{Id: 1, Status: StatusArchived, Level: LevelHigh}

	b, err := json.Marshal(in)
	require.NoError(t, err)
	require.JSONEq(t, `{"id":1,"status":"archived","level":"high"}`, string(b))

	var out Project
	require.NoError(t, json.Unmarshal(b, &out))
	require.Equal(t, in, out)
	require.True(t, out.Status.IsValid())
	require.True(t, out.Level.IsValid())
}

func TestStrictEnums_ZeroValueStaysValid(t *testing.T) {
	var s Status
	require.Equal(t, StatusActive, s)
	require.True(t, s.IsValid())
}

func TestStrictEnums_UnknownNameNeverDecodesToADeclaredValue(t *testing.T) {
	var p Project
	require.NoError(t, json.Unmarshal([]byte(`{"id":1,"status":"frozen","level":"extreme"}`), &p))

	require.False(t, p.Status.IsValid())
	require.NotEqual(t, StatusActive, p.Status)
	require.False(t, p.Status.Is(StatusActive, StatusArchived, StatusSuspended))
	require.Equal(t, "", p.Status.String())

	require.False(t, p.Level.IsValid())
	require.False(t, p.Level.Is(LevelLow, LevelHigh))
}

func TestStrictEnums_InvalidValueFailsToMarshal(t *testing.T) {
	var p Project
	require.NoError(t, json.Unmarshal([]byte(`{"id":1,"status":"frozen","level":"low"}`), &p))

	_, err := json.Marshal(p)
	require.Error(t, err)
	require.ErrorContains(t, err, "Status: invalid value")
}
