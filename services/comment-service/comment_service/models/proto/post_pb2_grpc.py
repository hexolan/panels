"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from . import post_pb2 as post__pb2

class PostServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreatePost = channel.unary_unary('/panels.post.v1.PostService/CreatePost', request_serializer=post__pb2.CreatePostRequest.SerializeToString, response_deserializer=post__pb2.Post.FromString)
        self.GetPost = channel.unary_unary('/panels.post.v1.PostService/GetPost', request_serializer=post__pb2.GetPostRequest.SerializeToString, response_deserializer=post__pb2.Post.FromString)
        self.GetPanelPost = channel.unary_unary('/panels.post.v1.PostService/GetPanelPost', request_serializer=post__pb2.GetPanelPostRequest.SerializeToString, response_deserializer=post__pb2.Post.FromString)
        self.UpdatePost = channel.unary_unary('/panels.post.v1.PostService/UpdatePost', request_serializer=post__pb2.UpdatePostRequest.SerializeToString, response_deserializer=post__pb2.Post.FromString)
        self.DeletePost = channel.unary_unary('/panels.post.v1.PostService/DeletePost', request_serializer=post__pb2.DeletePostRequest.SerializeToString, response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString)
        self.GetFeedPosts = channel.unary_unary('/panels.post.v1.PostService/GetFeedPosts', request_serializer=post__pb2.GetFeedPostsRequest.SerializeToString, response_deserializer=post__pb2.FeedPosts.FromString)
        self.GetUserPosts = channel.unary_unary('/panels.post.v1.PostService/GetUserPosts', request_serializer=post__pb2.GetUserPostsRequest.SerializeToString, response_deserializer=post__pb2.UserPosts.FromString)
        self.GetPanelPosts = channel.unary_unary('/panels.post.v1.PostService/GetPanelPosts', request_serializer=post__pb2.GetPanelPostsRequest.SerializeToString, response_deserializer=post__pb2.PanelPosts.FromString)

class PostServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def CreatePost(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetPost(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetPanelPost(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdatePost(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeletePost(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetFeedPosts(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetUserPosts(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetPanelPosts(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

def add_PostServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {'CreatePost': grpc.unary_unary_rpc_method_handler(servicer.CreatePost, request_deserializer=post__pb2.CreatePostRequest.FromString, response_serializer=post__pb2.Post.SerializeToString), 'GetPost': grpc.unary_unary_rpc_method_handler(servicer.GetPost, request_deserializer=post__pb2.GetPostRequest.FromString, response_serializer=post__pb2.Post.SerializeToString), 'GetPanelPost': grpc.unary_unary_rpc_method_handler(servicer.GetPanelPost, request_deserializer=post__pb2.GetPanelPostRequest.FromString, response_serializer=post__pb2.Post.SerializeToString), 'UpdatePost': grpc.unary_unary_rpc_method_handler(servicer.UpdatePost, request_deserializer=post__pb2.UpdatePostRequest.FromString, response_serializer=post__pb2.Post.SerializeToString), 'DeletePost': grpc.unary_unary_rpc_method_handler(servicer.DeletePost, request_deserializer=post__pb2.DeletePostRequest.FromString, response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString), 'GetFeedPosts': grpc.unary_unary_rpc_method_handler(servicer.GetFeedPosts, request_deserializer=post__pb2.GetFeedPostsRequest.FromString, response_serializer=post__pb2.FeedPosts.SerializeToString), 'GetUserPosts': grpc.unary_unary_rpc_method_handler(servicer.GetUserPosts, request_deserializer=post__pb2.GetUserPostsRequest.FromString, response_serializer=post__pb2.UserPosts.SerializeToString), 'GetPanelPosts': grpc.unary_unary_rpc_method_handler(servicer.GetPanelPosts, request_deserializer=post__pb2.GetPanelPostsRequest.FromString, response_serializer=post__pb2.PanelPosts.SerializeToString)}
    generic_handler = grpc.method_handlers_generic_handler('panels.post.v1.PostService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

class PostService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def CreatePost(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/panels.post.v1.PostService/CreatePost', post__pb2.CreatePostRequest.SerializeToString, post__pb2.Post.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetPost(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/panels.post.v1.PostService/GetPost', post__pb2.GetPostRequest.SerializeToString, post__pb2.Post.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetPanelPost(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/panels.post.v1.PostService/GetPanelPost', post__pb2.GetPanelPostRequest.SerializeToString, post__pb2.Post.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdatePost(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/panels.post.v1.PostService/UpdatePost', post__pb2.UpdatePostRequest.SerializeToString, post__pb2.Post.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeletePost(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/panels.post.v1.PostService/DeletePost', post__pb2.DeletePostRequest.SerializeToString, google_dot_protobuf_dot_empty__pb2.Empty.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetFeedPosts(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/panels.post.v1.PostService/GetFeedPosts', post__pb2.GetFeedPostsRequest.SerializeToString, post__pb2.FeedPosts.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetUserPosts(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/panels.post.v1.PostService/GetUserPosts', post__pb2.GetUserPostsRequest.SerializeToString, post__pb2.UserPosts.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetPanelPosts(request, target, options=(), channel_credentials=None, call_credentials=None, insecure=False, compression=None, wait_for_ready=None, timeout=None, metadata=None):
        return grpc.experimental.unary_unary(request, target, '/panels.post.v1.PostService/GetPanelPosts', post__pb2.GetPanelPostsRequest.SerializeToString, post__pb2.PanelPosts.FromString, options, channel_credentials, insecure, call_credentials, compression, wait_for_ready, timeout, metadata)