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
