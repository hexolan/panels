import logging
import traceback
from typing import Type

from google.protobuf import empty_pb2
from grpc import RpcContext, StatusCode

from auth_service.models.exceptions import ServiceException
from auth_service.models.service import AuthRepository, AuthToken
from auth_service.models.proto import auth_pb2, auth_pb2_grpc


class AuthServicer(auth_pb2_grpc.AuthServiceServicer):
    """Contains definitions for the service's RPC methods.
    
    The request attributes are validated and translated 
    from protobufs into business model form, then passed
    along to the service repository to handling the
    business logic.

    Responses from calls to methods in the service repository
    are then translated back to protobuf form to return
    to the user.

    Attributes:
        _svc_repo (Type[AuthRepository]): The highest level service repository.
    
    """
    def __init__(self, svc_repo: Type[AuthRepository]) -> None:
        self._svc_repo = svc_repo
    
    def _apply_error(self, context: RpcContext, code: StatusCode, msg: str) -> None:
        """Set an error on a given RPC request.
        
        Args:
            context (grpc.RpcContext): The context to apply the error to.
            code (grpc.StatusCode): The gRPC status code.
            msg (str): The error details.

        """
        context.set_code(code)
        context.set_details(msg)
    
    def _apply_unknown_error(self, context: RpcContext) -> None:
        """Apply a de facto error fallback message.
        
        Args:
            context (grpc.RpcContext): The context to apply the error to.
        
        """
        self._apply_error(context, StatusCode.UNKNOWN, "unknown error occured")

    async def AuthWithPassword(self, request: auth_pb2.PasswordAuthRequest, context: RpcContext) -> auth_pb2.AuthToken:
        """AuthWithPassword RPC Call
        
        Args:
            request (auth_pb2.PasswordAuthRequest): The request parameters.
            context (grpc.RpcContext): The context of the RPC call.

        Returns:
            auth_pb2.AuthToken: With a succesfully authentication.

        """
        # validate the request inputs
        if request.user_id == "":
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                msg="user not provided"
            )
            return

        if request.password == "":
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                msg="password not provided"
            )
            return

        if len(request.password) < 8:
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                msg="invalid password"
            )
            return

        # Attempt to authenticate the user
        try:
            token = await self._svc_repo.auth_with_password(request.user_id, request.password)
        except ServiceException as err:
            err.apply_to_rpc(context)
            return
        except Exception:
            logging.error(traceback.format_exc())
            self._apply_unknown_error(context)
            return
        
        # Convert token to protobuf
        return AuthToken.to_protobuf(token)

    async def SetPasswordAuth(self, request: auth_pb2.SetPasswordAuthMethod, context: RpcContext) -> empty_pb2.Empty:
        """SetPasswordAuth RPC Call
        
        Args:
            request (auth_pb2.SetPasswordAuthMethod): The request parameters.
            context (grpc.RpcContext): The context of the RPC call.

        Returns:
            empty_pb2.Empty: Empty protobuf response (in effect returns None).

        """
        # validate the request inputs
        if request.user_id == "":
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                msg="user id not provided"
            )
            return
        
        if request.password == "":
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                msg="password not provided"
            )
            return

        if len(request.password) < 8:
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                msg="password must be at least 8 characters"
            )
            return

        # Attempt to create the auth method
        try:
            await self._svc_repo.set_password_auth_method(request.user_id, request.password)
        except ServiceException as err:
            err.apply_to_rpc(context)
            return
        except Exception:
            logging.error(traceback.format_exc())
            self._apply_unknown_error(context)
            return
        
        # Success
        return empty_pb2.Empty()

    async def DeletePasswordAuth(self, request: auth_pb2.DeletePasswordAuthMethod, context: RpcContext) -> empty_pb2.Empty:
        """DeletePasswordAuth RPC Call
        
        Args:
            request (auth_pb2.DeletePasswordAuthMethod): The request parameters.
            context (grpc.RpcContext): The context of the RPC call.

        Returns:
            empty_pb2.Empty: Empty protobuf response (in effect returns None).

        """
        # Ensure a user id is provided
        if request.user_id == "":
            self._apply_error(
                context,
                code=StatusCode.INVALID_ARGUMENT,
                msg="user id not provided"
            )
            return
        
        # Attempt to delete the auth method
        try:
            await self._svc_repo.delete_password_auth_method(request.user_id, request.password)
        except ServiceException as err:
            err.apply_to_rpc(context)
            return
        except Exception:
            logging.error(traceback.format_exc())
            self._apply_unknown_error(context)
            return
        
        # Success
        return empty_pb2.Empty()