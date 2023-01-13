package tests

import (
	_ "embed"
)

//go:embed schema/test.ridl
var testSchema string

func GetSchema() string {
	return testSchema
}
