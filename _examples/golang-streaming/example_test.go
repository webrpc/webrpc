package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	client ExampleServiceClient
)

// func TestMain()

func init() {
	// go func() {
	// 	startServer()
	// }()

	url := "http://0.0.0.0:4242"
	// url := "https://pkgrok.0xhorizon.net"

	// client = NewExampleServiceClient(url, &http.Client{
	// 	Timeout: time.Duration(10 * time.Second),
	// })
	client = NewExampleServiceClient(url, &http.Client{})

	time.Sleep(time.Millisecond * 500)

}

func TestPing(t *testing.T) {
	err := client.Ping(context.Background())
	assert.NoError(t, err)
}

func TestGetUser(t *testing.T) {
	{
		user, err := client.GetUser(context.Background(), 12)
		assert.Equal(t, &User{Id: 12, Username: "hihi"}, user)
		assert.NoError(t, err)
	}

	{
		// Error case, expecting to receive an error
		user, err := client.GetUser(context.Background(), 911)
		fmt.Println(err.Error())
		// assert.True(t, IsErrorCode(err, ErrNotFound))
		assert.Nil(t, user)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "not found")
	}
}

func TestDownload(t *testing.T) {
	{
		stream, err := client.Download(context.Background(), "hi")
		assert.NoError(t, err)

		for {
			respBase64, err := stream.Read()

			if errors.Is(err, ErrStreamClosed) {
				fmt.Println("success. stream is done.")
				break
			}

			if errors.Is(err, ErrStreamLost) {
				fmt.Println("connection lost..", err)
				// 	// lets reconnect..
				// 	stream, err = client.Download(context.Background(), "hi")
				// 	if err != nil {
				// 		// ..
				// 	}
				// 	continue
				return
			}
			if err != nil {
				t.Fatal(err)
			}

			fmt.Println("=> resp:", respBase64)
		}

		respBase64, err := stream.Read()
		fmt.Println("=> ha,", respBase64)
		fmt.Println("=> err", err)

	}
}
