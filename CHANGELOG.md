- [webrpc v0.11.0](#webrpc-v0110)
  - [Feature: Define webrpc schema errors](#feature-define-webrpc-schema-errors)
  - [typescript@v0.11.0 breaking changes](#typescriptv0110-breaking-changes)
  - [golang@v0.11.0 breaking changes](#golangv0110-breaking-changes)
- [webrpc v0.10.0](#webrpc-v0100)
  - [Interoperability tests](#interoperability-tests)
  - [Breaking changes in webrpc-gen Go API](#breaking-changes-in-webrpc-gen-go-api)
- [webrpc v0.9.0](#webrpc-v090)
  - [Breaking changes](#breaking-changes)
    - [RIDL v0.9.0 changes](#ridl-v090-changes)
    - [JSON schema v0.9.0 changes](#json-schema-v090-changes)
    - [Template changes](#template-changes)
  - [Migration guide](#migration-guide)
    - [RIDL v0.9.0 migration guide](#ridl-v090-migration-guide)
    - [JSON schema v0.9.0 migration guide](#json-schema-v090-migration-guide)
    - [Generator templates v0.9.0 migration guide](#generator-templates-v090-migration-guide)

# webrpc v0.11.0

## Feature: Define webrpc schema errors

You can now define your own custom schema errors in RIDL file, for example:

```ridl
error   1 Unauthorized    "unauthorized"        HTTP 401
error   2 ExpiredToken    "expired token"       HTTP 401
error   3 InvalidToken    "invalid token"       HTTP 401
error   4 Deactivated     "account deactivated" HTTP 403
error   5 ConfirmAccount  "confirm your email"  HTTP 403
error   6 AccessDenied    "access denied"       HTTP 403
error   7 MissingArgument "missing argument"    HTTP 400
error   8 UnexpectedValue "unexpected value"    HTTP 400
error 100 RateLimited     "too many requests"   HTTP 429
error 101 DatabaseDown    "service outage"      HTTP 503
error 102 ElasticDown     "search is degraded"  HTTP 503
error 103 NotImplemented  "not implemented"     HTTP 501
error 200 UserNotFound    "user not found"
error 201 UserBusy        "user busy"
error 202 InvalidUsername "invalid username"
error 300 FileTooBig      "file is too big (max 1GB)"
error 301 FileInfected    "file is infected"
error 302 FileType        "unsupported file type"
```

Note: Unless specified, the default HTTP status for webrpc errors is `HTTP 400`.

## typescript@v0.11.0 breaking changes

- All errors thrown by webrpc client are now instance of `WebrpcError`, which extends JavaScript `Error`. No need to re-throw errors anymore.
- ~`error.msg`~ `error.message`
- by default, the error messages are "human-friendly", they don't contain any details about the backend error cause
- underlying backend error (for developers) is optionally available as `error.cause?`
- `error.code` or `error.message` can be used as input for user-friendly error i18n translations

You can now check for explicit error class instance (as defined in RIDL schema) or against a generic `WebrpcError` class.

```js
try {
  const resp = await testApiClient.getUser();
  // setUser(resp.user)
} catch (error) {
  if (error instanceof RateLimitedError) {
    // retry with back-off time
  }
  
  if (error instanceof UnauthorizedError) {
    // render sign-in page
  }
  
  if (error instanceof WebrpcError) {
    console.log(error.status) // print response HTTP status code (ie. 4xx or 5xx)
    console.log(error.code) // print unique schema error code; generic endpoint errors are 0
    console.log(error.message) // print error message
    console.log(error.cause) // print the underlying backend error -- ie. "DB error" - useful for debugging / reporting to Sentry
  }
  
  // setError(error.message)
}
```

## golang@v0.11.0 breaking changes

Note: You can turn on `-legacyErrors=true` flag on golang generator (ie. `webrpc-gen -target=golang -legacyErrors=true -pkg=proto`) in order to preserve the deprecated functions and sentinel errors (see below). This will allow you to migrate your codebase to the new custom schema errors gradually.

The following werbrpc error functions and sentinel errors are now deprecated or ~removed~:
- `proto.WrapError() // Deprecated.`
- `proto.Errorf() // Deprecated.`
- ~`proto.HTTPStatusFromErrorCode()`~
- ~`proto.IsErrorCode()`~
- `proto.ErrCanceled // Deprecated.`
- `proto.ErrUnknown // Deprecated.`
- `proto.ErrFail // Deprecated.`
- `proto.ErrInvalidArgument // Deprecated.`
- `proto.ErrDeadlineExceeded // Deprecated.`
- `proto.ErrNotFound // Deprecated.`
- `proto.ErrBadRoute // Deprecated.`
- `proto.ErrAlreadyExists // Deprecated.`
- `proto.ErrPermissionDenied // Deprecated.`
- `proto.ErrUnauthenticated // Deprecated.`
- `proto.ErrResourceExhausted // Deprecated.`
- `proto.ErrFailedPrecondition // Deprecated.`
- `proto.ErrAborted // Deprecated.`
- `proto.ErrOutOfRange // Deprecated.`
- `proto.ErrUnimplemented // Deprecated.`
- `proto.ErrInternal // Deprecated.`
- `proto.ErrUnavailable // Deprecated.`
- `proto.ErrDataLoss // Deprecated.`
- `proto.ErrNone // Deprecated.`

The schema errors can now be returned from the RPC endpoints via:

```diff
func (s *RPC) RemoveUser(ctx context.Context, userID int64) (bool, error) {
  r, _ := ctx.Value(proto.HTTPRequestCtxKey).(*http.Request)
  if s.IsRateLimited(r) {
-    return false, proto.Errorf(proto.ErrUnavailable, "rate limited")
+    return false, proto.ErrRateLimited // HTTP 429 per RIDL schema
  }

  _, err := s.DB.RemoveUser(ctx, userID)
  if err != nil {
    if errors.Is(err, pgx.ErrNoRows) {
-		  return false, proto.Errorf(proto.ErrNotFound, "no such user(%v)", userID)
+     return false, proto.ErrUserNotFound
    }
-		return false, proto.WrapError(proto.ErrInternal, err, "")
+		return false, proto.ErrorWithCause(proto.ErrDatabaseDown, err)
 	}
 
  return true, nil
}
```

You can also return any other Go error and webrpc will render generic `proto.ErrWebrpcEndpoint` error automatically along with `HTTP 400` status code.
```go
return fmt.Errorf("some error")
```

The RPC client(s) can now assert the schema error type by their unique error code:

```go
if err, ok := rpc.RemoveUser(ctx, userID); err != nil {
	if errors.Is(proto.ErrRateLimited) {
		// slow down; retry with back-off strategy
	}
	if errors.Is(proto.ErrUserNotFound) {
		// handle 
	}
	// etc.
}
```

# webrpc v0.10.0

## Interoperability tests

We have defined a new interoperability test suite implementing the following schema:

```
webrpc = v1

name = Test
version = v0.10.0

service TestApi
  - GetEmpty()
  - GetError()
  
  - GetOne() => (one: Simple)
  - SendOne(one: Simple)

  - GetMulti() => (one: Simple, two: Simple, three: Simple)
  - SendMulti(one: Simple, two: Simple, three: Simple)
  
  - GetComplex() => (complex: Complex)
  - SendComplex(complex: Complex)
```

All generators are expected to implement [TestApi schema](./tests/schema/test.ridl) and run client/server interoperability tests against a reference [webrpc-test binaries)](https://github.com/webrpc/webrpc/releases).

For more info, see [typescript](https://github.com/webrpc/gen-typescript/tree/master/tests) or [golang](https://github.com/webrpc/gen-golang/tree/master/tests) tests.

## Breaking changes in webrpc-gen Go API

```diff
-func NewParser(r *schema.Reader) *Parser
+func NewParser(fsys fs.FS, path string) *Parser
```

```diff
- func NewTemplateSource(proto *schema.WebRPCSchema, target string, config *Config) (*TemplateSource, error)
+ func NewTemplateSource(target string, config *Config) (*TemplateSource, error)
```

# webrpc v0.9.0

Towards reaching webrpc@v1.0.0, we have decided to make some breaking changes to webrpc schema and RIDL file format.

## Breaking changes

### RIDL v0.9.0 changes

Keyword "message" was renamed to "struct".

```diff
 webrpc = v1
 
 name = your-app
 version = v0.1.0
 
- message User
+ struct User
   - id: uint64
   - username: string
   - createdAt?: timestamp
```

### JSON schema v0.9.0 changes

- Field "messages" was renamed to "types"
- Field "type" was renamed to "kind"
- Enum type was moved from enum.fields[] to enum object

```diff
 {
   "webrpc": "v1",
   "name": "Test",
   "version": "v0.0.1",
   "imports": [],
-  "messages": [
+  "types": [
   {
     "name": "Status",
-    "type": "enum",
+    "kind": "enum",
+    "type": "uint32",
     "fields": [
     {
       "name": "AVAILABLE",
-      "type": "uint32",
-      "optional": false,
       "value": "0",
-      "meta": null
     },
     {
       "name": "NOT_AVAILABLE",
-      "type": "uint32",
-      "optional": false,
       "value": "1",
-      "meta": null
     }
   },
   {
    "name": "Empty",
-   "type": "struct",
+   "kind": "struct",
-   "fields": null
   }]
 }
```

### Template changes

You might see the following error when running your webrpc generator templates against webrpc-gen@v0.9.0+:

```
template: main.go.tmpl:88:57: executing "main" at <.Messages>: can't evaluate field Messages in type struct { *schema.WebRPCSchema; SchemaHash string; WebrpcGenVersion string; WebrpcGenCommand string; WebrpcTarget string; Opts map[string]interface {} }
```

To fix this, rename `{{.Messages}}` variable to `{{.Types}}` in your `*.go.tmpl` template files.

## Migration guide

### RIDL v0.9.0 migration guide

Run this command to migrate your RIDL files to webrpc@v0.9.0+:

```bash
#!/bin/bash

find . -name '*.ridl' -exec sed -i -e 's/^message /struct /g' {} \;
```

### JSON schema v0.9.0 migration guide

Run this Node.js script to migrate your `*webrpc.json` schema files to webrpc@v0.9.0+:

`node migrate.js schema.webrpc.json`

Contents of `migrate.js` file:
```javascript
const fs = require("fs");

if (process.argv.length != 3) {
  throw Error(`Usage: node ${process.argv[1]} <webrpc-schema.json>`);
}

const filePath = process.argv[2];

console.log(filePath);

fs.readFile(filePath, "utf8", (e, data) => {
  if (e) {
    throw e;
  }

  let schema = JSON.parse(data);
  schema = {
    webrpc: schema.webrpc,
    name: schema.name,
    version: schema.version,
    types: schema.messages.map((orig) => {
      let type = {
        name: orig.name,
        kind: orig.type,
        fields: orig.fields,
      };

      if (type.kind == "enum") {
        type = {
          name: orig.name,
          kind: orig.type,
          type: orig.fields[0].type,
          fields: orig.fields.map((field) => {
            return { name: field.name, value: field.value };
          }),
        };
      }
      return type;
    }),
    services: schema.services,
  };

  schema.types = fs.writeFile(
    filePath,
    JSON.stringify(schema, null, "\t"),
    (err) => {
      if (err) {
        console.error(err);
        console.log(schema);
      }
    }
  );
});
```

### Generator templates v0.9.0 migration guide

Run this command to migrate your `.go.tmpl` templates to webrpc@v0.9.0+:

```bash
#!/bin/bash

find . -name '*.go.tmpl' -exec sed -i -e 's/\.Messages/.Types/g' {} \;
find . -name '*.go.tmpl' -exec sed -i -e 's/"Messages"/"Types"/g' {} \;
```

