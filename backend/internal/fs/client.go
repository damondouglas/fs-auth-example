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

// Client wraps a firestore.Client
type Client struct {
	*firestore.Client
	Counts *Counts
}

// FromEnvironment instantiates a Client from details in the Environment configuration
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
