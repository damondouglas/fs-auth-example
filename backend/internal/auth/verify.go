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

// AccountInfo models user details provided from id token verification
type AccountInfo auth.Token

// Email provides the user email, if available, from Jwt token verify
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

// AccountInfoFromContext provides, if available, AccountInfo stored in context.Context
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

// verify oauth2.Token and provide AccountInfo associated with id token
func (auth *Authorizer) verify(ctx context.Context, tok *oauth2.Token) (*AccountInfo, error) {
	resp, err := auth.internal.VerifyIDToken(ctx, tok.AccessToken)
	if err != nil {
		logger.Println(fmt.Errorf("verify: %w", err))
		return nil, err
	}
	return (*AccountInfo)(resp), nil
}
