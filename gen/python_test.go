package gen

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

const pythonTestSchema = `
webrpc = v1

name = python-test
version = v1.0.0
basepath = /rpc

struct Item
  - id: uint64

service Items
  - GetItem(id: uint64) => (item: Item)
`

const pythonHeader = `"Webrpc": WEBRPC_HEADER_VALUE`

func TestPythonWebrpcHeader(t *testing.T) {
	tt := []struct {
		name string
		opts map[string]interface{}
		want bool
	}{
		{name: "default", opts: map[string]interface{}{"client": ""}, want: true},
		{name: "disabled", opts: map[string]interface{}{"client": "", "webrpcHeader": "false"}, want: false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out, err := Generate(parseCSharpTestSchema(t, pythonTestSchema), "python", &Config{TemplateOptions: tc.opts})
			if err != nil {
				t.Fatal(err)
			}
			if got := strings.Contains(out.Code, pythonHeader); got != tc.want {
				t.Errorf("expected generated code to contain %q: %v", pythonHeader, tc.want)
			}
		})
	}
}

// The templates exit the process on unsupported input, so the error cases run
// in a subprocess of the test binary.
func TestPythonErrors(t *testing.T) {
	if opts := os.Getenv("WEBRPC_TEST_PYTHON_OPTS"); opts != "" {
		var templateOptions map[string]interface{}
		if err := json.Unmarshal([]byte(opts), &templateOptions); err != nil {
			t.Fatal(err)
		}
		if _, err := Generate(parseCSharpTestSchema(t, os.Getenv("WEBRPC_TEST_PYTHON_SCHEMA")), "python", &Config{TemplateOptions: templateOptions}); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(0)
	}

	tt := []struct {
		name   string
		schema string
		opts   map[string]interface{}
		err    string
	}{
		{
			name:   "streaming",
			schema: strings.Replace(pythonTestSchema, "  - GetItem(id: uint64) => (item: Item)", "  - stream Send(id: uint64)", 1),
			err:    "python generator error: streaming methods are not supported (Items.Send)",
		},
		{
			name:   "file method",
			schema: pythonTestSchema + "  - Upload(data: file) => (size: uint64)\n",
			err:    "python generator error: the file type is not supported (Items.Upload)",
		},
		{
			name:   "file field",
			schema: pythonTestSchema + "\nstruct Blob\n  - data: file\n",
			err:    "python generator error: the file type is not supported (Blob.data)",
		},
		{
			name: "unknown option",
			opts: map[string]interface{}{"pkg": "api"},
			err:  `-pkg="api" is not supported target option`,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ridlSchema := tc.schema
			if ridlSchema == "" {
				ridlSchema = pythonTestSchema
			}
			opts, err := json.Marshal(tc.opts)
			if err != nil {
				t.Fatal(err)
			}

			cmd := exec.Command(os.Args[0], "-test.run=^TestPythonErrors$")
			cmd.Env = append(os.Environ(), "WEBRPC_TEST_PYTHON_OPTS="+string(opts), "WEBRPC_TEST_PYTHON_SCHEMA="+ridlSchema)
			out, err := cmd.CombinedOutput()

			exitErr, ok := err.(*exec.ExitError)
			if !ok || exitErr.ExitCode() != 1 {
				t.Fatalf("expected exit code 1, got %v:\n%s", err, out)
			}
			if !strings.Contains(string(out), tc.err) {
				t.Errorf("expected error %q, got:\n%s", tc.err, out)
			}
		})
	}
}
