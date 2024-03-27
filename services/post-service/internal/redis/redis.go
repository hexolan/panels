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

package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"

	"github.com/hexolan/panels/post-service/internal"
)

func NewRedisInterface(ctx context.Context, cfg internal.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisHost,
		Password: cfg.RedisPass,
		DB:       0,

		DialTimeout: time.Millisecond * 250,
		ReadTimeout: time.Millisecond * 500,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Warn().Err(err).Msg("failed Redis ping")
	}

	return rdb
}
