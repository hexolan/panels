import logging
import traceback
from typing import Type

from google.protobuf import empty_pb2
from grpc import RpcContext, StatusCode

from comment_service.models.exceptions import ServiceException
from comment_service.models.service import CommentRepository, Comment, CommentCreate, CommentUpdate
from comment_service.models.proto import comment_pb2, comment_pb2_grpc


class CommentServicer(comment_pb2_grpc.CommentServiceServicer):
    """Contains definitions for the service's RPC methods.
    
    Requests are converted from protobuf to business model
    form, then directed to the service repository, where the
    response is then translated back to protobuf.

    Attributes:
        _svc_repo (Type[CommentRepository]): The highest level service repository.
    
    """
    def __init__(self, svc_repo: Type[CommentRepository]) -> None:
        self._svc_repo = svc_repo
    
    def _apply_error(self, context: RpcContext, code: StatusCode, msg: str) -> None:
        """Apply an error to a given RPC context.
        
        Args:
            context (grpc.RpcContext): The context to apply the error to.
            code (grpc.StatusCode): The gRPC status code.
            msg (str): The error details.

        """
        context.set_code(code)
        context.set_details(msg)
    
    def _apply_unknown_error(self, context: RpcContext) -> None:
        """Apply a de facto error fallback message.
        
        Args:
            context (grpc.RpcContext): The context to apply the error to.
        
        """
        self._apply_error(context, StatusCode.UNKNOWN, "unknown error occured")

    async def CreateComment(self, request: comment_pb2.CreateCommentRequest, context: RpcContext) -> comment_pb2.Comment:
        """CreateComment RPC Call
        
        Args:
            request (comment_pb2.CreateCommentRequest): The request parameters.
            context (grpc.RpcContext): The context of the RPC call.

        Returns:
            comment_pb2.Comment: With a succesful comment creation.

        """
        # vaLidate the request inputs
        if request.post_id == "":
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                msg="post not provided"
            )
            return
        
        if request.author_id == "":
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                msg="author not provided"
            )
            return
        
        if request.data == None:
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                msg="malformed request"
            )
            return
        
        if request.data.message == "":
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                msg="comment message not provided"
            )
            return
        
        # convert to service model from protobuf
        try:
            data = CommentCreate.from_protobuf(request)
            comment = await self._svc_repo.create_comment(data)
        except ServiceException as err:
            err.apply_to_rpc(context)
            return
        except Exception:
            logging.error(traceback.format_exc())
            self._apply_unknown_error(context)
            return

        # convert comment to protobuf form
        return Comment.to_protobuf(comment)

    async def UpdateComment(self, request: comment_pb2.UpdateCommentRequest, context: RpcContext) -> comment_pb2.Comment:
        """UpdateComment RPC Call
        
        Args:
            request (comment_pb2.UpdateCommentRequest): The request parameters.
            context (grpc.RpcContext): The context of the RPC call.

        Returns:
            comment_pb2.Comment: The updated comment details (if succesfully updated).

        """
        # vaLidate the request inputs
        if request.id == "":
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                msg="comment not provided"
            )
            return
        elif not request.id.isnumeric():
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                msg="invalid comment id provided"
            )
            return
        
        if request.data == None:
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                msg="malformed request"
            )
            return
        
        if request.data.message == "":
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                msg="comment message not provided"
            )
            return
        
        # convert to service model from protobuf
        try:
            comment_id = int(request.id)
            data = CommentUpdate.from_protobuf(request)
            comment = await self._svc_repo.update_comment(comment_id, data)
        except ServiceException as err:
            err.apply_to_rpc(context)
            return
        except Exception:
            logging.error(traceback.format_exc())
            self._apply_unknown_error(context)
            return

        # convert comment to protobuf form
        return Comment.to_protobuf(comment)

    async def DeleteComment(self, request: comment_pb2.DeleteCommentRequest, context: RpcContext) -> empty_pb2.Empty:
        """DeleteComment RPC Call
        
        Args:
            request (comment_pb2.DeleteCommentRequest): The request parameters.
            context (grpc.RpcContext): The context of the RPC call.

        Returns:
            empty_pb2.Empty: Empty protobuf response (in effect returns None).

        """
        # vaLidate the request inputs
        if request.id == "":
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                msg="comment not provided"
            )
            return
        
        if not request.id.isnumeric():
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                msg="invalid comment id provided"
            )
            return

        # attempt to delete the comment
        try:
            comment_id = int(request.id)
            await self._svc_repo.delete_comment(comment_id)
        except ServiceException as err:
            err.apply_to_rpc(context)
            return
        except Exception:
            logging.error(traceback.format_exc())
            self._apply_unknown_error(context)
            return

        return empty_pb2.Empty()

    async def GetComment(self, request: comment_pb2.GetCommentRequest, context: RpcContext) -> comment_pb2.PostComments:
        """GetComment RPC Call
        
        Returns a comment by comment id.

        Args:
            request (comment_pb2.GetCommentRequest): The request parameters.
            context (grpc.RpcContext): The context of the RPC call.

        Returns:
            comment_pb2.Comment: The located comment

        """
        # vaLidate the request inputs
        if request.id == "":
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                msg="comment id not provided"
            )
            return
        
        if not request.id.isnumeric():
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                msg="invalid comment id provided"
            )
            return

        # attempt to get the comment
        try:
            comment = await self._svc_repo.get_comment(int(request.id))
        except ServiceException as err:
            err.apply_to_rpc(context)
            return
        except Exception:
            logging.error(traceback.format_exc())
            self._apply_unknown_error(context)
            return

        return Comment.to_protobuf(comment)

    async def GetPostComments(self, request: comment_pb2.GetPostCommentsRequest, context: RpcContext) -> comment_pb2.PostComments:
        """GetPostComments RPC Call
        
        Returns a list of comments that a post has.

        TODO: 
            Implement pagination (?after=comment_id or some effect)
            to return more comments from a post.

        Args:
            request (comment_pb2.UpdateCommentRequest): The request parameters.
            context (grpc.RpcContext): The context of the RPC call.

        Returns:
            comment_pb2.PostComments: containing a list of the post's comments

        """
        # vaLidate the request inputs
        if request.post_id == "":
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                msg="post id not provided"
            )
            return
        
        # attempt to get the comments
        try:
            comments = await self._svc_repo.get_post_comments(request.post_id)
        except ServiceException as err:
            err.apply_to_rpc(context)
            return
        except Exception:
            logging.error(traceback.format_exc())
            self._apply_unknown_error(context)
            return
        
        # convert to protobuf
        return comment_pb2.PostComments(comments=[Comment.to_protobuf(comment) for comment in comments])