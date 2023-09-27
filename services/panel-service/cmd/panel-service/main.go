package main

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/hexolan/panels/panel-service/internal"
	"github.com/hexolan/panels/panel-service/internal/postgres"
	"github.com/hexolan/panels/panel-service/internal/redis"
	"github.com/hexolan/panels/panel-service/internal/kafka"
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