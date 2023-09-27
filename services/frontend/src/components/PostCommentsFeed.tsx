import { Center, Loader, Text } from '@mantine/core'

import CommentsFeed from './CommentsFeed'
import { useGetPostCommentsQuery } from '../app/api/comments'
import type { Post } from '../app/types/common'

function PostCommentsFeed({ post }: { post: Post }) {
  const { data, isLoading } = useGetPostCommentsQuery({ postId: post.id })
  if (isLoading) {
    return (
      <Center>
        <Loader color='dark' size='sm' />
      </Center>
    )
  } else if (!data) {
    return <Text color='red' align='center'>Failed to Load Comments</Text>
  } else if (!data.length) {
    return null
  }

  return <CommentsFeed comments={data} />
}

export default PostCommentsFeed