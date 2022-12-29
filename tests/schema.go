package tests

import (
	_ "embed"
)

//go:embed schema/api.ridl
var schema string

func GetSchema() string {
	return schema
}
