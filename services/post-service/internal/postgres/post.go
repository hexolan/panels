package postgres

import (
	"context"
	"strings"
	"encoding/json"

	"github.com/rs/zerolog/log"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/hexolan/panels/post-service/internal"
)

type postDatabaseRepo struct {
	db *pgxpool.Pool
}

func NewPostRepository(db *pgxpool.Pool) internal.PostDBRepository {
	return postDatabaseRepo{
		db: db,
	}
}

func (r postDatabaseRepo) CreatePost(ctx context.Context, panelId string, authorId string, data internal.PostCreate) (*internal.Post, error) {
	var id internal.PostId
	err := r.db.QueryRow(ctx, "INSERT INTO posts (panel_id, author_id, title, content) VALUES ($1, $2, $3, $4) RETURNING id", panelId, authorId, data.Title, data.Content).Scan(&id)
	if err != nil {
		if strings.Contains(err.Error(), "failed to connect to") {
			return nil, internal.WrapServiceError(err, internal.ConnectionErrorCode, "failed to connect to database")
		}
		log.Error().Err(err).Msg("unaccounted error whilst creating post")
		return nil, internal.WrapServiceError(err, internal.UnknownErrorCode, "failed to create post")
	}

	return r.GetPost(ctx, id)
}

func (r postDatabaseRepo) GetPost(ctx context.Context, id internal.PostId) (*internal.Post, error) {
	var post internal.Post
	row := r.db.QueryRow(ctx, "SELECT id, panel_id, author_id, title, content, created_at, updated_at FROM posts WHERE id=$1", id)
	err := row.Scan(&post.Id, &post.PanelId, &post.AuthorId, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt)
    if err != nil {
		if err == pgx.ErrNoRows {
			return nil, internal.WrapServiceError(err, internal.NotFoundErrorCode, "post not found")
		} else if strings.Contains(err.Error(), "failed to connect to") {
			return nil, internal.WrapServiceError(err, internal.ConnectionErrorCode, "failed to connect to database")
		}
		log.Error().Err(err).Msg("unaccounted error whilst getting post")
		return nil, internal.WrapServiceError(err, internal.UnknownErrorCode, "failed to get post")
	}

	return &post, nil
}

func (r postDatabaseRepo) GetPanelPost(ctx context.Context, id internal.PostId, panelId string) (*internal.Post, error) {
	var post internal.Post
	row := r.db.QueryRow(ctx, "SELECT id, panel_id, author_id, title, content, created_at, updated_at FROM posts WHERE id=$1 AND panel_id=$2", id, panelId)
	err := row.Scan(&post.Id, &post.PanelId, &post.AuthorId, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt)
    if err != nil {
		if err == pgx.ErrNoRows {
			return nil, internal.WrapServiceError(err, internal.NotFoundErrorCode, "post not found on that panel")
		} else if strings.Contains(err.Error(), "failed to connect to") {
			return nil, internal.WrapServiceError(err, internal.ConnectionErrorCode, "failed to connect to database")
		}
		log.Error().Err(err).Msg("unaccounted error whilst getting post")
		return nil, internal.WrapServiceError(err, internal.UnknownErrorCode, "failed to get post")
	}

	return &post, nil
}

func (r postDatabaseRepo) UpdatePost(ctx context.Context, id internal.PostId, data internal.PostUpdate) (*internal.Post, error) {
	// Transform request to patch data (marshal to remove omitted keys)
	patchData := goqu.Record{"updated_at": goqu.L("timezone('utc', now())")}
	marshalled, _ := json.Marshal(data)
	_ = json.Unmarshal(marshalled, &patchData)
	
	// Build a statement to updated the post
	statement, args, _ := goqu.Dialect("postgres").Update("posts").Prepared(true).Set(patchData).Where(goqu.C("id").Eq(id)).ToSQL()

	// Execute the query
	result, err := r.db.Exec(ctx, statement, args...)
	if err != nil {
		if strings.Contains(err.Error(), "failed to connect to") {
			return nil, internal.WrapServiceError(err, internal.ConnectionErrorCode, "failed to connect to database")
		}
		log.Error().Err(err).Msg("unaccounted error whilst updating post")
		return nil, internal.WrapServiceError(err, internal.UnknownErrorCode, "failed to update post")
	}

	// Check if any rows were affected by the query
	rows_affected := result.RowsAffected()
	if rows_affected != 1 {
		return nil, internal.NewServiceError(internal.NotFoundErrorCode, "post not found")
	}

	// Return the updated post
	return r.GetPost(ctx, id)
}

func (r postDatabaseRepo) DeletePost(ctx context.Context, id internal.PostId) error {
	result, err := r.db.Exec(ctx, "DELETE FROM posts WHERE id=$1", id)
	if err != nil {
		if strings.Contains(err.Error(), "failed to connect to") {
			return internal.WrapServiceError(err, internal.ConnectionErrorCode, "failed to connect to database")
		}
		log.Error().Err(err).Msg("error whilst deleting post")
		return internal.WrapServiceError(err, internal.UnknownErrorCode, "failed to delete post")
	}

	// Check if any rows were affected by the query
	rows_affected := result.RowsAffected()
	if rows_affected != 1 {
		return internal.NewServiceError(internal.NotFoundErrorCode, "post not found")
	}

	return nil
}

