package ridl

import (
	"strings"
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/webrpc/webrpc/schema"
)

func parseRIDLString(src string) (*schema.WebRPCSchema, error) {
	fsys := fstest.MapFS{
		"test.ridl": {
			Data: []byte(src),
		},
	}
	return NewParser(fsys, "/", "test.ridl").Parse()
}

func TestMethodErrorsKeyword_JsonSchemaSupport(t *testing.T) {
	jsonContent := `{
  "webrpc": "v1",
  "name": "JsonMethodErrorsTest",
  "version": "v1.0.0",
  "types": [
    {
      "kind": "struct",
      "name": "User", 
      "fields": [
        {"name": "id", "type": "uint64"},
        {"name": "username", "type": "string"}
      ]
    }
  ],
  "errors": [
    {"code": 1, "name": "UserNotFound", "message": "user not found", "httpStatus": 404},
    {"code": 2, "name": "Unauthorized", "message": "unauthorized", "httpStatus": 401}
  ],
  "services": [
    {
      "name": "UserService",
      "methods": [
        {
          "name": "GetUser",
          "inputs": [{"name": "userID", "type": "uint64", "optional": false}],
          "outputs": [{"name": "user", "type": "User", "optional": false}],
          "errors": ["UserNotFound"]
        },
        {
          "name": "UpdateUser", 
          "inputs": [{"name": "user", "type": "User", "optional": false}],
          "outputs": [{"name": "user", "type": "User", "optional": false}],
          "errors": ["UserNotFound", "Unauthorized"]
        },
        {
          "name": "ListUsers",
          "inputs": [],
          "outputs": [{"name": "users", "type": "[]User", "optional": false}]
        }
      ]
    }
  ]
}`

	// Parse JSON schema with method errors
	parsedSchema, err := schema.ParseSchemaJSON([]byte(jsonContent))
	require.NoError(t, err, "JSON schema with method errors should parse successfully")
	require.NotNil(t, parsedSchema, "Schema should not be nil")

	// Validate service and methods
	require.Len(t, parsedSchema.Services, 1)
	service := parsedSchema.Services[0]
	require.Equal(t, "UserService", service.Name)
	require.Len(t, service.Methods, 3)

	// Check GetUser method (single error)
	getUserMethod := service.Methods[0]
	require.Equal(t, "GetUser", getUserMethod.Name)
	require.Len(t, getUserMethod.Errors, 1)
	require.Equal(t, "UserNotFound", getUserMethod.Errors[0])

	// Check UpdateUser method (multiple errors)
	updateMethod := service.Methods[1]
	require.Equal(t, "UpdateUser", updateMethod.Name)
	require.Len(t, updateMethod.Errors, 2)
	require.Equal(t, []string{"UserNotFound", "Unauthorized"}, updateMethod.Errors)

	// Check ListUsers method (no errors)
	listMethod := service.Methods[2]
	require.Equal(t, "ListUsers", listMethod.Name)
	require.Len(t, listMethod.Errors, 0)

	t.Logf("✅ JSON schema support works! Methods have proper errors specifications")
}

func TestMethodErrorsOption(t *testing.T) {
	const header = `webrpc = v1

name = MethodErrorsOptionTest
version = v1.0.0

error 1 ErrOne   "one"   HTTP 400
error 2 ErrTwo   "two"   HTTP 409
error 3 ErrThree "three" HTTP 500

struct User
  - id: uint64

service UserService
`

	tt := []struct {
		name       string
		methods    string
		wantErrors []string
		wantErr    string
	}{
		{
			name: "single error",
			methods: `  - GetUser(id: uint64) => (user: User)
    + error ErrOne
`,
			wantErrors: []string{"ErrOne"},
		},
		{
			name: "several errors on one line",
			methods: `  - GetUser(id: uint64) => (user: User)
    + error ErrOne|ErrTwo|ErrThree
`,
			wantErrors: []string{"ErrOne", "ErrTwo", "ErrThree"},
		},
		{
			name: "several errors with spaces around the bars",
			methods: `  - GetUser(id: uint64) => (user: User)
    + error ErrOne | ErrTwo | ErrThree
`,
			wantErrors: []string{"ErrOne", "ErrTwo", "ErrThree"},
		},
		{
			name: "several error lines append in order",
			methods: `  - GetUser(id: uint64) => (user: User)
    + error ErrOne
    + error ErrTwo|ErrThree
`,
			wantErrors: []string{"ErrOne", "ErrTwo", "ErrThree"},
		},

		{
			name: "duplicate within one line",
			methods: `  - GetUser(id: uint64) => (user: User)
    + error ErrOne|ErrOne
`,
			wantErr: "method 'GetUser' declares error 'ErrOne' twice",
		},
		{
			name: "empty name after a trailing bar",
			methods: `  - GetUser(id: uint64) => (user: User)
    + error ErrOne|
`,
			wantErr: "empty error name in 'ErrOne|'",
		},
		{
			name: "duplicate across lines",
			methods: `  - GetUser(id: uint64) => (user: User)
    + error ErrOne
    + error ErrOne
`,
			wantErr: "method 'GetUser' declares error 'ErrOne' twice",
		},
		{
			name: "unknown error name",
			methods: `  - GetUser(id: uint64) => (user: User)
    + error ErrFour
`,
			wantErr: "references undefined error 'ErrFour'",
		},
		{
			name: "before any method",
			methods: `  + error ErrOne
  - GetUser(id: uint64) => (user: User)
`,
			wantErr: "'+ error' must be declared under a service method",
		},
		{
			name: "missing error name",
			methods: `  - GetUser(id: uint64) => (user: User)
    + error
`,
			wantErr: "expecting error name after '+ error'",
		},
		{
			name: "trailing comment",
			methods: `  - GetUser(id: uint64) => (user: User)
    + error ErrOne  # scope the error
`,
			wantErrors: []string{"ErrOne"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s, err := parseRIDLString(header + tc.methods)
			if tc.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.wantErr)
				return
			}

			require.NoError(t, err)
			require.Len(t, s.Services[0].Methods, 1)
			assert.Equal(t, tc.wantErrors, s.Services[0].Methods[0].Errors)
		})
	}
}

func TestMethodErrorsClauseRejected(t *testing.T) {
	const header = `webrpc = v1

name = MethodErrorsClauseTest
version = v1.0.0

error 1 ErrOne "one" HTTP 400
error 2 ErrTwo "two" HTTP 409

struct User
  - id: uint64

service UserService
`

	// Every clause form is rejected with the migration hint, whatever the
	// separator or list contents.
	methods := []string{
		"  - GetUser(id: uint64) => (user: User) errors ErrOne|ErrTwo\n",
		"  - GetUser(id: uint64) => (user: User) errors ErrOne | ErrTwo\n",
		"  - GetUser(id: uint64) => (user: User) errors ErrOne|ErrOne\n",
		"  - GetUser(id: uint64) => (user: User) errors ErrOne|\n",
	}

	const wantErr = "'errors A|B' after the method signature is no longer supported: declare errors below the method as '+ error A|B'"
	for _, method := range methods {
		t.Run(strings.TrimSpace(method), func(t *testing.T) {
			_, err := parseRIDLString(header + method)
			require.Error(t, err)
			assert.Contains(t, err.Error(), wantErr)
		})
	}
}

func TestErrorDeclarationRejectsBarInName(t *testing.T) {
	const src = `webrpc = v1

name = ErrorNamePipeTest
version = v1.0.0

error 1 Foo|Bar "msg" HTTP 400
`
	_, err := parseRIDLString(src)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "error name may not contain '|'")
}
