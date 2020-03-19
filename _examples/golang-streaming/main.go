package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/pkg/errors"
)

func main() {
	err := startServer()
	if err != nil {
		log.Fatal(err)
	}
}

func startServer() error {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("."))
	})

	webrpcHandler := NewExampleServiceServer(&ExampleServiceRPC{})
	r.Handle("/*", webrpcHandler)

	return http.ListenAndServe(":4242", r)
}

type ExampleServiceRPC struct {
}

func (s *ExampleServiceRPC) Ping(ctx context.Context) error {
	return nil
}

func (s *ExampleServiceRPC) GetUser(ctx context.Context, userID uint64) (*User, error) {
	if userID == 911 {
		return nil, ErrorNotFound("fdf")
	}

	return &User{
		Id:       userID,
		Username: "hihi",
	}, nil
}

// func (s *ExampleServiceRPC) Upload(ctx context.Context, reader UploadReader) (bool, error) {
// 	return false, nil
// }

// UploadStream
// DownloadStream
// stream.Read() -- this works, cuz we can't have both kinds?
// stream.Write()

// TODO: it could return err
// which would close the conncetion etc...

func (s *ExampleServiceRPC) Download(ctx context.Context, file string, stream DownloadStreamWriter) error {
	// TODO: the middleware.Logger in chi isn't suitable here as well, it will be reading all of this, wrapping it etc.
	// and we don't want that.. ideally chi logger after amount of bytes stops tracking and drops count

	i := 0
	for {

		// TODO: add code that checks if the connection is open
		// ie, <-stream.Connected()
		// or something.. or <-stream.Disconnected() perhaps? or .Closed()
		// in case the client disconnects we must know this.
		// perhaps we do this test right inside of .Data() though, and return the err that client is gone, even cleaner..

		// select {
		// // case <-stream.ClientDisconnected():
		// case <-stream.Done():
		// 	fmt.Println("stream is done!")
		// 	return nil
		// 	// scenarios:
		// 	// 1. client disconnected unexpectedly but we're still sending stream? ok, not okay? .. they might reconnect, we can track them..
		// 	// 2. stream has been closed by us, which is fine.
		// default:
		// }

		// select {
		// case <-stream.Done():
		// 	fmt.Println("stream is done!")
		// 	return nil
		// default:
		// }

		fmt.Println("hi..")

		err := stream.Data(fmt.Sprintf("hiii send %d", i))
		if err == ErrStreamClosed {
			fmt.Println("client is gone.....")
			return nil
		}

		if err != nil {
			fmt.Println("ERR!!", err)
			return nil
		}
		time.Sleep(1 * time.Second)
		if i >= 6 {
			break
		}

		i += 1
	}

	// stream will close on its own when we return actually, nice.
	// we can close ourselves too if we want though
	// stream.Close()

	return nil

}

//--
// move this to the gen..

type httpStreamWriter struct {
	w             http.ResponseWriter
	headerWritten bool
	done          chan struct{}
	mu            sync.Mutex
}

// newHTTPStreamWriter() ? .. perhaps, set the flusher.. etc.. kinda nice.

func (s *httpStreamWriter) Write(payload []byte) error {
	select {
	case <-s.Done():
		return ErrStreamClosed
	default:
	}

	flusher, ok := s.w.(http.Flusher)
	if !ok {
		return errors.Errorf("expected http.ResponseWriter to be an http.Flusher")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	w := s.w
	// TODO: review, etc.
	if !s.headerWritten {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Transfer-Encoding", "chunked")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Cache-Control", "no-cache")
		s.headerWritten = true
		// TODO: this doesnt mean written, this is set..
		// perhaps we track number of flushes, etc. "n" then chcek if n == 0..?
	}

	// TODO: error check .....
	s.w.Write([]byte(fmt.Sprintf("%x\r\n", len(payload))))
	s.w.Write(payload)
	s.w.Write([]byte("\r\n"))

	fmt.Printf("SEND: %s\n", payload)

	flusher.Flush()

	return nil
}

func (s *httpStreamWriter) Error(err error) error {
	return nil
}

func (s *httpStreamWriter) Ping() error {
	return s.Write([]byte(`{"ping":true}`))
}

func (s *httpStreamWriter) Close() error {
	flusher, ok := s.w.(http.Flusher)
	if !ok {
		return errors.Errorf("expected http.ResponseWriter to be an http.Flusher")
	}

	select {
	case <-s.Done():
		return nil
	default:
	}

	s.mu.Lock()
	fmt.Fprintf(s.w, "0\r\n")
	flusher.Flush() // Trigger "chunked" encoding and send a chunk...
	close(s.done)
	s.mu.Unlock()

	return nil
}

func (s *httpStreamWriter) Done() <-chan struct{} {
	s.mu.Lock()
	if s.done == nil {
		s.done = make(chan struct{})
	}
	d := s.done
	s.mu.Unlock()
	return d
}

type downloadStreamWriter struct {
	httpStreamWriter
}

// var _ DownloadStreamWriter = &downloadStreamWriter{}

func (s *downloadStreamWriter) Data(base64 string) error {
	ret0 := base64

	respContent := struct {
		Ret0 string `json:"base64"`
	}{ret0}

	respBody, err := json.Marshal(respContent)
	if err != nil {
		err = WrapError(ErrInternal, err, "failed to marshal json response")
		RespondWithError(s.w, err)
		return nil //..
	}

	data := []byte(fmt.Sprintf(`{"data":%s}`, string(respBody)))
	return s.Write(data)
}
