// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: stock/v1/stock.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Stock_DeductStock_FullMethodName      = "/api.stock.v1.Stock/DeductStock"
	Stock_IncreaseStock_FullMethodName    = "/api.stock.v1.Stock/IncreaseStock"
	Stock_DeductIntegral_FullMethodName   = "/api.stock.v1.Stock/DeductIntegral"
	Stock_IncreaseIntegral_FullMethodName = "/api.stock.v1.Stock/IncreaseIntegral"
)

// StockClient is the client API for Stock service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StockClient interface {
	DeductStock(ctx context.Context, in *DeductStockRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	IncreaseStock(ctx context.Context, in *IncreaseStockRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeductIntegral(ctx context.Context, in *DeductIntegralRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	IncreaseIntegral(ctx context.Context, in *IncreaseIntegralRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type stockClient struct {
	cc grpc.ClientConnInterface
}

func NewStockClient(cc grpc.ClientConnInterface) StockClient {
	return &stockClient{cc}
}

func (c *stockClient) DeductStock(ctx context.Context, in *DeductStockRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Stock_DeductStock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockClient) IncreaseStock(ctx context.Context, in *IncreaseStockRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Stock_IncreaseStock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockClient) DeductIntegral(ctx context.Context, in *DeductIntegralRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Stock_DeductIntegral_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockClient) IncreaseIntegral(ctx context.Context, in *IncreaseIntegralRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Stock_IncreaseIntegral_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StockServer is the server API for Stock service.
// All implementations must embed UnimplementedStockServer
// for forward compatibility
type StockServer interface {
	DeductStock(context.Context, *DeductStockRequest) (*emptypb.Empty, error)
	IncreaseStock(context.Context, *IncreaseStockRequest) (*emptypb.Empty, error)
	DeductIntegral(context.Context, *DeductIntegralRequest) (*emptypb.Empty, error)
	IncreaseIntegral(context.Context, *IncreaseIntegralRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedStockServer()
}

// UnimplementedStockServer must be embedded to have forward compatible implementations.
type UnimplementedStockServer struct {
}

func (UnimplementedStockServer) DeductStock(context.Context, *DeductStockRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeductStock not implemented")
}
func (UnimplementedStockServer) IncreaseStock(context.Context, *IncreaseStockRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncreaseStock not implemented")
}
func (UnimplementedStockServer) DeductIntegral(context.Context, *DeductIntegralRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeductIntegral not implemented")
}
func (UnimplementedStockServer) IncreaseIntegral(context.Context, *IncreaseIntegralRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncreaseIntegral not implemented")
}
func (UnimplementedStockServer) mustEmbedUnimplementedStockServer() {}

// UnsafeStockServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StockServer will
// result in compilation errors.
type UnsafeStockServer interface {
	mustEmbedUnimplementedStockServer()
}

func RegisterStockServer(s grpc.ServiceRegistrar, srv StockServer) {
	s.RegisterService(&Stock_ServiceDesc, srv)
}

func _Stock_DeductStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeductStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServer).DeductStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Stock_DeductStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServer).DeductStock(ctx, req.(*DeductStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stock_IncreaseStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncreaseStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServer).IncreaseStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Stock_IncreaseStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServer).IncreaseStock(ctx, req.(*IncreaseStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stock_DeductIntegral_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeductIntegralRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServer).DeductIntegral(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Stock_DeductIntegral_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServer).DeductIntegral(ctx, req.(*DeductIntegralRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stock_IncreaseIntegral_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncreaseIntegralRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServer).IncreaseIntegral(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Stock_IncreaseIntegral_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServer).IncreaseIntegral(ctx, req.(*IncreaseIntegralRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Stock_ServiceDesc is the grpc.ServiceDesc for Stock service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Stock_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.stock.v1.Stock",
	HandlerType: (*StockServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeductStock",
			Handler:    _Stock_DeductStock_Handler,
		},
		{
			MethodName: "IncreaseStock",
			Handler:    _Stock_IncreaseStock_Handler,
		},
		{
			MethodName: "DeductIntegral",
			Handler:    _Stock_DeductIntegral_Handler,
		},
		{
			MethodName: "IncreaseIntegral",
			Handler:    _Stock_IncreaseIntegral_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stock/v1/stock.proto",
}
