// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: pkg/connectionmanager/grpc/quarkcmsvc.proto

package grpc

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

// QuarkCMServiceClient is the client API for QuarkCMService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuarkCMServiceClient interface {
	// Test
	TestPing(ctx context.Context, in *TestRequestMessage, opts ...grpc.CallOption) (*TestResponseMessage, error)
	ListNode(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*NodeListMessage, error)
	WatchNode(ctx context.Context, in *MaxResourceVersionMessage, opts ...grpc.CallOption) (QuarkCMService_WatchNodeClient, error)
	ListPod(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PodListMessage, error)
	WatchPod(ctx context.Context, in *MaxResourceVersionMessage, opts ...grpc.CallOption) (QuarkCMService_WatchPodClient, error)
}

type quarkCMServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQuarkCMServiceClient(cc grpc.ClientConnInterface) QuarkCMServiceClient {
	return &quarkCMServiceClient{cc}
}

func (c *quarkCMServiceClient) TestPing(ctx context.Context, in *TestRequestMessage, opts ...grpc.CallOption) (*TestResponseMessage, error) {
	out := new(TestResponseMessage)
	err := c.cc.Invoke(ctx, "/quarkcmsvc.QuarkCMService/TestPing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quarkCMServiceClient) ListNode(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*NodeListMessage, error) {
	out := new(NodeListMessage)
	err := c.cc.Invoke(ctx, "/quarkcmsvc.QuarkCMService/ListNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quarkCMServiceClient) WatchNode(ctx context.Context, in *MaxResourceVersionMessage, opts ...grpc.CallOption) (QuarkCMService_WatchNodeClient, error) {
	stream, err := c.cc.NewStream(ctx, &QuarkCMService_ServiceDesc.Streams[0], "/quarkcmsvc.QuarkCMService/WatchNode", opts...)
	if err != nil {
		return nil, err
	}
	x := &quarkCMServiceWatchNodeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type QuarkCMService_WatchNodeClient interface {
	Recv() (*NodeMessage, error)
	grpc.ClientStream
}

type quarkCMServiceWatchNodeClient struct {
	grpc.ClientStream
}

func (x *quarkCMServiceWatchNodeClient) Recv() (*NodeMessage, error) {
	m := new(NodeMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *quarkCMServiceClient) ListPod(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PodListMessage, error) {
	out := new(PodListMessage)
	err := c.cc.Invoke(ctx, "/quarkcmsvc.QuarkCMService/ListPod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quarkCMServiceClient) WatchPod(ctx context.Context, in *MaxResourceVersionMessage, opts ...grpc.CallOption) (QuarkCMService_WatchPodClient, error) {
	stream, err := c.cc.NewStream(ctx, &QuarkCMService_ServiceDesc.Streams[1], "/quarkcmsvc.QuarkCMService/WatchPod", opts...)
	if err != nil {
		return nil, err
	}
	x := &quarkCMServiceWatchPodClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type QuarkCMService_WatchPodClient interface {
	Recv() (*PodMessage, error)
	grpc.ClientStream
}

type quarkCMServiceWatchPodClient struct {
	grpc.ClientStream
}

func (x *quarkCMServiceWatchPodClient) Recv() (*PodMessage, error) {
	m := new(PodMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// QuarkCMServiceServer is the server API for QuarkCMService service.
// All implementations must embed UnimplementedQuarkCMServiceServer
// for forward compatibility
type QuarkCMServiceServer interface {
	// Test
	TestPing(context.Context, *TestRequestMessage) (*TestResponseMessage, error)
	ListNode(context.Context, *emptypb.Empty) (*NodeListMessage, error)
	WatchNode(*MaxResourceVersionMessage, QuarkCMService_WatchNodeServer) error
	ListPod(context.Context, *emptypb.Empty) (*PodListMessage, error)
	WatchPod(*MaxResourceVersionMessage, QuarkCMService_WatchPodServer) error
	mustEmbedUnimplementedQuarkCMServiceServer()
}

// UnimplementedQuarkCMServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQuarkCMServiceServer struct {
}

func (UnimplementedQuarkCMServiceServer) TestPing(context.Context, *TestRequestMessage) (*TestResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestPing not implemented")
}
func (UnimplementedQuarkCMServiceServer) ListNode(context.Context, *emptypb.Empty) (*NodeListMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNode not implemented")
}
func (UnimplementedQuarkCMServiceServer) WatchNode(*MaxResourceVersionMessage, QuarkCMService_WatchNodeServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchNode not implemented")
}
func (UnimplementedQuarkCMServiceServer) ListPod(context.Context, *emptypb.Empty) (*PodListMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPod not implemented")
}
func (UnimplementedQuarkCMServiceServer) WatchPod(*MaxResourceVersionMessage, QuarkCMService_WatchPodServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchPod not implemented")
}
func (UnimplementedQuarkCMServiceServer) mustEmbedUnimplementedQuarkCMServiceServer() {}

// UnsafeQuarkCMServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuarkCMServiceServer will
// result in compilation errors.
type UnsafeQuarkCMServiceServer interface {
	mustEmbedUnimplementedQuarkCMServiceServer()
}

func RegisterQuarkCMServiceServer(s grpc.ServiceRegistrar, srv QuarkCMServiceServer) {
	s.RegisterService(&QuarkCMService_ServiceDesc, srv)
}

func _QuarkCMService_TestPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuarkCMServiceServer).TestPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quarkcmsvc.QuarkCMService/TestPing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuarkCMServiceServer).TestPing(ctx, req.(*TestRequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuarkCMService_ListNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuarkCMServiceServer).ListNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quarkcmsvc.QuarkCMService/ListNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuarkCMServiceServer).ListNode(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuarkCMService_WatchNode_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MaxResourceVersionMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QuarkCMServiceServer).WatchNode(m, &quarkCMServiceWatchNodeServer{stream})
}

type QuarkCMService_WatchNodeServer interface {
	Send(*NodeMessage) error
	grpc.ServerStream
}

type quarkCMServiceWatchNodeServer struct {
	grpc.ServerStream
}

func (x *quarkCMServiceWatchNodeServer) Send(m *NodeMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _QuarkCMService_ListPod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuarkCMServiceServer).ListPod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quarkcmsvc.QuarkCMService/ListPod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuarkCMServiceServer).ListPod(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuarkCMService_WatchPod_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MaxResourceVersionMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QuarkCMServiceServer).WatchPod(m, &quarkCMServiceWatchPodServer{stream})
}

type QuarkCMService_WatchPodServer interface {
	Send(*PodMessage) error
	grpc.ServerStream
}

type quarkCMServiceWatchPodServer struct {
	grpc.ServerStream
}

func (x *quarkCMServiceWatchPodServer) Send(m *PodMessage) error {
	return x.ServerStream.SendMsg(m)
}

// QuarkCMService_ServiceDesc is the grpc.ServiceDesc for QuarkCMService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuarkCMService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "quarkcmsvc.QuarkCMService",
	HandlerType: (*QuarkCMServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestPing",
			Handler:    _QuarkCMService_TestPing_Handler,
		},
		{
			MethodName: "ListNode",
			Handler:    _QuarkCMService_ListNode_Handler,
		},
		{
			MethodName: "ListPod",
			Handler:    _QuarkCMService_ListPod_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchNode",
			Handler:       _QuarkCMService_WatchNode_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WatchPod",
			Handler:       _QuarkCMService_WatchPod_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/connectionmanager/grpc/quarkcmsvc.proto",
}
