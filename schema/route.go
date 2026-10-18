package schema

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

// MethodRoute is one REST route of a method, declared in RIDL as a meta line
// `+ VERB /absolute/path` under the method. A method with routes no longer
// serves its default webrpc path (POST {basepath}{Service}/{Method}).
type MethodRoute struct {
	Verb string `json:"verb"` // GET, POST, PUT, PATCH, DELETE or QUERY
	Path string `json:"path"` // absolute path template relative to {basepath}, e.g. /users/{userId}

	Success int `json:"success,omitempty"` // success HTTP status code declared by `+ @success NNN`; 0 means the default 200

	Method *Method `json:"-"` // denormalize/back-reference
}

// RouteVerbs are the HTTP verbs allowed on a method route. QUERY (RFC draft,
// a safe method with a request body) shares the POST/PUT/PATCH binding rules.
var RouteVerbs = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "QUERY"}

func IsRouteVerb(s string) bool {
	return slices.Contains(RouteVerbs, s)
}

// routeVerbHasBody reports whether request fields not consumed by path params
// travel in the JSON body (POST/PUT/PATCH/QUERY) or as query params (GET/DELETE).
func routeVerbHasBody(verb string) bool {
	return verb == "POST" || verb == "PUT" || verb == "PATCH" || verb == "QUERY"
}

