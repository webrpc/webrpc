# webrpc generators <!-- omit in toc -->

`webrpc-gen` uses Go [text/template](https://pkg.go.dev/text/template) language, along with the webrpc schema AST (abtract-syntax-tree)
to generate source code of the target's type system, client library and server handlers.

`webrpc-gen` can be invoked against templates in a local directory:
```
webrpc-gen -schema=api.ridl -target=./local/directory
```

- [Go templates](#go-templates)
  - [Create "main" template](#create-main-template)
  - [Require specific webrpc schema version](#require-specific-webrpc-schema-version)
  - [Print help on -Help flag](#print-help-on--help-flag)
  - [Set variables via custom CLI `-Flags`](#set-variables-via-custom-cli--flags)
  - [Set default values for your custom options](#set-default-values-for-your-custom-options)
  - [Map webrpc types to your type system](#map-webrpc-types-to-your-type-system)
  - [Split your template into sub-templates](#split-your-template-into-sub-templates)
  - [Create a recursive "type" template](#create-a-recursive-type-template)
- [Template variables](#template-variables)
  - [CLI variables](#cli-variables)
  - [Schema variables](#schema-variables)
- [Template functions](#template-functions)
  - [Base functions from Go text/template package](#base-functions-from-go-texttemplate-package)
  - [String manipulation](#string-manipulation)
    - [`.camelCase STRING`](#camelcase-string)

# Go templates

## Create "main" template

`webrpc-gen` expects at least one `*.go.tmpl` file with the entrypoint template called `"main"`.

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

Import a sub-template.

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

## CLI variables

| Variable                     | Description             | Example value                           |
|------------------------------|-------------------------|-----------------------------------------|
| `{{.WebrpcProtocolVersion}}` | webrpc protocol version | `"v1"`                                  |
| `{{.WebrpcGenVersion}}`      | webrpc-gen CLI version  | `"v0.7.0"`                              |
| `{{.WebrpcGenCmd}}`          | webrpc-gen command      | `"webrpc-gen ..."`                      |
| `{{.WebrpcTarget}}`          | webrpc-gen target       | `"github.com/webrpc/gen-golang@v0.7.0"` |

## Schema variables 

| Variable                                       | Description                    | Example value               |
|------------------------------------------------|--------------------------------|-----------------------------|
| `{{.SchemaName}}`                              | schema name                    | `"example schema"`          |
| `{{.SchemaVersion}}`                           | schema version                 | `"v0.0.1"`                  |
| `{{.SchemaHash}}`                              | `sha1` schema hash             | `483889fb084764e3a256`      |
| `{{.Imports}}`                                 | schema imports                 | array of imports            |
| `{{.Messages}}`                                | schema messages                | array of messages           |
| `{{.Services}}`                                | schema services                | array of services           |
| `{{.Messages[0].Name}}`                        | messages name                  | `"User"`                    |
| `{{.Messages[0].Type}}`                        | messages type                  | `"struct"`                  |
| `{{.Messages[0].Fields}}`                      | messages fields                | array                       |
| `{{.Messages[0].Fields[0].Name}}`              | message name                   | `"ID"`                      |
| `{{.Messages[0].Fields[0].Type}}`              | message type                   | `"int"`                     |
| `{{.Messages[0].Fields[0].Optional}}`          | messages fields                | `false`                     |
| `{{.Messages[0].Fields[0].Meta}}`              | messages fields                | array of `{"key": "value"}` |
| `{{.Services[0].Name}}`                        | service name                   | `"ExampleService"`          |
| `{{.Services[0].Methods}}`                     | service methods                | array of methods            |
| `{{.Services[0].Methods[0].Inputs}}`           | method inputs                  | array of inputs             |
| `{{.Services[0].Methods[0].Outputs}}`          | method outputs                 | array of outputs            |
| `{{.Services[0].Methods[0].Inputs[0].Name}}`   | input name                     | `"header"`                  |
| `{{.Services[0].Methods[0].Inputs[0].Type}}`   | input type                     | `"map<string,string>"`      |
| `{{.Services[0].Methods[0].Outputs[0].Name}}`  | output name                    | `"user"`                    |
| `{{.Services[0].Methods[0].Outputs[0].Type}}`  | output type                    | `"User"`                    |

See the [example schema JSON file](https://github.com/webrpc/webrpc/blob/master/_examples/golang-basics/example.webrpc.json).

For example, you can iterate over the schema methods and print their names:
```go
{{- range $i, $msg := .Messages -}}
  {{- range $i, $method := .Methods -}}
    method {{.Name}}()
  {{- end -}}
{{- end -}}
```

# Template functions

## Base functions from Go text/template package

See https://pkg.go.dev/text/template#hdr-Functions

## String manipulation

### `.camelCase STRING`

Converts input string to "camelCase" (lower camel case) naming convention.
Removes all whitespace and special characters. Supports Unicode characters.

```go
{{- range .Methods}}
  {{.Name | camelCase}} = () => {}
{{- end}}
```
