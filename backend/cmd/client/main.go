package main

import (
	"context"
	"flag"
	"fs-auth-example/backend/internal/connection"
	counter "fs-auth-example/backend/internal/counter/v1"
	"fs-auth-example/backend/internal/environment"
	"log"
	"os"
	"os/signal"
)

var (
	logger = log.New(os.Stdout, "CLIENT: ", log.LstdFlags)
	env    = environment.DefaultEnvironment
	help   = flag.Bool("help", false, "Show help")
)

func vars(ctx context.Context) error {
	if env.IsLocal() {
		if err := environment.LoadLocalEnv(); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	flag.Parse()
	if *help {
		environment.Usage()
		os.Exit(0)
	}
	ctx := context.Background()
	if err := vars(ctx); err != nil {
		panic(err)
	}
	if err := run(ctx); err != nil {
		panic(err)
	}
}

func run(ctx context.Context) error {
	errChan := make(chan error)
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()
	if env.IsVerbose() {
		logger.Println(env.String())
	}

	conn, err := connection.FromEnvironment(env).TCP(ctx)
	if err != nil {
		return err
	}

	client := counter.NewCounterServiceClient(conn)
	resp, err := client.ListCounts(ctx, &counter.ListCountsRequest{})
	if err != nil {
		return err
	}
	for _, item := range resp.Items {
		log.Println(item.String())
	}

	stream, err := client.StreamCounts(ctx, &counter.StreamCountsRequest{})
	if err != nil {
		return err
	}

	go func(ctx context.Context, errChan chan error) {
		for {
			count, err := stream.Recv()
			if err != nil {
				errChan <- err
				return
			}
			log.Println(count.String())
		}
	}(ctx, errChan)

	for {
		select {
		case err := <-errChan:
			return err
		case <-ctx.Done():
			logger.Println("interupt signal received, terminating...")
			return nil
		}
	}
}
