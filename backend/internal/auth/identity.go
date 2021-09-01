package auth

import (
	"context"
	"golang.org/x/oauth2"
	"google.golang.org/api/identitytoolkit/v3"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/oauth"
)

type Identity struct {
	internal *identitytoolkit.Service
}

func NewIdentity(ctx context.Context, opts ...option.ClientOption) (*Identity, error) {
	svc, err := identitytoolkit.NewService(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return &Identity{
		internal: svc,
	}, nil
}

func (identity *Identity) WithPasswordSignIn(ctx context.Context, email string, password string) (grpc.DialOption, error) {
	tok, err := identity.token(ctx, email, password)
	if err != nil {
		return nil, err
	}
	perRPC := oauth.NewOauthAccess(tok)
	return grpc.WithPerRPCCredentials(perRPC), nil
}

func (identity *Identity) token(ctx context.Context, email string, password string) (*oauth2.Token, error) {
	pwdId := &passwordIdentity{
		internal: identity.internal,
	}
	req := pwdId.request(ctx, email, password)
	resp, err := req.Do()
	if err != nil {
		return nil, err
	}
	return &oauth2.Token{
		AccessToken: resp.IdToken,
	}, nil
}

