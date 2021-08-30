package auth

import (
	"fs-auth-example/backend/internal/environment"

	"google.golang.org/grpc"
)

type authorizer struct {
	env *environment.Environment
}

func WithUnaryAuthorization(env *environment.Environment) grpc.ServerOption {
	auth := &authorizer{
		env: env,
	}
	return grpc.UnaryInterceptor(auth.validateUnary)
}

func WithStreamAuthorization(env *environment.Environment) grpc.ServerOption {
	auth := &authorizer{
		env: env,
	}
	return grpc.StreamInterceptor(auth.validateStream)
}
