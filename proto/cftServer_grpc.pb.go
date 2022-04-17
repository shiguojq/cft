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

// CftServerClient is the client API for CftServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CftServerClient interface {
	HandleHeartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartBeatReply, error)
	HandleSnapshot(ctx context.Context, in *SnapshotRequest, opts ...grpc.CallOption) (*SnapshotReply, error)
}

type cftServerClient struct {
	cc grpc.ClientConnInterface
}

func NewCftServerClient(cc grpc.ClientConnInterface) CftServerClient {
	return &cftServerClient{cc}
}

func (c *cftServerClient) HandleHeartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartBeatReply, error) {
	out := new(HeartBeatReply)
	err := c.cc.Invoke(ctx, "/proto.CftServer/HandleHeartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cftServerClient) HandleSnapshot(ctx context.Context, in *SnapshotRequest, opts ...grpc.CallOption) (*SnapshotReply, error) {
	out := new(SnapshotReply)
	err := c.cc.Invoke(ctx, "/proto.CftServer/HandleSnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CftServerServer is the server API for CftServer service.
// All implementations must embed UnimplementedCftServerServer
// for forward compatibility
type CftServerServer interface {
	HandleHeartbeat(context.Context, *HeartbeatRequest) (*HeartBeatReply, error)
	HandleSnapshot(context.Context, *SnapshotRequest) (*SnapshotReply, error)
	mustEmbedUnimplementedCftServerServer()
}

// UnimplementedCftServerServer must be embedded to have forward compatible implementations.
type UnimplementedCftServerServer struct {
}

func (UnimplementedCftServerServer) HandleHeartbeat(context.Context, *HeartbeatRequest) (*HeartBeatReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleHeartbeat not implemented")
}
func (UnimplementedCftServerServer) HandleSnapshot(context.Context, *SnapshotRequest) (*SnapshotReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleSnapshot not implemented")
}
func (UnimplementedCftServerServer) mustEmbedUnimplementedCftServerServer() {}

// UnsafeCftServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CftServerServer will
// result in compilation errors.
type UnsafeCftServerServer interface {
	mustEmbedUnimplementedCftServerServer()
}

func RegisterCftServerServer(s grpc.ServiceRegistrar, srv CftServerServer) {
	s.RegisterService(&CftServer_ServiceDesc, srv)
}

func _CftServer_HandleHeartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CftServerServer).HandleHeartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CftServer/HandleHeartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CftServerServer).HandleHeartbeat(ctx, req.(*HeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CftServer_HandleSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CftServerServer).HandleSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CftServer/HandleSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CftServerServer).HandleSnapshot(ctx, req.(*SnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CftServer_ServiceDesc is the grpc.ServiceDesc for CftServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CftServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CftServer",
	HandlerType: (*CftServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleHeartbeat",
			Handler:    _CftServer_HandleHeartbeat_Handler,
		},
		{
			MethodName: "HandleSnapshot",
			Handler:    _CftServer_HandleSnapshot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/cftServer.proto",
}
