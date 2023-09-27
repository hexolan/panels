import logging

from databases import Database

from auth_service.models.config import Config
from auth_service.postgres.repository import ServiceDBRepository


async def connect_database(config: Config) -> Database:
    """Opens a connection to the database.
    
    Args:
        config (Config): The app configuration.

    Returns:
        A connected databases.Database instance

    """
    db = Database(config.postgres_dsn)
    try:
        await db.connect()
    except Exception:
        logging.error("failed to connect to postgresql database")
        raise
    
    return db


async def create_db_repository(config: Config) -> ServiceDBRepository:
    """Create the database repository.
    
    Open a database connection and instantialise the
    database repository.

    Returns:
        ServiceDBRepository
    
    """
    db = await connect_database(config)
    return ServiceDBRepository(db)