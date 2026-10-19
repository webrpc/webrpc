package ridl

import (
	"io"
	"os"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/webrpc/webrpc/schema"
)

// captureStderr collects everything written to os.Stderr while fn runs.
func captureStderr(t *testing.T, fn func()) string {
	t.Helper()

	r, w, err := os.Pipe()
	require.NoError(t, err)
	orig := os.Stderr
	os.Stderr = w
	defer func() { os.Stderr = orig }()

	fn()

	require.NoError(t, w.Close())
	out, err := io.ReadAll(r)
	require.NoError(t, err)
	return string(out)
}

func parseStringWithWarnings(src string) (*schema.WebRPCSchema, []string, error) {
	fsys := fstest.MapFS{
		"main.ridl": {Data: []byte(src)},
	}
	p := NewParser(fsys, "/", "main.ridl")
	s, err := p.Parse()
	return s, p.Warnings(), err
}

const annotationsRIDLHeader = `webrpc = v1
name = example
version = v0.0.1

struct PingRequest
  - x: uint64

struct PingResponse
  - status: bool

`

func TestRIDLMethodAnnotationsNewSyntax(t *testing.T) {
	s, warnings, err := parseStringWithWarnings(annotationsRIDLHeader + `
service Example
  - Ping(PingRequest) => (PingResponse)
    + @internal
    + @auth:admin
    + @auth2:"cookies,authorization"
    + @rateLimit:100
    + GET /ping
`)
	require.NoError(t, err)
	assert.Empty(t, warnings)

	ping := s.Services[0].Methods[0]
	require.Contains(t, ping.Annotations, "internal")
	require.Contains(t, ping.Annotations, "auth")
	require.Contains(t, ping.Annotations, "auth2")
	require.Contains(t, ping.Annotations, "rateLimit")
	assert.Equal(t, "admin", ping.Annotations["auth"].Value)
	assert.Equal(t, "cookies,authorization", ping.Annotations["auth2"].Value)
	assert.Equal(t, "100", ping.Annotations["rateLimit"].Value)
	assert.Empty(t, ping.Annotations["internal"].Value)

	require.Len(t, ping.Routes, 1)
	assert.Equal(t, "GET", ping.Routes[0].Verb)
	assert.Empty(t, ping.Routes[0].Success)
}

func TestRIDLMethodAnnotationsBareValueToEOL(t *testing.T) {
	s, warnings, err := parseStringWithWarnings(annotationsRIDLHeader + `
service Example
  - Ping(PingRequest) => (PingResponse)
    + @deprecated:Use GetUserV2 instead.
    + @headers:X-Access-Key,S2S,Cookies
    + @public
`)
	require.NoError(t, err)
	assert.Empty(t, warnings)

	ping := s.Services[0].Methods[0]
	assert.Equal(t, "Use GetUserV2 instead.", ping.Annotations["deprecated"].Value)
	assert.Equal(t, "X-Access-Key,S2S,Cookies", ping.Annotations["headers"].Value)
	assert.Empty(t, ping.Annotations["public"].Value)
}

func TestRIDLMethodAnnotationsQuotedValueEscapeHatch(t *testing.T) {
	s, warnings, err := parseStringWithWarnings(annotationsRIDLHeader + `
service Example
  - Ping(PingRequest) => (PingResponse)
    + @padded:" leading and trailing "
    + @hashTag:"value with # comment look"
`)
	require.NoError(t, err)
	assert.Empty(t, warnings)

	ping := s.Services[0].Methods[0]
	assert.Equal(t, " leading and trailing ", ping.Annotations["padded"].Value)
	assert.Equal(t, "value with # comment look", ping.Annotations["hashTag"].Value)
}

func TestRIDLMethodAnnotationsMixedOldAndNewSyntax(t *testing.T) {
	s, warnings, err := parseStringWithWarnings(annotationsRIDLHeader + `
service Example
  @auth:admin
  - Ping(PingRequest) => (PingResponse)
    + @internal
    + GET /ping
`)
	require.NoError(t, err)
	require.Len(t, warnings, 1)
	assert.Contains(t, warnings[0], "warning: deprecated annotation syntax")
	assert.Contains(t, warnings[0], "main.ridl:13")

	ping := s.Services[0].Methods[0]
	require.Contains(t, ping.Annotations, "auth")
	require.Contains(t, ping.Annotations, "internal")
	assert.Equal(t, "admin", ping.Annotations["auth"].Value)
}

