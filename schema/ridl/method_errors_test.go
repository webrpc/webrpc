package ridl

import (
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

func TestMethodErrorsKeyword_SingleError(t *testing.T) {
	ridlContent := `webrpc = v1

name = MethodErrorsTest
version = v1.0.0
basepath = /rpc

# Define a custom error
error 1 UserNotFound "user not found" HTTP 404

struct User
  - id: uint64
  - username: string

service UserService
  # Method that scopes specific potential errors returned
  - GetUser(userID: uint64) => (user: User) errors UserNotFound
`

	// Now the parser should succeed since we implemented 'errors' support!
	schema, err := parseRIDLString(ridlContent)

	require.NoError(t, err, "Parser should now succeed with 'errors' keyword")
	require.NotNil(t, schema, "Schema should not be nil")

	// Validate that the schema has the expected service
	require.Len(t, schema.Services, 1, "Should have exactly one service")
	service := schema.Services[0]
	require.Equal(t, "UserService", service.Name)

	// Validate that the method has the expected error specification
	require.Len(t, service.Methods, 1, "Should have exactly one method")
	getUserMethod := service.Methods[0]
	require.Equal(t, "GetUser", getUserMethod.Name)

	// This is the key assertion - the method should have the error specified
	require.Len(t, getUserMethod.Errors, 1, "Method should throw exactly one error")
	require.Equal(t, "UserNotFound", getUserMethod.Errors[0], "Method should throw UserNotFound error")

	t.Logf("✅ Test passes! Method errors: %v", getUserMethod.Errors)
}

func TestMethodErrorsKeyword_MultipleErrors(t *testing.T) {
	ridlContent := `webrpc = v1

name = MultipleMethodErrorsTest
version = v0.1.0

# Define multiple custom errors
error 1 Unauthorized    "unauthorized"        HTTP 401
error 2 ValidationError "validation failed"   HTTP 400
error 3 ServiceError    "service error"       HTTP 500

struct User
  - id: uint64
  - username: string

service UserService
  # Method that errors multiple specific errors
  - UpdateUser(user: User) => (user: User) errors Unauthorized | ValidationError | ServiceError
`

	// Now the parser should succeed since we implemented 'errors' support!
	schema, err := parseRIDLString(ridlContent)

	require.NoError(t, err, "Parser should now succeed with 'errors' keyword")
	require.NotNil(t, schema, "Schema should not be nil")

	// Validate that the schema has the expected service
	require.Len(t, schema.Services, 1, "Should have exactly one service")
	service := schema.Services[0]
	require.Equal(t, "UserService", service.Name)

	// Validate that the method has the expected error specification
	require.Len(t, service.Methods, 1, "Should have exactly one method")
	updateMethod := service.Methods[0]
	require.Equal(t, "UpdateUser", updateMethod.Name)

	// This is the key assertion - the method should have multiple errors specified
	require.Len(t, updateMethod.Errors, 3, "Method should throw exactly three errors")
	assert.Equal(t, []string{"Unauthorized", "ValidationError", "ServiceError"}, updateMethod.Errors)

	t.Logf("✅ Test passes! Method errors: %v", updateMethod.Errors)
}

func TestMethodErrorsKeyword_BackwardCompatibility(t *testing.T) {
	ridlContent := `webrpc = v1

name = BackwardCompatTest
version = v0.1.0

# Define some errors (even though we won't use them in errors)
error 1 SomeError "some error" HTTP 400

struct User
  - id: uint64
  - username: string

service UserService
  # Method without errors clause - should work normally
  - GetUser(userID: uint64) => (user: User)
  # Another method without errors
  - ListUsers() => (users: []User)
`

	// This should SUCCEED because methods without 'errors' should work as before
	schema, err := parseRIDLString(ridlContent)

	require.NoError(t, err, "Methods without errors should parse successfully")
	require.NotNil(t, schema, "Schema should not be nil")

	// Validate the schema structure
	require.Len(t, schema.Services, 1)
	service := schema.Services[0]
	require.Equal(t, "UserService", service.Name)
	require.Len(t, service.Methods, 2)

	// Check first method
	getUserMethod := service.Methods[0]
	require.Equal(t, "GetUser", getUserMethod.Name)
	require.Len(t, getUserMethod.Errors, 0, "Method without errors should have no errors")

	// Check second method
	listMethod := service.Methods[1]
	require.Equal(t, "ListUsers", listMethod.Name)
	require.Len(t, listMethod.Errors, 0, "Method without errors should have no errors")

	t.Logf("✅ Backward compatibility maintained - methods without errors work fine")
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
			name: "merges with the errors clause",
			methods: `  - GetUser(id: uint64) => (user: User) errors ErrOne
    + error ErrTwo
`,
			wantErrors: []string{"ErrOne", "ErrTwo"},
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
			name: "empty name between bars",
			methods: `  - GetUser(id: uint64) => (user: User)
    + error ErrOne||ErrTwo
`,
			wantErr: "empty error name in 'ErrOne||ErrTwo'",
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
			name: "duplicate against the errors clause",
			methods: `  - GetUser(id: uint64) => (user: User) errors ErrOne
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
			name: "error names are case-sensitive",
			methods: `  - GetUser(id: uint64) => (user: User)
    + error errOne
`,
			wantErr: "references undefined error 'errOne'",
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

func TestMethodErrorsClause(t *testing.T) {
	const header = `webrpc = v1

name = MethodErrorsClauseTest
version = v1.0.0

error 1 ErrOne "one" HTTP 400
error 2 ErrTwo "two" HTTP 409

struct User
  - id: uint64

service UserService
`

	tt := []struct {
		name       string
		method     string
		wantErrors []string
		wantErr    string
	}{
		{
			name:       "list without spaces around the bars",
			method:     "  - GetUser(id: uint64) => (user: User) errors ErrOne|ErrTwo\n",
			wantErrors: []string{"ErrOne", "ErrTwo"},
		},
		{
			name:       "list with spaces around the bars",
			method:     "  - GetUser(id: uint64) => (user: User) errors ErrOne | ErrTwo\n",
			wantErrors: []string{"ErrOne", "ErrTwo"},
		},
		{
			name:    "duplicate within the clause",
			method:  "  - GetUser(id: uint64) => (user: User) errors ErrOne|ErrOne\n",
			wantErr: "method 'GetUser' declares error 'ErrOne' twice",
		},
		{
			name:    "empty name in the clause",
			method:  "  - GetUser(id: uint64) => (user: User) errors ErrOne|\n",
			wantErr: "empty error name in 'ErrOne|'",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s, err := parseRIDLString(header + tc.method)
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

func TestMethodErrorsEmptyNameReportsSegment(t *testing.T) {
	const src = `webrpc = v1

name = MethodErrorsPositionTest
version = v1.0.0

error 1 ErrOne "one" HTTP 400

struct User
  - id: uint64

service UserService
  - GetUser(id: uint64) => (user: User)
    + error ErrOne|
`
	_, err := parseRIDLString(src)
	require.Error(t, err)
	// The error points at the rejected list, not at the line break after it,
	// and is not wrapped in the generic '+ error' hint.
	assert.Contains(t, err.Error(), `error near "ErrOne|"`)
	assert.NotContains(t, err.Error(), "expecting error name after '+ error'")
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
