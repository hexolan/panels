import { useOutletContext } from 'react-router-dom'
import { Stack } from '@mantine/core'

import UserPostFeed from '../components/UserPostFeed'
import type { UserContext } from '../components/UserLayout'

function UserPage() {
  const { user } = useOutletContext<UserContext>()

  return (
    <Stack my='lg' spacing='md'>
      <UserPostFeed user={user} />
    </Stack>
  )
}

export default UserPage