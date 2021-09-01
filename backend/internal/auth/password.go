package auth

import (
	"context"
	"google.golang.org/api/identitytoolkit/v3"
)

type passwordIdentity struct {
	internal *identitytoolkit.Service
}

func (identity *passwordIdentity) request(ctx context.Context, email string, password string) *identitytoolkit.RelyingpartyVerifyPasswordCall {
	req := &identitytoolkit.IdentitytoolkitRelyingpartyVerifyPasswordRequest{
		Email:    email,
		Password: password,
	}
	return identity.internal.Relyingparty.VerifyPassword(req).Context(ctx)
}

