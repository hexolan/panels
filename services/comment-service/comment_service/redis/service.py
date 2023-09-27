from typing import Type

import redis.asyncio as redis

from comment_service.models.config import Config
from comment_service.models.service import CommentRepository
from comment_service.redis.repository import ServiceRedisRepository


async def connect_redis(config: Config) -> redis.Redis:
    """Opens a connection to Redis.
    
    Args:
        config (Config): The app configuration.

    Returns:
        A connected redis.Redis instance
    
    """
    host, port = config.redis_host, 5432
    try:
        host, port = config.redis_host.split(":")
        port = int(port)
    except Exception:
        pass

    conn = redis.Redis(host=host, port=port, password=config.redis_pass)
    await conn.ping()
    return conn


async def create_redis_repository(config: Config, downstream_repo: Type[CommentRepository]) -> ServiceRedisRepository:
    """Create the Redis repository.
    
    Open a Redis connection and instantialise the
    Redis repository.

    Args:
        downstream_repo (Type[CommentRepository]): The next downstream repository
            that will be called by the Redis repository. (in this case the database
            repository)

    Returns:
        ServiceRedisRepository
    
    """
    redis_conn = await connect_redis(config)
    return ServiceRedisRepository(redis_conn=redis_conn, downstream_repo=downstream_repo)