package main

import (
	"context"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/golang-cz/devslog"
	"github.com/golang-cz/transport"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/webrpc/webrpc/_example/golang-sse/proto"
)

var (
	client proto.ChatClient
)

func TestMain(m *testing.M) {
	srv := httptest.NewServer(router())
	defer srv.Close()

	opts := &devslog.Options{
		MaxSlicePrintSize: 4,
		SortKeys:          true,
		TimeFormat:        "[04:05]",
		NewLineAfterLog:   true,
		DebugColor:        devslog.Magenta,
	}

	logger := slog.New(devslog.NewHandler(os.Stdout, opts))
	slog.SetDefault(logger)

	client = proto.NewChatClient(srv.URL, &http.Client{
		Transport: transport.Chain(http.DefaultTransport, transport.DebugRequests),
	})
	time.Sleep(time.Millisecond * 500)

	os.Exit(m.Run())
}

func TestStream(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	stream, err := client.SubscribeMessages(ctx, 0)
	require.Nil(t, err)

	for {
		msg, err := stream.Read()
		if err != nil {
			assert.ErrorIs(t, err, proto.ErrWebrpcStreamFinished)
			break
		}
		t.Log(msg.Text)
	}
}

func TestStreamServerConnectionLost(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(router())
	defer srv.Close()

	go func() {
		<-time.After(2 * time.Second)
		srv.Config.Close()
	}()

	client := proto.NewChatClient(srv.URL, &http.Client{
		Transport: transport.Chain(http.DefaultTransport, transport.DebugRequests),
	})
	time.Sleep(time.Millisecond * 500)

	ctx := context.Background()

	stream, err := client.SubscribeMessages(ctx, 0)
	require.Nil(t, err)

	for {
		msg, err := stream.Read()
		if err != nil {
			assert.ErrorIs(t, err, proto.ErrWebrpcStreamLost)
			break
		}
		t.Log(msg.Text)
	}
}

func TestStreamServerTimeout(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	stream, err := client.SubscribeMessages(ctx, 2)
	require.Nil(t, err)

	for {
		msg, err := stream.Read()
		if err != nil {
			assert.ErrorIs(t, err, proto.ErrConnectionTooLong)
			break
		}
		t.Log(msg.Text)
	}
}

func TestStreamClientTimeout(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 2*time.Second)

	stream, err := client.SubscribeMessages(ctx, 0)
	require.Nil(t, err)

	for {
		msg, err := stream.Read()
		if err != nil {
			assert.ErrorIs(t, err, proto.ErrWebrpcRequestFailed)
			break
		}
		t.Log(msg.Text)
	}
}
