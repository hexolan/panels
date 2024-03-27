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

	"github.com/hexolan/panels/panel-service/internal"
	"github.com/hexolan/panels/panel-service/internal/kafka"
	"github.com/hexolan/panels/panel-service/internal/postgres"
	"github.com/hexolan/panels/panel-service/internal/redis"
	"github.com/hexolan/panels/panel-service/internal/rpc"
	"github.com/hexolan/panels/panel-service/internal/service"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Load the configuration
	cfg := internal.NewConfig()
	zerolog.SetGlobalLevel(cfg.GetLogLevel())

	// Loading the dependencies
	ctx := context.Background()
	db := postgres.NewPostgresInterface(ctx, cfg)
	rdb := redis.NewRedisInterface(ctx, cfg)
	events := kafka.NewPanelEventProducer(cfg)

	// Create the repositories and services
	databaseRepo := postgres.NewPanelRepository(db)
	cacheRepo := redis.NewPanelRepository(rdb, databaseRepo)
	service := service.NewPanelService(events, cacheRepo)

	// Create and serve RPC
	rpcServer := rpc.NewRPCServer(service)
	rpcServer.Serve()
}
