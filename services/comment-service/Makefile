migration-new:
	migrate create -ext sql -dir ./comment_service/postgres/migrations -seq ${MIGRATION_NAME}

migration-upgrade:
	docker run -v "${SERVICE_DIR}/comment_service/postgres/migrations:/migrations" --network host migrate/migrate:4 -path=/migrations/ -database postgresql://${POSTGRES_USER}:${POSTGRES_PASS}@${POSTGRES_HOST}/${POSTGRES_DATABASE}?sslmode=disable up

migration-downgrade:
	docker run -v "${SERVICE_DIR}/comment_service/postgres/migrations:/migrations" --network host migrate/migrate:4 -path=/migrations/ -database postgresql://${POSTGRES_USER}:${POSTGRES_PASS}@${POSTGRES_HOST}/${POSTGRES_DATABASE}?sslmode=disable down

protobufs-compile:
	poetry run python -m grpc_tools.protoc -I../../protobufs --python_out=./comment_service/models/proto --pyi_out=./comment_service/models/proto --grpc_python_out=./comment_service/models/proto comment.proto post.proto user.proto
	poetry run python -m protoletariat --dont-create-package --in-place --python-out ./comment_service/models/proto protoc --proto-path=../../protobufs comment.proto post.proto user.proto