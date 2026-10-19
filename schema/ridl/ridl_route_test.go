package ridl

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/webrpc/webrpc/schema"
)

func parseRouteRIDL(t *testing.T, input string) (*schema.WebRPCSchema, error) {
	t.Helper()
	return parseString(input)
}

const routeRIDLHeader = `webrpc = v1
name = example
version = v0.0.1
basepath = /api

struct User
  - id: uint64
  - name: string

struct GetUserRequest
  - userId: uint64
  - details?: bool

struct GetUserResponse
  - user: User

struct ListUsersRequest
  - q?: string
  - tags: []string

struct ListUsersResponse
  - users: []User

struct CreateUserRequest
  - user: User

struct CreateUserResponse
  - user: User

`

func TestRIDLServiceRoutes(t *testing.T) {
	s, err := parseRouteRIDL(t, routeRIDLHeader+`
service Users
  - Get(GetUserRequest) => (GetUserResponse)
      + GET /users/{userId}
  - List(ListUsersRequest) => (ListUsersResponse)
      + GET /users
  - Create(CreateUserRequest) => (CreateUserResponse)
      + POST /users
  - Remove(userId: uint64)
      + DELETE /users/{userId}
  - GetFriend(userId: uint64, friendId: uint64) => (friend: User)
      + GET /users/{userId}/friends/{friendId}
  - Ping()
`)
	require.NoError(t, err)
	require.Len(t, s.Services, 1)

	users := s.Services[0]
	assert.Equal(t, "Users", users.Name)
	require.Len(t, users.Methods, 6)

	get := users.Methods[0]
	require.Len(t, get.Routes, 1)
	assert.Equal(t, "GET", get.Routes[0].Verb)
	assert.Equal(t, "/users/{userId}", get.Routes[0].Path)
	assert.Equal(t, "/api/users/{userId}", schema.RoutePath(get.Routes[0]))

	list := users.Methods[1]
	assert.Equal(t, "/api/users", schema.RoutePath(list.Routes[0]))

	assert.Equal(t, "POST", users.Methods[2].Routes[0].Verb)
	assert.Equal(t, "DELETE", users.Methods[3].Routes[0].Verb)
	assert.Equal(t, "/api/users/{userId}/friends/{friendId}", schema.RoutePath(users.Methods[4].Routes[0]))

	assert.Empty(t, users.Methods[5].Routes)
}

func TestRIDLMultipleRoutesKeepOrder(t *testing.T) {
	s, err := parseRouteRIDL(t, strings.Replace(routeRIDLHeader, "/api", "/v1", 1)+`
service Intents
  - Get(GetUserRequest) => (GetUserResponse)
      + GET /intents/{userId}
      + POST /Intents/Get     # backward-compatible with pre-REST API route.
  - Ping()
`)
	require.NoError(t, err)

	get := s.Services[0].Methods[0]
	require.Len(t, get.Routes, 2)
	assert.Same(t, get.Routes[0], schema.MethodCanonicalRoute(get))
	assert.Equal(t, "GET", get.Routes[0].Verb)
	assert.Equal(t, "/intents/{userId}", get.Routes[0].Path)
	assert.Equal(t, "POST", get.Routes[1].Verb)
	assert.Equal(t, "/v1/Intents/Get", schema.RoutePath(get.Routes[1]))
	assert.Empty(t, s.Services[0].Methods[1].Routes)
}

func TestRIDLRouteWithAnnotationsAndComments(t *testing.T) {
	s, err := parseRouteRIDL(t, routeRIDLHeader+`
service Users
  # fetch a single user
  @auth:required
  - Get(GetUserRequest) => (GetUserResponse)
      + GET /users/{userId} # by id
  - Create(CreateUserRequest) => (CreateUserResponse)
`)
	require.NoError(t, err)

	users := s.Services[0]
	get := users.Methods[0]
	require.Len(t, get.Routes, 1)
	assert.Equal(t, "GET", get.Routes[0].Verb)
	assert.Equal(t, "/users/{userId}", get.Routes[0].Path)
	require.Contains(t, get.Annotations, "auth")
	assert.Equal(t, "required", get.Annotations["auth"].Value)

	assert.Empty(t, users.Methods[1].Routes)
}

