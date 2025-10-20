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
version = v0.1.0

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
  - UpdateUser(user: User) => (user: User) errors Unauthorized, ValidationError, ServiceError
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
