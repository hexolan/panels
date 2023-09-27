import { Text } from '@mantine/core'

import FeedPost from './FeedPost'
import SkeletonPostFeed from './SkeletonPostFeed'
import { useGetPanelPostsQuery } from '../app/api/posts'
import type { Panel } from '../app/types/common'

function PanelPostFeed({ panel }: { panel: Panel }) {
  const { data, isLoading } = useGetPanelPostsQuery({ panelId: panel.id })
  if (isLoading) {
    return <SkeletonPostFeed />
  } else if (!data) {
    return <Text align='center'>Failed to Load Posts</Text>
  } else if (!data.length) {
    // Check that there are posts.
    return <Text align='center'>No Posts Found!</Text>
  }

  return (
    <>
      {Object.values(data).map(post => {
        return <FeedPost key={post.id} post={post} hidePanel={true} />
      })}
    </>
  )
}

export default PanelPostFeed