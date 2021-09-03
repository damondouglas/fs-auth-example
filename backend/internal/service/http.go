package service

import (
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"net/http"
)

// Wrap a grpc.Server as an http.Handler
func Wrap(svc *grpc.Server) http.Handler {
	wrapped := grpcweb.WrapServer(svc)
	// TODO: determine minimal necessary CORS allowance
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
