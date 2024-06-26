# Copyright 2023 Declan Teevan
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

from time import time
from typing import Optional
from datetime import timedelta

from pydantic import BaseModel, Field, SecretStr

from auth_service.models.proto import auth_pb2
from auth_service.models.exceptions import ServiceException, ServiceErrorCode


class AuthRecord(BaseModel):
    user_id: str
    password: SecretStr  # Hashed Password


class AuthToken(BaseModel):
    token_type: str = "Bearer"
    access_token: str
    expires_in: int = Field(default=int(timedelta(minutes=30).total_seconds()))
    # refresh_token: str  # todo: implement functionality in the future

    @classmethod
    def to_protobuf(cls, auth_token: "AuthToken") -> auth_pb2.AuthToken:
        return auth_pb2.AuthToken(**auth_token.model_dump())


class AccessTokenClaims(BaseModel):
    sub: str
    iat: int = Field(default_factory=lambda: int(time()))
    exp: int = Field(default_factory=lambda: int(time() + timedelta(minutes=30).total_seconds()))


class AuthRepository:
    """Abstract repository interface"""
    async def auth_with_password(self, user_id: str, password: str) -> AuthToken:
        raise ServiceException("unimplemented internal repository method", ServiceErrorCode.SERVICE_ERROR)
    
    async def set_password_auth_method(self, user_id: str, password: str) -> None:
        raise ServiceException("unimplemented internal repository method", ServiceErrorCode.SERVICE_ERROR)

    async def delete_password_auth_method(self, user_id: str) -> None:
        raise ServiceException("unimplemented internal repository method", ServiceErrorCode.SERVICE_ERROR)


class AuthDBRepository(AuthRepository):
    """Abstract database repository interface"""
    async def get_auth_record(self, user_id: str) -> Optional[AuthRecord]:
        raise ServiceException("unimplemented internal repository method", ServiceErrorCode.SERVICE_ERROR)

    async def create_password_auth_method(self, user_id: str, password: str) -> None:
        raise ServiceException("unimplemented internal repository method", ServiceErrorCode.SERVICE_ERROR)
    
    async def update_password_auth_method(self, user_id: str, password: str) -> None:
        raise ServiceException("unimplemented internal repository method", ServiceErrorCode.SERVICE_ERROR)
    