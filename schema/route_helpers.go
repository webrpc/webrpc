package schema

// RouteParam is a request field bound to a path param (when the route path
// template contains "{Name}") or else to a query param (GET/DELETE/QUERY).
//
// This type is used as a template helper for adding REST route support in
// various codegen targets.
type RouteParam struct {
	Name     string // wire name: json field name (succinct) or argument name
	TypeExpr string // alias-resolved scalar type, the element type for lists
	Optional bool
	Repeated bool // []scalar query param, sent as repeated keys
}

// MethodHasRoute returns true if the method declares a REST route, served
// instead of its default webrpc path (POST {basepath}{Service}/{Method}).
func MethodHasRoute(m *Method) bool {
	return m != nil && m.Route != nil
}

// MethodRouteVerb returns the route's HTTP verb, or "" without a route.
func MethodRouteVerb(m *Method) string {
	if !MethodHasRoute(m) {
		return ""
	}
	return m.Route.Verb
}

// MethodRoutePath returns the full route path template, composed as
// {basepath}{service.path}{method.path} and normalized (leading slash, no
// trailing slash).
func MethodRoutePath(m *Method) string {
	if !MethodHasRoute(m) {
		return ""
	}
	basePath := ""
	servicePath := ""
	if m.Service != nil {
		servicePath = m.Service.Path
		if m.Service.Schema != nil {
			basePath = m.Service.Schema.BasePath
		}
	}
	return JoinRoutePath(basePath, servicePath, m.Route.Path)
}

// MethodRouteParams returns the request fields bound to the route's path
// params and, for GET/DELETE/QUERY routes, to its query params — in request
// field order. Body verbs send all other fields in the JSON body.
func MethodRouteParams(m *Method) []*RouteParam {
	if !MethodHasRoute(m) {
		return nil
	}
	pathParams, err := validateRoutePathSyntax(MethodRoutePath(m))
	if err != nil {
		return nil
	}
	inPath := map[string]bool{}
	for _, name := range pathParams {
		inPath[name] = true
	}

	var out []*RouteParam
	for _, f := range methodRouteFields(m) {
		if !inPath[f.name] && routeVerbHasBody(m.Route.Verb) {
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
