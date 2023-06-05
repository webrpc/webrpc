package tests

import (
	"embed"
	_ "embed"

	"github.com/webrpc/webrpc/schema/ridl"
)

//go:embed schema/test.ridl
var fs embed.FS

func GetRIDLSchema() string {
	data, _ := fs.ReadFile("schema/test.ridl")
	return string(data)
}

func GetJSONSchema() string {
	schema, _ := ridl.NewParser(fs, "schema/test.ridl").Parse()
	data, _ := schema.ToJSON()
	return data
}
