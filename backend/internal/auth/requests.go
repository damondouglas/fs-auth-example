package auth

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
)

func (auth *Authorizer) validateUnary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	c, err := auth.validate(ctx)
	if err != nil {
		logger.Println(fmt.Errorf("validateUnary: %w", err))
		return nil, err
	}
	ctx = context.WithValue(ctx, ClaimsKey, c)
	return handler(ctx, req)
}

func (auth *Authorizer) validateStream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	ctx := stream.Context()
	c, err := auth.validate(ctx)
	if err != nil {
		logger.Println(fmt.Errorf("validateStream: %w", err))
		return err
	}
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.New(map[string]string{})
	}
	ctx = context.WithValue(ctx, ClaimsKey, c)
	ctx = metadata.NewOutgoingContext(ctx, md)
	md, _ = metadata.FromOutgoingContext(ctx)
	if err := stream.SetHeader(md); err != nil {
		logger.Println(fmt.Errorf("validateStream: %w", err))
		return err
	}

	return handler(srv, stream)
}

func (auth *Authorizer) validate(ctx context.Context) (*AccountInfo, error) {
	tok, err := auth.TokenFromContext(ctx)
	if err != nil {
		logger.Println(fmt.Errorf("validate: %w", err))
		return nil, err
	}
	return auth.claims(tok)
}
