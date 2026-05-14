import { Text } from '@mantine/core'

import FeedPost from './FeedPost'
import SkeletonPostFeed from './SkeletonPostFeed'
import { useGetFeedPostsQuery } from '../app/api/posts'

function HomePostFeed() {
  const { data, isLoading } = useGetFeedPostsQuery()
  if (isLoading) {
    return <SkeletonPostFeed />
  } else if (!data) {
    return <Text ta='center' color='red'>Failed to Load Posts</Text>
  } else if (!data.length) {
    // Check that there are posts.
    return <Text ta='center'>No Posts Found!</Text>
  }

  return (
    <>
      {Object.values(data).map(post => {
        return <FeedPost key={post.id} post={post} />
      })}
    </>
  )
}

export default HomePostFeed