package service

import (
	"context"
	"fs-auth-example/backend/internal/environment"
	"net/http"
	"os"
	"os/signal"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/rs/cors"
	"google.golang.org/grpc"
)

func Wrap(svc *grpc.Server) http.Handler {
	wrapped := grpcweb.WrapServer(svc)
	return cors.AllowAll().Handler(&wrapper{
		svc: wrapped,
	})
}

type wrapper struct {
	svc *grpcweb.WrappedGrpcServer
}

func (wrap *wrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	wrap.svc.ServeHTTP(w, r)
}

func ListenHTTP(ctx context.Context, env *environment.Environment, svc *grpc.Server) error {
	errChan := make(chan error)
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	handler := Wrap(svc)

	go func(ctx context.Context, errChan chan error) {
		logger.Printf("Listening on HTTP at %s", env.HttpAddress())
		if err := http.ListenAndServe(env.HttpAddress(), handler); err != nil {
			errChan <- err
			return
		}
	}(ctx, errChan)

	for {
		select {
		case err := <-errChan:
			return err
		case <-ctx.Done():
			logger.Println("interupt signal received, terminating HTTP server...")
			return nil
		}
	}
}
