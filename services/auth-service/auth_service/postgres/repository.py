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

from typing import Optional

from databases import Database

from auth_service.models.service import AuthDBRepository, AuthRecord


class ServiceDBRepository(AuthDBRepository):
    """Database repository responsible for CRUD actions 
    on the database.
    
    This repository will be utilised by other upstream repositories
    or the Kafka event consumers.

    Attributes:
        _db (Database): The postgres database connection handler.

    """
    def __init__(self, db: Database) -> None:
        self._db = db

    async def get_auth_record(self, user_id: str) -> Optional[AuthRecord]:
        query = "SELECT user_id, password FROM auth_methods WHERE user_id = :user_id"
        result = await self._db.fetch_one(query=query, values={"user_id": user_id})
        if result is None:
            return None
        return AuthRecord(user_id=result["user_id"], password=result["password"])

    async def create_password_auth_method(self, user_id: str, password: str) -> None:
        query = "INSERT INTO auth_methods (user_id, password) VALUES (:user_id, :password)"
        await self._db.execute(query=query, values={"user_id": user_id, "password": password})

    async def update_password_auth_method(self, user_id: str, password: str) -> None:
        query = "UPDATE auth_methods SET password = :password WHERE user_id = :user_id"
        await self._db.execute(query=query, values={"user_id": user_id, "password": password})

    async def delete_password_auth_method(self, user_id: str) -> None:
        query = "DELETE FROM auth_methods WHERE user_id = :user_id"
        await self._db.execute(query=query, values={"user_id": user_id})