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

import base64
from typing import Any, List

from pydantic import computed_field
from pydantic.fields import FieldInfo
from pydantic_settings import BaseSettings, EnvSettingsSource
from pydantic_settings.main import BaseSettings
from pydantic_settings.sources import PydanticBaseSettingsSource


class ConfigSource(EnvSettingsSource):
    """Responsible for loading config options from environment variables."""
    def prepare_field_value(self, field_name: str, field: FieldInfo, value: Any, value_is_complex: bool) -> Any:
        if field_name == "kafka_brokers":
            # Comma delimit the kafka brokers.
            if value == None:
                return None
            return value.split(",")
        elif field_name == "jwt_public_key" or field_name == "jwt_private_key":
            # Decode the JWT public and private keys from base64.
            if value == None:
                return None
            return base64.standard_b64decode(value).decode(encoding="utf-8")

        return super().prepare_field_value(field_name, field, value, value_is_complex)


class Config(BaseSettings):
    """The service configuration loaded from environment
    variables.
    
    Attributes:
        postgres_user (str): Loaded from the 'POSTGRES_USER' envvar.
        postgres_pass (str): Loaded from the 'POSTGRES_PASS' envvar.
        postgres_host (str): Loaded from the 'POSTGRES_HOST' envvar.
        postgres_database (str): Loaded from the 'POSTGRES_DATABASE' envvar.
        kafka_brokers (list[str]): Loaded and comma delmited from the 'KAFKA_BROKERS' envvar.
        jwt_public_key (str): Loaded and decoded, from base64, from the 'JWT_PUBLIC_KEY' envvar.
        jwt_private_key (str): Loaded and decoded, from base64, from the 'JWT_PRIVATE_KEY' envvar.
        password_pepper (str): Loaded from the 'PASSWORD_PEPPER' envvar.
        postgres_dsn (str): Computed when accessed the first time. (@property)
    
    """
    postgres_user: str
    postgres_pass: str
    postgres_host: str
    postgres_database: str

    kafka_brokers: List[str]
    
    jwt_public_key: str
    jwt_private_key: str
    
    password_pepper: str

    @computed_field
    @property
    def postgres_dsn(self) -> str:
        """Uses the postgres_user, postgres_pass, postgres_host,
        and postgres_database options to assemble a DSN.

        Returns:
            str: DSN for connecting to the database.
        
        """
        return "postgresql+asyncpg://{user}:{password}@{host}/{db}".format(
            user=self.postgres_user,
            password=self.postgres_pass,
            host=self.postgres_host,
            db=self.postgres_database
        )
    
    @classmethod
    def settings_customise_sources(cls, settings_cls: type[BaseSettings], *args, **kwargs) -> tuple[PydanticBaseSettingsSource, ...]:
        return (ConfigSource(settings_cls), )