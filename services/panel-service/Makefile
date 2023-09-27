migration-new:
	migrate create -ext sql -dir internal/postgres/migrations -seq ${MIGRATION_NAME}

migration-upgrade:
	migrate -path internal/postgres/migrations/ -database "postgres://${POSTGRES_USER}:${POSTGRES_PASS}@${POSTGRES_HOST}/${POSTGRES_DATABASE}?sslmode=disable" up

migration-downgrade:
	migrate -path internal/postgres/migrations/ -database "postgres://${POSTGRES_USER}:${POSTGRES_PASS}@${POSTGRES_HOST}/${POSTGRES_DATABASE}?sslmode=disable" down
	
protobufs-compile:
	protoc --proto_path=../../protobufs/ --go_out=. --go_opt=Mpanel.proto=./internal/rpc/panelv1 --go-grpc_out=. --go-grpc_opt=Mpanel.proto=./internal/rpc/panelv1 panel.proto