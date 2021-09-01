package main

import (
	"crypto/tls"
	"fs-auth-example/backend/internal/auth"
	"fs-auth-example/backend/internal/environment"
	"fs-auth-example/backend/internal/service"
	"fs-auth-example/backend/testdata"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/examples/data"
	"log"
	"os"
	"os/signal"
)

var (
	rootCmd = &cobra.Command{
		Use:     "service",
		Short:   "Run gRPC backend service",
		PreRunE: preRun,
		RunE:    run,
	}

	local   bool
	verbose bool

	opts []grpc.ServerOption
)

func init() {
	rootCmd.Flags().BoolVar(&local, "local", false, "Run in a local server environment")
	rootCmd.Flags().BoolVar(&verbose, "verbose", false, "Run with verbose logging")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func preRun(c *cobra.Command, _ []string) error {
	if local {
		if err := testdata.LoadDotEnv(); err != nil {
			return err
		}
		cert, err := tls.LoadX509KeyPair(data.Path("x509/server_cert.pem"), data.Path("x509/server_key.pem"))
		if err != nil {
			return err
		}
		opts = append(opts, grpc.Creds(credentials.NewServerTLSFromCert(&cert)))
	}

	env := environment.DefaultEnvironment
	a, err := auth.NewAuthorizer(c.Context(), env)
	if err != nil {
		return err
	}

	opts = append(opts, a.WithStreamAuthorization(), a.WithUnaryAuthorization())

	return nil
}

func run(c *cobra.Command, _ []string) error {
	ctx, cancel := signal.NotifyContext(c.Context(), os.Interrupt)
	defer cancel()
	errChan := make(chan error)
	env := environment.DefaultEnvironment
	if verbose {
		log.Println(env.String())
	}

	svc, err := service.FromEnvironment(ctx, env, opts...)
	if err != nil {
		return err
	}

	go listenTcp(ctx, errChan, env, svc)

	handler := service.Wrap(svc)
	go listenHttp(ctx, errChan, env, handler)

	for {
		select {
		case err := <-errChan:
			return err
		case <-ctx.Done():
			log.Println("interrupt signal received; stopping...")
			return nil
		}
	}
}
