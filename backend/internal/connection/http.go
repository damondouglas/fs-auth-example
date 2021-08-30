package connection

import (
	"context"
	"net/http"

	"google.golang.org/api/option"
)

type Requester interface {
	Do() (*http.Response, error)
}

// TODO: consider HTTP client connection to test gRPC-web features
func HTTPConnection(ctx context.Context, opts ...option.ClientOption) (Requester, error) {
	return nil, nil
}
