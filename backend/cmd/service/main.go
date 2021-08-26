package main

import (
	"context"
	"fmt"
	"fs-auth-example/backend/internal/environment"
	"log"
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
}
