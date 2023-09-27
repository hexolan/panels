import logging

from comment_service.models.proto import user_pb2
from comment_service.events.base_consumer import EventConsumer


class UserEventConsumer(EventConsumer):
    """Consumer class responsible for 'user' events.
    
    Attributes:
        CONSUMER_TOPIC: The topic to consume events from.
        CONSUMER_EVENT_TYPE (user_pb2.UserEvent): Kafka messages are serialised to this type.
        _db_repo (CommentDBRepository): The repository interface for modifying data.
        _consumer (aiokafka.AIOKafkaConsumer): The underlying Kafka instance.

    """
    CONSUMER_TOPIC = "user"
    CONSUMER_EVENT_TYPE = user_pb2.UserEvent

    async def _process_event(self, event: user_pb2.UserEvent) -> None:
        """Process a recieved event.

        In response to a User deleted event, delete any comments
        created by that user.

        Args:
            event (user_pb2.UserEvent): The decoded protobuf message.
        
        """
        if event.type == "deleted":
            assert event.data.id != ""
            await self._db_repo.delete_user_comments(event.data.id)
            logging.info("succesfully processed UserEvent (type: 'deleted')")