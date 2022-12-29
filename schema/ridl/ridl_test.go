package ridl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
	"github.com/webrpc/webrpc/schema"
)

func parseString(src string) (*schema.WebRPCSchema, error) {
	fsys := fstest.MapFS{
		"main.ridl": {
			Data: []byte(src),
		},
	}
	return NewParser(fsys, "main.ridl").Parse()
}

func compactJSON(src []byte) string {
	buf := bytes.NewBuffer(nil)

	err := json.Compact(buf, src)
	if err != nil {
		panic(fmt.Sprintf("json.Compact: %v", err))
	}

	return buf.String()
}

func TestRIDLHeader(t *testing.T) {
	{
		buf := `
    name = myapi
  `
		_, err := parseString(buf)
		assert.Error(t, err, `"version" is required`)
	}

	{
		buf := `
    webrpc = v1

    name = myapi1
    name = myapi2
  `
		_, err := parseString(buf)
		assert.Error(t, err, `should not be able to declare "ridl" twice`)
	}

	{
		buf := `
    webrpc = v1 #comment
    # comment
  version = v0.1.1

  name= h_ello-webrpc
  `
		s, err := parseString(buf)
		assert.NoError(t, err)

		assert.Equal(t, "v1", s.WebrpcVersion)
		assert.Equal(t, "h_ello-webrpc", s.SchemaName)
		assert.Equal(t, "v0.1.1", s.SchemaVersion)
	}
}

func TestRIDLImports(t *testing.T) {
	fsys := fstest.MapFS{
		"schema/import-service.ridl": {Data: []byte(`
			webrpc = v1
			version = v0.1.1
			name = ImportService
	
			import
			- types.ridl
			`)},
		"schema/types.ridl": {Data: []byte(`
			webrpc = v1
			version = v1.0.0
			name = types

			import
			  - foo.ridl        # import from current directory
			  - subdir/bar.ridl # import from subdirectory
			  - ../common.ridl  # import from parent directory

			struct ExtraType
			  - name: string
		`)},
		"schema/foo.ridl": {Data: []byte(`
			webrpc = v1
			version = v0.8.0
			name = foo

			struct Foo
			- name: string	
		`)},
		"schema/subdir/bar.ridl": {Data: []byte(`
			webrpc = v1
			version = v0.8.0
			name = bar

			struct Bar
			- name: string	

			struct Baz
			- name: string	
			`)},
		"common.ridl": {Data: []byte(`
		webrpc = v1
		version = v1.0.0
		name = common

		struct Common
		- name: string	
		`)},
	}

	s, err := NewParser(fsys, "schema/import-service.ridl").Parse()
	assert.NoError(t, err)

	assert.Equal(t, "v1", s.WebrpcVersion)
	assert.Equal(t, "ImportService", s.SchemaName)
	assert.Equal(t, "v0.1.1", s.SchemaVersion)

	if assert.Equal(t, 5, len(s.Types)) {
		assert.Equal(t, "Foo", string(s.Types[0].Name))
		assert.Equal(t, "Bar", string(s.Types[1].Name))
		assert.Equal(t, "Baz", string(s.Types[2].Name))
		assert.Equal(t, "Common", string(s.Types[3].Name))
		assert.Equal(t, "ExtraType", string(s.Types[4].Name))
	}
}

func TestRIDLEnum(t *testing.T) {
	{
		input := `
    webrpc = v1
    version = v0.1.1
  name = hello-webrpc

          # this is a comment
            # yep
          enum Kind:uint32
            - USER = 33             # comment
            - ADMIN = 44            # comment..

        # or.. just..
        enum   KindTwo:    uint32
          - USER                 # aka, = 0
          - ADMIN         # aka, = 1
          - OTHER
  `
		s, err := parseString(input)
		assert.NoError(t, err)

		assert.Equal(t, "v1", s.WebrpcVersion)
		assert.Equal(t, "hello-webrpc", s.SchemaName)
		assert.Equal(t, "v0.1.1", s.SchemaVersion)

		assert.Equal(t, "Kind", string(s.Types[0].Name))
		assert.Equal(t, "enum", string(s.Types[0].Kind))
		assert.Equal(t, "uint32", string(s.Types[0].Type.String()))

		assert.Equal(t, "USER", string(s.Types[0].Fields[0].Name))
		assert.Equal(t, "ADMIN", string(s.Types[0].Fields[1].Name))

		assert.Equal(t, "33", string(s.Types[0].Fields[0].Value))
		assert.Equal(t, "44", string(s.Types[0].Fields[1].Value))

		assert.Equal(t, (*schema.VarType)(nil), s.Types[0].Fields[0].Type)
		assert.Equal(t, (*schema.VarType)(nil), s.Types[0].Fields[1].Type)

		assert.Equal(t, "KindTwo", string(s.Types[1].Name))
		assert.Equal(t, "enum", string(s.Types[1].Kind))
		assert.Equal(t, "uint32", string(s.Types[1].Type.String()))

		assert.Equal(t, (*schema.VarType)(nil), s.Types[1].Fields[0].Type)
		assert.Equal(t, (*schema.VarType)(nil), s.Types[1].Fields[1].Type)
		assert.Equal(t, (*schema.VarType)(nil), s.Types[1].Fields[2].Type)

		assert.Equal(t, "0", string(s.Types[1].Fields[0].Value))
		assert.Equal(t, "1", string(s.Types[1].Fields[1].Value))
		assert.Equal(t, "2", string(s.Types[1].Fields[2].Value))
	}
}

