import { Stack } from '@mantine/core'

import FeedComment from './FeedComment'
import type { Comment } from '../app/types/common'

function CommentsFeed({ comments }: { comments: Comment[] }) {
  return (
    <Stack spacing='sm'>
      {Object.values(comments).map(comment => <FeedComment key={comment.id} comment={comment} />)}
    </Stack>
  )
}

export default CommentsFeed