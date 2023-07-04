// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: proto/grpc_order/service.proto

package grpc_order

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
	MarketOrder_MarketOrderInit_FullMethodName = "/grpc_order.MarketOrder/MarketOrderInit"
)

// MarketOrderClient is the client API for MarketOrder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MarketOrderClient interface {
	MarketOrderInit(ctx context.Context, in *MarketOrderInput, opts ...grpc.CallOption) (*Ack, error)
}

type marketOrderClient struct {
	cc grpc.ClientConnInterface
}

func NewMarketOrderClient(cc grpc.ClientConnInterface) MarketOrderClient {
	return &marketOrderClient{cc}
}

func (c *marketOrderClient) MarketOrderInit(ctx context.Context, in *MarketOrderInput, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, MarketOrder_MarketOrderInit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarketOrderServer is the server API for MarketOrder service.
// All implementations must embed UnimplementedMarketOrderServer
// for forward compatibility
type MarketOrderServer interface {
	MarketOrderInit(context.Context, *MarketOrderInput) (*Ack, error)
	mustEmbedUnimplementedMarketOrderServer()
}

// UnimplementedMarketOrderServer must be embedded to have forward compatible implementations.
type UnimplementedMarketOrderServer struct {
}

func (UnimplementedMarketOrderServer) MarketOrderInit(context.Context, *MarketOrderInput) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarketOrderInit not implemented")
}
func (UnimplementedMarketOrderServer) mustEmbedUnimplementedMarketOrderServer() {}

// UnsafeMarketOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MarketOrderServer will
// result in compilation errors.
type UnsafeMarketOrderServer interface {
	mustEmbedUnimplementedMarketOrderServer()
}

func RegisterMarketOrderServer(s grpc.ServiceRegistrar, srv MarketOrderServer) {
	s.RegisterService(&MarketOrder_ServiceDesc, srv)
}

func _MarketOrder_MarketOrderInit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketOrderInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketOrderServer).MarketOrderInit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MarketOrder_MarketOrderInit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketOrderServer).MarketOrderInit(ctx, req.(*MarketOrderInput))
	}
	return interceptor(ctx, in, info, handler)
}

// MarketOrder_ServiceDesc is the grpc.ServiceDesc for MarketOrder service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MarketOrder_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_order.MarketOrder",
	HandlerType: (*MarketOrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MarketOrderInit",
			Handler:    _MarketOrder_MarketOrderInit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/grpc_order/service.proto",
}

const (
	LimitOrder_LimitOrderInit_FullMethodName = "/grpc_order.LimitOrder/LimitOrderInit"
)

// LimitOrderClient is the client API for LimitOrder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LimitOrderClient interface {
	LimitOrderInit(ctx context.Context, in *LimitOrderInput, opts ...grpc.CallOption) (*Ack, error)
}

type limitOrderClient struct {
	cc grpc.ClientConnInterface
}

func NewLimitOrderClient(cc grpc.ClientConnInterface) LimitOrderClient {
	return &limitOrderClient{cc}
}

func (c *limitOrderClient) LimitOrderInit(ctx context.Context, in *LimitOrderInput, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, LimitOrder_LimitOrderInit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LimitOrderServer is the server API for LimitOrder service.
// All implementations must embed UnimplementedLimitOrderServer
// for forward compatibility
type LimitOrderServer interface {
	LimitOrderInit(context.Context, *LimitOrderInput) (*Ack, error)
	mustEmbedUnimplementedLimitOrderServer()
}

// UnimplementedLimitOrderServer must be embedded to have forward compatible implementations.
type UnimplementedLimitOrderServer struct {
}

func (UnimplementedLimitOrderServer) LimitOrderInit(context.Context, *LimitOrderInput) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LimitOrderInit not implemented")
}
func (UnimplementedLimitOrderServer) mustEmbedUnimplementedLimitOrderServer() {}

// UnsafeLimitOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LimitOrderServer will
// result in compilation errors.
type UnsafeLimitOrderServer interface {
	mustEmbedUnimplementedLimitOrderServer()
}

func RegisterLimitOrderServer(s grpc.ServiceRegistrar, srv LimitOrderServer) {
	s.RegisterService(&LimitOrder_ServiceDesc, srv)
}

