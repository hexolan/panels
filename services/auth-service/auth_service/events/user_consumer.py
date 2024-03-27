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

import logging

from auth_service.models.proto import user_pb2
from auth_service.events.base_consumer import EventConsumer


class UserEventConsumer(EventConsumer):
    """Consumer class responsible for 'user' events.
    
    Attributes:
        CONSUMER_TOPIC: The topic to consume events from.
        CONSUMER_EVENT_TYPE (user_pb2.UserEvent): Kafka messages are serialised to this type.
        _db_repo (AuthDBRepository): The repository interface for modifying data.
        _consumer (aiokafka.AIOKafkaConsumer): The underlying Kafka instance.

    """
    CONSUMER_TOPIC = "user"
    CONSUMER_EVENT_TYPE = user_pb2.UserEvent

    async def _process_event(self, event: user_pb2.UserEvent) -> None:
        """Process a recieved event.

        In response to a User deleted event, delete any auth methods
        this service has in relation to that user.

        Args:
            event (user_pb2.UserEvent): The decoded protobuf message.
        
        """
        if event.type == "deleted":
            await self._db_repo.delete_password_auth_method(event.data.id)
            logging.info("succesfully processed UserEvent (type: 'deleted')")