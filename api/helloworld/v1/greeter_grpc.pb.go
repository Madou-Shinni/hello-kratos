// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: helloworld/v1/greeter.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	v1 "helloword/api/stock/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Greeter_SayHello_FullMethodName    = "/helloworld.v1.Greeter/SayHello"
	Greeter_DeductStock_FullMethodName = "/helloworld.v1.Greeter/DeductStock"
	Greeter_AddStock_FullMethodName    = "/helloworld.v1.Greeter/AddStock"
)

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	// 扣库存
	DeductStock(ctx context.Context, in *v1.DeductStockRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 增加库存
	AddStock(ctx context.Context, in *v1.IncreaseStockRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, Greeter_SayHello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) DeductStock(ctx context.Context, in *v1.DeductStockRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Greeter_DeductStock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) AddStock(ctx context.Context, in *v1.IncreaseStockRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Greeter_AddStock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	// 扣库存
	DeductStock(context.Context, *v1.DeductStockRequest) (*emptypb.Empty, error)
	// 增加库存
	AddStock(context.Context, *v1.IncreaseStockRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreeterServer) DeductStock(context.Context, *v1.DeductStockRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeductStock not implemented")
}
func (UnimplementedGreeterServer) AddStock(context.Context, *v1.IncreaseStockRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStock not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_DeductStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.DeductStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).DeductStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_DeductStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).DeductStock(ctx, req.(*v1.DeductStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_AddStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.IncreaseStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).AddStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_AddStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).AddStock(ctx, req.(*v1.IncreaseStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.v1.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
		{
			MethodName: "DeductStock",
			Handler:    _Greeter_DeductStock_Handler,
		},
		{
			MethodName: "AddStock",
			Handler:    _Greeter_AddStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld/v1/greeter.proto",
}
