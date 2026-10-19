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

// RouteHasBodyFields reports whether the route has request fields that travel
// in the JSON body, i.e. request fields not bound to path params.
func RouteHasBodyFields(r *MethodRoute) bool {
	if r == nil || r.Method == nil || !RouteHasBody(r) {
		return false
	}
	pathParams, err := validateRoutePathSyntax(r.Path)
	if err != nil {
		return false
	}
	inPath := map[string]bool{}
	for _, name := range pathParams {
		inPath[name] = true
	}
	for _, f := range methodRouteFields(r.Method) {
		if !inPath[f.name] {
			return true
		}
	}
	return false
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

// RouteSuccessStatus returns the route's default success status code: the
// first declared `+ @success` code, or 200 when the route has no `+ @success`
// line.
func RouteSuccessStatus(r *MethodRoute) int {
	codes := RouteSuccessCodes(r)
	return codes[0]
}

// RouteSuccessCodes returns the route's full declared set of success status
// codes, or [200] when the route has no `+ @success` line.
func RouteSuccessCodes(r *MethodRoute) []int {
	if r == nil || len(r.Success) == 0 {
		return []int{200}
	}
	return r.Success
}

// MethodHasEmptyResponse reports whether the method's success response carries
// no body: no outputs at all, or a single empty struct output (`=> ()`).
func MethodHasEmptyResponse(m *Method) bool {
	if len(m.Outputs) == 0 {
		return true
	}
	if len(m.Outputs) == 1 && m.Outputs[0].Type != nil && m.Outputs[0].Type.Type == T_Struct {
		if st := m.Outputs[0].Type.Struct; st != nil && st.Type != nil && len(st.Type.Fields) == 0 {
			return true
		}
	}
	return false
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
	RPC    []string // default webrpc paths of methods without routes
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
		// precedence: more literal segments first. Ties keep declaration order;
		// validation rejects the overlapping ties where the winner matters.
		slices.SortStableFunc(verbCase.Params, func(a, b *RestRoute) int {
			return RouteShapeOf(b.Route).LiteralSegs() - RouteShapeOf(a.Route).LiteralSegs()
		})
		router.Verbs = append(router.Verbs, verbCase)
	}

	// Param shapes: one entry per distinct literal structure, carrying the
	// union of its own routes' verbs.
	type paramShape struct {
		shape *RouteShape
		segs  []string
		verbs map[string]bool
	}
	var paramShapes []paramShape
	shapeIndex := map[string]int{}
	for _, verb := range RouteVerbs {
		for _, r := range verbRoutes[verb] {
			shape := RouteShapeOf(r.Route)
			if len(shape.Params) == 0 {
				continue
			}
			key := strings.Join(shape.Literals, "\x00")
			i, ok := shapeIndex[key]
			if !ok {
				i = len(paramShapes)
				shapeIndex[key] = i
				paramShapes = append(paramShapes, paramShape{
					shape: shape,
					segs:  routePathSegments(RoutePath(r.Route)),
					verbs: map[string]bool{},
				})
			}
			paramShapes[i].verbs[verb] = true
		}
	}
	// Overlapping shapes serve some concrete path together (e.g. GET /{x}/b
	// and DELETE /a/{y} both match /a/b), and the generated Allow table
	// returns the first match, so each shape unions the other's verbs.
	for i := range paramShapes {
		for j := i + 1; j < len(paramShapes); j++ {
			a, b := &paramShapes[i], &paramShapes[j]
			if overlap, _, _ := compareRouteSegments(a.segs, b.segs); overlap {
				for verb := range a.verbs {
					b.verbs[verb] = true
				}
				for verb := range b.verbs {
					a.verbs[verb] = true
				}
			}
		}
	}
	// staticVerbs unions a static path's own verbs with every param shape
	// that also matches it: such a request really dispatches on the shape
	// too, so the Allow value must list those methods.
	staticVerbs := func(path string) map[string]bool {
		verbs := map[string]bool{}
		for verb := range allowVerbsByPath[path] {
			verbs[verb] = true
		}
		segs := routePathSegments(path)
		for _, ps := range paramShapes {
			if overlap, _, _ := compareRouteSegments(ps.segs, segs); overlap {
				for verb := range ps.verbs {
					verbs[verb] = true
				}
			}
		}
		return verbs
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
				Allow: restAllowValue(staticVerbs(r.Path)),
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
			for verb := range staticVerbs(p) {
				verbs[verb] = true
			}
		}
		router.Allow = append(router.Allow, &RestAllowCase{Paths: rpcPaths, Allow: restAllowValue(verbs)})
	}
	for _, ps := range paramShapes {
		router.Shapes = append(router.Shapes, &RestShapeCase{
			Shape: ps.shape,
			Allow: restAllowValue(ps.verbs),
		})
	}
	for _, m := range svc.Methods {
		if len(m.Routes) == 0 {
			router.RPC = append(router.RPC, MethodRPCPath(m))
		}
	}
	return router
}

