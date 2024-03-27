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