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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: post.proto

package postv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PanelId   string                 `protobuf:"bytes,2,opt,name=panel_id,json=panelId,proto3" json:"panel_id,omitempty"`    // External Ref: Panel Id
	AuthorId  string                 `protobuf:"bytes,3,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"` // External Ref: User Id
	Title     string                 `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Content   string                 `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{0}
}

func (x *Post) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Post) GetPanelId() string {
	if x != nil {
		return x.PanelId
	}
	return ""
}

func (x *Post) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *Post) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Post) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Post) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Post) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type PostMutable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   *string `protobuf:"bytes,1,opt,name=title,proto3,oneof" json:"title,omitempty"`
	Content *string `protobuf:"bytes,2,opt,name=content,proto3,oneof" json:"content,omitempty"`
}

func (x *PostMutable) Reset() {
	*x = PostMutable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostMutable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostMutable) ProtoMessage() {}

func (x *PostMutable) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostMutable.ProtoReflect.Descriptor instead.
func (*PostMutable) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{1}
}

func (x *PostMutable) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *PostMutable) GetContent() string {
	if x != nil && x.Content != nil {
		return *x.Content
	}
	return ""
}

type CreatePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PanelId string       `protobuf:"bytes,1,opt,name=panel_id,json=panelId,proto3" json:"panel_id,omitempty"` // External Ref: Panel Id
	UserId  string       `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`    // External Ref: User Id
	Data    *PostMutable `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreatePostRequest) Reset() {
	*x = CreatePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostRequest) ProtoMessage() {}

func (x *CreatePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostRequest.ProtoReflect.Descriptor instead.
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePostRequest) GetPanelId() string {
	if x != nil {
		return x.PanelId
	}
	return ""
}

func (x *CreatePostRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreatePostRequest) GetData() *PostMutable {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPostRequest) Reset() {
	*x = GetPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostRequest) ProtoMessage() {}

func (x *GetPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostRequest.ProtoReflect.Descriptor instead.
func (*GetPostRequest) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{3}
}

func (x *GetPostRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetPanelPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PanelId string `protobuf:"bytes,1,opt,name=panel_id,json=panelId,proto3" json:"panel_id,omitempty"` // External Ref: Panel Id
	Id      string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPanelPostRequest) Reset() {
	*x = GetPanelPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPanelPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPanelPostRequest) ProtoMessage() {}

func (x *GetPanelPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPanelPostRequest.ProtoReflect.Descriptor instead.
func (*GetPanelPostRequest) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{4}
}

func (x *GetPanelPostRequest) GetPanelId() string {
	if x != nil {
		return x.PanelId
	}
	return ""
}

func (x *GetPanelPostRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdatePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Data *PostMutable `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UpdatePostRequest) Reset() {
	*x = UpdatePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePostRequest) ProtoMessage() {}

func (x *UpdatePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePostRequest.ProtoReflect.Descriptor instead.
func (*UpdatePostRequest) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{5}
}

func (x *UpdatePostRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdatePostRequest) GetData() *PostMutable {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeletePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePostRequest) Reset() {
	*x = DeletePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostRequest) ProtoMessage() {}

func (x *DeletePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePostRequest.ProtoReflect.Descriptor instead.
func (*DeletePostRequest) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{6}
}

func (x *DeletePostRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetFeedPostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFeedPostsRequest) Reset() {
	*x = GetFeedPostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeedPostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeedPostsRequest) ProtoMessage() {}

func (x *GetFeedPostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeedPostsRequest.ProtoReflect.Descriptor instead.
func (*GetFeedPostsRequest) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{7}
}

type FeedPosts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*Post `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *FeedPosts) Reset() {
	*x = FeedPosts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedPosts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedPosts) ProtoMessage() {}

func (x *FeedPosts) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedPosts.ProtoReflect.Descriptor instead.
func (*FeedPosts) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{8}
}

func (x *FeedPosts) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

type GetUserPostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // External Ref: User Id
}

func (x *GetUserPostsRequest) Reset() {
	*x = GetUserPostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserPostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserPostsRequest) ProtoMessage() {}

func (x *GetUserPostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserPostsRequest.ProtoReflect.Descriptor instead.
func (*GetUserPostsRequest) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{9}
}

func (x *GetUserPostsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserPosts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*Post `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *UserPosts) Reset() {
	*x = UserPosts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPosts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPosts) ProtoMessage() {}

