package main

import (
	"context"
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

	client = NewExampleServiceClient("http://0.0.0.0:4242", &http.Client{
		Timeout: time.Duration(20 * time.Second),
	})
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

// TODO: lets do specific http1.1 test and another http2 test

func TestDownload(t *testing.T) {
	{
		stream, err := client.Download(context.Background(), "hi")
		assert.NoError(t, err)

		// time.Sleep(100 * time.Millisecond)

		// for {
		// 	select {
		// 	case <-reader.Done():
		// 		break
		// 	default:
		// 	}

		//
		// stream.Next()
		// stream.Read() ..
		// check again io.EOF ? or .Done() with chan thing.. maybe..

		for {
			resp, err := stream.Read()
			if err != nil {
				t.Fatal(err)
				panic("test-error")
			}
			fmt.Println("=> resp:", resp)
			if resp == "" {
				fmt.Println("emptym, done", err)
				break
			}
			// time.Sleep(3500 * time.Millisecond)
		}

		/*
			base64, err := reader.Read() // returns (base64 string, err error)
			assert.NoError(t, err)
			fmt.Println("==> a", base64)
			// time.Sleep(100 * time.Millisecond)

			base64, err = reader.Read() // returns (base64 string, err error)
			assert.NoError(t, err)
			fmt.Println("==> b", base64)
			// time.Sleep(100 * time.Millisecond)

			base64, err = reader.Read() // returns (base64 string, err error)
			assert.NoError(t, err)
			fmt.Println("==> c", base64)
			// time.Sleep(100 * time.Millisecond)
		*/

		// TODO: query until EOF..
		// or a disconnect..

		// }
	}
}
