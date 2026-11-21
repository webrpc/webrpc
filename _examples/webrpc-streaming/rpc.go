package main

import (
	"cmp"
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/go-chi/httplog/v3"
	"github.com/golang-cz/devslog"
	"github.com/golang-cz/ringbuf"
	"github.com/webrpc/webrpc/_example/golang-sse/proto"
)

type ChatServer struct {
	mu       sync.Mutex
	msgId    uint64
	messages *ringbuf.RingBuffer[*proto.Message]
	logger   *slog.Logger
}

func NewChatServer() *ChatServer {
	return &ChatServer{
		messages: ringbuf.New[*proto.Message](10000),
		logger: slog.New(devslog.NewHandler(os.Stdout, &devslog.Options{
			HandlerOptions: &slog.HandlerOptions{
				ReplaceAttr: httplog.SchemaECS.Concise(true).ReplaceAttr,
			},
		})),
	}
}

func (s *ChatServer) close(ctx context.Context) {
	s.SendMessage(ctx, "SYSTEM", "Server shutting down...")
	s.messages.Close()
}

func (s *ChatServer) SendMessage(ctx context.Context, username string, text string) error {
	now := time.Now()

	s.mu.Lock()
	defer s.mu.Unlock()

	msg := &proto.Message{
		Id:        s.msgId,
		Username:  username,
		Text:      text,
		CreatedAt: now,
	}
	s.msgId++

	s.messages.Write(msg)

	return nil
}

func (s *ChatServer) SubscribeMessages(ctx context.Context, username string, lastMessageId *uint64, stream proto.SubscribeMessagesStreamWriter) error {
	if username == "" {
		return proto.ErrEmptyUsername
	}

	s.SendMessage(context.Background(), "", fmt.Sprintf("%v joined", username))
	defer s.SendMessage(context.Background(), "", fmt.Sprintf("%v left", username))

	sub := s.messages.Subscribe(ctx, &ringbuf.SubscribeOpts{
		Name:        username,
		StartBehind: s.messages.Size() / 2,
		MaxLag:      3 * s.messages.Size() / 4,
	})

	if lastMessageId != nil {
		// Reconnection logic.
		sub.SeekAfter(func(msg *proto.Message) int {
			return cmp.Compare(msg.Id, *lastMessageId)
		})
	}

	// ringbuf blocks in cond.Wait() and wakes on the next Write, not on ctx cancel.
	// Read aside, so a dead client returns (and publishes "left") right away.
	errCh := make(chan error, 1)
	go func() {
		for msg := range sub.Iter() {
			if err := stream.Write(msg); err != nil {
				errCh <- err
				return
			}
		}
		errCh <- sub.Err()
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errCh:
		if errors.Is(err, io.EOF) {
			return proto.ErrWebrpcStreamFinished.WithCausef("stream finished: %w", err)
		}
		return err
	}
}
