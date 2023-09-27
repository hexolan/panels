package internal

import (
	"os"
	"fmt"
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
		RedisHost: optFromEnvRequire("REDIS_HOST"),
		RedisPass: optFromEnvRequire("REDIS_PASS"),
		KafkaBrokers: kafkaBrokers,
		LogLevel: logLvl,
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