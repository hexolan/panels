"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from . import user_pb2 as user__pb2

class UserServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateUser = channel.unary_unary('/panels.user.v1.UserService/CreateUser', request_serializer=user__pb2.CreateUserRequest.SerializeToString, response_deserializer=user__pb2.User.FromString)
        self.GetUser = channel.unary_unary('/panels.user.v1.UserService/GetUser', request_serializer=user__pb2.GetUserByIdRequest.SerializeToString, response_deserializer=user__pb2.User.FromString)
        self.GetUserByName = channel.unary_unary('/panels.user.v1.UserService/GetUserByName', request_serializer=user__pb2.GetUserByNameRequest.SerializeToString, response_deserializer=user__pb2.User.FromString)
        self.UpdateUser = channel.unary_unary('/panels.user.v1.UserService/UpdateUser', request_serializer=user__pb2.UpdateUserByIdRequest.SerializeToString, response_deserializer=user__pb2.User.FromString)
        self.UpdateUserByName = channel.unary_unary('/panels.user.v1.UserService/UpdateUserByName', request_serializer=user__pb2.UpdateUserByNameRequest.SerializeToString, response_deserializer=user__pb2.User.FromString)
        self.DeleteUser = channel.unary_unary('/panels.user.v1.UserService/DeleteUser', request_serializer=user__pb2.DeleteUserByIdRequest.SerializeToString, response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString)
        self.DeleteUserByName = channel.unary_unary('/panels.user.v1.UserService/DeleteUserByName', request_serializer=user__pb2.DeleteUserByNameRequest.SerializeToString, response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString)

class UserServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def CreateUser(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetUser(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetUserByName(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateUser(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateUserByName(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteUser(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteUserByName(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_UserServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'CreateUser': grpc.unary_unary_rpc_method_handler(servicer.CreateUser, request_deserializer=user__pb2.CreateUserRequest.FromString, response_serializer=user__pb2.User.SerializeToString), 'GetUser': grpc.unary_unary_rpc_method_handler(servicer.GetUser, request_deserializer=user__pb2.GetUserByIdRequest.FromString, response_serializer=user__pb2.User.SerializeToString), 'GetUserByName': grpc.unary_unary_rpc_method_handler(servicer.GetUserByName, request_deserializer=user__pb2.GetUserByNameRequest.FromString, response_serializer=user__pb2.User.SerializeToString), 'UpdateUser': grpc.unary_unary_rpc_method_handler(servicer.UpdateUser, request_deserializer=user__pb2.UpdateUserByIdRequest.FromString, response_serializer=user__pb2.User.SerializeToString), 'UpdateUserByName': grpc.unary_unary_rpc_method_handler(servicer.UpdateUserByName, request_deserializer=user__pb2.UpdateUserByNameRequest.FromString, response_serializer=user__pb2.User.SerializeToString), 'DeleteUser': grpc.unary_unary_rpc_method_handler(servicer.DeleteUser, request_deserializer=user__pb2.DeleteUserByIdRequest.FromString, response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString), 'DeleteUserByName': grpc.unary_unary_rpc_method_handler(servicer.DeleteUserByName, request_deserializer=user__pb2.DeleteUserByNameRequest.FromString, response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('panels.user.v1.UserService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class UserService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def CreateUser(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/panels.user.v1.UserService/CreateUser', user__pb2.CreateUserRequest.SerializeToString, user__pb2.User.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetUser(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/panels.user.v1.UserService/GetUser', user__pb2.GetUserByIdRequest.SerializeToString, user__pb2.User.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetUserByName(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/panels.user.v1.UserService/GetUserByName', user__pb2.GetUserByNameRequest.SerializeToString, user__pb2.User.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateUser(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/panels.user.v1.UserService/UpdateUser', user__pb2.UpdateUserByIdRequest.SerializeToString, user__pb2.User.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateUserByName(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/panels.user.v1.UserService/UpdateUserByName', user__pb2.UpdateUserByNameRequest.SerializeToString, user__pb2.User.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteUser(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/panels.user.v1.UserService/DeleteUser', user__pb2.DeleteUserByIdRequest.SerializeToString, google_dot_protobuf_dot_empty__pb2.Empty.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteUserByName(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/panels.user.v1.UserService/DeleteUserByName', user__pb2.DeleteUserByNameRequest.SerializeToString, google_dot_protobuf_dot_empty__pb2.Empty.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)