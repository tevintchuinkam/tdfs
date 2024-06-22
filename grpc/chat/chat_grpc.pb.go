// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: chat.proto

package chat

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

// ChatSerrviceClient is the client API for ChatSerrvice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatSerrviceClient interface {
	SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type chatSerrviceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatSerrviceClient(cc grpc.ClientConnInterface) ChatSerrviceClient {
	return &chatSerrviceClient{cc}
}

func (c *chatSerrviceClient) SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/ChatSerrvice/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatSerrviceServer is the server API for ChatSerrvice service.
// All implementations must embed UnimplementedChatSerrviceServer
// for forward compatibility
type ChatSerrviceServer interface {
	SayHello(context.Context, *Message) (*Message, error)
	mustEmbedUnimplementedChatSerrviceServer()
}

// UnimplementedChatSerrviceServer must be embedded to have forward compatible implementations.
type UnimplementedChatSerrviceServer struct {
}

func (UnimplementedChatSerrviceServer) SayHello(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedChatSerrviceServer) mustEmbedUnimplementedChatSerrviceServer() {}

// UnsafeChatSerrviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatSerrviceServer will
// result in compilation errors.
type UnsafeChatSerrviceServer interface {
	mustEmbedUnimplementedChatSerrviceServer()
}

func RegisterChatSerrviceServer(s grpc.ServiceRegistrar, srv ChatSerrviceServer) {
	s.RegisterService(&ChatSerrvice_ServiceDesc, srv)
}

func _ChatSerrvice_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatSerrviceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatSerrvice/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatSerrviceServer).SayHello(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatSerrvice_ServiceDesc is the grpc.ServiceDesc for ChatSerrvice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatSerrvice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ChatSerrvice",
	HandlerType: (*ChatSerrviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _ChatSerrvice_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
