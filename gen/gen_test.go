package gen

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/webrpc/webrpc/schema"
)

// Minimal schema; the templates os.Exit on invalid schemas, which would kill
// the test binary.
var testSchema = &schema.WebRPCSchema{WebrpcVersion: "v1", SchemaName: "test"}

func captureStderr(t *testing.T, fn func()) string {
	t.Helper()

	var buf bytes.Buffer
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	origStderr := os.Stderr
	os.Stderr = w
	defer func() { os.Stderr = origStderr }()

	done := make(chan struct{})
	go func() {
		io.Copy(&buf, r)
		close(done)
	}()

	fn()

	w.Close()
	<-done
	r.Close()
	return buf.String()
}

func TestJsoniterDeprecationWarning(t *testing.T) {
	config := &Config{
		TemplateOptions: map[string]interface{}{"json": "jsoniter"},
	}

	stderr := captureStderr(t, func() {
		if _, err := Generate(testSchema, "golang", config); err != nil {
			t.Fatal(err)
		}
	})

	if !strings.Contains(stderr, "warning: -json=jsoniter is deprecated and unsupported") {
		t.Errorf("expected jsoniter deprecation warning on stderr, got: %q", stderr)
	}
}

func TestNoJsoniterWarningByDefault(t *testing.T) {
	stderr := captureStderr(t, func() {
		if _, err := Generate(testSchema, "golang", &Config{}); err != nil {
			t.Fatal(err)
		}
	})

	if strings.Contains(stderr, "warning:") {
		t.Errorf("unexpected warning with default json option, got: %q", stderr)
	}
}