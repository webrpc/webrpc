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

	url := "http://0.0.0.0:4242"
	// url := "https://pkgrok.0xhorizon.net"

	// client = NewExampleServiceClient(url, &http.Client{
	// 	Timeout: time.Duration(31 * time.Second),
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

// TODO: lets do specific http1.1 test and another http2 test

func TestDownload(t *testing.T) {
	{
		stream, err := client.Download(context.Background(), "hi")
		assert.NoError(t, err)

	loop:
		for {
			respBase64, err := stream.Read()

			switch err {
			case ErrStreamClosed:
				fmt.Println("we done.")
				break loop
			default:
				// some error,
				t.Fatal(err)
				panic(err)
			case nil:
				// no error
			}

			fmt.Println("=> resp:", respBase64)
		}

		respBase64, err := stream.Read()
		fmt.Println("=> ha,", respBase64)
		fmt.Println("=> err", err)

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

	}
}
