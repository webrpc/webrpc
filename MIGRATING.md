# Migrating between webrpc versions

Breaking changes per release, newest first. Each entry starts with the error you see after regenerating with the new `webrpc-gen`, so you can search this page for it.

## v0.52.0

### Go server: `MethodNameCtxKey` and `ServiceNameCtxKey` removed

**Error**

```
undefined: proto.MethodNameCtxKey
undefined: proto.ServiceNameCtxKey
```

`proto` stands for your generated package.

**Why**

The generated server no longer stores the service and method name in the request context. `MethodCtx(ctx)` resolves them from the request. This saves three allocations per request.

**Reading the name**

Before:

```go
method, _ := ctx.Value(proto.MethodNameCtxKey).(string)
service, _ := ctx.Value(proto.ServiceNameCtxKey).(string)
```

After:

```go
if m, ok := proto.MethodCtx(ctx); ok {
	method, service := m.Name(), m.Service()
}
```

**Tests that injected a name**

Inject the request instead. `MethodCtx` resolves the method from its path.

Before:

```go
ctx = context.WithValue(ctx, proto.MethodNameCtxKey, "Ping")
```

After:

```go
r := httptest.NewRequest("POST", "/rpc/Example/Ping", nil)
ctx = context.WithValue(ctx, proto.HTTPRequestCtxKey, r)
```

Use the path the method is served on, which depends on the schema's basepath. `proto.WebrpcMethods()` lists every path as a map key.

**Also changed**

- `MethodNameFromContext` and `ServiceNameFromContext` still work. They are deprecated in favor of `MethodCtx`.
- On a path that is not a known method, both return `""`. Before, `ServiceNameFromContext` returned the service name.