func TestRIDLRouteAnnotationAttachesToNextMethod(t *testing.T) {
	s, err := parseRouteRIDL(t, routeRIDLHeader+`
service Users
  - Get(GetUserRequest) => (GetUserResponse)
      + GET /users/{userId}
  @auth:required
  - List(ListUsersRequest) => (ListUsersResponse)
`)
	require.NoError(t, err)

	users := s.Services[0]
	assert.Empty(t, users.Methods[0].Annotations)
	require.Contains(t, users.Methods[1].Annotations, "auth")
	assert.Equal(t, "required", users.Methods[1].Annotations["auth"].Value)
	assert.Empty(t, users.Methods[1].Routes)
}

func TestRIDLRouteAfterErrorsClause(t *testing.T) {
	s, err := parseRouteRIDL(t, `webrpc = v1
name = example
version = v0.0.1

error 100 NotFound "not found" HTTP 404

service Users
  - Get(userId: uint64) => (name: string) errors NotFound
      + GET /users/{userId}
`)
	require.NoError(t, err)
	get := s.Services[0].Methods[0]
	assert.Equal(t, []string{"NotFound"}, get.Errors)
	require.Len(t, get.Routes, 1)
}

func TestRIDLRouteGetAndQueryOnOneMethod(t *testing.T) {
	s, err := parseRouteRIDL(t, routeRIDLHeader+`
service Users
  - List(ListUsersRequest) => (ListUsersResponse)
      + GET /users
      + QUERY /users
`)
	require.NoError(t, err)
	routes := s.Services[0].Methods[0].Routes
	require.Len(t, routes, 2)
	assert.Equal(t, "GET", routes[0].Verb)
	assert.Equal(t, "QUERY", routes[1].Verb)
}

func TestRIDLLegacyRouteAliasDoesNotConflictWithOwnDefaultPath(t *testing.T) {
	_, err := parseRouteRIDL(t, routeRIDLHeader+`
service Repos
  - GetRepo(org: string, repo: string) => (name: string)
      + GET /{org}/{repo}
      + POST /Repos/GetRepo
`)
	require.NoError(t, err)
}

