package service

import (
	"context"
	"fmt"
	counter "fs-auth-example/backend/internal/counter/v1"
	"fs-auth-example/backend/internal/fs"
)

// ListCounts list counts in the counter collection
func (svc *service) ListCounts(ctx context.Context, req *counter.ListCountsRequest) (*counter.ListCountsResponse, error) {
	items, err := svc.fsClient.Counts.List(ctx, req)
	if err != nil {
		logger.Println(fmt.Errorf("ListCounts: %w", err))
		return nil, err
	}

	return &counter.ListCountsResponse{
		Items: items,
	}, nil
}

// UpdateCount updates a counter document count for a user associated from
// a validated id token.  The underlying counter document is created if it doesn't already exist
func (svc *service) UpdateCount(ctx context.Context, req *counter.UpdateCountRequest) (*counter.UpdateCountResponse, error) {
	accountInfo := svc.authorizer.AccountInfoFromContext(ctx)
	if accountInfo == nil {
		err := fmt.Errorf("UpdateCount: account info empty from context")
		logger.Println(err)
		return nil, err
	}
	name := accountInfo.Email()
	if name == "" {
		err := fmt.Errorf("UpdateCount: account info Email empty from context")
		logger.Println(err)
		return nil, err
	}
	ref, err := svc.fsClient.Counts.GetOrCreate(ctx, name)
	if err != nil {
		err = fmt.Errorf("UpdateCount: %w", err)
		logger.Println(err)
		return nil, err
	}

	if err := ref.Update(ctx, req.Count); err != nil {
		err = fmt.Errorf("UpdateCount: %w", err)
		logger.Println(err)
		return nil, err
	}

	return &counter.UpdateCountResponse{
		Count: req.Count,
	}, nil
}

// StreamCounts provides realtime udpates to the underlying counts collection
func (svc *service) StreamCounts(req *counter.StreamCountsRequest, server counter.CounterService_StreamCountsServer) error {
	ctx, cancel := context.WithCancel(server.Context())
	countChan := make(chan *fs.Count)
	errChan := make(chan error)
	go svc.fsClient.Counts.Stream(ctx, countChan, errChan)
	defer cancel()
	for {
		select {
		case err := <- errChan:
			return err
		case ref := <- countChan:
			data, err := ref.Data()
			if err != nil {
				return err
			}
			if err = server.Send(data); err != nil {
				return err
			}
		case <-ctx.Done():
			return nil
		}
	}
}
