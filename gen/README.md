# webrpc generators <!-- omit in toc -->

`webrpc-gen` uses Go [text/template](https://pkg.go.dev/text/template) language, along with the webrpc schema AST (abtract-syntax-tree) to generate source code of the target's type system, client library and server handlers.

The Go templates are used in many popular projects including [Hugo](https://gohugo.io/) and [Helm](https://helm.sh). Hugo has a [nice introduction to Go templates](https://gohugo.io/templates/introduction/).

- [Developing a new generator](#developing-a-new-generator)
- [Interoperability tests](#interoperability-tests)
- [Template structure](#template-structure)
  - [Create "main" template](#create-main-template)
  - [Require specific webrpc protocol version](#require-specific-webrpc-protocol-version)
  - [Require specific webrpc-gen version](#require-specific-webrpc-gen-version)
  - [Print help on -help flag](#print-help-on--help-flag)
  - [Set default values for your custom generator options](#set-default-values-for-your-custom-generator-options)
  - [Map webrpc types to your type system](#map-webrpc-types-to-your-type-system)
  - [Split your template into sub-templates](#split-your-template-into-sub-templates)
  - [Create a recursive "type" template](#create-a-recursive-type-template)
- [Template variables](#template-variables)
  - [Default CLI variables](#default-cli-variables)
  - [Custom CLI variables](#custom-cli-variables)
  - [Schema variables](#schema-variables)
- [Template functions](#template-functions)
  - [Go text/template functions](#go-texttemplate-functions)
  - [webrpc-gen functions](#webrpc-gen-functions)


# Developing a new generator

`webrpc-gen` can be invoked against templates located in a local directory:
```
webrpc-gen -schema=api.ridl -target=./local/directory
```

# Interoperability tests

All webrpc generators are expected to implement reference [TestApi schema](../tests/schema/test.ridl) and run client/server interoperability tests against the official [webrpc-test binaries](https://github.com/webrpc/webrpc/releases).

For more info, see [typescript](https://github.com/webrpc/gen-typescript/tree/master/tests) or [golang](https://github.com/webrpc/gen-golang/tree/master/tests) tests.

# Template structure

## Create "main" template

`webrpc-gen` expects at least one `*.go.tmpl` file with the entrypoint template called `"main"`.

```go
{{- define "main" -}}

{{/* Your generator code */}}

{{- end -}}
```

## Require specific webrpc protocol version

```go
{{- if ne .WebrpcVersion "v1" -}}
  {{- stderrPrintf "%s generator error: unsupported webrpc protocol version %s\n" .WebrpcTarget .WebrpcVersion -}}
  {{- exit 1 -}}
{{- end -}}
```

## Require specific webrpc-gen version

Require specific `webrpc-gen` version to ensure the API of the template functions.

```go
{{- if not (minVersion .WebrpcGenVersion "v0.7.0") -}}
  {{- stderrPrintf "%s generator error: unsupported webrpc-gen version %s, please update\n" .WebrpcTarget .WebrpcGenVersion -}}
  {{- exit 1 -}}
{{- end -}}
```

## Print help on -help flag

`webrpc-gen -schema=proto.ridl -target=golang -h`

```go
{{- if exists .Opts "help" -}}
  {{- template "help" $opts -}}
  {{- exit 0 -}}
{{- end -}}
```

## Set default values for your custom generator options

```go
{{- $opts := dict -}}
{{- set $opts "pkg" (default .Opts.pkg "proto") -}}
{{- set $opts "client" (ternary (in .Opts.client "" "true") true false) -}}
{{- set $opts "server" (ternary (in .Opts.server "" "true") true false) -}}

{{- /* Print help on unsupported option. */ -}}
{{- range $k, $v := .Opts }}
  {{- if not (exists $opts $k) -}}
    {{- stderrPrintf "-%v=%q is not supported target option\n\nUsage:\n" $k $v -}}
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
{{- else if isCoreType .Type -}}
    {{ get $typeMap .Type }}
{{- else -}}
    *{{.Type}}
{{- end -}}

{{- end -}}
```

# Template variables

## Default CLI variables

| Variable                     | Description             | Example value                           |
|------------------------------|-------------------------|-----------------------------------------|
| `{{.WebrpcVersion}}`         | webrpc protocol version | `"v1"`                                  |
| `{{.WebrpcGenVersion}}`      | webrpc-gen CLI version  | `"v0.7.0"`                              |
| `{{.WebrpcGenCmd}}`          | webrpc-gen command      | `"webrpc-gen ..."`                      |
| `{{.WebrpcTarget}}`          | webrpc-gen target       | `"github.com/webrpc/gen-golang@v0.7.0"` |

## Custom CLI variables

You can let users pass custom variables into your template by adding custom `-options` to `webrpc-gen` CLI.

| webrpc-gen -option          | Template variable         |
|-----------------------------|---------------------------|
| `-name=HelloService`        | `{{.Opts.name}}`          |
| `-description="some value"` | `{{.Opts.description}}`   |
| `-enableFeature`            | `{{.Opts.someFeature}}`   |

Example:

`webrpc-gen -schema=proto.ridl -target=./custom-template -name=Hello -description="some value" -enableFeature`

will pass `{{.Opts.name}}`, `{{.Opts.description}}` and `{{.Opts.enableFeature}}` variables into your template.

## Schema variables

| Variable                                       | Description                    | Example value               |
|------------------------------------------------|--------------------------------|-----------------------------|
| `{{.SchemaName}}`                              | schema name                    | `"example schema"`          |
| `{{.SchemaVersion}}`                           | schema version                 | `"v0.0.1"`                  |
| `{{.SchemaHash}}`                              | `sha1` schema hash             | `483889fb084764e3a256`      |
| `{{.Errors}}`                                  | schema errors                  | array of schema errors      |
| `{{.Errors[0].Code}}`                          | unique error code              | `1001"` (positive number)   |
| `{{.Errors[0].Name}}`                          | unique error name              | `"RateLimited"`             |
| `{{.Errors[0].Message}}`                       | error description              | `"rate limited, slow down"` |
| `{{.Errors[0].HTTPStatus}}`                    | HTTP response status code      | `429` (number `100`-`599`)  |
| `{{.Types}}`                                   | types                          | array of types              |
| `{{.Types[0].Name}}`                           | type name                      | `"User"`                    |
| `{{.Types[0].Type}}`                           | type                           | `"struct"`                  |
| `{{.Types[0].Fields}}`                         | type fields                    | array of fields             |
| `{{.Types[0].Fields[0].Name}}`                 | field name                     | `"ID"`                      |
| `{{.Types[0].Fields[0].Type}}`                 | field type                     | `"int"`                     |
| `{{.Types[0].Fields[0].Optional}}`             | field optional?                | `false`                     |
| `{{.Types[0].Fields[0].Meta}}`                 | field metadata                 | array of `{"key": "value"}` |
| `{{.Services}}`                                | schema services                | array of services           |
| `{{.Services[0].Name}}`                        | service name                   | `"ExampleService"`          |
| `{{.Services[0].Methods}}`                     | service methods                | array of methods            |
| `{{.Services[0].Methods[0].Inputs}}`           | method inputs                  | array of method inputs      |
| `{{.Services[0].Methods[0].Outputs}}`          | method outputs                 | array of method outputs     |
| `{{.Services[0].Methods[0].Inputs[0].Name}}`   | method input name              | `"header"`                  |
| `{{.Services[0].Methods[0].Inputs[0].Type}}`   | method input type              | `"map<string,string>"`      |
| `{{.Services[0].Methods[0].Outputs[0].Name}}`  | method output name             | `"user"`                    |
| `{{.Services[0].Methods[0].Outputs[0].Type}}`  | method output type             | `"User"`                    |

See the [example schema JSON file](https://github.com/webrpc/webrpc/blob/master/_examples/golang-basics/example.webrpc.json).

For example, you can iterate over the schema methods and print their names:
```go
{{- range $_, $msg := .Services -}}
  {{- range $_, $method := .Methods -}}
    method {{.Name}}()
  {{- end -}}
{{- end -}}
```

# Template functions

## Go text/template functions

| Function                                       | Description                    |
|------------------------------------------------|--------------------------------|
| `and EXPR` | Returns the boolean AND of its arguments by returning the first empty argument or the last argument. That is, "and x y" behaves as "if x then y else x." Evaluation proceeds through the arguments left to right and returns when the result is determined. |
| `call FUNC ARGS...` | Returns the result of calling the first argument, which must be a function, with the remaining arguments as parameters. Thus "call .X.Y 1 2" is, in Go notation, dot.X.Y(1, 2) where Y is a func-valued field, map entry, or the like. The first argument must be the result of an evaluation that yields a value of function type (as distinct from a predefined function such as print). The function must return either one or two result values, the second of which is of type error. If the arguments don't match the function or the returned error value is non-nil, execution stops. |
| `html STRING` | Returns the escaped HTML equivalent of the textual representation of its arguments. This function is unavailable in html/template, with a few exceptions. |
| `index ARRAY 1` | Returns the result of indexing its first argument by the following arguments. Thus "index x 1 2 3" is, in Go syntax, x[1][2][3]. Each indexed item must be a map, slice, or array. |
| `slice ARRAY 1 2` | slice returns the result of slicing its first argument by the remaining arguments. Thus "slice x 1 2" is, in Go syntax, x[1:2], while "slice x" is x[:], "slice x 1" is x[1:], and "slice x 1 2 3" is x[1:2:3]. The first argument must be a string, slice, or array. |
| `js STRING` | Returns the escaped JavaScript equivalent of the textual representation of its arguments. |
| `len ARRAY\|MAP` | Returns the integer length of its argument. |
| `not EXPR` | Returns the boolean negation of its single argument. |
| `or EXPR` | Returns the boolean OR of its arguments by returning the first non-empty argument or the last argument, that is, "or x y" behaves as "if x then x else y". Evaluation proceeds through the arguments left to right and returns when the result is determined. |
| `print "format %v" ARGS...` | Print and format, see Go's [fmt.Sprint()](https://pkg.go.dev/fmt) |
| `printf "format %v" ARGS...` | Print and format, see Go's [fmt.Sprintf()](https://pkg.go.dev/fmt) |
| `println "format %v" ARGS...` | Print and format, see Go's [fmt.Sprintln()](https://pkg.go.dev/fmt) |
| `urlquery STRING` | Returns the escaped value of the textual representation of its arguments in a form suitable for embedding in a URL query. This function is unavailable in html/template, with a few exceptions. |

See https://pkg.go.dev/text/template#hdr-Functions

## webrpc-gen functions

| Template flow                                  | Description                    | webrpc-gen |
|------------------------------------------------|-------------------------------------------------|-------------|
| `minVersion {{.WebrpcVersion}} v1.4`           | Returns `boolean` if the given major/minor semver is at least v1.4 | v0.7.0 |
| `stderrPrintf "format %v" ARGS...`             | `printf` to `webrpc-gen` stderr | v0.7.0 |
| `exit INT`                                     | Terminate template execution, useful for fatal errors | v0.7.0 |

| Schema type helpers                            | Description                    | webrpc-gen |
|------------------------------------------------|-------------------------------------------------|-------------|
| `isBasicType .Type`                            | Returns `true` if `.Type` is [core type](https://github.com/webrpc/webrpc/tree/master/schema#core-types) | v0.7.0 (deprecated) |
| `isCoreType .Type`                            | Returns `true` if `.Type` is [core type](https://github.com/webrpc/webrpc/tree/master/schema#core-types) | v0.9.0 |
| `isStructType .Type`                           | Returns `true` if `.Type` is [struct](https://github.com/webrpc/webrpc/tree/master/schema#struct) | v0.7.0 |
| `isEnumType .Type`                             | Returns `true` if `.Type` is [enum](https://github.com/webrpc/webrpc/tree/master/schema#enum) | v0.7.0 |
| `isMapType .Type`                              | Returns `true` if `.Type` is [map](https://github.com/webrpc/webrpc/tree/master/schema#map) (`map<T1,T2>`) | v0.7.0 |
| `isListType .Type`                             | Returns `true` if `.Type` is [list](https://github.com/webrpc/webrpc/tree/master/schema#list-array) (`[]T`) | v0.7.0 |
| `mapKeyType .MapType`                          | Returns map's key type (ie. `T1` from `map<T1,T2>`) | v0.7.0 |
| `mapValueType .MapType`                        | Returns map's value type (ie. `T2` from `map<T1,T2>`) | v0.7.0 |
| `listElemType .ListType`                       | Returns list's element type (ie. `T` from `[]T`) | v0.7.0 |

| Dictionary (`map[string]any`)                  | Description                    | webrpc-gen |
|------------------------------------------------|-------------------------------------------------|-------------|
| `dict [KEY VALUE]...`                          | Create a new dictionary (`map[string]any`) | v0.7.0 |
| `get $dict KEY`                                | Get value for the given KEY in dictionary | v0.7.0 |
| `set $dict KEY VALUE`                          | Set value for the given KEY in dictionary | v0.7.0 |
| `exists $dict KEY`                             | Returns `true` if the KEY exists in the given dictionary | v0.7.0 |

| String arrays                                  | Description                    | webrpc-gen |
|------------------------------------------------|-------------------------------------------------|-------------|
dev/strings#Join)) | v0.7.0 |
| `array [ELEMENTS]...`                          | Create a new string array | v0.8.0 |
| `append ARRAY [ELEMENTS]...`                   | Append elements to existing string array | v0.8.0 |
| `first ARRAY`                                  | Return first element from the given array | v0.7.0 |
| `join ARRAY SEPARATOR`                         | Join array items with a separator (see [strings.Join()](https://pkg.go.dev/strings#Join)) | v0.7.0 |
| `last ARRAY`                                   | Return last element from the given array | v0.7.0 |
| `sort ARRAY`                                   | Return sorted copy of the given array (ascending order) | v0.8.0 |
| `split SEPARATOR STRING`                       | Split string by a separator into array `[]string` | v0.7.0 |

| String utils                                   | Description                    | webrpc-gen |
|------------------------------------------------|-------------------------------------------------|-------------|
| `hasPrefix STRING PREFIX`                      | Returns `true` if the given string starts with PREFIX | v0.8.0 |
| `hasSuffix STRING SUFFIX`                      | Returns `true` if the given string ends with SUFFIX | v0.8.0 |
| `trimPrefix STRING PREFIX`                     | Trim prefix from a given string | v0.8.0 |
| `trimSuffix STRING SUFFIX`                     | Trim suffix from a given string | v0.8.0 |
| `toLower STRING`                               | Converts input to `"lower case"` | v0.7.0 |
| `toUpper STRING`                               | Converts input to `"UPPER CASE"` | v0.7.0 |
| `firstLetterToLower STRING`                    | Converts first letter to lower case | v0.7.0 |
| `firstLetterToUpper STRING`                    | Converts first letter to UPPER CASE | v0.7.0 |
| `camelCase STRING`                             | Converts input to `"camelCase"` | v0.7.0 |
| `pascalCase STRING`                            | Converts input to `"PascalCase"` | v0.7.0 |
| `snakeCase STRING`                             | Converts input to `"snake_case"` | v0.7.0 |
| `kebabCase STRING`                             | Converts input to `"kebab-case"` | v0.7.0 |

| Generic utils                                  | Description                    | webrpc-gen |
|------------------------------------------------|-------------------------------------------------|-------------|
| `coalesce VALUES...`                           | Returns first non-empty value | v0.7.0 |
| `default VALUE DEFAULT`                        | Returns `DEFAULT` value if given `VALUE` is empty | v0.7.0 |
| `in FIRST VALUES...`                           | Returns `true` if any of the given VALUES match the `first` value | v0.7.0 |
| `ternary BOOL FIRST SECOND`                    | Ternary if-else. Returns first value if `true`, second value if `false` | v0.7.0 |

