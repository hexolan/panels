"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from . import comment_pb2 as comment__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2

class CommentServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateComment = channel.unary_unary('/panels.comment.v1.CommentService/CreateComment', request_serializer=comment__pb2.CreateCommentRequest.SerializeToString, response_deserializer=comment__pb2.Comment.FromString)
        self.UpdateComment = channel.unary_unary('/panels.comment.v1.CommentService/UpdateComment', request_serializer=comment__pb2.UpdateCommentRequest.SerializeToString, response_deserializer=comment__pb2.Comment.FromString)
        self.DeleteComment = channel.unary_unary('/panels.comment.v1.CommentService/DeleteComment', request_serializer=comment__pb2.DeleteCommentRequest.SerializeToString, response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString)
        self.GetComment = channel.unary_unary('/panels.comment.v1.CommentService/GetComment', request_serializer=comment__pb2.GetCommentRequest.SerializeToString, response_deserializer=comment__pb2.Comment.FromString)
        self.GetPostComments = channel.unary_unary('/panels.comment.v1.CommentService/GetPostComments', request_serializer=comment__pb2.GetPostCommentsRequest.SerializeToString, response_deserializer=comment__pb2.PostComments.FromString)

class CommentServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def CreateComment(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateComment(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteComment(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetComment(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetPostComments(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_CommentServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'CreateComment': grpc.unary_unary_rpc_method_handler(servicer.CreateComment, request_deserializer=comment__pb2.CreateCommentRequest.FromString, response_serializer=comment__pb2.Comment.SerializeToString), 'UpdateComment': grpc.unary_unary_rpc_method_handler(servicer.UpdateComment, request_deserializer=comment__pb2.UpdateCommentRequest.FromString, response_serializer=comment__pb2.Comment.SerializeToString), 'DeleteComment': grpc.unary_unary_rpc_method_handler(servicer.DeleteComment, request_deserializer=comment__pb2.DeleteCommentRequest.FromString, response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString), 'GetComment': grpc.unary_unary_rpc_method_handler(servicer.GetComment, request_deserializer=comment__pb2.GetCommentRequest.FromString, response_serializer=comment__pb2.Comment.SerializeToString), 'GetPostComments': grpc.unary_unary_rpc_method_handler(servicer.GetPostComments, request_deserializer=comment__pb2.GetPostCommentsRequest.FromString, response_serializer=comment__pb2.PostComments.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('panels.comment.v1.CommentService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class CommentService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def CreateComment(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/panels.comment.v1.CommentService/CreateComment', comment__pb2.CreateCommentRequest.SerializeToString, comment__pb2.Comment.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateComment(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/panels.comment.v1.CommentService/UpdateComment', comment__pb2.UpdateCommentRequest.SerializeToString, comment__pb2.Comment.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteComment(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/panels.comment.v1.CommentService/DeleteComment', comment__pb2.DeleteCommentRequest.SerializeToString, google_dot_protobuf_dot_empty__pb2.Empty.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetComment(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/panels.comment.v1.CommentService/GetComment', comment__pb2.GetCommentRequest.SerializeToString, comment__pb2.Comment.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetPostComments(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/panels.comment.v1.CommentService/GetPostComments', comment__pb2.GetPostCommentsRequest.SerializeToString, comment__pb2.PostComments.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)