package ridl

import (
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
  path = /users

  - Get(GetUserRequest) => (GetUserResponse)
      GET /{userId}
  - List(ListUsersRequest) => (ListUsersResponse)
      GET /
  - Create(CreateUserRequest) => (CreateUserResponse)
      POST /
  - Remove(userId: uint64)
      DELETE /{userId}
  - GetFriend(userId: uint64, friendId: uint64) => (friend: User)
      GET /{userId}/friends/{friendId}
  - Ping()
`)
	require.NoError(t, err)
	require.Len(t, s.Services, 1)

	users := s.Services[0]
	assert.Equal(t, "Users", users.Name)
	assert.Equal(t, "/users", users.Path)
	require.Len(t, users.Methods, 6)

	get := users.Methods[0]
	require.NotNil(t, get.Route)
	assert.Equal(t, "GET", get.Route.Verb)
	assert.Equal(t, "/{userId}", get.Route.Path)
	assert.Equal(t, "/api/users/{userId}", schema.MethodRoutePath(get))

	list := users.Methods[1]
	require.NotNil(t, list.Route)
	assert.Equal(t, "GET", list.Route.Verb)
	assert.Equal(t, "/", list.Route.Path)
	assert.Equal(t, "/api/users", schema.MethodRoutePath(list))

	create := users.Methods[2]
	require.NotNil(t, create.Route)
	assert.Equal(t, "POST", create.Route.Verb)

	remove := users.Methods[3]
	require.NotNil(t, remove.Route)
	assert.Equal(t, "DELETE", remove.Route.Verb)
	assert.Equal(t, "/{userId}", remove.Route.Path)

	friend := users.Methods[4]
	require.NotNil(t, friend.Route)
	assert.Equal(t, "/api/users/{userId}/friends/{friendId}", schema.MethodRoutePath(friend))

	ping := users.Methods[5]
	assert.Nil(t, ping.Route)
}

func TestRIDLRouteWithAnnotationsAndComments(t *testing.T) {
	s, err := parseRouteRIDL(t, routeRIDLHeader+`
service Users
  path = /users # user endpoints

  # fetch a single user
  @auth:required
  - Get(GetUserRequest) => (GetUserResponse)
      GET /{userId} # by id
  - Create(CreateUserRequest) => (CreateUserResponse)
`)
	require.NoError(t, err)

	users := s.Services[0]
	assert.Equal(t, "/users", users.Path)

	get := users.Methods[0]
	require.NotNil(t, get.Route)
	assert.Equal(t, "GET", get.Route.Verb)
	assert.Equal(t, "/{userId}", get.Route.Path)
	require.Contains(t, get.Annotations, "auth")
	assert.Equal(t, "required", get.Annotations["auth"].Value)

	create := users.Methods[1]
	assert.Nil(t, create.Route)
}

func TestRIDLRouteQueryVerb(t *testing.T) {
	s, err := parseRouteRIDL(t, routeRIDLHeader+`
service Users
  path = /users

  - List(ListUsersRequest) => (ListUsersResponse)
      QUERY /
`)
	require.NoError(t, err)
	assert.Equal(t, "QUERY", s.Services[0].Methods[0].Route.Verb)
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
  path = /users
  GET /{userId}
  - Get(GetUserRequest) => (GetUserResponse)
`,
			wantErr: "must be declared under a service method",
		},
		{
			name: "duplicate route on a method",
			input: routeRIDLHeader + `
service Users
  - Get(GetUserRequest) => (GetUserResponse)
      GET /users/{userId}
      GET /people/{userId}
`,
			wantErr: "already declares a route",
		},
		{
			name: "duplicate service path",
			input: routeRIDLHeader + `
service Users
  path = /users
  path = /people
  - Get(GetUserRequest) => (GetUserResponse)
`,
			wantErr: "previously declared",
		},
		{
			name: "service path after methods",
			input: routeRIDLHeader + `
service Users
  - Get(GetUserRequest) => (GetUserResponse)
  path = /users
`,
			wantErr: "must be declared before the service methods",
		},
		{
			name: "route on stream method",
			input: routeRIDLHeader + `
service Users
  - Watch(GetUserRequest) => stream (GetUserResponse)
      GET /users/{userId}/watch
`,
			wantErr: "stream methods are POST-only",
		},
		{
			name: "unknown verb ends the service block",
			input: routeRIDLHeader + `
service Users
  - Get(GetUserRequest) => (GetUserResponse)
      FETCH /users/{userId}
`,
			wantErr: `error near "FETCH": unexpected token`,
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
