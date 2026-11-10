package tests

//go:generate webrpc-gen -schema=./enum/enum.ridl -target=golang -pkg=enum -out=./enum/enum.gen.go

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/webrpc/webrpc/tests/enum"
)

func TestEnumHelpers(t *testing.T) {
	t.Run("int enum", func(t *testing.T) {
		require.Equal(t, []enum.Status{enum.StatusAvailable, enum.StatusNotAvailable}, enum.Status(0).Values())
		require.True(t, enum.StatusAvailable.IsValid())
		require.True(t, enum.StatusNotAvailable.IsValid())
		require.False(t, enum.Status(2).IsValid())
	})

	t.Run("int enum with json wire values", func(t *testing.T) {
		require.Equal(t, []enum.Kind{enum.KindUser, enum.KindAdmin}, enum.Kind(0).Values())
		require.True(t, enum.KindAdmin.IsValid())
		require.False(t, enum.Kind(2).IsValid())
	})

	t.Run("string enum", func(t *testing.T) {
		require.Equal(t, []enum.Tier{enum.TierFree, enum.TierPro}, enum.Tier("").Values())
		require.Equal(t, enum.Tier_values, enum.Tier("").Values())
		require.True(t, enum.TierPro.IsValid())
		for _, s := range []string{"", "Pro", "bogus"} {
			require.False(t, enum.Tier(s).IsValid(), s)
		}
	})

	t.Run("string enum without members", func(t *testing.T) {
		require.Empty(t, enum.Empty("").Values())
		require.False(t, enum.Empty("").IsValid())
		require.False(t, enum.Empty("anything").IsValid())
	})

	t.Run("maps are untouched", func(t *testing.T) {
		require.Equal(t, "AVAILABLE", enum.Status_name[enum.StatusAvailable])
		require.Equal(t, enum.StatusNotAvailable, enum.Status_value["NOT_AVAILABLE"])
	})
}
