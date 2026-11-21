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

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)

	stream, err := client.SubscribeMessages(ctx, t.Name(), nil)
	require.Nil(t, err)

	msgCount := 10000
	sendErr := make(chan error, 1)
	go func() {
		for i := 0; i < msgCount; i++ {
			if err := client.SendMessage(ctx, t.Name(), fmt.Sprintf("Hello %v", i)); err != nil {
				sendErr <- err
				return
			}
		}
		sendErr <- nil
	}()

	// Skip system messages like "joined".
	for received := 0; received < msgCount; {
		msg, err := stream.Read()
		require.Nil(t, err)
		if msg.Username == t.Name() {
			received++
		}
	}
	require.NoError(t, <-sendErr)

	cancel() // stop subscription

	_, err = stream.Read()
	if err != nil {
		assert.ErrorIs(t, err, proto.ErrWebrpcClientAborted)
	}
}

// A client that disconnects from an idle chat is announced as "left" right away.
func TestSubscriberLeavesIdleChat(t *testing.T) {
	t.Parallel()

	rpc := NewChatServer()
	srv := httptest.NewServer(rpc.Router())
	defer srv.Close()
	defer rpc.close(context.Background()) // Release parked handlers, so a failure doesn't hang srv.Close.

	client := proto.NewChatClient(srv.URL, &http.Client{})

	// Shorter than the keepalive interval, so only the disconnect can explain "left".
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	watcher, err := client.SubscribeMessages(ctx, "watcher", nil)
	require.NoError(t, err)

	leaverCtx, disconnect := context.WithCancel(ctx)
	_, err = client.SubscribeMessages(leaverCtx, "leaver", nil)
	require.NoError(t, err)

	waitFor := func(text string) {
		for {
			msg, err := watcher.Read()
			require.NoError(t, err)
			if msg.Text == text {
				return
			}
		}
	}

	waitFor("leaver joined")
	disconnect()
	waitFor("leaver left")
}

// A reconnecting client gets every buffered message after lastMessageId.
func TestReconnectDeliversMissedMessages(t *testing.T) {
	t.Parallel()

	rpc := NewChatServer()
	srv := httptest.NewServer(rpc.Router())
	defer srv.Close()

	client := proto.NewChatClient(srv.URL, &http.Client{})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for i := range 3 {
		require.NoError(t, client.SendMessage(ctx, t.Name(), fmt.Sprintf("Hello %v", i))) // Ids 0, 1, 2.
	}

	lastMessageId := uint64(0)
	stream, err := client.SubscribeMessages(ctx, t.Name(), &lastMessageId)
	require.NoError(t, err)

	for _, id := range []uint64{1, 2} {
		msg, err := stream.Read()
		require.NoError(t, err)
		require.Equal(t, id, msg.Id)
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

	stream, err := client.SubscribeMessages(ctx, t.Name(), nil)
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

	stream, err := client.SubscribeMessages(ctx, "", nil) // empty username
	require.Nil(t, err)                                   // only network/connection errors should come back here

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

		stream, err := client.SubscribeMessages(ctx, t.Name(), nil)
		require.Nil(t, err)

		for {
			msg, err := stream.Read()
			if err != nil {
				assert.ErrorIs(t, err, proto.ErrWebrpcClientAborted)
				break
			}
			t.Log(msg.Text)
		}
	}
}
