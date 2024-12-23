// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: parser.proto

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

// ModelParserClient is the client API for ModelParser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModelParserClient interface {
	// Receives a file and returns the parsed structure
	ParseModel(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (*ParseResponse, error)
}

type modelParserClient struct {
	cc grpc.ClientConnInterface
}

func NewModelParserClient(cc grpc.ClientConnInterface) ModelParserClient {
	return &modelParserClient{cc}
}

func (c *modelParserClient) ParseModel(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (*ParseResponse, error) {
	out := new(ParseResponse)
	err := c.cc.Invoke(ctx, "/parser.ModelParser/ParseModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModelParserServer is the server API for ModelParser service.
// All implementations should embed UnimplementedModelParserServer
// for forward compatibility
type ModelParserServer interface {
	// Receives a file and returns the parsed structure
	ParseModel(context.Context, *ParseRequest) (*ParseResponse, error)
}

// UnimplementedModelParserServer should be embedded to have forward compatible implementations.
type UnimplementedModelParserServer struct {
}

func (UnimplementedModelParserServer) ParseModel(context.Context, *ParseRequest) (*ParseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseModel not implemented")
}

// UnsafeModelParserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModelParserServer will
// result in compilation errors.
type UnsafeModelParserServer interface {
	mustEmbedUnimplementedModelParserServer()
}

func RegisterModelParserServer(s grpc.ServiceRegistrar, srv ModelParserServer) {
	s.RegisterService(&ModelParser_ServiceDesc, srv)
}

func _ModelParser_ParseModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelParserServer).ParseModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parser.ModelParser/ParseModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelParserServer).ParseModel(ctx, req.(*ParseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ModelParser_ServiceDesc is the grpc.ServiceDesc for ModelParser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModelParser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "parser.ModelParser",
	HandlerType: (*ModelParserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ParseModel",
			Handler:    _ModelParser_ParseModel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "parser.proto",
}
