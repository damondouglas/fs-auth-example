package auth

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	bearerPrefix = "Bearer "
	claimsKey    = "claims"
)

func (auth *authorizer) validateUnary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	c, err := auth.validate(ctx)
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, claimsKey, c)
	return handler(ctx, req)
}

func (auth *authorizer) validateStream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	ctx := stream.Context()
	c, err := auth.validate(ctx)
	if err != nil {
		return err
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing metadata from context")
	}

	if err != nil {
		return err
	}
	if err := stream.SetHeader(md); err != nil {

	}
	return handler(srv, stream)
}

func (auth *authorizer) validate(ctx context.Context) (*claims, error) {
	tok, err := auth.tokenFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return auth.claims(ctx, tok)
}
