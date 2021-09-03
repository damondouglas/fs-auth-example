package auth

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"fs-auth-example/backend/internal/environment"
	"log"
	"os"

	"google.golang.org/grpc"
)

var (
	logger = log.New(os.Stderr, "AUTH: ", log.LstdFlags)
)

type Authorizer struct {
	internal *auth.Client
	env *environment.Environment
}

func NewAuthorizer(ctx context.Context, env *environment.Environment) (*Authorizer, error) {
	project, err := env.Project()
	if err != nil {
		return nil, err
	}
	fb, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: project,
	})
	if err != nil {
		return nil, err
	}
	svc, err := fb.Auth(ctx)
	if err != nil {
		return nil, err
	}
	return &Authorizer{
		internal: svc,
		env: env,
	}, nil
}

func (auth *Authorizer) WithUnaryAuthorization() grpc.ServerOption {
	return grpc.UnaryInterceptor(auth.validateUnary)
}

func (auth *Authorizer) WithStreamAuthorization() grpc.ServerOption {
	return grpc.StreamInterceptor(auth.validateStream)
}
