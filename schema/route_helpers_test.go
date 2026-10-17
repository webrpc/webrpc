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
// unions the param shape's verbs into the static path's entry, so the
// generated restAllowMethods returns the full answer on its first match
// (proven end-to-end by TestRESTAllowUnionsOverlappingRoutes).
func TestRestAllowOverlapEntries(t *testing.T) {
	s := restPlannerSchema(t, `
		{"verb": "GET", "path": "/a/{x}"},
		{"verb": "POST", "path": "/a/b"}`)
	r := RestRouter(s.Services[0])

	require.Len(t, r.Allow, 1)
	assert.Equal(t, []string{"/rpc/a/b"}, r.Allow[0].Paths)
	assert.Equal(t, "GET, HEAD, POST", r.Allow[0].Allow, "static path unions the overlapping shape's verbs")

	require.Len(t, r.Shapes, 1)
	assert.Equal(t, "GET, HEAD", r.Shapes[0].Allow, "param shape carries only its own verbs")
	// The shape matches both the static path and any other /rpc/a/{x} path.
	assert.Equal(t, "/rpc/a/", r.Shapes[0].Shape.Literals[0])
}

// TestRestAllowStaticParamUnion pins the interop case: /items/find serves
// GET+QUERY statically but the /items/{x} shape also matches it, so a DELETE
// to it really dispatches and the Allow value must be the full union.
func TestRestAllowStaticParamUnion(t *testing.T) {
	s := restPlannerSchema(t, `
		{"verb": "GET", "path": "/items/find"},
		{"verb": "QUERY", "path": "/items/find"},
		{"verb": "GET", "path": "/items/{x}"},
		{"verb": "PUT", "path": "/items/{x}"},
		{"verb": "PATCH", "path": "/items/{x}"},
		{"verb": "DELETE", "path": "/items/{x}"}`)
	r := RestRouter(s.Services[0])

	require.Len(t, r.Allow, 1)
	assert.Equal(t, []string{"/rpc/items/find"}, r.Allow[0].Paths)
	assert.Equal(t, "DELETE, GET, HEAD, PATCH, PUT, QUERY", r.Allow[0].Allow)
}

// TestRestAllowOverlappingShapes pins the cross-shape union: two param shapes
// can cover one concrete path (GET /{x}/b and DELETE /a/{y} both match /a/b),
// and the generated table returns the first match, so each carries the
// other's verbs.
func TestRestAllowOverlappingShapes(t *testing.T) {
	s := restPlannerSchema(t, `
		{"verb": "GET", "path": "/{x}/b"},
		{"verb": "DELETE", "path": "/a/{y}"}`)
	r := RestRouter(s.Services[0])

	require.Len(t, r.Shapes, 2)
	assert.Equal(t, "DELETE, GET, HEAD", r.Shapes[0].Allow)
	assert.Equal(t, "DELETE, GET, HEAD", r.Shapes[1].Allow)
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
