ridl, pronounced "riddle"
=========================

a ridl file is a "rpc interface design language" schema file that
describes a webrpc client/server app.

## Define webrpc schema errors

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
