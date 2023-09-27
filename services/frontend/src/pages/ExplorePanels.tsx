import { Link } from 'react-router-dom'
import { IconMessages, IconTableOff } from '@tabler/icons-react'
import { Container, Stack, Paper, Box, Title, Text, Anchor, Divider, ThemeIcon, Group } from '@mantine/core'

import { useAppSelector } from '../app/hooks'

const ExplorePanelsPage = () => {
  const currentUser = useAppSelector((state) => state.auth.currentUser)

  return (
    <Container mt='xl'>
      <Title>Explore Panels</Title>
      {currentUser && (
        <Text color='dimmed' size='sm' mt={5}>
          Alternatively you could <Anchor size='sm' component={Link} to='/panels/new'>create your own.</Anchor>
        </Text>
      )}
      <Divider my='md' variant='dotted' />

      <Stack spacing='sm' align='stretch'>
        <Paper shadow='xl' radius='md' p='md' withBorder component={Link} to='/panel/Panel'>
          <Group>
            <ThemeIcon color='teal' variant='light' size='xl'><IconMessages /></ThemeIcon>
            <Box>
              <Text weight={600}>Panel</Text>
              <Text>The first and therefore defacto primary panel.</Text>
              <Text color='dimmed' size='xs' mt={3}>Click to View</Text>
            </Box>
          </Group>
        </Paper>

        <Paper shadow='xl' radius='md' p='md' withBorder component='a' href='https://github.com/hexolan/Panels/'>
          <Group>
            <ThemeIcon color='red' variant='light' size='xl'><IconTableOff /></ThemeIcon>
            <Box>
              <Text weight={600}>Note</Text>
              <Text>This page is exemplary as this feature is currently unimplemented.</Text>
              <Text color='dimmed' size='xs' mt={3}>Planned Functionality</Text>
            </Box>
          </Group>
        </Paper>
      </Stack>
    </Container>
  )
}

export default ExplorePanelsPage