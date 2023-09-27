package main

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/hexolan/panels/post-service/internal"
	"github.com/hexolan/panels/post-service/internal/postgres"
	"github.com/hexolan/panels/post-service/internal/redis"
	"github.com/hexolan/panels/post-service/internal/kafka"
	"github.com/hexolan/panels/post-service/internal/kafka/producer"
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