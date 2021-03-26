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

// FrameworkServiceClient is the client API for FrameworkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FrameworkServiceClient interface {
	GetVersionName(ctx context.Context, in *ServiceEmpty, opts ...grpc.CallOption) (*ServiceOutputString, error)
	GetSemver(ctx context.Context, in *ServiceEmpty, opts ...grpc.CallOption) (*ServiceOutputString, error)
	Detect(ctx context.Context, in *ServiceInput, opts ...grpc.CallOption) (*ServiceOutputBool, error)
	IsFrameworkFound(ctx context.Context, in *ServiceInput, opts ...grpc.CallOption) (*ServiceOutputBool, error)
	IsFrameworkUsed(ctx context.Context, in *ServiceInput, opts ...grpc.CallOption) (*ServiceOutputBool, error)
	PercentOfFrameworkUsed(ctx context.Context, in *ServiceInput, opts ...grpc.CallOption) (*ServiceOutputInt, error)
	GetFrameworkName(ctx context.Context, in *ServiceEmpty, opts ...grpc.CallOption) (*ServiceOutputString, error)
}

type frameworkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFrameworkServiceClient(cc grpc.ClientConnInterface) FrameworkServiceClient {
	return &frameworkServiceClient{cc}
}

func (c *frameworkServiceClient) GetVersionName(ctx context.Context, in *ServiceEmpty, opts ...grpc.CallOption) (*ServiceOutputString, error) {
	out := new(ServiceOutputString)
	err := c.cc.Invoke(ctx, "/proto.FrameworkService/GetVersionName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frameworkServiceClient) GetSemver(ctx context.Context, in *ServiceEmpty, opts ...grpc.CallOption) (*ServiceOutputString, error) {
	out := new(ServiceOutputString)
	err := c.cc.Invoke(ctx, "/proto.FrameworkService/GetSemver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frameworkServiceClient) Detect(ctx context.Context, in *ServiceInput, opts ...grpc.CallOption) (*ServiceOutputBool, error) {
	out := new(ServiceOutputBool)
	err := c.cc.Invoke(ctx, "/proto.FrameworkService/Detect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frameworkServiceClient) IsFrameworkFound(ctx context.Context, in *ServiceInput, opts ...grpc.CallOption) (*ServiceOutputBool, error) {
	out := new(ServiceOutputBool)
	err := c.cc.Invoke(ctx, "/proto.FrameworkService/IsFrameworkFound", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frameworkServiceClient) IsFrameworkUsed(ctx context.Context, in *ServiceInput, opts ...grpc.CallOption) (*ServiceOutputBool, error) {
	out := new(ServiceOutputBool)
	err := c.cc.Invoke(ctx, "/proto.FrameworkService/IsFrameworkUsed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frameworkServiceClient) PercentOfFrameworkUsed(ctx context.Context, in *ServiceInput, opts ...grpc.CallOption) (*ServiceOutputInt, error) {
	out := new(ServiceOutputInt)
	err := c.cc.Invoke(ctx, "/proto.FrameworkService/PercentOfFrameworkUsed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frameworkServiceClient) GetFrameworkName(ctx context.Context, in *ServiceEmpty, opts ...grpc.CallOption) (*ServiceOutputString, error) {
	out := new(ServiceOutputString)
	err := c.cc.Invoke(ctx, "/proto.FrameworkService/GetFrameworkName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FrameworkServiceServer is the server API for FrameworkService service.
// All implementations must embed UnimplementedFrameworkServiceServer
// for forward compatibility
type FrameworkServiceServer interface {
	GetVersionName(context.Context, *ServiceEmpty) (*ServiceOutputString, error)
	GetSemver(context.Context, *ServiceEmpty) (*ServiceOutputString, error)
	Detect(context.Context, *ServiceInput) (*ServiceOutputBool, error)
	IsFrameworkFound(context.Context, *ServiceInput) (*ServiceOutputBool, error)
	IsFrameworkUsed(context.Context, *ServiceInput) (*ServiceOutputBool, error)
	PercentOfFrameworkUsed(context.Context, *ServiceInput) (*ServiceOutputInt, error)
	GetFrameworkName(context.Context, *ServiceEmpty) (*ServiceOutputString, error)
	//mustEmbedUnimplementedFrameworkServiceServer()
}

// UnimplementedFrameworkServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFrameworkServiceServer struct {
}

func (UnimplementedFrameworkServiceServer) GetVersionName(context.Context, *ServiceEmpty) (*ServiceOutputString, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersionName not implemented")
}
func (UnimplementedFrameworkServiceServer) GetSemver(context.Context, *ServiceEmpty) (*ServiceOutputString, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSemver not implemented")
}
func (UnimplementedFrameworkServiceServer) Detect(context.Context, *ServiceInput) (*ServiceOutputBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Detect not implemented")
}
func (UnimplementedFrameworkServiceServer) IsFrameworkFound(context.Context, *ServiceInput) (*ServiceOutputBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsFrameworkFound not implemented")
}
func (UnimplementedFrameworkServiceServer) IsFrameworkUsed(context.Context, *ServiceInput) (*ServiceOutputBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsFrameworkUsed not implemented")
}
func (UnimplementedFrameworkServiceServer) PercentOfFrameworkUsed(context.Context, *ServiceInput) (*ServiceOutputInt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PercentOfFrameworkUsed not implemented")
}
func (UnimplementedFrameworkServiceServer) GetFrameworkName(context.Context, *ServiceEmpty) (*ServiceOutputString, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFrameworkName not implemented")
}
func (UnimplementedFrameworkServiceServer) mustEmbedUnimplementedFrameworkServiceServer() {}

// UnsafeFrameworkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FrameworkServiceServer will
// result in compilation errors.
type UnsafeFrameworkServiceServer interface {
	mustEmbedUnimplementedFrameworkServiceServer()
}

func RegisterFrameworkServiceServer(s grpc.ServiceRegistrar, srv FrameworkServiceServer) {
	s.RegisterService(&FrameworkService_ServiceDesc, srv)
}

func _FrameworkService_GetVersionName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrameworkServiceServer).GetVersionName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FrameworkService/GetVersionName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrameworkServiceServer).GetVersionName(ctx, req.(*ServiceEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FrameworkService_GetSemver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrameworkServiceServer).GetSemver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FrameworkService/GetSemver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrameworkServiceServer).GetSemver(ctx, req.(*ServiceEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FrameworkService_Detect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrameworkServiceServer).Detect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FrameworkService/Detect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrameworkServiceServer).Detect(ctx, req.(*ServiceInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _FrameworkService_IsFrameworkFound_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrameworkServiceServer).IsFrameworkFound(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FrameworkService/IsFrameworkFound",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrameworkServiceServer).IsFrameworkFound(ctx, req.(*ServiceInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _FrameworkService_IsFrameworkUsed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrameworkServiceServer).IsFrameworkUsed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FrameworkService/IsFrameworkUsed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrameworkServiceServer).IsFrameworkUsed(ctx, req.(*ServiceInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _FrameworkService_PercentOfFrameworkUsed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrameworkServiceServer).PercentOfFrameworkUsed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FrameworkService/PercentOfFrameworkUsed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrameworkServiceServer).PercentOfFrameworkUsed(ctx, req.(*ServiceInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _FrameworkService_GetFrameworkName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrameworkServiceServer).GetFrameworkName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FrameworkService/GetFrameworkName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrameworkServiceServer).GetFrameworkName(ctx, req.(*ServiceEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

// FrameworkService_ServiceDesc is the grpc.ServiceDesc for FrameworkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FrameworkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.FrameworkService",
	HandlerType: (*FrameworkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersionName",
			Handler:    _FrameworkService_GetVersionName_Handler,
		},
		{
			MethodName: "GetSemver",
			Handler:    _FrameworkService_GetSemver_Handler,
		},
		{
			MethodName: "Detect",
			Handler:    _FrameworkService_Detect_Handler,
		},
		{
			MethodName: "IsFrameworkFound",
			Handler:    _FrameworkService_IsFrameworkFound_Handler,
		},
		{
			MethodName: "IsFrameworkUsed",
			Handler:    _FrameworkService_IsFrameworkUsed_Handler,
		},
		{
			MethodName: "PercentOfFrameworkUsed",
			Handler:    _FrameworkService_PercentOfFrameworkUsed_Handler,
		},
		{
			MethodName: "GetFrameworkName",
			Handler:    _FrameworkService_GetFrameworkName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/framework.proto",
}
