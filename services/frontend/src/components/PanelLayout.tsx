import { Suspense, useState } from 'react'
import { Link, Outlet, useParams } from 'react-router-dom'
import { Paper, Container, Group, Box, Text, Button, rem } from '@mantine/core'

import LoadingBar from '../components/LoadingBar'
import { useAppSelector } from '../app/hooks'
import { useGetPanelByNameQuery } from '../app/api/panels'
import type { Panel } from '../app/types/common'
import type { ErrorResponse } from '../app/types/api'

export type PanelContext = {
  panel: Panel;
  setPanel: React.Dispatch<Panel>;
}

type PanelParams = {
  panelName: string
}

const PanelLayoutComponent = ({ panel, setPanel }: { panel: Panel, setPanel: React.Dispatch<Panel> }) => {
  const currentUser = useAppSelector((state) => state.auth.currentUser)
  
  return (
    <>
      <Paper py={rem(50)} shadow='md' sx={{ borderBottom: '1px' }}>
        <Container>
          <Group position='apart'>
            <Box component={Link} to={`/panel/${panel.name}`} style={{ textDecoration: 'none' }}>
              <Text size='lg' color='black'>{panel.name}</Text>
              <Text size='sm' color='dimmed'>{panel.description}</Text>
            </Box>

            <Group spacing='sm'>
              {currentUser && <Button size='xs' variant='filled' color='teal' component={Link} to={`/panel/${panel.name}/posts/new`}>Create Post</Button>}
              {currentUser && currentUser.isAdmin && <Button size='xs' variant='outline' color='green' component={Link} to={`/panel/${panel.name}/settings`}>Manage Panel</Button>}
            </Group>
          </Group>
        </Container>
      </Paper>
      <Container mt='xl'>
        <Suspense>
          <Outlet context={{ panel: panel, setPanel: setPanel } satisfies PanelContext} />
        </Suspense>
      </Container>
    </>
  )
}

const PanelLayoutComponentWrapper = ({ panel: initialPanel }: { panel: Panel }) => {
  const [panel, setPanel] = useState<Panel>(initialPanel)
  return <PanelLayoutComponent panel={panel} setPanel={setPanel} />
}

function PanelLayout() {
  const { panelName } = useParams<PanelParams>();
  if (panelName === undefined) {
    throw Error('panel name not provided')
  }
  
  const { data, error, isLoading } = useGetPanelByNameQuery({ name: panelName })
  if (isLoading) {
    return <LoadingBar />;
  } else if (!data) {
    if (!error) {
      throw Error('Unknown error occured')
    } else if ('data' in error) {
      const errResponse = error.data as ErrorResponse
      if (errResponse.msg) {
        throw Error(errResponse.msg)
      } else {
        throw Error('Unexpected API error occured')
      }
    } else {
      throw Error('Failed to access the API')
    }
  }
  
  return <PanelLayoutComponentWrapper panel={data} />
}

export default PanelLayout