package service

import (
	"context"
	counter "fs-auth-example/backend/internal/counter/v1"
	"fs-auth-example/backend/internal/environment"
	"log"
	"os"

	"google.golang.org/grpc"
)

var (
	logger = log.New(os.Stdout, "SERVICE: ", log.LstdFlags)
)

type service struct {
	counter.UnimplementedCounterServiceServer
}

func FromEnvironment(ctx context.Context, env *environment.Environment, opts ...grpc.ServerOption) *grpc.Server {
	grpcServer := grpc.NewServer(opts...)
	counter.RegisterCounterServiceServer(grpcServer, &service{})
	return grpcServer
}
