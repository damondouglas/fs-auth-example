package auth

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"google.golang.org/grpc/metadata"
	"strings"
)

const (
	authorizationKey = "authorization"
	bearerPrefix     = "Bearer "
	ClaimsKey        = "AccountInfo"
)

// TokenFromContext provides, if available, oauth2.Token stored in a context.Context
func (auth *Authorizer) TokenFromContext(ctx context.Context) (*oauth2.Token, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata from context")
	}
	return auth.tokenFromMetadata(md)
}

// tokenFromMetadata provides, if available, oauth2.Token stored in metadata.MD
func (auth *Authorizer) tokenFromMetadata(md metadata.MD) (*oauth2.Token, error) {
	if md == nil {
		return nil, fmt.Errorf("metadata is nil from context")
	}

	authorization := md.Get(authorizationKey)
	if len(authorization) == 0 {
		return nil, fmt.Errorf("metadata missing %s from context", authorizationKey)
	}

	authData := strings.TrimPrefix(authorization[0], bearerPrefix)
	return &oauth2.Token{
		AccessToken:  authData,
		TokenType:    "bearer",
	}, nil
}
