// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.10
// source: bgp/bgp.proto

package bgp

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

// BGPClient is the client API for BGP service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BGPClient interface {
	ClearBGPNeighbor(ctx context.Context, in *ClearBGPNeighborRequest, opts ...grpc.CallOption) (*ClearBGPNeighborResponse, error)
}

type bGPClient struct {
	cc grpc.ClientConnInterface
}

func NewBGPClient(cc grpc.ClientConnInterface) BGPClient {
	return &bGPClient{cc}
}

func (c *bGPClient) ClearBGPNeighbor(ctx context.Context, in *ClearBGPNeighborRequest, opts ...grpc.CallOption) (*ClearBGPNeighborResponse, error) {
	out := new(ClearBGPNeighborResponse)
	err := c.cc.Invoke(ctx, "/gnoi.bgp.BGP/ClearBGPNeighbor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BGPServer is the server API for BGP service.
// All implementations must embed UnimplementedBGPServer
// for forward compatibility
type BGPServer interface {
	ClearBGPNeighbor(context.Context, *ClearBGPNeighborRequest) (*ClearBGPNeighborResponse, error)
	mustEmbedUnimplementedBGPServer()
}

// UnimplementedBGPServer must be embedded to have forward compatible implementations.
type UnimplementedBGPServer struct {
}

func (UnimplementedBGPServer) ClearBGPNeighbor(context.Context, *ClearBGPNeighborRequest) (*ClearBGPNeighborResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearBGPNeighbor not implemented")
}
func (UnimplementedBGPServer) mustEmbedUnimplementedBGPServer() {}

// UnsafeBGPServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BGPServer will
// result in compilation errors.
type UnsafeBGPServer interface {
	mustEmbedUnimplementedBGPServer()
}

func RegisterBGPServer(s grpc.ServiceRegistrar, srv BGPServer) {
	s.RegisterService(&BGP_ServiceDesc, srv)
}

func _BGP_ClearBGPNeighbor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearBGPNeighborRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BGPServer).ClearBGPNeighbor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.bgp.BGP/ClearBGPNeighbor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BGPServer).ClearBGPNeighbor(ctx, req.(*ClearBGPNeighborRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BGP_ServiceDesc is the grpc.ServiceDesc for BGP service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BGP_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gnoi.bgp.BGP",
	HandlerType: (*BGPServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClearBGPNeighbor",
			Handler:    _BGP_ClearBGPNeighbor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bgp/bgp.proto",
}
