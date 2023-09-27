"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from . import auth_pb2 as auth__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2

class AuthServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.AuthWithPassword = channel.unary_unary('/panels.auth.v1.AuthService/AuthWithPassword', request_serializer=auth__pb2.PasswordAuthRequest.SerializeToString, response_deserializer=auth__pb2.AuthToken.FromString)
        self.SetPasswordAuth = channel.unary_unary('/panels.auth.v1.AuthService/SetPasswordAuth', request_serializer=auth__pb2.SetPasswordAuthMethod.SerializeToString, response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString)
        self.DeletePasswordAuth = channel.unary_unary('/panels.auth.v1.AuthService/DeletePasswordAuth', request_serializer=auth__pb2.DeletePasswordAuthMethod.SerializeToString, response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString)

class AuthServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def AuthWithPassword(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetPasswordAuth(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeletePasswordAuth(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_AuthServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'AuthWithPassword': grpc.unary_unary_rpc_method_handler(servicer.AuthWithPassword, request_deserializer=auth__pb2.PasswordAuthRequest.FromString, response_serializer=auth__pb2.AuthToken.SerializeToString), 'SetPasswordAuth': grpc.unary_unary_rpc_method_handler(servicer.SetPasswordAuth, request_deserializer=auth__pb2.SetPasswordAuthMethod.FromString, response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString), 'DeletePasswordAuth': grpc.unary_unary_rpc_method_handler(servicer.DeletePasswordAuth, request_deserializer=auth__pb2.DeletePasswordAuthMethod.FromString, response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('panels.auth.v1.AuthService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class AuthService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def AuthWithPassword(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/panels.auth.v1.AuthService/AuthWithPassword', auth__pb2.PasswordAuthRequest.SerializeToString, auth__pb2.AuthToken.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SetPasswordAuth(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/panels.auth.v1.AuthService/SetPasswordAuth', auth__pb2.SetPasswordAuthMethod.SerializeToString, google_dot_protobuf_dot_empty__pb2.Empty.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeletePasswordAuth(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/panels.auth.v1.AuthService/DeletePasswordAuth', auth__pb2.DeletePasswordAuthMethod.SerializeToString, google_dot_protobuf_dot_empty__pb2.Empty.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)