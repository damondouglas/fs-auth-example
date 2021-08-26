package main

import (
	"context"
	"fmt"
	"fs-auth-example/backend/internal/environment"
	"fs-auth-example/backend/internal/service"
	"log"
	"net"
	"os"
	"os/signal"
)

var (
	env     *environment.Environment
	address string
)

func init() {
	if err := vars(context.Background()); err != nil {
		panic(err)
	}
}

func vars(ctx context.Context) error {
	var err error
	env, err = environment.New(environment.EnvPort)
	if err != nil {
		return err
	}
	address = fmt.Sprintf(":%s", env.Port.Value())
	return nil
}

func main() {
	log.Printf("Listening at %s\n", address)
	ctx := context.Background()
	if err := run(ctx); err != nil {
		panic(err)
	}
}

func run(ctx context.Context) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()
	errChan := make(chan error)
	svc := service.FromEnvironment(ctx, env)

	go func(ctx context.Context, errChan chan error) {
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
			log.Println("interupt signal received; shutting down ...")
			return nil
		}
	}
}
