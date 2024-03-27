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

from databases import Database

from comment_service.models.config import Config
from comment_service.events.producer import CommentEventProducer
from comment_service.postgres.repository import ServiceDBRepository


async def connect_database(config: Config) -> Database:
    """Opens a connection to the database.
    
    Args:
        config (Config): The app configuration.

    Returns:
        A connected databases.Database instance

    """
    db = Database(config.postgres_dsn)
    try:
        await db.connect()
    except Exception:
        logging.error("failed to connect to postgresql database")
        raise
    
    return db


async def create_db_repository(config: Config, event_producer: CommentEventProducer) -> ServiceDBRepository:
    """Create the database repository.
    
    Open a database connection and instantialise the
    database repository.

    Returns:
        ServiceDBRepository
    
    """
    db = await connect_database(config)
    return ServiceDBRepository(db, event_producer)