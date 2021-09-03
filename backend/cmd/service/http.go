package main

import (
	"context"
	"fs-auth-example/backend/internal/environment"
	"log"
	"net/http"
)

// listenHttp binds the handler on the TCP network address provided by the Environment
func listenHttp(ctx context.Context, errChan chan error, env *environment.Environment, handler http.Handler) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	address := env.HttpAddress()
	log.Printf("listening HTTP at %s\n", address)
	if err := http.ListenAndServe(address, handler); err != nil {
		errChan <- err
		return
	}
	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}
