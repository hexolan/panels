import asyncio
import logging
from sys import stdout

from comment_service.models.config import Config
from comment_service.events.service import create_producer, create_consumers
from comment_service.postgres.service import create_db_repository
from comment_service.redis.service import create_redis_repository
from comment_service.rpc.service import create_rpc_server
from comment_service.service import ServiceRepository


async def main() -> None:
    config = Config()

    event_prod = await create_producer(config)
    db_repo = await create_db_repository(config, event_producer=event_prod)
    redis_repo = await create_redis_repository(config, downstream_repo=db_repo)
    svc_repo = ServiceRepository(downstream_repo=redis_repo)

    rpc_server = create_rpc_server(svc_repo)
    event_consumers = create_consumers(config, db_repo=db_repo)

    await asyncio.gather(
        rpc_server.start(),
        event_consumers.start()
    )


if __name__ == "__main__":
    logging.basicConfig(stream=stdout, level=logging.INFO)
    asyncio.run(main())