func (r postDatabaseRepo) GetFeedPosts(ctx context.Context) ([]*internal.Post, error) {
	// todo: pagination
	rows, err := r.db.Query(ctx, "SELECT id, panel_id, author_id, title, content, created_at, updated_at FROM posts ORDER BY created_at DESC LIMIT 25")
    if err != nil {
		if strings.Contains(err.Error(), "failed to connect to") {
			return nil, internal.WrapServiceError(err, internal.ConnectionErrorCode, "failed to connect to database")
		}
		log.Error().Err(err).Msg("unaccounted error whilst getting posts")
		return nil, internal.WrapServiceError(err, internal.UnknownErrorCode, "failed to get posts")
	}

	posts := []*internal.Post{}
	for rows.Next() {
		var post internal.Post
		err := rows.Scan(&post.Id, &post.PanelId, &post.AuthorId, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, internal.WrapServiceError(err, internal.UnknownErrorCode, "failed to scan posts")
		}
		posts = append(posts, &post)
	}

	if rows.Err() != nil {
		log.Error().Err(err).Msg("query error whilst retrieving posts")
		return nil, internal.WrapServiceError(err, internal.UnknownErrorCode, "failed to get posts")
	}

	return posts, nil
}

func (r postDatabaseRepo) GetUserPosts(ctx context.Context, userId string) ([]*internal.Post, error) {
	// todo: pagination
	rows, err := r.db.Query(ctx, "SELECT id, panel_id, author_id, title, content, created_at, updated_at FROM posts WHERE author_id=$1 ORDER BY created_at DESC LIMIT 25", userId)
    if err != nil {
		if strings.Contains(err.Error(), "failed to connect to") {
			return nil, internal.WrapServiceError(err, internal.ConnectionErrorCode, "failed to connect to database")
		}
		log.Error().Err(err).Msg("unaccounted error whilst getting posts by user")
		return nil, internal.WrapServiceError(err, internal.UnknownErrorCode, "failed to get posts")
	}

	posts := []*internal.Post{}
	for rows.Next() {
		var post internal.Post
		err := rows.Scan(&post.Id, &post.PanelId, &post.AuthorId, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, internal.WrapServiceError(err, internal.UnknownErrorCode, "failed to scan posts")
		}
		posts = append(posts, &post)
	}

	if rows.Err() != nil {
		log.Error().Err(err).Msg("query error whilst retrieving posts by user")
		return nil, internal.WrapServiceError(err, internal.UnknownErrorCode, "failed to get posts")
	}

	return posts, nil
}

func (r postDatabaseRepo) GetPanelPosts(ctx context.Context, panelId string) ([]*internal.Post, error) {
	// todo: pagination
	rows, err := r.db.Query(ctx, "SELECT id, panel_id, author_id, title, content, created_at, updated_at FROM posts WHERE panel_id=$1 ORDER BY created_at DESC LIMIT 25", panelId)
    if err != nil {
		if strings.Contains(err.Error(), "failed to connect to") {
			return nil, internal.WrapServiceError(err, internal.ConnectionErrorCode, "failed to connect to database")
		}
		log.Error().Err(err).Msg("unaccounted error whilst getting posts from panel")
		return nil, internal.WrapServiceError(err, internal.UnknownErrorCode, "failed to get posts")
	}

	posts := []*internal.Post{}
	for rows.Next() {
		var post internal.Post
		err := rows.Scan(&post.Id, &post.PanelId, &post.AuthorId, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, internal.WrapServiceError(err, internal.UnknownErrorCode, "failed to scan posts")
		}
		posts = append(posts, &post)
	}

	if rows.Err() != nil {
		log.Error().Err(err).Msg("query error whilst retrieving posts from panel")
		return nil, internal.WrapServiceError(err, internal.UnknownErrorCode, "failed to get posts")
	}

	return posts, nil
}

func (r postDatabaseRepo) DeletePostsByUser(ctx context.Context, userId string) ([]internal.PostId, error) {
	// Get post IDs for sending events
	deletedIds := []internal.PostId{}
	rows, err := r.db.Query(ctx, "SELECT id FROM posts WHERE author_id=$1", userId)
	if err == nil {
		for rows.Next() {
			var postId internal.PostId
			err = rows.Scan(&postId)
			if err == nil {
				deletedIds = append(deletedIds, postId)
			}
		}
	}

	// Delete posts
	_, err = r.db.Exec(ctx, "DELETE FROM posts WHERE author_id=$1", userId)
	if err != nil {
		if strings.Contains(err.Error(), "failed to connect to") {
			return nil, internal.WrapServiceError(err, internal.ConnectionErrorCode, "failed to connect to database")
		}
		log.Error().Err(err).Msg("error whilst deleting posts by user")
		return nil, internal.WrapServiceError(err, internal.UnknownErrorCode, "failed to delete posts")
	}

	return deletedIds, nil
}

func (r postDatabaseRepo) DeletePostsOnPanel(ctx context.Context, panelId string) ([]internal.PostId, error) {
	// Get post IDs for sending events
	deletedIds := []internal.PostId{}
	rows, err := r.db.Query(ctx, "SELECT id FROM posts WHERE panel_id=$1", panelId)
	if err == nil {
		for rows.Next() {
			var postId internal.PostId
			err = rows.Scan(&postId)
			if err == nil {
				deletedIds = append(deletedIds, postId)
			}
		}
	}

	// Delete posts
	_, err = r.db.Exec(ctx, "DELETE FROM posts WHERE panel_id=$1", panelId)
	if err != nil {
		if strings.Contains(err.Error(), "failed to connect to") {
			return nil, internal.WrapServiceError(err, internal.ConnectionErrorCode, "failed to connect to database")
		}
		log.Error().Err(err).Msg("error whilst deleting posts from panel")
		return nil, internal.WrapServiceError(err, internal.UnknownErrorCode, "failed to delete posts")
	}

	return deletedIds, nil
}