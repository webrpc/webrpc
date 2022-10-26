# `gen` - webrpc code-generation for multiple language targets

`gen` uses a Go [text/template](https://pkg.go.dev/text/template) language, along with the webrpc schema AST (abtract-syntax-tree)
to generate source code of the Web service's type system, client library and server handlers.

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

## Set template-specific variables as CLI -Flags

| webrpc-gen CLI flag          | Template variable              |
|------------------------------|--------------------------------|
| `-UpperName=Hello`           | `{{.Opts.UpperName}}`          |
| `-PassYourOwn="value"`       | `{{.Opts.PassYourOwn}}`        |

Example:
`webrpc-gen -schema=proto.ridl -target=./custom-template -UpperName=Hello -PassYourOwn="value"`

will pass `{{.UpperName}}` and `{{PassYourOwn}}` variables into the custom template generator.

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
