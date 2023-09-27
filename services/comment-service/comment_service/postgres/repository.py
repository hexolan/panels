from typing import List

from databases import Database

from comment_service.events.producer import CommentEventProducer
from comment_service.models.exceptions import ServiceException, ServiceErrorCode
from comment_service.models.service import CommentDBRepository, Comment, CommentCreate, CommentUpdate


class ServiceDBRepository(CommentDBRepository):
    """Database repository responsible for actions
    relating to the postgres database.
    
    This repository will be utilised by other upstream repositories
    or the Kafka event consumers.

    Attributes:
        _db (Database): The postgres database connection handler.
        _event_prod (CommentEventProducer): Used to dispatch events upon execution of a CRUD action.

    """
    def __init__(self, db: Database, event_producer: CommentEventProducer) -> None:
        self._db = db
        self._event_prod = event_producer

    def _result_to_comment(self, result) -> Comment:
        return Comment(
            id=result.id,
            post_id=result.post_id,
            author_id=result.author_id,
            message=result.message,
            created_at=result.created_at,
            updated_at=result._mapping.get("updated_at", None)
        )

    async def get_post_comments(self, post_id: str) -> List[Comment]:
        query = "SELECT id, post_id, author_id, message, created_at, updated_at FROM comments WHERE post_id = :post_id"
        rows = await self._db.fetch_all(query=query, values={"post_id": post_id})
        return [self._result_to_comment(result) for result in rows]
    
    async def create_comment(self, data: CommentCreate) -> Comment:
        query = "INSERT INTO comments (post_id, author_id, message) VALUES (:post_id, :author_id, :message) RETURNING id"
        result = await self._db.execute(query=query, values={"post_id": data.post_id, "author_id": data.author_id, "message": data.message})

        comment = await self.get_comment(result)
        await self._event_prod.send_created_event(comment)
        return comment

    async def update_comment(self, comment_id: int, data: CommentUpdate) -> Comment:
        query = "UPDATE comments SET message = :message, updated_at = now() WHERE id = :comment_id"
        await self._db.execute(query=query, values={"message": data.message, "comment_id": comment_id})

        comment = await self.get_comment(comment_id)
        await self._event_prod.send_updated_event(comment)
        return comment
    
    async def delete_comment(self, comment_id: int) -> None:
        comment = await self.get_comment(comment_id)

        query = "DELETE FROM comments WHERE id = :comment_id"
        await self._db.execute(query=query, values={"comment_id": comment_id})

        await self._event_prod.send_deleted_event(comment)
    
    async def get_comment(self, comment_id: int) -> Comment:
        query = "SELECT id, post_id, author_id, message, created_at, updated_at FROM comments WHERE id = :comment_id"
        result = await self._db.fetch_one(query=query, values={"comment_id": comment_id})
        if result is None:
            raise ServiceException(message="no comment found", error_code=ServiceErrorCode.NOT_FOUND)
        
        return self._result_to_comment(result)

    async def delete_post_comments(self, post_id: str) -> None:
        comments = await self.get_post_comments(post_id)

        query = "DELETE FROM comments WHERE post_id = :post_id"
        await self._db.execute(query=query, values={"post_id": post_id})

        for comment in comments:
            await self._event_prod.send_deleted_event(comment)
    
    async def delete_user_comments(self, user_id: str) -> None:
        comments = await self.get_post_comments(user_id)

        query = "DELETE FROM comments WHERE author_id = :author_id"
        await self._db.execute(query=query, values={"author_id": user_id})

        for comment in comments:
            await self._event_prod.send_deleted_event(comment)

    async def _get_user_comments(self, user_id: str) -> List[Comment]:
        query = "SELECT id, post_id, author_id, message, created_at, updated_at FROM comments WHERE author_id = :user_id"
        rows = await self._db.fetch_all(query=query, values={"user_id": user_id})
        return [self._result_to_comment(result) for result in rows]