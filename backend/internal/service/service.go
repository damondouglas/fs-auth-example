package service

import (
	"context"
	"fs-auth-example/backend/internal/auth"
	counter "fs-auth-example/backend/internal/counter/v1"
	"fs-auth-example/backend/internal/environment"
	"fs-auth-example/backend/internal/fs"
	"google.golang.org/grpc"
	"log"
	"os"
)

var (
	logger = log.New(os.Stdout, "SERVICE: ", log.LstdFlags)
)

type service struct {
	authorizer *auth.Authorizer
	fsClient *fs.Client
	env      *environment.Environment
}

func FromEnvironment(ctx context.Context, env *environment.Environment, opts ...grpc.ServerOption) (*grpc.Server, error) {
	authorizer, err := auth.NewAuthorizer(ctx, env)
	if err != nil {
		return nil, err
	}
	fsClient, err := fs.FromEnvironment(ctx, env)
	if err != nil {
		return nil, err
	}
	grpcServer := grpc.NewServer(opts...)
	counter.RegisterCounterServiceServer(grpcServer, &service{
		authorizer: authorizer,
		env:      env,
		fsClient: fsClient,
	})
	return grpcServer, nil
}
