import { Link } from 'react-router-dom'
import { Paper, Skeleton, Box, Badge, Text, Group, ThemeIcon } from '@mantine/core'
import { IconUser, IconMessages } from '@tabler/icons-react'

import { useGetUserByIdQuery } from '../app/api/users'
import { useGetPanelByIdQuery } from '../app/api/panels'
import type { Post } from '../app/types/common'

const FeedPost = ({ post, hidePanel, hideAuthor }: { post: Post, hidePanel?: boolean, hideAuthor?: boolean }) => {
  // Fetch panel info
  let panelElement: React.ReactNode = null
  const { data: panelData, isLoading: panelIsLoading } = useGetPanelByIdQuery({ id: post.panelId })
  if (!hidePanel) {
    if (panelIsLoading) {
      panelElement = <Skeleton height={8} mt={6} width='20%' radius='xl' />
    } else if (!panelData) {
      panelElement = <Text color='red' size='xs'>Error Loading Panel Data</Text>
    } else {
      panelElement = (
        <Badge
        pl={0}
        color='orange'
        leftSection={
          <ThemeIcon color='orange' size={24} radius='xl' mr={5}>
              <IconMessages size={12} />
            </ThemeIcon>
          }
          component={Link}
          to={`/panel/${panelData.name}`}
        >
          {`panel/${panelData.name}`}
        </Badge>
      )
    }
  }
  
  // Fetch author info
  let authorElement: React.ReactNode = null
  const { data: authorData, isLoading: authorIsLoading } = useGetUserByIdQuery({ id: post.authorId })
  if (!hideAuthor) {
    if (authorIsLoading) {
      authorElement = <Skeleton height={8} mt={6} width='20%' radius='xl' />
    } else if (!authorData) {
      authorElement = <Text color='red' size='xs'>Error Loading Author Data</Text>
    } else {
      authorElement = (
        <Badge
          pl={0}
          color='teal'
          leftSection={
            <ThemeIcon color='teal' size={24} radius='xl' mr={5}>
              <IconUser size={12} />
            </ThemeIcon>
          }
          component={Link}
          to={`/user/${authorData.username}`}
        >
          {`user/${authorData.username}`}
        </Badge>
      )
    }
  }

  return (
    <Paper shadow='xl' radius='lg' p='lg' withBorder>
      <Group spacing='xs'>
        {panelElement}
        {authorElement}
      </Group>
      <Box component={Link} to={panelData ? `/panel/${panelData.name}/post/${post.id}` : '#'} style={{ textDecoration: 'none', color: 'inherit' }}>
        <Text mt={3} weight={600} lineClamp={1}>{post.title}</Text>
        <Text size='sm' lineClamp={2}>{post.content}</Text>
        <Text size='xs' color='dimmed' mt={3}>Click to View</Text>
      </Box>
    </Paper>
  )
}

export default FeedPost