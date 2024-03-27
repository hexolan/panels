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

import asyncio
from typing import Type

from comment_service.models.config import Config
from comment_service.models.service import CommentDBRepository
from comment_service.events.producer import CommentEventProducer
from comment_service.events.post_consumer import PostEventConsumer
from comment_service.events.user_consumer import UserEventConsumer


class EventConsumersWrapper:
    """A wrapper class for starting the event consumers.
    
    Attributes:
        _post_consumer (PostEventConsumer): A wrapped consumer.
        _user_consumer (UserEventConsumer): A wrapped consumer.

    """

    def __init__(self, post_consumer: PostEventConsumer, user_consumer: UserEventConsumer) -> None:
        """Add the consumers to the wrapper

        Args:
            post_consumer (PostEventConsumer): Initialised post consumer.
            user_consumer (UserEventConsumer): Initialised user consumer.

        """
        self._post_consumer = post_consumer
        self._user_consumer = user_consumer

    async def start(self) -> None:
        """Begin consuming events on all the event consumers."""
        await asyncio.gather(
            self._post_consumer.start(),
            self._user_consumer.start()
        )


def create_consumers(config: Config, db_repo: Type[CommentDBRepository]) -> EventConsumersWrapper:
    """Initialse the event consumers and return them in a wrapper.
    
    Args:
        config (Config): The app configuration instance.
        db_repo (Type[CommentDBRepository]): The database repo to pass to the consumers.
    
    Returns:
        EventConsumerWrapper

    """
    post_consumer = PostEventConsumer(config, db_repo)
    user_consumer = UserEventConsumer(config, db_repo)
    return EventConsumersWrapper(post_consumer=post_consumer, user_consumer=user_consumer)


async def create_producer(config: Config) -> CommentEventProducer:
    """Create an event producer for the service.
    
    Args:
        config (Config): The app configuration instance.

    Returns:
        CommentEventProducer
    
    """
    producer = CommentEventProducer(config)
    await producer.start()
    return producer