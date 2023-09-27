import { useState } from 'react'
import { useNavigate, useOutletContext } from 'react-router-dom'
import { useForm, hasLength } from '@mantine/form'
import { Paper, Center, Stack, Group, Text, TextInput, Textarea, Button } from '@mantine/core'

import { useAppSelector } from '../app/hooks'
import { useDeletePanelByIdMutation, useUpdatePanelByIdMutation } from '../app/api/panels'
import type { Panel } from '../app/types/common'
import type { UpdatePanelData } from '../app/types/panels'
import type { PanelContext } from '../components/PanelLayout'

const UpdatePanelForm = ({
  panel,
  setPanel,
  setModifying,
  setErrorMsg
}: {
  panel: Panel,
  setPanel: React.Dispatch<Panel>,
  setModifying: React.Dispatch<boolean>,
  setErrorMsg: React.Dispatch<string>
}) => {
  const panelForm = useForm<UpdatePanelData>({
    initialValues: {
      name: panel.name,
      description: panel.description,
    },
    validate: {
      name: hasLength({ min: 3, max: 20 }, 'Name must be between 3 and 20 characters long'),
      description: hasLength({ min: 3, max: 512 }, 'Description must be between 3 and 512 characters'),
    }
  })

  const [updatePanel, { isLoading }] = useUpdatePanelByIdMutation()
  const submitUpdatePanel = async (values: UpdatePanelData) => {
    await updatePanel({
      id: panel.id,
      data: values
    }).unwrap().then((panelInfo) => {
      setErrorMsg('')
      setModifying(false)
      setPanel(panelInfo)
    }).catch((error) => {
      if (!error.data) {
        setErrorMsg('Failed to access the API')
      } else {
        setErrorMsg(error.data.msg)
      }
    })
  }

  return (
    <form onSubmit={panelForm.onSubmit(submitUpdatePanel)}>
      <Stack spacing='md'>
        <TextInput label='Name' {...panelForm.getInputProps('name')} />
        <Textarea label='Description' {...panelForm.getInputProps('description')} />

        <Button type='submit' variant='outline' color='teal' disabled={isLoading} fullWidth>Update</Button>
      </Stack>
    </form>
  )
}

function PanelSettingsPage() {
  const navigate = useNavigate()
  const [errorMsg, setErrorMsg] = useState<string>('')
  const [modifying, setModifying] = useState<boolean>(false)
  const { panel, setPanel } = useOutletContext<PanelContext>()

  // Check permissions
  const currentUser = useAppSelector((state) => state.auth.currentUser)
  if (!currentUser || !currentUser.isAdmin) {
    throw new Error('You do not have permission to view that page.')
  }

  // Panel Deletion
  const [deletePanel, { isLoading: isLoadingDelPanel }] = useDeletePanelByIdMutation()
  const submitDeletePanel = async () => {
    await deletePanel({ id: panel.id }).unwrap().then(() => {
      navigate('/')
    }).catch((error) => {
      if (!error.data) {
        setErrorMsg('Failed to access the API')
      } else {
        setErrorMsg(error.data.msg)
      }
    })
  }

  return (
    <Paper mt='md' radius='lg' shadow='md' p='lg' withBorder>
      <Center>
        <Group spacing='sm'>
          {modifying 
            ? <Button color='teal' onClick={() => { setModifying(false); setErrorMsg('') }}>Stop Modifying</Button>
            : <Button color='teal' onClick={() => setModifying(true)}>Modify Panel</Button>
          }
          <Button color='red' onClick={() => submitDeletePanel()} disabled={isLoadingDelPanel}>Delete Panel</Button>
        </Group>
      </Center>

      {modifying && <UpdatePanelForm panel={panel} setPanel={setPanel} setModifying={setModifying} setErrorMsg={setErrorMsg} />}

      {errorMsg && <Text color='red' mt='sm'>{'Error: ' + errorMsg}</Text>}
    </Paper>
  )
}

export default PanelSettingsPage