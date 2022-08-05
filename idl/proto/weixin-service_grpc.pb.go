// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// WeixinServiceClient is the client API for WeixinService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WeixinServiceClient interface {
	//Ping接口
	Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingRsp, error)
	//群发文章消息
	SendNews(ctx context.Context, in *SendNewsReq, opts ...grpc.CallOption) (*SendNewsResp, error)
}

type weixinServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWeixinServiceClient(cc grpc.ClientConnInterface) WeixinServiceClient {
	return &weixinServiceClient{cc}
}

func (c *weixinServiceClient) Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingRsp, error) {
	out := new(PingRsp)
	err := c.cc.Invoke(ctx, "/WeixinService.WeixinService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weixinServiceClient) SendNews(ctx context.Context, in *SendNewsReq, opts ...grpc.CallOption) (*SendNewsResp, error) {
	out := new(SendNewsResp)
	err := c.cc.Invoke(ctx, "/WeixinService.WeixinService/SendNews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WeixinServiceServer is the server API for WeixinService service.
// All implementations should embed UnimplementedWeixinServiceServer
// for forward compatibility
type WeixinServiceServer interface {
	//Ping接口
	Ping(context.Context, *PingReq) (*PingRsp, error)
	//群发文章消息
	SendNews(context.Context, *SendNewsReq) (*SendNewsResp, error)
}

// UnimplementedWeixinServiceServer should be embedded to have forward compatible implementations.
type UnimplementedWeixinServiceServer struct {
}

func (UnimplementedWeixinServiceServer) Ping(context.Context, *PingReq) (*PingRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedWeixinServiceServer) SendNews(context.Context, *SendNewsReq) (*SendNewsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendNews not implemented")
}

// UnsafeWeixinServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WeixinServiceServer will
// result in compilation errors.
type UnsafeWeixinServiceServer interface {
	mustEmbedUnimplementedWeixinServiceServer()
}

func RegisterWeixinServiceServer(s grpc.ServiceRegistrar, srv WeixinServiceServer) {
	s.RegisterService(&WeixinService_ServiceDesc, srv)
}

func _WeixinService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeixinServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WeixinService.WeixinService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeixinServiceServer).Ping(ctx, req.(*PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WeixinService_SendNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendNewsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeixinServiceServer).SendNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WeixinService.WeixinService/SendNews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeixinServiceServer).SendNews(ctx, req.(*SendNewsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// WeixinService_ServiceDesc is the grpc.ServiceDesc for WeixinService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WeixinService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "WeixinService.WeixinService",
	HandlerType: (*WeixinServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _WeixinService_Ping_Handler,
		},
		{
			MethodName: "SendNews",
			Handler:    _WeixinService_SendNews_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "weixin-service.proto",
}
