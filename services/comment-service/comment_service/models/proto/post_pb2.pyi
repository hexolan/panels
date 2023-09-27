from google.protobuf import empty_pb2 as _empty_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class Post(_message.Message):
    __slots__ = ['id', 'panel_id', 'author_id', 'title', 'content', 'created_at', 'updated_at']
    ID_FIELD_NUMBER: _ClassVar[int]
    PANEL_ID_FIELD_NUMBER: _ClassVar[int]
    AUTHOR_ID_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    CONTENT_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    UPDATED_AT_FIELD_NUMBER: _ClassVar[int]
    id: str
    panel_id: str
    author_id: str
    title: str
    content: str
    created_at: _timestamp_pb2.Timestamp
    updated_at: _timestamp_pb2.Timestamp

    def __init__(self, id: _Optional[str]=..., panel_id: _Optional[str]=..., author_id: _Optional[str]=..., title: _Optional[str]=..., content: _Optional[str]=..., created_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]]=..., updated_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]]=...) -> None:
        ...

class PostMutable(_message.Message):
    __slots__ = ['title', 'content']
    TITLE_FIELD_NUMBER: _ClassVar[int]
    CONTENT_FIELD_NUMBER: _ClassVar[int]
    title: str
    content: str

    def __init__(self, title: _Optional[str]=..., content: _Optional[str]=...) -> None:
        ...

class CreatePostRequest(_message.Message):
    __slots__ = ['panel_id', 'user_id', 'data']
    PANEL_ID_FIELD_NUMBER: _ClassVar[int]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    panel_id: str
    user_id: str
    data: PostMutable

    def __init__(self, panel_id: _Optional[str]=..., user_id: _Optional[str]=..., data: _Optional[_Union[PostMutable, _Mapping]]=...) -> None:
        ...

class GetPostRequest(_message.Message):
    __slots__ = ['id']
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str

    def __init__(self, id: _Optional[str]=...) -> None:
        ...

class GetPanelPostRequest(_message.Message):
    __slots__ = ['panel_id', 'id']
    PANEL_ID_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    panel_id: str
    id: str

    def __init__(self, panel_id: _Optional[str]=..., id: _Optional[str]=...) -> None:
        ...

class UpdatePostRequest(_message.Message):
    __slots__ = ['id', 'data']
    ID_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    id: str
    data: PostMutable

    def __init__(self, id: _Optional[str]=..., data: _Optional[_Union[PostMutable, _Mapping]]=...) -> None:
        ...

class DeletePostRequest(_message.Message):
    __slots__ = ['id']
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str

    def __init__(self, id: _Optional[str]=...) -> None:
        ...

class GetFeedPostsRequest(_message.Message):
    __slots__ = []

    def __init__(self) -> None:
        ...

class FeedPosts(_message.Message):
    __slots__ = ['posts']
    POSTS_FIELD_NUMBER: _ClassVar[int]
    posts: _containers.RepeatedCompositeFieldContainer[Post]

    def __init__(self, posts: _Optional[_Iterable[_Union[Post, _Mapping]]]=...) -> None:
        ...

class GetUserPostsRequest(_message.Message):
    __slots__ = ['user_id']
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: str

    def __init__(self, user_id: _Optional[str]=...) -> None:
        ...

class UserPosts(_message.Message):
    __slots__ = ['posts']
    POSTS_FIELD_NUMBER: _ClassVar[int]
    posts: _containers.RepeatedCompositeFieldContainer[Post]

    def __init__(self, posts: _Optional[_Iterable[_Union[Post, _Mapping]]]=...) -> None:
        ...

class GetPanelPostsRequest(_message.Message):
    __slots__ = ['panel_id']
    PANEL_ID_FIELD_NUMBER: _ClassVar[int]
    panel_id: str

    def __init__(self, panel_id: _Optional[str]=...) -> None:
        ...

class PanelPosts(_message.Message):
    __slots__ = ['posts']
    POSTS_FIELD_NUMBER: _ClassVar[int]
    posts: _containers.RepeatedCompositeFieldContainer[Post]

    def __init__(self, posts: _Optional[_Iterable[_Union[Post, _Mapping]]]=...) -> None:
        ...

class PostEvent(_message.Message):
    __slots__ = ['type', 'data']
    TYPE_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    type: str
    data: Post

    def __init__(self, type: _Optional[str]=..., data: _Optional[_Union[Post, _Mapping]]=...) -> None:
        ...