import { useState } from 'react'
import { useOutletContext, useNavigate } from 'react-router-dom'
import { useForm, hasLength } from '@mantine/form'
import { Stack, Paper, Text, TextInput, Textarea, Button } from '@mantine/core'

import { useAppSelector } from '../app/hooks'
import { useCreatePanelPostMutation } from '../app/api/posts'
import type { CreatePostData } from '../app/types/posts'
import type { PanelContext } from '../components/PanelLayout'

const NewPanelPostPage = () => {
  const { panel } = useOutletContext<PanelContext>()
  const [errorMsg, setErrorMsg] = useState('')
  const navigate = useNavigate()

  // Ensure the user is authenticated
  const currentUser = useAppSelector((state) => state.auth.currentUser)
  if (currentUser === null) {
    throw Error('You must be authenticated to create posts')
  }

  const createPostForm = useForm<CreatePostData>({
    initialValues: {
      title: '',
      content: '',
    },
    validate: {
      title: hasLength({ min: 3, max: 512 }, 'Title must be between 3 and 512 characters'),
      content: hasLength({ min: 3, max: 2048 }, 'Content must be between 3 and 2048 characters'),
    }
  })

  const [createPost, { isLoading }] = useCreatePanelPostMutation()
  const submitPost = async (values: CreatePostData) => {
    await createPost({
      panelId: panel.id,
      data: values
    }).unwrap().then((post) => {
      navigate(`/panel/${panel.name}/post/${post.id}`)
    }).catch((error) => {
      if (!error.data) {
        setErrorMsg('Failed to access the API')
      } else {
        setErrorMsg(error.data.msg)
      }
    })
  }

  return (
    <Paper shadow='md' radius='md' p='lg' withBorder>
      <form onSubmit={createPostForm.onSubmit(submitPost)}>
        <Stack spacing='md'>
          <TextInput 
            label='Title'
            placeholder='Post Title'
            {...createPostForm.getInputProps('title')}
          />

          <Textarea
            label='Content'
            placeholder='Post Content'
            {...createPostForm.getInputProps('content')}
          />
          
          { errorMsg && <Text color='red' align='center'>{'Error: ' + errorMsg}</Text> }

          <Button type='submit' variant='outline' color='teal' disabled={isLoading} fullWidth>
            Create Post
          </Button>
        </Stack>
      </form>
    </Paper>
  )
}

export default NewPanelPostPage