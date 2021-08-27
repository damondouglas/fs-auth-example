package service

import (
	"context"
	"fs-auth-example/backend/internal/environment"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

func ListenTCP(ctx context.Context, env *environment.Environment, svc *grpc.Server) error {
	errChan := make(chan error)
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()
	lis, err := net.Listen("tcp", env.TcpAddress())
	if err != nil {
		return err
	}
	go func(ctx context.Context, errChan chan error) {
		logger.Printf("Listening on TCP at %s", env.TcpAddress())
		if err := svc.Serve(lis); err != nil {
			errChan <- err
			return
		}
	}(ctx, errChan)

	for {
		select {
		case err := <-errChan:
			return err
		case <-ctx.Done():
			logger.Println("interupt signal received, terminating TCP server...")
			return nil
		}
	}
}
