from typing import Type

from auth_service.models.config import Config
from auth_service.models.service import AuthDBRepository
from auth_service.events.user_consumer import UserEventConsumer


class EventConsumersWrapper:
    """A wrapper class for starting the event consumers.
    
    Attributes:
        _user_consumer (UserEventConsumer): Wrapped consumer.

    """

    def __init__(self, user_consumer: UserEventConsumer) -> None:
        """Add the consumers to the wrapper
        
        Args:
            user_consumer (UserEventConsumer): Initialised user consumer.
        
        """
        self._user_consumer = user_consumer

    async def start(self) -> None:
        """Begin consuming events on all the event consumers."""
        await self._user_consumer.start()


def create_consumers(config: Config, db_repo: Type[AuthDBRepository]) -> EventConsumersWrapper:
    """Initialse the event consumers and return them in a wrapper.
    
    Args:
        config (Config): The app configuration instance.
        db_repo (Type[AuthDBRepository]): The database repo to pass to the consumers.
    
    Returns:
        EventConsumerWrapper

    """
    user_consumer = UserEventConsumer(config, db_repo)
    return EventConsumersWrapper(user_consumer=user_consumer)