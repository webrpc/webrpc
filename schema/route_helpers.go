package schema

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

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

// RouteShape splits a route's full path into its literal and param segments:
// Literals[i] precedes Params[i], and the final literal trails the last param
// ("" when the path ends with one). Generated code matches paths without a
// router by cutting on the literals.
type RouteShape struct {
	Literals []string
	Params   []string
}

// RouteShapeOf splits a route's full path (basepath included) into literals
// and params.
func RouteShapeOf(r *MethodRoute) *RouteShape {
	shape := &RouteShape{}
	if r == nil {
		return shape
	}
	parts := strings.Split(strings.TrimPrefix(RoutePath(r), "/"), "/")
	lit := "/"
	for i, seg := range parts {
		if name, ok := strings.CutPrefix(seg, "{"); ok {
			if name, ok = strings.CutSuffix(name, "}"); ok {
				shape.Literals = append(shape.Literals, lit)
				shape.Params = append(shape.Params, name)
				lit = ""
				continue
			}
		}
		if lit == "" {
			lit = "/"
		}
		lit += seg
		if i < len(parts)-1 {
			lit += "/"
		}
	}
	shape.Literals = append(shape.Literals, lit)
	return shape
}

// SegN is the path param count, the N of the restSegN matcher.
func (s *RouteShape) SegN() int { return len(s.Params) }

// LiteralSegs counts the route's literal (non-param) path segments. More
// literal segments make a param route more specific, like http.ServeMux
// precedence: GET /a/{x}/b beats GET /a/{x}/{y} for /a/v/b.
func (s *RouteShape) LiteralSegs() int {
	n := 0
	for _, lit := range s.Literals {
		if segs := strings.Trim(lit, "/"); segs != "" {
			n += strings.Count(segs, "/") + 1
		}
	}
	return n
}

// Blanks joins the matcher's blank return variables, e.g. "_, _".
func (s *RouteShape) Blanks() string {
	return strings.TrimSuffix(strings.Repeat("_, ", len(s.Params)), ", ")
}

// Args renders the quoted literal arguments of a restSegN call: all literals
// but the trailing one.
func (s *RouteShape) Args() string {
	parts := make([]string, len(s.Params))
	for i := range parts {
		parts[i] = strconv.Quote(s.Literals[i])
	}
	return strings.Join(parts, ", ")
}

// Trailing is the literal after the last param, "" when the path ends with one.
func (s *RouteShape) Trailing() string { return s.Literals[len(s.Params)] }

// HasTrailing reports whether a literal follows the last path param; such
// routes match with strings.CutSuffix before the restSegN call.
func (s *RouteShape) HasTrailing() bool { return s.Trailing() != "" }

// RouteHandlerField returns the generated Go field name of a route's handler:
// "rest<Method>" for the first route, and additional routes suffixed with
// "RPC" when they revive the method's default RPC path, else the verb — with
// a 2-based ordinal when several alias routes share a verb, so names stay
// unique: GET /a, GET /b, GET /c give restX, restXGet, restXGet2.
func RouteHandlerField(r *MethodRoute) string {
	m := r.Method
	if m == nil {
		return ""
	}
	name := "rest" + m.Name
	idx := slices.Index(m.Routes, r)
	if idx == 0 {
		return name
	}
	if RoutePath(r) == MethodRPCPath(m) {
		return name + "RPC"
	}
	suffix := strings.ToUpper(r.Verb[:1]) + strings.ToLower(r.Verb[1:])
	n := 0
	for _, prev := range m.Routes[1:idx] {
		if prev.Verb == r.Verb && RoutePath(prev) != MethodRPCPath(m) {
			n++
		}
	}
	if n > 0 {
		suffix += strconv.Itoa(n + 1)
	}
	return name + suffix
}

// restVarReserved are matcher-local names a path param must not take: Go
// keywords are handled by the target's KeywordMap, these are the locals the
// generated matchers themselves use.
var restVarReserved = map[string]bool{"ok": true, "err": true, "s": true, "path": true}

