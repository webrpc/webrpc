package server

import (
	"context"
	"fmt"
	"io"
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
	return fmt.Errorf("internal error")
}

func (c *TestServer) GetOne(ctx context.Context) (*Simple, error) {
	return &fixtureOne, nil
}

func (c *TestServer) SendOne(ctx context.Context, one *Simple) error {
	if !cmp.Equal(&fixtureOne, one) {
		return ErrorWithCause(ErrUnexpectedValue, fmt.Errorf("%q:\n%s", "one", cmp.Diff(&fixtureOne, one)))
	}

	return nil
}

func (c *TestServer) GetMulti(ctx context.Context) (*Simple, *Simple, *Simple, error) {
	return &fixtureOne, &fixtureTwo, &fixtureThree, nil
}

func (c *TestServer) SendMulti(ctx context.Context, one, two, three *Simple) error {
	if !cmp.Equal(&fixtureOne, one) {
		return ErrorWithCause(ErrUnexpectedValue, fmt.Errorf("%q:\n%s", "one", cmp.Diff(&fixtureOne, one)))
	}
	if !cmp.Equal(&fixtureTwo, two) {
		return ErrorWithCause(ErrUnexpectedValue, fmt.Errorf("%q:\n%s", "two", cmp.Diff(&fixtureTwo, two)))
	}
	if !cmp.Equal(&fixtureThree, three) {
		return ErrorWithCause(ErrUnexpectedValue, fmt.Errorf("%q:\n%s", "three", cmp.Diff(&fixtureThree, three)))
	}

	return nil
}

func (c *TestServer) GetComplex(ctx context.Context) (*Complex, error) {
	return &fixtureComplex, nil
}

func (c *TestServer) SendComplex(ctx context.Context, complex *Complex) error {
	if !cmp.Equal(&fixtureComplex, complex) {
		return ErrorWithCause(ErrUnexpectedValue, fmt.Errorf("%q:\n%s", "complex", cmp.Diff(&fixtureComplex, complex)))
	}

	return nil
}

func (c *TestServer) GetEnumList(ctx context.Context) ([]Status, error) {
	return fixtureEnums.List, nil
}

func (c *TestServer) GetEnumMap(ctx context.Context) (map[Access]uint64, error) {
	return fixtureEnums.Dict, nil
}

func (c *TestServer) GetSchemaError(ctx context.Context, code int) error {
	switch code {
	case 0:
		return fmt.Errorf("failed to read file: %w", io.ErrUnexpectedEOF)
	case 1:
		return ErrorWithCause(ErrUnauthorized, fmt.Errorf("failed to verify JWT token"))
	case 2:
		return ErrExpiredToken
	case 3:
		return ErrInvalidToken
	case 4:
		return ErrDeactivated
	case 5:
		return ErrConfirmAccount
	case 6:
		return ErrAccessDenied
	case 7:
		return ErrMissingArgument
	case 8:
		return ErrUnexpectedValue
	case 100:
		return ErrorWithCause(ErrRateLimited, fmt.Errorf("1000 req/min exceeded"))
	case 101:
		return ErrDatabaseDown
	case 102:
		return ErrElasticDown
	case 103:
		return ErrNotImplemented
	case 200:
		return ErrUserNotFound
	case 201:
		return ErrUserBusy
	case 202:
		return ErrInvalidUsername
	case 300:
		return ErrFileTooBig
	case 301:
		return ErrFileInfected
	case 302:
		return ErrorWithCause(ErrFileType, fmt.Errorf(".wav is not supported"))
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
		"2": float64(24), // Go JSON unmarshaler uses float64 for numbers by default.
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
		Status:            Status_AVAILABLE,
	}

	fixtureEnums = EnumData{
		Dict: map[Access]uint64{Access_READ: 1, Access_WRITE: 2},
		List: []Status{Status_AVAILABLE, Status_NOT_AVAILABLE},
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

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		select {
		case <-srv.closed:

		case <-ctx.Done():
			// 1s graceful shutdown
			gracefulShutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			srv.err = srv.Shutdown(gracefulShutdownCtx)
			close(srv.closed)
		}
	}()

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
