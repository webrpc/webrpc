package main

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"sync"
	"time"

	"github.com/webrpc/webrpc/_example/golang-sse/proto"
)

type ChatServer struct {
	mu            sync.Mutex
	lastId        uint64
	subscriptions map[uint64]chan *proto.Message
}

func NewChatServer() *ChatServer {
	return &ChatServer{
		subscriptions: map[uint64]chan *proto.Message{},
	}
}

func (s *ChatServer) SendMessage(ctx context.Context, username string, text string) error {
	msg := &proto.Message{
		Id:       uint64(rand.Uint64()),
		Username: username,
		Text:     text,
	}

	slog.Info("broadcasting message",
		"author", msg.Username,
		"text", msg.Text,
		"subscribers", len(s.subscriptions),
	)
	for _, sub := range s.subscriptions {
		sub := sub
		go func() {
			sub <- msg
		}()
	}

	return nil
}

func (s *ChatServer) SubscribeMessages(ctx context.Context, username string, stream proto.SubscribeMessagesStreamWriter) error {
	maxConnectionDuration := 30 * time.Minute
	ctx, _ = context.WithTimeout(ctx, maxConnectionDuration)

	msgs := make(chan *proto.Message, 10)
	defer s.unsubscribe(s.subscribe(msgs))

	for {
		select {
		case <-ctx.Done():
			switch err := ctx.Err(); err {
			case context.Canceled:
				return fmt.Errorf("client disconnected")
			default:
				return proto.ErrConnectionTooLong.WithCause(fmt.Errorf("timed out after %vs", maxConnectionDuration))
			}

		case msg := <-msgs:
			if err := stream.Write(msg); err != nil {
				return err
			}
		}
	}
}

func (s *ChatServer) subscribe(c chan *proto.Message) uint64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := s.lastId
	s.subscriptions[id] = c
	s.lastId++

	return id
}

func (s *ChatServer) unsubscribe(subscriptionId uint64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.subscriptions, subscriptionId)
}
