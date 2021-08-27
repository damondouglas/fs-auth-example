package service

import (
	"net/http"

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
