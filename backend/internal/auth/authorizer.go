package auth

import (
	"context"
	"fs-auth-example/backend/internal/environment"
	"log"
	"os"

	"google.golang.org/grpc"
)

var (
	logger = log.New(os.Stderr, "AUTH: ", log.LstdFlags)
)

type Authorizer struct {
	env *environment.Environment
	id *Identity
}

func NewAuthorizer(ctx context.Context, env *environment.Environment) (*Authorizer, error) {
	id, err := NewIdentity(ctx)
	if err != nil {
		return nil, err
	}
	return &Authorizer{
		env: env,
		id: id,
	}, nil
}

func (auth *Authorizer) WithUnaryAuthorization() grpc.ServerOption {
	return grpc.UnaryInterceptor(auth.validateUnary)
}

func (auth *Authorizer) WithStreamAuthorization() grpc.ServerOption {
	return grpc.StreamInterceptor(auth.validateStream)
}
