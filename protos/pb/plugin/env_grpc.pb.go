// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package plugin

import (
	helpers "code-analyser/protos/pb/helpers"
	languageSpecific "code-analyser/protos/pb/output/languageSpecific"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EnvClient is the client API for Env service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EnvClient interface {
	Detect(ctx context.Context, in *helpers.Input, opts ...grpc.CallOption) (*languageSpecific.Envs, error)
}

type envClient struct {
	cc grpc.ClientConnInterface
}

func NewEnvClient(cc grpc.ClientConnInterface) EnvClient {
	return &envClient{cc}
}

func (c *envClient) Detect(ctx context.Context, in *helpers.Input, opts ...grpc.CallOption) (*languageSpecific.Envs, error) {
	out := new(languageSpecific.Envs)
	err := c.cc.Invoke(ctx, "/proto.Env/Detect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnvServer is the server API for Env service.
// All implementations should embed UnimplementedEnvServer
// for forward compatibility
type EnvServer interface {
	Detect(context.Context, *helpers.Input) (*languageSpecific.Envs, error)
}

// UnimplementedEnvServer should be embedded to have forward compatible implementations.
type UnimplementedEnvServer struct {
}

func (UnimplementedEnvServer) Detect(context.Context, *helpers.Input) (*languageSpecific.Envs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Detect not implemented")
}

// UnsafeEnvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EnvServer will
// result in compilation errors.
type UnsafeEnvServer interface {
	mustEmbedUnimplementedEnvServer()
}

func RegisterEnvServer(s grpc.ServiceRegistrar, srv EnvServer) {
	s.RegisterService(&Env_ServiceDesc, srv)
}

func _Env_Detect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(helpers.Input)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvServer).Detect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Env/Detect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvServer).Detect(ctx, req.(*helpers.Input))
	}
	return interceptor(ctx, in, info, handler)
}

// Env_ServiceDesc is the grpc.ServiceDesc for Env service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Env_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Env",
	HandlerType: (*EnvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Detect",
			Handler:    _Env_Detect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/plugin/env.proto",
}