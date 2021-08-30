package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
	"google.golang.org/grpc/metadata"
)

const (
	authorizationKey = "authorization"
)

func (auth *authorizer) claims(ctx context.Context, tok *oauth2.Token) (*claims, error) {
	var result *claims
	if tok == nil {
		return nil, fmt.Errorf("oauth token empty from context")
	}

	payload, err := idtoken.Validate(ctx, tok.AccessToken, auth.env.ClientId.String())
	if err != nil {
		return nil, err
	}

	expiry := time.Unix(payload.Expires, 0)
	if expiry.Before(time.Now()) {
		return nil, fmt.Errorf("oauth token expired from context")
	}

	b, err := json.Marshal(payload.Claims)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (auth *authorizer) tokenFromContext(ctx context.Context) (*oauth2.Token, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata from context")
	}
	return auth.tokenFromMetadata(md)
}

func (auth *authorizer) tokenFromMetadata(md metadata.MD) (*oauth2.Token, error) {
	if md == nil {
		return nil, fmt.Errorf("metadata is nil from context")
	}

	authorization := md.Get(authorizationKey)
	if len(authorization) == 0 {
		return nil, fmt.Errorf("metadata missing %s from context", authorizationKey)
	}

	authData := strings.TrimPrefix(authorization[0], bearerPrefix)
	return auth.tokenFromJsonKey([]byte(authData))
}

func (auth *authorizer) tokenFromJsonKey(jsonKey []byte) (*oauth2.Token, error) {
	if len(jsonKey) == 0 {
		return nil, fmt.Errorf("jsonKey data received from metadata is empty from context")
	}
	src, err := google.JWTAccessTokenSourceFromJSON(jsonKey, auth.env.ClientId.String())
	if err != nil {
		return nil, err
	}
	if src == nil {
		return nil, fmt.Errorf("authorization token source received from metadata is empty from context")
	}

	return src.Token()
}
