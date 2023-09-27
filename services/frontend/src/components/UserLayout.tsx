import { Suspense } from 'react'
import { Outlet, useLocation, useNavigate, useParams } from 'react-router-dom'
import { Container, Box, Center, Paper, Flex, Avatar, Text, Tabs } from '@mantine/core'
import { IconMessageCircle, IconAddressBook, IconSettings } from '@tabler/icons-react'

import LoadingBar from './LoadingBar'
import { User } from '../app/types/common'
import { useAppSelector } from '../app/hooks'
import { useGetUserByNameQuery } from '../app/api/users'
import type { ErrorResponse, QueryError } from '../app/types/api'

export type UserContext = {
  user: User
}

type UserPageParams = {
  username: string;
}

function UserLayout() {
  const navigate = useNavigate()
  const currentUser = useAppSelector((state) => state.auth.currentUser)
  
  const { pathname } = useLocation()
  let currentTab = 'posts'
  if (pathname.endsWith('about') || pathname.endsWith('about/')) {
    currentTab = 'about'
  } else if (pathname.endsWith('settings') || pathname.endsWith('settings/')) {
    currentTab = 'settings'
  }

  const { username } = useParams<UserPageParams>();
  if (username === undefined) {
    throw Error('Username not provided')
  }

  const { data, error, isLoading } = useGetUserByNameQuery({ username: username })
  if (isLoading) {
    return <LoadingBar />;
  } else if (!data) {
    if (!error) {
      throw Error('Unknown error occured')
    } else if ((error as QueryError).data) {
      const errResponse = (error as QueryError).data as ErrorResponse
      if (errResponse.msg) {
        throw Error(errResponse.msg)
      } else {
        throw Error('Unexpected API error occured')
      }
    } else {
      throw Error('Failed to access the API')
    }
  }

  return (
    <Container mt='xl'>
      <Tabs color='teal' radius='md' value={currentTab} onTabChange={(tab) => navigate(`/user/${data.username}${tab === 'posts' ? '' : `/${tab}`}`)}>
        <Paper shadow='md' radius='md' mt='md' withBorder>
          <Flex>
            <Avatar radius='md' size={200} color='lime' />
            <Paper w='100%'>
              <Box h='100%' pos='relative'>
                <Center h='100%'>
                  <Text weight={600} mr={3}>User:</Text>
                  <Text>{data.username}</Text>
                </Center>

                <Tabs.List pos='absolute' bottom={0} w='100%' grow>
                  <Tabs.Tab value='posts' icon={<IconMessageCircle size='0.8rem' />}>Posts</Tabs.Tab>
                  <Tabs.Tab value='about' icon={<IconAddressBook size='0.8rem' />}>About</Tabs.Tab>
                  {(currentUser && (currentUser.id == data.id || currentUser.isAdmin)) && <Tabs.Tab value='settings' icon={<IconSettings size='0.8rem' />}>Settings</Tabs.Tab>}
                </Tabs.List>
              </Box>
            </Paper>
          </Flex>
        </Paper>
      </Tabs>

      <Suspense>
        <Outlet context={{ user: data } satisfies UserContext} />
      </Suspense>
    </Container>
  )
}

export default UserLayout