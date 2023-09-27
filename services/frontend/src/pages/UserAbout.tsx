import { useOutletContext } from 'react-router-dom'
import { Paper, Text } from '@mantine/core'

import type { UserContext } from '../components/UserLayout'

function UserAboutPage() {
  const { user } = useOutletContext<UserContext>()

  return (
    <Paper mt='md' radius='lg' shadow='md' p='lg' withBorder>
      <Text weight={500}>About {user.username}</Text>
      {user.createdAt && <Text>Signed up {new Date(user.createdAt).toUTCString()}</Text>}
    </Paper>
  )
}

export default UserAboutPage