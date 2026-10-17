package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// restPlannerSchema builds a minimal schema with the given routes on one
// service, so planner tests can declare routes inline.
func restPlannerSchema(t *testing.T, routes string) *WebRPCSchema {
	t.Helper()
	s, err := ParseSchemaJSON([]byte(`{
		"webrpc": "v1",
		"name": "example",
		"version": "v0.0.1",
		"basepath": "/rpc",
		"services": [{
			"name": "Items",
			"methods": [
				{"name": "GetItem", "inputs": [
					{"name": "x", "type": "uint64"},
					{"name": "y", "type": "uint64"},
					{"name": "type", "type": "string"}
				], "outputs": [], "routes": [` + routes + `]}
			]
		}]
	}`))
	require.NoError(t, err)
	require.NoError(t, s.Validate())
	return s
}

// TestRestRouteDispatchOrder pins ServeMux-style specificity: with the same
// verb, a route with more literal segments wins, so GET /a/{x}/b serves
// /a/v/b even though GET /a/{x}/{y} also matches it.
func TestRestRouteDispatchOrder(t *testing.T) {
	s := restPlannerSchema(t, `
		{"verb": "GET", "path": "/a/{x}/{y}"},
		{"verb": "GET", "path": "/a/{x}/b"}`)
	r := RestRouter(s.Services[0])
	require.Len(t, r.Verbs, 1)
	require.Len(t, r.Verbs[0].Params, 2)
	assert.Equal(t, "/rpc/a/{x}/b", r.Verbs[0].Params[0].Path, "more literal segments first")
	assert.Equal(t, "/rpc/a/{x}/{y}", r.Verbs[0].Params[1].Path)
}

// TestRestRouteHandlerFieldNames covers the field-naming scheme: the first
// route is bare, the RPC-style alias gets RPC, other aliases get the verb,
// and same-verb aliases get a 2-based ordinal so names stay unique.
func TestRestRouteHandlerFieldNames(t *testing.T) {
	s := restPlannerSchema(t, `
		{"verb": "GET", "path": "/a"},
		{"verb": "GET", "path": "/Items/GetItem"},
		{"verb": "GET", "path": "/b"},
		{"verb": "GET", "path": "/c"},
		{"verb": "POST", "path": "/d"},
		{"verb": "POST", "path": "/e"}`)
	method := s.Services[0].Methods[0]
	require.Len(t, method.Routes, 6)

	assert.Equal(t, "restGetItem", RouteHandlerField(method.Routes[0]))
	assert.Equal(t, "restGetItemRPC", RouteHandlerField(method.Routes[1]), "RPC-style alias")
	assert.Equal(t, "restGetItemGet", RouteHandlerField(method.Routes[2]))
	assert.Equal(t, "restGetItemGet2", RouteHandlerField(method.Routes[3]), "second same-verb alias")
	assert.Equal(t, "restGetItemPost", RouteHandlerField(method.Routes[4]))
	assert.Equal(t, "restGetItemPost2", RouteHandlerField(method.Routes[5]), "second same-verb alias")
}

// TestRestAllowOverlapEntries pins the static/param overlap: the planner
// keeps one entry per static path and one per param shape, each carrying its
// own verbs; the generated restAllowMethods accumulates every entry that
// matches the request path, so the Allow value is the union (proven
// end-to-end by TestRESTAllowUnionsOverlappingRoutes).
func TestRestAllowOverlapEntries(t *testing.T) {
	s := restPlannerSchema(t, `
		{"verb": "GET", "path": "/a/{x}"},
		{"verb": "POST", "path": "/a/b"}`)
	r := RestRouter(s.Services[0])

	require.Len(t, r.Allow, 1)
	assert.Equal(t, []string{"/rpc/a/b"}, r.Allow[0].Paths)
	assert.Equal(t, "POST", r.Allow[0].Allow, "static entry carries only its own verbs")

	require.Len(t, r.Shapes, 1)
	assert.Equal(t, "GET, HEAD", r.Shapes[0].Allow, "param shape carries only its own verbs")
	// The shape matches both the static path and any other /rpc/a/{x} path,
	// so both entries accumulate for either request path.
	assert.Equal(t, "/rpc/a/", r.Shapes[0].Shape.Literals[0])
}

// TestRestVarNameKeyword sanitizes Go keywords and matcher-local names for
// generated locals, while the wire name stays untouched.
func TestRestVarNameKeyword(t *testing.T) {
	keywords := map[string]interface{}{"type": true, "var": true}
	assert.Equal(t, "type_", RestVarName("type", keywords))
	assert.Equal(t, "var_", RestVarName("var", keywords))
	assert.Equal(t, "id", RestVarName("id", keywords))
	// Reserved even without a keyword map.
	for _, name := range []string{"ok", "err", "s", "path"} {
		assert.Equal(t, name+"_", RestVarName(name, keywords), "matcher-local name")
	}
}

// TestRestRouteShapeKeywordParam keeps the schema name in the wire positions
// (SetPathValue key, restPath name) while the Go local is sanitized.
func TestRestRouteShapeKeywordParam(t *testing.T) {
	s := restPlannerSchema(t, `
		{"verb": "GET", "path": "/items/{type}"}`)
	shape := RouteShapeOf(s.Services[0].Methods[0].Routes[0])
	require.Equal(t, []string{"type"}, shape.Params, "schema name")
	assert.Equal(t, "type_", RestVarName(shape.Params[0], map[string]interface{}{"type": true}))
	assert.Equal(t, "\"/rpc/items/\"", shape.Args())
}

// TestRestSegAdjacentParams pins adjacent-param matching: /a/{x}/{y} matches
// /a/v/w, which the generated restSeg2 handles via an empty middle literal.
func TestRestSegAdjacentParams(t *testing.T) {
	s := restPlannerSchema(t, `
		{"verb": "GET", "path": "/a/{x}/{y}"}`)
	shape := RouteShapeOf(s.Services[0].Methods[0].Routes[0])
	assert.Equal(t, []string{"/rpc/a/", ""}, shape.Literals[:2])
	assert.Equal(t, `"/rpc/a/", ""`, shape.Args(), "empty middle reaches restSeg2")
	// The generated helper cuts on "/" for an empty middle; assert via the
	// restSeg1-style semantics on the shape: 2 params, prefix "/rpc/a/".
	assert.Equal(t, 2, shape.SegN())
}
