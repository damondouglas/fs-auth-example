package main

import (
	"context"
	"flag"
	"fs-auth-example/backend/internal/environment"
	"log"
	"os"
)

var (
	help = flag.Bool("help", false, "Show help")
	env  = environment.DefaultEnvironment
)

func vars(ctx context.Context) error {
	if (env.IsLocal()) {
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
	log.Println(env.String())
	return nil
	// lis, err := net.Listen("tcp", address)
	// if err != nil {
	// 	return err
	// }
	// ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	// defer cancel()
	// errChan := make(chan error)
	// svc := service.FromEnvironment(ctx, env)

	// go func(ctx context.Context, errChan chan error) {
	// 	if err := svc.Serve(lis); err != nil {
	// 		errChan <- err
	// 		return
	// 	}
	// }(ctx, errChan)

	// for {
	// 	select {
	// 	case err := <-errChan:
	// 		return err
	// 	case <-ctx.Done():
	// 		log.Println("interupt signal received; shutting down ...")
	// 		return nil
	// 	}
	// }
}
