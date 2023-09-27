import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useForm, hasLength } from '@mantine/form'
import { Center, Container, Paper, Title, Text, TextInput, Textarea, Stack, Button } from '@mantine/core'

import { useAppSelector } from '../app/hooks'
import { useCreatePanelMutation } from '../app/api/panels'
import type { CreatePanelData } from '../app/types/panels'

const NewPanelPage = () => {
  const navigate = useNavigate()
  const [errorMsg, setErrorMsg] = useState('')

  // Ensure the user is authenticated
  const currentUser = useAppSelector((state) => state.auth.currentUser)
  if (currentUser === null) {
    throw Error('Authentication Required')
  }

  const panelForm = useForm<CreatePanelData>({
    initialValues: {
      name: '',
      description: '',
    },
    validate: {
      name: hasLength({ min: 3, max: 20 }, 'Panel name must be between 3 and 20 characters long'),
      description: hasLength({ min: 3, max: 512 }, 'Description must be between 3 and 512 characters'),
    }
  })

  const [createPanel, { isLoading }] = useCreatePanelMutation()
  const submitPanelForm = async (values: CreatePanelData) => {
    await createPanel({
      ...values
    }).unwrap().then((panel) => {
      navigate(`/panel/${panel.name}`)
    }).catch((error) => {
      if (!error.data) {
        setErrorMsg('Failed to access the API')
      } else {
        setErrorMsg(error.data.msg)
      }
    })
  }

  return (
    <Center h='95%'>
      <Container>
        <Title align='center' weight={900}>Create a Panel</Title>

        <Paper withBorder shadow='md' radius='md' p={30} mt={30}>
          <form onSubmit={panelForm.onSubmit(submitPanelForm)}>
            <Stack spacing='md'>
              <TextInput 
                label='Name'
                placeholder='e.g. music, programming, football'
                {...panelForm.getInputProps('name')}
              />

              <Textarea
                label='Description'
                placeholder='e.g. The place to talk about all things music related...'
                {...panelForm.getInputProps('description')}
              />

              { errorMsg && <Text color='red' align='center'>{'Error: ' + errorMsg}</Text> }

              <Button type='submit' variant='outline' color='teal' disabled={isLoading} fullWidth>Create Panel</Button>
            </Stack>
          </form>
        </Paper>
      </Container>
    </Center>
  )
}

export default NewPanelPage