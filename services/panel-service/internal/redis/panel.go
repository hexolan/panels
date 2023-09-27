package redis

import (
	"time"
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"
	"github.com/redis/go-redis/v9"

	"github.com/hexolan/panels/panel-service/internal"
)

type panelCacheRepo struct {
	rdb *redis.Client

	repo internal.PanelRepository
}

func NewPanelRepository(rdb *redis.Client, repo internal.PanelRepository) internal.PanelRepository {
	return panelCacheRepo{
		rdb: rdb,
		repo: repo,
	}
}

func (r panelCacheRepo) getCachedPanel(ctx context.Context, id int64) *internal.Panel {
	value, err := r.rdb.Get(ctx, internal.StringifyPanelId(id)).Result()
	if err == redis.Nil {
		return nil
	} else if err != nil {
		log.Error().Err(err).Msg("failed to get cached panel")
		return nil
	}

	var panel internal.Panel
	err = json.Unmarshal([]byte(value), &panel)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal cached panel")
		return nil
	}

	return &panel
}

func (r panelCacheRepo) purgeCachedPanel(ctx context.Context, id int64) {
	err := r.rdb.Del(ctx, internal.StringifyPanelId(id)).Err()
	if err != nil && err != redis.Nil {
		log.Error().Err(err).Msg("error while purging cached panel")
	}
}

func (r panelCacheRepo) cachePanel(ctx context.Context, panel *internal.Panel) {
	value, err := json.Marshal(panel)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal panel for caching")
		return
	}

	err = r.rdb.Set(ctx, internal.StringifyPanelId(panel.Id), string(value), 5 * time.Minute).Err()
	if err != nil {
		log.Error().Err(err).Msg("failed to cache panel")
		return
	}
}

func (r panelCacheRepo) GetPanelIdFromName(ctx context.Context, name string) (*int64, error) {
	// This is not cached for safety with UpdatePanel and DeletePanel methods.
	return r.repo.GetPanelIdFromName(ctx, name)
}

func (r panelCacheRepo) CreatePanel(ctx context.Context, data internal.PanelCreate) (*internal.Panel, error) {
	// Create the panel with the downstream DB repo.
	panel, err := r.repo.CreatePanel(ctx, data)
	if err != nil {
		return panel, err
	}

	// Cache and return the created panel.
	r.cachePanel(ctx, panel)
	return panel, err
}

func (r panelCacheRepo) GetPanel(ctx context.Context, id int64) (*internal.Panel, error) {
	// Check for a cached version of the panel.
	if panel := r.getCachedPanel(ctx, id); panel != nil {
		return panel, nil
	}

	// Panel is not cached. Fetch from the DB repo.
	panel, err := r.repo.GetPanel(ctx, id)
	if err != nil {
		return panel, err
	}

	// Cache and return the fetched panel.
	r.cachePanel(ctx, panel)
	return panel, err
}

func (r panelCacheRepo) UpdatePanel(ctx context.Context, id int64, data internal.PanelUpdate) (*internal.Panel, error) {
	// Update the panel at the downstream repo.
	panel, err := r.repo.UpdatePanel(ctx, id, data)
	if err != nil {
		return panel, err
	}

	// Cache and return the updated panel.
	r.cachePanel(ctx, panel)
	return panel, err
}

func (r panelCacheRepo) DeletePanel(ctx context.Context, id int64) error {
	// Delete the panel downstream.
	err := r.repo.DeletePanel(ctx, id)
	if err != nil {
		return err
	}

	// Purge any cached version of the panel.
	r.purgeCachedPanel(ctx, id)
	return err
}