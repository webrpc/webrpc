package proto

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/webrpc/webrpc-go"
)

// TODO: this file should be code-generated from the proto.json file
// we write it by hand first to get a sense of what we're designing for.
// .. with ideas+help from twirp

type GetUserRequest struct {
	UserID uint64 `json:"userID"`
}

type User struct {
	ID       uint64 `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
}

/*
TODO:
=====

1. error responses / handling
2. json marshalling type schema mapping + code-generation -- inspired by TypeScript -- TS types to json marshalling?
3. optional types, and defaults.. ie. jsonpb from protobuf has `EmitDefaults: true`
4. fast json code-generation, look into ffjson, etc.
5. runtime map of service+methods for ExampleService, to get a lits of methods on service at runtime? could be useful for acl stuff..

*/

//--

type ExampleService interface {
	Ping(context.Context) (*bool, error)
	GetUser(context.Context, *GetUserRequest) (*User, error)
}

//
// Client
//

type exampleServiceClient struct {
	client HTTPClient
	urls   [2]string
}

func NewExampleServiceClient(addr string, client HTTPClient) ExampleService {
	prefix := urlBase(addr) + ExampleServicePathPrefix
	urls := [2]string{
		prefix + "Ping",
		prefix + "GetUser",
	}
	return &exampleServiceClient{
		client: client,
		urls:   urls,
	}
}

func (c *exampleServiceClient) Ping(ctx context.Context) (*bool, error) {
	// out := new(Empty)
	var out *bool // TODO: should we support primitives at all in the first version? or require a struct/object for each input/output like gRPC?
	err := doJSONRequest(ctx, c.client, c.urls[0], nil, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleServiceClient) GetUser(ctx context.Context, in *GetUserRequest) (*User, error) {
	out := new(User)
	err := doJSONRequest(ctx, c.client, c.urls[1], in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

//
// Server
//

const ExampleServicePathPrefix = "/rpc/ExampleService/"

type exampleServiceServer struct {
	ExampleService
}

func NewExampleServiceServer(svc ExampleService) WebRPCServer {
	return &exampleServiceServer{
		ExampleService: svc,
	}
}

func (s *exampleServiceServer) WebRPCVersion() string {
	return "v0.0.1"
}

func (s *exampleServiceServer) ServiceVersion() string {
	return "v0.1.0"
}

// TODO(future): create a ServiceDescriptor() []byte on ExampleService interface
// where we return the entire schema, so a client can test if it supports it easily
// or, maybe we just embed the schema json entirely, and also compute
// a hash of it and include it in both client/server? that can work too.
// but the schema json would need to be ordered/sorted, so testability is robust

func (s *exampleServiceServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = webrpc.WithResponseWriter(ctx, w)
	ctx = webrpc.WithServiceName(ctx, "ExampleService")

	if r.Method != "POST" {
		msg := fmt.Sprintf("unsupported method %q (only POST is allowed)", r.Method)
		// err = badRouteError(msg, req.Method, req.URL.Path)
		// s.writeError(ctx, resp, err)
		panic(msg) // TODO
		return
	}

	switch r.URL.Path {
	case "/rpc/ExampleService/Ping":
		s.servePing(ctx, w, r)
		return
	case "/rpc/ExampleService/GetUser":
		s.serveGetUser(ctx, w, r)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", r.URL.Path)
		// err = badRouteError(msg, req.Method, req.URL.Path)
		// s.writeError(ctx, resp, err)
		panic(msg) // TODO
		return
	}
}

func (s *exampleServiceServer) servePing(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.servePingJSON(ctx, w, r)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", r.Header.Get("Content-Type"))
		panic(msg) // TODO
		// twerr := badRouteError(msg, req.Method, req.URL.Path)
		// s.writeError(ctx, resp, twerr)
	}
}

func (s *exampleServiceServer) servePingJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = webrpc.WithMethodName(ctx, "Ping")

	// reqContent := new(Empty)
	// unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	// if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
	// 	err = wrapErr(err, "failed to parse request json")
	// 	s.writeError(ctx, resp, twirp.InternalErrorWith(err))
	// 	return
	// }

	// Call service method
	var respContent *bool
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if r := recover(); r != nil {
				// s.writeError(ctx, resp, twirp.InternalError("Internal service panic"))
				panic(r)
			}
		}()
		respContent, err = s.ExampleService.Ping(ctx)
	}()

	if err != nil {
		// s.writeError(ctx, w, err)
		panic(err) // TODO
		return
	}
	// if respContent == nil {
	// 	// s.writeError(ctx, resp, twirp.InternalError("received a nil *Empty and nil error while calling Ping. nil responses are not supported"))
	// 	return
	// }

	respBody, err := json.Marshal(respContent)
	if err != nil {
		// 	err = wrapErr(err, "failed to marshal json response")
		// 	s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		panic(err) // TODO
		return
	}

	// var buf bytes.Buffer
	// err = json.Marshal()
	// marshaler := &jsonpb.Marshaler{OrigName: true}
	// if err = marshaler.Marshal(&buf, respContent); err != nil {
	// 	err = wrapErr(err, "failed to marshal json response")
	// 	s.writeError(ctx, resp, twirp.InternalErrorWith(err))
	// 	return
	// }

	// ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if n, err := w.Write(respBody); err != nil {
		_ = n
		// msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		// twerr := twirp.NewError(twirp.Unknown, msg)
		panic(err) // TODO
	}
}

func (s *exampleServiceServer) serveGetUser(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}

	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveGetUserJSON(ctx, w, r)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", r.Header.Get("Content-Type"))
		panic(msg) // TODO
		// twerr := badRouteError(msg, req.Method, req.URL.Path)
		// s.writeError(ctx, resp, twerr)
	}
}

func (s *exampleServiceServer) serveGetUserJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = webrpc.WithMethodName(ctx, "GetUser")

	reqContent := new(GetUserRequest)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err) // TODO
	}
	defer r.Body.Close()
	err = json.Unmarshal(reqBody, reqContent)
	if err != nil {
		panic(err) // TODO......
	}

	// TODO: multiple args..?
	// perhaps, but need to name it and make implicit types essentially here..
	// for named args, and ordered args, its possible..

	// Call service method
	var respContent *User
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if r := recover(); r != nil {
				// s.writeError(ctx, resp, twirp.InternalError("Internal service panic"))
				panic(r)
			}
		}()
		respContent, err = s.ExampleService.GetUser(ctx, reqContent)
	}()

	if err != nil {
		// s.writeError(ctx, w, err)
		panic(err) // TODO
		return
	}
	// if respContent == nil {
	// 	// s.writeError(ctx, resp, twirp.InternalError("received a nil *Empty and nil error while calling Ping. nil responses are not supported"))
	// 	return
	// }

	respBody, err := json.Marshal(respContent)
	if err != nil {
		// 	err = wrapErr(err, "failed to marshal json response")
		// 	s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		panic(err) // TODO
		return
	}

	// var buf bytes.Buffer
	// err = json.Marshal()
	// marshaler := &jsonpb.Marshaler{OrigName: true}
	// if err = marshaler.Marshal(&buf, respContent); err != nil {
	// 	err = wrapErr(err, "failed to marshal json response")
	// 	s.writeError(ctx, resp, twirp.InternalErrorWith(err))
	// 	return
	// }

	// ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if n, err := w.Write(respBody); err != nil {
		_ = n
		// msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		// twerr := twirp.NewError(twirp.Unknown, msg)
		panic(err) // TODO
	}
}

//
// Helpers
//

// HTTPClient is the interface used by generated clients to send HTTP requests.
// It is fulfilled by *(net/http).Client, which is sufficient for most users.
// Users can provide their own implementation for special retry policies.
//
// HTTPClient implementations should not follow redirects. Redirects are
// automatically disabled if *(net/http).Client is passed to client
// constructors. See the withoutRedirects function in this file for more
// details.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type WebRPCServer interface {
	http.Handler
	WebRPCVersion() string
	ServiceVersion() string
}

// urlBase helps ensure that addr specifies a scheme. If it is unparsable
// as a URL, it returns addr unchanged.
func urlBase(addr string) string {
	// If the addr specifies a scheme, use it. If not, default to
	// http. If url.Parse fails on it, return it unchanged.
	url, err := url.Parse(addr)
	if err != nil {
		return addr
	}
	if url.Scheme == "" {
		url.Scheme = "http"
	}
	return url.String()
}

// newRequest makes an http.Request from a client, adding common headers.
func newRequest(ctx context.Context, url string, reqBody io.Reader, contentType string) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return nil, err
	}
	// req = req.WithContext(ctx)
	// if customHeader := getCustomHTTPReqHeaders(ctx); customHeader != nil {
	// 	req.Header = customHeader
	// }
	req.Header.Set("Accept", contentType)
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

// doJSONRequest is common code to make a request to the remote twirp service.
func doJSONRequest(ctx context.Context, client HTTPClient, url string, in, out interface{}) error {
	//  reqBody := bytes.NewBuffer(nil)
	// 	marshaler := &jsonpb.Marshaler{OrigName: true}
	// 	if err = marshaler.Marshal(reqBody, in); err != nil {
	// 		return clientError("failed to marshal json request", err)
	// 	}
	// 	if err = ctx.Err(); err != nil {
	// 		return clientError("aborted because context was done", err)
	// 	}

	reqBody, err := json.Marshal(in)
	if err != nil {
		return err
		// 		return clientError("failed to marshal json request", err)
	}
	// 	if err = ctx.Err(); err != nil {
	// 		return clientError("aborted because context was done", err)
	// 	}

	req, err := newRequest(ctx, url, bytes.NewBuffer(reqBody), "application/json")
	if err != nil {
		return err
		// return clientError("could not build request", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
		// return clientError("failed to do request", err)
	}

	defer func() {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = cerr
			// err = clientError("failed to close response body", cerr)
		}
	}()

	// 	if err = ctx.Err(); err != nil {
	// 		return clientError("aborted because context was done", err)
	// 	}

	if resp.StatusCode != 200 {
		return errors.New("todo..")
		// return errorFromResponse(resp)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// TODO
		return err
	}

	err = json.Unmarshal(respBody, &out)
	if err != nil {
		// TODO
		return err
	}

	// 	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	// 	if err = unmarshaler.Unmarshal(resp.Body, out); err != nil {
	// 		return clientError("failed to unmarshal json response", err)
	// 	}
	// 	if err = ctx.Err(); err != nil {
	// 		return clientError("aborted because context was done", err)
	// 	}

	return nil
}
