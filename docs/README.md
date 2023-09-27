# Panels Documentation

## Contributing

This project is a work-in-progress hobby project, however contributions are welcome. Please feel free to open issues or make a pull request.

## Deployment

### Using [Kubernetes](https://kubernetes.io/):

The application can be deployed to Kubernetes using [Skaffold](https://skaffold.dev/):

> skaffold run

Note that when deploying to Kubernetes you will need to have deployed Redis, Postgres and Mongo instances off cluster and adjusted your configuration for the services.

### Using [Docker](https://www.docker.com/):

The application can be easily deployed with [Docker Compose](https://docs.docker.com/compose/) using the following command from the root directory:

> docker compose up

When deploying using Docker Compose, the default container configuration (exposed as environment variables in [docker-compose.yaml](/docker-compose.yaml)) can be left as is.

This is presuming that the [docker-compose.override.yaml](/docker-compose.override.yaml) file, which contains specification for the instances that each service requires, is also being used. However, if solely the [docker-compose.yaml](/docker-compose.yaml) is being used then the configuration will need to be changed to point to your instances of the databases.

## Configuration

For an outline on the environment variables that each service requires, or events that are produced/consumed by the services, view the documentation for the individual services (located in the ``README.md`` files of each service folder). 

Here is a table for easy access:

| Service | Documentation | Example Configuration |
| --- | --- | --- |
| [frontend](/services/frontend) | [README.md](/services/frontend/README.md) | [.env.example](/services/frontend/.env.example) |
| [gateway-service](/services/gateway-service) | [README.md](/services/gateway-service/README.md) | [.env.example](/services/gateway-service/.env.example) |
| [panel-service](/services/panel-service) | [README.md](/services/panel-service/README.md) | [.env.example](/services/panel-service/.env.example) |
| [post-service](/services/post-service) | [README.md](/services/post-service/README.md) | [.env.example](/services/post-service/.env.example) |
| [user-service](/services/user-service) | [README.md](/services/user-service/README.md) | [.env.example](/services/user-service/.env.example) |
| [auth-service](/services/auth-service) | [README.md](/services/auth-service/README.md) | [.env.example](/services/auth-service/.env.example) |
| [comment-service](/services/comment-service) | [README.md](/services/comment-service/README.md) | [.env.example](/services/comment-service/.env.example) |