// RestFeatures reports which REST helpers the generated Go code calls, so the
// templates emit only those. Server helpers cover every route of every
// service; client helpers only each method's canonical route.
type RestFeatures struct {
	Segs       int             // highest path-param count; 0 when no route has path params
	PathParams bool            // restPath binder
	Query      bool            // restQuery: required scalar query param
	QueryOpt   bool            // restQueryOpt: optional scalar query param
	QueryMaybe bool            // restQueryMaybe: optional query param, non-pointer field
	QueryList  bool            // restQueryList: []scalar query param
	Parsers    map[string]bool // string, bool, uint, int, float, text, json, jsonString

	ClientPath  bool            // restURL.setPath: canonical route has path params
	ClientQuery bool            // restURL.addQuery: canonical route has query params
	ClientKinds map[string]bool // string, bool, uint, int, float, text, custom, customString
}

// RestFeaturesOf scans the services' routes for the helpers they call.
func RestFeaturesOf(services []*Service) *RestFeatures {
	f := &RestFeatures{Parsers: map[string]bool{}, ClientKinds: map[string]bool{}}
	for _, svc := range services {
		for _, m := range svc.Methods {
			customs := routeFieldCustomTypes(m)
			for _, route := range m.Routes {
				inPath := map[string]bool{}
				if names, err := validateRoutePathSyntax(route.Path); err == nil {
					for _, name := range names {
						inPath[name] = true
					}
				}
				if n := RouteShapeOf(route).SegN(); n > f.Segs {
					f.Segs = n
				}
				for _, p := range RouteParams(route) {
					custom := customs[p.Name]
					if inPath[p.Name] {
						f.PathParams = true
					} else {
						switch {
						case p.Repeated:
							f.QueryList = true
						case p.Optional && custom != "":
							f.QueryMaybe = true
						case p.Optional:
							f.QueryOpt = true
						default:
							f.Query = true
						}
					}
					f.Parsers[restParamParser(p, custom)] = true
				}
			}
			if route := MethodCanonicalRoute(m); route != nil {
				canonPathParams, _ := validateRoutePathSyntax(route.Path)
				for _, p := range RouteParams(route) {
					custom := customs[p.Name]
					if slices.Contains(canonPathParams, p.Name) {
						f.ClientPath = true
					} else {
						f.ClientQuery = true
					}
					switch kind := restParamParser(p, custom); kind {
					case "json":
						f.ClientKinds["custom"] = true
					case "jsonString":
						f.ClientKinds["customString"] = true
					default:
						f.ClientKinds[kind] = true
					}
				}
			}
		}
	}
	// restJSONString wraps restJSON, so the base parser must be emitted too.
	if f.Parsers["jsonString"] {
		f.Parsers["json"] = true
	}
	return f
}

// routeFieldCustomTypes maps each bindable request field's wire name to its
// go.field.type override, mirroring the generator's param binding.
func routeFieldCustomTypes(m *Method) map[string]string {
	out := map[string]string{}
	custom := func(metas []TypeFieldMeta) string {
		v := ""
		for _, meta := range metas {
			if s, ok := meta["go.field.type"].(string); ok {
				v = s
			}
		}
		return v
	}
	if m.Succinct && len(m.Inputs) == 1 && m.Inputs[0].Type != nil && m.Inputs[0].Type.Type == T_Struct {
		for _, fld := range m.Inputs[0].Type.Struct.Type.Fields {
			out[fieldJSONName(fld)] = custom(fld.Meta)
		}
		return out
	}
	for _, in := range m.Inputs {
		out[in.Name] = custom(in.Meta)
	}
	return out
}

// restParamParser names the parser a param's type binds with, mirroring the
// generator's binder templates: a go.field.type override parses as JSON, the
// core scalars with their strconv parser, everything else as text.
func restParamParser(p *RouteParam, custom string) string {
	switch {
	case custom != "":
		if p.TypeExpr == "string" {
			return "jsonString"
		}
		return "json"
	case p.TypeExpr == "string":
		return "string"
	case p.TypeExpr == "bool":
		return "bool"
	case strings.HasPrefix(p.TypeExpr, "uint"):
		return "uint"
	case strings.HasPrefix(p.TypeExpr, "int"):
		return "int"
	case strings.HasPrefix(p.TypeExpr, "float"):
		return "float"
	default:
		return "text"
	}
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
func RPCPathCondition(paths []string) string {
	parts := make([]string, len(paths))
	for i, path := range paths {
		parts[i] = fmt.Sprintf("path == %s", strconv.Quote(path))
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
