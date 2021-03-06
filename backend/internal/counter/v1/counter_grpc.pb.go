// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CounterServiceClient is the client API for CounterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CounterServiceClient interface {
	ListCounts(ctx context.Context, in *ListCountsRequest, opts ...grpc.CallOption) (*ListCountsResponse, error)
	UpdateCount(ctx context.Context, in *UpdateCountRequest, opts ...grpc.CallOption) (*UpdateCountResponse, error)
	StreamCounts(ctx context.Context, in *StreamCountsRequest, opts ...grpc.CallOption) (CounterService_StreamCountsClient, error)
}

type counterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCounterServiceClient(cc grpc.ClientConnInterface) CounterServiceClient {
	return &counterServiceClient{cc}
}

func (c *counterServiceClient) ListCounts(ctx context.Context, in *ListCountsRequest, opts ...grpc.CallOption) (*ListCountsResponse, error) {
	out := new(ListCountsResponse)
	err := c.cc.Invoke(ctx, "/counter.v1.CounterService/ListCounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterServiceClient) UpdateCount(ctx context.Context, in *UpdateCountRequest, opts ...grpc.CallOption) (*UpdateCountResponse, error) {
	out := new(UpdateCountResponse)
	err := c.cc.Invoke(ctx, "/counter.v1.CounterService/UpdateCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterServiceClient) StreamCounts(ctx context.Context, in *StreamCountsRequest, opts ...grpc.CallOption) (CounterService_StreamCountsClient, error) {
	stream, err := c.cc.NewStream(ctx, &CounterService_ServiceDesc.Streams[0], "/counter.v1.CounterService/StreamCounts", opts...)
	if err != nil {
		return nil, err
	}
	x := &counterServiceStreamCountsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CounterService_StreamCountsClient interface {
	Recv() (*Count, error)
	grpc.ClientStream
}

type counterServiceStreamCountsClient struct {
	grpc.ClientStream
}

func (x *counterServiceStreamCountsClient) Recv() (*Count, error) {
	m := new(Count)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CounterServiceServer is the server API for CounterService service.
// All implementations should embed UnimplementedCounterServiceServer
// for forward compatibility
type CounterServiceServer interface {
	ListCounts(context.Context, *ListCountsRequest) (*ListCountsResponse, error)
	UpdateCount(context.Context, *UpdateCountRequest) (*UpdateCountResponse, error)
	StreamCounts(*StreamCountsRequest, CounterService_StreamCountsServer) error
}

// UnimplementedCounterServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCounterServiceServer struct {
}

func (UnimplementedCounterServiceServer) ListCounts(context.Context, *ListCountsRequest) (*ListCountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCounts not implemented")
}
func (UnimplementedCounterServiceServer) UpdateCount(context.Context, *UpdateCountRequest) (*UpdateCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCount not implemented")
}
func (UnimplementedCounterServiceServer) StreamCounts(*StreamCountsRequest, CounterService_StreamCountsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamCounts not implemented")
}

// UnsafeCounterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CounterServiceServer will
// result in compilation errors.
type UnsafeCounterServiceServer interface {
	mustEmbedUnimplementedCounterServiceServer()
}

func RegisterCounterServiceServer(s grpc.ServiceRegistrar, srv CounterServiceServer) {
	s.RegisterService(&CounterService_ServiceDesc, srv)
}

func _CounterService_ListCounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).ListCounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/counter.v1.CounterService/ListCounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).ListCounts(ctx, req.(*ListCountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CounterService_UpdateCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).UpdateCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/counter.v1.CounterService/UpdateCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).UpdateCount(ctx, req.(*UpdateCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CounterService_StreamCounts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamCountsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CounterServiceServer).StreamCounts(m, &counterServiceStreamCountsServer{stream})
}

type CounterService_StreamCountsServer interface {
	Send(*Count) error
	grpc.ServerStream
}

type counterServiceStreamCountsServer struct {
	grpc.ServerStream
}

func (x *counterServiceStreamCountsServer) Send(m *Count) error {
	return x.ServerStream.SendMsg(m)
}

// CounterService_ServiceDesc is the grpc.ServiceDesc for CounterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CounterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "counter.v1.CounterService",
	HandlerType: (*CounterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCounts",
			Handler:    _CounterService_ListCounts_Handler,
		},
		{
			MethodName: "UpdateCount",
			Handler:    _CounterService_UpdateCount_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamCounts",
			Handler:       _CounterService_StreamCounts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "counter/v1/counter.proto",
}
