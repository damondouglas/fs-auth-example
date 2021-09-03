package main

import (
	"context"
	"fs-auth-example/backend/internal/environment"
	"google.golang.org/grpc"
	"log"
	"net"
)

// listenTcp binds the grpc.Server to the TCP address provided by the Environment
func listenTcp(ctx context.Context, errChan chan error, env *environment.Environment, svc *grpc.Server) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	address := env.TcpAddress()
	log.Printf("listening TCP at %s\n", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		errChan <- err
		return
	}
	if err := svc.Serve(lis); err != nil {
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
