package main

import (
	"context"
	"flag"
	"fs-auth-example/backend/internal/environment"
	"fs-auth-example/backend/internal/service"
	"os"
	"sync"
)

var (
	help = flag.Bool("help", false, "Show help")
	env  = environment.DefaultEnvironment
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
	wg := &sync.WaitGroup{}
	svc, err := service.FromEnvironment(ctx, env)
	if err != nil {
		return err
	}
	if env.PortHttp.Exists() {
		wg.Add(1)
		go func(ctx context.Context) {
			service.ListenHTTP(ctx, env, svc)
			wg.Done()
		}(ctx)
	}
	if env.PortTcp.Exists() {
		wg.Add(1)
		go func(ctx context.Context) {
			service.ListenTCP(ctx, env, svc)
			wg.Done()
		}(ctx)
	}

	wg.Wait()
	return nil
}
