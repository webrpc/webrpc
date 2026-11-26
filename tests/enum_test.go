package tests

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/webrpc/webrpc/tests/enum"
)

func TestEnumStrict(t *testing.T) {
	var a enum.Animal

	err := a.UnmarshalText([]byte("ELEPHANT"))
	require.Error(t, err)
	assert.Equal(t, enum.Animal(0), a)

	err = a.UnmarshalText([]byte(""))
	require.Error(t, err)

	err = a.UnmarshalText([]byte("UNKNOWN"))
	require.Error(t, err)

	err = a.UnmarshalText([]byte("CAT"))
	require.NoError(t, err)
	assert.Equal(t, enum.AnimalCat, a)

	b, err := enum.AnimalCat.MarshalText()
	require.NoError(t, err)
	assert.Equal(t, "CAT", string(b))

	_, err = enum.Animal(999).MarshalText()
	require.Error(t, err)

	c := enum.EnumContainer{Animal: enum.AnimalDog}
	data, err := json.Marshal(c)
	require.NoError(t, err)
	assert.JSONEq(t, `{"animal":"DOG","fallbackFirst":"UNKNOWN","fallbackLast":"CAT","signedFallbackLast":"CAT"}`, string(data))

	var decoded enum.EnumContainer
	require.NoError(t, json.Unmarshal(data, &decoded))
	assert.Equal(t, enum.AnimalDog, decoded.Animal)

	err = json.Unmarshal([]byte(`{"animal":"ELEPHANT"}`), &decoded)
	require.Error(t, err)
}

func TestEnumFallbackFirst(t *testing.T) {
	var f enum.FallbackFirst

	err := f.UnmarshalText([]byte("ELEPHANT"))
	require.NoError(t, err)
	assert.Equal(t, enum.FallbackFirstUnknown, f)

	err = f.UnmarshalText([]byte(""))
	require.NoError(t, err)
	assert.Equal(t, enum.FallbackFirstUnknown, f)

	err = f.UnmarshalText([]byte("UNKNOWN"))
	require.NoError(t, err)
	assert.Equal(t, enum.FallbackFirstUnknown, f)

	err = f.UnmarshalText([]byte("CAT"))
	require.NoError(t, err)
	assert.Equal(t, enum.FallbackFirstCat, f)

	b, err := enum.FallbackFirstUnknown.MarshalText()
	require.NoError(t, err)
	assert.Equal(t, "UNKNOWN", string(b))

	b, err = enum.FallbackFirst(999).MarshalText()
	require.NoError(t, err)
	assert.Equal(t, "UNKNOWN", string(b))

	assert.Equal(t, uint32(0), uint32(enum.FallbackFirstUnknown))
}

func TestEnumFallbackLast(t *testing.T) {
	var f enum.FallbackLast

	err := f.UnmarshalText([]byte("ELEPHANT"))
	require.NoError(t, err)
	assert.Equal(t, enum.FallbackLastUnknown, f)

	err = f.UnmarshalText([]byte(""))
	require.NoError(t, err)
	assert.Equal(t, enum.FallbackLastUnknown, f)

	err = f.UnmarshalText([]byte("CAT"))
	require.NoError(t, err)
	assert.Equal(t, enum.FallbackLastCat, f)

	b, err := enum.FallbackLast(999).MarshalText()
	require.NoError(t, err)
	assert.Equal(t, "UNKNOWN", string(b))

	assert.Equal(t, uint32(math.MaxUint32), uint32(enum.FallbackLastUnknown))
}

func TestEnumSignedFallbackLast(t *testing.T) {
	var s enum.SignedFallbackLast

	err := s.UnmarshalText([]byte("ELEPHANT"))
	require.NoError(t, err)
	assert.Equal(t, enum.SignedFallbackLastUnknown, s)

	err = s.UnmarshalText([]byte(""))
	require.NoError(t, err)
	assert.Equal(t, enum.SignedFallbackLastUnknown, s)

	err = s.UnmarshalText([]byte("CAT"))
	require.NoError(t, err)
	assert.Equal(t, enum.SignedFallbackLastCat, s)

	b, err := enum.SignedFallbackLast(999).MarshalText()
	require.NoError(t, err)
	assert.Equal(t, "UNKNOWN", string(b))

	assert.Equal(t, int32(-1), int32(enum.SignedFallbackLastUnknown))
}
