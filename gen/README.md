# webrpc generator templates

`webrpc-gen` uses Go [text/template](https://pkg.go.dev/text/template) language, along with the webrpc schema AST (abtract-syntax-tree)
to generate source code of the target's type system, client library and server handlers.

`webrpc-gen` can be invoked against templates in local directory:
```
webrpc-gen -schema=api.ridl -target=./local/directory
```

## Create "main" template

`webrpc-gen` expects at least one `*.go.tmpl` file with entrypoint template called `"main"`.

```go
{{- define "main" -}}

{{/* Your generator code */}}

{{- end -}}
```

## Require specific webrpc schema version

```go
{{- if ne .WebrpcVersion "v1" -}}
	{{- stderrPrintf "%s generator error: unsupported webrpc version %s\n" .WebrpcTarget .WebrpcVersion -}}
	{{- exit 1 -}}
{{- end -}}
```

## Print help on -Help flag

`webrpc-gen -schema=proto.ridl -target=golang -Help`

```go
{{- if exists .Opts "Help" -}}
  {{ template "help" $opts }}
  {{- exit 0 -}}
{{- end -}}
```

## Set variables via custom CLI `-Flags`

You can pass custom variables into your template by adding CLI `-Flags` to `webrpc-gen` binary.

| webrpc-gen CLI flag          | Template variable              |
|------------------------------|--------------------------------|
| `-UpperName=Hello`           | `{{.Opts.UpperName}}`          |
| `-PassYourOwn="value"`       | `{{.Opts.PassYourOwn}}`        |

Example:

`webrpc-gen -schema=proto.ridl -target=./custom-template -UpperName=Hello -PassYourOwn="value"`

will pass `{{.UpperName}}` and `{{.PassYourOwn}}` variables into your template.

## Set default values for your custom options

```go
{{- /* Options with default values. */ -}}
{{- $opts := dict -}}
{{- set $opts "Pkg" (coalesce .Opts.Pkg "proto") -}}
{{- set $opts "Client" (coalesce .Opts.Version false) -}}
{{- set $opts "Server" (coalesce .Opts.Version false) -}}

{{- /* Print help on unsupported option. */ -}}
{{- range $k, $v := .Opts }}
  {{- if not (exists $opts $k) -}}
    {{- stderrPrintf "  -%v=%q is not supported option\n" $k $v -}}
    {{- template "help" $opts -}}
    {{- exit 1 -}}
  {{- end -}}
{{- end -}}
```

## Map webrpc types to your type system
```go
{{- /* Type mapping. */ -}}
{{- $typeMap := dict }}
{{- set $typeMap "null" "null" -}}
{{- set $typeMap "any" "object" -}}
{{- set $typeMap "byte" "string" -}}
{{- set $typeMap "bool" "boolean" -}}
{{- set $typeMap "uint" "number" -}}
{{- set $typeMap "uint8" "number" -}}
{{- set $typeMap "uint16" "number" -}}
{{- set $typeMap "uint32" "number" -}}
{{- set $typeMap "uint64" "number" -}}
{{- set $typeMap "int" "number" -}}
{{- set $typeMap "int8" "number" -}}
{{- set $typeMap "int16" "number" -}}
{{- set $typeMap "int32" "number" -}}
{{- set $typeMap "int64" "number" -}}
{{- set $typeMap "float32" "number" -}}
{{- set $typeMap "float64" "number" -}}
{{- set $typeMap "string" "string" -}}
{{- set $typeMap "timestamp" "string" -}}
{{- set $typeMap "map" "object" -}}
{{- set $typeMap "[]" "array" -}}
```

Call `{{ get $typeMap .Type }}` to print your type.

## Split your template into sub-templates

You can import template.
```go
{{ template "sub-template" }}
```

Use `dict` function to pass multiple variables into the sub-template:

```go
{{ template "sub-template" dict "Type" .Type "TypeMap" $typeMap }}
```

## Create a recursive "type" template

Base webrpc types can be nested (ie. `map<string,map<string,User>>`), so you will need to render them recursively.

```go
{{- define "type" -}}
{{- $typeMap := .TypeMap -}}

{{- if isMapType .Type -}}
    map[{{mapKeyType .Type}}]{{template "type" dict "Type" (mapValueType .Type) "TypeMap" $typeMap}}
{{- else if isArrayType .Type -}}
    []{{template "type" dict "Type" (arrayItemType .Type) "TypeMap" $typeMap}}
{{- else if isBaseType .Type -}}
    {{ get $typeMap .Type }}
{{- else -}}
    *{{.Type}}
{{- end -}}

{{- end -}}
```

# Template variables

## Built-in template variables

| Variable                     | Description             | Example value                           |
|------------------------------|-------------------------|-----------------------------------------|
| `{{.WebrpcProtocolVersion}}` | webrpc protocol version | `"v1"`                                  |
| `{{.WebrpcGenVersion}}`      | webrpc-gen CLI version  | `"v0.7.0"`                              |
| `{{.WebrpcGenCmd}}`          | webrpc-gen command      | `"webrpc-gen ..."`                      |
| `{{.WebrpcTarget}}`          | webrpc-gen target       | `"github.com/webrpc/gen-golang@v0.7.0"` |

## Variables 
| Variable                     | Description                    | Example value              |
|------------------------------|--------------------------------|----------------------------|
| `{{.SchemaName}}`            | schema name                    | `"example schema"`         |
| `{{.SchemaVersion}}`         | schema version                 | `"v0.0.1"`                 |
| `{{.SchemaHash}}`            | `sha1` schema hash             | object                     |
| `{{.Imports}}`               | schema imports                 | object                     |
| `{{.Messages}}`              | schema messages                | object                     |
| `{{.Services}}`              | schema services                | object                     |

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
