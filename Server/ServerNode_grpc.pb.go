// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package MiniProject3

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

// ServerNodeClient is the client API for ServerNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerNodeClient interface {
	Bid(ctx context.Context, in *BidRequest, opts ...grpc.CallOption) (*BidReply, error)
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error)
	Finished(ctx context.Context, in *FinishMessage, opts ...grpc.CallOption) (*FinishReply, error)
	ExportInformation(ctx context.Context, in *InfoMessage, opts ...grpc.CallOption) (*EmptyReply, error)
}

type serverNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewServerNodeClient(cc grpc.ClientConnInterface) ServerNodeClient {
	return &serverNodeClient{cc}
}

func (c *serverNodeClient) Bid(ctx context.Context, in *BidRequest, opts ...grpc.CallOption) (*BidReply, error) {
	out := new(BidReply)
	err := c.cc.Invoke(ctx, "/Server.ServerNode/Bid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverNodeClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/Server.ServerNode/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverNodeClient) Finished(ctx context.Context, in *FinishMessage, opts ...grpc.CallOption) (*FinishReply, error) {
	out := new(FinishReply)
	err := c.cc.Invoke(ctx, "/Server.ServerNode/Finished", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverNodeClient) ExportInformation(ctx context.Context, in *InfoMessage, opts ...grpc.CallOption) (*EmptyReply, error) {
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, "/Server.ServerNode/ExportInformation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerNodeServer is the server API for ServerNode service.
// All implementations must embed UnimplementedServerNodeServer
// for forward compatibility
type ServerNodeServer interface {
	Bid(context.Context, *BidRequest) (*BidReply, error)
	Status(context.Context, *StatusRequest) (*StatusReply, error)
	Finished(context.Context, *FinishMessage) (*FinishReply, error)
	ExportInformation(context.Context, *InfoMessage) (*EmptyReply, error)
	mustEmbedUnimplementedServerNodeServer()
}

// UnimplementedServerNodeServer must be embedded to have forward compatible implementations.
type UnimplementedServerNodeServer struct {
}

func (UnimplementedServerNodeServer) Bid(context.Context, *BidRequest) (*BidReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bid not implemented")
}
func (UnimplementedServerNodeServer) Status(context.Context, *StatusRequest) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedServerNodeServer) Finished(context.Context, *FinishMessage) (*FinishReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Finished not implemented")
}
func (UnimplementedServerNodeServer) ExportInformation(context.Context, *InfoMessage) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportInformation not implemented")
}
func (UnimplementedServerNodeServer) mustEmbedUnimplementedServerNodeServer() {}

// UnsafeServerNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerNodeServer will
// result in compilation errors.
type UnsafeServerNodeServer interface {
	mustEmbedUnimplementedServerNodeServer()
}

func RegisterServerNodeServer(s grpc.ServiceRegistrar, srv ServerNodeServer) {
	s.RegisterService(&ServerNode_ServiceDesc, srv)
}

func _ServerNode_Bid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerNodeServer).Bid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Server.ServerNode/Bid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerNodeServer).Bid(ctx, req.(*BidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerNode_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerNodeServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Server.ServerNode/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerNodeServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerNode_Finished_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerNodeServer).Finished(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Server.ServerNode/Finished",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerNodeServer).Finished(ctx, req.(*FinishMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerNode_ExportInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerNodeServer).ExportInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Server.ServerNode/ExportInformation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerNodeServer).ExportInformation(ctx, req.(*InfoMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// ServerNode_ServiceDesc is the grpc.ServiceDesc for ServerNode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerNode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Server.ServerNode",
	HandlerType: (*ServerNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Bid",
			Handler:    _ServerNode_Bid_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _ServerNode_Status_Handler,
		},
		{
			MethodName: "Finished",
			Handler:    _ServerNode_Finished_Handler,
		},
		{
			MethodName: "ExportInformation",
			Handler:    _ServerNode_ExportInformation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Server/ServerNode.proto",
}
