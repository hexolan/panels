package rpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	"github.com/hexolan/panels/post-service/internal"
	pb "github.com/hexolan/panels/post-service/internal/rpc/postv1"
)

type postServer struct {
	pb.UnimplementedPostServiceServer
	
	service internal.PostService
}

func NewPostServer(service internal.PostService) postServer {
	return postServer{service: service}
}

func (svr *postServer) CreatePost(ctx context.Context, request *pb.CreatePostRequest) (*pb.Post, error) {
	// Ensure the required args are provided
	if request.GetData() == nil {
		return nil, status.Error(codes.InvalidArgument, "malformed request")
	}

	if request.GetPanelId() == "" {
		return nil, status.Error(codes.InvalidArgument, "panel id not provided")
	}

	if request.GetUserId() == "" {
		return nil, status.Error(codes.InvalidArgument, "user id not provided")
	}

	// Convert to service model
	data := pb.PostCreateFromProto(request.GetData())
	
	// Pass to service method for creation
	post, err := svr.service.CreatePost(ctx, request.GetPanelId(), request.GetUserId(), data)
	if err != nil {
		return nil, err
	}

	return pb.PostToProto(post), nil
}

func (svr *postServer) GetPost(ctx context.Context, request *pb.GetPostRequest) (*pb.Post, error) {
	// Ensure the required args are provided
	if request.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "post id not provided")
	}

	// Convert to business model
	id, err := internal.NewPostIdFromRepr(request.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid post id")
	}

	// Pass to service method for retrieval
	post, err := svr.service.GetPost(ctx, *id)
	if err != nil {
		return nil, err
	}
	return pb.PostToProto(post), nil
}

func (svr *postServer) GetPanelPost(ctx context.Context, request *pb.GetPanelPostRequest) (*pb.Post, error) {
	// Ensure the required args are provided
	if request.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "post id not provided")
	}

	if request.GetPanelId() == "" {
		return nil, status.Error(codes.InvalidArgument, "panel id not provided")
	}

	// Convert to service model
	id, err := internal.NewPostIdFromRepr(request.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid post id")
	}

	// Pass to service method for retrieval
	post, err := svr.service.GetPanelPost(ctx, *id, request.GetPanelId())
	if err != nil {
		return nil, err
	}
	return pb.PostToProto(post), nil
}

func (svr *postServer) UpdatePost(ctx context.Context, request *pb.UpdatePostRequest) (*pb.Post, error) {
	// Ensure the required args are provided
	if request.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "post id not provided")
	}

	if request.GetData() == nil {
		return nil, status.Error(codes.InvalidArgument, "malformed request")
	}

	// Convert to service models
	id, err := internal.NewPostIdFromRepr(request.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid post id")
	}

	data := pb.PostUpdateFromProto(request.GetData())

	// Pass to service method for update
	post, err := svr.service.UpdatePost(ctx, *id, data)
	if err != nil {
		return nil, err
	}

	return pb.PostToProto(post), nil
}

func (svr *postServer) DeletePost(ctx context.Context, request *pb.DeletePostRequest) (*emptypb.Empty, error) {
	// Ensure the required args are provided
	if request.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "post id not provided")
	}

	// Convert to service model
	id, err := internal.NewPostIdFromRepr(request.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid post id")
	}

	// Pass to service method for deletion
	err = svr.service.DeletePost(ctx, *id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (svr *postServer) GetFeedPosts(ctx context.Context, request *pb.GetFeedPostsRequest) (*pb.FeedPosts, error) {
	// Get the posts
	posts, err := svr.service.GetFeedPosts(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.FeedPosts{Posts: pb.PostsToProto(posts)}, nil
}

func (svr *postServer) GetUserPosts(ctx context.Context, request *pb.GetUserPostsRequest) (*pb.UserPosts, error) {
	// Ensure the required args are provided
	if request.GetUserId() == "" {
		return nil, status.Error(codes.InvalidArgument, "user id not provided")
	}

	// Get the posts
	posts, err := svr.service.GetUserPosts(ctx, request.GetUserId())
	if err != nil {
		return nil, err
	}

	return &pb.UserPosts{Posts: pb.PostsToProto(posts)}, nil
}

func (svr *postServer) GetPanelPosts(ctx context.Context, request *pb.GetPanelPostsRequest) (*pb.PanelPosts, error) {
	// Ensure the required args are provided
	if request.GetPanelId() == "" {
		return nil, status.Error(codes.InvalidArgument, "panel id not provided")
	}

	// Get the posts
	posts, err := svr.service.GetPanelPosts(ctx, request.GetPanelId())
	if err != nil {
		return nil, err
	}

	return &pb.PanelPosts{Posts: pb.PostsToProto(posts)}, nil
}