func TestRIDLRouteErrors(t *testing.T) {
	tt := []struct {
		name    string
		input   string
		wantErr string
	}{
		{
			name: "route before any method",
			input: routeRIDLHeader + `
service Users
  + GET /users/{userId}
  - Get(GetUserRequest) => (GetUserResponse)
`,
			wantErr: "must be declared under a service method",
		},
		{
			name: "old bare route line",
			input: routeRIDLHeader + `
service Users
  - Get(GetUserRequest) => (GetUserResponse)
      GET /users/{userId}
`,
			wantErr: "route 'GET' must start with '+'",
		},
		{
			name: "service-level path",
			input: routeRIDLHeader + `
service Users
  path = /users
  - Get(GetUserRequest) => (GetUserResponse)
`,
			wantErr: "service-level 'path' is no longer supported",
		},
		{
			name: "service-level path after methods",
			input: routeRIDLHeader + `
service Users
  - Get(GetUserRequest) => (GetUserResponse)
  path = /users
`,
			wantErr: "service-level 'path' is no longer supported",
		},
		{
			name: "unknown verb",
			input: routeRIDLHeader + `
service Users
  - Get(GetUserRequest) => (GetUserResponse)
      + FETCH /users/{userId}
`,
			wantErr: `expecting route '+ VERB /path'`,
		},
		{
			name: "missing path",
			input: routeRIDLHeader + `
service Users
  - Get(GetUserRequest) => (GetUserResponse)
      + GET
`,
			wantErr: "expecting route path after 'GET'",
		},
		{
			name: "relative path",
			input: routeRIDLHeader + `
service Users
  - Get(GetUserRequest) => (GetUserResponse)
      + GET users/{userId}
`,
			wantErr: "must start with '/'",
		},
		{
			name: "dot-dot segment in route",
			input: routeRIDLHeader + `
service Items
  - Lookup() => ()
      + GET /items/../lookup
`,
			wantErr: "dot segments are not allowed",
		},
		{
			name: "plus without verb",
			input: routeRIDLHeader + `
service Users
  - Get(GetUserRequest) => (GetUserResponse)
      +
`,
			wantErr: "expecting route verb after '+'",
		},
		{
			name: "unclosed param in route",
			input: routeRIDLHeader + `
service Users
  - Get(GetUserRequest) => (GetUserResponse)
      + GET /users/{userId
`,
			wantErr: "invalid route of method 'Get'",
		},
		{
			name: "empty param in route",
			input: routeRIDLHeader + `
service Users
  - Get(GetUserRequest) => (GetUserResponse)
      + GET /users/{}
`,
			wantErr: "invalid route of method 'Get'",
		},
		{
			name: "annotation between method and route",
			input: routeRIDLHeader + `
service Users
  - Get(GetUserRequest) => (GetUserResponse)
      @auth:x
      + GET /users/{userId}
`,
			wantErr: "route must directly follow its method",
		},
		{
			name: "annotation above a success route line",
			input: routeRIDLHeader + `
service Users
  - Get(GetUserRequest) => (GetUserResponse)
      @auth:x
      + @success 201
      + GET /users/{userId}
`,
			wantErr: "annotation '@auth' must be placed directly above a '- Method' line",
		},
		{
			name: "annotation after last method's routes at end of service",
			input: routeRIDLHeader + `
service Users
  - Get(GetUserRequest) => (GetUserResponse)
      + GET /users/{userId}
  @auth:x
`,
			wantErr: "annotation '@auth' must be placed directly above a '- Method' line",
		},
		{
			name: "annotation after routeless last method at end of service",
			input: routeRIDLHeader + `
service Users
  - Get(GetUserRequest) => (GetUserResponse)
  @auth:x
`,
			wantErr: "annotation '@auth' must be placed directly above a '- Method' line",
		},
		{
			name: "annotation after last method's routes at EOF without trailing newline",
			input: routeRIDLHeader + `
service Users
  - Get(GetUserRequest) => (GetUserResponse)
      + GET /users/{userId}
  @auth:x`,
			wantErr: "annotation '@auth' must be placed directly above a '- Method' line",
		},
		{
			name: "route on stream method",
			input: routeRIDLHeader + `
service Users
  - Watch(GetUserRequest) => stream (GetUserResponse)
      + GET /users/{userId}/watch
`,
			wantErr: "stream methods are POST-only",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := parseRouteRIDL(t, tc.input)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tc.wantErr)
		})
	}
}

func TestRIDLRouteSuccessStatus(t *testing.T) {
	s, err := parseRouteRIDL(t, routeRIDLHeader+`
service Users
  - Create(CreateUserRequest) => (CreateUserResponse)
      + POST /users
      + @success 201
  - Commit(GetUserRequest) => (GetUserResponse)
      + POST /commits
      + @success 202
      + POST /Users/Commit     # the alias keeps the default 200
  - Upsert(User) => (User)
      + PUT /users
      + @success 200|201
  - Delete()
      + DELETE /users
      + @success 200|204
`)
	require.NoError(t, err)

	create := s.Services[0].Methods[0]
	require.Len(t, create.Routes, 1)
	assert.Equal(t, []int{201}, create.Routes[0].Success)

	// The method's @success set fills every route of the method.
	commit := s.Services[0].Methods[1]
	require.Len(t, commit.Routes, 2)
	assert.Equal(t, []int{202}, commit.Routes[0].Success)
	assert.Equal(t, []int{202}, commit.Routes[1].Success)

	// A pipe-separated set declares every allowed success code; the first is
	// the default. 204 requires an empty response.
	upsert := s.Services[0].Methods[2]
	require.Len(t, upsert.Routes, 1)
	assert.Equal(t, []int{200, 201}, upsert.Routes[0].Success)
	assert.Equal(t, 200, schema.RouteSuccessStatus(upsert.Routes[0]))
	deleteMethod := s.Services[0].Methods[3]
	require.Len(t, deleteMethod.Routes, 1)
	assert.Equal(t, []int{200, 204}, deleteMethod.Routes[0].Success)
}

