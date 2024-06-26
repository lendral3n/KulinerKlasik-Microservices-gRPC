// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.0
// source: menu.proto

package pb

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

// MenuServiceClient is the client API for MenuService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MenuServiceClient interface {
	CreateMenu(ctx context.Context, in *CreateMenuRequest, opts ...grpc.CallOption) (*CreateMenuResponse, error)
	FindOne(ctx context.Context, in *FindOneRequest, opts ...grpc.CallOption) (*FindOneResponse, error)
	DecreaseStock(ctx context.Context, in *DecreaseStockRequest, opts ...grpc.CallOption) (*DecreaseStockResponse, error)
}

type menuServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMenuServiceClient(cc grpc.ClientConnInterface) MenuServiceClient {
	return &menuServiceClient{cc}
}

func (c *menuServiceClient) CreateMenu(ctx context.Context, in *CreateMenuRequest, opts ...grpc.CallOption) (*CreateMenuResponse, error) {
	out := new(CreateMenuResponse)
	err := c.cc.Invoke(ctx, "/menu.MenuService/CreateMenu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) FindOne(ctx context.Context, in *FindOneRequest, opts ...grpc.CallOption) (*FindOneResponse, error) {
	out := new(FindOneResponse)
	err := c.cc.Invoke(ctx, "/menu.MenuService/FindOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) DecreaseStock(ctx context.Context, in *DecreaseStockRequest, opts ...grpc.CallOption) (*DecreaseStockResponse, error) {
	out := new(DecreaseStockResponse)
	err := c.cc.Invoke(ctx, "/menu.MenuService/DecreaseStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MenuServiceServer is the server API for MenuService service.
// All implementations must embed UnimplementedMenuServiceServer
// for forward compatibility
type MenuServiceServer interface {
	CreateMenu(context.Context, *CreateMenuRequest) (*CreateMenuResponse, error)
	FindOne(context.Context, *FindOneRequest) (*FindOneResponse, error)
	DecreaseStock(context.Context, *DecreaseStockRequest) (*DecreaseStockResponse, error)
	mustEmbedUnimplementedMenuServiceServer()
}

// UnimplementedMenuServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMenuServiceServer struct {
}

func (UnimplementedMenuServiceServer) CreateMenu(context.Context, *CreateMenuRequest) (*CreateMenuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMenu not implemented")
}
func (UnimplementedMenuServiceServer) FindOne(context.Context, *FindOneRequest) (*FindOneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOne not implemented")
}
func (UnimplementedMenuServiceServer) DecreaseStock(context.Context, *DecreaseStockRequest) (*DecreaseStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecreaseStock not implemented")
}
func (UnimplementedMenuServiceServer) mustEmbedUnimplementedMenuServiceServer() {}

// UnsafeMenuServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MenuServiceServer will
// result in compilation errors.
type UnsafeMenuServiceServer interface {
	mustEmbedUnimplementedMenuServiceServer()
}

func RegisterMenuServiceServer(s grpc.ServiceRegistrar, srv MenuServiceServer) {
	s.RegisterService(&MenuService_ServiceDesc, srv)
}

func _MenuService_CreateMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).CreateMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/menu.MenuService/CreateMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).CreateMenu(ctx, req.(*CreateMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_FindOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).FindOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/menu.MenuService/FindOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).FindOne(ctx, req.(*FindOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_DecreaseStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecreaseStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).DecreaseStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/menu.MenuService/DecreaseStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).DecreaseStock(ctx, req.(*DecreaseStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MenuService_ServiceDesc is the grpc.ServiceDesc for MenuService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MenuService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "menu.MenuService",
	HandlerType: (*MenuServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMenu",
			Handler:    _MenuService_CreateMenu_Handler,
		},
		{
			MethodName: "FindOne",
			Handler:    _MenuService_FindOne_Handler,
		},
		{
			MethodName: "DecreaseStock",
			Handler:    _MenuService_DecreaseStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "menu.proto",
}
