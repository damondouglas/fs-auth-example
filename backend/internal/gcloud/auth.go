package gcloud

var (
	// Auth wraps the gcloud auth command
	Auth = &GcloudAuth{}
)

// GcloudAuth wraps the gcloud auth command
type GcloudAuth struct{}

// IdentityToken provides the gcloud auth print-identity-token command
func (auth *GcloudAuth) IdentityToken() (string, error) {
	return output("auth", "print-identity-token")
}
