from google.protobuf import empty_pb2 as _empty_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional
DESCRIPTOR: _descriptor.FileDescriptor

class SetPasswordAuthMethod(_message.Message):
    __slots__ = ['user_id', 'password']
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    user_id: str
    password: str

    def __init__(self, user_id: _Optional[str]=..., password: _Optional[str]=...) -> None:
        ...

class DeletePasswordAuthMethod(_message.Message):
    __slots__ = ['user_id']
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: str

    def __init__(self, user_id: _Optional[str]=...) -> None:
        ...

class PasswordAuthRequest(_message.Message):
    __slots__ = ['user_id', 'password']
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    user_id: str
    password: str

    def __init__(self, user_id: _Optional[str]=..., password: _Optional[str]=...) -> None:
        ...

class AuthToken(_message.Message):
    __slots__ = ['token_type', 'access_token', 'refresh_token', 'expires_in']
    TOKEN_TYPE_FIELD_NUMBER: _ClassVar[int]
    ACCESS_TOKEN_FIELD_NUMBER: _ClassVar[int]
    REFRESH_TOKEN_FIELD_NUMBER: _ClassVar[int]
    EXPIRES_IN_FIELD_NUMBER: _ClassVar[int]
    token_type: str
    access_token: str
    refresh_token: str
    expires_in: int

    def __init__(self, token_type: _Optional[str]=..., access_token: _Optional[str]=..., refresh_token: _Optional[str]=..., expires_in: _Optional[int]=...) -> None:
        ...