package environment

import (
	"crypto/tls"
	"crypto/x509"
	"fs-auth-example/backend/internal/gcloud"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

func (env *Environment) ApplyDialOptions(opts *[]grpc.DialOption) error {
	if env.IsLocal() {
		*opts = append(*opts, grpc.WithInsecure())
		return nil
	}

	tok, err := gcloud.Auth.IdentityToken()
	if err != nil {
		return err
	}
	creds, err := oauth.NewJWTAccessFromKey([]byte(tok))
	if err != nil {
		return err
	}
	*opts = append(*opts, grpc.WithPerRPCCredentials(creds))

	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		return err
	}

	tlsCreds := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})
	*opts = append(*opts, grpc.WithTransportCredentials(tlsCreds))

	return nil
}
