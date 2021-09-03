package testdata

import (
	"fs-auth-example/backend/internal/gcloud"
)

func IdToken() (string, error) {
	idToken, err := gcloud.Auth.IdentityToken()
	if err != nil {
		return "", err
	}
	return idToken, nil
}