// RestVarName sanitizes a path param name for use as a generated Go local:
// keywords and matcher-local names get a trailing underscore, mirroring the
// generator's KeywordMap convention. The wire name is unaffected.
func RestVarName(name string, keywords map[string]interface{}) string {
	if v, ok := keywords[name]; (ok && truthy(v)) || restVarReserved[name] {
		return name + "_"
	}
	return name
}

// truthy matches template truthiness for KeywordMap values (all set to true).
func truthy(v interface{}) bool {
	b, ok := v.(bool)
	return ok && b
}

// MethodRPCPath is a method's default webrpc path: {basepath}{Service}/{Method}.
func MethodRPCPath(m *Method) string {
	svc := m.Service
	if svc == nil {
		return ""
	}
	base := ""
	if svc.Schema != nil {
		base = svc.Schema.BasePath
	}
	return base + svc.Name + "/" + m.Name
}

// RestRoute is one entry of a generated REST router's dispatch tables.
type RestRoute struct {
	Verb    string
	Path    string // literal request path
	Pattern string // r.Pattern value; empty for plain RPC methods
	Handler string // generated Go handler, e.g. "s.restGetItem" or "s.handleRpcMethod"
	Method  *Method
	Route   *MethodRoute // nil for a plain RPC method
}

// RestVerbCase groups one verb's static and param routes for ServeHTTP.
type RestVerbCase struct {
	Verb    string // the generated case label; "GET" also matches HEAD
	Statics []*RestRoute
	Params  []*RestRoute
}

// RestAllowCase is one static Allow table entry; Paths share one Allow value.
type RestAllowCase struct {
	Paths []string
	Allow string
}

// RestShapeCase is one param matcher entry of the Allow table.
type RestShapeCase struct {
	Shape *RouteShape
	Allow string
}

// RestRouterPlan is the computed dispatch plan for one service's generated REST
// router: ServeHTTP's verb cases, the static Allow table, and the param
// matcher entries. RPC lists the service's methods without routes, which keep
// their default POST path.
type RestRouterPlan struct {
	Verbs  []*RestVerbCase
	Allow  []*RestAllowCase
	Shapes []*RestShapeCase
	RPC    []*RestRoute
}

// restAllowVerbOrder scans static routes for the Allow table, safe verbs
// first, so the generated entries read GET/QUERY before the body verbs.
var restAllowVerbOrder = []string{"GET", "QUERY", "POST", "PUT", "PATCH", "DELETE"}

