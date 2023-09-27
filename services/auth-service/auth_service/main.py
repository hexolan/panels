import asyncio
import logging
from sys import stdout

from auth_service.models.config import Config
from auth_service.events.service import create_consumers
from auth_service.postgres.service import create_db_repository
from auth_service.service import ServiceRepository
from auth_service.rpc.service import create_rpc_server


async def main() -> None:
    config = Config()

    db_repo = await create_db_repository(config)
    svc_repo = ServiceRepository(config, downstream_repo=db_repo)

    rpc_server = create_rpc_server(svc_repo)
    event_consumers = create_consumers(config, db_repo=db_repo)

    await asyncio.gather(
        rpc_server.start(),
        event_consumers.start()
    )


if __name__ == "__main__":
    logging.basicConfig(stream=stdout, level=logging.INFO)
    asyncio.run(main())
