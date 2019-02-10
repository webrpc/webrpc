# `gen` - webrpc code-generation for multiple language targets

`gen` uses a basic templating language, along with the webrpc schema AST (abtract-syntax-tree)
to generate source code of the Web service's type system, client library and server handlers.

## Supported targets

* `go` - gen/golang
* `ts` - gen/typescript
* `js` - gen/javascript


## Adding a new target

Adding a new target is easy, just add a folder under `gen/` and copy one of the existing
targets and start adapting it for your language. Run `make build` to re-generate all templates
and execute the code-generator via the `go:generate` in the target package you make.
