import { apiSlice } from '../features/api'
import { convertRawComment } from '../types/comments'

import type { Comment } from '../types/common'
import type {
  RawComment, RawCommentResponse, RawCommentsResponse,
  GetPostCommentsRequest,
  UpdatePostCommentRequest,
  DeletePostCommentRequest,
  CreatePostCommentRequest
} from '../types/comments'

export const commentsApiSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    getPostComments: builder.query<Comment[], GetPostCommentsRequest>({
      query: data => ({ url: `/v1/posts/${data.postId}/comments` }),
      transformResponse: (response: RawCommentsResponse) => {
        if (response.data === undefined) {
          throw Error('invalid comments response')
        } else if (!response.data.comments) {
          return []
        }

        return response.data.comments.map<Comment>((rawComment: RawComment) => convertRawComment(rawComment))
      }
    }),

    updatePostComment: builder.mutation<Comment, UpdatePostCommentRequest>({
      query: req => ({
        url: `/v1/posts/${req.postId}/comments/${req.id}`,
        method: 'PATCH',
        body: { ...req.data }
      }),
      transformResponse: (response: RawCommentResponse) => {
        if (response.data === undefined) { throw Error('invalid comment response') }

        return convertRawComment(response.data)
      }
    }),

    deletePostComment: builder.mutation<void, DeletePostCommentRequest>({
      query: req => ({
        url: `/v1/posts/${req.postId}/comments/${req.id}`,
        method: 'DELETE'
      })
    }),

    createPostComment: builder.mutation<Comment, CreatePostCommentRequest>({
      query: req => ({
        url: `/v1/posts/${req.postId}/comments`,
        method: 'POST',
        body: { ...req.data }
      }),
      transformResponse: (response: RawCommentResponse) => {
        if (response.data === undefined) { throw Error('invalid comment response') }

        return convertRawComment(response.data)
      }
    }),
  })
})

export const {
  useGetPostCommentsQuery,
  useUpdatePostCommentMutation,
  useDeletePostCommentMutation,
  useCreatePostCommentMutation
} = commentsApiSlice