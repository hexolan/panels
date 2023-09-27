from google.protobuf import empty_pb2 as _empty_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union
DESCRIPTOR: _descriptor.FileDescriptor

class User(_message.Message):
    __slots__ = ['id', 'username', 'is_admin', 'created_at', 'updated_at']
    ID_FIELD_NUMBER: _ClassVar[int]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    IS_ADMIN_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    UPDATED_AT_FIELD_NUMBER: _ClassVar[int]
    id: str
    username: str
    is_admin: bool
    created_at: _timestamp_pb2.Timestamp
    updated_at: _timestamp_pb2.Timestamp

    def __init__(self, id: _Optional[str]=..., username: _Optional[str]=..., is_admin: bool=..., created_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]]=..., updated_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]]=...) -> None:
        ...

class UserMutable(_message.Message):
    __slots__ = ['username']
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    username: str

    def __init__(self, username: _Optional[str]=...) -> None:
        ...

class CreateUserRequest(_message.Message):
    __slots__ = ['data']
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: UserMutable

    def __init__(self, data: _Optional[_Union[UserMutable, _Mapping]]=...) -> None:
        ...

class GetUserByIdRequest(_message.Message):
    __slots__ = ['id']
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str

    def __init__(self, id: _Optional[str]=...) -> None:
        ...

class GetUserByNameRequest(_message.Message):
    __slots__ = ['username']
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    username: str

    def __init__(self, username: _Optional[str]=...) -> None:
        ...

class UpdateUserByIdRequest(_message.Message):
    __slots__ = ['id', 'data']
    ID_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    id: str
    data: UserMutable

    def __init__(self, id: _Optional[str]=..., data: _Optional[_Union[UserMutable, _Mapping]]=...) -> None:
        ...

class UpdateUserByNameRequest(_message.Message):
    __slots__ = ['username', 'data']
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    username: str
    data: UserMutable

    def __init__(self, username: _Optional[str]=..., data: _Optional[_Union[UserMutable, _Mapping]]=...) -> None:
        ...

class DeleteUserByIdRequest(_message.Message):
    __slots__ = ['id']
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str

    def __init__(self, id: _Optional[str]=...) -> None:
        ...

class DeleteUserByNameRequest(_message.Message):
    __slots__ = ['username']
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    username: str

    def __init__(self, username: _Optional[str]=...) -> None:
        ...

class UserEvent(_message.Message):
    __slots__ = ['type', 'data']
    TYPE_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    type: str
    data: User

    def __init__(self, type: _Optional[str]=..., data: _Optional[_Union[User, _Mapping]]=...) -> None:
        ...