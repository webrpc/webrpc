package tests

import (
	"flag"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/webrpc/webrpc/schema/ridl"
)

var updateFlag = flag.String("update", "", "update golden file to match tests' current behavior")

func TestRIDLSchemaAgainstJSON(t *testing.T) {
	exampleDirFS := os.DirFS("./schema")

	schema, err := ridl.NewParser(exampleDirFS, "/", "test.ridl").Parse()
	assert.NoError(t, err)

	jsonSchema, err := schema.ToJSON()
	assert.NoError(t, err)

	current := []byte(jsonSchema)

	golden, err := os.ReadFile("./_testdata/test.golden.json")
	assert.NoError(t, err)

	if *updateFlag == "./_testdata/test.golden.json" {
		assert.NoError(t, os.WriteFile("./_testdata/test.golden.json", current, 0644))
		return
	}

	if !cmp.Equal(golden, current) {
		t.Error(cmp.Diff(golden, current))
		t.Log("To update the golden file, run `go test ./tests -update=./_testdata/test.golden.json'")
	}
}
