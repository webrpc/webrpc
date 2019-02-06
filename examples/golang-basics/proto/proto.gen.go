package proto

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/webrpc/webrpc/lib/webrpc-go"
)

// TODO: this file should be code-generated from the proto.json file
// we write it by hand first to get a sense of what we're designing for.
// .. with ideas+help from twirp/grpc

type GetUserRequest struct {
	UserID uint64 `json:"userID"`
}

type User struct {
	ID        uint64     `json:"id" db:"id"`
	Username  string     `json:"USERNAME" db:"username"`
	CreatedAt *time.Time `json:"created_at,omitempty" db:"created_at"`
}

//--

type ExampleService interface {
	Ping(context.Context) (*bool, error)
	GetUser(context.Context, *GetUserRequest) (*User, error)
}

var Services = map[string][]string{
	"ExampleService": {"Ping", "GetUser"},
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
	// otherwise, we'd make an Empty message like protobuf people do, which is an empty object, and take it as input
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
		err := webrpc.Errorf(webrpc.ErrBadRoute, "unsupported method %q (only POST is allowed)", r.Method)
		writeJSONError(ctx, w, r, err)
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
		err := webrpc.Errorf(webrpc.ErrBadRoute, "no handler for path %q", r.URL.Path)
		writeJSONError(ctx, w, r, err)
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
		err := webrpc.Errorf(webrpc.ErrBadRoute, "unexpected Content-Type: %q", r.Header.Get("Content-Type"))
		writeJSONError(ctx, w, r, err)
	}
}

func (s *exampleServiceServer) servePingJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = webrpc.WithMethodName(ctx, "Ping")

	// NOTE: Ping method has no input arguments,
	// TODO: is that cool or should we enforce some at schema level?
	// lets try without for now, should be fine

	// Call service method
	var respContent *bool
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				writeJSONError(ctx, w, r, webrpc.ErrorInternal("internal service panic"))
				panic(rr)
			}
		}()
		respContent, err = s.ExampleService.Ping(ctx)
	}()

	if err != nil {
		writeJSONError(ctx, w, r, err)
		return
	}
	// if respContent == nil {
	// 	// s.writeError(ctx, resp, twirp.InternalError("received a nil *Empty and nil error while calling Ping. nil responses are not supported"))
	// 	return
	// }

	respBody, err := json.Marshal(respContent)
	if err != nil {
		err = webrpc.WrapError(webrpc.ErrInternal, err, "failed to marshal json response")
		writeJSONError(ctx, w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if n, err := w.Write(respBody); err != nil {
		_ = n
		_ = err
		// TODO: failing silently for now..
		// msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
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
		err := webrpc.Errorf(webrpc.ErrBadRoute, "unexpected Content-Type: %q", r.Header.Get("Content-Type"))
		writeJSONError(ctx, w, r, err)
	}
}

func (s *exampleServiceServer) serveGetUserJSON(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var err error
	ctx = webrpc.WithMethodName(ctx, "GetUser")

	reqContent := new(GetUserRequest)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = webrpc.WrapError(webrpc.ErrInternal, err, "failed to read request data")
		writeJSONError(ctx, w, r, err)
		return
	}
	defer r.Body.Close()
	err = json.Unmarshal(reqBody, reqContent)
	if err != nil {
		err = webrpc.WrapError(webrpc.ErrInternal, err, "failed to unmarshal request data")
		writeJSONError(ctx, w, r, err)
		return
	}

	// TODO: multiple args..?
	// perhaps, but need to name it and make implicit types essentially here..
	// for named args, and ordered args, its possible..
	// for now, just stick to single output, we can enforce
	// this at schema level for now if we want.. gRPC only allows single output
	// or they have you wrap multiple response types another message type for that response..
	// lets just do what gRPC does for now can iterate after

	// Call service method
	var respContent *User
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if rr := recover(); rr != nil {
				writeJSONError(ctx, w, r, webrpc.ErrorInternal("internal service panic"))
				panic(rr)
			}
		}()
		respContent, err = s.ExampleService.GetUser(ctx, reqContent)
	}()

	if err != nil {
		writeJSONError(ctx, w, r, err)
		return
	}
	// if respContent == nil {
	// 	// s.writeError(ctx, resp, twirp.InternalError("received a nil *Empty and nil error while calling Ping. nil responses are not supported"))
	// 	return
	// }

	respBody, err := json.Marshal(respContent)
	if err != nil {
		err = webrpc.WrapError(webrpc.ErrInternal, err, "failed to marshal json response")
		writeJSONError(ctx, w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if n, err := w.Write(respBody); err != nil {
		_ = n
		_ = err
		// TODO: failing silently for now..
		// msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
	}
}

//
// Helpers
//

// HTTPClient is the interface used by generated clients to send HTTP requests.
// It is fulfilled by *(net/http).Client, which is sufficient for most users.
// Users can provide their own implementation for special retry policies.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type WebRPCServer interface {
	http.Handler
	WebRPCVersion() string
	ServiceVersion() string
}

type errResponse struct {
	Status int    `json:"status"`
	Code   string `json:"code"`
	Msg    string `json:"msg"`
}

func writeJSONError(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	rpcErr, ok := err.(webrpc.Error)
	if !ok {
		rpcErr = webrpc.WrapError(webrpc.ErrInternal, err, "webrpc error")
	}

	statusCode := webrpc.HTTPStatusFromErrorCode(rpcErr.Code())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errResp := errResponse{
		Status: statusCode,
		Code:   string(rpcErr.Code()),
		Msg:    rpcErr.Error(),
	}
	respBody, _ := json.Marshal(errResp)
	w.Write(respBody)
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
	req.Header.Set("Accept", contentType)
	req.Header.Set("Content-Type", contentType)
	return req, nil
}

// doJSONRequest is common code to make a request to the remote twirp service.
func doJSONRequest(ctx context.Context, client HTTPClient, url string, in, out interface{}) error {
	// TODO: return webrpc.Error every chance we get, with the proper cause..

	reqBody, err := json.Marshal(in)
	if err != nil {
		return err
		// return clientError("failed to marshal json request", err)
	}
	// if err = ctx.Err(); err != nil {
	// 	return clientError("aborted because context was done", err)
	// }

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

	// if err = ctx.Err(); err != nil {
	// 	return clientError("aborted because context was done", err)
	// }

	if resp.StatusCode != 200 {

		// TODO ......
		var respErr errResponse
		respBody, _ := ioutil.ReadAll(resp.Body)
		_ = json.Unmarshal(respBody, &respErr)
		return errors.New(respErr.Msg)
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

	return nil
}
