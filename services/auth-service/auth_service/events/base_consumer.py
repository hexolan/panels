import logging
from typing import Type

from google.protobuf import message
from aiokafka import AIOKafkaConsumer
from aiokafka.structs import ConsumerRecord

from auth_service.models.config import Config
from auth_service.models.service import AuthDBRepository


class EventConsumer:
    """An abstract consumer base class.
    
    Attributes:
        CONSUMER_TOPIC: The topic to consume events from.
        CONSUMER_EVENT_TYPE (Type[message.Message]): The protobuf class type of the event msgs (used for deserialisation).
        _db_repo (Type[AuthDBRepository]): The repository interface for modifying data.
        _consumer (aiokafka.AIOKafkaConsumer): The underlying Kafka instance.
        
    """
    CONSUMER_TOPIC: str
    CONSUMER_EVENT_TYPE: Type[message.Message]

    def __init__(self, config: Config, db_repo: Type[AuthDBRepository]) -> None:
        """Initialise the event consumer.
        
        Args:
            config (Config): The app configuration instance (to access brokers list).
            db_repo (Type[AuthDBRepository]): The repository interface for updating data.

        """
        self._db_repo = db_repo
        self._consumer = AIOKafkaConsumer(
            self.CONSUMER_TOPIC,
            bootstrap_servers=config.kafka_brokers,
            group_id="auth-service"
        )

    async def start(self) -> None:
        """Begin consuming messages."""
        await self._consumer.start()
        try:
            async for msg in self._consumer:
                await self._process_msg(msg)
        except Exception as e:
            logging.error(f"error whilst consuming messages (on topic '{self.CONSUMER_TOPIC}'): {e}")
        finally:
            await self._consumer.stop()

    async def _process_msg(self, msg: ConsumerRecord) -> None:
        """Process a recieved message.
        
        The messages are deserialise from bytes into their protobuf form,
        then passed to the `_process_event` method to handle the logic.
        
        Args:
            msg (kafka.Message): The event to process.
        
        """
        try:    
            event = self.CONSUMER_EVENT_TYPE()
            event.ParseFromString(msg.value)
            assert event.type != ""
            await self._process_event(event)
        except AssertionError:
            logging.error("invalid event recieved")
            return
        except Exception as e:
            logging.error("error whilst processing recieved event:", e)
            return

    async def _process_event(self, event: Type[message.Message]) -> None:
        """Process a recieved event.

        Args:
            event (Type[message.Message]): The event serialised to protobuf form.
        
        """
        raise NotImplementedError("required consumer method (_process_event) not implemented")