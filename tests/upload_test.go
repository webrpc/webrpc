package tests

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webrpc/webrpc/tests/client"
	"github.com/webrpc/webrpc/tests/server"
)

// outOfOrderTestServer overrides UploadFiles to read the uploaded files out
// of wire order: file parts stream sequentially off the wire, so reading a
// later file auto-drains the unread files before it, whose Bodies must then
// return a clear error.
type outOfOrderTestServer struct {
	*server.TestServer
}

func (s *outOfOrderTestServer) UploadFiles(ctx context.Context, files []*server.File) (uint32, error) {
	if len(files) != 2 {
		return 0, fmt.Errorf("expected 2 files, got %v", len(files))
	}

	// Reading the second file first must still yield its full content.
	second, err := io.ReadAll(files[1].Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read second file: %w", err)
	}
	if string(second) != "second file content" {
		return 0, fmt.Errorf("unexpected second file content: %q", second)
	}

	// The first file was drained while seeking to the second one; reading it
	// now must fail with the wire order error.
	if _, err := io.ReadAll(files[0].Body); err == nil {
		return 0, fmt.Errorf("expected an error reading the drained first file, got none")
	} else if !strings.Contains(err.Error(), "read in wire order") {
		return 0, fmt.Errorf("expected a wire order error reading the drained first file, got: %v", err)
	}

	return uint32(len(files)), nil
}

// TestUploadFilesOutOfOrderRead covers the streaming multipart ordering
// semantic: multiple files are usable without buffering as long as they are
// read in wire order, and an out-of-order read of a drained file fails with
// a clear error rather than returning wrong data.
func TestUploadFilesOutOfOrderRead(t *testing.T) {
	srv := httptest.NewServer(server.NewTestApiServer(&outOfOrderTestServer{&server.TestServer{}}))
	defer srv.Close()

	api := client.NewTestApiClient(srv.URL, &http.Client{})
	count, err := api.UploadFiles(context.Background(), []*client.File{
		{
			Name:        "first.txt",
			ContentType: "text/plain",
			Size:        int64(len("first file content")),
			Body:        io.NopCloser(bytes.NewReader([]byte("first file content"))),
		},
		{
			Name:        "second.txt",
			ContentType: "text/plain",
			Size:        int64(len("second file content")),
			Body:        io.NopCloser(bytes.NewReader([]byte("second file content"))),
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, uint32(2), count)
}

// TestUploadFilesWithoutJsonPart covers the wire format requirement for
// []file methods: the leading "json" part must declare the []file argument
// as an array with one null per file part, so the server knows how many
// parts to expect and can keep them streaming. A request without that shape
// is rejected with a clear HTTP 400 error.
func TestUploadFilesWithoutJsonPart(t *testing.T) {
	srv := httptest.NewServer(server.NewTestApiServer(&server.TestServer{}))
	defer srv.Close()

	var body bytes.Buffer
	form := multipart.NewWriter(&body)
	for i, content := range [][]byte{[]byte("first file"), []byte("second file\x00\xff")} {
		part, err := form.CreateFormFile("files", fmt.Sprintf("file%v.bin", i))
		assert.NoError(t, err)
		_, err = part.Write(content)
		assert.NoError(t, err)
	}
	assert.NoError(t, form.Close())

	resp, err := http.Post(srv.URL+"/rpc/TestApi/UploadFiles", form.FormDataContentType(), &body)
	assert.NoError(t, err)
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode, string(respBody))
	assert.Contains(t, string(respBody), `must declare \"files\" as a JSON array with one null per file part`)
}