// RestRouter computes the dispatch plan for a service's REST routes, or nil
// when the service has none.
func RestRouter(svc *Service) *RestRouterPlan {
	hasRoutes := false
	for _, m := range svc.Methods {
		if MethodHasRoute(m) {
			hasRoutes = true
			break
		}
	}
	if !hasRoutes {
		return nil
	}

	verbRoutes := map[string][]*RestRoute{}
	allowVerbsByPath := map[string]map[string]bool{}
	rpcPathSet := map[string]bool{}
	for _, m := range svc.Methods {
		rpcPathSet[MethodRPCPath(m)] = true
	}

	addRoute := func(r *RestRoute, shape *RouteShape) {
		verbRoutes[r.Verb] = append(verbRoutes[r.Verb], r)
		if len(shape.Params) == 0 {
			if allowVerbsByPath[r.Path] == nil {
				allowVerbsByPath[r.Path] = map[string]bool{}
			}
			allowVerbsByPath[r.Path][r.Verb] = true
		}
	}

	for _, m := range svc.Methods {
		if len(m.Routes) == 0 {
			addRoute(&RestRoute{Verb: "POST", Path: MethodRPCPath(m), Handler: "s.handle" + m.Name, Method: m}, &RouteShape{})
			continue
		}
		for _, route := range m.Routes {
			addRoute(&RestRoute{
				Verb:    route.Verb,
				Path:    RoutePath(route),
				Pattern: RoutePattern(route),
				Handler: "s." + RouteHandlerField(route),
				Method:  m,
				Route:   route,
			}, RouteShapeOf(route))
		}
	}

	router := &RestRouterPlan{}
	for _, verb := range RouteVerbs {
		routes := verbRoutes[verb]
		if len(routes) == 0 {
			continue
		}
		verbCase := &RestVerbCase{Verb: verb}
		for _, r := range routes {
			if len(RouteShapeOf(r.Route).Params) > 0 {
				verbCase.Params = append(verbCase.Params, r)
			} else {
				verbCase.Statics = append(verbCase.Statics, r)
			}
		}
		// Param matchers run most specific first, like http.ServeMux
		// precedence: more literal segments, then a longer literal prefix, then
		// fewer path params.
		slices.SortStableFunc(verbCase.Params, func(a, b *RestRoute) int {
			sa, sb := RouteShapeOf(a.Route), RouteShapeOf(b.Route)
			if d := sb.LiteralSegs() - sa.LiteralSegs(); d != 0 {
				return d
			}
			if d := len(sb.Literals[0]) - len(sa.Literals[0]); d != 0 {
				return d
			}
			if d := len(sa.Params) - len(sb.Params); d != 0 {
				return d
			}
			return strings.Compare(a.Path, b.Path)
		})
		router.Verbs = append(router.Verbs, verbCase)
	}

	// Static Allow entries: one per literal route path, then the RPC paths
	// (route aliases that revive a method's default path, and plain RPC
	// methods) grouped into one trailing entry.
	seen := map[string]bool{}
	for _, verb := range restAllowVerbOrder {
		for _, r := range verbRoutes[verb] {
			if len(RouteShapeOf(r.Route).Params) > 0 || seen[r.Path] || rpcPathSet[r.Path] {
				continue
			}
			seen[r.Path] = true
			router.Allow = append(router.Allow, &RestAllowCase{
				Paths: []string{r.Path},
				Allow: restAllowValue(allowVerbsByPath[r.Path]),
			})
		}
	}
	var rpcPaths []string
	for _, m := range svc.Methods {
		if len(m.Routes) == 0 {
			if !slices.Contains(rpcPaths, MethodRPCPath(m)) {
				rpcPaths = append(rpcPaths, MethodRPCPath(m))
			}
			continue
		}
		for _, route := range m.Routes {
			if p := RoutePath(route); rpcPathSet[p] && !slices.Contains(rpcPaths, p) {
				rpcPaths = append(rpcPaths, p)
			}
		}
	}
	if len(rpcPaths) > 0 {
		verbs := map[string]bool{}
		for _, p := range rpcPaths {
			for verb := range allowVerbsByPath[p] {
				verbs[verb] = true
			}
		}
		router.Allow = append(router.Allow, &RestAllowCase{Paths: rpcPaths, Allow: restAllowValue(verbs)})
	}

	// Param matcher Allow entries: most specific first, verbs unioned per shape.
	shapeByKey := map[string]*RouteShape{}
	var shapes []*RouteShape
	shapeVerbs := map[string]map[string]bool{}
	for _, verb := range RouteVerbs {
		for _, r := range verbRoutes[verb] {
			shape := RouteShapeOf(r.Route)
			if len(shape.Params) == 0 {
				continue
			}
			key := strings.Join(shape.Literals, "\x00")
			if shapeByKey[key] == nil {
				shapeByKey[key] = shape
				shapes = append(shapes, shape)
				shapeVerbs[key] = map[string]bool{}
			}
			shapeVerbs[key][verb] = true
		}
	}
	slices.SortStableFunc(shapes, func(a, b *RouteShape) int {
		if d := b.LiteralSegs() - a.LiteralSegs(); d != 0 {
			return d
		}
		if d := len(b.Literals[0]) - len(a.Literals[0]); d != 0 {
			return d
		}
		if d := len(a.Params) - len(b.Params); d != 0 {
			return d
		}
		return strings.Compare(a.Literals[0], b.Literals[0])
	})
	for _, shape := range shapes {
		router.Shapes = append(router.Shapes, &RestShapeCase{
			Shape: shape,
			Allow: restAllowValue(shapeVerbs[strings.Join(shape.Literals, "\x00")]),
		})
	}
	for _, m := range svc.Methods {
		if len(m.Routes) == 0 {
			router.RPC = append(router.RPC, &RestRoute{Path: MethodRPCPath(m), Handler: "handle" + m.Name, Method: m})
		}
	}
	return router
}

