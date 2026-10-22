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

// The contextual keyword yield is a valid namespace segment.
func TestCSharpNamespaceYield(t *testing.T) {
	out, err := Generate(parseCSharpTestSchema(t, csharpTestSchema), "csharp", &Config{TemplateOptions: map[string]interface{}{"client": "", "namespace": "Finance.yield"}})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out.Code, "namespace Finance.yield\n") {
		t.Errorf("expected namespace Finance.yield")
	}
}

// Generated code references its runtime helpers by their full name, so schema
// names (enum values, arguments) can't hide them.
func TestCSharpQualifiedHelpers(t *testing.T) {
	out, err := Generate(parseCSharpTestSchema(t, `
webrpc = v1

name = helpers
version = v1.0.0
basepath = /rpc

enum Kind: string
  - WebrpcJson

service Svc
  - Ping(WebrpcJson: string, WebrpcHttp: string) => (kind: Kind)
`), "csharp", &Config{TemplateOptions: map[string]interface{}{"client": ""}})
	if err != nil {
		t.Fatal(err)
	}

	for _, want := range []string{
		"global::Helpers.WebrpcJson.ReadEnum(",
		"global::Helpers.WebrpcJson.Serialize(__args)",
		"global::Helpers.WebrpcHttp.NewRequest(",
		"global::Helpers.WebrpcHttp.CallAsync<",
	} {
		if !strings.Contains(out.Code, want) {
			t.Errorf("expected generated code to contain %q", want)
		}
	}
}

// Arguments named like the generated parameters, locals and fields are renamed.
func TestCSharpArgumentNames(t *testing.T) {
	out, err := Generate(parseCSharpTestSchema(t, `
webrpc = v1

name = args
version = v1.0.0
basepath = /rpc

service Svc
  - Do(__json: string, _baseUrl: string, cancellationToken: string, cancellationToken_: string) => (ok: bool)
`), "csharp", &Config{TemplateOptions: map[string]interface{}{"client": ""}})
	if err != nil {
		t.Fatal(err)
	}

	want := "DoAsync(string __json_, string _baseUrl_, string cancellationToken__, string cancellationToken_, global::System.Threading.CancellationToken cancellationToken = default)"
	if !strings.Contains(out.Code, want) {
		t.Errorf("expected generated code to contain %q", want)
	}
	if !strings.Contains(out.Code, `__args["cancellationToken"] = cancellationToken__;`) {
		t.Errorf("expected the renamed argument to keep its JSON name")
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

# Field order doesn't matter: toString_ keeps its name.
struct Reordered
  - toString_: string
  - toString: string

# A field left out with + json = - doesn't take a name.
struct Hidden
  - equals: string
    + json = -
  - equals_: string

enum Flag: string
  - Value
  - Value_

# Names the compiler reserves for the generated accessors and operators.
enum Accessor: string
  - get_Value
  - op_Implicit

# On an integer enum value, + json = - is a wire value, not an omission.
enum Status: uint32
  - Unknown
    + json = -

# The last + json wins.
struct Renamed
  - id: string
    + json = -
    + json = ident

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
		`public static readonly Flag Value_ = new Flag("Value_");`,
		`public static readonly Flag Value__ = new Flag("Value");`,
		"public string Equals_ { get; set; }",
		`public static readonly Accessor get_Value_ = new Accessor("get_Value");`,
		`public static readonly Accessor op_Implicit_ = new Accessor("op_Implicit");`,
		"JsonPropertyName(\"ident\")]\n        public string Id { get; set; }",
		`public static readonly Status Unknown = new Status("-");`,
	} {
		if !strings.Contains(out.Code, want) {
			t.Errorf("expected generated code to contain %q", want)
		}
	}
	// The JSON field of each renamed property doesn't depend on field order.
	for _, want := range []string{
		"JsonPropertyName(\"toString_\")]\n        public string ToString_ { get; set; }",
		"JsonPropertyName(\"toString\")]\n        public string ToString__ { get; set; }",
	} {
		if n := strings.Count(out.Code, want); n != 2 {
			t.Errorf("expected %q in Item and Reordered, got %d", want, n)
		}
	}

	// The GetResponse class renames its outputs the same way.
	for _, prop := range []string{"ToString_", "ToString__"} {
		if n := strings.Count(out.Code, "public string "+prop+" { get; set; }"); n != 3 {
			t.Errorf("expected %s in Item, Reordered and GetResponse, got %d", prop, n)
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
		{name: "namespace __arglist", opts: map[string]interface{}{"namespace": "My.__arglist"}, err: `"__arglist" is a C# keyword`},
		{name: "namespace await", opts: map[string]interface{}{"namespace": "My.await"}, err: `-namespace="My.await" is not a valid C# namespace: "await" is a C# keyword`},
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
			err:    `"WebrpcJson", a schema type, clashes with a generated runtime helper; rename it in the schema`,
		},
		{
			name:   "duplicate arguments",
			schema: strings.Replace(csharpTestSchema, "  - GetItem(id: uint64) => (item: Item)", "  - GetItem(id: uint64, id: string) => (item: Item)", 1),
			opts:   map[string]interface{}{"client": ""},
			err:    `method Items.GetItem has two arguments named "id"`,
		},
		{
			name:   "duplicate outputs",
			schema: strings.Replace(csharpTestSchema, "  - GetItem(id: uint64) => (item: Item)", "  - GetItem(id: uint64) => (item: Item, item: Item)", 1),
			opts:   map[string]interface{}{"client": ""},
			err:    `method Items.GetItem has two outputs named "item"`,
		},
		{
			name: "ignored route field",
			schema: strings.Replace(csharpTestSchema, "  - GetItem(id: uint64) => (item: Item)", "  - Find(FindRequest) => (FindResponse)\n      + GET /items/{id}", 1) +
				"\nstruct FindRequest\n  - ident: string\n    + json = -\n    + go.tag.json = id\n\nstruct FindResponse\n  - ok: bool\n",
			opts: map[string]interface{}{"client": ""},
			err:  "field FindRequest.ident is left out with + json = -, but a REST route or file upload of a method uses it",
		},
		{
			name:   "no unique name",
			schema: csharpTestSchema + "\nenum Kind: string\n  - Value\n" + csharpUnderscoreChain("  - Value", 33),
			err:    `can't find a unique C# name for "Value" in type Kind`,
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

// csharpUnderscoreChain renders the RIDL lines prefix_, prefix__, ... up to n underscores.
func csharpUnderscoreChain(prefix string, n int) string {
	var b strings.Builder
	for i := 1; i <= n; i++ {
		b.WriteString(prefix + strings.Repeat("_", i) + "\n")
	}
	return b.String()
}
