import { createElement } from 'react'

import { NavLink } from 'react-router-dom'
import { Navbar, ThemeIcon, Group, Text, UnstyledButton, rem } from '@mantine/core'
import { IconTrendingUp, IconSearch, IconMessages } from '@tabler/icons-react'

const NavbarButton = ({ text, page, icon }: { text: string, page: string, icon: JSX.ElementType }) => (
  <NavLink to={page} style={{ textDecoration: 'none' }}>
    {({ isActive }) => (
      <UnstyledButton
        p='xs'
        sx={(theme) => ({
          display: 'block',
          width: '100%',
          borderRadius: theme.radius.sm,

          backgroundColor: isActive ? theme.colors.gray[0] : 'inherit',
          '&:hover': {
            backgroundColor: theme.colors.gray[0],
          },
        })}
      >
        <Group>
          <ThemeIcon color='teal' variant='light'>
            { createElement(icon, { size: '1rem' }) }
          </ThemeIcon>

          <Text size='sm'>{text}</Text>
        </Group>
      </UnstyledButton>
    )}
  </NavLink>
)

function AppNavbar() {
  return (
    <Navbar width={{ base: 300 }} p='xs'>
      <Navbar.Section py='xs'>
        <Text size='xs' color='dimmed' my='xs' weight={500}>Browse</Text>
        <NavbarButton text='Feed' page='/' icon={IconTrendingUp} />
        <NavbarButton text='Find Panels' page='/panels' icon={IconSearch} />
      </Navbar.Section>

      <Navbar.Section 
        grow
        pt='xs'
        sx={(theme) => ({
          borderTop: `${rem(1)} solid ${theme.colors.gray[3]}`
        })}
      >
        <Text size='xs' color='dimmed' m='xs' weight={500}>Suggested Panels</Text>
        <NavbarButton text='panel/Panel' page='/panel/Panel' icon={IconMessages} />
      </Navbar.Section>
    </Navbar>
  )
}

export default AppNavbar