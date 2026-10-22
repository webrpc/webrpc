package tests

//go:generate webrpc-gen -schema=./enum/enum.ridl -target=golang -pkg=enum -out=./enum/enum.gen.go

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/webrpc/webrpc/tests/enum"
	"github.com/webrpc/webrpc/tests/server"
)

func TestEnumHelpers(t *testing.T) {
	t.Run("int enum", func(t *testing.T) {
		require.Equal(t, []enum.Status{enum.StatusAvailable, enum.StatusNotAvailable}, enum.Status(0).Values())
		require.True(t, enum.StatusNotAvailable.IsValid())
		require.False(t, enum.Status(2).IsValid())

		var got enum.Status
		require.NoError(t, got.Parse("NOT_AVAILABLE"))
		require.Equal(t, enum.StatusNotAvailable, got)

		for _, s := range []string{"", "bogus", "not_available"} {
			require.Error(t, got.Parse(s), s)
			require.Equal(t, enum.StatusNotAvailable, got, "a failed Parse must not touch the receiver")
		}
		require.EqualError(t, got.Parse("bogus"), `invalid enum Status "bogus"`)
	})

	t.Run("int enum with json wire values", func(t *testing.T) {
		var got enum.Kind
		require.NoError(t, got.Parse("ADMIN"))
		require.Equal(t, enum.KindAdmin, got)
		require.Equal(t, "ADMIN", got.String())

		for _, s := range []string{"Admin", "admin", ""} {
			require.Error(t, got.Parse(s), s)
		}
	})

	t.Run("string enum", func(t *testing.T) {
		require.Equal(t, []enum.Tier{enum.TierFree, enum.TierPro}, enum.Tier("").Values())
		require.Equal(t, enum.Tier_values, enum.Tier("").Values())
		require.True(t, enum.TierPro.IsValid())
		require.False(t, enum.Tier("Pro").IsValid())

		var got enum.Tier
		require.NoError(t, got.Parse("pro"))
		require.Equal(t, enum.TierPro, got)

		for _, s := range []string{"", "Pro", "bogus"} {
			require.Error(t, got.Parse(s), s)
			require.Equal(t, enum.TierPro, got, "a failed Parse must not touch the receiver")
		}
	})

	t.Run("string enum without members", func(t *testing.T) {
		require.Empty(t, enum.Empty("").Values())
		require.False(t, enum.Empty("").IsValid())

		var got enum.Empty
		require.Error(t, got.Parse("anything"))
	})

	t.Run("maps are untouched", func(t *testing.T) {
		require.Equal(t, "AVAILABLE", enum.Status_name[enum.StatusAvailable])
		require.Equal(t, enum.StatusNotAvailable, enum.Status_value["NOT_AVAILABLE"])
	})

	t.Run("wrapped in a WebRPCError", func(t *testing.T) {
		var got enum.Status
		err := got.Parse("bogus")
		rpcErr := server.ErrWebrpcBadRequest.WithCause(err)
		require.ErrorIs(t, rpcErr, server.ErrWebrpcBadRequest)
		require.ErrorContains(t, rpcErr, `invalid enum Status "bogus"`)
	})
}
