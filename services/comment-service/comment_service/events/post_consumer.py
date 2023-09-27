import logging

from comment_service.models.proto import post_pb2
from comment_service.events.base_consumer import EventConsumer


class PostEventConsumer(EventConsumer):
    """Consumer class responsible for 'post' events.
    
    Attributes:
        CONSUMER_TOPIC: The topic to consume events from.
        CONSUMER_EVENT_TYPE (post_pb2.PostEvent): Kafka messages are serialised to this type.
        _db_repo (CommentDBRepository): The repository interface for modifying data.
        _consumer (aiokafka.AIOKafkaConsumer): The underlying Kafka instance.

    """
    CONSUMER_TOPIC = "post"
    CONSUMER_EVENT_TYPE = post_pb2.PostEvent

    async def _process_event(self, event: post_pb2.PostEvent) -> None:
        """Process a recieved event.

        In response to a Post deleted event, delete any comments
        that fall under that post.

        Args:
            event (post_pb2.PostEvent): The decoded protobuf message.
        
        """
        if event.type == "deleted":
            assert event.data.id != ""
            await self._db_repo.delete_post_comments(event.data.id)
            logging.info("succesfully processed PostEvent (type: 'deleted')")