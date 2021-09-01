package fs

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	counter "fs-auth-example/backend/internal/counter/v1"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
	"time"
)

const (
	eqOp     = "=="
	countKey = "Count"
	nameKey  = "Name"
)

type Counts struct {
	firestore.CollectionRef
	client *firestore.Client
}

type Count struct {
	firestore.DocumentSnapshot
	client *firestore.Client
}

func (counts *Counts) List(ctx context.Context, request *counter.ListCountsRequest) ([]*counter.Count, error) {
	var result []*counter.Count
	var n int64 = -1
	if request.Limit > 0 {
		n = request.Limit
	}
	itr := counts.Documents(ctx)
	for {
		var v *counter.Count
		snap, err := itr.Next()
		if err == iterator.Done {
			return result, nil
		}
		if err != nil {
			return nil, err
		}
		if err = snap.DataTo(&v); err != nil {
			return nil, err
		}
		result = append(result, v)
		n--
		if n == 0 {
			return result, nil
		}
	}
}

func (counts *Counts) snapshots(ctx context.Context, countChan chan *Count, errChan chan error) {
	itr := counts.Snapshots(ctx)
	for {
		snap, err := itr.Next()
		if status.Code(err) == codes.DeadlineExceeded {
			return
		}
		if err != nil {
			errChan <- err
		}
		if snap == nil {
			continue
		}
		for _, k := range snap.Changes {
			countChan <- &Count{
				DocumentSnapshot: *k.Doc,
				client:           counts.client,
			}
		}
	}
}

func (counts *Counts) Stream(ctx context.Context, countChan chan *Count, errChan chan error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	tick := time.NewTicker(time.Second)
	for {
		select {
		case <- tick.C:
			counts.snapshots(ctx, countChan, errChan)
		case <- ctx.Done():
			return
		}
	}
}

func (counts *Counts) get(ctx context.Context, name string) (*Count, error) {
	itr := counts.Where(nameKey, eqOp, name).Documents(ctx)
	snaps, err := itr.GetAll()
	if err != nil {
		return nil, err
	}
	if len(snaps) == 0 {
		return nil, nil
	}
	if len(snaps) > 1 {
		var paths []string
		for _, k := range snaps {
			paths = append(paths, k.Ref.Path)
		}
		return nil, fmt.Errorf("fatal: query of %s/{%s=%s} yielded >1 document at: %s", countsPath, nameKey, name, strings.Join(paths, ", "))
	}
	return &Count{
		DocumentSnapshot: *snaps[0],
		client:           counts.client,
	}, nil
}

func (counts *Counts) create(ctx context.Context, name string) (*Count, error) {
	colRef := counts.CollectionRef
	docRef, _, err := colRef.Add(ctx, &counter.Count{
		Name:  name,
		Count: 0,
	})
	if err != nil {
		return nil, err
	}

	snap, err := docRef.Get(ctx)
	if err != nil {
		return nil, err
	}
	return &Count{
		DocumentSnapshot: *snap,
		client:           counts.client,
	}, nil
}

func (counts *Counts) GetOrCreate(ctx context.Context, name string) (*Count, error) {
	count, err := counts.get(ctx, name)
	if err != nil {
		return nil, err
	}
	if count != nil {
		return count, nil
	}
	return counts.create(ctx, name)
}

func (count *Count) Data() (*counter.Count, error) {
	var result *counter.Count
	snap := count.DocumentSnapshot
	if err := snap.DataTo(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func (count *Count) Update(ctx context.Context, value int64) error {
	return count.client.RunTransaction(ctx, func(ctx context.Context, transaction *firestore.Transaction) error {
		if err := transaction.Update(count.Ref, []firestore.Update{
			{
				Path:  countKey,
				Value: value,
			},
		}); err != nil {
			return err
		}
		return nil
	})
}