func _LimitOrder_LimitOrderInit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LimitOrderInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LimitOrderServer).LimitOrderInit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LimitOrder_LimitOrderInit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LimitOrderServer).LimitOrderInit(ctx, req.(*LimitOrderInput))
	}
	return interceptor(ctx, in, info, handler)
}

// LimitOrder_ServiceDesc is the grpc.ServiceDesc for LimitOrder service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LimitOrder_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_order.LimitOrder",
	HandlerType: (*LimitOrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LimitOrderInit",
			Handler:    _LimitOrder_LimitOrderInit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/grpc_order/service.proto",
}

const (
	CancelOrder_CancelOrder_FullMethodName = "/grpc_order.CancelOrder/CancelOrder"
)

// CancelOrderClient is the client API for CancelOrder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CancelOrderClient interface {
	CancelOrder(ctx context.Context, in *OrderCancellation, opts ...grpc.CallOption) (*Ack, error)
}

type cancelOrderClient struct {
	cc grpc.ClientConnInterface
}

func NewCancelOrderClient(cc grpc.ClientConnInterface) CancelOrderClient {
	return &cancelOrderClient{cc}
}

func (c *cancelOrderClient) CancelOrder(ctx context.Context, in *OrderCancellation, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, CancelOrder_CancelOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CancelOrderServer is the server API for CancelOrder service.
// All implementations must embed UnimplementedCancelOrderServer
// for forward compatibility
type CancelOrderServer interface {
	CancelOrder(context.Context, *OrderCancellation) (*Ack, error)
	mustEmbedUnimplementedCancelOrderServer()
}

// UnimplementedCancelOrderServer must be embedded to have forward compatible implementations.
type UnimplementedCancelOrderServer struct {
}

func (UnimplementedCancelOrderServer) CancelOrder(context.Context, *OrderCancellation) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedCancelOrderServer) mustEmbedUnimplementedCancelOrderServer() {}

// UnsafeCancelOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CancelOrderServer will
// result in compilation errors.
type UnsafeCancelOrderServer interface {
	mustEmbedUnimplementedCancelOrderServer()
}

func RegisterCancelOrderServer(s grpc.ServiceRegistrar, srv CancelOrderServer) {
	s.RegisterService(&CancelOrder_ServiceDesc, srv)
}

func _CancelOrder_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderCancellation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CancelOrderServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CancelOrder_CancelOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CancelOrderServer).CancelOrder(ctx, req.(*OrderCancellation))
	}
	return interceptor(ctx, in, info, handler)
}

// CancelOrder_ServiceDesc is the grpc.ServiceDesc for CancelOrder service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CancelOrder_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_order.CancelOrder",
	HandlerType: (*CancelOrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CancelOrder",
			Handler:    _CancelOrder_CancelOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/grpc_order/service.proto",
}

const (
	Event_OrderPlacementEvent_FullMethodName          = "/grpc_order.Event/OrderPlacementEvent"
	Event_OrderPlacementFailedEvent_FullMethodName    = "/grpc_order.Event/OrderPlacementFailedEvent"
	Event_OrderCancellationEvent_FullMethodName       = "/grpc_order.Event/OrderCancellationEvent"
	Event_OrderCancellationFailedEvent_FullMethodName = "/grpc_order.Event/OrderCancellationFailedEvent"
	Event_OrderMatchingEvent_FullMethodName           = "/grpc_order.Event/OrderMatchingEvent"
	Event_OrderMatchingFailedEvent_FullMethodName     = "/grpc_order.Event/OrderMatchingFailedEvent"
	Event_OrderPartialFillEvent_FullMethodName        = "/grpc_order.Event/OrderPartialFillEvent"
	Event_OrderFulfillmentEvent_FullMethodName        = "/grpc_order.Event/OrderFulfillmentEvent"
	Event_OrderInitializeEvent_FullMethodName         = "/grpc_order.Event/OrderInitializeEvent"
	Event_BalanceUpdateEvent_FullMethodName           = "/grpc_order.Event/BalanceUpdateEvent"
)

