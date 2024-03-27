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
from typing import Type

import grpc
from grpc_health.v1 import health, health_pb2_grpc

from comment_service.rpc.comment import CommentServicer
from comment_service.models.proto import comment_pb2_grpc
from comment_service.models.service import CommentRepository


class RPCServerWrapper:
    """A wrapper class for the RPC server.

    Attributes:
        _grpc_server (grpc.aio.Server): The gRPC server instance.
    
    """
    def __init__(self, svc_repo: Type[CommentRepository]) -> None:
        """Creates the gRPC server and adds the servicers.
        
        Args:
            svc_repo (Type[CommentRepository]): The service repository to pass to the servicers.

        """
        self._grpc_server = grpc.aio.server()
        self._grpc_server.add_insecure_port("[::]:9090")

        comment_servicer = CommentServicer(svc_repo)
        comment_pb2_grpc.add_CommentServiceServicer_to_server(comment_servicer, self._grpc_server)

        health_servicer = health.aio.HealthServicer()
        health_pb2_grpc.add_HealthServicer_to_server(health_servicer, self._grpc_server)

    async def start(self) -> None:
        """Begin serving RPC asynchronously."""
        logging.info("attempting to serve RPC...")
        await self._grpc_server.start()
        await self._grpc_server.wait_for_termination()


def create_rpc_server(svc_repo: Type[CommentRepository]) -> RPCServerWrapper:
    """Instantialise the RPC server wrapper.
    
    Args:
        svc_repo (Type[CommentRepository]): The service repository for the RPC servicers to interface with.

    Returns:
        RPCServerWrapper

    """
    return RPCServerWrapper(svc_repo)