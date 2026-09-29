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
		require.Equal(t, []enum.Status{enum.StatusAvailable, enum.StatusNotAvailable}, enum.StatusValues())
		require.True(t, enum.StatusNotAvailable.IsValid())
		require.False(t, enum.Status(2).IsValid())

		got, err := enum.ParseStatus("NOT_AVAILABLE")
		require.NoError(t, err)
		require.Equal(t, enum.StatusNotAvailable, got)

		for _, s := range []string{"", "bogus", "not_available"} {
			got, err := enum.ParseStatus(s)
			require.ErrorIs(t, err, enum.ErrInvalidEnum, s)
			require.Zero(t, got, s)
		}
		_, err = enum.ParseStatus("bogus")
		require.EqualError(t, err, `invalid enum Status "bogus"`)
	})

	t.Run("int enum with json wire values", func(t *testing.T) {
		got, err := enum.ParseKind("ADMIN")
		require.NoError(t, err)
		require.Equal(t, enum.KindAdmin, got)
		require.Equal(t, "ADMIN", got.String())

		for _, s := range []string{"Admin", "admin", ""} {
			_, err := enum.ParseKind(s)
			require.ErrorIs(t, err, enum.ErrInvalidEnum, s)
		}
	})

	t.Run("string enum", func(t *testing.T) {
		require.Equal(t, []enum.Tier{enum.TierFree, enum.TierPro}, enum.TierValues())
		require.True(t, enum.TierPro.IsValid())
		require.False(t, enum.Tier("Pro").IsValid())

		got, err := enum.ParseTier("pro")
		require.NoError(t, err)
		require.Equal(t, enum.TierPro, got)

		for _, s := range []string{"", "Pro", "bogus"} {
			got, err := enum.ParseTier(s)
			require.ErrorIs(t, err, enum.ErrInvalidEnum, s)
			require.Zero(t, got, s)
		}
	})

	t.Run("string enum without members", func(t *testing.T) {
		require.Empty(t, enum.EmptyValues())
		require.False(t, enum.Empty("").IsValid())

		_, err := enum.ParseEmpty("anything")
		require.ErrorIs(t, err, enum.ErrInvalidEnum)
	})

	t.Run("wrapped in a WebRPCError", func(t *testing.T) {
		_, err := enum.ParseStatus("bogus")
		rpcErr := server.ErrWebrpcBadRequest.WithCause(err)
		require.ErrorIs(t, rpcErr, enum.ErrInvalidEnum)
		require.ErrorIs(t, rpcErr, server.ErrWebrpcBadRequest)
	})
}