// EventClient is the client API for Event service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventClient interface {
	OrderPlacementEvent(ctx context.Context, in *OrderPlacement, opts ...grpc.CallOption) (*Ack, error)
	OrderPlacementFailedEvent(ctx context.Context, in *OrderPlacementFailed, opts ...grpc.CallOption) (*Ack, error)
	OrderCancellationEvent(ctx context.Context, in *OrderCancelled, opts ...grpc.CallOption) (*Ack, error)
	OrderCancellationFailedEvent(ctx context.Context, in *OrderCancellationFailed, opts ...grpc.CallOption) (*Ack, error)
	OrderMatchingEvent(ctx context.Context, in *OrderMatching, opts ...grpc.CallOption) (*Ack, error)
	OrderMatchingFailedEvent(ctx context.Context, in *OrderMatchingFailed, opts ...grpc.CallOption) (*Ack, error)
	OrderPartialFillEvent(ctx context.Context, in *OrderPartialFill, opts ...grpc.CallOption) (*Ack, error)
	OrderFulfillmentEvent(ctx context.Context, in *OrderFulfillment, opts ...grpc.CallOption) (*Ack, error)
	OrderInitializeEvent(ctx context.Context, in *OrderInitialize, opts ...grpc.CallOption) (*Ack, error)
	BalanceUpdateEvent(ctx context.Context, in *BalanceUpdate, opts ...grpc.CallOption) (*Ack, error)
}

type eventClient struct {
	cc grpc.ClientConnInterface
}

func NewEventClient(cc grpc.ClientConnInterface) EventClient {
	return &eventClient{cc}
}

func (c *eventClient) OrderPlacementEvent(ctx context.Context, in *OrderPlacement, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, Event_OrderPlacementEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) OrderPlacementFailedEvent(ctx context.Context, in *OrderPlacementFailed, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, Event_OrderPlacementFailedEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) OrderCancellationEvent(ctx context.Context, in *OrderCancelled, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, Event_OrderCancellationEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) OrderCancellationFailedEvent(ctx context.Context, in *OrderCancellationFailed, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, Event_OrderCancellationFailedEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) OrderMatchingEvent(ctx context.Context, in *OrderMatching, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, Event_OrderMatchingEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) OrderMatchingFailedEvent(ctx context.Context, in *OrderMatchingFailed, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, Event_OrderMatchingFailedEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) OrderPartialFillEvent(ctx context.Context, in *OrderPartialFill, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, Event_OrderPartialFillEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) OrderFulfillmentEvent(ctx context.Context, in *OrderFulfillment, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, Event_OrderFulfillmentEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) OrderInitializeEvent(ctx context.Context, in *OrderInitialize, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, Event_OrderInitializeEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) BalanceUpdateEvent(ctx context.Context, in *BalanceUpdate, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, Event_BalanceUpdateEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServer is the server API for Event service.
// All implementations must embed UnimplementedEventServer
// for forward compatibility
type EventServer interface {
	OrderPlacementEvent(context.Context, *OrderPlacement) (*Ack, error)
	OrderPlacementFailedEvent(context.Context, *OrderPlacementFailed) (*Ack, error)
	OrderCancellationEvent(context.Context, *OrderCancelled) (*Ack, error)
	OrderCancellationFailedEvent(context.Context, *OrderCancellationFailed) (*Ack, error)
	OrderMatchingEvent(context.Context, *OrderMatching) (*Ack, error)
	OrderMatchingFailedEvent(context.Context, *OrderMatchingFailed) (*Ack, error)
	OrderPartialFillEvent(context.Context, *OrderPartialFill) (*Ack, error)
	OrderFulfillmentEvent(context.Context, *OrderFulfillment) (*Ack, error)
	OrderInitializeEvent(context.Context, *OrderInitialize) (*Ack, error)
	BalanceUpdateEvent(context.Context, *BalanceUpdate) (*Ack, error)
	mustEmbedUnimplementedEventServer()
}

// UnimplementedEventServer must be embedded to have forward compatible implementations.
type UnimplementedEventServer struct {
}

func (UnimplementedEventServer) OrderPlacementEvent(context.Context, *OrderPlacement) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderPlacementEvent not implemented")
}
func (UnimplementedEventServer) OrderPlacementFailedEvent(context.Context, *OrderPlacementFailed) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderPlacementFailedEvent not implemented")
}
func (UnimplementedEventServer) OrderCancellationEvent(context.Context, *OrderCancelled) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderCancellationEvent not implemented")
}
func (UnimplementedEventServer) OrderCancellationFailedEvent(context.Context, *OrderCancellationFailed) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderCancellationFailedEvent not implemented")
}
func (UnimplementedEventServer) OrderMatchingEvent(context.Context, *OrderMatching) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderMatchingEvent not implemented")
}
func (UnimplementedEventServer) OrderMatchingFailedEvent(context.Context, *OrderMatchingFailed) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderMatchingFailedEvent not implemented")
}
func (UnimplementedEventServer) OrderPartialFillEvent(context.Context, *OrderPartialFill) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderPartialFillEvent not implemented")
}
func (UnimplementedEventServer) OrderFulfillmentEvent(context.Context, *OrderFulfillment) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderFulfillmentEvent not implemented")
}
func (UnimplementedEventServer) OrderInitializeEvent(context.Context, *OrderInitialize) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderInitializeEvent not implemented")
}
func (UnimplementedEventServer) BalanceUpdateEvent(context.Context, *BalanceUpdate) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BalanceUpdateEvent not implemented")
}
func (UnimplementedEventServer) mustEmbedUnimplementedEventServer() {}

// UnsafeEventServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventServer will
// result in compilation errors.
type UnsafeEventServer interface {
	mustEmbedUnimplementedEventServer()
}

func RegisterEventServer(s grpc.ServiceRegistrar, srv EventServer) {
	s.RegisterService(&Event_ServiceDesc, srv)
}

func _Event_OrderPlacementEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderPlacement)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).OrderPlacementEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Event_OrderPlacementEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).OrderPlacementEvent(ctx, req.(*OrderPlacement))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_OrderPlacementFailedEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderPlacementFailed)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).OrderPlacementFailedEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Event_OrderPlacementFailedEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).OrderPlacementFailedEvent(ctx, req.(*OrderPlacementFailed))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_OrderCancellationEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderCancelled)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).OrderCancellationEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Event_OrderCancellationEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).OrderCancellationEvent(ctx, req.(*OrderCancelled))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_OrderCancellationFailedEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderCancellationFailed)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).OrderCancellationFailedEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Event_OrderCancellationFailedEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).OrderCancellationFailedEvent(ctx, req.(*OrderCancellationFailed))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_OrderMatchingEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderMatching)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).OrderMatchingEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Event_OrderMatchingEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).OrderMatchingEvent(ctx, req.(*OrderMatching))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_OrderMatchingFailedEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderMatchingFailed)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).OrderMatchingFailedEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Event_OrderMatchingFailedEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).OrderMatchingFailedEvent(ctx, req.(*OrderMatchingFailed))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_OrderPartialFillEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderPartialFill)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).OrderPartialFillEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Event_OrderPartialFillEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).OrderPartialFillEvent(ctx, req.(*OrderPartialFill))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_OrderFulfillmentEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderFulfillment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).OrderFulfillmentEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Event_OrderFulfillmentEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).OrderFulfillmentEvent(ctx, req.(*OrderFulfillment))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_OrderInitializeEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderInitialize)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).OrderInitializeEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Event_OrderInitializeEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).OrderInitializeEvent(ctx, req.(*OrderInitialize))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_BalanceUpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).BalanceUpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Event_BalanceUpdateEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).BalanceUpdateEvent(ctx, req.(*BalanceUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

// Event_ServiceDesc is the grpc.ServiceDesc for Event service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Event_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_order.Event",
	HandlerType: (*EventServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OrderPlacementEvent",
			Handler:    _Event_OrderPlacementEvent_Handler,
		},
		{
			MethodName: "OrderPlacementFailedEvent",
			Handler:    _Event_OrderPlacementFailedEvent_Handler,
		},
		{
			MethodName: "OrderCancellationEvent",
			Handler:    _Event_OrderCancellationEvent_Handler,
		},
		{
			MethodName: "OrderCancellationFailedEvent",
			Handler:    _Event_OrderCancellationFailedEvent_Handler,
		},
		{
			MethodName: "OrderMatchingEvent",
			Handler:    _Event_OrderMatchingEvent_Handler,
		},
		{
			MethodName: "OrderMatchingFailedEvent",
			Handler:    _Event_OrderMatchingFailedEvent_Handler,
		},
		{
			MethodName: "OrderPartialFillEvent",
			Handler:    _Event_OrderPartialFillEvent_Handler,
		},
		{
			MethodName: "OrderFulfillmentEvent",
			Handler:    _Event_OrderFulfillmentEvent_Handler,
		},
		{
			MethodName: "OrderInitializeEvent",
			Handler:    _Event_OrderInitializeEvent_Handler,
		},
		{
			MethodName: "BalanceUpdateEvent",
			Handler:    _Event_BalanceUpdateEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/grpc_order/service.proto",
}

const (
	OrderBook_GetOrderBook_FullMethodName  = "/grpc_order.OrderBook/GetOrderBook"
	OrderBook_PushOrderBook_FullMethodName = "/grpc_order.OrderBook/PushOrderBook"
)

// OrderBookClient is the client API for OrderBook service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderBookClient interface {
	GetOrderBook(ctx context.Context, in *GetOrderBookInput, opts ...grpc.CallOption) (*GetOrderBookOutput, error)
	PushOrderBook(ctx context.Context, in *PushOrderBookInput, opts ...grpc.CallOption) (*Ack, error)
}

type orderBookClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderBookClient(cc grpc.ClientConnInterface) OrderBookClient {
	return &orderBookClient{cc}
}

func (c *orderBookClient) GetOrderBook(ctx context.Context, in *GetOrderBookInput, opts ...grpc.CallOption) (*GetOrderBookOutput, error) {
	out := new(GetOrderBookOutput)
	err := c.cc.Invoke(ctx, OrderBook_GetOrderBook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderBookClient) PushOrderBook(ctx context.Context, in *PushOrderBookInput, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, OrderBook_PushOrderBook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderBookServer is the server API for OrderBook service.
// All implementations must embed UnimplementedOrderBookServer
// for forward compatibility
type OrderBookServer interface {
	GetOrderBook(context.Context, *GetOrderBookInput) (*GetOrderBookOutput, error)
	PushOrderBook(context.Context, *PushOrderBookInput) (*Ack, error)
	mustEmbedUnimplementedOrderBookServer()
}

// UnimplementedOrderBookServer must be embedded to have forward compatible implementations.
type UnimplementedOrderBookServer struct {
}

func (UnimplementedOrderBookServer) GetOrderBook(context.Context, *GetOrderBookInput) (*GetOrderBookOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderBook not implemented")
}
func (UnimplementedOrderBookServer) PushOrderBook(context.Context, *PushOrderBookInput) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushOrderBook not implemented")
}
func (UnimplementedOrderBookServer) mustEmbedUnimplementedOrderBookServer() {}

// UnsafeOrderBookServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderBookServer will
// result in compilation errors.
type UnsafeOrderBookServer interface {
	mustEmbedUnimplementedOrderBookServer()
}

func RegisterOrderBookServer(s grpc.ServiceRegistrar, srv OrderBookServer) {
	s.RegisterService(&OrderBook_ServiceDesc, srv)
}

func _OrderBook_GetOrderBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderBookInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderBookServer).GetOrderBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderBook_GetOrderBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderBookServer).GetOrderBook(ctx, req.(*GetOrderBookInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderBook_PushOrderBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushOrderBookInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderBookServer).PushOrderBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderBook_PushOrderBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderBookServer).PushOrderBook(ctx, req.(*PushOrderBookInput))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderBook_ServiceDesc is the grpc.ServiceDesc for OrderBook service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderBook_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_order.OrderBook",
	HandlerType: (*OrderBookServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrderBook",
			Handler:    _OrderBook_GetOrderBook_Handler,
		},
		{
			MethodName: "PushOrderBook",
			Handler:    _OrderBook_PushOrderBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/grpc_order/service.proto",
}
