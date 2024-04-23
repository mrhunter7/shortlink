// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: infrastructure/rpc/cqrs/link/v1/link_command.proto

package v1

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

const (
	LinkCommandService_Add_FullMethodName    = "/infrastructure.rpc.cqrs.link.v1.LinkCommandService/Add"
	LinkCommandService_Update_FullMethodName = "/infrastructure.rpc.cqrs.link.v1.LinkCommandService/Update"
	LinkCommandService_Delete_FullMethodName = "/infrastructure.rpc.cqrs.link.v1.LinkCommandService/Delete"
)

// LinkCommandServiceClient is the client API for LinkCommandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LinkCommandServiceClient interface {
	// Add adds a new link
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	// Update updates an existing link
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Delete deletes an existing link
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type linkCommandServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLinkCommandServiceClient(cc grpc.ClientConnInterface) LinkCommandServiceClient {
	return &linkCommandServiceClient{cc}
}

func (c *linkCommandServiceClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, LinkCommandService_Add_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkCommandServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, LinkCommandService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkCommandServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, LinkCommandService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LinkCommandServiceServer is the server API for LinkCommandService service.
// All implementations must embed UnimplementedLinkCommandServiceServer
// for forward compatibility
type LinkCommandServiceServer interface {
	// Add adds a new link
	Add(context.Context, *AddRequest) (*AddResponse, error)
	// Update updates an existing link
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Delete deletes an existing link
	Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedLinkCommandServiceServer()
}

// UnimplementedLinkCommandServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLinkCommandServiceServer struct {
}

func (UnimplementedLinkCommandServiceServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedLinkCommandServiceServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedLinkCommandServiceServer) Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedLinkCommandServiceServer) mustEmbedUnimplementedLinkCommandServiceServer() {}

// UnsafeLinkCommandServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LinkCommandServiceServer will
// result in compilation errors.
type UnsafeLinkCommandServiceServer interface {
	mustEmbedUnimplementedLinkCommandServiceServer()
}

func RegisterLinkCommandServiceServer(s grpc.ServiceRegistrar, srv LinkCommandServiceServer) {
	s.RegisterService(&LinkCommandService_ServiceDesc, srv)
}

func _LinkCommandService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkCommandServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkCommandService_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkCommandServiceServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkCommandService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkCommandServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkCommandService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkCommandServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkCommandService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkCommandServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkCommandService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkCommandServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LinkCommandService_ServiceDesc is the grpc.ServiceDesc for LinkCommandService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LinkCommandService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infrastructure.rpc.cqrs.link.v1.LinkCommandService",
	HandlerType: (*LinkCommandServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _LinkCommandService_Add_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _LinkCommandService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _LinkCommandService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "infrastructure/rpc/cqrs/link/v1/link_command.proto",
}