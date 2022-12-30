package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/google/go-cmp/cmp"
)

type TestServer struct{}

func (c *TestServer) GetEmpty(ctx context.Context) error {
	return nil
}

func (c *TestServer) GetError(ctx context.Context) error {
	return ErrorInternal("internal error")
}

func (c *TestServer) GetOne(ctx context.Context) (*Simple, error) {
	return &fixtureOne, nil
}

func (c *TestServer) SendOne(ctx context.Context, one *Simple) error {
	if !cmp.Equal(&fixtureOne, one) {
		return Errorf(ErrInvalidArgument, "%q:\n%s", "one", cmp.Diff(&fixtureOne, one))
	}

	return nil
}

func (c *TestServer) GetMulti(ctx context.Context) (*Simple, *Simple, *Simple, error) {
	return &fixtureOne, &fixtureTwo, &fixtureThree, nil
}

func (c *TestServer) SendMulti(ctx context.Context, one, two, three *Simple) error {
	if !cmp.Equal(&fixtureOne, one) {
		return Errorf(ErrInvalidArgument, "%q:\n%s", "one", cmp.Diff(&fixtureOne, one))
	}
	if !cmp.Equal(&fixtureTwo, two) {
		return Errorf(ErrInvalidArgument, "%q:\n%s", "two", cmp.Diff(&fixtureTwo, two))
	}
	if !cmp.Equal(&fixtureThree, three) {
		return Errorf(ErrInvalidArgument, "%q:\n%s", "three", cmp.Diff(&fixtureThree, three))
	}

	return nil
}

func (c *TestServer) GetComplex(ctx context.Context) (*Complex, error) {
	return &fixtureComplex, nil
}

func (c *TestServer) SendComplex(ctx context.Context, complex *Complex) error {
	if !cmp.Equal(&fixtureComplex, complex) {
		return Errorf(ErrInvalidArgument, "%q:\n%s", "complex", cmp.Diff(&fixtureComplex, complex))
	}

	return nil
}

// Fixtures
var (
	fixtureOne = Simple{
		Id:   1,
		Name: "one",
	}
	fixtureTwo = Simple{
		Id:   2,
		Name: "two",
	}
	fixtureThree = Simple{
		Id:   3,
		Name: "three",
	}

	meta = map[string]interface{}{
		"1": "23",
		"2": 24,
	}
	metaNested = map[string]map[string]uint32{
		"1": {
			"2": 1,
		},
	}
	namesList   = []string{"John", "Alice", "Jakob"}
	numsList    = []int64{1, 2, 3, 4534643543}
	doubleArray = [][]string{{"testing"}, {"api"}}
	listOfMaps  = []map[string]uint32{
		{
			"john":  1,
			"alice": 2,
			"Jakob": 251,
		},
	}
	listOfUsers = []*User{
		{
			ID:       1,
			Username: "John-Doe",
			Role:     "admin",
		},
	}
	mapOfUsers = map[string]*User{
		"admin": {
			ID:       1,
			Username: "John-Doe",
			Role:     "admin",
		},
	}
	user = &User{
		ID:       1,
		Username: "John-Doe",
		Role:     "admin",
	}
	available = Status_AVAILABLE

	fixtureComplex = Complex{
		Meta:              meta,
		MetaNestedExample: metaNested,
		NamesList:         namesList,
		NumsList:          numsList,
		DoubleArray:       doubleArray,
		ListOfMaps:        listOfMaps,
		ListOfUsers:       listOfUsers,
		MapOfUsers:        mapOfUsers,
		User:              user,
		Enum:              &available,
	}
)

func RunTestServer(addr string, timeout time.Duration) (*testServer, error) {
	srv := &testServer{
		Server: &http.Server{
			Addr:    addr,
			Handler: NewTestApiServer(&TestServer{}),
		},
		closed: make(chan struct{}),
	}

	if timeout > 0 {
		go func() {
			timeoutCtx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			select {
			case <-srv.closed:

			case <-timeoutCtx.Done():
				gracefulShutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()

				srv.err = srv.Shutdown(gracefulShutdownCtx)
				close(srv.closed)
			}
		}()
	}

	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to bind %v: %w", addr, err)
	}

	go srv.Serve(l)

	return srv, nil
}

type testServer struct {
	*http.Server
	closed chan struct{}
	err    error
}

func (srv *testServer) Wait() error {
	<-srv.closed
	return srv.err
}

func (srv *testServer) Close() error {
	gracefulShutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return srv.Shutdown(gracefulShutdownCtx)
}
