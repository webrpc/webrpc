package schema

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func parseRouteTestSchema(t *testing.T, basePath string, types string, services string) (*WebRPCSchema, error) {
	t.Helper()
	input := fmt.Sprintf(`{
		"webrpc": "v1",
		"name": "example",
		"version": "v0.0.1",
		"basepath": "%s",
		"types": [%s],
		"services": [%s]
	}`, basePath, types, services)
	return ParseSchemaJSON([]byte(input))
}

const userTypes = `
	{
		"kind": "struct",
		"name": "GetUserRequest",
		"fields": [
			{"name": "userId", "type": "uint64"},
			{"name": "details", "type": "bool", "optional": true}
		]
	},
	{
		"kind": "struct",
		"name": "User",
		"fields": [
			{"name": "id", "type": "uint64"},
			{"name": "name", "type": "string"}
		]
	},
	{
		"kind": "struct",
		"name": "ListUsersRequest",
		"fields": [
			{"name": "q", "type": "string", "optional": true},
			{"name": "tags", "type": "[]string"}
		]
	},
	{
		"kind": "struct",
		"name": "CreateUserRequest",
		"fields": [
			{"name": "user", "type": "User"}
		]
	}
`

func TestRouteValidSchema(t *testing.T) {
	s, err := parseRouteTestSchema(t, "/api", userTypes, `
		{
			"name": "Users",
			"path": "/users",
			"methods": [
				{
					"name": "Get",
					"succinct": true,
					"route": {"verb": "GET", "path": "/{userId}"},
					"inputs": [{"name": "getUserRequest", "type": "GetUserRequest"}],
					"outputs": [{"name": "user", "type": "User"}]
				},
				{
					"name": "List",
					"succinct": true,
					"route": {"verb": "GET", "path": "/"},
					"inputs": [{"name": "listUsersRequest", "type": "ListUsersRequest"}],
					"outputs": [{"name": "users", "type": "[]User"}]
				},
				{
					"name": "Create",
					"succinct": true,
					"route": {"verb": "POST", "path": "/"},
					"inputs": [{"name": "createUserRequest", "type": "CreateUserRequest"}],
					"outputs": [{"name": "user", "type": "User"}]
				},
				{
					"name": "Delete",
					"route": {"verb": "DELETE", "path": "/{userId}"},
					"inputs": [{"name": "userId", "type": "uint64"}],
					"outputs": []
				},
				{
					"name": "Ping",
					"inputs": [],
					"outputs": []
				}
			]
		}
	`)
	require.NoError(t, err)

	users := s.Services[0]
	assert.Equal(t, "/users", users.Path)

	get := users.Methods[0]
	require.NotNil(t, get.Route)
	assert.Equal(t, "GET", get.Route.Verb)
	assert.Equal(t, "/{userId}", get.Route.Path)
	assert.Equal(t, "/api/users/{userId}", MethodRoutePath(get))

	assert.Equal(t, []*RouteParam{
		{Name: "userId", TypeExpr: "uint64"},
		{Name: "details", TypeExpr: "bool", Optional: true},
	}, MethodRouteParams(get))

	list := users.Methods[1]
	assert.Equal(t, "/api/users", MethodRoutePath(list))
	assert.Equal(t, []*RouteParam{
		{Name: "q", TypeExpr: "string", Optional: true},
		{Name: "tags", TypeExpr: "string", Repeated: true},
	}, MethodRouteParams(list))

	// POST sends every non-path field in the JSON body.
	create := users.Methods[2]
	assert.Equal(t, "/api/users", MethodRoutePath(create))
	assert.Empty(t, MethodRouteParams(create))

	assert.Equal(t, []*RouteParam{{Name: "userId", TypeExpr: "uint64"}}, MethodRouteParams(users.Methods[3]))

	ping := users.Methods[4]
	assert.False(t, MethodHasRoute(ping))
}

