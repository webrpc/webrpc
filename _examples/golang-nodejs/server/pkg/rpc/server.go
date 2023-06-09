package rpc

import (
	"context"
)

type ExampleServiceRPC struct {
}

func (s *ExampleServiceRPC) Ping(ctx context.Context) (bool, error) {
	return true, nil
}

func (s *ExampleServiceRPC) GetUser(ctx context.Context, req *GetUserRequest) (*User, error) {
	if req.UserID == 911 {
		return nil, ErrUserNotFound
	}

	return &User{
		ID:       req.UserID,
		Username: "hihi",
	}, nil
}
