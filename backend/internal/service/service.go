package service

import (
	"context"
	"fs-auth-example/backend/internal/environment"

	"google.golang.org/grpc"
)

func FromEnvironment(ctx context.Context, env *environment.Environment) (*grpc.Server, error) {
	return nil, nil
}