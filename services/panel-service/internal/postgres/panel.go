package postgres

import (
	"context"
	"errors"
	"strings"
	"encoding/json"

	"github.com/rs/zerolog/log"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/hexolan/panels/panel-service/internal"
)

type panelDatabaseRepo struct {
	db *pgxpool.Pool
}

func NewPanelRepository(db *pgxpool.Pool) internal.PanelRepository {
	return panelDatabaseRepo{
		db: db,
	}
}

func (r panelDatabaseRepo) transformToPatchData(data internal.PanelUpdate) goqu.Record {
	// Ensure updated_at field is changed
	patchData := goqu.Record{"updated_at": goqu.L("timezone('utc', now())")}

	// Marshal the data to remove omitted keys
	marshalled, _ := json.Marshal(data)
	_ = json.Unmarshal(marshalled, &patchData)
	return patchData
}

func (r panelDatabaseRepo) GetPanelIdFromName(ctx context.Context, name string) (*int64, error) {
	var id int64
	err := r.db.QueryRow(ctx, "SELECT id FROM panels WHERE LOWER(name)=LOWER($1)", name).Scan(&id)
    if err != nil {
		if err == pgx.ErrNoRows {
			return nil, internal.WrapServiceError(err, internal.NotFoundErrorCode, "panel not found")
		} else if strings.Contains(err.Error(), "failed to connect to") {
			return nil, internal.WrapServiceError(err, internal.ConnectionErrorCode, "failed to connect to database")
		}
		log.Error().Err(err).Msg("unaccounted error whilst getting panel ID from name")
		return nil, internal.WrapServiceError(err, internal.UnknownErrorCode, "failed to get panel")
	}

	return &id, nil
}

func (r panelDatabaseRepo) CreatePanel(ctx context.Context, data internal.PanelCreate) (*internal.Panel, error) {
	var id int64
	err := r.db.QueryRow(ctx, "INSERT INTO panels (name, description) VALUES ($1, $2) RETURNING id", data.Name, data.Description).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
				return nil, internal.WrapServiceError(err, internal.ConflictErrorCode, "panel name not unique")
			}
		} else if strings.Contains(err.Error(), "failed to connect to") {
			return nil, internal.WrapServiceError(err, internal.ConnectionErrorCode, "failed to connect to database")
		}
		log.Error().Err(err).Msg("unaccounted error whilst creating panel")
		return nil, internal.WrapServiceError(err, internal.UnknownErrorCode, "failed to create panel")
	}

	return r.GetPanel(ctx, id)
}

func (r panelDatabaseRepo) GetPanel(ctx context.Context, id int64) (*internal.Panel, error) {
	var panel internal.Panel
	row := r.db.QueryRow(ctx, "SELECT id, name, description, created_at, updated_at FROM panels WHERE id=$1", id)
	err := row.Scan(&panel.Id, &panel.Name, &panel.Description, &panel.CreatedAt, &panel.UpdatedAt)
    if err != nil {
		if err == pgx.ErrNoRows {
			return nil, internal.WrapServiceError(err, internal.NotFoundErrorCode, "panel not found")
		} else if strings.Contains(err.Error(), "failed to connect to") {
			return nil, internal.WrapServiceError(err, internal.ConnectionErrorCode, "failed to connect to database")
		}
		log.Error().Err(err).Msg("unaccounted error whilst getting panel")
		return nil, internal.WrapServiceError(err, internal.UnknownErrorCode, "failed to get panel")
	}

	return &panel, nil
}

func (r panelDatabaseRepo) UpdatePanel(ctx context.Context, id int64, data internal.PanelUpdate) (*internal.Panel, error) {
	patchData := r.transformToPatchData(data)

	// Build a statement to update the panel
	statement, args, _ := goqu.Dialect("postgres").Update("panels").Prepared(true).Set(patchData).Where(goqu.C("id").Eq(id)).ToSQL()

	// Execute the query
	result, err := r.db.Exec(ctx, statement, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
				return nil, internal.WrapServiceError(err, internal.ConflictErrorCode, "panel name not unique")
			}
		} else if strings.Contains(err.Error(), "failed to connect to") {
			return nil, internal.WrapServiceError(err, internal.ConnectionErrorCode, "failed to connect to database")
		}
		log.Error().Err(err).Msg("unaccounted error whilst updating panel")
		return nil, internal.WrapServiceError(err, internal.UnknownErrorCode, "failed to update panel")
	}

	// Check if any rows were affected from the query
	rows_affected := result.RowsAffected()
	if rows_affected != 1 {
		return nil, internal.NewServiceError(internal.NotFoundErrorCode, "panel not found")
	}

	return r.GetPanel(ctx, id)
}

func (r panelDatabaseRepo) DeletePanel(ctx context.Context, id int64) error {
	// Attempt to delete the panel
	result, err := r.db.Exec(ctx, "DELETE FROM panels WHERE id=$1", id)
	if err != nil {
		if strings.Contains(err.Error(), "failed to connect to") {
			return internal.WrapServiceError(err, internal.ConnectionErrorCode, "failed to connect to database")
		}
		log.Error().Err(err).Msg("unaccounted error whilst deleting panel")
		return internal.WrapServiceError(err, internal.UnknownErrorCode, "failed to delete panel")
	}

	// Check if any rows were affected
	rows_affected := result.RowsAffected()
	if rows_affected != 1 {
		return internal.NewServiceError(internal.NotFoundErrorCode, "panel not found")
	}

	return nil
}