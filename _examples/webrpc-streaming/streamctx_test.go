package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/webrpc/webrpc/_example/golang-sse/proto"
)

// blockingChatServer blocks SubscribeMessages until ctx is cancelled, then
// reports how it was cancelled.
type blockingChatServer struct {
	started  chan struct{}
	returned chan struct{}
	ctxErr   error
}

func (s *blockingChatServer) SendMessage(ctx context.Context, username string, text string) error {
	return nil
}

func (s *blockingChatServer) SubscribeMessages(ctx context.Context, username string, lastMessageId *uint64, stream proto.SubscribeMessagesStreamWriter) error {
	close(s.started)
	<-ctx.Done()
	s.ctxErr = ctx.Err()
	close(s.returned)
	return ctx.Err()
}

// TestCleanCloseCancelsHandlerContext proves case 1: a client FIN makes
// net/http cancel r.Context(), unblocking the service handler.
func TestCleanCloseCancelsHandlerContext(t *testing.T) {
	svc := &blockingChatServer{started: make(chan struct{}), returned: make(chan struct{})}
	srv := httptest.NewServer(proto.NewChatServer(svc))
	defer srv.Close()

	conn, err := net.Dial("tcp", srv.Listener.Addr().String())
	require.NoError(t, err)

	body := `{"username":"test"}`
	req := fmt.Sprintf("POST /v1/Chat/SubscribeMessages HTTP/1.1\r\nHost: %s\r\nContent-Type: application/json\r\nContent-Length: %d\r\n\r\n%s",
		srv.Listener.Addr().String(), len(body), body)
	_, err = conn.Write([]byte(req))
	require.NoError(t, err)

	buf := make([]byte, 64)
	require.NoError(t, conn.SetReadDeadline(time.Now().Add(2*time.Second)))
	n, err := conn.Read(buf)
	require.NoError(t, err)
	require.Greater(t, n, 0)

	select {
	case <-svc.started:
	case <-time.After(2 * time.Second):
		t.Fatal("service handler never started")
	}

	// Close without a request body terminator: the FIN cancels the stream.
	require.NoError(t, conn.Close())

	select {
	case <-svc.returned:
	case <-time.After(2 * time.Second):
		t.Fatal("handler context was not cancelled on clean close")
	}
}

// failWriter implements http.Flusher so ResponseController can flush it.
// Writes succeed failAfter times, then fail as a dead peer would.
type failWriter struct {
	hdr       http.Header
	mu        sync.Mutex
	writes    int
	failAfter int
}

func (f *failWriter) Header() http.Header {
	if f.hdr == nil {
		f.hdr = make(http.Header)
	}
	return f.hdr
}

func (f *failWriter) WriteHeader(int) {}
func (f *failWriter) Flush()          {}

func (f *failWriter) Write(p []byte) (int, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.writes++
	if f.failAfter > 0 && f.writes > f.failAfter {
		return 0, errors.New("write: broken pipe")
	}
	return len(p), nil
}

// TestKeepAliveCancelsHandlerOnWriteError proves case 2: a wrapped
// ResponseWriter that reports a write error without signalling the HTTP
// server still unblocks the handler, because keepAlive cancels the ctx.
func TestKeepAliveCancelsHandlerOnWriteError(t *testing.T) {
	// Not t.Parallel: the interval is a package global shared with the
	// streaming tests, which rely on the default.
	old := proto.StreamKeepAliveInterval
	proto.StreamKeepAliveInterval = 50 * time.Millisecond
	t.Cleanup(func() { proto.StreamKeepAliveInterval = old })

	svc := &blockingChatServer{started: make(chan struct{}), returned: make(chan struct{})}
	handler := proto.NewChatServer(svc)

	req := httptest.NewRequest(http.MethodPost, "/v1/Chat/SubscribeMessages", strings.NewReader(`{"username":"test"}`))
	req.Header.Set("Content-Type", "application/json")

	// First write is the stream's initial ping; the next ping fails.
	w := &failWriter{failAfter: 1}

	go handler.ServeHTTP(w, req)

	select {
	case <-svc.returned:
	case <-time.After(2 * time.Second):
		t.Fatal("handler did not return after keepalive write error")
	}

	require.ErrorIs(t, svc.ctxErr, context.Canceled)
}

// TestListenAcceptsTCP proves the example listener is a live TCP listener.
func TestListenAcceptsTCP(t *testing.T) {
	ln, err := listen("127.0.0.1:0")
	require.NoError(t, err)
	defer ln.Close()

	conn, err := net.Dial("tcp", ln.Addr().String())
	require.NoError(t, err)
	require.NoError(t, conn.Close())
}
