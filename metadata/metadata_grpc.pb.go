// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: metadata.proto

package metadata

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

// MetadataServiceClient is the client API for MetadataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetadataServiceClient interface {
	RegisterFileCreation(ctx context.Context, in *RecRequest, opts ...grpc.CallOption) (*RecResponse, error)
	GetLocation(ctx context.Context, in *LocRequest, opts ...grpc.CallOption) (*LocResponse, error)
	OpenDir(ctx context.Context, in *OpenDirRequest, opts ...grpc.CallOption) (*OpenDirResponse, error)
	ReadDir(ctx context.Context, in *ReadDirRequest, opts ...grpc.CallOption) (*FileInfo, error)
	ReadDirAll(ctx context.Context, in *ReadDirRequest, opts ...grpc.CallOption) (*ReadDirAllResponse, error)
	MkDir(ctx context.Context, in *MkDirRequest, opts ...grpc.CallOption) (*MkDirResponse, error)
	DeleteAllData(ctx context.Context, in *DeleteAllDataRequest, opts ...grpc.CallOption) (*DeleteAllDataReponse, error)
}

type metadataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetadataServiceClient(cc grpc.ClientConnInterface) MetadataServiceClient {
	return &metadataServiceClient{cc}
}

func (c *metadataServiceClient) RegisterFileCreation(ctx context.Context, in *RecRequest, opts ...grpc.CallOption) (*RecResponse, error) {
	out := new(RecResponse)
	err := c.cc.Invoke(ctx, "/metadata.MetadataService/RegisterFileCreation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) GetLocation(ctx context.Context, in *LocRequest, opts ...grpc.CallOption) (*LocResponse, error) {
	out := new(LocResponse)
	err := c.cc.Invoke(ctx, "/metadata.MetadataService/GetLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) OpenDir(ctx context.Context, in *OpenDirRequest, opts ...grpc.CallOption) (*OpenDirResponse, error) {
	out := new(OpenDirResponse)
	err := c.cc.Invoke(ctx, "/metadata.MetadataService/OpenDir", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) ReadDir(ctx context.Context, in *ReadDirRequest, opts ...grpc.CallOption) (*FileInfo, error) {
	out := new(FileInfo)
	err := c.cc.Invoke(ctx, "/metadata.MetadataService/ReadDir", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) ReadDirAll(ctx context.Context, in *ReadDirRequest, opts ...grpc.CallOption) (*ReadDirAllResponse, error) {
	out := new(ReadDirAllResponse)
	err := c.cc.Invoke(ctx, "/metadata.MetadataService/ReadDirAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) MkDir(ctx context.Context, in *MkDirRequest, opts ...grpc.CallOption) (*MkDirResponse, error) {
	out := new(MkDirResponse)
	err := c.cc.Invoke(ctx, "/metadata.MetadataService/MkDir", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataServiceClient) DeleteAllData(ctx context.Context, in *DeleteAllDataRequest, opts ...grpc.CallOption) (*DeleteAllDataReponse, error) {
	out := new(DeleteAllDataReponse)
	err := c.cc.Invoke(ctx, "/metadata.MetadataService/DeleteAllData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetadataServiceServer is the server API for MetadataService service.
// All implementations must embed UnimplementedMetadataServiceServer
// for forward compatibility
type MetadataServiceServer interface {
	RegisterFileCreation(context.Context, *RecRequest) (*RecResponse, error)
	GetLocation(context.Context, *LocRequest) (*LocResponse, error)
	OpenDir(context.Context, *OpenDirRequest) (*OpenDirResponse, error)
	ReadDir(context.Context, *ReadDirRequest) (*FileInfo, error)
	ReadDirAll(context.Context, *ReadDirRequest) (*ReadDirAllResponse, error)
	MkDir(context.Context, *MkDirRequest) (*MkDirResponse, error)
	DeleteAllData(context.Context, *DeleteAllDataRequest) (*DeleteAllDataReponse, error)
	mustEmbedUnimplementedMetadataServiceServer()
}

// UnimplementedMetadataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMetadataServiceServer struct {
}

func (UnimplementedMetadataServiceServer) RegisterFileCreation(context.Context, *RecRequest) (*RecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterFileCreation not implemented")
}
func (UnimplementedMetadataServiceServer) GetLocation(context.Context, *LocRequest) (*LocResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocation not implemented")
}
func (UnimplementedMetadataServiceServer) OpenDir(context.Context, *OpenDirRequest) (*OpenDirResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenDir not implemented")
}
func (UnimplementedMetadataServiceServer) ReadDir(context.Context, *ReadDirRequest) (*FileInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadDir not implemented")
}
func (UnimplementedMetadataServiceServer) ReadDirAll(context.Context, *ReadDirRequest) (*ReadDirAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadDirAll not implemented")
}
func (UnimplementedMetadataServiceServer) MkDir(context.Context, *MkDirRequest) (*MkDirResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MkDir not implemented")
}
func (UnimplementedMetadataServiceServer) DeleteAllData(context.Context, *DeleteAllDataRequest) (*DeleteAllDataReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAllData not implemented")
}
func (UnimplementedMetadataServiceServer) mustEmbedUnimplementedMetadataServiceServer() {}

// UnsafeMetadataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetadataServiceServer will
// result in compilation errors.
type UnsafeMetadataServiceServer interface {
	mustEmbedUnimplementedMetadataServiceServer()
}

func RegisterMetadataServiceServer(s grpc.ServiceRegistrar, srv MetadataServiceServer) {
	s.RegisterService(&MetadataService_ServiceDesc, srv)
}

func _MetadataService_RegisterFileCreation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).RegisterFileCreation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.MetadataService/RegisterFileCreation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).RegisterFileCreation(ctx, req.(*RecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_GetLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).GetLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.MetadataService/GetLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).GetLocation(ctx, req.(*LocRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_OpenDir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenDirRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).OpenDir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.MetadataService/OpenDir",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).OpenDir(ctx, req.(*OpenDirRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_ReadDir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadDirRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).ReadDir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.MetadataService/ReadDir",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).ReadDir(ctx, req.(*ReadDirRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_ReadDirAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadDirRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).ReadDirAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.MetadataService/ReadDirAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).ReadDirAll(ctx, req.(*ReadDirRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_MkDir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MkDirRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).MkDir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.MetadataService/MkDir",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).MkDir(ctx, req.(*MkDirRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataService_DeleteAllData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAllDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServiceServer).DeleteAllData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metadata.MetadataService/DeleteAllData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServiceServer).DeleteAllData(ctx, req.(*DeleteAllDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MetadataService_ServiceDesc is the grpc.ServiceDesc for MetadataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetadataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "metadata.MetadataService",
	HandlerType: (*MetadataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterFileCreation",
			Handler:    _MetadataService_RegisterFileCreation_Handler,
		},
		{
			MethodName: "GetLocation",
			Handler:    _MetadataService_GetLocation_Handler,
		},
		{
			MethodName: "OpenDir",
			Handler:    _MetadataService_OpenDir_Handler,
		},
		{
			MethodName: "ReadDir",
			Handler:    _MetadataService_ReadDir_Handler,
		},
		{
			MethodName: "ReadDirAll",
			Handler:    _MetadataService_ReadDirAll_Handler,
		},
		{
			MethodName: "MkDir",
			Handler:    _MetadataService_MkDir_Handler,
		},
		{
			MethodName: "DeleteAllData",
			Handler:    _MetadataService_DeleteAllData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metadata.proto",
}
