# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from protos.plugin import common_pb2 as protos_dot_plugin_dot_common__pb2


class DependenciesServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetDependencies = channel.unary_unary(
                '/proto.DependenciesService/GetDependencies',
                request_serializer=protos_dot_plugin_dot_common__pb2.ServiceInput.SerializeToString,
                response_deserializer=protos_dot_plugin_dot_common__pb2.ServiceOutputStringMap.FromString,
                )


class DependenciesServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetDependencies(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_DependenciesServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetDependencies': grpc.unary_unary_rpc_method_handler(
                    servicer.GetDependencies,
                    request_deserializer=protos_dot_plugin_dot_common__pb2.ServiceInput.FromString,
                    response_serializer=protos_dot_plugin_dot_common__pb2.ServiceOutputStringMap.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'proto.DependenciesService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class DependenciesService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetDependencies(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/proto.DependenciesService/GetDependencies',
            protos_dot_plugin_dot_common__pb2.ServiceInput.SerializeToString,
            protos_dot_plugin_dot_common__pb2.ServiceOutputStringMap.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)