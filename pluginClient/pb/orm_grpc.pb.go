// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// OrmServiceClient is the client API for OrmService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrmServiceClient interface {
	Detect(ctx context.Context, in *ServiceInput, opts ...grpc.CallOption) (*ServiceOutputDetectOrm, error)
	IsORMUsed(ctx context.Context, in *ServiceInput, opts ...grpc.CallOption) (*ServiceOutputBool, error)
	PercentOfORMUsed(ctx context.Context, in *ServiceInput, opts ...grpc.CallOption) (*ServiceOutputFloat, error)
}

type ormServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrmServiceClient(cc grpc.ClientConnInterface) OrmServiceClient {
	return &ormServiceClient{cc}
}

func (c *ormServiceClient) Detect(ctx context.Context, in *ServiceInput, opts ...grpc.CallOption) (*ServiceOutputDetectOrm, error) {
	out := new(ServiceOutputDetectOrm)
	err := c.cc.Invoke(ctx, "/proto.OrmService/Detect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ormServiceClient) IsORMUsed(ctx context.Context, in *ServiceInput, opts ...grpc.CallOption) (*ServiceOutputBool, error) {
	out := new(ServiceOutputBool)
	err := c.cc.Invoke(ctx, "/proto.OrmService/IsORMUsed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ormServiceClient) PercentOfORMUsed(ctx context.Context, in *ServiceInput, opts ...grpc.CallOption) (*ServiceOutputFloat, error) {
	out := new(ServiceOutputFloat)
	err := c.cc.Invoke(ctx, "/proto.OrmService/PercentOfORMUsed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrmServiceServer is the server API for OrmService service.
// All implementations must embed UnimplementedOrmServiceServer
// for forward compatibility
type OrmServiceServer interface {
	Detect(context.Context, *ServiceInput) (*ServiceOutputDetectOrm, error)
	IsORMUsed(context.Context, *ServiceInput) (*ServiceOutputBool, error)
	PercentOfORMUsed(context.Context, *ServiceInput) (*ServiceOutputFloat, error)
	//mustEmbedUnimplementedOrmServiceServer()
}

// UnimplementedOrmServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrmServiceServer struct {
}

func (UnimplementedOrmServiceServer) Detect(context.Context, *ServiceInput) (*ServiceOutputDetectOrm, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Detect not implemented")
}
func (UnimplementedOrmServiceServer) IsORMUsed(context.Context, *ServiceInput) (*ServiceOutputBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsORMUsed not implemented")
}
func (UnimplementedOrmServiceServer) PercentOfORMUsed(context.Context, *ServiceInput) (*ServiceOutputFloat, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PercentOfORMUsed not implemented")
}
func (UnimplementedOrmServiceServer) mustEmbedUnimplementedOrmServiceServer() {}

// UnsafeOrmServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrmServiceServer will
// result in compilation errors.
type UnsafeOrmServiceServer interface {
	mustEmbedUnimplementedOrmServiceServer()
}

func RegisterOrmServiceServer(s grpc.ServiceRegistrar, srv OrmServiceServer) {
	s.RegisterService(&OrmService_ServiceDesc, srv)
}

func _OrmService_Detect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrmServiceServer).Detect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OrmService/Detect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrmServiceServer).Detect(ctx, req.(*ServiceInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrmService_IsORMUsed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrmServiceServer).IsORMUsed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OrmService/IsORMUsed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrmServiceServer).IsORMUsed(ctx, req.(*ServiceInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrmService_PercentOfORMUsed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrmServiceServer).PercentOfORMUsed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OrmService/PercentOfORMUsed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrmServiceServer).PercentOfORMUsed(ctx, req.(*ServiceInput))
	}
	return interceptor(ctx, in, info, handler)
}

// OrmService_ServiceDesc is the grpc.ServiceDesc for OrmService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrmService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.OrmService",
	HandlerType: (*OrmServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Detect",
			Handler:    _OrmService_Detect_Handler,
		},
		{
			MethodName: "IsORMUsed",
			Handler:    _OrmService_IsORMUsed_Handler,
		},
		{
			MethodName: "PercentOfORMUsed",
			Handler:    _OrmService_PercentOfORMUsed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/orm.proto",
}
