package postgres

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/hexolan/panels/panel-service/internal"
)

func NewPostgresInterface(ctx context.Context, cfg internal.Config) *pgxpool.Pool {
	db, err := pgxpool.New(ctx, cfg.GetPostgresURL())
	if err != nil {
		log.Panic().Err(err).Caller().Msg("")
	}
	
	err = db.Ping(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("failed Postgres ping")
	}

	return db
}