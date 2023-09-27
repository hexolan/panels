migration-new:
	migrate create -ext sql -dir internal/postgres/migrations -seq ${MIGRATION_NAME}

migration-upgrade:
	migrate -path internal/postgres/migrations/ -database "postgres://${POSTGRES_USER}:${POSTGRES_PASS}@${POSTGRES_HOST}/${POSTGRES_DATABASE}?sslmode=disable" up

migration-downgrade:
	migrate -path internal/postgres/migrations/ -database "postgres://${POSTGRES_USER}:${POSTGRES_PASS}@${POSTGRES_HOST}/${POSTGRES_DATABASE}?sslmode=disable" down

protobufs-compile:
	protoc --proto_path=../../protobufs/ --go_out=. --go_opt=Mpost.proto=./internal/rpc/postv1 --go-grpc_out=. --go-grpc_opt=Mpost.proto=./internal/rpc/postv1 post.proto
	protoc --proto_path=../../protobufs/ --go_out=. --go_opt=Muser.proto=./internal/kafka/consumers/userv1 --go-grpc_out=. --go-grpc_opt=Muser.proto=./internal/kafka/consumers/userv1 user.proto
	protoc --proto_path=../../protobufs/ --go_out=. --go_opt=Mpanel.proto=./internal/kafka/consumers/panelv1 --go-grpc_out=. --go-grpc_opt=Mpanel.proto=./internal/kafka/consumers/panelv1 panel.proto