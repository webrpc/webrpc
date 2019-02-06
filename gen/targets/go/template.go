package gen

const goStruct = `
{{if .Messages}}
  {{range .Messages}}
    {{if .Type | isEnum}}
      {{$enumName := .Name}}
      {{range .Fields}}
        // {{$enumName}}_{{.Name}} = {{.Type}} {{.Value}}
      {{end}}
    {{end}}
    {{if .Type | isStruct  }}
      type {{.Name}} struct {
        {{range .Fields}}
          {{.Name | exportedField}} {{.Type | fieldType}} {{.Meta | fieldTags}}
        {{end}}
      }
    {{end}}
  {{end}}
{{end}}
`

const goClient = `
{{if .Services}}
  // Client

  {{range .Services}}
  const {{.Name | constPathPrefix}} = "/rpc/{{.Name}}/"
  {{end}}

  {{range .Services}}
    {{ $serviceName := .Name | clientServiceName}}
    type {{$serviceName}} struct {
      client HTTPClient
      urls   [{{.Methods | countMethods}}]string
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

    {{range $i, $method := .Methods}}
      func (c *{{$serviceName}}) {{.Name}}({{.Inputs | methodInputs}}) ({{.Outputs | methodOutputs }}) {
        {{range .Outputs}}
          out := new({{.Type | fieldType}})
        {{end}}
        {{if .Inputs | len}}
        err := doJSONRequest(ctx, c.client, c.urls[{{$i}}], req, out)
        {{else}}
        err := doJSONRequest(ctx, c.client, c.urls[{{$i}}], nil, out)
        {{end}}
        if err != nil {
          return nil, err
        }
        return out, nil
      }
    {{end}}
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
  //   return clientError("aborted because context was done", err)
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
  //   return clientError("aborted because context was done", err)
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
{{if .Services}}
  {{range .Services}}
    type {{.Name}} interface {
      {{range .Methods}}
        {{.Name}}({{.Inputs | methodInputs}}) ({{.Outputs | methodOutputs}})
      {{end}}
    }
  {{end}}
  var Services = map[string][]string{
    {{range .Services}}
      "{{.Name}}": {
        {{range .Methods}}
          "{{.Name}}",
        {{end}}
      },
    {{end}}
  }
{{end}}
`

const goServices = `
{{if .Services}}
  // Server
  {{range .Services}}
    {{$name := .Name}}
    {{$serviceName := .Name | serverServiceName}}

    type {{$serviceName}} struct {
      {{.Name}}
    }

    func {{ .Name | newServerServiceName }}(svc {{.Name}}) WebRPCServer {
      return &{{$serviceName}}{
        {{.Name}}: svc,
      }
    }


    func (s *{{$serviceName}}) WebRPCVersion() string {
      return "v0.0.1"
    }

    func (s *{{$serviceName}}) ServiceVersion() string {
      return "v0.1.0"
    }

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
      {{range .Methods}}
      case "/rpc/{{$name}}/{{.Name}}":
        s.{{.Name | serviceMethodName}}(ctx, w, r)
        return
      {{end}}
      default:
        err := webrpc.Errorf(webrpc.ErrBadRoute, "no handler for path %q", r.URL.Path)
        writeJSONError(ctx, w, r, err)
        return
      }
    }

    {{range .Methods}}
      func (s *{{$serviceName}}) {{.Name | serviceMethodName}}(ctx context.Context, w http.ResponseWriter, r *http.Request) {
        header := r.Header.Get("Content-Type")
        i := strings.Index(header, ";")
        if i == -1 {
          i = len(header)
        }

        switch strings.TrimSpace(strings.ToLower(header[:i])) {
        case "application/json":
          s.{{ .Name | serviceMethodJSONName }}(ctx, w, r)
        default:
          err := webrpc.Errorf(webrpc.ErrBadRoute, "unexpected Content-Type: %q", r.Header.Get("Content-Type"))
          writeJSONError(ctx, w, r, err)
        }
      }

      func (s *{{$serviceName}}) {{.Name | serviceMethodJSONName}}(ctx context.Context, w http.ResponseWriter, r *http.Request) {
        var err error
        ctx = webrpc.WithMethodName(ctx, "{{.Name}}")

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
    {{end}}
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
  "github.com/webrpc/webrpc/lib/webrpc-go"
)

` + goStruct + `
` + goServiceInterface + `
` + goClient + `
` + goServices + `

` + goBaseCode + `
`
