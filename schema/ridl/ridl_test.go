package ridl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		_, err := Parse(buf)
		assert.Error(t, err, `"version" is required`)
	}

	{
		buf := `
    webrpc = v1

    name = myapi1
    name = myapi2
  `
		_, err := Parse(buf)
		assert.Error(t, err, `should not be able to declare "ridl" twice`)
	}

	{
		buf := `
    webrpc = v1 #comment
    # comment
  version = v0.1.1

  name= h_ello-webrpc
  `
		s, err := Parse(buf)
		assert.NoError(t, err)

		assert.Equal(t, "v1", s.WebRPCVersion)
		assert.Equal(t, "h_ello-webrpc", s.Name)
		assert.Equal(t, "v0.1.1", s.SchemaVersion)
	}
}

func TestRIDLImport(t *testing.T) {
	{
		input := `
    webrpc = v1
      version = v0.1.1
  name = hello-webrpc

    import
    - foo # ko ment
    # ko ment

      - bar
      # comment
    `
		s, err := Parse(input)
		assert.NoError(t, err)

		assert.Equal(t, "v1", s.WebRPCVersion)
		assert.Equal(t, "hello-webrpc", s.Name)
		assert.Equal(t, "v0.1.1", s.SchemaVersion)

		assert.Equal(t, "foo", s.Imports[0])
		assert.Equal(t, "bar", s.Imports[1])
	}

	{
		input := `
    webrpc = v1
    version = v0.1.1 # version number
  name     = hello-webrpc

  import # import line
  - foo1 # foo-comment with spaces
    - bar2 # # # bar-comment
  `
		s, err := Parse(input)
		assert.NoError(t, err)

		assert.Equal(t, "v1", s.WebRPCVersion)
		assert.Equal(t, "hello-webrpc", s.Name)
		assert.Equal(t, "v0.1.1", s.SchemaVersion)

		assert.Equal(t, "foo1", s.Imports[0])
		assert.Equal(t, "bar2", s.Imports[1])
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
		s, err := Parse(input)
		assert.NoError(t, err)

		assert.Equal(t, "v1", s.WebRPCVersion)
		assert.Equal(t, "hello-webrpc", s.Name)
		assert.Equal(t, "v0.1.1", s.SchemaVersion)

		assert.Equal(t, "Kind", string(s.Messages[0].Name))
		assert.Equal(t, "enum", string(s.Messages[0].Type))

		assert.Equal(t, "USER", string(s.Messages[0].Fields[0].Name))
		assert.Equal(t, "ADMIN", string(s.Messages[0].Fields[1].Name))

		assert.Equal(t, "33", string(s.Messages[0].Fields[0].Value))
		assert.Equal(t, "44", string(s.Messages[0].Fields[1].Value))

		assert.Equal(t, "uint32", string(s.Messages[0].Fields[0].Type.String()))
		assert.Equal(t, "uint32", string(s.Messages[0].Fields[1].Type.String()))

		assert.Equal(t, "KindTwo", string(s.Messages[1].Name))
		assert.Equal(t, "enum", string(s.Messages[1].Type))

		assert.Equal(t, "uint32", string(s.Messages[1].Fields[0].Type.String()))
		assert.Equal(t, "uint32", string(s.Messages[1].Fields[1].Type.String()))
		assert.Equal(t, "uint32", string(s.Messages[1].Fields[2].Type.String()))

		assert.Equal(t, "0", string(s.Messages[1].Fields[0].Value))
		assert.Equal(t, "1", string(s.Messages[1].Fields[1].Value))
		assert.Equal(t, "2", string(s.Messages[1].Fields[2].Value))
	}
}

