webrpc-gen Golang templates
===============================

This repo contains the templates used by the `webrpc-gen` cli to code-generate
webrpc Go server and client code.


## Usage

```
webrpc-gen -schema=example.ridl -target=golang -out=./example.gen.go -pkg=main -server -client

# or 
webrpc-gen -schema=example.ridl -target=github.com/webrpc/gen-golang@v0.6.0 -out=./example.gen.go -pkg=main -server -client

# or
webrpc-gen -schema=example.ridl -target=./local-go-templates-on-disk -out=./example.gen.go -pkg=main -server -client
```

As you can see, the `-target` supports default `golang`, any git URI, or a local folder :)

### Set custom template variables
Change any of the following values by passing `-option="Value"` CLI flag to `webrpc-gen`.

| webrpc-gen -option    | Default   | Description                                                                 | Added in |
|-----------------------|-----------|-----------------------------------------------------------------------------|----------|
| `-pkg=<name>`         | `"proto"` | package name                                                                | v0.5.0   |
| `-client`             | `false`   | generate client code                                                        | v0.5.0   |
| `-server`             | `false`   | generate server code                                                        | v0.5.0   |
| `-types=false`        | `true`    | don't generate types                                                        | v0.13.0  |
| `-json=sonic`         |           | use [sonic](https://github.com/bytedance/sonic) for JSON encoding           | v0.18.0  |
| `-json=jsoniter`      |           | use [jsoniter](https://github.com/json-iterator/go) for JSON encoding       | v0.12.0  |
| `-json=<pkg>`         |           | use alternative drop-in replacement import path for JSON encoding package   | v0.18.0  |
| `-fixEmptyArrays`     | `false`   | `encoding/json`: fix `null` arrays with reflect (see Go [#27589][go27589])  | v0.13.0  |
| `-errorStackTrace`    | `false`   | enables error stack traces                                                  | v0.14.0  |
| `-webrpcHeader=false` | `true`    | enable client send webrpc version in http headers                           | v0.16.0  |
| `-schemaHash=false`   | `true`    | don't emit schema hash + version helper funcs (avoids merge conflicts)      | v0.30.0  |

Example:
```
webrpc-gen -schema=./proto.json -target=golang -out server.gen.go -pkg=main -server
```

## Set custom Go field meta tags in your RIDL file

| CLI option flag                              | Description                                                      |
|----------------------------------------------|------------------------------------------------------------------|
| `+ go.field.name = ID`                       | Set custom field name                                            |
| `+ go.field.type = uuid.UUID`                | Set custom field type (must be able to JSON unmarshal the value) |
| `+ go.type.import = github.com/google/uuid`  | Set custom field type's import path                              |
| `+ go.tag.json = id`                         | Set `json:"id"` struct tag                                       |
| `+ go.tag.db = id`                           | Set `db:"id"` struct tag                                         |

Example:
```ridl
struct User
  - ID: int64
    + go.tag.db = id
    + go.tag.json = id
  - UUID: string
    + go.field.type = uuid.UUID
    + go.type.import = github.com/google/uuid
    + go.tag.json = uuid
    + go.tag.db = uuid
  - Age: int
    + go.tag.db: age
  - Name: string
    + go.tag.db = name
  - PasswordHash: string
    + go.tag.db = passwd_hash
```

will result in

```go
import "github.com/google/uuid"

type User struct {
	ID           int64     `json:"id" db:"id"`
	UUID         uuid.UUID `json:"uuid" db:"uuid"`
	Name         string    `db:"name"`
	Age          int       `db:"age"`
	PasswordHash string    `json:"-" db:"passwd_hash"`
}
```

## Examples

See [_examples](./_examples)

## LICENSE

[MIT LICENSE](./LICENSE)

[go27589]: https://github.com/golang/go/issues/27589