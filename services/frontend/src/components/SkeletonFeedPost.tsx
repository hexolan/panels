import { Paper, Skeleton, Box, Stack, Group } from '@mantine/core'

const SkeletonFeedPost = () => (
  <Paper shadow='xl' radius='lg' p='lg' withBorder>
    <Group spacing='xs'>
      <Skeleton height={8} mt={6} width='20%' radius='xl' />
      <Skeleton height={8} mt={6} width='20%' radius='xl' />
    </Group>
    <Stack mt={2} spacing={1}>
      <Box>
        <Skeleton h='md' radius='xl' w='60%' />
        <Skeleton h='md' radius='xl' mt='sm' />
        <Skeleton h='xs' radius='xl' mt='sm' w='20%' />
      </Box>
    </Stack>
  </Paper>
)

export default SkeletonFeedPost