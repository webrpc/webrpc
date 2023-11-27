package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/webrpc/webrpc/_example/golang-sse/proto"
)

type ChatServer struct {
}

func (s *ChatServer) SendMessage(ctx context.Context, authorName string, text string) error {
	return nil
}

func (s *ChatServer) SubscribeMessages(ctx context.Context, serverTimeoutSec int, stream proto.SubscribeMessagesStreamWriter) error {
	if serverTimeoutSec > 0 {
		ctx, _ = context.WithTimeout(ctx, time.Duration(serverTimeoutSec)*time.Second)
	}

	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			switch err := ctx.Err(); err {
			case context.Canceled:
				return fmt.Errorf("client disconnected")
			default:
				return proto.ErrConnectionTooLong.WithCause(fmt.Errorf("timed out after %vs", serverTimeoutSec))
			}

		case <-time.After(time.Duration(rand.Intn(1000)) * time.Millisecond):
			// Simulate work. Delay each message by 0-1000ms.
		}

		msg := &proto.Message{
			Id:         rand.Uint64(),
			AuthorName: "Alice",
			Text:       fmt.Sprintf("Message %v", i),
		}

		if err := stream.Write(msg); err != nil {
			return err
		}
	}

	return nil
}
