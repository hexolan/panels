package redis

import (
	"time"
	"context"

	"github.com/rs/zerolog/log"
	"github.com/redis/go-redis/v9"

	"github.com/hexolan/panels/post-service/internal"
)

func NewRedisInterface(ctx context.Context, cfg internal.Config) *redis.Client {
    rdb := redis.NewClient(&redis.Options{
        Addr: cfg.RedisHost,
        Password: cfg.RedisPass,
        DB: 0,

		DialTimeout: time.Millisecond * 250,
		ReadTimeout: time.Millisecond * 500,
    })

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Warn().Err(err).Msg("failed Redis ping")
	}

	return rdb
}