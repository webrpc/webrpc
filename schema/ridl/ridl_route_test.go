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
