package server

import (
	"context"
	"fmt"
)

// Impl echoes the decoded request, so tests can check each binding.
type Impl struct{}

func (Impl) Index(ctx context.Context) error { return nil }

func (Impl) Search(ctx context.Context, req SearchRequest) (*SearchResponse, error) {
	return &SearchResponse{Echo: fmt.Sprintf("key=%s tag=%s limit=%d", req.Key, req.Tag, req.Limit)}, nil
}

func (Impl) Scale(ctx context.Context, factor float32) (float32, error) {
	return factor * 2, nil
}
