//go:build go1.27 || goexperiment.jsonv2

package streaming

import (
	"context"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type streamStub struct {
	StreamServiceServer
	err error
}

func (s streamStub) Subscribe(_ context.Context, channel string, stream SubscribeStreamWriter) error {
	for _, msg := range []string{channel + "-1", channel + "-2"} {
		if err := stream.Write(msg); err != nil {
			return err
		}
	}
	return s.err
}

func TestSubscribe(t *testing.T) {
	srv := httptest.NewServer(NewStreamServiceServer(streamStub{}))
	defer srv.Close()

	stream, err := NewStreamServiceClient(srv.URL, srv.Client()).Subscribe(context.Background(), "news")
	require.NoError(t, err)

	msg, err := stream.Read()
	require.NoError(t, err)
	assert.Equal(t, "news-1", msg)

	msg, err = stream.Read()
	require.NoError(t, err)
	assert.Equal(t, "news-2", msg)

	_, err = stream.Read()
	assert.ErrorIs(t, err, ErrWebrpcStreamFinished)
}

func TestSubscribeEndpointError(t *testing.T) {
	srv := httptest.NewServer(NewStreamServiceServer(streamStub{err: errors.New("boom")}))
	defer srv.Close()

	stream, err := NewStreamServiceClient(srv.URL, srv.Client()).Subscribe(context.Background(), "news")
	require.NoError(t, err)

	_, err = stream.Read()
	require.NoError(t, err)
	_, err = stream.Read()
	require.NoError(t, err)
	_, err = stream.Read()
	assert.ErrorIs(t, err, ErrWebrpcEndpoint)
}
