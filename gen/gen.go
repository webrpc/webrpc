package gen

import (
	"github.com/webrpc/webrpc/schema"
)

type Generator interface {
	Gen(proto *schema.WebRPCSchema, opts TargetOptions) (string, error)
}

var Generators = map[string]Generator{}

func Register(target string, generator Generator) {
	Generators[target] = generator
}

func GetGenerator(target string) Generator {
	g, _ := Generators[target]
	return g
}

type TargetOptions struct {
	PkgName   string
	Client    bool
	Server    bool
	Extra     string
	Websocket bool
}
