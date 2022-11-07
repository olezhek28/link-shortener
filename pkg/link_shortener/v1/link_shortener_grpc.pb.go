// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package linkShortenerV1

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

// LinkShortenerV1Client is the client API for LinkShortenerV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LinkShortenerV1Client interface {
	AddLink(ctx context.Context, in *AddLinkRequest, opts ...grpc.CallOption) (*AddLinkResponse, error)
	GetLongLink(ctx context.Context, in *GetLongLinkRequest, opts ...grpc.CallOption) (*GetLongLinkResponse, error)
}

type linkShortenerV1Client struct {
	cc grpc.ClientConnInterface
}

func NewLinkShortenerV1Client(cc grpc.ClientConnInterface) LinkShortenerV1Client {
	return &linkShortenerV1Client{cc}
}

func (c *linkShortenerV1Client) AddLink(ctx context.Context, in *AddLinkRequest, opts ...grpc.CallOption) (*AddLinkResponse, error) {
	out := new(AddLinkResponse)
	err := c.cc.Invoke(ctx, "/api.link_shortener.v1.LinkShortenerV1/AddLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkShortenerV1Client) GetLongLink(ctx context.Context, in *GetLongLinkRequest, opts ...grpc.CallOption) (*GetLongLinkResponse, error) {
	out := new(GetLongLinkResponse)
	err := c.cc.Invoke(ctx, "/api.link_shortener.v1.LinkShortenerV1/GetLongLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LinkShortenerV1Server is the server API for LinkShortenerV1 service.
// All implementations must embed UnimplementedLinkShortenerV1Server
// for forward compatibility
type LinkShortenerV1Server interface {
	AddLink(context.Context, *AddLinkRequest) (*AddLinkResponse, error)
	GetLongLink(context.Context, *GetLongLinkRequest) (*GetLongLinkResponse, error)
	mustEmbedUnimplementedLinkShortenerV1Server()
}

// UnimplementedLinkShortenerV1Server must be embedded to have forward compatible implementations.
type UnimplementedLinkShortenerV1Server struct {
}

func (UnimplementedLinkShortenerV1Server) AddLink(context.Context, *AddLinkRequest) (*AddLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLink not implemented")
}
func (UnimplementedLinkShortenerV1Server) GetLongLink(context.Context, *GetLongLinkRequest) (*GetLongLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLongLink not implemented")
}
func (UnimplementedLinkShortenerV1Server) mustEmbedUnimplementedLinkShortenerV1Server() {}

// UnsafeLinkShortenerV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LinkShortenerV1Server will
// result in compilation errors.
type UnsafeLinkShortenerV1Server interface {
	mustEmbedUnimplementedLinkShortenerV1Server()
}

func RegisterLinkShortenerV1Server(s grpc.ServiceRegistrar, srv LinkShortenerV1Server) {
	s.RegisterService(&LinkShortenerV1_ServiceDesc, srv)
}

func _LinkShortenerV1_AddLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkShortenerV1Server).AddLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.link_shortener.v1.LinkShortenerV1/AddLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkShortenerV1Server).AddLink(ctx, req.(*AddLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkShortenerV1_GetLongLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLongLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkShortenerV1Server).GetLongLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.link_shortener.v1.LinkShortenerV1/GetLongLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkShortenerV1Server).GetLongLink(ctx, req.(*GetLongLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LinkShortenerV1_ServiceDesc is the grpc.ServiceDesc for LinkShortenerV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LinkShortenerV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.link_shortener.v1.LinkShortenerV1",
	HandlerType: (*LinkShortenerV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddLink",
			Handler:    _LinkShortenerV1_AddLink_Handler,
		},
		{
			MethodName: "GetLongLink",
			Handler:    _LinkShortenerV1_GetLongLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "link_shortener.proto",
}
