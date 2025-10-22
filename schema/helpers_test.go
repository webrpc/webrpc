package schema

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSchemaBigIntFieldsByType(t *testing.T) {
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
						"name": "balance",
						"type": "bigint",
						"optional": false
					},
					{
						"name": "extra",
						"type": "Extra",
						"optional": false
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
					},
					{
						"name": "byBN",
						"type": "bigint",
						"optional": false
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
					},
					{
						"name": "content",
						"type": "string",
						"optional": false
					},
					{
						"name": "largeNum",
						"type": "bigint",
						"optional": false
					}
				]
			},
			{
				"kind": "struct",
				"name": "Extra",
				"fields": [
					{
						"name": "info",
						"type": "string",
						"optional": false
					},
					{
						"name": "amount",
						"type": "bigint",
						"optional": false
					},
					{
						"name": "points",
						"type": "[]bigint",
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
						"inputs": [
							{
								"name": "counter",
								"type": "bigint",
								"optional": true
							}
						],
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
	require.NoError(t, err)

	bigIntFields := SchemaBigIntFieldsByType(schema)

	expected := map[string][]any{
		"GetArticleRequest":  {"byBN"},
		"GetArticleResponse": {"largeNum"},
		"GetUserResponse":    {[2]string{"user", "User"}},
		"User":               {"balance", [2]string{"extra", "Extra"}},
		"Extra":              {"amount", "points[]"},
		"RandomStuff":        {[2]string{"listOfUsers", "User[]"}, [2]string{"user", "User"}},
		"PingRequest":        {"counter"},
	}

	require.Equal(t, expected, bigIntFields)

}
