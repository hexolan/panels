import logging
from typing import Type

import grpc
from grpc_health.v1 import health, health_pb2_grpc

from auth_service.models.proto import auth_pb2_grpc
from auth_service.models.service import AuthRepository
from auth_service.rpc.auth import AuthServicer


class RPCServerWrapper:
    """A wrapper class for the RPC server.

    Attributes:
        _grpc_server (grpc.aio.Server): The gRPC server instance.
    
    """
    def __init__(self, svc_repo: Type[AuthRepository]) -> None:
        """Creates the gRPC server and adds the servicers.
        
        Args:
            svc_repo (Type[AuthRepository]): The service repository to pass to the servicers.

        """
        self._grpc_server = grpc.aio.server()
        self._grpc_server.add_insecure_port("[::]:9090")

        auth_servicer = AuthServicer(svc_repo)
        auth_pb2_grpc.add_AuthServiceServicer_to_server(auth_servicer, self._grpc_server)

        health_servicer = health.aio.HealthServicer()
        health_pb2_grpc.add_HealthServicer_to_server(health_servicer, self._grpc_server)

    async def start(self) -> None:
        """Begin serving RPC asynchronously."""
        logging.info("attempting to serve RPC...")
        await self._grpc_server.start()
        await self._grpc_server.wait_for_termination()


def create_rpc_server(svc_repo: Type[AuthRepository]) -> RPCServerWrapper:
    """Instantialise the RPC server wrapper.
    
    Args:
        svc_repo (Type[AuthRepository]): The service repository for the RPC servicers to interface with.

    Returns:
        RPCServerWrapper

    """
    return RPCServerWrapper(svc_repo)