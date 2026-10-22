package gen

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/webrpc/webrpc/schema"
	"github.com/webrpc/webrpc/schema/ridl"
)

const csharpTestSchema = `
webrpc = v1

name = csharp-test
version = v1.0.0
basepath = /rpc

struct Item
  - id: uint64

service Items
  - GetItem(id: uint64) => (item: Item)
`

func parseCSharpTestSchema(t testing.TB, ridlSchema string) *schema.WebRPCSchema {
	t.Helper()

	fsys := fstest.MapFS{"schema.ridl": {Data: []byte(ridlSchema)}}
	s, err := ridl.NewParser(fsys, "/", "schema.ridl").Parse()
	if err != nil {
		t.Fatal(err)
	}
	return s
}

func TestCSharpDefaults(t *testing.T) {
	out, err := Generate(parseCSharpTestSchema(t, csharpTestSchema), "csharp", &Config{TemplateOptions: map[string]interface{}{"client": ""}})
	if err != nil {
		t.Fatal(err)
	}

	for _, want := range []string{
		"namespace CsharpTest\n",
		"public partial class Item",
		"[global::System.Text.Json.Serialization.JsonPropertyName(\"id\")]",
		"public interface IItemsClient",
		"public partial class ItemsClient : IItemsClient",
		"_baseUrl + \"/rpc/Items/GetItem\"",
		"request.Headers.TryAddWithoutValidation(Webrpc.Header, Webrpc.HeaderValue);",
		"public const string SchemaHash = ",
	} {
		if !strings.Contains(out.Code, want) {
			t.Errorf("expected generated code to contain %q", want)
		}
	}
}

func TestCSharpOptions(t *testing.T) {
	out, err := Generate(parseCSharpTestSchema(t, csharpTestSchema), "csharp", &Config{TemplateOptions: map[string]interface{}{
		"client":       "",
		"namespace":    "My.Api",
		"json":         "newtonsoft",
		"visibility":   "internal",
		"schemaHash":   "false",
		"webrpcHeader": "false",
	}})
	if err != nil {
		t.Fatal(err)
	}

	for _, want := range []string{
		"namespace My.Api\n",
		"internal partial class Item",
		"internal partial class ItemsClient : IItemsClient",
		"[global::Newtonsoft.Json.JsonProperty(\"id\")]",
	} {
		if !strings.Contains(out.Code, want) {
			t.Errorf("expected generated code to contain %q", want)
		}
	}
	for _, unwanted := range []string{"System.Text.Json", "SchemaHash", "TryAddWithoutValidation(Webrpc.Header"} {
		if strings.Contains(out.Code, unwanted) {
			t.Errorf("expected generated code not to contain %q", unwanted)
		}
	}
}

func TestCSharpTypesOnly(t *testing.T) {
	out, err := Generate(parseCSharpTestSchema(t, csharpTestSchema), "csharp", &Config{})
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(out.Code, "public partial class Item") {
		t.Errorf("expected types without -client")
	}
	if strings.Contains(out.Code, "ItemsClient") || strings.Contains(out.Code, "WebrpcHttp") {
		t.Errorf("expected no client without -client")
	}
}

// Members renamed to avoid a clash get "_" suffixes until they are unique.
func TestCSharpMemberNames(t *testing.T) {
	schema := parseCSharpTestSchema(t, `
webrpc = v1

name = names
version = v1.0.0
basepath = /rpc

struct Item
  - toString: string
  - toString_: string
  - item: string
  - item_: string

enum Flag: string
  - Value
  - Value_

service Items
  - Get() => (toString: string, toString_: string)
`)
	out, err := Generate(schema, "csharp", &Config{TemplateOptions: map[string]interface{}{"client": ""}})
	if err != nil {
		t.Fatal(err)
	}

	for _, want := range []string{
		"public string ToString_ { get; set; }",
		"public string ToString__ { get; set; }",
		"public string Item_ { get; set; }",
		"public string Item__ { get; set; }",
		`public static readonly Flag Value_ = new Flag("Value");`,
		`public static readonly Flag Value__ = new Flag("Value_");`,
	} {
		if !strings.Contains(out.Code, want) {
			t.Errorf("expected generated code to contain %q", want)
		}
	}
	// The GetResponse class renames its outputs the same way.
	for _, prop := range []string{"ToString_", "ToString__"} {
		if n := strings.Count(out.Code, "public string "+prop+" { get; set; }"); n != 2 {
			t.Errorf("expected %s in Item and GetResponse, got %d", prop, n)
		}
	}
}

