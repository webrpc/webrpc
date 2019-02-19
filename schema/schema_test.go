package schema

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchema(t *testing.T) {
	input := `{
		"webrpc": "v1",
		"name": "example",
		"version": "v0.0.1",
		"messages": [
			{
				"name": "Kind",
				"type": "enum",
				"fields": [
					{
						"name": "USER",
						"type": "uint32",
						"value": "1"
					},
					{
						"name": "ADMIN",
						"type": "uint32",
						"value": "2"
					}
				]
			},
			{
				"name": "Empty",
				"type": "struct",
				"fields": [
				]
			},
			{
				"name": "GetUserRequest",
				"type": "struct",
				"fields": [
					{
						"name": "userID",
						"type": "uint64",
						"optional": false
					}
				]
			},
			{
				"name": "User",
				"type": "struct",
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
				"name": "RandomStuff",
				"type": "struct",
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
			}
		],
		"services": [
			{
				"name": "ExampleService",
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
					}
				]
			}
		]
	}
`

	schema, err := ParseSchemaJSON([]byte(input))
	assert.NoError(t, err)

	// spew.Dump(schema)

	jout, err := schema.ToJSON(true)
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
