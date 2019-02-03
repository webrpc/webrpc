package webrpc

import (
	"context"
	"net/http"
)

var (
	HTTPResponseWriterCtxKey = &contextKey{"ResponseWriter"}

	ServiceNameCtxKey = &contextKey{"ServiceName"}

	MethodNameCtxKey = &contextKey{"MethodName"}
)

func WithResponseWriter(ctx context.Context, w http.ResponseWriter) context.Context {
	return context.WithValue(ctx, HTTPResponseWriterCtxKey, w)
}

func WithServiceName(ctx context.Context, serviceName string) context.Context {
	return context.WithValue(ctx, ServiceNameCtxKey, serviceName)
}

func WithMethodName(ctx context.Context, methodName string) context.Context {
	return context.WithValue(ctx, MethodNameCtxKey, methodName)
}

// contextKey is a value for use with context.WithValue. It's used as
// a pointer so it fits in an interface{} without allocation. This technique
// for defining context keys was copied from Go 1.7's new use of context in net/http.
type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "webrpc context value " + k.name
}
