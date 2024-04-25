// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: proto/Auth/ResetAuthData.proto

package reset_auth_data

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

// ResetAuthDataClient is the client API for ResetAuthData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResetAuthDataClient interface {
	ResetPassword(ctx context.Context, in *ResetPasswordRequests, opts ...grpc.CallOption) (*ResetResponse, error)
	ResetEmail(ctx context.Context, in *ResetEmailRequests, opts ...grpc.CallOption) (*ResetResponse, error)
	ResetSteamId(ctx context.Context, in *ResetSteamIdRequests, opts ...grpc.CallOption) (*ResetResponse, error)
}

type resetAuthDataClient struct {
	cc grpc.ClientConnInterface
}

func NewResetAuthDataClient(cc grpc.ClientConnInterface) ResetAuthDataClient {
	return &resetAuthDataClient{cc}
}

func (c *resetAuthDataClient) ResetPassword(ctx context.Context, in *ResetPasswordRequests, opts ...grpc.CallOption) (*ResetResponse, error) {
	out := new(ResetResponse)
	err := c.cc.Invoke(ctx, "/resetAuthData.ResetAuthData/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resetAuthDataClient) ResetEmail(ctx context.Context, in *ResetEmailRequests, opts ...grpc.CallOption) (*ResetResponse, error) {
	out := new(ResetResponse)
	err := c.cc.Invoke(ctx, "/resetAuthData.ResetAuthData/ResetEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resetAuthDataClient) ResetSteamId(ctx context.Context, in *ResetSteamIdRequests, opts ...grpc.CallOption) (*ResetResponse, error) {
	out := new(ResetResponse)
	err := c.cc.Invoke(ctx, "/resetAuthData.ResetAuthData/ResetSteamId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResetAuthDataServer is the server API for ResetAuthData service.
// All implementations must embed UnimplementedResetAuthDataServer
// for forward compatibility
type ResetAuthDataServer interface {
	ResetPassword(context.Context, *ResetPasswordRequests) (*ResetResponse, error)
	ResetEmail(context.Context, *ResetEmailRequests) (*ResetResponse, error)
	ResetSteamId(context.Context, *ResetSteamIdRequests) (*ResetResponse, error)
	mustEmbedUnimplementedResetAuthDataServer()
}

// UnimplementedResetAuthDataServer must be embedded to have forward compatible implementations.
type UnimplementedResetAuthDataServer struct {
}

func (UnimplementedResetAuthDataServer) ResetPassword(context.Context, *ResetPasswordRequests) (*ResetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedResetAuthDataServer) ResetEmail(context.Context, *ResetEmailRequests) (*ResetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetEmail not implemented")
}
func (UnimplementedResetAuthDataServer) ResetSteamId(context.Context, *ResetSteamIdRequests) (*ResetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetSteamId not implemented")
}
func (UnimplementedResetAuthDataServer) mustEmbedUnimplementedResetAuthDataServer() {}

// UnsafeResetAuthDataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResetAuthDataServer will
// result in compilation errors.
type UnsafeResetAuthDataServer interface {
	mustEmbedUnimplementedResetAuthDataServer()
}

func RegisterResetAuthDataServer(s grpc.ServiceRegistrar, srv ResetAuthDataServer) {
	s.RegisterService(&ResetAuthData_ServiceDesc, srv)
}

func _ResetAuthData_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordRequests)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResetAuthDataServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resetAuthData.ResetAuthData/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResetAuthDataServer).ResetPassword(ctx, req.(*ResetPasswordRequests))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResetAuthData_ResetEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetEmailRequests)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResetAuthDataServer).ResetEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resetAuthData.ResetAuthData/ResetEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResetAuthDataServer).ResetEmail(ctx, req.(*ResetEmailRequests))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResetAuthData_ResetSteamId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetSteamIdRequests)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResetAuthDataServer).ResetSteamId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resetAuthData.ResetAuthData/ResetSteamId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResetAuthDataServer).ResetSteamId(ctx, req.(*ResetSteamIdRequests))
	}
	return interceptor(ctx, in, info, handler)
}

// ResetAuthData_ServiceDesc is the grpc.ServiceDesc for ResetAuthData service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResetAuthData_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "resetAuthData.ResetAuthData",
	HandlerType: (*ResetAuthDataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ResetPassword",
			Handler:    _ResetAuthData_ResetPassword_Handler,
		},
		{
			MethodName: "ResetEmail",
			Handler:    _ResetAuthData_ResetEmail_Handler,
		},
		{
			MethodName: "ResetSteamId",
			Handler:    _ResetAuthData_ResetSteamId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/Auth/ResetAuthData.proto",
}