func TestRIDLMessages(t *testing.T) {
	{
		input := `
    webrpc = v1
    version = v0.1.1
  name = hello-webrpc

  message Empty
  `
		s, err := Parse(input)
		assert.NoError(t, err)

		assert.Equal(t, "Empty", string(s.Messages[0].Name))
		assert.Equal(t, "struct", string(s.Messages[0].Type))
	}

	{
		input := `
    webrpc = v1
    version = v0.1.1
  name = hello-webrpc

  message Empty # with a, comment
  `
		s, err := Parse(input)
		assert.NoError(t, err)

		assert.Equal(t, "Empty", string(s.Messages[0].Name))
		assert.Equal(t, "struct", string(s.Messages[0].Type))

	}

	{
		input := `
    webrpc = v1
    version = v0.1.1
  name = hello-webrpc

  message Simple # with a, comment
    - ID: uint32
    - Value?: uint32
  `
		s, err := Parse(input)
		assert.NoError(t, err)

		assert.Equal(t, "Simple", string(s.Messages[0].Name))
		assert.Equal(t, "struct", string(s.Messages[0].Type))

		assert.Equal(t, "ID", string(s.Messages[0].Fields[0].Name))
		assert.Equal(t, "uint32", string(s.Messages[0].Fields[0].Type.String()))
		assert.Equal(t, false, s.Messages[0].Fields[0].Optional)

		assert.Equal(t, "Value", string(s.Messages[0].Fields[1].Name))
		assert.Equal(t, "uint32", string(s.Messages[0].Fields[1].Type.String()))
		assert.Equal(t, true, s.Messages[0].Fields[1].Optional)
	}

	{
		input := `
    webrpc = v1
    version = v0.1.1
  name = hello-webrpc

  message Simple # with a-comment an,d meta fields
    - ID: uint32
  - Field2: uint64 # one two #t
      + json = field_2 # a comment
      #
        + go.tag.db = field_3
        #
  - Field3: bool
    + go.tag.db = - # omits the field from db


  message Simple2 # with a-comment an,d meta fields
  `
		s, err := Parse(input)
		assert.NoError(t, err)

		assert.Equal(t, "Simple", string(s.Messages[0].Name))
		assert.Equal(t, "struct", string(s.Messages[0].Type))

		assert.Equal(t, "Simple2", string(s.Messages[1].Name))
		assert.Equal(t, "struct", string(s.Messages[1].Type))

		assert.Equal(t, "ID", string(s.Messages[0].Fields[0].Name))
		assert.Equal(t, "Field2", string(s.Messages[0].Fields[1].Name))
		assert.Equal(t, "Field3", string(s.Messages[0].Fields[2].Name))

		assert.Equal(t, "field_2", s.Messages[0].Fields[1].Meta[0]["json"])
		assert.Equal(t, "field_3", s.Messages[0].Fields[1].Meta[1]["go.tag.db"])

		assert.Equal(t, "-", s.Messages[0].Fields[2].Meta[0]["go.tag.db"])
	}

	{
		input := `
    webrpc = v1
    version = v0.1.1
  name = hello-webrpc

  message Simple # with a-comment an,d meta fields
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

  message Simple2 # with a-comment an,d meta fields
  `
		s, err := Parse(input)
		assert.NoError(t, err)

		assert.Equal(t, "map<string,string>", string(s.Messages[0].Fields[1].Type.String()))
		assert.Equal(t, "[]bool", string(s.Messages[0].Fields[2].Type.String()))
		assert.Equal(t, "[][][]bool", string(s.Messages[0].Fields[3].Type.String()))
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
		s, err := Parse(input)
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

		s, err := Parse(input)
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
		s, err := Parse(input)
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
		s, err := Parse(input)
		assert.NoError(t, err)

		assert.Equal(t, "map<string,[][]string>", s.Services[0].Methods[0].Inputs[0].Type.String())
		assert.Equal(t, "[]map<uint64,map<int32,string>>", s.Services[0].Methods[2].Inputs[1].Type.String())
	}
}

func TestRIDLParse(t *testing.T) {
	fp, err := os.Open("example1.ridl")
	assert.NoError(t, err)

	buf, err := ioutil.ReadAll(fp)
	assert.NoError(t, err)

	s, err := Parse(string(buf))
	assert.NoError(t, err)

	jout, err := s.ToJSON(true)
	assert.NoError(t, err)

	assert.NotZero(t, jout)
}