func TestRIDLMethodAnnotationsOldSyntaxStillParsesWithWarning(t *testing.T) {
	s, warnings, err := parseStringWithWarnings(annotationsRIDLHeader + `
service Example
  @internal
  @deprecated:"Use PingV2 instead."
  - Ping(PingRequest) => (PingResponse)
`)
	require.NoError(t, err)
	require.Len(t, warnings, 2)
	assert.Contains(t, warnings[0], "main.ridl:13: warning: deprecated annotation syntax")
	assert.Contains(t, warnings[1], "main.ridl:14: warning: deprecated annotation syntax")

	ping := s.Services[0].Methods[0]
	require.Contains(t, ping.Annotations, "internal")
	require.Contains(t, ping.Annotations, "deprecated")
}

func TestRIDLMethodAnnotationsCommentPlacement(t *testing.T) {
	s, _, err := parseStringWithWarnings(annotationsRIDLHeader + `
service Example
  # ping does ping things
  - Ping(PingRequest) => (PingResponse)
    + @internal
  - Status(PingRequest) => (PingResponse)
`)
	require.NoError(t, err)

	assert.Equal(t, []string{"ping does ping things"}, s.Services[0].Methods[0].Comments)
	assert.Empty(t, s.Services[0].Methods[1].Comments)
}

func TestRIDLMethodAnnotationWithoutMethodError(t *testing.T) {
	_, _, err := parseStringWithWarnings(annotationsRIDLHeader + `
service Example
  + @internal
  - Ping(PingRequest) => (PingResponse)
`)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "annotation must be declared under a service method")
}

func TestRIDLMethodAnnotationDuplicateNewSyntaxError(t *testing.T) {
	_, _, err := parseStringWithWarnings(annotationsRIDLHeader + `
service Example
  - Ping(PingRequest) => (PingResponse)
    + @internal
    + @internal
`)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "duplicate annotation type: internal")
}

func TestRIDLMethodAnnotationDuplicateMixedFormsError(t *testing.T) {
	_, _, err := parseStringWithWarnings(annotationsRIDLHeader + `
service Example
  @internal
  - Ping(PingRequest) => (PingResponse)
    + @internal
`)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "duplicate annotation type: internal")
}

func TestRIDLMethodAnnotationsOnePerLineError(t *testing.T) {
	_, _, err := parseStringWithWarnings(annotationsRIDLHeader + `
service Example
  - Ping(PingRequest) => (PingResponse)
    + @public @cache:60
`)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected '@': one annotation per line")
}

func TestRIDLMethodAnnotationsEmptyValueError(t *testing.T) {
	_, _, err := parseStringWithWarnings(annotationsRIDLHeader + `
service Example
  - Ping(PingRequest) => (PingResponse)
    + @internal:
`)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "expecting value after ':' in annotation '@internal'")
}

func TestRIDLMethodAnnotationTrailingComment(t *testing.T) {
	s, _, err := parseStringWithWarnings(annotationsRIDLHeader + `
service Example
  - Ping(PingRequest) => (PingResponse)
    + @auth:admin # only admins
    + @see:"https://a.com/#frag"
    + @label:hello é and more
`)
	require.NoError(t, err)

	ping := s.Services[0].Methods[0]
	assert.Equal(t, "admin", ping.Annotations["auth"].Value)
	assert.Equal(t, "https://a.com/#frag", ping.Annotations["see"].Value)
	assert.Equal(t, "hello é and more", ping.Annotations["label"].Value)
}

func TestRIDLDeprecationWarningQuotesUnrepresentableValues(t *testing.T) {
	_, warnings, err := parseStringWithWarnings(annotationsRIDLHeader + `
service Example
  @see:"https://a.com/#frag"
  @plain:simple
  - Ping(PingRequest) => (PingResponse)
`)
	require.NoError(t, err)
	require.Len(t, warnings, 2)
	assert.Contains(t, warnings[0], `+ @see:"https://a.com/#frag"`)
	assert.Contains(t, warnings[1], "+ @plain:simple")
}