func TestRouteNormalFormBinding(t *testing.T) {
	// Path and query params bind to method arguments in normal (non-succinct) form.
	s, err := parseRouteTestSchema(t, "/api", "", `
		{
			"name": "Files",
			"methods": [
				{
					"name": "Get",
					"route": {"verb": "GET", "path": "/files/{id}"},
					"inputs": [
						{"name": "id", "type": "uint64"},
						{"name": "raw", "type": "bool", "optional": true}
					],
					"outputs": [{"name": "name", "type": "string"}]
				}
			]
		}
	`)
	require.NoError(t, err)

	assert.Equal(t, []*RouteParam{
		{Name: "id", TypeExpr: "uint64"},
		{Name: "raw", TypeExpr: "bool", Optional: true},
	}, MethodRouteParams(s.Services[0].Methods[0]))
}

func TestRouteServicePathParam(t *testing.T) {
	// A path param declared in the service path binds per method.
	_, err := parseRouteTestSchema(t, "", "", `
		{
			"name": "Orgs",
			"path": "/orgs/{orgId}",
			"methods": [
				{
					"name": "GetMember",
					"route": {"verb": "GET", "path": "/members/{memberId}"},
					"inputs": [
						{"name": "orgId", "type": "string"},
						{"name": "memberId", "type": "uint64"}
					],
					"outputs": []
				}
			]
		}
	`)
	require.NoError(t, err)
}

