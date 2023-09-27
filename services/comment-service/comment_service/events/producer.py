import asyncio

from aiokafka import AIOKafkaProducer

from comment_service.models.config import Config
from comment_service.models.service import Comment
from comment_service.models.proto import comment_pb2


class CommentEventProducer:
    """Service event producer.
    
    Attributes:
        _producer (aiokafka.AIOKafkaProducer): The underlying Kafka instance.
        _pending_sends (set): Background events that are still awaiting a response.
        
    """
    def __init__(self, config: Config) -> None:
        """Initialise the event producer.
        
        Args:
            config (Config): The app configuration instance (to access Kafka brokers list).

        """
        self._producer = AIOKafkaProducer(
            bootstrap_servers=config.kafka_brokers,
            client_id="auth-service"
        )
        self._pending_sends = set()

    async def start(self) -> None:
        """Start the underlying Kafka instance (open a connection)."""
        await self._producer.start()

    async def _send_event(self, event: comment_pb2.CommentEvent) -> None:
        """Send an event.
        
        Execute the sending action as a background task to avoid
        delaying a response to the request.

        Args:
            event (comment_pb2.CommentEvent): The protobuf class for the event.

        """
        bg_task = asyncio.create_task(self._producer.send(
            topic="comment",
            value=event.SerializeToString(),
        ))
        self._pending_sends.add(bg_task)
        bg_task.add_done_callback(self._pending_sends.discard)

    async def send_created_event(self, comment: Comment) -> None:
        """Send a comment created event.
        
        Args:
            comment (Comment): The comment to reference.
        
        """
        event = comment_pb2.CommentEvent(
            type="created",
            data=Comment.to_protobuf(comment)
        )
        await self._send_event(event)

    async def send_updated_event(self, comment: Comment) -> None:
        """Send a comment updated event.
        
        Args:
            comment (Comment): The comment to reference.
        
        """
        event = comment_pb2.CommentEvent(
            type="updated",
            data=Comment.to_protobuf(comment)
        )
        await self._send_event(event)

    async def send_deleted_event(self, comment: Comment) -> None:
        """Send a comment deleted event.
        
        Args:
            comment (Comment): The comment to reference.
        
        """
        event = comment_pb2.CommentEvent(
            type="deleted",
            data=Comment.to_protobuf(comment)
        )
        await self._send_event(event)