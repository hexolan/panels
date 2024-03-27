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

package internal

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewConfig() Config {
	// Parse the log level
	logLvl, err := zerolog.ParseLevel(optFromEnvFallback("LOG_LEVEL", "info"))
	if err != nil {
		log.Fatal().Err(err).Msg("invalid log level specified")
	}

	// Parse the kafka brokers
	kafkaBrokers := strings.Split(optFromEnvRequire("KAFKA_BROKERS"), ",")
	if len(kafkaBrokers) == 0 {
		log.Fatal().Err(err).Msg("no kafka brokers provided in configuration")
	}

	// Create the config
	cfg := Config{
		RedisHost:    optFromEnvRequire("REDIS_HOST"),
		RedisPass:    optFromEnvRequire("REDIS_PASS"),
		KafkaBrokers: kafkaBrokers,
		LogLevel:     logLvl,
	}

	// Assemble the Config.PostgresURL
	cfg.SetPostgresURL(
		optFromEnvRequire("POSTGRES_USER"),
		optFromEnvRequire("POSTGRES_PASS"),
		optFromEnvRequire("POSTGRES_HOST"),
		optFromEnvRequire("POSTGRES_DATABASE"),
	)

	return cfg
}

func optFromEnv(opt string) *string {
	optValue, exists := os.LookupEnv(opt)
	if !exists || optValue == "" {
		return nil
	}
	return &optValue
}

func optFromEnvRequire(opt string) string {
	optValue := optFromEnv(opt)
	if optValue == nil {
		log.Fatal().Str("option", opt).Msg("failed to load required config option")
	}
	return *optValue
}

func optFromEnvFallback(opt string, fallback string) string {
	optValue := optFromEnv(opt)
	if optValue == nil {
		return fallback
	}
	return *optValue
}

type Config struct {
	PostgresURL string

	RedisHost string
	RedisPass string

	KafkaBrokers []string

	LogLevel zerolog.Level
}

func (cfg *Config) SetPostgresURL(user string, pass string, host string, db string) {
	cfg.PostgresURL = fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", user, pass, host, db)
}

func (cfg Config) GetPostgresURL() string {
	return cfg.PostgresURL
}

func (cfg Config) GetLogLevel() zerolog.Level {
	return cfg.LogLevel
}