func TestRIDLRouteSuccessStatusErrors(t *testing.T) {
	tt := []struct {
		name    string
		input   string
		wantErr string
	}{
		{
			name: "method without routes",
			input: routeRIDLHeader + `
service Users
  - Create(CreateUserRequest) => (CreateUserResponse)
      + @success 201
`,
			wantErr: "@success requires at least one route on method 'Create'",
		},
		{
			name: "before any method",
			input: routeRIDLHeader + `
service Users
  + @success 201
  - Create(CreateUserRequest) => (CreateUserResponse)
`,
			wantErr: "@success must be declared under a service method",
		},
		{
			name: "non-2xx code",
			input: routeRIDLHeader + `
service Users
  - Create(CreateUserRequest) => (CreateUserResponse)
      + POST /users
      + @success 404
`,
			wantErr: "invalid @success status code '404' of method 'Create': must be a 2xx code (200-299)",
		},
		{
			name: "1xx code",
			input: routeRIDLHeader + `
service Users
  - Create(CreateUserRequest) => (CreateUserResponse)
      + POST /users
      + @success 100
`,
			wantErr: "must be a 2xx code (200-299)",
		},
		{
			name: "204 on a non-empty response",
			input: routeRIDLHeader + `
service Users
  - Create(CreateUserRequest) => (CreateUserResponse)
      + POST /users
      + @success 204
`,
			wantErr: "invalid @success status code '204' of route 'POST /users' of method 'Users.Create': 204/205 require an empty response (no outputs or an empty struct)",
		},
		{
			name: "205 on a multi-output response",
			input: routeRIDLHeader + `
service Users
  - Create() => (a: string, b: string)
      + POST /users
      + @success 205
`,
			wantErr: "invalid @success status code '205' of method 'Create': 204/205 require an empty response (no outputs or an empty struct)",
		},
		{
			name: "duplicate code in the set",
			input: routeRIDLHeader + `
service Users
  - Create(CreateUserRequest) => (CreateUserResponse)
      + POST /users
      + @success 200|200
`,
			wantErr: "duplicate @success status code '200' of method 'Create'",
		},
		{
			name: "non-2xx code in the set",
			input: routeRIDLHeader + `
service Users
  - Create(CreateUserRequest) => (CreateUserResponse)
      + POST /users
      + @success 200|404
`,
			wantErr: "invalid @success status code '404' of method 'Create': must be a 2xx code (200-299)",
		},
		{
			name: "non-numeric code in the set",
			input: routeRIDLHeader + `
service Users
  - Create(CreateUserRequest) => (CreateUserResponse)
      + POST /users
      + @success 200|abc
`,
			wantErr: "expecting '@success' status code to be a number but got '200|abc'",
		},
		{
			name: "non-numeric code",
			input: routeRIDLHeader + `
service Users
  - Create(CreateUserRequest) => (CreateUserResponse)
      + POST /users
      + @success created
`,
			wantErr: "expecting '@success' status code to be a number but got 'created'",
		},
		{
			name: "missing code",
			input: routeRIDLHeader + `
service Users
  - Create(CreateUserRequest) => (CreateUserResponse)
      + POST /users
      + @success
`,
			wantErr: "expecting '@success' status code to be a number",
		},
		{
			name: "unknown at-meta",
			input: routeRIDLHeader + `
service Users
  - Create(CreateUserRequest) => (CreateUserResponse)
      + POST /users
      + @created 201
`,
			wantErr: "expecting ':' before the annotation value, e.g. '+ @created:201'",
		},
		{
			name: "duplicate @success on one method",
			input: routeRIDLHeader + `
service Users
  - Create(CreateUserRequest) => (CreateUserResponse)
      + POST /users
      + @success 201
      + @success 202
`,
			wantErr: "method 'Create' already declares @success",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			_, err := parseRouteRIDL(t, tc.input)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tc.wantErr)
		})
	}
}

func TestRIDLRoutePathSyntaxErrorHasPosition(t *testing.T) {
	_, err := parseRouteRIDL(t, routeRIDLHeader+`
service Users
  - Get(GetUserRequest) => (GetUserResponse)
      + GET /users/../{userId}
`)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "dot segments are not allowed")
	assert.Regexp(t, `^\d+:\d+:`, err.Error())
}