func TestRIDLTypes(t *testing.T) {
	{
		input := `
    webrpc = v1
    version = v0.1.1
  name = hello-webrpc

  struct Empty
  `
		s, err := parseString(input)
		assert.NoError(t, err)

		assert.Equal(t, "struct", string(s.Types[0].Kind))
		assert.Equal(t, "Empty", string(s.Types[0].Name))
	}

	{
		input := `
    webrpc = v1
    version = v0.1.1
  name = hello-webrpc

  struct Empty # with a, comment
  `
		s, err := parseString(input)
		assert.NoError(t, err)

		assert.Equal(t, "struct", string(s.Types[0].Kind))
		assert.Equal(t, "Empty", string(s.Types[0].Name))

	}

	{
		input := `
    webrpc = v1
    version = v0.1.1
  name = hello-webrpc

  struct Simple # with a, comment
    - ID: uint32
    - Value?: uint32
  `
		s, err := parseString(input)
		assert.NoError(t, err)

		assert.Equal(t, "struct", string(s.Types[0].Kind))
		assert.Equal(t, "Simple", string(s.Types[0].Name))

		assert.Equal(t, "ID", string(s.Types[0].Fields[0].Name))
		assert.Equal(t, "uint32", string(s.Types[0].Fields[0].Type.String()))
		assert.Equal(t, false, s.Types[0].Fields[0].Optional)

		assert.Equal(t, "Value", string(s.Types[0].Fields[1].Name))
		assert.Equal(t, "uint32", string(s.Types[0].Fields[1].Type.String()))
		assert.Equal(t, true, s.Types[0].Fields[1].Optional)
	}

	{
		input := `
    webrpc = v1
    version = v0.1.1
  name = hello-webrpc

  struct Simple # with a-comment an,d meta fields
    - ID: uint32
  - Field2: uint64 # one two #t
      + json = field_2 # a comment
      #
        + go.tag.db = field_3
        #
  - Field3: bool
    + go.tag.db = - # omits the field from db


  struct Simple2 # with a-comment an,d meta fields
  `
		s, err := parseString(input)
		assert.NoError(t, err)

		assert.Equal(t, "struct", string(s.Types[0].Kind))
		assert.Equal(t, "Simple", string(s.Types[0].Name))

		assert.Equal(t, "struct", string(s.Types[1].Kind))
		assert.Equal(t, "Simple2", string(s.Types[1].Name))

		assert.Equal(t, "ID", string(s.Types[0].Fields[0].Name))
		assert.Equal(t, "Field2", string(s.Types[0].Fields[1].Name))
		assert.Equal(t, "Field3", string(s.Types[0].Fields[2].Name))

		assert.Equal(t, "field_2", s.Types[0].Fields[1].Meta[0]["json"])
		assert.Equal(t, "field_3", s.Types[0].Fields[1].Meta[1]["go.tag.db"])

		assert.Equal(t, "-", s.Types[0].Fields[2].Meta[0]["go.tag.db"])
	}

	{
		input := `
    webrpc = v1
    version = v0.1.1
  name = hello-webrpc

  struct Simple # with a-comment an,d meta fields
    - ID: uint32
  - Field2: map<string, string> # one two #t
      + json = field_2 # a comment
        + go.tag.db = field_2
  - Field3: []bool # one two #t
      + json = field_2 # a comment
        + go.tag.db = field_2
  - Field4: [][][]bool # one two #t
      + json = field_2 # a comment
        + go.tag.db = field_2

  struct Simple2 # with a-comment an,d meta fields
  `
		s, err := parseString(input)
		assert.NoError(t, err)

		assert.Equal(t, "map<string,string>", string(s.Types[0].Fields[1].Type.String()))
		assert.Equal(t, "[]bool", string(s.Types[0].Fields[2].Type.String()))
		assert.Equal(t, "[][][]bool", string(s.Types[0].Fields[3].Type.String()))
	}

	{
		input := `
    webrpc = v1
    version = v0.1.1
  name = hello-webrpc

  struct Simple # with a-comment an,d meta fields
    - ID: uint32
  - Field2: map<string, string> # one two #t
      + json = field_2 # a comment
        + go.tag.db = field_2 # a comment
    + go.tag.db.1 = default**:**now**()**,use_zero#000 # # # a comment

        + go.tag.db.2 = default**:**now**()**,use_zero,"// # a comment
        + go.tag.db.3 = "default**:**now**()**,use_zero,// # not a comment" # a comment
        + go.tag.db.4 = default**:**now**()**,use_zero`
		s, err := parseString(input)
		assert.NoError(t, err)

		assert.Equal(t, "map<string,string>", string(s.Types[0].Fields[1].Type.String()))
		assert.Equal(t, "field_2", s.Types[0].Fields[1].Meta[1]["go.tag.db"])
		assert.Equal(t, "default**:**now**()**,use_zero#000", s.Types[0].Fields[1].Meta[2]["go.tag.db.1"])
		assert.Equal(t, `default**:**now**()**,use_zero,"//`, s.Types[0].Fields[1].Meta[3]["go.tag.db.2"])
		assert.Equal(t, "default**:**now**()**,use_zero,// # not a comment", s.Types[0].Fields[1].Meta[4]["go.tag.db.3"])
		assert.Equal(t, "default**:**now**()**,use_zero", s.Types[0].Fields[1].Meta[5]["go.tag.db.4"])
	}
}

