package auth

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"fmt"
	"golang.org/x/oauth2"
)

const (
	emailKey = "email"
)

type AccountInfo auth.Token

func (info *AccountInfo) Email() string {
	itr, ok := info.Claims[emailKey]
	if !ok {
		return ""
	}
	if v, ok := itr.(string); ok {
		return v
	}
	return ""
}

func (auth *Authorizer) AccountInfoFromContext(ctx context.Context) *AccountInfo {
	itr := ctx.Value(ClaimsKey)
	if itr == nil {
		return nil
	}
	if c, ok := itr.(*AccountInfo); ok {
		return c
	}
	return nil
}

func (auth *Authorizer) claims(ctx context.Context, tok *oauth2.Token) (*AccountInfo, error) {
	resp, err := auth.internal.VerifyIDToken(ctx, tok.AccessToken)
	if err != nil {
		logger.Println(fmt.Errorf("claims: %w", err))
		return nil, err
	}
	return (*AccountInfo)(resp), nil
}
