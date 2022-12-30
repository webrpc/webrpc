package tests

import (
	_ "embed"
)

//go:embed schema/test.ridl
var schema string

func GetSchema() string {
	return schema
}
