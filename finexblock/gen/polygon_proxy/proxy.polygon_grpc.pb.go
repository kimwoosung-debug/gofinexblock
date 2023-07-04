// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: proto/proxy/proxy.polygon.proto

package polygon_proxy

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

const (
	PolygonProxy_GetReceipt_FullMethodName      = "/polygon_proxy.PolygonProxy/GetReceipt"
	PolygonProxy_SendTransaction_FullMethodName = "/polygon_proxy.PolygonProxy/SendTransaction"
	PolygonProxy_Transfer_FullMethodName        = "/polygon_proxy.PolygonProxy/Transfer"
	PolygonProxy_GetBlockNumber_FullMethodName  = "/polygon_proxy.PolygonProxy/GetBlockNumber"
	PolygonProxy_GetBlocks_FullMethodName       = "/polygon_proxy.PolygonProxy/GetBlocks"
	PolygonProxy_CreateWallet_FullMethodName    = "/polygon_proxy.PolygonProxy/CreateWallet"
	PolygonProxy_GetBalance_FullMethodName      = "/polygon_proxy.PolygonProxy/GetBalance"
)

// PolygonProxyClient is the client API for PolygonProxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PolygonProxyClient interface {
	GetReceipt(ctx context.Context, in *GetReceiptInput, opts ...grpc.CallOption) (*GetReceiptOutput, error)
	SendTransaction(ctx context.Context, in *SendTransactionInput, opts ...grpc.CallOption) (*SendTransactionOutput, error)
	Transfer(ctx context.Context, in *TransferInput, opts ...grpc.CallOption) (*TransferOutput, error)
	GetBlockNumber(ctx context.Context, in *GetBlockNumberInput, opts ...grpc.CallOption) (*GetBlockNumberOutput, error)
	GetBlocks(ctx context.Context, in *GetBlocksInput, opts ...grpc.CallOption) (*GetBlocksOutput, error)
	CreateWallet(ctx context.Context, in *CreateWalletInput, opts ...grpc.CallOption) (*CreateWalletOutput, error)
	GetBalance(ctx context.Context, in *GetBalanceInput, opts ...grpc.CallOption) (*GetBalanceOutput, error)
}

type polygonProxyClient struct {
	cc grpc.ClientConnInterface
}

func NewPolygonProxyClient(cc grpc.ClientConnInterface) PolygonProxyClient {
	return &polygonProxyClient{cc}
}

