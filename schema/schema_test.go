package schema

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestSchema(t *testing.T) {
	input := `{
		"webrpc": "v1",
		"name": "example",
		"version": "v0.0.1",
		"types": [
			{
				"kind": "enum",
				"name": "Kind",
				"type": "uint32",
				"fields": [
					{
						"name": "USER",
						"value": "1"
					},
					{
						"name": "ADMIN",
						"value": "2"
					}
				]
			},
			{
				"kind": "struct",
				"name": "Empty",
				"fields": [
				]
			},
			{
				"kind": "struct",
				"name": "GetUserRequest",
				"fields": [
					{
						"name": "userID",
						"type": "uint64",
						"optional": false
					}
				]
			},
			{
				"kind": "struct",
				"name": "RandomStuff",
				"fields": [
					{
						"name": "meta",
						"type": "map<string,any>"
					},
					{
						"name": "metaNestedExample",
						"type": "map<string,map<string,uint32>>"
					},
					{
						"name": "namesList",
						"type": "[]string"
					},
					{
						"name": "numsList",
						"type": "[]int64"
					},
					{
						"name": "doubleArray",
						"type": "[][]string"
					},
					{
						"name": "listOfMaps",
						"type": "[]map<string,uint32>"
					},
					{
						"name": "listOfUsers",
						"type": "[]User"
					},
					{
						"name": "mapOfUsers",
						"type": "map<string,User>"
					},
					{
						"name": "user",
						"type": "User"
					}
				]
			},
			{
				"kind": "struct",
				"name": "User",
				"fields": [
					{
						"name": "ID",
						"type": "uint64",
						"optional": false,
						"meta": [
							{ "json": "id" },
							{ "go.tag.db": "id" }
						]
					},
					{
						"name": "username",
						"type": "string",
						"optional": false,
						"meta": [
							{ "json": "USERNAME" },
							{ "go.tag.db": "username" }
						]
					},
					{
						"name": "createdAt",
						"type": "timestamp",
						"optional": true,
						"meta": [
							{ "go.tag.json": "created_at,omitempty" },
							{ "go.tag.db": "created_at" }
						]
					}
				]
			},
			{
				"kind": "struct",
				"name": "GetArticleRequest",
				"fields": [
					{
						"name": "articleId",
						"type": "uint32",
						"optional": false,
						"meta": [
							{ "go.field.name": "ArticleID" }
						]
					}
				]
			},
			{
				"kind": "struct",
				"name": "GetArticleResponse",
				"fields": [
					{
						"name": "title",
						"type": "string",
						"optional": false
					}
				]
			}
		],
		"services": [
			{
				"name": "Example",
				"methods": [
					{
						"name": "Ping",
						"inputs": [],
						"outputs": [
							{
								"name": "status",
								"type": "bool"
							}
						]
					},
					{
						"name": "GetUser",
						"inputs": [
							{
								"name": "header",
								"type": "map<string,string>"
							},
							{
								"name": "req",
								"type": "GetUserRequest"
							}
						],
						"outputs": [
							{
								"name": "code",
								"type": "uint32"
							},
							{
								"name": "user",
								"type": "User"
							}
						]
					},
					{
						"name": "GetArticle",
						"succinct": true,
						"inputs": [
							{
								"name": "getArticleRequest",
								"type": "GetArticleRequest"
							}
						],
						"outputs": [
							{
								"name": "getArticleResponse",
								"type": "GetArticleResponse"
							}
						]
					}
				]
			}
		]
	}
`

	schema, err := ParseSchemaJSON([]byte(input))
	assert.NoError(t, err)

	// spew.Dump(schema)

	jout, err := schema.ToJSON()
	assert.NoError(t, err)
	fmt.Println("schema JSON:", jout)
}

func TestFoo(t *testing.T) {
	assert.True(t, IsValidArgName("a"))
	assert.True(t, IsValidArgName("a1"))
	assert.False(t, IsValidArgName("1a"))
	assert.False(t, IsValidArgName("12222"))
	assert.False(t, IsValidArgName("1_"))
	assert.False(t, IsValidArgName("1_2"))
	assert.False(t, IsValidArgName(""))
	assert.True(t, IsValidArgName("a"))
	assert.False(t, IsValidArgName("-1_2"))
	assert.False(t, IsValidArgName("_ddd"))
	assert.False(t, IsValidArgName("aAA--"))
	assert.True(t, IsValidArgName("a3_44aass"))
	assert.True(t, IsValidArgName("a55_____cdDDDD"))
	assert.False(t, IsValidArgName("asSS_E_##$"))
}

