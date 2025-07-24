package main

import (
	"context"
	"errors"
	"fmt"
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

func (s *ChatServer) SubscribeMessages(ctx context.Context, username string, stream proto.SubscribeMessagesStreamWriter) error {
	if username == "" {
		return proto.ErrEmptyUsername
	}

	s.SendMessage(context.Background(), "", fmt.Sprintf("%v joined", username))
	defer s.SendMessage(context.Background(), "", fmt.Sprintf("%v left", username))

	sub := s.messages.Subscribe(ctx, &ringbuf.SubscribeOpts{Name: "sub1"})

	for msg := range sub.Seq {
		if err := stream.Write(msg); err != nil {
			return err
		}
	}

	if err := sub.Err(); !errors.Is(err, context.Canceled) {
		return proto.ErrWebrpcInternalError.WithCausef("failed to stream messages: %w", err)
	}

	return nil
}
