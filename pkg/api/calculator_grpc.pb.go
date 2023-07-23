// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: calculator.proto

package api

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

// CalculatingServiceClient is the client API for CalculatingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatingServiceClient interface {
	Calculate(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponse, error)
}

type calculatingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatingServiceClient(cc grpc.ClientConnInterface) CalculatingServiceClient {
	return &calculatingServiceClient{cc}
}

func (c *calculatingServiceClient) Calculate(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponse, error) {
	out := new(CalculateResponse)
	err := c.cc.Invoke(ctx, "/api.CalculatingService/Calculate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatingServiceServer is the server API for CalculatingService service.
// All implementations must embed UnimplementedCalculatingServiceServer
// for forward compatibility
type CalculatingServiceServer interface {
	Calculate(context.Context, *CalculateRequest) (*CalculateResponse, error)
	mustEmbedUnimplementedCalculatingServiceServer()
}

// UnimplementedCalculatingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalculatingServiceServer struct {
}

func (UnimplementedCalculatingServiceServer) Calculate(context.Context, *CalculateRequest) (*CalculateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calculate not implemented")
}
func (UnimplementedCalculatingServiceServer) mustEmbedUnimplementedCalculatingServiceServer() {}

// UnsafeCalculatingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatingServiceServer will
// result in compilation errors.
type UnsafeCalculatingServiceServer interface {
	mustEmbedUnimplementedCalculatingServiceServer()
}

func RegisterCalculatingServiceServer(s grpc.ServiceRegistrar, srv CalculatingServiceServer) {
	s.RegisterService(&CalculatingService_ServiceDesc, srv)
}

func _CalculatingService_Calculate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatingServiceServer).Calculate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CalculatingService/Calculate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatingServiceServer).Calculate(ctx, req.(*CalculateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CalculatingService_ServiceDesc is the grpc.ServiceDesc for CalculatingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculatingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.CalculatingService",
	HandlerType: (*CalculatingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calculate",
			Handler:    _CalculatingService_Calculate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calculator.proto",
}
