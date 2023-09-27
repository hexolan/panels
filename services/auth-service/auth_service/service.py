import hmac
import hashlib
from typing import Type

import jwt
from argon2 import PasswordHasher
from argon2.exceptions import VerificationError

from auth_service.models.config import Config
from auth_service.models.exceptions import ServiceException, ServiceErrorCode
from auth_service.models.service import AuthRepository, AuthDBRepository, AuthToken, AccessTokenClaims


class ServiceRepository(AuthRepository):
    """The service repository responsible for handling
    the business logic.

    As I've developed most of the services in this project
    using a repository pattern, this provides a level
    of abstraction, allowing me to swap out the downstream
    repositories to add additional layers.

    For example, instead of have the next downstream
    repository (`_repo` attribute) point directly
    to the DB repo instance, I could point it to a
    Redis repository to use as a caching layer, as I
    have done so in some of the other services in this
    project.

    Attributes:
        _repo (AuthDBRepository): The downstream database repository.
        _jwt_private (str): RSA private key for producing JWT tokens.
        _hasher (argon2.PasswordHasher): Library class utilised for hashing passwords.
        _pepper (bytes): An additional 'secret salt' applied to all passwords when hashed.
    
    """

    def __init__(self, config: Config, downstream_repo: Type[AuthDBRepository]) -> None:
        self._repo = downstream_repo

        self._jwt_private = config.jwt_private_key

        self._hasher = PasswordHasher()  # Use default RFC_9106_LOW_MEMORY profile
        self._pepper = bytes(config.password_pepper, encoding="utf-8")

    def _apply_pepper(self, password: str) -> str:
        """Apply the pepper ('secret salt') to
        a given password using HMAC.

        Args:
            password (str): The password to apply to pepper to.

        Returns:
            str: The password with applied pepper (in hexadecimal form). 
        
        """
        return hmac.new(key=self._pepper, msg=bytes(password, encoding="utf-8"), digestmod=hashlib.sha256).hexdigest()

    def _hash_password(self, password: str) -> str:
        """Hash a given password.

        The pepper is applied to the password, and the result
        is then hashed using Argon2id.
        
        Args:
            password (str): The password to hash.

        Returns:
            str: The hashed password.
        
        """
        return self._hasher.hash(self._apply_pepper(password))

    def _verify_password(self, hash: str, password: str) -> None:
        """Verify that an input password matches
        a hash.

        The pepper is applied to the input password
        and compared to the stored hash.

        Args:
            hash (str): The existing hashed password.
            password (str): The password to verify against the hash.

        Raises:
            argon2.Exceptions.VerificationError: if the password does not match
        
        """
        self._hasher.verify(hash, self._apply_pepper(password))

    def _issue_auth_token(self, user_id: str) -> AuthToken:
        """Issue an auth token.
        
        Args:
            user_id (str): The user to issue the tokens to.

        Returns:
            AuthToken: A response containing the generated access token,
                token type and expiry time.        
        """
        claims = AccessTokenClaims(sub=user_id)
        access_token = jwt.encode(claims.model_dump(), self._jwt_private, algorithm="RS256")
        return AuthToken(access_token=access_token)

    async def auth_with_password(self, user_id: str, password: str) -> AuthToken:
        """Attempt to authenticate a user with a
        provided password.

        Args:
            user_id (str): The user to attempt to authenticate.
            password (str): The password to verify against.
         
        """
        # Get the auth record for that user
        auth_record = await self._repo.get_auth_record(user_id)
        if not auth_record:
            raise ServiceException("invalid user id or password", ServiceErrorCode.INVALID_CREDENTIALS)

        # Verify the password hashes
        try:
            self._verify_password(auth_record.password.get_secret_value(), password)
        except VerificationError:
            raise ServiceException("invalid user id or password", ServiceErrorCode.INVALID_CREDENTIALS)
        
        # Update the auth record if the password needs rehash
        if self._hasher.check_needs_rehash(auth_record.password.get_secret_value()):
            await self._repo.update_password_auth_method(auth_record.user_id, self._hash_password(password))

        # Issue a token for the user
        return self._issue_auth_token(auth_record.user_id)

    async def set_password_auth_method(self, user_id: str, password: str) -> None:
        """Set a user's password for use when authenticating.
        
        If the specified user does not already have an existing
        authentication method, then insert a record, otherwise
        update their existing record.

        Args:
            user_id (str): The referenced user.
            password (str): The new password for that user.
        
        """
        # Hash the password
        password = self._hash_password(password)

        # Update the auth method or create one if it doesn't exist
        auth_record = await self._repo.get_auth_record(user_id)
        if auth_record is not None:
            await self._repo.update_password_auth_method(user_id, password)
        else:
            await self._repo.create_password_auth_method(user_id, password)

    async def delete_password_auth_method(self, user_id: str) -> None:
        """Delete a user's authentication method.

        Args:
            user_id (str): The referenced user.        
        
        """
        await self._repo.delete_password_auth_method(user_id)