func (x *UserPosts) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPosts.ProtoReflect.Descriptor instead.
func (*UserPosts) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{10}
}

func (x *UserPosts) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

type GetPanelPostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PanelId string `protobuf:"bytes,1,opt,name=panel_id,json=panelId,proto3" json:"panel_id,omitempty"` // External Ref: Panel Id
}

func (x *GetPanelPostsRequest) Reset() {
	*x = GetPanelPostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPanelPostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPanelPostsRequest) ProtoMessage() {}

func (x *GetPanelPostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPanelPostsRequest.ProtoReflect.Descriptor instead.
func (*GetPanelPostsRequest) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{11}
}

func (x *GetPanelPostsRequest) GetPanelId() string {
	if x != nil {
		return x.PanelId
	}
	return ""
}

type PanelPosts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*Post `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *PanelPosts) Reset() {
	*x = PanelPosts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PanelPosts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PanelPosts) ProtoMessage() {}

func (x *PanelPosts) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PanelPosts.ProtoReflect.Descriptor instead.
func (*PanelPosts) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{12}
}

func (x *PanelPosts) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

// Kafka Event Schema
type PostEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Data *Post  `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PostEvent) Reset() {
	*x = PostEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostEvent) ProtoMessage() {}

func (x *PostEvent) ProtoReflect() protoreflect.Message {
	mi := &file_post_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostEvent.ProtoReflect.Descriptor instead.
func (*PostEvent) Descriptor() ([]byte, []int) {
	return file_post_proto_rawDescGZIP(), []int{13}
}

func (x *PostEvent) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PostEvent) GetData() *Post {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_post_proto protoreflect.FileDescriptor

var file_post_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x61,
	0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4, 0x01, 0x0a, 0x04, 0x50,
	0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x5d, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x4d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x19, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x22, 0x78, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x73,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4d, 0x75, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x40, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x54,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x46, 0x65, 0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x37, 0x0a, 0x09, 0x46, 0x65, 0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x2a, 0x0a,
	0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70,
	0x61, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x2e, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x37, 0x0a, 0x09, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x2a, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x73,
	0x74, 0x73, 0x22, 0x31, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x50, 0x6f,
	0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61,
	0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61,
	0x6e, 0x65, 0x6c, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x0a, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x50, 0x6f,
	0x73, 0x74, 0x73, 0x12, 0x2a, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x22,
	0x49, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xf3, 0x04, 0x0a, 0x0b, 0x50,
	0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c,
	0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x61,
	0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73,
	0x74, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x1e,
	0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x61, 0x6e,
	0x65, 0x6c, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x23, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x73, 0x2e,
	0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x6e, 0x65, 0x6c,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x61,
	0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73,
	0x74, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73,
	0x74, 0x12, 0x21, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x70, 0x61, 0x6e,
	0x65, 0x6c, 0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x46, 0x65,
	0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x23, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x73,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64,
	0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70,
	0x61, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65,
	0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x23, 0x2e, 0x70, 0x61, 0x6e, 0x65,
	0x6c, 0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x24, 0x2e, 0x70,
	0x61, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x6e, 0x65, 0x6c, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x00,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_post_proto_rawDescOnce sync.Once
	file_post_proto_rawDescData = file_post_proto_rawDesc
)

func file_post_proto_rawDescGZIP() []byte {
	file_post_proto_rawDescOnce.Do(func() {
		file_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_post_proto_rawDescData)
	})
	return file_post_proto_rawDescData
}

