package gen

const goStruct = `
{{range .Messages}}
	{{if .Type eq "struct" }}
		type {{.Name}} struct {
			{{range .Fields}}
				{{.Name}} {{.Type | fieldType}} {{.Tags | fieldTags}}
			{{end}}
		}
	{{end}}
{{end}}
`

const goClient = `
// Client

{{range .Services}}
const {{.Name | constPathPrefix}} = "/rpc/{{.Name}}/"
{{end}}

{{range .Services}}
	{{ $serviceName := .Name | clientServiceName}}
	type {{$serviceName}} struct {

	}

	func {{.Name | newClientServiceName }}(addr string, client HTTPClient) {{.Name}} {
		prefix := urlBase(addr) + ExampleServicePathPrefix
		urls := [{{.Methods | countMethods}}]string{
			{{range .Methods}}
			prefix + "{{.Name}}",
			{{end}}
		}
		return &exampleServiceClient{
			client: client,
			urls:   urls,
		}
	}

	{{range .Methods}}
		func (c *{{$serviceName}}) {{.Name}}({{.Inputs | methodInputs}}) {{.Outputs | methodOutputs }} {
			//
		}
	{{end}}
{{end}}
`

const goBaseCode = `
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
	Status int    ` + "`json:\"status\"`" + `
	Code   string ` + "`json:\"code\"`" + `
	Msg    string ` + "`json:\"msg\"`" + `
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
`

const goServiceInterface = `
{{range .Services}}
	type {{.Name}} interface {
		{{range .Methods}}
			{{.Name}}({{.Inputs | methodInputs}}) {{.Outputs | methodOutputs}}
		{{end}}
	}
{{end}}
`

const goServices = `
// Server
{{range .Services}}
	type {{.Name}} struct {

	}

	func New{{.Name}}Service() {{.Name}} {

	}

	{{range .Methods}}
		func (c *{{.Name}}) {{.Name | serviceMethodName}}(ctx context.Context, w http.ResponseWriter, r *http.Request) {
			ctx = webrpc.WithMethodName(ctx, "{{.Name}}")

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
	{{end}}
{{end}}
`

const goTemplate = `
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
	"github.com/webrpc/webrpc-go"
)

` + goStruct + `
` + goServiceInterface + `
` + goClient + `
` + goServices + `
` + goBaseCode + `
`
