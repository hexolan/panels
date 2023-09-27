from google.protobuf import empty_pb2 as _empty_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class Comment(_message.Message):
    __slots__ = ['id', 'post_id', 'author_id', 'message', 'created_at', 'updated_at']
    ID_FIELD_NUMBER: _ClassVar[int]
    POST_ID_FIELD_NUMBER: _ClassVar[int]
    AUTHOR_ID_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    UPDATED_AT_FIELD_NUMBER: _ClassVar[int]
    id: str
    post_id: str
    author_id: str
    message: str
    created_at: _timestamp_pb2.Timestamp
    updated_at: _timestamp_pb2.Timestamp

    def __init__(self, id: _Optional[str]=..., post_id: _Optional[str]=..., author_id: _Optional[str]=..., message: _Optional[str]=..., created_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]]=..., updated_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]]=...) -> None:
        ...

class CommentMutable(_message.Message):
    __slots__ = ['message']
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    message: str

    def __init__(self, message: _Optional[str]=...) -> None:
        ...

class CreateCommentRequest(_message.Message):
    __slots__ = ['post_id', 'author_id', 'data']
    POST_ID_FIELD_NUMBER: _ClassVar[int]
    AUTHOR_ID_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    post_id: str
    author_id: str
    data: CommentMutable

    def __init__(self, post_id: _Optional[str]=..., author_id: _Optional[str]=..., data: _Optional[_Union[CommentMutable, _Mapping]]=...) -> None:
        ...

class UpdateCommentRequest(_message.Message):
    __slots__ = ['id', 'data']
    ID_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    id: str
    data: CommentMutable

    def __init__(self, id: _Optional[str]=..., data: _Optional[_Union[CommentMutable, _Mapping]]=...) -> None:
        ...

class DeleteCommentRequest(_message.Message):
    __slots__ = ['id']
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str

    def __init__(self, id: _Optional[str]=...) -> None:
        ...

class GetCommentRequest(_message.Message):
    __slots__ = ['id']
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str

    def __init__(self, id: _Optional[str]=...) -> None:
        ...

class GetPostCommentsRequest(_message.Message):
    __slots__ = ['post_id']
    POST_ID_FIELD_NUMBER: _ClassVar[int]
    post_id: str

    def __init__(self, post_id: _Optional[str]=...) -> None:
        ...

class PostComments(_message.Message):
    __slots__ = ['comments']
    COMMENTS_FIELD_NUMBER: _ClassVar[int]
    comments: _containers.RepeatedCompositeFieldContainer[Comment]

    def __init__(self, comments: _Optional[_Iterable[_Union[Comment, _Mapping]]]=...) -> None:
        ...

class CommentEvent(_message.Message):
    __slots__ = ['type', 'data']
    TYPE_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    type: str
    data: Comment

    def __init__(self, type: _Optional[str]=..., data: _Optional[_Union[Comment, _Mapping]]=...) -> None:
        ...