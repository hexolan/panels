migration-new:
	migrate create -ext sql -dir ./auth_service/postgres/migrations -seq ${MIGRATION_NAME}

migration-upgrade:
	docker run -v "${SERVICE_DIR}/auth_service/postgres/migrations:/migrations" --network host migrate/migrate:4 -path=/migrations/ -database postgresql://${POSTGRES_USER}:${POSTGRES_PASS}@${POSTGRES_HOST}/${POSTGRES_DATABASE}?sslmode=disable up

migration-downgrade:
	docker run -v "${SERVICE_DIR}/auth_service/postgres/migrations:/migrations" --network host migrate/migrate:4 -path=/migrations/ -database postgresql://${POSTGRES_USER}:${POSTGRES_PASS}@${POSTGRES_HOST}/${POSTGRES_DATABASE}?sslmode=disable down

protobufs-compile:
	poetry run python -m grpc_tools.protoc -I../../protobufs --python_out=./auth_service/models/proto --pyi_out=./auth_service/models/proto --grpc_python_out=./auth_service/models/proto auth.proto user.proto
	poetry run python -m protoletariat --dont-create-package --in-place --python-out ./auth_service/models/proto protoc --proto-path=../../protobufs auth.proto user.proto