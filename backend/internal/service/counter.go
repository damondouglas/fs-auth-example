package service

import (
	"context"
	counter "fs-auth-example/backend/internal/counter/v1"
	"math/rand"
	"time"
)

func (svc *service) ListCounts(ctx context.Context, req *counter.ListCountsRequest) (*counter.ListCountsResponse, error) {
	return &counter.ListCountsResponse{}, nil
}

func (svc *service) UpdateCount(ctx context.Context, req *counter.UpdateCountRequest) (*counter.UpdateCountResponse, error) {
	return &counter.UpdateCountResponse{}, nil
}

func (svc *service) StreamCounts(req *counter.StreamCountsRequest, server counter.CounterService_StreamCountsServer) error {
	ctx, cancel := context.WithCancel(server.Context())
	defer cancel()
	tick := time.NewTicker(time.Second)
	for {
		select {
		case <-tick.C:
			if err := server.Send(&counter.Count{
				Name:  "foo",
				Count: rand.Int63(),
			}); err != nil {
				return err
			}

		case <-ctx.Done():
			tick.Stop()
			return nil
		}
	}
}
