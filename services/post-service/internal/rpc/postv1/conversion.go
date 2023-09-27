// Handles conversion between Protobuf types and service types
package postv1

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	"github.com/hexolan/panels/post-service/internal"
)

// Post -> Protobuf 'Post'
func PostToProto(post *internal.Post) *Post {
	proto := Post{
		Id: post.Id.GetReprId(),

		PanelId: post.PanelId,
		AuthorId: post.AuthorId,

		Title: post.Title,
		Content: post.Content,

		CreatedAt: timestamppb.New(post.CreatedAt.Time),
	}

	// convert nullable attributes to PB form (if present)
	if post.UpdatedAt.Valid == true {
		proto.UpdatedAt = timestamppb.New(post.UpdatedAt.Time)
	}

	return &proto
}

// []Post -> []Protobuf 'Post'
func PostsToProto(posts []*internal.Post) []*Post {
	protoPosts := []*Post{}
	for _, post := range posts {
		protoPosts = append(protoPosts, PostToProto(post))
	}
	return protoPosts
} 

// Protobuf 'PostMutable' -> PostCreate
func PostCreateFromProto(proto *PostMutable) internal.PostCreate {
	return internal.PostCreate{
		Title: proto.GetTitle(),
		Content: proto.GetContent(),
	}
}

// Protobuf 'PostMutable' -> PostUpdate
func PostUpdateFromProto(proto *PostMutable) internal.PostUpdate {
	return internal.PostUpdate{
		Title: proto.Title,
		Content: proto.Content,
	}
}