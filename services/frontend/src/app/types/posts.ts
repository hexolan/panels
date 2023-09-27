import { convertRawTimestamp } from './api'

import type { Post } from './common'
import type { RawResponse, RawTimestamp } from './api'

// Request Data
export type CreatePostData = {
  title: string;
  content: string;
}

export type UpdatePostData = Partial<CreatePostData>

// API Request Paramaters
export type GetPanelPostRequest = {
  id: string;
  panelId: string;
}

export type GetUserPostsRequest = {
  userId: string;
}

export type GetPanelPostsRequest = {
  panelId: string;
}

export type UpdatePostRequest = {
  id: string;
  data: UpdatePostData;
}

export type DeletePostRequest = {
  id: string;
}

export type CreatePostRequest = {
  panelId: string;
  data: CreatePostData;
}

// API Responses
export type RawPost = {
  id: string;
  panel_id: string;
  author_id: string;
  title: string;
  content: string;
  created_at: RawTimestamp;
  updated_at?: RawTimestamp;
}

export type RawPostResponse = RawResponse & {
  data?: RawPost;
}

export type RawPostsResponse = RawResponse & {
  data?: {
    posts: RawPost[];
  };
}

// API Response Conversion
export const convertRawPost = (rawPost: RawPost): Post => ({
  id: rawPost.id,
  panelId: rawPost.panel_id,
  authorId: rawPost.author_id,
  title: rawPost.title,
  content: rawPost.content,
  createdAt: convertRawTimestamp(rawPost.created_at),
  updatedAt: (rawPost.updated_at ? convertRawTimestamp(rawPost.updated_at) : undefined),
})