// The templates exit the process on invalid options, so the error cases run
// in a subprocess of the test binary.
func TestCSharpErrors(t *testing.T) {
	if opts := os.Getenv("WEBRPC_TEST_CSHARP_OPTS"); opts != "" {
		var templateOptions map[string]interface{}
		if err := json.Unmarshal([]byte(opts), &templateOptions); err != nil {
			t.Fatal(err)
		}
		if _, err := Generate(parseCSharpTestSchema(t, os.Getenv("WEBRPC_TEST_CSHARP_SCHEMA")), "csharp", &Config{TemplateOptions: templateOptions}); err != nil {
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
		{name: "server", opts: map[string]interface{}{"server": ""}, err: "the C# target generates clients only, -server is not supported"},
		{name: "json", opts: map[string]interface{}{"json": "jsoniter"}, err: `-json="jsoniter" is not supported`},
		{name: "visibility", opts: map[string]interface{}{"visibility": "private"}, err: `-visibility="private" is not supported`},
		{name: "namespace", opts: map[string]interface{}{"namespace": "My.1Api"}, err: `-namespace="My.1Api" is not a valid C# namespace`},
		{name: "unknown option", opts: map[string]interface{}{"pkg": "api"}, err: `-pkg="api" is not supported target option`},
		{name: "namespace keyword", opts: map[string]interface{}{"namespace": "My.class"}, err: `-namespace="My.class" is not a valid C# namespace: "class" is a C# keyword`},
		{
			name:   "client names",
			schema: csharpTestSchema + "\nservice IItems\n  - Ping()\n",
			opts:   map[string]interface{}{"client": ""},
			err:    `"IItemsClient", the client class of service IItems, clashes with the client interface of service Items`,
		},
		{
			name:   "type and client",
			schema: csharpTestSchema + "\nstruct ItemsClient\n  - id: uint64\n",
			opts:   map[string]interface{}{"client": ""},
			err:    `"ItemsClient", a schema type, clashes with the client class of service Items`,
		},
		{
			name:   "type and error",
			schema: csharpTestSchema + "\nstruct NotFoundException\n  - id: uint64\n\nerror 1 NotFound \"not found\" HTTP 404\n",
			opts:   map[string]interface{}{"client": ""},
			err:    `"NotFoundException", a schema type, clashes with the exception of error NotFound`,
		},
		{
			name:   "type and runtime helper",
			schema: csharpTestSchema + "\nstruct WebrpcJson\n  - id: uint64\n",
			err:    `"WebrpcJson", a schema type, clashes with a generated runtime helper`,
		},
		{
			name:   "stream input",
			schema: strings.Replace(csharpTestSchema, "  - GetItem(id: uint64) => (item: Item)", "  - stream Send(id: uint64)", 1),
			opts:   map[string]interface{}{"client": ""},
			err:    "streaming input is not supported (Items.Send)",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ridlSchema := tc.schema
			if ridlSchema == "" {
				ridlSchema = csharpTestSchema
			}
			opts, err := json.Marshal(tc.opts)
			if err != nil {
				t.Fatal(err)
			}

			cmd := exec.Command(os.Args[0], "-test.run=^TestCSharpErrors$")
			cmd.Env = append(os.Environ(), "WEBRPC_TEST_CSHARP_OPTS="+string(opts), "WEBRPC_TEST_CSHARP_SCHEMA="+ridlSchema)
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
