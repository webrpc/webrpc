package main

import (
	"context"
	"math/rand"
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

	ctx := context.Background()

	stream, err := client.SubscribeMessages(ctx, t.Name())
	require.Nil(t, err)

	go func() {
		for i := 0; i < 10000; i++ {
			if err := client.SendMessage(ctx, t.Name(), "Hello"); err != nil {
				t.Fatal(err)
			}
		}
	}()

	for i := 0; i < 10000; i++ {
		_, err := stream.Read()
		if err != nil {
			assert.ErrorIs(t, err, proto.ErrWebrpcStreamFinished)
			break
		}
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

	stream, err := client.SubscribeMessages(ctx, "")
	require.Nil(t, err)

	_, err = stream.Read()
	require.Error(t, err)
	require.ErrorIs(t, err, proto.ErrEmptyUsername)
}

func TestStreamClientTimeout(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Duration(rand.Int63n(15)+1)*time.Second)

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
