package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

type TestServer struct{}

func (c *TestServer) GetComplex(ctx context.Context) (*ComplexType, error) {
	status := Status_AVAILABLE

	return &ComplexType{
		Meta:              c.GetMeta(),
		MetaNestedExample: c.GetMetaNestedExample(),
		NamesList:         c.GetNamesList(),
		NumsList:          c.GetNumsList(),
		DoubleArray:       c.GetDoubleArray(),
		ListOfMaps:        c.ListOfMaps(),
		ListOfUsers:       c.ListOfUsers(),
		MapOfUsers:        c.MapOfUsers(),
		User:              c.GetUser(),
		Enum:              &status,
	}, nil
}

func (*TestServer) GetMeta() map[string]interface{} {
	return map[string]interface{}{
		"1": "23",
		"2": 24,
	}
}

func (*TestServer) GetMetaNestedExample() map[string]map[string]uint32 {
	return map[string]map[string]uint32{
		"1": {
			"2": 1,
		},
	}
}

func (*TestServer) GetNamesList() []string {
	return []string{"John", "Alice", "Jakob"}
}

func (*TestServer) GetNumsList() []int64 {
	return []int64{1, 2, 3, 4534643543}
}

func (*TestServer) GetDoubleArray() [][]string {
	return [][]string{{"testing"}, {"api"}}
}

func (*TestServer) ListOfMaps() []map[string]uint32 {
	return []map[string]uint32{
		{
			"john":  1,
			"alice": 2,
			"Jakob": 251,
		},
	}
}

func (*TestServer) ListOfUsers() []*User {
	return []*User{
		{
			ID:       1,
			Username: "John-Doe",
			Role:     "admin",
		},
	}
}

func (*TestServer) MapOfUsers() map[string]*User {
	return map[string]*User{
		"admin": {
			ID:       1,
			Username: "John-Doe",
			Role:     "admin",
		},
	}
}

func (*TestServer) GetUser() *User {
	return &User{
		ID:       1,
		Username: "John-Doe",
		Role:     "admin",
	}
}

func (c *TestServer) SendComplex(ctx context.Context, complexType *ComplexType) error {
	if _, ok := complexType.Meta["1"]; !ok {
		return Errorf(ErrInvalidArgument, "meta: key %q does not exist", 1)
	}
	if complexType.Meta["1"] != "23" {
		return Errorf(ErrInvalidArgument, "meta: value in key '1' is not equal to '23'")
	}

	if _, ok := complexType.Meta["2"]; !ok {
		return Errorf(ErrInvalidArgument, "meta: key %q does not exist", 2)
	}
	if complexType.Meta["2"] != float64(24) {
		return Errorf(ErrInvalidArgument, "meta: value in key '2' is not equal to '24'")
	}

	metaNestedExample := c.GetMetaNestedExample()
	if _, ok := metaNestedExample["1"]; !ok {
		return Errorf(ErrInvalidArgument, "meta nested: key [%q] does not exist in meta nested structure", 1)
	}
	if _, ok := metaNestedExample["1"]["2"]; !ok {
		return Errorf(ErrInvalidArgument, "meta nested: key [1][%q] does not exist", 2)
	}
	if metaNestedExample["1"]["2"] != 1 {
		return Errorf(ErrInvalidArgument, "meta nested: expected value 1, actual: %v", metaNestedExample["1"]["2"])
	}

	namesList := c.GetNamesList()
	for index, val := range complexType.NamesList {
		if val != namesList[index] {
			return Errorf(ErrInvalidArgument, "names list: expected value %q, given %q", val, namesList[index])
		}
	}

	numsList := c.GetNumsList()
	for index, val := range complexType.NumsList {
		if val != numsList[index] {
			return Errorf(ErrInvalidArgument, "nums list: expected value %q, given %q", val, numsList[index])
		}
	}

	doubleArray := c.GetDoubleArray()
	for i, arr := range complexType.DoubleArray {
		if arr[0] != doubleArray[i][0] {
			return Errorf(ErrInvalidArgument, "double array: expected value %q, given %q", arr[0], doubleArray[i][0])
		}
	}

	listOfMaps := c.ListOfMaps()
	for i, m := range complexType.ListOfMaps {
		if m["john"] != listOfMaps[i]["john"] {
			return Errorf(ErrInvalidArgument, "list of maps: expected value %q, given %q", m["john"], listOfMaps[i]["john"])
		}
		if m["alice"] != listOfMaps[i]["alice"] {
			return Errorf(ErrInvalidArgument, "list of maps: expected value %q, given %q", m["alice"], listOfMaps[i]["alice"])
		}
		if m["Jakob"] != listOfMaps[i]["Jakob"] {
			return Errorf(ErrInvalidArgument, "list of maps: expected value %q, given %q", m["Jakob"], listOfMaps[i]["Jakob"])
		}
	}

	listOfUsers := c.ListOfUsers()
	if complexType.ListOfUsers[0].ID != listOfUsers[0].ID {
		return Errorf(ErrInvalidArgument, "list of users: expected value %q, given %q", listOfUsers[0].ID, complexType.ListOfUsers[0].ID)
	}
	if complexType.ListOfUsers[0].Role != listOfUsers[0].Role {
		return Errorf(ErrInvalidArgument, "list of users: expected value %q, given %q", listOfUsers[0].Role, complexType.ListOfUsers[0].Role)
	}
	if complexType.ListOfUsers[0].Username != listOfUsers[0].Username {
		return Errorf(ErrInvalidArgument, "list of users: expected value %q, given %q", listOfUsers[0].Username, complexType.ListOfUsers[0].Username)
	}

	mapOfUsers := c.MapOfUsers()
	if _, ok := complexType.MapOfUsers["admin"]; !ok {
		return Errorf(ErrInvalidArgument, "map of users: expected 'admin' key to exist in map")
	}
	if complexType.MapOfUsers["admin"].ID != mapOfUsers["admin"].ID {
		return Errorf(ErrInvalidArgument, "map of users: expected value %q, given %q", mapOfUsers["admin"].ID, complexType.MapOfUsers["admin"].ID)
	}
	if complexType.MapOfUsers["admin"].Role != mapOfUsers["admin"].Role {
		return Errorf(ErrInvalidArgument, "map of users: expected value %q, given %q", mapOfUsers["admin"].Role, complexType.MapOfUsers["admin"].Role)
	}
	if complexType.MapOfUsers["admin"].Username != mapOfUsers["admin"].Username {
		return Errorf(ErrInvalidArgument, "map of users: expected value %q, given %q", mapOfUsers["admin"].Username, complexType.MapOfUsers["admin"].Username)
	}

	user := c.GetUser()
	if complexType.User.ID != user.ID {
		return Errorf(ErrInvalidArgument, "user: expected value %q, given %q", complexType.User.ID, user.ID)
	}
	if complexType.User.Role != user.Role {
		return Errorf(ErrInvalidArgument, "user: expected value %q, given %q", complexType.User.Role, user.Role)
	}
	if complexType.User.Username != user.Username {
		return Errorf(ErrInvalidArgument, "user: expected value %q, given %q", complexType.User.Username, user.Username)
	}

	if complexType.Enum.String() != Status_AVAILABLE.String() {
		return Errorf(ErrInvalidArgument, "enum: expected value %q, given %q", Status_AVAILABLE, complexType.Enum.String())
	}

	return nil
}

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