var file_post_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_post_proto_goTypes = []interface{}{
	(*Post)(nil),                  // 0: panels.post.v1.Post
	(*PostMutable)(nil),           // 1: panels.post.v1.PostMutable
	(*CreatePostRequest)(nil),     // 2: panels.post.v1.CreatePostRequest
	(*GetPostRequest)(nil),        // 3: panels.post.v1.GetPostRequest
	(*GetPanelPostRequest)(nil),   // 4: panels.post.v1.GetPanelPostRequest
	(*UpdatePostRequest)(nil),     // 5: panels.post.v1.UpdatePostRequest
	(*DeletePostRequest)(nil),     // 6: panels.post.v1.DeletePostRequest
	(*GetFeedPostsRequest)(nil),   // 7: panels.post.v1.GetFeedPostsRequest
	(*FeedPosts)(nil),             // 8: panels.post.v1.FeedPosts
	(*GetUserPostsRequest)(nil),   // 9: panels.post.v1.GetUserPostsRequest
	(*UserPosts)(nil),             // 10: panels.post.v1.UserPosts
	(*GetPanelPostsRequest)(nil),  // 11: panels.post.v1.GetPanelPostsRequest
	(*PanelPosts)(nil),            // 12: panels.post.v1.PanelPosts
	(*PostEvent)(nil),             // 13: panels.post.v1.PostEvent
	(*timestamppb.Timestamp)(nil), // 14: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 15: google.protobuf.Empty
}
var file_post_proto_depIdxs = []int32{
	14, // 0: panels.post.v1.Post.created_at:type_name -> google.protobuf.Timestamp
	14, // 1: panels.post.v1.Post.updated_at:type_name -> google.protobuf.Timestamp
	1,  // 2: panels.post.v1.CreatePostRequest.data:type_name -> panels.post.v1.PostMutable
	1,  // 3: panels.post.v1.UpdatePostRequest.data:type_name -> panels.post.v1.PostMutable
	0,  // 4: panels.post.v1.FeedPosts.posts:type_name -> panels.post.v1.Post
	0,  // 5: panels.post.v1.UserPosts.posts:type_name -> panels.post.v1.Post
	0,  // 6: panels.post.v1.PanelPosts.posts:type_name -> panels.post.v1.Post
	0,  // 7: panels.post.v1.PostEvent.data:type_name -> panels.post.v1.Post
	2,  // 8: panels.post.v1.PostService.CreatePost:input_type -> panels.post.v1.CreatePostRequest
	3,  // 9: panels.post.v1.PostService.GetPost:input_type -> panels.post.v1.GetPostRequest
	4,  // 10: panels.post.v1.PostService.GetPanelPost:input_type -> panels.post.v1.GetPanelPostRequest
	5,  // 11: panels.post.v1.PostService.UpdatePost:input_type -> panels.post.v1.UpdatePostRequest
	6,  // 12: panels.post.v1.PostService.DeletePost:input_type -> panels.post.v1.DeletePostRequest
	7,  // 13: panels.post.v1.PostService.GetFeedPosts:input_type -> panels.post.v1.GetFeedPostsRequest
	9,  // 14: panels.post.v1.PostService.GetUserPosts:input_type -> panels.post.v1.GetUserPostsRequest
	11, // 15: panels.post.v1.PostService.GetPanelPosts:input_type -> panels.post.v1.GetPanelPostsRequest
	0,  // 16: panels.post.v1.PostService.CreatePost:output_type -> panels.post.v1.Post
	0,  // 17: panels.post.v1.PostService.GetPost:output_type -> panels.post.v1.Post
	0,  // 18: panels.post.v1.PostService.GetPanelPost:output_type -> panels.post.v1.Post
	0,  // 19: panels.post.v1.PostService.UpdatePost:output_type -> panels.post.v1.Post
	15, // 20: panels.post.v1.PostService.DeletePost:output_type -> google.protobuf.Empty
	8,  // 21: panels.post.v1.PostService.GetFeedPosts:output_type -> panels.post.v1.FeedPosts
	10, // 22: panels.post.v1.PostService.GetUserPosts:output_type -> panels.post.v1.UserPosts
	12, // 23: panels.post.v1.PostService.GetPanelPosts:output_type -> panels.post.v1.PanelPosts
	16, // [16:24] is the sub-list for method output_type
	8,  // [8:16] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_post_proto_init() }
func file_post_proto_init() {
	if File_post_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostMutable); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePostRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPanelPostRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePostRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePostRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFeedPostsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedPosts); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserPostsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPosts); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPanelPostsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PanelPosts); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_post_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_post_proto_goTypes,
		DependencyIndexes: file_post_proto_depIdxs,
		MessageInfos:      file_post_proto_msgTypes,
	}.Build()
	File_post_proto = out.File
	file_post_proto_rawDesc = nil
	file_post_proto_goTypes = nil
	file_post_proto_depIdxs = nil
}
