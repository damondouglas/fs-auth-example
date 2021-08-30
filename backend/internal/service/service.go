package service

import (
	"context"
	"fs-auth-example/backend/internal/auth"
	counter "fs-auth-example/backend/internal/counter/v1"
	"fs-auth-example/backend/internal/environment"
	"fs-auth-example/backend/internal/fs"
	"log"
	"os"

	"google.golang.org/grpc"
)

var (
	logger = log.New(os.Stdout, "SERVICE: ", log.LstdFlags)
)

type service struct {
	env      *environment.Environment
	fsClient *fs.Client
}

func FromEnvironment(ctx context.Context, env *environment.Environment, opts ...grpc.ServerOption) (*grpc.Server, error) {
	if env.IsVerbose() {
		logger.Printf("environment: %s\n", env.String())
	}
	opts = append(opts, auth.WithUnaryAuthorization(env), auth.WithStreamAuthorization(env))
	grpcServer := grpc.NewServer(opts...)
	fsClient, err := fs.FromEnvironment(ctx, env)
	if err != nil {
		logger.Fatalf("fatal: could not instantiate firestore client from environment: %s; err: %s\n", env.String(), err)
		return nil, err
	}
	counter.RegisterCounterServiceServer(grpcServer, &service{
		fsClient: fsClient,
		env:      env,
	})
	return grpcServer, nil
}
