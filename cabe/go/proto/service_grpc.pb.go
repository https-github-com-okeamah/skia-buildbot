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

// AnalysisClient is the client API for Analysis service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnalysisClient interface {
	// Get results of a performance experiment analysis.
	GetAnalysis(ctx context.Context, in *GetAnalysisRequest, opts ...grpc.CallOption) (*GetAnalysisResponse, error)
}

type analysisClient struct {
	cc grpc.ClientConnInterface
}

func NewAnalysisClient(cc grpc.ClientConnInterface) AnalysisClient {
	return &analysisClient{cc}
}

func (c *analysisClient) GetAnalysis(ctx context.Context, in *GetAnalysisRequest, opts ...grpc.CallOption) (*GetAnalysisResponse, error) {
	out := new(GetAnalysisResponse)
	err := c.cc.Invoke(ctx, "/cabe.v1.Analysis/GetAnalysis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnalysisServer is the server API for Analysis service.
// All implementations must embed UnimplementedAnalysisServer
// for forward compatibility
type AnalysisServer interface {
	// Get results of a performance experiment analysis.
	GetAnalysis(context.Context, *GetAnalysisRequest) (*GetAnalysisResponse, error)
	mustEmbedUnimplementedAnalysisServer()
}

// UnimplementedAnalysisServer must be embedded to have forward compatible implementations.
type UnimplementedAnalysisServer struct {
}

func (UnimplementedAnalysisServer) GetAnalysis(context.Context, *GetAnalysisRequest) (*GetAnalysisResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnalysis not implemented")
}
func (UnimplementedAnalysisServer) mustEmbedUnimplementedAnalysisServer() {}

// UnsafeAnalysisServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnalysisServer will
// result in compilation errors.
type UnsafeAnalysisServer interface {
	mustEmbedUnimplementedAnalysisServer()
}

func RegisterAnalysisServer(s grpc.ServiceRegistrar, srv AnalysisServer) {
	s.RegisterService(&Analysis_ServiceDesc, srv)
}

func _Analysis_GetAnalysis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAnalysisRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalysisServer).GetAnalysis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cabe.v1.Analysis/GetAnalysis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalysisServer).GetAnalysis(ctx, req.(*GetAnalysisRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Analysis_ServiceDesc is the grpc.ServiceDesc for Analysis service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Analysis_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cabe.v1.Analysis",
	HandlerType: (*AnalysisServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAnalysis",
			Handler:    _Analysis_GetAnalysis_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cabe/proto/v1/service.proto",
}
