import { useState } from 'react'
import { Paper, Flex, Textarea, ActionIcon } from '@mantine/core'
import { useForm, hasLength } from '@mantine/form'
import { IconWriting } from '@tabler/icons-react'

import { useCreatePostCommentMutation } from '../app/api/comments'
import type { Comment, Post } from '../app/types/common'
import type { CreateCommentData } from '../app/types/comments'

const CreateComment = ({ post, addNewComment }: { post: Post, addNewComment: (comment: Comment) => void }) => {
  const [errorMsg, setErrorMsg] = useState('')

  const [createComment, { isLoading }] = useCreatePostCommentMutation()
  const submitComment = async (values: CreateCommentData) => {
    await createComment({
      postId: post.id,
      data: values
    }).unwrap().then((comment) => {
      // Display the new comment
      addNewComment(comment)
      setErrorMsg('')
    }).catch((error) => {
      if (!error.data) {
        setErrorMsg('Failed to access the API')
      } else {
        setErrorMsg(error.data.msg)
      }
    })
  }

  const commentForm = useForm<CreateCommentData>({
    initialValues: {
      message: '',
    },
    validate: {
      message: hasLength({ min: 3, max: 512 }, 'Message must be between 3 and 512 characters'),
    }
  })

  return (
    <Paper shadow='sm' radius='md' p='md' withBorder>
      <form onSubmit={commentForm.onSubmit(submitComment)}>
        <Flex gap='sm' align='center' direction='row' wrap='nowrap'>
          <Textarea 
            size='xs'
            w='100%'
            radius='lg'
            variant='filled'
            placeholder='Input comment...'
            error={errorMsg}
            {...commentForm.getInputProps('message')}
          />
          
          <ActionIcon type='submit' radius='lg' color='teal' variant='outline' size='xl' aria-label='Post Comment' disabled={isLoading}>
            <IconWriting />
          </ActionIcon>
        </Flex>
      </form>
    </Paper>
  )
}

export default CreateComment