var (
	routeParamNameRexp = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*$`)
	routeSegmentRexp   = regexp.MustCompile(`^[a-zA-Z0-9._~-]+$`)
)

// JoinRoutePath joins path parts into a single normalized path: leading
// slash, single slashes between segments, no trailing slash (except "/").
func JoinRoutePath(parts ...string) string {
	segments := []string{}
	for _, part := range parts {
		for _, seg := range strings.Split(part, "/") {
			if seg != "" {
				segments = append(segments, seg)
			}
		}
	}
	return "/" + strings.Join(segments, "/")
}

// routePathSegments splits a normalized path template into its segments.
func routePathSegments(path string) []string {
	return strings.FieldsFunc(path, func(r rune) bool { return r == '/' })
}

// validateRoutePathSyntax checks a path template: it must start with "/", and
// each segment is either a static segment or a whole-segment "{param}".
// Returns the param names in order of appearance.
func validateRoutePathSyntax(path string) ([]string, error) {
	if !strings.HasPrefix(path, "/") {
		return nil, fmt.Errorf("path %q must start with '/'", path)
	}
	if strings.ContainsAny(path, " \t\r\n?#") {
		return nil, fmt.Errorf("path %q must not contain whitespace, query or fragment", path)
	}

	var params []string
	for _, seg := range routePathSegments(path) {
		if seg == "." || seg == ".." {
			return nil, fmt.Errorf("path %q has invalid segment %q: dot segments are not allowed", path, seg)
		}
		if name, ok := strings.CutPrefix(seg, "{"); ok {
			if name, ok = strings.CutSuffix(name, "}"); !ok || !routeParamNameRexp.MatchString(name) {
				return nil, fmt.Errorf("path %q has invalid param segment %q", path, seg)
			}
			params = append(params, name)
		} else if !routeSegmentRexp.MatchString(seg) {
			return nil, fmt.Errorf("path %q has invalid segment %q: params must span a whole segment, e.g. /{userId}", path, seg)
		}
	}
	return params, nil
}

// CheckRoutePathSyntax validates a route path template, so parsers can fail with a source position.
func CheckRoutePathSyntax(path string) error {
	_, err := validateRoutePathSyntax(path)
	return err
}

// routeParamCoreTypes are the types a request field must have to be bound to
// a path param or serialized as a query param.
var routeParamCoreTypes = []CoreType{
	T_Bool,
	T_Uint, T_Uint8, T_Uint16, T_Uint32, T_Uint64,
	T_Int, T_Int8, T_Int16, T_Int32, T_Int64,
	T_Float32, T_Float64,
	T_String, T_Timestamp,
	T_Enum,
}

// resolveRouteScalar resolves aliases and reports whether the type is a
// scalar bindable to a path/query param.
func resolveRouteScalar(vt *VarType) (*VarType, bool) {
	for vt != nil && vt.Type == T_Alias {
		if vt.Alias == nil || vt.Alias.Type == nil {
			return nil, false
		}
		vt = vt.Alias.Type.Type
	}
	if vt == nil {
		return nil, false
	}
	return vt, slices.Contains(routeParamCoreTypes, vt.Type)
}

// routeField is one bindable field of a method's request: a succinct input
// struct field (by its json name), or a method input argument (by its name).
type routeField struct {
	name     string
	typ      *VarType
	optional bool
}

// methodRouteFields lists the request fields a route can bind to. For a
// succinct method these are the input struct's fields; otherwise the method's
// input arguments.
func methodRouteFields(m *Method) []*routeField {
	var fields []*routeField
	if m.Succinct && len(m.Inputs) == 1 && m.Inputs[0].Type != nil && m.Inputs[0].Type.Type == T_Struct {
		for _, f := range m.Inputs[0].Type.Struct.Type.Fields {
			fields = append(fields, &routeField{name: fieldJSONName(f), typ: f.Type, optional: f.Optional})
		}
		return fields
	}
	for _, in := range m.Inputs {
		fields = append(fields, &routeField{name: in.Name, typ: in.Type, optional: in.Optional})
	}
	return fields
}

// validateRoutes enforces the REST route rules across the whole schema. It
// runs after services are parsed, so all back-references and types are set.
func (s *WebRPCSchema) validateRoutes() error {
	type declaredRoute struct {
		id       string // "Service.Method"
		service  string
		verb     string
		path     string // full path template, including basepath
		segments []string
	}
	var routes []declaredRoute
	var defaults []declaredRoute

	if basePath := JoinRoutePath(s.BasePath); basePath != "/" {
		params, err := validateRoutePathSyntax(basePath)
		if err == nil && len(params) > 0 {
			err = fmt.Errorf("path %q must not contain params", basePath)
		}
		if err != nil && s.hasRoutes() {
			return fmt.Errorf("schema error: invalid basepath for REST routes: %w", err)
		}
	}

	for _, svc := range s.Services {
		for _, m := range svc.Methods {
			id := fmt.Sprintf("%s.%s", svc.Name, m.Name)
			if len(m.Routes) == 0 {
				defaultPath := JoinRoutePath(s.BasePath, svc.Name, m.Name)
				defaults = append(defaults, declaredRoute{id: id, service: svc.Name, path: defaultPath, segments: routePathSegments(defaultPath)})
				continue
			}
			for _, route := range m.Routes {
				if err := s.validateMethodRoute(svc, m, route); err != nil {
					return err
				}
				path := RoutePath(route)
				routes = append(routes, declaredRoute{id: id, service: svc.Name, verb: route.Verb, path: path, segments: routePathSegments(path)})
			}
		}
	}

	for i, r := range routes {
		// Unrouted methods' default webrpc paths share the REST mux (http.ServeMux conflict).
		for _, d := range defaults {
			if overlap, _, _ := compareRouteSegments(r.segments, d.segments); overlap {
				return fmt.Errorf("schema error: route '%s %s' of method '%s' overlaps the default webrpc dispatch path '%s' of method '%s'", r.verb, r.path, r.id, d.path, d.id)
			}
		}

		for _, o := range routes[:i] {
			overlap, rMoreStatic, oMoreStatic := compareRouteSegments(r.segments, o.segments)
			if !overlap {
				continue
			}
			if o.service != r.service {
				// Per-service handlers claim every request matching one of their
				// route paths (405 on verb mismatch), so services can't overlap.
				return fmt.Errorf("schema error: route '%s %s' of method '%s' overlaps route '%s %s' of method '%s' in another service: route paths cannot overlap across services", r.verb, r.path, r.id, o.verb, o.path, o.id)
			}
			if o.verb != r.verb {
				continue
			}
			if !rMoreStatic && !oMoreStatic {
				if o.id == r.id {
					return fmt.Errorf("schema error: duplicate route '%s %s' declared twice by method '%s'", r.verb, r.path, r.id)
				}
				return fmt.Errorf("schema error: duplicate route '%s %s' declared by methods '%s' and '%s'", r.verb, r.path, o.id, r.id)
			}
			if rMoreStatic && oMoreStatic {
				// http.ServeMux panics on such pairs.
				return fmt.Errorf("schema error: routes '%s %s' of method '%s' and '%s %s' of method '%s' are ambiguous: a request can match both and neither is more specific", o.verb, o.path, o.id, r.verb, r.path, r.id)
			}
		}
	}

	return nil
}

// compareRouteSegments reports whether two path templates can match the same
// request path, and whether each one has a static segment where the other has
// a param (i.e. is more specific at some position).
func compareRouteSegments(a, b []string) (overlap, aMoreStatic, bMoreStatic bool) {
	if len(a) != len(b) {
		return false, false, false
	}
	for i := range a {
		aParam := strings.HasPrefix(a[i], "{")
		bParam := strings.HasPrefix(b[i], "{")
		switch {
		case !aParam && !bParam:
			if a[i] != b[i] {
				return false, false, false
			}
		case !aParam && bParam:
			aMoreStatic = true
		case aParam && !bParam:
			bMoreStatic = true
		}
	}
	return true, aMoreStatic, bMoreStatic
}

// validateMethodRoute checks one route of a method: verb, path syntax, and
// the field binding rules for path and query params, by the route's own verb.
func (s *WebRPCSchema) validateMethodRoute(svc *Service, m *Method, route *MethodRoute) error {
	methodID := fmt.Sprintf("%s.%s", svc.Name, m.Name)

	if !IsRouteVerb(route.Verb) {
		return fmt.Errorf("schema error: method '%s' declares invalid route verb '%s': must be one of %s", methodID, route.Verb, strings.Join(RouteVerbs, ", "))
	}
	if m.StreamInput || m.StreamOutput {
		return fmt.Errorf("schema error: method '%s' cannot declare a route: stream methods are POST-only", methodID)
	}
	if MethodHasFileUpload(m) || MethodHasFileDownload(m) {
		return fmt.Errorf("schema error: method '%s' cannot declare a route: methods using the file type keep the default webrpc dispatch", methodID)
	}

	params, err := validateRoutePathSyntax(route.Path)
	if err != nil {
		return fmt.Errorf("schema error: invalid route of method '%s': %w", methodID, err)
	}
	if route.Success != 0 && (route.Success < 200 || route.Success > 299) {
		return fmt.Errorf("schema error: invalid @success status code '%d' of route '%s %s' of method '%s': must be a 2xx code (200-299)", route.Success, route.Verb, route.Path, methodID)
	}
	if route.Success == 204 || route.Success == 205 {
		return fmt.Errorf("schema error: invalid @success status code '%d' of route '%s %s' of method '%s': bodyless 204/205 responses are not supported (yet)", route.Success, route.Verb, route.Path, methodID)
	}

	seenParams := map[string]bool{}
	for _, p := range params {
		if seenParams[p] {
			return fmt.Errorf("schema error: route '%s %s' of method '%s' declares duplicate path param '{%s}'", route.Verb, route.Path, methodID, p)
		}
		seenParams[p] = true
	}

	fields := methodRouteFields(m)
	fieldsByName := map[string]*routeField{}
	for _, f := range fields {
		fieldsByName[f.name] = f
	}

	for _, p := range params {
		f, ok := fieldsByName[p]
		if !ok {
			return fmt.Errorf("schema error: path param '{%s}' of method '%s' does not match any request field", p, methodID)
		}
		if f.optional {
			return fmt.Errorf("schema error: path param '{%s}' of method '%s' must bind to a required field, but '%s' is optional", p, methodID, f.name)
		}
		if _, ok := resolveRouteScalar(f.typ); !ok {
			return fmt.Errorf("schema error: path param '{%s}' of method '%s' must bind to a scalar field, but '%s' is of type '%s'", p, methodID, f.name, f.typ)
		}
	}

	// GET/DELETE requests have no body: every field not consumed by a path
	// param is sent as a query param, and must be a scalar or []scalar.
	if !routeVerbHasBody(route.Verb) {
		for _, f := range fields {
			if seenParams[f.name] {
				continue
			}
			if _, ok := resolveRouteScalar(f.typ); ok {
				continue
			}
			if f.typ != nil && f.typ.Type == T_List && f.typ.List != nil {
				if _, ok := resolveRouteScalar(f.typ.List.Elem); ok {
					continue
				}
			}
			return fmt.Errorf("schema error: field '%s' of method '%s' cannot be sent as a query param of '%s %s': only scalars and scalar arrays are supported in %s requests, which have no body", f.name, methodID, route.Verb, route.Path, route.Verb)
		}
	}

	return nil
}

func (s *WebRPCSchema) hasRoutes() bool {
	for _, svc := range s.Services {
		for _, m := range svc.Methods {
			if len(m.Routes) > 0 {
				return true
			}
		}
	}
	return false
}
