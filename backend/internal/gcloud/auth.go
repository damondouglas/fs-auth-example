package gcloud

var (
	Auth = &GcloudAuth{}
)

type GcloudAuth struct{}

func (auth *GcloudAuth) IdentityToken() (string, error) {
	return output("auth", "print-identity-token")
}
