package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/webrpc/webrpc/_example/golang-sse/proto"
)

type ChatServer struct {
	mu        sync.Mutex
	msgId     uint64
	lastSubId uint64
	subs      map[uint64]chan *proto.Message
}

func NewChatServer() *ChatServer {
	return &ChatServer{
		subs: map[uint64]chan *proto.Message{},
	}
}

func (s *ChatServer) SendMessage(ctx context.Context, username string, text string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	msg := &proto.Message{
		Id:        s.msgId,
		Username:  username,
		Text:      text,
		CreatedAt: time.Now(),
	}
	s.msgId++

	for _, sub := range s.subs {
		sub := sub
		go func() {
			sub <- msg
		}()
	}

	return nil
}

func (s *ChatServer) SubscribeMessages(ctx context.Context, username string, stream proto.SubscribeMessagesStreamWriter) error {
	if username == "" {
		return proto.ErrEmptyUsername
	}

	s.SendMessage(ctx, "SYSTEM", fmt.Sprintf("%v joined", username))
	defer s.SendMessage(ctx, "SYSTEM", fmt.Sprintf("%v left", username))

	msgs := make(chan *proto.Message, 10)
	defer s.unsubscribe(s.subscribe(msgs))

	for {
		select {
		case <-ctx.Done():
			switch err := ctx.Err(); err {
			case context.Canceled:
				return proto.ErrWebrpcClientDisconnected
			default:
				return proto.ErrWebrpcInternalError
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

	id := s.lastSubId
	s.subs[id] = c
	s.lastSubId++

	return id
}

func (s *ChatServer) unsubscribe(subscriptionId uint64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.subs, subscriptionId)
}
