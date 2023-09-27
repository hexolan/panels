package service

import (
	"context"
	"strings"

	"github.com/hexolan/panels/post-service/internal"
	"github.com/hexolan/panels/post-service/internal/kafka/producer"
)

type postService struct {
	events producer.PostEventProducer

	repo internal.PostRepository
}

func NewPostService(events producer.PostEventProducer, repo internal.PostRepository) internal.PostService {
	return postService{
		events: events,
		repo: repo,
	}
}

func (srv postService) CreatePost(ctx context.Context, panelId string, authorId string, data internal.PostCreate) (*internal.Post, error) {
	// Validate the data
	err := data.Validate()
	if err != nil {
		return nil, internal.NewServiceErrorf(internal.InvalidArgumentErrorCode, "invalid argument: %s", err.Error())
	}

	// Create the post
	post, err := srv.repo.CreatePost(ctx, panelId, authorId, data)
	
	// Dispatch post created event
	if err == nil {
		srv.events.DispatchCreatedEvent(post)
	}

	return post, err
}

func (srv postService) GetPost(ctx context.Context, id internal.PostId) (*internal.Post, error) {
	return srv.repo.GetPost(ctx, id)
}

func (srv postService) GetPanelPost(ctx context.Context, id internal.PostId, panelId string) (*internal.Post, error) {
	panelId = strings.ToLower(panelId)  // Panel IDs are case insensitive
	return srv.repo.GetPanelPost(ctx, id, panelId)
}

func (srv postService) UpdatePost(ctx context.Context, id internal.PostId, data internal.PostUpdate) (*internal.Post, error) {
	// Validate the data
	if data == (internal.PostUpdate{}) {
		return nil, internal.NewServiceError(internal.InvalidArgumentErrorCode, "no values provided")
	}

	err := data.Validate()
	if err != nil {
		return nil, internal.NewServiceErrorf(internal.InvalidArgumentErrorCode, "invalid argument: %s", err.Error())
	}

	// Update the post
	post, err := srv.repo.UpdatePost(ctx, id, data)

	// Dispatch post created event
	if err == nil {
		srv.events.DispatchUpdatedEvent(post)
	}

	return post, err
}

func (srv postService) DeletePost(ctx context.Context, id internal.PostId) error {
	err := srv.repo.DeletePost(ctx, id)
	
	// Dispatch post deleted event
	if err == nil {
		srv.events.DispatchDeletedEvent(id)
	}

	return err
}

func (srv postService) GetFeedPosts(ctx context.Context) ([]*internal.Post, error) {
	return srv.repo.GetFeedPosts(ctx)
}

func (srv postService) GetUserPosts(ctx context.Context, userId string) ([]*internal.Post, error) {
	return srv.repo.GetUserPosts(ctx, userId)
}

func (srv postService) GetPanelPosts(ctx context.Context, panelId string) ([]*internal.Post, error) {
	return srv.repo.GetPanelPosts(ctx, panelId)
}