func TestMatchServices(t *testing.T) {
	type args struct {
		s        *WebRPCSchema
		services []string
	}
	tests := []struct {
		name string
		args args
		want *WebRPCSchema
	}{
		{
			name: "match services",
			args: args{
				s: &WebRPCSchema{
					WebrpcVersion: "v0.1.0",
					SchemaName:    "dev",
					SchemaVersion: "dev-v0.1.0",
					Services: []*Service{
						{Name: "ExampleService"},
						{Name: "AdminService"},
						{Name: "PublicService"},
					},
				},
				services: []string{"ExampleService", "AdminService"},
			},
			want: &WebRPCSchema{
				WebrpcVersion: "v0.1.0",
				SchemaName:    "dev",
				SchemaVersion: "dev-v0.1.0",
				Services: []*Service{
					{Name: "ExampleService"},
					{Name: "AdminService"},
				},
			},
		},
		{
			name: "no services to match",
			args: args{
				s: &WebRPCSchema{
					WebrpcVersion: "v0.1.0",
					SchemaName:    "dev",
					SchemaVersion: "dev-v0.1.0",
					Services: []*Service{
						{Name: "ExampleService"},
						{Name: "AdminService"},
					},
				},
				services: []string{},
			},
			want: &WebRPCSchema{
				WebrpcVersion: "v0.1.0",
				SchemaName:    "dev",
				SchemaVersion: "dev-v0.1.0",
				Services: []*Service{
					{Name: "ExampleService"},
					{Name: "AdminService"},
				},
			},
		},
		{
			name: "no matching services",
			args: args{
				s: &WebRPCSchema{
					WebrpcVersion: "v0.1.0",
					SchemaName:    "dev",
					SchemaVersion: "dev-v0.1.0",
					Services: []*Service{
						{Name: "ExampleService"},
						{Name: "AdminService"},
					},
				},
				services: []string{"NonExistentService"},
			},
			want: &WebRPCSchema{
				WebrpcVersion: "v0.1.0",
				SchemaName:    "dev",
				SchemaVersion: "dev-v0.1.0",
				Services:      []*Service(nil),
			},
		},
		{
			name: "all services match",
			args: args{
				s: &WebRPCSchema{
					WebrpcVersion: "v0.1.0",
					SchemaName:    "dev",
					SchemaVersion: "dev-v0.1.0",
					Services: []*Service{
						{Name: "ExampleService"},
						{Name: "AdminService"},
					},
				},
				services: []string{"ExampleService", "AdminService"},
			},
			want: &WebRPCSchema{
				WebrpcVersion: "v0.1.0",
				SchemaName:    "dev",
				SchemaVersion: "dev-v0.1.0",
				Services: []*Service{
					{Name: "ExampleService"},
					{Name: "AdminService"},
				},
			},
		},
		{
			name: "empty service list in schema",
			args: args{
				s: &WebRPCSchema{
					WebrpcVersion: "v0.1.0",
					SchemaName:    "dev",
					SchemaVersion: "dev-v0.1.0",
					Services:      []*Service{}, // No services in schema
				},
				services: []string{"ExampleService", "AdminService"},
			},
			want: &WebRPCSchema{
				WebrpcVersion: "v0.1.0",
				SchemaName:    "dev",
				SchemaVersion: "dev-v0.1.0",
				Services:      []*Service(nil), // Nothing to match, return empty
			},
		},
		{
			name: "nil schema",
			args: args{
				s:        nil, // Schema is nil
				services: []string{"ExampleService", "AdminService"},
			},
			want: nil, // Expect nil in return
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.EqualValues(t, tt.want, MatchServices(tt.args.s, tt.args.services))
		})
	}
}