func TestRouteInvalidSchemas(t *testing.T) {
	tt := []struct {
		name     string
		basePath string
		types    string
		services string
		wantErr  string
	}{
		{
			name: "unknown verb",
			services: `{
				"name": "Users",
				"methods": [{
					"name": "Get",
					"route": {"verb": "FETCH", "path": "/users/{id}"},
					"inputs": [{"name": "id", "type": "uint64"}],
					"outputs": []
				}]
			}`,
			wantErr: "invalid route verb 'FETCH'",
		},
		{
			name: "path without leading slash",
			services: `{
				"name": "Users",
				"methods": [{
					"name": "Get",
					"route": {"verb": "GET", "path": "users/{id}"},
					"inputs": [{"name": "id", "type": "uint64"}],
					"outputs": []
				}]
			}`,
			wantErr: "must start with '/'",
		},
		{
			name: "partial segment param",
			services: `{
				"name": "Users",
				"methods": [{
					"name": "Get",
					"route": {"verb": "GET", "path": "/user-{id}"},
					"inputs": [{"name": "id", "type": "uint64"}],
					"outputs": []
				}]
			}`,
			wantErr: "params must span a whole segment",
		},
		{
			name: "duplicate path param",
			services: `{
				"name": "Users",
				"methods": [{
					"name": "Get",
					"route": {"verb": "GET", "path": "/users/{id}/friends/{id}"},
					"inputs": [{"name": "id", "type": "uint64"}],
					"outputs": []
				}]
			}`,
			wantErr: "duplicate path param '{id}'",
		},
		{
			name: "path param without matching field",
			services: `{
				"name": "Users",
				"methods": [{
					"name": "Get",
					"route": {"verb": "GET", "path": "/users/{userId}"},
					"inputs": [{"name": "id", "type": "uint64"}],
					"outputs": []
				}]
			}`,
			wantErr: "path param '{userId}' of method 'Users.Get' does not match any request field",
		},
		{
			name: "path param bound to optional field",
			services: `{
				"name": "Users",
				"methods": [{
					"name": "Get",
					"route": {"verb": "GET", "path": "/users/{id}"},
					"inputs": [{"name": "id", "type": "uint64", "optional": true}],
					"outputs": []
				}]
			}`,
			wantErr: "must bind to a required field",
		},
		{
			name:  "path param bound to struct field",
			types: userTypes,
			services: `{
				"name": "Users",
				"methods": [{
					"name": "Create",
					"route": {"verb": "POST", "path": "/users/{user}"},
					"inputs": [{"name": "user", "type": "User"}],
					"outputs": []
				}]
			}`,
			wantErr: "must bind to a scalar field",
		},
		{
			name:  "struct query param in GET",
			types: userTypes,
			services: `{
				"name": "Users",
				"methods": [{
					"name": "Find",
					"route": {"verb": "GET", "path": "/users"},
					"inputs": [{"name": "user", "type": "User"}],
					"outputs": []
				}]
			}`,
			wantErr: "only scalars and scalar arrays are supported in GET requests",
		},
		{
			name:  "struct array query param in QUERY",
			types: userTypes,
			services: `{
				"name": "Users",
				"methods": [{
					"name": "Find",
					"route": {"verb": "QUERY", "path": "/users"},
					"inputs": [{"name": "users", "type": "[]User"}],
					"outputs": []
				}]
			}`,
			wantErr: "only scalars and scalar arrays are supported in QUERY requests",
		},
		{
			name:  "struct body field in POST is allowed",
			types: userTypes,
			services: `{
				"name": "Users",
				"methods": [{
					"name": "Create",
					"route": {"verb": "POST", "path": "/users"},
					"inputs": [{"name": "user", "type": "User"}],
					"outputs": []
				}]
			}`,
		},
		{
			name: "route on stream method",
			services: `{
				"name": "Chat",
				"methods": [{
					"name": "Subscribe",
					"streamOutput": true,
					"route": {"verb": "GET", "path": "/chat"},
					"inputs": [],
					"outputs": [{"name": "msg", "type": "string"}]
				}]
			}`,
			wantErr: "stream methods are POST-only",
		},
		{
			name: "route on file upload method",
			services: `{
				"name": "Files",
				"methods": [{
					"name": "Upload",
					"route": {"verb": "POST", "path": "/files"},
					"inputs": [{"name": "data", "type": "file"}],
					"outputs": []
				}]
			}`,
			wantErr: "methods using the file type keep the default webrpc dispatch",
		},
		{
			name: "duplicate route across services",
			services: `{
				"name": "Users",
				"methods": [{
					"name": "Get",
					"route": {"verb": "GET", "path": "/users/{id}"},
					"inputs": [{"name": "id", "type": "uint64"}],
					"outputs": []
				}]
			},
			{
				"name": "Admin",
				"methods": [{
					"name": "GetUser",
					"route": {"verb": "GET", "path": "/users/{userId}"},
					"inputs": [{"name": "userId", "type": "uint64"}],
					"outputs": []
				}]
			}`,
			wantErr: "cannot overlap across services",
		},
		{
			name: "same path different verbs in one service is allowed",
			services: `{
				"name": "Users",
				"methods": [
					{
						"name": "List",
						"route": {"verb": "GET", "path": "/users"},
						"inputs": [],
						"outputs": []
					},
					{
						"name": "Create",
						"route": {"verb": "POST", "path": "/users"},
						"inputs": [{"name": "name", "type": "string"}],
						"outputs": []
					}
				]
			}`,
		},
		{
			name: "same path across services",
			services: `{
				"name": "Users",
				"methods": [{
					"name": "List",
					"route": {"verb": "GET", "path": "/users"},
					"inputs": [],
					"outputs": []
				}]
			},
			{
				"name": "Admin",
				"methods": [{
					"name": "Create",
					"route": {"verb": "POST", "path": "/users"},
					"inputs": [{"name": "name", "type": "string"}],
					"outputs": []
				}]
			}`,
			wantErr: "cannot overlap across services",
		},
		{
			name:     "route collides with default dispatch path",
			basePath: "/rpc",
			services: `{
				"name": "Users",
				"methods": [
					{
						"name": "Get",
						"inputs": [],
						"outputs": []
					},
					{
						"name": "GetHTTP",
						"route": {"verb": "POST", "path": "/Users/Get"},
						"inputs": [],
						"outputs": []
					}
				]
			}`,
			wantErr: "overlaps the default webrpc dispatch path '/rpc/Users/Get'",
		},
		{
			name: "duplicate route in one service",
			services: `{
				"name": "Users",
				"methods": [
					{
						"name": "Get",
						"route": {"verb": "GET", "path": "/users/{id}"},
						"inputs": [{"name": "id", "type": "uint64"}],
						"outputs": []
					},
					{
						"name": "GetAgain",
						"route": {"verb": "GET", "path": "/users/{userId}"},
						"inputs": [{"name": "userId", "type": "uint64"}],
						"outputs": []
					}
				]
			}`,
			wantErr: "duplicate route 'GET /api/users/{userId}'",
		},
		{
			name: "ambiguous routes",
			services: `{
				"name": "Things",
				"methods": [
					{
						"name": "A",
						"route": {"verb": "GET", "path": "/a/{x}/c"},
						"inputs": [{"name": "x", "type": "string"}],
						"outputs": []
					},
					{
						"name": "B",
						"route": {"verb": "GET", "path": "/a/b/{y}"},
						"inputs": [{"name": "y", "type": "string"}],
						"outputs": []
					}
				]
			}`,
			wantErr: "are ambiguous",
		},
		{
			name: "static route more specific than param route is allowed",
			services: `{
				"name": "Things",
				"methods": [
					{
						"name": "A",
						"route": {"verb": "GET", "path": "/a/{x}"},
						"inputs": [{"name": "x", "type": "string"}],
						"outputs": []
					},
					{
						"name": "B",
						"route": {"verb": "GET", "path": "/a/b"},
						"inputs": [],
						"outputs": []
					}
				]
			}`,
		},
		{
			name: "overlapping paths across services with different verbs",
			services: `{
				"name": "Users",
				"methods": [{
					"name": "Get",
					"route": {"verb": "GET", "path": "/users/{id}"},
					"inputs": [{"name": "id", "type": "uint64"}],
					"outputs": []
				}]
			},
			{
				"name": "Admin",
				"methods": [{
					"name": "Special",
					"route": {"verb": "POST", "path": "/users/special"},
					"inputs": [],
					"outputs": []
				}]
			}`,
			wantErr: "cannot overlap across services",
		},
		{
			name:     "GET route shadowed by default dispatch path",
			basePath: "/rpc",
			services: `{
				"name": "Users",
				"methods": [
					{
						"name": "Get",
						"inputs": [],
						"outputs": []
					},
					{
						"name": "Fetch",
						"route": {"verb": "GET", "path": "/Users/Get"},
						"inputs": [],
						"outputs": []
					}
				]
			}`,
			wantErr: "overlaps the default webrpc dispatch path '/rpc/Users/Get'",
		},
		{
			name:     "param route overlapping another service's default dispatch path",
			basePath: "/rpc",
			services: `{
				"name": "Users",
				"methods": [{
					"name": "Ping",
					"inputs": [],
					"outputs": []
				}]
			},
			{
				"name": "Things",
				"methods": [{
					"name": "Get",
					"route": {"verb": "GET", "path": "/{svc}/Ping"},
					"inputs": [{"name": "svc", "type": "string"}],
					"outputs": []
				}]
			}`,
			wantErr: "overlaps the default webrpc dispatch path '/rpc/Users/Ping'",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			basePath := tc.basePath
			if basePath == "" {
				basePath = "/api"
			}
			_, err := parseRouteTestSchema(t, basePath, tc.types, tc.services)
			if tc.wantErr == "" {
				require.NoError(t, err)
				return
			}
			require.Error(t, err)
			assert.Contains(t, err.Error(), tc.wantErr)
		})
	}
}

func TestJoinRoutePath(t *testing.T) {
	assert.Equal(t, "/", JoinRoutePath("", "", ""))
	assert.Equal(t, "/api", JoinRoutePath("/api/", "", "/"))
	assert.Equal(t, "/api/users/{id}", JoinRoutePath("/api/", "/users", "/{id}"))
	assert.Equal(t, "/users", JoinRoutePath("", "/users", "/"))
	assert.Equal(t, "/rpc/users/{id}", JoinRoutePath("/rpc", "users", "{id}"))
}
