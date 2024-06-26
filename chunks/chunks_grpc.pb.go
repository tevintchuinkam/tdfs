// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: chunks.proto

package chunks

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

// ChunkServiceClient is the client API for ChunkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChunkServiceClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	StoreChunk(ctx context.Context, in *StoreChunkRequest, opts ...grpc.CallOption) (*StoreChunkResponse, error)
	GetChunk(ctx context.Context, in *GetChunkRequest, opts ...grpc.CallOption) (*GetChunkReponse, error)
}

type chunkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChunkServiceClient(cc grpc.ClientConnInterface) ChunkServiceClient {
	return &chunkServiceClient{cc}
}

func (c *chunkServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/chunks.ChunkService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chunkServiceClient) StoreChunk(ctx context.Context, in *StoreChunkRequest, opts ...grpc.CallOption) (*StoreChunkResponse, error) {
	out := new(StoreChunkResponse)
	err := c.cc.Invoke(ctx, "/chunks.ChunkService/StoreChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chunkServiceClient) GetChunk(ctx context.Context, in *GetChunkRequest, opts ...grpc.CallOption) (*GetChunkReponse, error) {
	out := new(GetChunkReponse)
	err := c.cc.Invoke(ctx, "/chunks.ChunkService/GetChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChunkServiceServer is the server API for ChunkService service.
// All implementations must embed UnimplementedChunkServiceServer
// for forward compatibility
type ChunkServiceServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	StoreChunk(context.Context, *StoreChunkRequest) (*StoreChunkResponse, error)
	GetChunk(context.Context, *GetChunkRequest) (*GetChunkReponse, error)
	mustEmbedUnimplementedChunkServiceServer()
}

// UnimplementedChunkServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChunkServiceServer struct {
}

func (UnimplementedChunkServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedChunkServiceServer) StoreChunk(context.Context, *StoreChunkRequest) (*StoreChunkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreChunk not implemented")
}
func (UnimplementedChunkServiceServer) GetChunk(context.Context, *GetChunkRequest) (*GetChunkReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChunk not implemented")
}
func (UnimplementedChunkServiceServer) mustEmbedUnimplementedChunkServiceServer() {}

// UnsafeChunkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChunkServiceServer will
// result in compilation errors.
type UnsafeChunkServiceServer interface {
	mustEmbedUnimplementedChunkServiceServer()
}

func RegisterChunkServiceServer(s grpc.ServiceRegistrar, srv ChunkServiceServer) {
	s.RegisterService(&ChunkService_ServiceDesc, srv)
}

func _ChunkService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunkServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chunks.ChunkService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunkServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChunkService_StoreChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreChunkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunkServiceServer).StoreChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chunks.ChunkService/StoreChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunkServiceServer).StoreChunk(ctx, req.(*StoreChunkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChunkService_GetChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChunkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunkServiceServer).GetChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chunks.ChunkService/GetChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunkServiceServer).GetChunk(ctx, req.(*GetChunkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChunkService_ServiceDesc is the grpc.ServiceDesc for ChunkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChunkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chunks.ChunkService",
	HandlerType: (*ChunkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _ChunkService_Ping_Handler,
		},
		{
			MethodName: "StoreChunk",
			Handler:    _ChunkService_StoreChunk_Handler,
		},
		{
			MethodName: "GetChunk",
			Handler:    _ChunkService_GetChunk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chunks.proto",
}
