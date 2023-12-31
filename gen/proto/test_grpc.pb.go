// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: test.proto

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

const (
	TestApi_Encrypt_FullMethodName     = "/main.TestApi/Encrypt"
	TestApi_Decrypt_FullMethodName     = "/main.TestApi/Decrypt"
	TestApi_GetFiles_FullMethodName    = "/main.TestApi/GetFiles"
	TestApi_GetFileInfo_FullMethodName = "/main.TestApi/GetFileInfo"
	TestApi_UploadFile_FullMethodName  = "/main.TestApi/UploadFile"
	TestApi_UpdateFile_FullMethodName  = "/main.TestApi/UpdateFile"
	TestApi_DeleteFile_FullMethodName  = "/main.TestApi/DeleteFile"
)

// TestApiClient is the client API for TestApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestApiClient interface {
	Encrypt(ctx context.Context, in *Crypt, opts ...grpc.CallOption) (*Crypt, error)
	Decrypt(ctx context.Context, in *Crypt, opts ...grpc.CallOption) (*Crypt, error)
	GetFiles(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*FilesResponse, error)
	GetFileInfo(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*File, error)
	UploadFile(ctx context.Context, in *UploadFileRequest, opts ...grpc.CallOption) (*EmptyMessage, error)
	UpdateFile(ctx context.Context, in *UpdateFileRequest, opts ...grpc.CallOption) (*EmptyMessage, error)
	DeleteFile(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*EmptyMessage, error)
}

type testApiClient struct {
	cc grpc.ClientConnInterface
}

func NewTestApiClient(cc grpc.ClientConnInterface) TestApiClient {
	return &testApiClient{cc}
}

func (c *testApiClient) Encrypt(ctx context.Context, in *Crypt, opts ...grpc.CallOption) (*Crypt, error) {
	out := new(Crypt)
	err := c.cc.Invoke(ctx, TestApi_Encrypt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testApiClient) Decrypt(ctx context.Context, in *Crypt, opts ...grpc.CallOption) (*Crypt, error) {
	out := new(Crypt)
	err := c.cc.Invoke(ctx, TestApi_Decrypt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testApiClient) GetFiles(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*FilesResponse, error) {
	out := new(FilesResponse)
	err := c.cc.Invoke(ctx, TestApi_GetFiles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testApiClient) GetFileInfo(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*File, error) {
	out := new(File)
	err := c.cc.Invoke(ctx, TestApi_GetFileInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testApiClient) UploadFile(ctx context.Context, in *UploadFileRequest, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := c.cc.Invoke(ctx, TestApi_UploadFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testApiClient) UpdateFile(ctx context.Context, in *UpdateFileRequest, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := c.cc.Invoke(ctx, TestApi_UpdateFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testApiClient) DeleteFile(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := c.cc.Invoke(ctx, TestApi_DeleteFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestApiServer is the server API for TestApi service.
// All implementations must embed UnimplementedTestApiServer
// for forward compatibility
type TestApiServer interface {
	Encrypt(context.Context, *Crypt) (*Crypt, error)
	Decrypt(context.Context, *Crypt) (*Crypt, error)
	GetFiles(context.Context, *EmptyMessage) (*FilesResponse, error)
	GetFileInfo(context.Context, *IdRequest) (*File, error)
	UploadFile(context.Context, *UploadFileRequest) (*EmptyMessage, error)
	UpdateFile(context.Context, *UpdateFileRequest) (*EmptyMessage, error)
	DeleteFile(context.Context, *IdRequest) (*EmptyMessage, error)
	mustEmbedUnimplementedTestApiServer()
}

// UnimplementedTestApiServer must be embedded to have forward compatible implementations.
type UnimplementedTestApiServer struct {
}

func (UnimplementedTestApiServer) Encrypt(context.Context, *Crypt) (*Crypt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Encrypt not implemented")
}
func (UnimplementedTestApiServer) Decrypt(context.Context, *Crypt) (*Crypt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decrypt not implemented")
}
func (UnimplementedTestApiServer) GetFiles(context.Context, *EmptyMessage) (*FilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFiles not implemented")
}
func (UnimplementedTestApiServer) GetFileInfo(context.Context, *IdRequest) (*File, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileInfo not implemented")
}
func (UnimplementedTestApiServer) UploadFile(context.Context, *UploadFileRequest) (*EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedTestApiServer) UpdateFile(context.Context, *UpdateFileRequest) (*EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFile not implemented")
}
func (UnimplementedTestApiServer) DeleteFile(context.Context, *IdRequest) (*EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFile not implemented")
}
func (UnimplementedTestApiServer) mustEmbedUnimplementedTestApiServer() {}

// UnsafeTestApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestApiServer will
// result in compilation errors.
type UnsafeTestApiServer interface {
	mustEmbedUnimplementedTestApiServer()
}

func RegisterTestApiServer(s grpc.ServiceRegistrar, srv TestApiServer) {
	s.RegisterService(&TestApi_ServiceDesc, srv)
}

func _TestApi_Encrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Crypt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiServer).Encrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestApi_Encrypt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiServer).Encrypt(ctx, req.(*Crypt))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestApi_Decrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Crypt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiServer).Decrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestApi_Decrypt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiServer).Decrypt(ctx, req.(*Crypt))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestApi_GetFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiServer).GetFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestApi_GetFiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiServer).GetFiles(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestApi_GetFileInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiServer).GetFileInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestApi_GetFileInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiServer).GetFileInfo(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestApi_UploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiServer).UploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestApi_UploadFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiServer).UploadFile(ctx, req.(*UploadFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestApi_UpdateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiServer).UpdateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestApi_UpdateFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiServer).UpdateFile(ctx, req.(*UpdateFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestApi_DeleteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiServer).DeleteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestApi_DeleteFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiServer).DeleteFile(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TestApi_ServiceDesc is the grpc.ServiceDesc for TestApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.TestApi",
	HandlerType: (*TestApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Encrypt",
			Handler:    _TestApi_Encrypt_Handler,
		},
		{
			MethodName: "Decrypt",
			Handler:    _TestApi_Decrypt_Handler,
		},
		{
			MethodName: "GetFiles",
			Handler:    _TestApi_GetFiles_Handler,
		},
		{
			MethodName: "GetFileInfo",
			Handler:    _TestApi_GetFileInfo_Handler,
		},
		{
			MethodName: "UploadFile",
			Handler:    _TestApi_UploadFile_Handler,
		},
		{
			MethodName: "UpdateFile",
			Handler:    _TestApi_UpdateFile_Handler,
		},
		{
			MethodName: "DeleteFile",
			Handler:    _TestApi_DeleteFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}