func TestRIDLTables(t *testing.T) {
	table := []struct {
		Input  string
		Output []byte
	}{
		{
			// Whitespace bug
			"webrpc = v1\n \nname = test\n   \nversion=v1.1\n",
			[]byte(`
        {
         "webrpc": "v1",
         "name": "test",
         "version": "v1.1",
         "imports": [],
         "messages": [],
         "services": []
        }
    `),
		},
		{
			"webrpc = v1\n \nname = test\n",
			[]byte(`
        {
         "webrpc": "v1",
         "name": "test",
         "version": "",
         "imports": [],
         "messages": [],
         "services": []
        }
    `),
		},
		{
			`
        webrpc = v1

        name = hello-webrpc
        version = v0.0.1

        import
          - ./blah.ridl
          - ./abc.json
      `,
			[]byte(`
        {
            "webrpc": "v1",
            "name": "hello-webrpc",
            "version": "v0.0.1",
            "imports": [
                "./blah.ridl",
                "./abc.json"
            ],
            "messages": [],
            "services": []
        }
      `),
		},
		{
			`
      webrpc = v1

      name = hello-webrpc
      version = v0.0.1


      import
        - ./blah.ridl
        - ./abc.json


      # this is a comment
      # yep
      enum Kind: uint32
        - USER = 1             # comment
        - ADMIN = 2            # comment..

      # or.. just..
      enum Kind2: uint32
        - USER                 # aka, = 0
        - ADMIN                # aka, = 1


      message Empty


      message GetUserRequest
        - userID: uint64


      message Role
        - name: string

        - users: map<uint64,User>

        - perms: []string # comment

        - other: map<uint64,map<string,string>> # comment


      message User
        - ID: uint64
          + json = id
          + go.tag.db = id

        - username: string
          + json = USERNAME
          + go.tag.db = username

        - role: string
          + go.tag.db = -

        - createdAt?: timestamp
          + json = created_at
          + go.tag.json = created_at,omitempty
          + go.tag.db = created_at


      message Notice
        - msg: string


      service ExampleService

        # comment can go here
        # too .. :)
        - Ping() => (status: bool)

        - GetUser(req: GetUserRequest) => (user: User)

        - stream Recv(req: string)

        - stream Send() => (resp: string)

        - stream SendAndRecv(req: string) => stream (resp: string)

        - Broadcast() => stream (resp: Notice)

        -  VerifyUsers(seq: int32, header?: map<string,[]string>, ids: []uint64) => (code: bool, ids: []bool)

        - MoreTest(n: uint64, stuff: []map<uint64,string>, etc: string) => (code?: bool)
      `,
			[]byte(`
        {
         "webrpc": "v1",
         "name": "hello-webrpc",
         "version": "v0.0.1",
         "imports": [
          "./blah.ridl",
          "./abc.json"
         ],
         "messages": [
          {
           "name": "Kind",
           "type": "enum",
           "fields": [
            {
             "name": "USER",
             "type": "uint32",
             "optional": false,
             "value": "1",
             "meta": null
            },
            {
             "name": "ADMIN",
             "type": "uint32",
             "optional": false,
             "value": "2",
             "meta": null
            }
           ]
          },
          {
           "name": "Kind2",
           "type": "enum",
           "fields": [
            {
             "name": "USER",
             "type": "uint32",
             "optional": false,
             "value": "0",
             "meta": null
            },
            {
             "name": "ADMIN",
             "type": "uint32",
             "optional": false,
             "value": "1",
             "meta": null
            }
           ]
          },
          {
           "name": "Empty",
           "type": "struct",
           "fields": null
          },
          {
           "name": "GetUserRequest",
           "type": "struct",
           "fields": [
            {
             "name": "userID",
             "type": "uint64",
             "optional": false,
             "value": "",
             "meta": null
            }
           ]
          },
          {
           "name": "Role",
           "type": "struct",
           "fields": [
            {
             "name": "name",
             "type": "string",
             "optional": false,
             "value": "",
             "meta": null
            },
            {
             "name": "users",
             "type": "map<uint64,User>",
             "optional": false,
             "value": "",
             "meta": null
            },
            {
             "name": "perms",
             "type": "[]string",
             "optional": false,
             "value": "",
             "meta": null
            },
            {
             "name": "other",
             "type": "map<uint64,map<string,string>>",
             "optional": false,
             "value": "",
             "meta": null
            }
           ]
          },
          {
           "name": "User",
           "type": "struct",
           "fields": [
            {
             "name": "ID",
             "type": "uint64",
             "optional": false,
             "value": "",
             "meta": [
              {
               "json": "id"
              },
              {
               "go.tag.db": "id"
              }
             ]
            },
            {
             "name": "username",
             "type": "string",
             "optional": false,
             "value": "",
             "meta": [
              {
               "json": "USERNAME"
              },
              {
               "go.tag.db": "username"
              }
             ]
            },
            {
             "name": "role",
             "type": "string",
             "optional": false,
             "value": "",
             "meta": [
              {
               "go.tag.db": "-"
              }
             ]
            },
            {
             "name": "createdAt",
             "type": "timestamp",
             "optional": true,
             "value": "",
             "meta": [
              {
               "json": "created_at"
              },
              {
               "go.tag.json": "created_at,omitempty"
              },
              {
               "go.tag.db": "created_at"
              }
             ]
            }
           ]
          },
          {
           "name": "Notice",
           "type": "struct",
           "fields": [
            {
             "name": "msg",
             "type": "string",
             "optional": false,
             "value": "",
             "meta": null
            }
           ]
          }
         ],
         "services": [
          {
           "name": "ExampleService",
           "methods": [
            {
             "name": "Ping",
             "inputs": [],
             "outputs": [
              {
               "name": "status",
               "type": "bool",
               "optional": false,
               "stream": false
              }
             ]
            },
            {
             "name": "GetUser",
             "inputs": [
              {
               "name": "req",
               "type": "GetUserRequest",
               "optional": false,
               "stream": false
              }
             ],
             "outputs": [
              {
               "name": "user",
               "type": "User",
               "optional": false,
               "stream": false
              }
             ]
            },
            {
             "name": "Recv",
             "inputs": [
              {
               "name": "req",
               "type": "string",
               "optional": false,
               "stream": false
              }
             ],
             "outputs": []
            },
            {
             "name": "Send",
             "inputs": [],
             "outputs": [
              {
               "name": "resp",
               "type": "string",
               "optional": false,
               "stream": false
              }
             ]
            },
            {
             "name": "SendAndRecv",
             "inputs": [
              {
               "name": "req",
               "type": "string",
               "optional": false,
               "stream": false
              }
             ],
             "outputs": [
              {
               "name": "resp",
               "type": "string",
               "optional": false,
               "stream": false
              }
             ]
            },
            {
             "name": "Broadcast",
             "inputs": [],
             "outputs": [
              {
               "name": "resp",
               "type": "Notice",
               "optional": false,
               "stream": false
              }
             ]
            },
            {
             "name": "VerifyUsers",
             "inputs": [
              {
               "name": "seq",
               "type": "int32",
               "optional": false,
               "stream": false
              },
              {
               "name": "header",
               "type": "map<string,[]string>",
               "optional": true,
               "stream": false
              },
              {
               "name": "ids",
               "type": "[]uint64",
               "optional": false,
               "stream": false
              }
             ],
             "outputs": [
              {
               "name": "code",
               "type": "bool",
               "optional": false,
               "stream": false
              },
              {
               "name": "ids",
               "type": "[]bool",
               "optional": false,
               "stream": false
              }
             ]
            },
            {
             "name": "MoreTest",
             "inputs": [
              {
               "name": "n",
               "type": "uint64",
               "optional": false,
               "stream": false
              },
              {
               "name": "stuff",
               "type": "[]map<uint64,string>",
               "optional": false,
               "stream": false
              },
              {
               "name": "etc",
               "type": "string",
               "optional": false,
               "stream": false
              }
             ],
             "outputs": [
              {
               "name": "code",
               "type": "bool",
               "optional": true,
               "stream": false
              }
             ]
            }
           ]
          }
         ]
        }
      `),
		},
	}

	for i := range table {
		s, err := Parse(table[i].Input)
		assert.NoError(t, err)

		jout, err := s.ToJSON(true)
		assert.NoError(t, err)

		assert.Equal(t, compactJSON(table[i].Output), compactJSON([]byte(jout)), fmt.Sprintf("GOT:\n\n%s\n\nEXPECTING:\n\n%s\n\n", jout, string(table[i].Output)))
	}
}
