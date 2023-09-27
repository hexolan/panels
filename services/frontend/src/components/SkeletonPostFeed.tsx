import SkeletonFeedPost from './SkeletonFeedPost'

const SkeletonPostFeed = () => [...Array(10)].map(() => <SkeletonFeedPost />)

export default SkeletonPostFeed