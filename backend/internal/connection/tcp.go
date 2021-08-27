package connection

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

func (conn *Connection) TCP(ctx context.Context, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if conn.env.TcpAddress() == "" {
		return nil, fmt.Errorf("fatal: could not derive TCP connection from empty address; environment: %s", conn.env.String())
	}
	if conn.env.IsLocal() {
		opts = append(opts, grpc.WithInsecure())
	}
	return grpc.DialContext(ctx, conn.env.TcpAddress(), opts...)
}
