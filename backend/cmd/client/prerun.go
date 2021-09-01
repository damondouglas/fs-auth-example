package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fs-auth-example/backend/internal/auth"
	counter "fs-auth-example/backend/internal/counter/v1"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/examples/data"
	"strings"
)

func preRun(c *cobra.Command, _ []string) error {
	ctx := c.Context()
	if err := applyOpts(ctx); err != nil {
		return err
	}
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		return err
	}
	counterClient = counter.NewCounterServiceClient(conn)
	return nil
}

func applyOpts(ctx context.Context) error {
	if err := applyCredOpt(ctx); err != nil {
		return err
	}
	if err := applyTLSOpt(); err != nil {
		return err
	}
	return nil
}

func applyCredOpt(ctx context.Context) error {
	email, password, err := promptCredentials()
	if err != nil {
		return err
	}
	id, err := auth.NewIdentity(ctx)
	if err != nil {
		return err
	}
	opt, err := id.WithPasswordSignIn(ctx, email, password)
	if err != nil {
		return err
	}
	opts = append(opts, opt)
	return nil
}

func promptCredentials() (email string, password string, err error) {
	email = "user1@example.com"
	password = "q1w2e3r4t5"
	//var b []byte
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("email: ")
	//email, err = reader.ReadString('\n')
	//if err != nil {
	//	return
	//}
	//email = strings.TrimSpace(email)
	//
	//fmt.Print("password: ")
	//b, err = terminal.ReadPassword(0)
	//if err != nil {
	//	return
	//}
	//b = bytes.TrimSpace(b)
	//password = string(b)

	return
}

func applyTLSOpt() error {
	if strings.Contains(endpoint, "localhost") {
		creds, err := credentials.NewClientTLSFromFile(data.Path("x509/ca_cert.pem"), "*.test.example.com")
		if err != nil {
			return err
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
		return nil
	}
	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		return err
	}
	creds := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})
	opts = append(opts, grpc.WithTransportCredentials(creds))

	return nil
}
