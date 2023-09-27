# Panels

A proof-of-concept forum application created using a microservices architecture.

![Panels Banner](/docs/img-1.png)

## About

This application was created as a hobby project in spare time. In current form the application serves as a proof of concept containing bare-bones functionality.

There are definitely aspects that I'd like to improve or polish further, in addition to some ideas for functionality that I would like to implement at some point. As such, the project should be considered a work-in-progress at this stage.

## Architecture

Users are served the React site from the ``frontend`` and make calls to the REST API exposed by the ``gateway-service``. The ``gateway-service`` makes calls to the relevant services for the request.

Interservice communication is handled through RPC calls (utilising [gRPC](https://grpc.io/)) and [event sourcing](https://microservices.io/patterns/data/event-sourcing.html) (utilising [Kafka](https://kafka.apache.org/)).

![Architecture](/docs/img-2.png)

## Services

I used a variety of languages for writing the services as this was created as a hobby project and the architecture gave me some room to play around with.

| Service | Language | Datastores | Description |
| --- | --- | --- | --- |
| [frontend](/services/frontend) | TypeScript (React) | N/A | Web-based user interface |
| [gateway-service](/services/gateway-service) | Golang | N/A | Exposes a HTTP REST API to allow users to communicate with the application. |
| [panel-service](/services/panel-service) | Golang | [PostgreSQL](https://www.postgresql.org/), [Redis](https://redis.io/) | Responsible for operations related to panels |
| [post-service](/services/post-service) | Golang | [PostgreSQL](https://www.postgresql.org/), [Redis](https://redis.io/) | Responsible for operations related to posts |
| [user-service](/services/user-service) | TypeScript (Node) | [MongoDB](https://www.mongodb.com/) | Responsible for operations related to users |
| [auth-service](/services/auth-service) | Python | [PostgreSQL](https://www.postgresql.org/) | Responsible for authenticating users |
| [comment-service](/services/comment-service) | Python | [PostgreSQL](https://www.postgresql.org/), [Redis](https://redis.io/) | Responsible for operations related to comments |

## Deployment and Configuration

For more information about configuration and deployment, please view the [documentation](/docs/README.md) located in the ``/docs`` folder.

## License

**Acknowledgments:**

* Logo Icon: [Tabler Icons](https://github.com/tabler/tabler-icons) ([MIT License](https://github.com/tabler/tabler-icons/blob/master/LICENSE))
* Logo Font: [Oregano](https://fonts.google.com/specimen/Oregano) ([Open Font License](https://scripts.sil.org/OFL))

This repository is licensed under the [Apache License v2.0](/LICENSE).
