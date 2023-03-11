package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func RunTests(ctx context.Context, serverURL string) error {
	var errs []error // Note: We can't use Go 1.20's errors.Join() until we drop support for older Go versions.

	testApi := NewTestApiClient(serverURL, &http.Client{})

	if err := testApi.GetEmpty(ctx); err != nil {
		errs = append(errs, fmt.Errorf("GetEmpty(): %w", err))
	}

	if err := testApi.GetError(ctx); err == nil {
		errs = append(errs, fmt.Errorf("GetError(): expected error, got nil"))
	}

	one, err := testApi.GetOne(ctx)
	if err != nil {
		errs = append(errs, fmt.Errorf("GetOne(): %w", err))
	}
	if err := testApi.SendOne(ctx, one); err != nil {
		errs = append(errs, fmt.Errorf("SendOne(): %w", err))
	}

	one, two, three, err := testApi.GetMulti(ctx)
	if err != nil {
		errs = append(errs, fmt.Errorf("GetMulti(): %w", err))
	}
	if err := testApi.SendMulti(ctx, one, two, three); err != nil {
		errs = append(errs, fmt.Errorf("SendMulti(): %w", err))
	}

	complex, err := testApi.GetComplex(ctx)
	if err != nil {
		errs = append(errs, fmt.Errorf("GetComplex(): %w", err))
	}
	if err := testApi.SendComplex(ctx, complex); err != nil {
		errs = append(errs, fmt.Errorf("SendComplex(): %w", err))
	}

	schemaErrs := testSchemaErrors(ctx, testApi)
	errs = append(errs, schemaErrs...)

	if len(errs) > 0 {
		var b strings.Builder
		fmt.Fprintf(&b, "Failed tests:\n")
		for _, err := range errs {
			fmt.Fprintf(&b, "%v\n", err)
		}
		return fmt.Errorf(b.String())
	}

	return nil
}

func testSchemaErrors(ctx context.Context, testApi TestApi) []error {
	tt := []struct {
		code           int
		err            WebRPCError
		name           string
		msg            string
		httpStatusCode int
		cause          string
	}{
		{code: 0, err: ErrWebrpcEndpoint, name: "WebrpcEndpoint", msg: "endpoint error", httpStatusCode: 400, cause: "failed to read file: unexpected EOF"},
		{code: 1, err: ErrUnauthorized, name: "Unauthorized", msg: "unauthorized", httpStatusCode: 401, cause: "failed to verify JWT token"},
		{code: 2, err: ErrExpiredToken, name: "ExpiredToken", msg: "expired token", httpStatusCode: 401},
		{code: 3, err: ErrInvalidToken, name: "InvalidToken", msg: "invalid token", httpStatusCode: 401},
		{code: 4, err: ErrDeactivated, name: "Deactivated", msg: "account deactivated", httpStatusCode: 403},
		{code: 5, err: ErrConfirmAccount, name: "ConfirmAccount", msg: "confirm your email", httpStatusCode: 403},
		{code: 6, err: ErrAccessDenied, name: "AccessDenied", msg: "access denied", httpStatusCode: 403},
		{code: 7, err: ErrMissingArgument, name: "MissingArgument", msg: "missing argument", httpStatusCode: 400},
		{code: 8, err: ErrUnexpectedValue, name: "UnexpectedValue", msg: "unexpected value", httpStatusCode: 400},
		{code: 100, err: ErrRateLimited, name: "RateLimited", msg: "too many requests", httpStatusCode: 429, cause: "1000 req/min exceeded"},
		{code: 101, err: ErrDatabaseDown, name: "DatabaseDown", msg: "service outage", httpStatusCode: 503},
		{code: 102, err: ErrElasticDown, name: "ElasticDown", msg: "search is degraded", httpStatusCode: 503},
		{code: 103, err: ErrNotImplemented, name: "NotImplemented", msg: "not implemented", httpStatusCode: 501},
		{code: 200, err: ErrUserNotFound, name: "UserNotFound", msg: "user not found", httpStatusCode: 400},
		{code: 201, err: ErrUserBusy, name: "UserBusy", msg: "user busy", httpStatusCode: 400},
		{code: 202, err: ErrInvalidUsername, name: "InvalidUsername", msg: "invalid username", httpStatusCode: 400},
		{code: 300, err: ErrFileTooBig, name: "FileTooBig", msg: "file is too big (max 1GB)", httpStatusCode: 400},
		{code: 301, err: ErrFileInfected, name: "FileInfected", msg: "file is infected", httpStatusCode: 400},
		{code: 302, err: ErrFileType, name: "FileType", msg: "unsupported file type", httpStatusCode: 400, cause: ".wav is not supported"},
	}

	var errs []error

	for _, tc := range tt {
		err := testApi.GetSchemaError(ctx, tc.code)
		if !errors.Is(err, tc.err) {
			errs = append(errs, fmt.Errorf("unexpected error for code=%v:\nexpected: %#v,\ngot:      %#v", tc.code, tc.err, err))
		}

		rpcErr, _ := err.(WebRPCError)
		if rpcErr.Code != tc.code {
			errs = append(errs, fmt.Errorf("unexpected error code: expected: %v, got: %v", tc.code, rpcErr.Code))
		}
		if rpcErr.Name != tc.name {
			errs = append(errs, fmt.Errorf("unexpected error name: expected: %q, got: %q", tc.name, rpcErr.Name))
		}
		if rpcErr.Message != tc.msg {
			errs = append(errs, fmt.Errorf("unexpected error message: expected: %q, got: %q", tc.msg, rpcErr.Message))
		}
		if rpcErr.HTTPStatus != tc.httpStatusCode {
			errs = append(errs, fmt.Errorf("unexpected error HTTP status code: expected: %v, got: %v", tc.httpStatusCode, rpcErr.HTTPStatus))
		}
		if cause := rpcErr.Unwrap(); cause != nil && cause.Error() != tc.cause {
			errs = append(errs, fmt.Errorf("unexpected error cause: expected %q, got %q", tc.cause, cause.Error()))
		}
	}

	return errs
}
