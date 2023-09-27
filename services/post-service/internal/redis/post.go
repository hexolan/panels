package redis

import (
	"time"
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"
	"github.com/redis/go-redis/v9"

	"github.com/hexolan/panels/post-service/internal"
)

type postCacheRepo struct {
	rdb *redis.Client

	repo internal.PostRepository
}

func NewPostRepository(rdb *redis.Client, repo internal.PostRepository) internal.PostRepository {
	return postCacheRepo{
		rdb: rdb,
		repo: repo,
	}
}

func (r postCacheRepo) getCachedPost(ctx context.Context, id internal.PostId) *internal.Post {
	value, err := r.rdb.Get(ctx, id.GetReprId()).Result()
	if err == redis.Nil {
		return nil
	} else if err != nil {
		log.Error().Err(err).Msg("failed to get cached post")
		return nil
	}

	var post internal.Post
	err = json.Unmarshal([]byte(value), &post)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal cached post")
		return nil
	}

	return &post
}

func (r postCacheRepo) cachePost(ctx context.Context, post *internal.Post) {
	value, err := json.Marshal(post)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal post for caching")
		return
	}

	err = r.rdb.Set(ctx, post.Id.GetReprId(), string(value), 2 * time.Minute).Err()
	if err != nil {
		log.Error().Err(err).Msg("failed to cache post")
		return
	}
}

func (r postCacheRepo) purgeCachedPost(ctx context.Context, id internal.PostId) {
	err := r.rdb.Del(ctx, id.GetReprId()).Err()
	if err != nil && err != redis.Nil {
		log.Error().Err(err).Msg("error while purging cached post")
	}
}

func (r postCacheRepo) CreatePost(ctx context.Context, panelId string, authorId string, data internal.PostCreate) (*internal.Post, error) {
	// Create the post (using downstream DB repo)
	post, err := r.repo.CreatePost(ctx, panelId, authorId, data)
	if err != nil {
		return post, err
	}

	// Cache and return the created post.
	r.cachePost(ctx, post)
	return post, err
}

func (r postCacheRepo) GetPost(ctx context.Context, id internal.PostId) (*internal.Post, error) {
	// Check for a cached version of the post
	if post := r.getCachedPost(ctx, id); post != nil {
		return post, nil
	}

	// Post is not cached (fetch from DB)
	post, err := r.repo.GetPost(ctx, id)
	if err != nil {
		return post, err
	}

	// Cache and return the fetched post
	r.cachePost(ctx, post)
	return post, err
}

func (r postCacheRepo) GetPanelPost(ctx context.Context, id internal.PostId, panelId string) (*internal.Post, error) {
	// Check for a cached version of the post
	if post := r.getCachedPost(ctx, id); post != nil {
		// The post is cached. Ensure panelId is a match.
		if post.PanelId != panelId {
			return nil, internal.NewServiceError(internal.NotFoundErrorCode, "post not found on that panel")
		}
		return post, nil
	}

	// Post is not cached (fetch from DB)
	post, err := r.repo.GetPanelPost(ctx, id, panelId)
	if err != nil {
		return post, err
	}

	// Cache and return the fetched post
	r.cachePost(ctx, post)
	return post, err
}

func (r postCacheRepo) UpdatePost(ctx context.Context, id internal.PostId, data internal.PostUpdate) (*internal.Post, error) {
	// Update the post at the downstream repo.
	post, err := r.repo.UpdatePost(ctx, id, data)
	if err != nil {
		return post, err
	}

	// Cache and return the updated post.
	r.cachePost(ctx, post)
	return post, err
}

func (r postCacheRepo) DeletePost(ctx context.Context, id internal.PostId) error {
	// Delete the post downstream.
	err := r.repo.DeletePost(ctx, id)
	if err != nil {
		return err
	}

	// Purge any cached version of the post.
	r.purgeCachedPost(ctx, id)
	return err
}

func (r postCacheRepo) GetFeedPosts(ctx context.Context) ([]*internal.Post, error) {
	return r.repo.GetFeedPosts(ctx)
}

func (r postCacheRepo) GetUserPosts(ctx context.Context, userId string) ([]*internal.Post, error) {
	return r.repo.GetUserPosts(ctx, userId)
}

func (r postCacheRepo) GetPanelPosts(ctx context.Context, panelId string) ([]*internal.Post, error) {
	return r.repo.GetPanelPosts(ctx, panelId)
}