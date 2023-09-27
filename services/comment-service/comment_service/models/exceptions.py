from enum import Enum, auto

from grpc import RpcContext, StatusCode


class ServiceErrorCode(Enum):
    """Error codes used for classifying ServiceExceptions."""
    INVALID_ARGUMENT = auto()
    CONFLICT = auto()
    NOT_FOUND = auto()
    INVALID_CREDENTIALS = auto()
    SERVICE_ERROR = auto()

    __RPC_CODE_MAP__ = {
        INVALID_ARGUMENT: StatusCode.INVALID_ARGUMENT,
        CONFLICT: StatusCode.ALREADY_EXISTS,
        NOT_FOUND: StatusCode.NOT_FOUND,
        INVALID_CREDENTIALS: StatusCode.UNAUTHENTICATED,
        SERVICE_ERROR: StatusCode.INTERNAL
    }

    def to_rpc_code(self) -> StatusCode:
        """Convert a service error code to a gRPC status code.

        Returns:
            The mapped RPC status code, if found, otherwise gRPC Unknown status code.

        """
        return self.__class__.__RPC_CODE_MAP__.get(self.value, StatusCode.UNKNOWN)


class ServiceException(Exception):
    """This exception provides an interface to convert service errors 
    into gRPC errors, which can then be returned to the caller.

    Args:
        msg (str): Error message.
        error_code (ServiceErrorCode): Categorisation code for the error.
    
    Attributes:
        msg (str): The error message.
        error_code (ServiceErrorCode): Categorisation code for the error.
    
    """

    def __init__(self, msg: str, error_code: ServiceErrorCode) -> None:
        super().__init__(msg)
        self.msg = msg
        self.error_code = error_code

    def apply_to_rpc(self, context: RpcContext) -> None:
        """Apply the exception to an RPC context.
        
        Args:
            context (grpc.RpcContext): The context to apply to.
        
        """
        context.set_code(self.error_code.to_rpc_code())
        context.set_details(self.msg)