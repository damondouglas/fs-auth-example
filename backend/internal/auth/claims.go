package auth

import (
	"context"
	"golang.org/x/oauth2"
	"google.golang.org/api/identitytoolkit/v3"
)

type AccountInfo identitytoolkit.GetAccountInfoResponse

func (identity *Identity) claims(tok *oauth2.Token) (*AccountInfo, error) {
	resp, err := identity.internal.Relyingparty.GetAccountInfo(&identitytoolkit.IdentitytoolkitRelyingpartyGetAccountInfoRequest{
		IdToken: tok.AccessToken,
	}).Do()
	if err != nil {
		return nil, err
	}

	return (*AccountInfo)(resp), nil
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
