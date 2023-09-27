import { Link } from 'react-router-dom'
import { Header, Flex, Button, Group, Avatar, Text, Menu } from '@mantine/core'
import { IconChevronDown, IconUserEdit, IconLogout } from '@tabler/icons-react'

import panelsLogo from '../assets/logo.svg'
import { useAppSelector, useAppDispatch } from '../app/hooks'
import { setUnauthed } from '../app/features/auth'

function AppHeader() {
  const currentUser = useAppSelector((state) => state.auth.currentUser)
  const dispatch = useAppDispatch();

  const signoutUser = () => {
    dispatch(setUnauthed())
  }

  return (
    <Header height={60} p={20}>
      <Flex justify='space-between' align='center' h='100%'>
        <Link to='/'>
          <img src={panelsLogo} height={30} alt='Panels Logo' />
        </Link>
        {!currentUser ? (
          <Button color='teal' component={Link} to='/signin'>Sign In</Button>
        ) : (
          <Menu>
            <Menu.Target>
              <Button color='teal' variant='outline'>
                <Group spacing={7}>
                  <Avatar color='teal' radius='xl' size={25} />
                  <Text weight={500} size='sm' sx={{ lineHeight: 1 }} mr={3}>
                    {currentUser.username}
                  </Text>
                  <IconChevronDown size={20} />
                </Group>
              </Button>
            </Menu.Target>

            <Menu.Dropdown>
              <Menu.Label>User Actions</Menu.Label>
              <Menu.Item icon={<IconUserEdit />} component={Link} to={'/user/' + currentUser.username}>My Profile</Menu.Item>
              <Menu.Item color='red' icon={<IconLogout />} onClick={signoutUser}>Sign Out</Menu.Item>
            </Menu.Dropdown>
          </Menu>
        )}
      </Flex>
    </Header>
  )
}

export default AppHeader