func TestIgnoreMethodsWithAnnotations(t *testing.T) {
	type args struct {
		s                 *WebRPCSchema
		ignoreAnnotations map[string]string
	}
	tests := []struct {
		name string
		args args
		want *WebRPCSchema
	}{
		{
			name: "no methods to ignore",
			args: args{
				s: &WebRPCSchema{
					Services: []*Service{
						{
							Name: "ExampleService",
							Methods: []*Method{
								{
									Name:        "Ping",
									Annotations: map[string]*Annotation{},
								},
								{
									Name:        "Version",
									Annotations: map[string]*Annotation{},
								},
							},
						},
					},
				},
				ignoreAnnotations: map[string]string{"deprecated": ""},
			},
			want: &WebRPCSchema{
				Services: []*Service{
					{
						Name: "ExampleService",
						Methods: []*Method{
							{
								Name:        "Ping",
								Annotations: map[string]*Annotation{},
							},
							{
								Name:        "Version",
								Annotations: map[string]*Annotation{},
							},
						},
					},
				},
			},
		},
		{
			name: "all methods ignored",
			args: args{
				s: &WebRPCSchema{
					Services: []*Service{
						{
							Name: "ExampleService",
							Methods: []*Method{
								{
									Name: "Ping",
									Annotations: map[string]*Annotation{
										"deprecated": {
											AnnotationType: "deprecated",
											Value:          "PingV2",
										},
									},
								},
								{
									Name: "Version",
									Annotations: map[string]*Annotation{
										"deprecated": {
											AnnotationType: "deprecated",
											Value:          "",
										},
									},
								},
							},
						},
					},
				},
				ignoreAnnotations: map[string]string{"deprecated": ""},
			},
			want: &WebRPCSchema{
				Services: []*Service{
					{
						Name:    "ExampleService",
						Methods: []*Method(nil), // All methods are ignored
					},
				},
			},
		},
		{
			name: "no matching annotations",
			args: args{
				s: &WebRPCSchema{
					Services: []*Service{
						{
							Name: "ExampleService",
							Methods: []*Method{
								{
									Name: "Ping",
									Annotations: map[string]*Annotation{
										"internal": {
											AnnotationType: "internal",
											Value:          "",
										},
									},
								},
							},
						},
					},
				},
				ignoreAnnotations: map[string]string{"deprecated": ""},
			},
			want: &WebRPCSchema{
				Services: []*Service{
					{
						Name: "ExampleService",
						Methods: []*Method{
							{
								Name: "Ping",
								Annotations: map[string]*Annotation{
									"internal": {
										AnnotationType: "internal",
										Value:          "",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "empty schema",
			args: args{
				s: &WebRPCSchema{
					Services: []*Service{}, // No services in the schema
				},
				ignoreAnnotations: map[string]string{"deprecated": ""},
			},
			want: &WebRPCSchema{
				Services: []*Service{}, // Expect no services
			},
		},
		{
			name: "empty annotations map",
			args: args{
				s: &WebRPCSchema{
					Services: []*Service{
						{
							Name: "ExampleService",
							Methods: []*Method{
								{
									Name:        "Ping",
									Annotations: map[string]*Annotation{},
								},
							},
						},
					},
				},
				ignoreAnnotations: map[string]string{},
			},
			want: &WebRPCSchema{
				Services: []*Service{
					{
						Name: "ExampleService",
						Methods: []*Method{
							{
								Name:        "Ping",
								Annotations: map[string]*Annotation{},
							},
						},
					},
				},
			},
		},
		{
			name: "nil schema",
			args: args{
				s:                 nil,
				ignoreAnnotations: map[string]string{"deprecated": ""},
			},
			want: nil, // Expect nil when the schema is nil
		},
		{
			name: "nil annotations map",
			args: args{
				s: &WebRPCSchema{
					Services: []*Service{
						{
							Name: "ExampleService",
							Methods: []*Method{
								{
									Name: "Ping",
									Annotations: map[string]*Annotation{
										"deprecated": {
											AnnotationType: "deprecated",
											Value:          "v1",
										},
									},
								},
							},
						},
					},
				},
				ignoreAnnotations: nil, // Annotations map is nil
			},
			want: &WebRPCSchema{
				Services: []*Service{
					{
						Name: "ExampleService",
						Methods: []*Method{
							{
								Name: "Ping",
								Annotations: map[string]*Annotation{
									"deprecated": {
										AnnotationType: "deprecated",
										Value:          "v1",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "filter deprecated methods",
			args: args{
				s: &WebRPCSchema{
					Services: []*Service{
						{
							Name: "ExampleService",
							Methods: []*Method{
								{
									Name: "Ping",
									Annotations: map[string]*Annotation{
										"internal": {
											AnnotationType: "internal",
											Value:          "",
										},
										"auth": {
											AnnotationType: "auth",
											Value:          "Cookies,Jwt,Authorization",
										},
									},
								},
								{
									Name: "Version",
									Annotations: map[string]*Annotation{
										"internal": {
											AnnotationType: "internal",
											Value:          "",
										},
										"deprecated": {
											AnnotationType: "deprecated",
										},
									},
								},
								{
									Name:        "Version2",
									Annotations: map[string]*Annotation{},
								},
							},
						},
					},
				},
				ignoreAnnotations: map[string]string{"deprecated": ""},
			},
			want: &WebRPCSchema{
				Services: []*Service{
					{
						Name: "ExampleService",
						Methods: []*Method{
							{
								Name: "Ping",
								Annotations: map[string]*Annotation{
									"internal": {
										AnnotationType: "internal",
										Value:          "",
									},
									"auth": {
										AnnotationType: "auth",
										Value:          "Cookies,Jwt,Authorization",
									},
								},
							},
							{
								Name:        "Version2",
								Annotations: map[string]*Annotation{},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.EqualValues(t, tt.want, IgnoreMethodsWithAnnotations(tt.args.s, tt.args.ignoreAnnotations))
		})
	}
}

func TestMatchMethodsWithAnnotations(t *testing.T) {
	type args struct {
		s                *WebRPCSchema
		matchAnnotations map[string]string
	}
	tests := []struct {
		name string
		args args
		want *WebRPCSchema
	}{
		{
			name: "no methods to ignore",
			args: args{
				s: &WebRPCSchema{
					Services: []*Service{
						{
							Name: "ExampleService",
							Methods: []*Method{
								{
									Name:        "Ping",
									Annotations: map[string]*Annotation{},
								},
								{
									Name: "Version",
									Annotations: map[string]*Annotation{
										"deprecated": {
											AnnotationType: "deprecated",
											Value:          "",
										},
									},
								},
							},
						},
					},
				},
				matchAnnotations: map[string]string{"deprecated": ""},
			},
			want: &WebRPCSchema{
				Services: []*Service{
					{
						Name: "ExampleService",
						Methods: []*Method{
							{
								Name: "Version",
								Annotations: map[string]*Annotation{
									"deprecated": {
										AnnotationType: "deprecated",
										Value:          "",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.EqualValues(t, tt.want, MatchMethodsWithAnnotations(tt.args.s, tt.args.matchAnnotations))
		})
	}
}
