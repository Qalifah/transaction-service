// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// TransactionClient is the client API for Transaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionClient interface {
	CreditAccount(ctx context.Context, in *Transfer, opts ...grpc.CallOption) (*TransferResponse, error)
	DebitAccount(ctx context.Context, in *Transfer, opts ...grpc.CallOption) (*TransferResponse, error)
}

type transactionClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionClient(cc grpc.ClientConnInterface) TransactionClient {
	return &transactionClient{cc}
}

func (c *transactionClient) CreditAccount(ctx context.Context, in *Transfer, opts ...grpc.CallOption) (*TransferResponse, error) {
	out := new(TransferResponse)
	err := c.cc.Invoke(ctx, "/pb.Transaction/CreditAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionClient) DebitAccount(ctx context.Context, in *Transfer, opts ...grpc.CallOption) (*TransferResponse, error) {
	out := new(TransferResponse)
	err := c.cc.Invoke(ctx, "/pb.Transaction/DebitAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServer is the server API for Transaction service.
// All implementations must embed UnimplementedTransactionServer
// for forward compatibility
type TransactionServer interface {
	CreditAccount(context.Context, *Transfer) (*TransferResponse, error)
	DebitAccount(context.Context, *Transfer) (*TransferResponse, error)
	mustEmbedUnimplementedTransactionServer()
}

// UnimplementedTransactionServer must be embedded to have forward compatible implementations.
type UnimplementedTransactionServer struct {
}

func (UnimplementedTransactionServer) CreditAccount(context.Context, *Transfer) (*TransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreditAccount not implemented")
}
func (UnimplementedTransactionServer) DebitAccount(context.Context, *Transfer) (*TransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DebitAccount not implemented")
}
func (UnimplementedTransactionServer) mustEmbedUnimplementedTransactionServer() {}

// UnsafeTransactionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionServer will
// result in compilation errors.
type UnsafeTransactionServer interface {
	mustEmbedUnimplementedTransactionServer()
}

func RegisterTransactionServer(s grpc.ServiceRegistrar, srv TransactionServer) {
	s.RegisterService(&Transaction_ServiceDesc, srv)
}

func _Transaction_CreditAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transfer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).CreditAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Transaction/CreditAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).CreditAccount(ctx, req.(*Transfer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transaction_DebitAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transfer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).DebitAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Transaction/DebitAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).DebitAccount(ctx, req.(*Transfer))
	}
	return interceptor(ctx, in, info, handler)
}

// Transaction_ServiceDesc is the grpc.ServiceDesc for Transaction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Transaction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Transaction",
	HandlerType: (*TransactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreditAccount",
			Handler:    _Transaction_CreditAccount_Handler,
		},
		{
			MethodName: "DebitAccount",
			Handler:    _Transaction_DebitAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/transaction.proto",
}