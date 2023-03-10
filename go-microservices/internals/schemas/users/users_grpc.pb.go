// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: protofiles/users.proto

package users

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

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersClient interface {
	LoginAuth(ctx context.Context, in *LoginDTO, opts ...grpc.CallOption) (*ApiResponse, error)
	RegisterAuth(ctx context.Context, in *RegisterDTO, opts ...grpc.CallOption) (*ApiResponse, error)
}

type usersClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersClient(cc grpc.ClientConnInterface) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) LoginAuth(ctx context.Context, in *LoginDTO, opts ...grpc.CallOption) (*ApiResponse, error) {
	out := new(ApiResponse)
	err := c.cc.Invoke(ctx, "/users.Users/LoginAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) RegisterAuth(ctx context.Context, in *RegisterDTO, opts ...grpc.CallOption) (*ApiResponse, error) {
	out := new(ApiResponse)
	err := c.cc.Invoke(ctx, "/users.Users/RegisterAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
// All implementations must embed UnimplementedUsersServer
// for forward compatibility
type UsersServer interface {
	LoginAuth(context.Context, *LoginDTO) (*ApiResponse, error)
	RegisterAuth(context.Context, *RegisterDTO) (*ApiResponse, error)
	mustEmbedUnimplementedUsersServer()
}

// UnimplementedUsersServer must be embedded to have forward compatible implementations.
type UnimplementedUsersServer struct {
}

func (UnimplementedUsersServer) LoginAuth(context.Context, *LoginDTO) (*ApiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginAuth not implemented")
}
func (UnimplementedUsersServer) RegisterAuth(context.Context, *RegisterDTO) (*ApiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAuth not implemented")
}
func (UnimplementedUsersServer) mustEmbedUnimplementedUsersServer() {}

// UnsafeUsersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServer will
// result in compilation errors.
type UnsafeUsersServer interface {
	mustEmbedUnimplementedUsersServer()
}

func RegisterUsersServer(s grpc.ServiceRegistrar, srv UsersServer) {
	s.RegisterService(&Users_ServiceDesc, srv)
}

func _Users_LoginAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginDTO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).LoginAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.Users/LoginAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).LoginAuth(ctx, req.(*LoginDTO))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_RegisterAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterDTO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).RegisterAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.Users/RegisterAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).RegisterAuth(ctx, req.(*RegisterDTO))
	}
	return interceptor(ctx, in, info, handler)
}

// Users_ServiceDesc is the grpc.ServiceDesc for Users service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Users_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "users.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginAuth",
			Handler:    _Users_LoginAuth_Handler,
		},
		{
			MethodName: "RegisterAuth",
			Handler:    _Users_RegisterAuth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protofiles/users.proto",
}
