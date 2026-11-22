package main

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/webrpc/webrpc/_example/golang-sse/proto"
)

// subscribeFunc implements proto.ChatServer with a stub SubscribeMessages.
type subscribeFunc func(ctx context.Context, stream proto.SubscribeMessagesStreamWriter) error

func (f subscribeFunc) SendMessage(ctx context.Context, username string, text string) error {
	return nil
}

func (f subscribeFunc) SubscribeMessages(ctx context.Context, username string, lastMessageId *uint64, stream proto.SubscribeMessagesStreamWriter) error {
	return f(ctx, stream)
}

func waitForCancel(ctx context.Context, stream proto.SubscribeMessagesStreamWriter) error {
	<-ctx.Done()
	return ctx.Err()
}

// serve runs a SubscribeMessages request against svc, pinging the client
// every keepAlive, and returns the error the server reported, if any.
func serve(t *testing.T, ctx context.Context, w http.ResponseWriter, keepAlive time.Duration, svc subscribeFunc) proto.WebRPCError {
	t.Helper()

	var rpcErr proto.WebRPCError
	srv := proto.NewChatServer(svc, &proto.Options{
		OnError:                 func(r *http.Request, err *proto.WebRPCError) { rpcErr = *err },
		StreamKeepAliveInterval: keepAlive,
	})

	req := httptest.NewRequest(http.MethodPost, "/v1/Chat/SubscribeMessages", strings.NewReader(`{"username":"test"}`))
	req.Header.Set("Content-Type", "application/json")

	done := make(chan struct{})
	go func() {
		defer close(done)
		srv.ServeHTTP(w, req.WithContext(ctx))
	}()

	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("handler did not return")
	}
	return rpcErr
}

// failWriter fails every Write after the stream's initial ping, like a dead
// peer behind middleware that doesn't cancel the request context.
type failWriter struct {
	*httptest.ResponseRecorder
	writes atomic.Int32
}

func (f *failWriter) Write(p []byte) (int, error) {
	if f.writes.Add(1) > 1 {
		return 0, errors.New("write: broken pipe")
	}
	return len(p), nil
}

// flushFailWriter fails every FlushError after the stream's initial ping.
type flushFailWriter struct {
	*httptest.ResponseRecorder
	flushes atomic.Int32
}

func (f *flushFailWriter) FlushError() error {
	if f.flushes.Add(1) > 1 {
		return errors.New("flush: broken pipe")
	}
	return nil
}

func TestKeepAliveWriteErrorCancelsStream(t *testing.T) {
	w := &failWriter{ResponseRecorder: httptest.NewRecorder()}
	err := serve(t, context.Background(), w, 10*time.Millisecond, waitForCancel)

	require.ErrorIs(t, err, proto.ErrWebrpcClientAborted)
}

func TestKeepAliveFlushErrorCancelsStream(t *testing.T) {
	w := &flushFailWriter{ResponseRecorder: httptest.NewRecorder()}
	err := serve(t, context.Background(), w, 10*time.Millisecond, waitForCancel)

	require.ErrorIs(t, err, proto.ErrWebrpcClientAborted)
}

func TestKeepAliveDisabled(t *testing.T) {
	ctx, disconnect := context.WithCancel(context.Background())
	err := serve(t, ctx, httptest.NewRecorder(), -1, func(ctx context.Context, stream proto.SubscribeMessagesStreamWriter) error {
		disconnect()
		return waitForCancel(ctx, stream)
	})

	require.ErrorIs(t, err, proto.ErrWebrpcClientAborted)
}

// A context error of the service's own, e.g. a database timeout, is not a client abort.
func TestServiceContextErrorIsEndpointError(t *testing.T) {
	err := serve(t, context.Background(), httptest.NewRecorder(), 0, func(ctx context.Context, stream proto.SubscribeMessagesStreamWriter) error {
		return context.DeadlineExceeded
	})

	require.ErrorIs(t, err, proto.ErrWebrpcEndpoint)
}

// A stream writer that outlives the service method must not write to the response.
func TestNoWriteAfterErrorFrame(t *testing.T) {
	var leaked proto.SubscribeMessagesStreamWriter
	rec := httptest.NewRecorder()

	serve(t, context.Background(), rec, 0, func(ctx context.Context, stream proto.SubscribeMessagesStreamWriter) error {
		leaked = stream
		return errors.New("failed")
	})
	body := rec.Body.String()

	require.ErrorIs(t, leaked.Write(&proto.Message{}), proto.ErrWebrpcStreamFinished)
	require.Equal(t, body, rec.Body.String())

	frames := strings.Split(strings.TrimSpace(body), "\n")
	require.Contains(t, frames[len(frames)-1], `"webrpcError"`)
}

func TestNonFlushableWriterFailsBeforeHeaders(t *testing.T) {
	rec := httptest.NewRecorder()
	called := false

	// The anonymous struct hides the recorder's Flush method.
	serve(t, context.Background(), struct{ http.ResponseWriter }{rec}, 0, func(ctx context.Context, stream proto.SubscribeMessagesStreamWriter) error {
		called = true
		return nil
	})

	require.Equal(t, http.StatusInternalServerError, rec.Code)
	require.False(t, called)
}
