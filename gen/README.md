# `gen` - webrpc code-generation for multiple language targets

`gen` uses a Go [text/template](https://pkg.go.dev/text/template) language, along with the webrpc schema AST (abtract-syntax-tree)
to generate source code of the Web service's type system, client library and server handlers.
