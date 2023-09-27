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