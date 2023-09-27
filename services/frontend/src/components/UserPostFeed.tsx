import { Text } from '@mantine/core'

import FeedPost from './FeedPost'
import SkeletonPostFeed from './SkeletonPostFeed'
import { useGetUserPostsQuery } from '../app/api/posts'
import type { User } from '../app/types/common'

function UserPostFeed({ user }: { user: User }) {
  const { data, isLoading } = useGetUserPostsQuery({ userId: user.id })
  if (isLoading) {
    return <SkeletonPostFeed />
  } else if (!data) {
    return <Text align='center' color='red'>Failed to Load Posts</Text>
  } else if (!data.length) {
    // Check that there are posts.
    return <Text align='center'>No Posts Found!</Text>
  }

  return (
    <>
      {Object.values(data).map(post => {
        return <FeedPost key={post.id} post={post} hideAuthor={true} />
      })}
    </>
  )
}

export default UserPostFeed