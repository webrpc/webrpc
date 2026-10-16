package schema

// RouteParam is a request field bound to a path param (when the route path
// template contains "{Name}") or else to a query param (GET/DELETE).
//
// This type is used as a template helper for adding REST route support in
// various codegen targets.
type RouteParam struct {
	Name     string // wire name: json field name (succinct) or argument name
	TypeExpr string // alias-resolved scalar type, the element type for lists
	Optional bool
	Repeated bool // []scalar query param, sent as repeated keys
}

// MethodHasRoute returns true if the method declares REST routes, served
// instead of its default webrpc path (POST {basepath}{Service}/{Method}).
func MethodHasRoute(m *Method) bool {
	return m != nil && len(m.Routes) > 0
}

// MethodCanonicalRoute returns the first declared route, the one generated
// clients call, or nil without routes. Later routes are server-only aliases.
func MethodCanonicalRoute(m *Method) *MethodRoute {
	if !MethodHasRoute(m) {
		return nil
	}
	return m.Routes[0]
}

// RouteHasBody reports whether the route's non-path fields travel in the
// JSON body (POST/PUT/PATCH/QUERY) rather than the query string (GET/DELETE).
func RouteHasBody(r *MethodRoute) bool {
	return r != nil && routeVerbHasBody(r.Verb)
}

// RoutePath returns the full route path template, {basepath}{route.path},
// normalized (leading slash, no trailing slash).
func RoutePath(r *MethodRoute) string {
	if r == nil {
		return ""
	}
	basePath := ""
	if r.Method != nil && r.Method.Service != nil && r.Method.Service.Schema != nil {
		basePath = r.Method.Service.Schema.BasePath
	}
	return JoinRoutePath(basePath, r.Path)
}

// RoutePattern returns the ServeMux pattern "VERB /path" of the route. The
// root path gets "/{$}" so it matches only "/", not the subtree.
func RoutePattern(r *MethodRoute) string {
	if r == nil {
		return ""
	}
	path := RoutePath(r)
	if path == "/" {
		path = "/{$}"
	}
	return r.Verb + " " + path
}

// RouteParams returns the request fields bound to the route's path params
// and, for GET/DELETE routes, to its query params, in request field order.
// Body verbs (see RouteHasBody) send all other fields in the body.
func RouteParams(r *MethodRoute) []*RouteParam {
	if r == nil || r.Method == nil {
		return nil
	}
	pathParams, err := validateRoutePathSyntax(r.Path)
	if err != nil {
		return nil
	}
	inPath := map[string]bool{}
	for _, name := range pathParams {
		inPath[name] = true
	}

	var out []*RouteParam
	for _, f := range methodRouteFields(r.Method) {
		if !inPath[f.name] && routeVerbHasBody(r.Verb) {
			continue
		}
		p := &RouteParam{Name: f.name, Optional: f.optional}
		typ := f.typ
		if typ != nil && typ.Type == T_List && typ.List != nil {
			p.Repeated = true
			typ = typ.List.Elem
		}
		if resolved, ok := resolveRouteScalar(typ); ok {
			p.TypeExpr = resolved.String()
		}
		out = append(out, p)
	}
	return out
}