func TestRIDLSuccessBeforeBetweenAfterRoutes(t *testing.T) {
	for _, tc := range []struct {
		name string
		ridl string
	}{
		{
			name: "before",
			ridl: `
  - Ping(PingRequest) => (PingResponse)
    + @success 200|201
    + GET /ping
    + POST /pings
`,
		},
		{
			name: "between",
			ridl: `
  - Ping(PingRequest) => (PingResponse)
    + GET /ping
    + @success 200|201
    + POST /pings
`,
		},
		{
			name: "after",
			ridl: `
  - Ping(PingRequest) => (PingResponse)
    + GET /ping
    + POST /pings
    + @success 200|201
`,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			s, warnings, err := parseStringWithWarnings(annotationsRIDLHeader + `
service Example
` + tc.ridl)
			require.NoError(t, err)
			assert.Empty(t, warnings)

			ping := s.Services[0].Methods[0]
			require.Len(t, ping.Routes, 2)
			assert.Equal(t, []int{200, 201}, ping.Routes[0].Success)
			assert.Equal(t, []int{200, 201}, ping.Routes[1].Success)
		})
	}
}

func TestRIDLSuccessWithoutRoutesError(t *testing.T) {
	_, _, err := parseStringWithWarnings(annotationsRIDLHeader + `
service Example
  - Ping(PingRequest) => (PingResponse)
    + @success 200
`)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "@success requires at least one route on method 'Ping'")
}

func TestRIDLSuccessDuplicateMethodLevelError(t *testing.T) {
	_, _, err := parseStringWithWarnings(annotationsRIDLHeader + `
service Example
  - Ping(PingRequest) => (PingResponse)
    + GET /ping
    + @success 200
    + @success 201
`)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "method 'Ping' already declares @success")
}

// An imported file's warnings must reach both Warnings() and stderr, and stderr
// exactly once.
func TestRIDLImportedFileWarnings(t *testing.T) {
	fsys := fstest.MapFS{
		"main.ridl": {Data: []byte(`webrpc = v1
name = example
version = v0.0.1

import
- types.ridl

service Example
  - Ping() => (ok: bool)
`)},
		"types.ridl": {Data: []byte(`webrpc = v1
name = types
version = v0.0.1

service TypesService
  @internal
  - Ping() => (ok: bool)
`)},
	}

	var (
		warnings []string
		parseErr error
	)
	stderr := captureStderr(t, func() {
		p := NewParser(fsys, "/", "main.ridl")
		_, parseErr = p.Parse()
		warnings = p.Warnings()
	})

	require.NoError(t, parseErr)
	require.Len(t, warnings, 1)
	assert.Contains(t, warnings[0], "types.ridl:6: warning: deprecated annotation syntax")
	assert.Equal(t, 1, strings.Count(stderr, "deprecated annotation syntax"))
}

func TestRIDLMethodAnnotationOnProxyMethod(t *testing.T) {
	s, warnings, err := parseStringWithWarnings(annotationsRIDLHeader + `
service Example
  - proxy ContactsService.GetContact
    + @internal
`)
	require.NoError(t, err)
	assert.Empty(t, warnings)

	method := s.Services[0].Methods[0]
	require.Contains(t, method.Annotations, "internal")
}

func TestRIDLMethodAnnotationMissingNameError(t *testing.T) {
	_, _, err := parseStringWithWarnings(annotationsRIDLHeader + `
service Example
  - proxy ContactsService.GetContact
    + @
`)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "expecting annotation name after '@'")
}

func TestRIDLMethodAnnotationLastMethodBeforeNextService(t *testing.T) {
	s, warnings, err := parseStringWithWarnings(annotationsRIDLHeader + `
service Example
  - Ping(PingRequest) => (PingResponse)
    + @internal

service Second
  - Pong(PingRequest) => (PingResponse)
`)
	require.NoError(t, err)
	assert.Empty(t, warnings)

	require.Len(t, s.Services, 2)
	require.Contains(t, s.Services[0].Methods[0].Annotations, "internal")
	assert.Empty(t, s.Services[1].Methods[0].Annotations)
}

func TestRIDLMethodAnnotationMissingColonError(t *testing.T) {
	_, _, err := parseStringWithWarnings(annotationsRIDLHeader + `
service Example
  - Ping(PingRequest) => (PingResponse)
    + @auth admin
`)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "expecting ':' before the annotation value, e.g. '+ @auth:admin'")
}

func TestRIDLMethodAnnotationEmptyValueAtEOFError(t *testing.T) {
	t.Run("new syntax", func(t *testing.T) {
		_, _, err := parseStringWithWarnings(annotationsRIDLHeader + `
service Example
  - Ping(PingRequest) => (PingResponse)
    + @auth:`)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "expecting value after ':' in annotation '@auth'")
	})

	t.Run("old syntax", func(t *testing.T) {
		_, _, err := parseStringWithWarnings(annotationsRIDLHeader + `
service Example
  @auth:`)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "expecting value after ':' in annotation '@auth'")
	})
}
