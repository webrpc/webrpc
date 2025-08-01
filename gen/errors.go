package gen

import "github.com/webrpc/webrpc/schema"

var WebrpcErrors = []*schema.Error{
	{Code: 0, Name: "WebrpcEndpoint", Message: "endpoint error", HTTPStatus: 400},
	{Code: -1, Name: "WebrpcRequestFailed", Message: "request failed", HTTPStatus: 400},
	{Code: -2, Name: "WebrpcBadRoute", Message: "bad route", HTTPStatus: 404},
	{Code: -3, Name: "WebrpcBadMethod", Message: "bad method", HTTPStatus: 405},
	{Code: -4, Name: "WebrpcBadRequest", Message: "bad request", HTTPStatus: 400},
	{Code: -5, Name: "WebrpcBadResponse", Message: "bad response", HTTPStatus: 500},
	{Code: -6, Name: "WebrpcServerPanic", Message: "server panic", HTTPStatus: 500},
	{Code: -7, Name: "WebrpcInternalError", Message: "internal error", HTTPStatus: 500},
	{Code: -8, Name: "WebrpcClientAborted", Message: "request aborted by client", HTTPStatus: 400},
	{Code: -9, Name: "WebrpcStreamLost", Message: "stream lost", HTTPStatus: 400},
	{Code: -10, Name: "WebrpcStreamFinished", Message: "stream finished", HTTPStatus: 200},
	// Note: Do not change existing values. Append only.
	// Keep the list short. Code and Name must be unique.
}
