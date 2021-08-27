package fs

import (
	"context"
	"fs-auth-example/backend/internal/environment"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type Client struct {
	*firestore.Client
}

func FromEnvironment(ctx context.Context, env *environment.Environment, opts ...option.ClientOption) (*Client, error) {
	project, err := env.Project()
	if err != nil {
		return nil, err
	}

	if err := env.ApplyFirestoreOpts(&opts); err != nil {
		return nil, err
	}

	client, err := firestore.NewClient(ctx, project, opts...)
	if err != nil {
		return nil, err
	}
	return &Client{
		Client: client,
	}, nil
}
