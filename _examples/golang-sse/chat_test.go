package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	client Chat
)

func TestMain(m *testing.M) {
	port := 4849

	go func() {
		if err := startServer(port); err != nil {
			log.Fatal("startServer():", err)
		}
	}()

	client = NewChatClient(fmt.Sprintf("http://0.0.0.0:%v", port), &http.Client{
		Timeout: time.Duration(2 * time.Second),
	})
	time.Sleep(time.Millisecond * 500)

	os.Exit(m.Run())
}

func TestStream(t *testing.T) {
	resp, err := client.GetNewMessages(context.Background())
	require.NoError(t, err)
	assert.NotNil(t, resp)
}
