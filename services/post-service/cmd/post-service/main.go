// Copyright 2023 Declan Teevan
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/hexolan/panels/post-service/internal"
	"github.com/hexolan/panels/post-service/internal/kafka"
	"github.com/hexolan/panels/post-service/internal/kafka/producer"
	"github.com/hexolan/panels/post-service/internal/postgres"
	"github.com/hexolan/panels/post-service/internal/redis"
	"github.com/hexolan/panels/post-service/internal/rpc"
	"github.com/hexolan/panels/post-service/internal/service"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Load the configuration
	cfg := internal.NewConfig()
	zerolog.SetGlobalLevel(cfg.GetLogLevel())

	// Create the interface dependencies
	ctx := context.Background()
	db := postgres.NewPostgresInterface(ctx, cfg)
	rdb := redis.NewRedisInterface(ctx, cfg)
	events := producer.NewPostEventProducer(cfg)

	// Create the repositories and services
	dbRepo := postgres.NewPostRepository(db)
	cacheRepo := redis.NewPostRepository(rdb, dbRepo)
	service := service.NewPostService(events, cacheRepo)

	// Start the event consumers
	eventConsumers := kafka.NewEventConsumers(cfg, dbRepo, events)
	eventConsumers.Start()

	// Create and serve RPC
	rpcServer := rpc.NewRPCServer(service)
	rpcServer.Serve()
}
