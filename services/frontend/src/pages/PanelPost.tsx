import { useState } from 'react'
import { useParams, useOutletContext } from 'react-router-dom'
import { Box, Stack, Divider } from '@mantine/core'

import PagePost from '../components/PagePost'
import CommentsFeed from '../components/CommentsFeed'
import PostCommentsFeed from '../components/PostCommentsFeed'
import CreateComment from '../components/CreateComment'
import LoadingBar from '../components/LoadingBar'
import { useAppSelector } from '../app/hooks'
import { useGetPanelPostQuery } from '../app/api/posts'
import type { PanelContext } from '../components/PanelLayout'
import type { ErrorResponse } from '../app/types/api'
import type { Comment } from '../app/types/common'

type PanelPostPageParams = {
  panelName: string;
  postId: string;
}

function PanelPostPage() {
  const { panel } = useOutletContext<PanelContext>()
  const { postId } = useParams<PanelPostPageParams>();
  if (postId === undefined) { throw Error('post id not provided') }

  const [newComments, setNewComments] = useState<Comment[]>([])
  const addNewComment = (comment: Comment) => {
    setNewComments([comment].concat(newComments))
  }

  const currentUser = useAppSelector((state) => state.auth.currentUser)

  // Fetch the post
  const { data, error, isLoading } = useGetPanelPostQuery({ panelId: panel.id, id: postId })
  if (isLoading) {
    return <LoadingBar />;
  } else if (!data) {
    if (!error) {
      throw Error('Unknown error occured')
    } else if ('data' in error) {
      const errResponse = error.data as ErrorResponse
      if (errResponse.msg) {
        throw Error(errResponse.msg)
      } else {
        throw Error('Unexpected API error occured')
      }
    } else {
      throw Error('Failed to access the API')
    }
  }

  return (
    <Box mb='lg'>
      <PagePost post={data} />
      <Divider my='md' variant='none' />
      <Stack spacing='sm'>
        { currentUser && <CreateComment post={data} addNewComment={addNewComment} /> }
        { newComments.length > 0 && <CommentsFeed comments={newComments} /> }
        <PostCommentsFeed post={data} />
      </Stack>
    </Box>
  )
}

export default PanelPostPage