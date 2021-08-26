package service

import (
	"context"
	counter "fs-auth-example/backend/internal/counter/v1"
	"fs-auth-example/backend/internal/environment"

	"google.golang.org/grpc"
)

type service struct {
}

func FromEnvironment(ctx context.Context, env *environment.Environment, opts ...grpc.ServerOption) *grpc.Server {
	grpcServer := grpc.NewServer()
	counter.RegisterCounterServiceServer(grpcServer, &service{})
	return grpcServer
}

func (svc *service) Count(ctx context.Context, req *counter.CountRequest) (*counter.CountResponse, error) {
	return &counter.CountResponse{
		Count: 100,
	}, nil
}
