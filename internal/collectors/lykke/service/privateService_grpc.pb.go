// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: privateService.proto

package service

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

// PrivateServiceClient is the client API for PrivateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrivateServiceClient interface {
	GetBalances(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*BalancesResponse, error)
	PlaceLimitOrder(ctx context.Context, in *LimitOrderRequest, opts ...grpc.CallOption) (*LimitOrderResponse, error)
	PlaceBulkLimitOrder(ctx context.Context, in *BulkLimitOrderRequest, opts ...grpc.CallOption) (*BulkLimitOrderResponse, error)
	PlaceMarketOrder(ctx context.Context, in *MarketOrderRequest, opts ...grpc.CallOption) (*MarketOrderResponse, error)
	GetOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error)
	GetActiveOrders(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (*OrdersResponse, error)
	GetClosedOrders(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (*OrdersResponse, error)
	CancelAllOrders(ctx context.Context, in *CancelOrdersRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error)
	CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error)
	GetTrades(ctx context.Context, in *TradesRequest, opts ...grpc.CallOption) (*TradesResponse, error)
	GetOrderTrades(ctx context.Context, in *OrderTradesRequest, opts ...grpc.CallOption) (*TradesResponse, error)
	GetBalanceUpdates(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (PrivateService_GetBalanceUpdatesClient, error)
	GetTradeUpdates(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (PrivateService_GetTradeUpdatesClient, error)
	GetOperationsHistory(ctx context.Context, in *GetOperationsHistoryRequest, opts ...grpc.CallOption) (*GetOperationsHistoryResponse, error)
	CreateDepositAddresses(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CreateDepositAddressesResponse, error)
	GetDepositAddress(ctx context.Context, in *GetDepositAddressRequest, opts ...grpc.CallOption) (*GetDepositAddressResponse, error)
	GetDepositAddresses(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetDepositAddressesResponse, error)
	CreateWithdrawal(ctx context.Context, in *CreateWithdrawalRequest, opts ...grpc.CallOption) (*CreateWithdrawalResponse, error)
}

type privateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPrivateServiceClient(cc grpc.ClientConnInterface) PrivateServiceClient {
	return &privateServiceClient{cc}
}

func (c *privateServiceClient) GetBalances(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*BalancesResponse, error) {
	out := new(BalancesResponse)
	err := c.cc.Invoke(ctx, "/hft.PrivateService/GetBalances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateServiceClient) PlaceLimitOrder(ctx context.Context, in *LimitOrderRequest, opts ...grpc.CallOption) (*LimitOrderResponse, error) {
	out := new(LimitOrderResponse)
	err := c.cc.Invoke(ctx, "/hft.PrivateService/PlaceLimitOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateServiceClient) PlaceBulkLimitOrder(ctx context.Context, in *BulkLimitOrderRequest, opts ...grpc.CallOption) (*BulkLimitOrderResponse, error) {
	out := new(BulkLimitOrderResponse)
	err := c.cc.Invoke(ctx, "/hft.PrivateService/PlaceBulkLimitOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateServiceClient) PlaceMarketOrder(ctx context.Context, in *MarketOrderRequest, opts ...grpc.CallOption) (*MarketOrderResponse, error) {
	out := new(MarketOrderResponse)
	err := c.cc.Invoke(ctx, "/hft.PrivateService/PlaceMarketOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateServiceClient) GetOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, "/hft.PrivateService/GetOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateServiceClient) GetActiveOrders(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (*OrdersResponse, error) {
	out := new(OrdersResponse)
	err := c.cc.Invoke(ctx, "/hft.PrivateService/GetActiveOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateServiceClient) GetClosedOrders(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (*OrdersResponse, error) {
	out := new(OrdersResponse)
	err := c.cc.Invoke(ctx, "/hft.PrivateService/GetClosedOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateServiceClient) CancelAllOrders(ctx context.Context, in *CancelOrdersRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error) {
	out := new(CancelOrderResponse)
	err := c.cc.Invoke(ctx, "/hft.PrivateService/CancelAllOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateServiceClient) CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error) {
	out := new(CancelOrderResponse)
	err := c.cc.Invoke(ctx, "/hft.PrivateService/CancelOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateServiceClient) GetTrades(ctx context.Context, in *TradesRequest, opts ...grpc.CallOption) (*TradesResponse, error) {
	out := new(TradesResponse)
	err := c.cc.Invoke(ctx, "/hft.PrivateService/GetTrades", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateServiceClient) GetOrderTrades(ctx context.Context, in *OrderTradesRequest, opts ...grpc.CallOption) (*TradesResponse, error) {
	out := new(TradesResponse)
	err := c.cc.Invoke(ctx, "/hft.PrivateService/GetOrderTrades", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateServiceClient) GetBalanceUpdates(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (PrivateService_GetBalanceUpdatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &PrivateService_ServiceDesc.Streams[0], "/hft.PrivateService/GetBalanceUpdates", opts...)
	if err != nil {
		return nil, err
	}
	x := &privateServiceGetBalanceUpdatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PrivateService_GetBalanceUpdatesClient interface {
	Recv() (*BalanceUpdate, error)
	grpc.ClientStream
}

type privateServiceGetBalanceUpdatesClient struct {
	grpc.ClientStream
}

func (x *privateServiceGetBalanceUpdatesClient) Recv() (*BalanceUpdate, error) {
	m := new(BalanceUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *privateServiceClient) GetTradeUpdates(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (PrivateService_GetTradeUpdatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &PrivateService_ServiceDesc.Streams[1], "/hft.PrivateService/GetTradeUpdates", opts...)
	if err != nil {
		return nil, err
	}
	x := &privateServiceGetTradeUpdatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PrivateService_GetTradeUpdatesClient interface {
	Recv() (*TradeUpdate, error)
	grpc.ClientStream
}

type privateServiceGetTradeUpdatesClient struct {
	grpc.ClientStream
}

func (x *privateServiceGetTradeUpdatesClient) Recv() (*TradeUpdate, error) {
	m := new(TradeUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *privateServiceClient) GetOperationsHistory(ctx context.Context, in *GetOperationsHistoryRequest, opts ...grpc.CallOption) (*GetOperationsHistoryResponse, error) {
	out := new(GetOperationsHistoryResponse)
	err := c.cc.Invoke(ctx, "/hft.PrivateService/GetOperationsHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateServiceClient) CreateDepositAddresses(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CreateDepositAddressesResponse, error) {
	out := new(CreateDepositAddressesResponse)
	err := c.cc.Invoke(ctx, "/hft.PrivateService/CreateDepositAddresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateServiceClient) GetDepositAddress(ctx context.Context, in *GetDepositAddressRequest, opts ...grpc.CallOption) (*GetDepositAddressResponse, error) {
	out := new(GetDepositAddressResponse)
	err := c.cc.Invoke(ctx, "/hft.PrivateService/GetDepositAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateServiceClient) GetDepositAddresses(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetDepositAddressesResponse, error) {
	out := new(GetDepositAddressesResponse)
	err := c.cc.Invoke(ctx, "/hft.PrivateService/GetDepositAddresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *privateServiceClient) CreateWithdrawal(ctx context.Context, in *CreateWithdrawalRequest, opts ...grpc.CallOption) (*CreateWithdrawalResponse, error) {
	out := new(CreateWithdrawalResponse)
	err := c.cc.Invoke(ctx, "/hft.PrivateService/CreateWithdrawal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrivateServiceServer is the server API for PrivateService service.
// All implementations should embed UnimplementedPrivateServiceServer
// for forward compatibility
type PrivateServiceServer interface {
	GetBalances(context.Context, *emptypb.Empty) (*BalancesResponse, error)
	PlaceLimitOrder(context.Context, *LimitOrderRequest) (*LimitOrderResponse, error)
	PlaceBulkLimitOrder(context.Context, *BulkLimitOrderRequest) (*BulkLimitOrderResponse, error)
	PlaceMarketOrder(context.Context, *MarketOrderRequest) (*MarketOrderResponse, error)
	GetOrder(context.Context, *OrderRequest) (*OrderResponse, error)
	GetActiveOrders(context.Context, *OrdersRequest) (*OrdersResponse, error)
	GetClosedOrders(context.Context, *OrdersRequest) (*OrdersResponse, error)
	CancelAllOrders(context.Context, *CancelOrdersRequest) (*CancelOrderResponse, error)
	CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error)
	GetTrades(context.Context, *TradesRequest) (*TradesResponse, error)
	GetOrderTrades(context.Context, *OrderTradesRequest) (*TradesResponse, error)
	GetBalanceUpdates(*emptypb.Empty, PrivateService_GetBalanceUpdatesServer) error
	GetTradeUpdates(*emptypb.Empty, PrivateService_GetTradeUpdatesServer) error
	GetOperationsHistory(context.Context, *GetOperationsHistoryRequest) (*GetOperationsHistoryResponse, error)
	CreateDepositAddresses(context.Context, *emptypb.Empty) (*CreateDepositAddressesResponse, error)
	GetDepositAddress(context.Context, *GetDepositAddressRequest) (*GetDepositAddressResponse, error)
	GetDepositAddresses(context.Context, *emptypb.Empty) (*GetDepositAddressesResponse, error)
	CreateWithdrawal(context.Context, *CreateWithdrawalRequest) (*CreateWithdrawalResponse, error)
}

// UnimplementedPrivateServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPrivateServiceServer struct {
}

func (UnimplementedPrivateServiceServer) GetBalances(context.Context, *emptypb.Empty) (*BalancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalances not implemented")
}
func (UnimplementedPrivateServiceServer) PlaceLimitOrder(context.Context, *LimitOrderRequest) (*LimitOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceLimitOrder not implemented")
}
func (UnimplementedPrivateServiceServer) PlaceBulkLimitOrder(context.Context, *BulkLimitOrderRequest) (*BulkLimitOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceBulkLimitOrder not implemented")
}
func (UnimplementedPrivateServiceServer) PlaceMarketOrder(context.Context, *MarketOrderRequest) (*MarketOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceMarketOrder not implemented")
}
func (UnimplementedPrivateServiceServer) GetOrder(context.Context, *OrderRequest) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedPrivateServiceServer) GetActiveOrders(context.Context, *OrdersRequest) (*OrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveOrders not implemented")
}
func (UnimplementedPrivateServiceServer) GetClosedOrders(context.Context, *OrdersRequest) (*OrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClosedOrders not implemented")
}
func (UnimplementedPrivateServiceServer) CancelAllOrders(context.Context, *CancelOrdersRequest) (*CancelOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelAllOrders not implemented")
}
func (UnimplementedPrivateServiceServer) CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedPrivateServiceServer) GetTrades(context.Context, *TradesRequest) (*TradesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrades not implemented")
}
func (UnimplementedPrivateServiceServer) GetOrderTrades(context.Context, *OrderTradesRequest) (*TradesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderTrades not implemented")
}
func (UnimplementedPrivateServiceServer) GetBalanceUpdates(*emptypb.Empty, PrivateService_GetBalanceUpdatesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBalanceUpdates not implemented")
}
func (UnimplementedPrivateServiceServer) GetTradeUpdates(*emptypb.Empty, PrivateService_GetTradeUpdatesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTradeUpdates not implemented")
}
func (UnimplementedPrivateServiceServer) GetOperationsHistory(context.Context, *GetOperationsHistoryRequest) (*GetOperationsHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOperationsHistory not implemented")
}
func (UnimplementedPrivateServiceServer) CreateDepositAddresses(context.Context, *emptypb.Empty) (*CreateDepositAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDepositAddresses not implemented")
}
func (UnimplementedPrivateServiceServer) GetDepositAddress(context.Context, *GetDepositAddressRequest) (*GetDepositAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDepositAddress not implemented")
}
func (UnimplementedPrivateServiceServer) GetDepositAddresses(context.Context, *emptypb.Empty) (*GetDepositAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDepositAddresses not implemented")
}
func (UnimplementedPrivateServiceServer) CreateWithdrawal(context.Context, *CreateWithdrawalRequest) (*CreateWithdrawalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWithdrawal not implemented")
}

// UnsafePrivateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrivateServiceServer will
// result in compilation errors.
type UnsafePrivateServiceServer interface {
	mustEmbedUnimplementedPrivateServiceServer()
}

func RegisterPrivateServiceServer(s grpc.ServiceRegistrar, srv PrivateServiceServer) {
	s.RegisterService(&PrivateService_ServiceDesc, srv)
}

func _PrivateService_GetBalances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServiceServer).GetBalances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hft.PrivateService/GetBalances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServiceServer).GetBalances(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateService_PlaceLimitOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LimitOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServiceServer).PlaceLimitOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hft.PrivateService/PlaceLimitOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServiceServer).PlaceLimitOrder(ctx, req.(*LimitOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateService_PlaceBulkLimitOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkLimitOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServiceServer).PlaceBulkLimitOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hft.PrivateService/PlaceBulkLimitOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServiceServer).PlaceBulkLimitOrder(ctx, req.(*BulkLimitOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateService_PlaceMarketOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServiceServer).PlaceMarketOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hft.PrivateService/PlaceMarketOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServiceServer).PlaceMarketOrder(ctx, req.(*MarketOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateService_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hft.PrivateService/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServiceServer).GetOrder(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateService_GetActiveOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServiceServer).GetActiveOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hft.PrivateService/GetActiveOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServiceServer).GetActiveOrders(ctx, req.(*OrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateService_GetClosedOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServiceServer).GetClosedOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hft.PrivateService/GetClosedOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServiceServer).GetClosedOrders(ctx, req.(*OrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateService_CancelAllOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServiceServer).CancelAllOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hft.PrivateService/CancelAllOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServiceServer).CancelAllOrders(ctx, req.(*CancelOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateService_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServiceServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hft.PrivateService/CancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServiceServer).CancelOrder(ctx, req.(*CancelOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateService_GetTrades_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServiceServer).GetTrades(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hft.PrivateService/GetTrades",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServiceServer).GetTrades(ctx, req.(*TradesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateService_GetOrderTrades_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderTradesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServiceServer).GetOrderTrades(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hft.PrivateService/GetOrderTrades",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServiceServer).GetOrderTrades(ctx, req.(*OrderTradesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateService_GetBalanceUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PrivateServiceServer).GetBalanceUpdates(m, &privateServiceGetBalanceUpdatesServer{stream})
}

type PrivateService_GetBalanceUpdatesServer interface {
	Send(*BalanceUpdate) error
	grpc.ServerStream
}

type privateServiceGetBalanceUpdatesServer struct {
	grpc.ServerStream
}

func (x *privateServiceGetBalanceUpdatesServer) Send(m *BalanceUpdate) error {
	return x.ServerStream.SendMsg(m)
}

func _PrivateService_GetTradeUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PrivateServiceServer).GetTradeUpdates(m, &privateServiceGetTradeUpdatesServer{stream})
}

type PrivateService_GetTradeUpdatesServer interface {
	Send(*TradeUpdate) error
	grpc.ServerStream
}

type privateServiceGetTradeUpdatesServer struct {
	grpc.ServerStream
}

func (x *privateServiceGetTradeUpdatesServer) Send(m *TradeUpdate) error {
	return x.ServerStream.SendMsg(m)
}

func _PrivateService_GetOperationsHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperationsHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServiceServer).GetOperationsHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hft.PrivateService/GetOperationsHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServiceServer).GetOperationsHistory(ctx, req.(*GetOperationsHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateService_CreateDepositAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServiceServer).CreateDepositAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hft.PrivateService/CreateDepositAddresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServiceServer).CreateDepositAddresses(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateService_GetDepositAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDepositAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServiceServer).GetDepositAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hft.PrivateService/GetDepositAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServiceServer).GetDepositAddress(ctx, req.(*GetDepositAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateService_GetDepositAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServiceServer).GetDepositAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hft.PrivateService/GetDepositAddresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServiceServer).GetDepositAddresses(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrivateService_CreateWithdrawal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWithdrawalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivateServiceServer).CreateWithdrawal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hft.PrivateService/CreateWithdrawal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivateServiceServer).CreateWithdrawal(ctx, req.(*CreateWithdrawalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PrivateService_ServiceDesc is the grpc.ServiceDesc for PrivateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PrivateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hft.PrivateService",
	HandlerType: (*PrivateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBalances",
			Handler:    _PrivateService_GetBalances_Handler,
		},
		{
			MethodName: "PlaceLimitOrder",
			Handler:    _PrivateService_PlaceLimitOrder_Handler,
		},
		{
			MethodName: "PlaceBulkLimitOrder",
			Handler:    _PrivateService_PlaceBulkLimitOrder_Handler,
		},
		{
			MethodName: "PlaceMarketOrder",
			Handler:    _PrivateService_PlaceMarketOrder_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _PrivateService_GetOrder_Handler,
		},
		{
			MethodName: "GetActiveOrders",
			Handler:    _PrivateService_GetActiveOrders_Handler,
		},
		{
			MethodName: "GetClosedOrders",
			Handler:    _PrivateService_GetClosedOrders_Handler,
		},
		{
			MethodName: "CancelAllOrders",
			Handler:    _PrivateService_CancelAllOrders_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _PrivateService_CancelOrder_Handler,
		},
		{
			MethodName: "GetTrades",
			Handler:    _PrivateService_GetTrades_Handler,
		},
		{
			MethodName: "GetOrderTrades",
			Handler:    _PrivateService_GetOrderTrades_Handler,
		},
		{
			MethodName: "GetOperationsHistory",
			Handler:    _PrivateService_GetOperationsHistory_Handler,
		},
		{
			MethodName: "CreateDepositAddresses",
			Handler:    _PrivateService_CreateDepositAddresses_Handler,
		},
		{
			MethodName: "GetDepositAddress",
			Handler:    _PrivateService_GetDepositAddress_Handler,
		},
		{
			MethodName: "GetDepositAddresses",
			Handler:    _PrivateService_GetDepositAddresses_Handler,
		},
		{
			MethodName: "CreateWithdrawal",
			Handler:    _PrivateService_CreateWithdrawal_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetBalanceUpdates",
			Handler:       _PrivateService_GetBalanceUpdates_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetTradeUpdates",
			Handler:       _PrivateService_GetTradeUpdates_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "privateService.proto",
}