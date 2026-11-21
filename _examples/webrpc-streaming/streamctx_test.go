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

// flushFailWriter accepts every write, but FlushError fails after failAfter flushes.
type flushFailWriter struct {
	hdr       http.Header
	mu        sync.Mutex
	flushes   int
	failAfter int
}

func (f *flushFailWriter) Header() http.Header {
	if f.hdr == nil {
		f.hdr = make(http.Header)
	}
	return f.hdr
}

func (f *flushFailWriter) WriteHeader(int)             {}
func (f *flushFailWriter) Write(p []byte) (int, error) { return len(p), nil }

func (f *flushFailWriter) FlushError() error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.flushes++
	if f.flushes > f.failAfter {
		return errors.New("flush: broken pipe")
	}
	return nil
}

// noFlushWriter supports neither Flush nor FlushError.
type noFlushWriter struct {
	hdr    http.Header
	status int
}

func (n *noFlushWriter) Header() http.Header {
	if n.hdr == nil {
		n.hdr = make(http.Header)
	}
	return n.hdr
}

func (n *noFlushWriter) WriteHeader(status int)      { n.status = status }
func (n *noFlushWriter) Write(p []byte) (int, error) { return len(p), nil }

func newSubscribeRequest() *http.Request {
	req := httptest.NewRequest(http.MethodPost, "/v1/Chat/SubscribeMessages", strings.NewReader(`{"username":"test"}`))
	req.Header.Set("Content-Type", "application/json")
	return req
}

// serveAndWait runs the handler until both the service method and ServeHTTP return.
func serveAndWait(t *testing.T, svc *blockingChatServer, w http.ResponseWriter) {
	t.Helper()

	served := make(chan struct{})
	go func() {
		defer close(served)
		proto.NewChatServer(svc).ServeHTTP(w, newSubscribeRequest())
	}()

	select {
	case <-svc.returned:
	case <-time.After(2 * time.Second):
		t.Fatal("handler did not return after keepalive failure")
	}
	<-served

	require.ErrorIs(t, svc.ctxErr, context.Canceled)
}

func withKeepAliveInterval(t *testing.T, d time.Duration) {
	// Not t.Parallel: the interval is a package global shared with the
	// streaming tests, which rely on the default.
	old := proto.StreamKeepAliveInterval
	proto.StreamKeepAliveInterval = d
	t.Cleanup(func() { proto.StreamKeepAliveInterval = old })
}

// TestKeepAliveCancelsHandlerOnWriteError proves case 2: a wrapped
// ResponseWriter that reports a write error without signalling the HTTP
// server still unblocks the handler, because keepAlive cancels the ctx.
func TestKeepAliveCancelsHandlerOnWriteError(t *testing.T) {
	withKeepAliveInterval(t, 50*time.Millisecond)

	svc := &blockingChatServer{started: make(chan struct{}), returned: make(chan struct{})}

	// First write is the stream's initial ping; the next ping fails.
	serveAndWait(t, svc, &failWriter{failAfter: 1})
}

// TestKeepAliveCancelsHandlerOnFlushError proves a failed flush on a dead peer
// unblocks the handler.
func TestKeepAliveCancelsHandlerOnFlushError(t *testing.T) {
	withKeepAliveInterval(t, 50*time.Millisecond)

	svc := &blockingChatServer{started: make(chan struct{}), returned: make(chan struct{})}

	// First flush is the stream's initial ping; the next ping fails.
	serveAndWait(t, svc, &flushFailWriter{failAfter: 1})
}

// TestNonFlushableWriterFailsBeforeHeaders proves misconfiguration yields a
// clear 500, not a 200 followed by an error frame.
func TestNonFlushableWriterFailsBeforeHeaders(t *testing.T) {
	svc := &blockingChatServer{started: make(chan struct{}), returned: make(chan struct{})}
	w := &noFlushWriter{}

	proto.NewChatServer(svc).ServeHTTP(w, newSubscribeRequest())

	require.Equal(t, http.StatusInternalServerError, w.status)
	select {
	case <-svc.started:
		t.Fatal("service method ran on a non-flushable writer")
	default:
	}
}
