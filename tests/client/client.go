package client

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

func RunTests(ctx context.Context, serverURL string) error {
	// Can't use Go 1.20's errors.Join() until we drop support for older Go versions.
	errs := []error{}

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

	if len(errs) > 0 {
		var b strings.Builder
		fmt.Fprintf(&b, "%v test errors:\n", len(errs))
		for _, err := range errs {
			fmt.Fprintf(&b, "- %v\n", err)
		}
		return errors.Errorf(b.String())
	}

	return nil
}
