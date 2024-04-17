package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/webrpc/webrpc/_example/golang-sse/proto"
)

var (
	client proto.ChatClient
)

func TestMain(m *testing.M) {
	rpc := NewChatServer()
	srv := httptest.NewServer(rpc.Router())
	defer srv.Close()

	// opts := &devslog.Options{
	// 	MaxSlicePrintSize: 4,
	// 	SortKeys:          true,
	// 	TimeFormat:        "[04:05:06]",
	// 	NewLineAfterLog:   true,
	// 	DebugColor:        devslog.Magenta,
	// }

	// logger := slog.New(devslog.NewHandler(os.Stdout, opts))
	// slog.SetDefault(logger)

	client = proto.NewChatClient(srv.URL, &http.Client{})
	time.Sleep(time.Millisecond * 500)

	os.Exit(m.Run())
}

func TestStream10k(t *testing.T) {
	t.Parallel()

	rpc := NewChatServer()
	srv := httptest.NewServer(rpc.Router())
	defer srv.Close()

	client := proto.NewChatClient(srv.URL, &http.Client{})

	ctx, cancel := context.WithCancel(context.Background())

	stream, err := client.SubscribeMessages(ctx, t.Name())
	require.Nil(t, err)

	msgCount := 10000
	go func() {
		for i := 0; i < msgCount; i++ {
			if err := client.SendMessage(ctx, t.Name(), fmt.Sprintf("Hello %v", i)); err != nil {
				t.Fatal(err)
			}
		}
	}()

	for i := 0; i < msgCount; i++ {
		_, err := stream.Read()
		require.Nil(t, err)
	}

	cancel() // stop subscription

	_, err = stream.Read()
	if err != nil {
		assert.ErrorIs(t, err, proto.ErrWebrpcClientDisconnected)
	}
}

func TestStreamServerConnectionLost(t *testing.T) {
	t.Parallel()

	rpc := NewChatServer()
	srv := httptest.NewServer(rpc.Router())
	defer srv.Close()

	go func() {
		<-time.After(2 * time.Second)
		srv.Config.Close()
	}()

	client := proto.NewChatClient(srv.URL, &http.Client{})
	time.Sleep(time.Millisecond * 500)

	ctx := context.Background()

	stream, err := client.SubscribeMessages(ctx, t.Name())
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

func TestStreamCustomError(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	stream, err := client.SubscribeMessages(ctx, "") // empty username
	require.Nil(t, err)                              // only network/connection errors should come back here

	_, err = stream.Read() // we should receive RPC handler errors (e.g. EmptyUsername) only when we start reading the data
	require.Error(t, err)
	require.ErrorIs(t, err, proto.ErrEmptyUsername)
}

func TestStreamClientTimeout(t *testing.T) {
	t.Parallel()

	t.Run("0.1s", testStreamClientTimeout(t, 100*time.Millisecond))
	t.Run("6s", testStreamClientTimeout(t, 5*time.Second))
	t.Run("12s", testStreamClientTimeout(t, 12*time.Second)) // over 10s (ping-alive)
}

func testStreamClientTimeout(t *testing.T, timeout time.Duration) func(t *testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		stream, err := client.SubscribeMessages(ctx, t.Name())
		require.Nil(t, err)

		for {
			msg, err := stream.Read()
			if err != nil {
				assert.ErrorIs(t, err, proto.ErrWebrpcClientDisconnected)
				break
			}
			t.Log(msg.Text)
		}
	}
}
