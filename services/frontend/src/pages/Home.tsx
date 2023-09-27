import { Container, Title, Stack, Divider } from '@mantine/core'

import HomePostFeed from '../components/HomePostFeed'

const Homepage = () => {
  return (
    <Container mt='xl'>
      <Title>Feed</Title>
      <Divider my='md' variant='dotted' />
      <Stack my='lg' spacing='md'>
        <HomePostFeed />
      </Stack>
    </Container>
  )
}

export default Homepage