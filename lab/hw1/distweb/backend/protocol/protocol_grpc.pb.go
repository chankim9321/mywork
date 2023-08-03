// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: backend/protocol/protocol.proto

package protocol_update

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

// FileTransferServiceClient is the client API for FileTransferService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileTransferServiceClient interface {
	GetFile(ctx context.Context, in *FileMetadata, opts ...grpc.CallOption) (FileTransferService_GetFileClient, error)
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (FileTransferService_UploadFileClient, error)
	GetFileList(ctx context.Context, in *DispatchInfo, opts ...grpc.CallOption) (*FileInfoList, error)
}

type fileTransferServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileTransferServiceClient(cc grpc.ClientConnInterface) FileTransferServiceClient {
	return &fileTransferServiceClient{cc}
}

func (c *fileTransferServiceClient) GetFile(ctx context.Context, in *FileMetadata, opts ...grpc.CallOption) (FileTransferService_GetFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileTransferService_ServiceDesc.Streams[0], "/protocol.FileTransferService/GetFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileTransferServiceGetFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileTransferService_GetFileClient interface {
	Recv() (*FileChunk, error)
	grpc.ClientStream
}

type fileTransferServiceGetFileClient struct {
	grpc.ClientStream
}

func (x *fileTransferServiceGetFileClient) Recv() (*FileChunk, error) {
	m := new(FileChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileTransferServiceClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (FileTransferService_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileTransferService_ServiceDesc.Streams[1], "/protocol.FileTransferService/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileTransferServiceUploadFileClient{stream}
	return x, nil
}

type FileTransferService_UploadFileClient interface {
	Send(*FileChunk) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type fileTransferServiceUploadFileClient struct {
	grpc.ClientStream
}

func (x *fileTransferServiceUploadFileClient) Send(m *FileChunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileTransferServiceUploadFileClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileTransferServiceClient) GetFileList(ctx context.Context, in *DispatchInfo, opts ...grpc.CallOption) (*FileInfoList, error) {
	out := new(FileInfoList)
	err := c.cc.Invoke(ctx, "/protocol.FileTransferService/GetFileList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileTransferServiceServer is the server API for FileTransferService service.
// All implementations must embed UnimplementedFileTransferServiceServer
// for forward compatibility
type FileTransferServiceServer interface {
	GetFile(*FileMetadata, FileTransferService_GetFileServer) error
	UploadFile(FileTransferService_UploadFileServer) error
	GetFileList(context.Context, *DispatchInfo) (*FileInfoList, error)
	mustEmbedUnimplementedFileTransferServiceServer()
}

// UnimplementedFileTransferServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFileTransferServiceServer struct {
}

func (UnimplementedFileTransferServiceServer) GetFile(*FileMetadata, FileTransferService_GetFileServer) error {
	return status.Errorf(codes.Unimplemented, "method GetFile not implemented")
}
func (UnimplementedFileTransferServiceServer) UploadFile(FileTransferService_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedFileTransferServiceServer) GetFileList(context.Context, *DispatchInfo) (*FileInfoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileList not implemented")
}
func (UnimplementedFileTransferServiceServer) mustEmbedUnimplementedFileTransferServiceServer() {}

// UnsafeFileTransferServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileTransferServiceServer will
// result in compilation errors.
type UnsafeFileTransferServiceServer interface {
	mustEmbedUnimplementedFileTransferServiceServer()
}

func RegisterFileTransferServiceServer(s grpc.ServiceRegistrar, srv FileTransferServiceServer) {
	s.RegisterService(&FileTransferService_ServiceDesc, srv)
}

func _FileTransferService_GetFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FileMetadata)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileTransferServiceServer).GetFile(m, &fileTransferServiceGetFileServer{stream})
}

type FileTransferService_GetFileServer interface {
	Send(*FileChunk) error
	grpc.ServerStream
}

type fileTransferServiceGetFileServer struct {
	grpc.ServerStream
}

func (x *fileTransferServiceGetFileServer) Send(m *FileChunk) error {
	return x.ServerStream.SendMsg(m)
}

func _FileTransferService_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileTransferServiceServer).UploadFile(&fileTransferServiceUploadFileServer{stream})
}

type FileTransferService_UploadFileServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*FileChunk, error)
	grpc.ServerStream
}

type fileTransferServiceUploadFileServer struct {
	grpc.ServerStream
}

func (x *fileTransferServiceUploadFileServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileTransferServiceUploadFileServer) Recv() (*FileChunk, error) {
	m := new(FileChunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FileTransferService_GetFileList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DispatchInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileTransferServiceServer).GetFileList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.FileTransferService/GetFileList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileTransferServiceServer).GetFileList(ctx, req.(*DispatchInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// FileTransferService_ServiceDesc is the grpc.ServiceDesc for FileTransferService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileTransferService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.FileTransferService",
	HandlerType: (*FileTransferServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFileList",
			Handler:    _FileTransferService_GetFileList_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetFile",
			Handler:       _FileTransferService_GetFile_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UploadFile",
			Handler:       _FileTransferService_UploadFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "backend/protocol/protocol.proto",
}
