//go:build tools
package main

import (
	// Protobuf Generation
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"

	// DB Migrations
    _ "github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/postgres"
)