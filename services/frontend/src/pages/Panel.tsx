import { useOutletContext } from 'react-router-dom'
import { Stack } from '@mantine/core'

import PanelPostFeed from '../components/PanelPostFeed'
import type { PanelContext } from '../components/PanelLayout'

function PanelPage() {
  const { panel } = useOutletContext<PanelContext>()

  return (
    <Stack my='lg' spacing='md'>
      <PanelPostFeed panel={panel} />
    </Stack>
  )
}

export default PanelPage