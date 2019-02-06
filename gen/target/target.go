package target

import (
	"github.com/webrpc/webrpc/schema"
)

type Generator interface {
	Gen(proto *schema.WebRPCSchema, opts Options) (string, error)
}

var Generators = map[string]Generator{}

func Register(target string, generator Generator) {
	Generators[target] = generator
}

func GetGenerator(target string) Generator {
	g, _ := Generators[target]
	return g
}

type Options struct {
	PkgName   string
	Client    bool
	Server    bool
	Websocket bool
}
