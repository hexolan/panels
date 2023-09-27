from datetime import datetime
from typing import List, Optional

from pydantic import BaseModel
from google.protobuf import timestamp_pb2

from comment_service.models.proto import comment_pb2
from comment_service.models.exceptions import ServiceException, ServiceErrorCode


# Validators
def is_valid_comment_msg(message: str) -> bool:
    if len(message) < 3 or len(message) > 512:
        return False
    return True


# Service Models
class Comment(BaseModel):
    id: int

    post_id: str
    author_id: str
    message: str

    created_at: datetime
    updated_at: Optional[datetime] = None

    @classmethod
    def to_protobuf(cls, comment: "Comment") -> comment_pb2.Comment:
        created_at = timestamp_pb2.Timestamp()
        created_at.FromDatetime(comment.created_at)

        updated_at = None
        if comment.updated_at is not None:
            updated_at = timestamp_pb2.Timestamp()
            updated_at.FromDatetime(comment.updated_at)

        return comment_pb2.Comment(
            id=str(comment.id),
            post_id=comment.post_id,
            author_id=comment.author_id,
            message=comment.message,

            created_at=created_at,
            updated_at=updated_at,
        )


class CommentCreate(BaseModel):
    post_id: str
    author_id: str
    message: str  # todo: validation on message (wrap validator for is_valid_comment_msg)

    @classmethod
    def from_protobuf(cls, request: comment_pb2.CreateCommentRequest) -> "CommentCreate":
        return cls(
            post_id=request.post_id,
            author_id=request.author_id,
            message=request.data.message,
        )


class CommentUpdate(BaseModel):
    message: Optional[str] = None  # todo: validation on message (if set use validator is_valid_comment_msg)

    @classmethod
    def from_protobuf(cls, request: comment_pb2.UpdateCommentRequest) -> "CommentUpdate":
        return cls(
            message=request.data.message
        )


# Repository Interfaces
class CommentRepository:
    async def get_comment(self, comment_id: int) -> Comment:
        raise ServiceException("unimplemented internal repository method", ServiceErrorCode.SERVICE_ERROR)

    async def get_post_comments(self, post_id: str) -> List[Comment]:
        raise ServiceException("unimplemented internal repository method", ServiceErrorCode.SERVICE_ERROR)
    
    async def create_comment(self, data: CommentCreate) -> Comment:
        raise ServiceException("unimplemented internal repository method", ServiceErrorCode.SERVICE_ERROR)
    
    async def update_comment(self, comment_id: int, data: CommentUpdate) -> Comment:
        raise ServiceException("unimplemented internal repository method", ServiceErrorCode.SERVICE_ERROR)
    
    async def delete_comment(self, comment_id: int) -> None:
        raise ServiceException("unimplemented internal repository method", ServiceErrorCode.SERVICE_ERROR)


class CommentDBRepository(CommentRepository):
    async def delete_post_comments(self, post_id: str) -> None:
        raise ServiceException("unimplemented internal repository method", ServiceErrorCode.SERVICE_ERROR)
    
    async def delete_user_comments(self, user_id: str) -> None:
        raise ServiceException("unimplemented internal repository method", ServiceErrorCode.SERVICE_ERROR)