func TestRIDLService(t *testing.T) {
	{
		input := `
    webrpc = v1
    version = v0.1.1
  name = hello-webrpc

  service Pinger
    - Ping()
  `
		s, err := parseString(input)
		assert.NoError(t, err)

		assert.Equal(t, "Pinger", string(s.Services[0].Name))
		assert.Equal(t, "Ping", string(s.Services[0].Methods[0].Name))
	}

	{
		input := `
    webrpc = v1
    version = v0.1.1
          name = hello-webrpc

          service Simple
          - Ping()
          -  Status() => (status: bool)
          -  StatusStream(q: string) => stream (status: bool)`

		s, err := parseString(input)
		assert.NoError(t, err)

		assert.Equal(t, "Ping", string(s.Services[0].Methods[0].Name))
		assert.Equal(t, "Status", string(s.Services[0].Methods[1].Name))
		assert.Equal(t, "StatusStream", string(s.Services[0].Methods[2].Name))

		assert.Equal(t, 0, len(s.Services[0].Methods[1].Inputs))
		assert.Equal(t, "status", string(s.Services[0].Methods[1].Outputs[0].Name))
		assert.Equal(t, "bool", s.Services[0].Methods[1].Outputs[0].Type.String())

		assert.Equal(t, "q", string(s.Services[0].Methods[2].Inputs[0].Name))
		assert.Equal(t, "string", s.Services[0].Methods[2].Inputs[0].Type.String())

		assert.Equal(t, "status", string(s.Services[0].Methods[2].Outputs[0].Name))
		assert.Equal(t, "bool", s.Services[0].Methods[2].Outputs[0].Type.String())
	}

	{
		input := `
    webrpc = v1
    version = v0.1.1
    name = hello-webrpc

    service Simple
      -  stream Ping(code?: uint32) => (code: bool)
      -  PingStream(text: string) => stream (code?: bool)
    `
		s, err := parseString(input)
		assert.NoError(t, err)

		assert.Equal(t, "Ping", string(s.Services[0].Methods[0].Name))
		assert.True(t, s.Services[0].Methods[0].StreamInput)
		assert.False(t, s.Services[0].Methods[0].StreamOutput)
		assert.True(t, s.Services[0].Methods[0].Inputs[0].Optional)

		assert.False(t, s.Services[0].Methods[1].StreamInput)
		assert.True(t, s.Services[0].Methods[1].StreamOutput)
		assert.True(t, s.Services[0].Methods[1].Outputs[0].Optional)

	}

	{
		input := `
    webrpc = v1
    version = v0.1.1
    name = hello-webrpc

    service Simple
    -  Ping(header: map<string,[][]string>) => (code: bool)
      -  stream VerifyUsers(seq: int32, header?: map<string,[]string>, ids: []uint64) => (code?: bool, ids: []bool)
    - MoreTest(n: uint64, stuff: []map<uint64,   map<int32,             string>>, etc: string) => (code: bool)`
		s, err := parseString(input)
		assert.NoError(t, err)

		assert.Equal(t, "map<string,[][]string>", s.Services[0].Methods[0].Inputs[0].Type.String())
		assert.Equal(t, "[]map<uint64,map<int32,string>>", s.Services[0].Methods[2].Inputs[1].Type.String())
	}
}

func TestRIDLParse(t *testing.T) {
	fp, err := os.Open("_example/example0.ridl")
	assert.NoError(t, err)

	buf, err := ioutil.ReadAll(fp)
	assert.NoError(t, err)

	s, err := parseString(string(buf))
	assert.NoError(t, err)

	jout, err := s.ToJSON(true)
	assert.NoError(t, err)

	assert.NotZero(t, jout)
}

func TestRIDLImportsExampleDir(t *testing.T) {
	exampleDirFS := os.DirFS("./_example")

	r := NewParser(exampleDirFS, "example1.ridl")
	s, err := r.Parse()
	assert.NoError(t, err)

	jout, err := s.ToJSON(true)
	assert.NoError(t, err)

	assert.NotZero(t, jout)

	golden, err := ioutil.ReadFile("./_example/example1-golden.json")
	assert.NoError(t, err)

	a := compactJSON(golden)
	b := compactJSON([]byte(jout))

	//ioutil.WriteFile("example1-golden.json", []byte(jout), 0644)

	// fmt.Println("==> GOLDEN:", a)
	// fmt.Println("==> PARSED:", b)

	assert.JSONEq(t, a, b)
}
