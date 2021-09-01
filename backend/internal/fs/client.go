package fs

import (
	"context"
	"fs-auth-example/backend/internal/environment"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

const (
	countsPath = "counts"
)

type Client struct {
	*firestore.Client
	Counts *Counts
}

func FromEnvironment(ctx context.Context, env *environment.Environment, opts ...option.ClientOption) (*Client, error) {
	project, err := env.Project()
	if err != nil {
		return nil, err
	}

	client, err := firestore.NewClient(ctx, project, opts...)
	if err != nil {
		return nil, err
	}
	counts := client.Collection(countsPath)
	return &Client{
		Client: client,
		Counts: &Counts{
			CollectionRef: *counts,
			client: client,
		},
	}, nil
}
