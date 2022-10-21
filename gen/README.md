# `gen` - webrpc code-generation for multiple language targets

`gen` uses a Go [text/template](https://pkg.go.dev/text/template) language, along with the webrpc schema AST (abtract-syntax-tree)
to generate source code of the Web service's type system, client library and server handlers.

# Template functions

## Base functions from Go text/template package

See https://pkg.go.dev/text/template#hdr-Functions

## String manipulation

### `.camelCase STRING`

Converts input string to "camelCase" (lower camel case) naming convention.
Removes all whitespace and special characters. Supports Unicode characters.

```
{{- range .Methods}}
  {{.Name | camelCase}} = () => {}
{{- end}}
```
