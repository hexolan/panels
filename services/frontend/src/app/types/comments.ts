import { convertRawTimestamp } from './api';

import type { Comment } from './common';
import type { RawResponse, RawTimestamp } from './api';

// Request Data
export type CreateCommentData = {
  message: string;
}

export type UpdateCommentData = Partial<CreateCommentData>

// API Request Paramaters
export type GetPostCommentsRequest = {
  postId: string;
}

type UpdateCommentRequest = {
  id: string;
  data: UpdateCommentData;
}

export type UpdatePostCommentRequest = UpdateCommentRequest & {
  postId: string;
}

type DeleteCommentRequest = {
  id: string;
}

export type DeletePostCommentRequest = DeleteCommentRequest & {
  postId: string;
}

export type CreatePostCommentRequest = {
  postId: string;
  data: CreateCommentData;
}

// API Responses
export type RawComment = {
  id: string;
  post_id: string;
  author_id: string;
  message: string;
  created_at: RawTimestamp;
  updated_at?: RawTimestamp;
}

export type RawCommentResponse = RawResponse & {
  data?: RawComment;
}

export type RawCommentsResponse = RawResponse & {
  data?: {
    comments: RawComment[];
  };
}

// API Response Conversion
export const convertRawComment = (rawComment: RawComment): Comment => ({
  id: rawComment.id,
  postId: rawComment.post_id,
  authorId: rawComment.author_id,
  message: rawComment.message,
  createdAt: convertRawTimestamp(rawComment.created_at),
  updatedAt: (rawComment.updated_at ? convertRawTimestamp(rawComment.updated_at) : undefined),
})