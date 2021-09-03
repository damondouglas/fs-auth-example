package service

import (
	"context"
	"fmt"
	counter "fs-auth-example/backend/internal/counter/v1"
	"fs-auth-example/backend/internal/fs"
)

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

func (svc *service) UpdateCount(ctx context.Context, req *counter.UpdateCountRequest) (*counter.UpdateCountResponse, error) {
	accountInfo := svc.authorizer.AccountInfoFromContext(ctx)
	if accountInfo == nil {
		err := fmt.Errorf("UpdateCount: account info empty from context")
		logger.Println(err)
		return nil, err
	}
	//if len(accountInfo.Users) == 0 {
	//	err := fmt.Errorf("UpdateCount: no users found with account info retrieved from context")
	//	logger.Println(err)
	//	return nil, err
	//}
	//user := accountInfo.Users[0]
	//name := user.Email
	name := accountInfo.Email()
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