// restAllowValue joins a path's verbs into an Allow header value, sorted, with
// HEAD next to GET (the dispatcher serves HEAD on GET routes).
func restAllowValue(verbs map[string]bool) string {
	allow := make([]string, 0, len(verbs)+1)
	for verb := range verbs {
		allow = append(allow, verb)
		if verb == "GET" {
			allow = append(allow, "HEAD")
		}
	}
	slices.Sort(allow)
	return strings.Join(allow, ", ")
}

// QuotedList renders Go string literals: ("/a", "/b") -> `"/a", "/b"`.
func QuotedList(paths []string) string {
	parts := make([]string, len(paths))
	for i, p := range paths {
		parts[i] = strconv.Quote(p)
	}
	return strings.Join(parts, ", ")
}

// RPCPathCondition renders the restMiss guard that keeps the legacy "only
// POST is allowed" message on plain RPC paths, evaluated on the canonical
// escaped path: (1 path) -> `path == "/rpc/Svc/M"`.
func RPCPathCondition(routes []*RestRoute) string {
	parts := make([]string, len(routes))
	for i, r := range routes {
		parts[i] = fmt.Sprintf("path == %s", strconv.Quote(r.Path))
	}
	return strings.Join(parts, " || ")
}

// RestSegHelpers renders the Go source of restSeg3..restSegN, the segment
// matchers for routes with 3 or more path params (restSeg1 and restSeg2 are
// template-emitted). Returns "" when no route needs them.
func RestSegHelpers(maxParams int) string {
	if maxParams < 3 {
		return ""
	}
	var b strings.Builder
	for k := 3; k <= maxParams; k++ {
		args := []string{"path", "prefix"}
		for i := 1; i <= k-1; i++ {
			args = append(args, fmt.Sprintf("middle%d", i))
		}
		zeros := ""
		for i := 0; i < k; i++ {
			if i > 0 {
				zeros += ", "
			}
			zeros += `""`
		}
		var rets []string
		for i := 0; i < k; i++ {
			rets = append(rets, "string")
		}
		fmt.Fprintf(&b, "\n// restSeg%d matches a literal prefix, %d middle literal(s), and %d path parameter segments.\n", k, k-1, k)
		fmt.Fprintf(&b, "func restSeg%d(%s string) (%s, bool) {\n", k, strings.Join(args, ", "), strings.Join(rets, ", "))
		fmt.Fprintf(&b, "\ts, ok := strings.CutPrefix(path, prefix)\n")
		fmt.Fprintf(&b, "\tif !ok {\n\t\treturn %s, false\n\t}\n", zeros)
		var segs []string
		for i := 1; i <= k-1; i++ {
			seg := fmt.Sprintf("seg%d", i)
			segs = append(segs, seg)
			// An empty middle literal means adjacent params; the boundary is
			// then the segment separator itself.
			fmt.Fprintf(&b, "\tsep%d := middle%d\n", i, i)
			fmt.Fprintf(&b, "\tif sep%d == \"\" {\n\t\tsep%d = \"/\"\n\t}\n", i, i)
			fmt.Fprintf(&b, "\t%s, s, ok := strings.Cut(s, sep%d)\n", seg, i)
			fmt.Fprintf(&b, "\tif !ok || %s == \"\" || strings.Contains(%s, \"/\") {\n\t\treturn %s, false\n\t}\n", seg, seg, zeros)
			decl := "="
			if i == 1 {
				decl = ":="
			}
			fmt.Fprintf(&b, "\t%s, err %s url.PathUnescape(%s)\n", seg, decl, seg)
			fmt.Fprintf(&b, "\tif err != nil {\n\t\treturn %s, false\n\t}\n", zeros)
		}
		fmt.Fprintf(&b, "\tif s == \"\" || strings.Contains(s, \"/\") {\n\t\treturn %s, false\n\t}\n", zeros)
		fmt.Fprintf(&b, "\ts, err = url.PathUnescape(s)\n")
		fmt.Fprintf(&b, "\tif err != nil {\n\t\treturn %s, false\n\t}\n", zeros)
		fmt.Fprintf(&b, "\treturn %s, s, true\n}\n", strings.Join(segs, ", "))
	}
	return b.String()
}