func (c *polygonProxyClient) GetReceipt(ctx context.Context, in *GetReceiptInput, opts ...grpc.CallOption) (*GetReceiptOutput, error) {
	out := new(GetReceiptOutput)
	err := c.cc.Invoke(ctx, PolygonProxy_GetReceipt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polygonProxyClient) SendTransaction(ctx context.Context, in *SendTransactionInput, opts ...grpc.CallOption) (*SendTransactionOutput, error) {
	out := new(SendTransactionOutput)
	err := c.cc.Invoke(ctx, PolygonProxy_SendTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polygonProxyClient) Transfer(ctx context.Context, in *TransferInput, opts ...grpc.CallOption) (*TransferOutput, error) {
	out := new(TransferOutput)
	err := c.cc.Invoke(ctx, PolygonProxy_Transfer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polygonProxyClient) GetBlockNumber(ctx context.Context, in *GetBlockNumberInput, opts ...grpc.CallOption) (*GetBlockNumberOutput, error) {
	out := new(GetBlockNumberOutput)
	err := c.cc.Invoke(ctx, PolygonProxy_GetBlockNumber_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polygonProxyClient) GetBlocks(ctx context.Context, in *GetBlocksInput, opts ...grpc.CallOption) (*GetBlocksOutput, error) {
	out := new(GetBlocksOutput)
	err := c.cc.Invoke(ctx, PolygonProxy_GetBlocks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polygonProxyClient) CreateWallet(ctx context.Context, in *CreateWalletInput, opts ...grpc.CallOption) (*CreateWalletOutput, error) {
	out := new(CreateWalletOutput)
	err := c.cc.Invoke(ctx, PolygonProxy_CreateWallet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polygonProxyClient) GetBalance(ctx context.Context, in *GetBalanceInput, opts ...grpc.CallOption) (*GetBalanceOutput, error) {
	out := new(GetBalanceOutput)
	err := c.cc.Invoke(ctx, PolygonProxy_GetBalance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolygonProxyServer is the server API for PolygonProxy service.
// All implementations must embed UnimplementedPolygonProxyServer
// for forward compatibility
type PolygonProxyServer interface {
	GetReceipt(context.Context, *GetReceiptInput) (*GetReceiptOutput, error)
	SendTransaction(context.Context, *SendTransactionInput) (*SendTransactionOutput, error)
	Transfer(context.Context, *TransferInput) (*TransferOutput, error)
	GetBlockNumber(context.Context, *GetBlockNumberInput) (*GetBlockNumberOutput, error)
	GetBlocks(context.Context, *GetBlocksInput) (*GetBlocksOutput, error)
	CreateWallet(context.Context, *CreateWalletInput) (*CreateWalletOutput, error)
	GetBalance(context.Context, *GetBalanceInput) (*GetBalanceOutput, error)
	mustEmbedUnimplementedPolygonProxyServer()
}

// UnimplementedPolygonProxyServer must be embedded to have forward compatible implementations.
type UnimplementedPolygonProxyServer struct {
}

func (UnimplementedPolygonProxyServer) GetReceipt(context.Context, *GetReceiptInput) (*GetReceiptOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReceipt not implemented")
}
func (UnimplementedPolygonProxyServer) SendTransaction(context.Context, *SendTransactionInput) (*SendTransactionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTransaction not implemented")
}
func (UnimplementedPolygonProxyServer) Transfer(context.Context, *TransferInput) (*TransferOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transfer not implemented")
}
func (UnimplementedPolygonProxyServer) GetBlockNumber(context.Context, *GetBlockNumberInput) (*GetBlockNumberOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockNumber not implemented")
}
func (UnimplementedPolygonProxyServer) GetBlocks(context.Context, *GetBlocksInput) (*GetBlocksOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlocks not implemented")
}
func (UnimplementedPolygonProxyServer) CreateWallet(context.Context, *CreateWalletInput) (*CreateWalletOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWallet not implemented")
}
func (UnimplementedPolygonProxyServer) GetBalance(context.Context, *GetBalanceInput) (*GetBalanceOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedPolygonProxyServer) mustEmbedUnimplementedPolygonProxyServer() {}

// UnsafePolygonProxyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PolygonProxyServer will
// result in compilation errors.
type UnsafePolygonProxyServer interface {
	mustEmbedUnimplementedPolygonProxyServer()
}

func RegisterPolygonProxyServer(s grpc.ServiceRegistrar, srv PolygonProxyServer) {
	s.RegisterService(&PolygonProxy_ServiceDesc, srv)
}

func _PolygonProxy_GetReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReceiptInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolygonProxyServer).GetReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolygonProxy_GetReceipt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolygonProxyServer).GetReceipt(ctx, req.(*GetReceiptInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolygonProxy_SendTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTransactionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolygonProxyServer).SendTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolygonProxy_SendTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolygonProxyServer).SendTransaction(ctx, req.(*SendTransactionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolygonProxy_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolygonProxyServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolygonProxy_Transfer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolygonProxyServer).Transfer(ctx, req.(*TransferInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolygonProxy_GetBlockNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockNumberInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolygonProxyServer).GetBlockNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolygonProxy_GetBlockNumber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolygonProxyServer).GetBlockNumber(ctx, req.(*GetBlockNumberInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolygonProxy_GetBlocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlocksInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolygonProxyServer).GetBlocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolygonProxy_GetBlocks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolygonProxyServer).GetBlocks(ctx, req.(*GetBlocksInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolygonProxy_CreateWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWalletInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolygonProxyServer).CreateWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolygonProxy_CreateWallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolygonProxyServer).CreateWallet(ctx, req.(*CreateWalletInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolygonProxy_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolygonProxyServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolygonProxy_GetBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolygonProxyServer).GetBalance(ctx, req.(*GetBalanceInput))
	}
	return interceptor(ctx, in, info, handler)
}

// PolygonProxy_ServiceDesc is the grpc.ServiceDesc for PolygonProxy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PolygonProxy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "polygon_proxy.PolygonProxy",
	HandlerType: (*PolygonProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetReceipt",
			Handler:    _PolygonProxy_GetReceipt_Handler,
		},
		{
			MethodName: "SendTransaction",
			Handler:    _PolygonProxy_SendTransaction_Handler,
		},
		{
			MethodName: "Transfer",
			Handler:    _PolygonProxy_Transfer_Handler,
		},
		{
			MethodName: "GetBlockNumber",
			Handler:    _PolygonProxy_GetBlockNumber_Handler,
		},
		{
			MethodName: "GetBlocks",
			Handler:    _PolygonProxy_GetBlocks_Handler,
		},
		{
			MethodName: "CreateWallet",
			Handler:    _PolygonProxy_CreateWallet_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _PolygonProxy_GetBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/proxy/proxy.polygon.proto",
}
