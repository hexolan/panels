import { useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { Paper, Stack, Badge, ThemeIcon, Text, Group, Menu, ActionIcon, TextInput, Textarea, Button } from '@mantine/core'
import { useForm, hasLength } from '@mantine/form'
import { IconUser, IconMenu2, IconPencil, IconTrash } from '@tabler/icons-react'

import { useAppSelector } from '../app/hooks'
import { useGetUserByIdQuery } from '../app/api/users'
import { useDeletePostMutation, useUpdatePostMutation } from '../app/api/posts'
import type { Post } from '../app/types/common'
import type { UpdatePostData } from '../app/types/posts'

const ModifyPostForm = ({
  post,
  setPost,
  setModifying
}: {
  post: Post,
  setPost: React.Dispatch<Post>,
  setModifying: React.Dispatch<boolean>
}) => {
  const [errorMsg, setErrorMsg] = useState('')
  const updatePostForm = useForm<UpdatePostData>({
    initialValues: {
      title: post.title,
      content: post.content,
    },
    validate: {
      title: hasLength({ min: 3, max: 512 }, 'Title must be between 3 and 512 characters'),
      content: hasLength({ min: 3, max: 2048 }, 'Content must be between 3 and 2048 characters'),
    }
  })

  const [updatePost, { isLoading }] = useUpdatePostMutation()
  const submitUpdatePost = async (values: UpdatePostData) => {
    await updatePost({
      id: post.id,
      data: values
    }).unwrap().then((postInfo) => {
      setErrorMsg('')
      setPost(postInfo)
      setModifying(false)
    }).catch((error) => {
      if (!error.data) {
        setErrorMsg('Failed to access the API')
      } else {
        setErrorMsg(error.data.msg)
      }
    })
  }

  return (
    <form onSubmit={updatePostForm.onSubmit(submitUpdatePost)}>
      <Stack spacing='sm'>
        <TextInput
          label='Title'
          placeholder='Post Title'
          {...updatePostForm.getInputProps('title')}
        />

        <Textarea
          label='Content'
          placeholder='Post Content'
          {...updatePostForm.getInputProps('content')}
        />

        {errorMsg && <Text color='red' align='center'>{'Error: ' + errorMsg}</Text>}

        <Button type='submit' variant='outline' color='teal' disabled={isLoading} fullWidth>
          Update Post
        </Button>
      </Stack>
    </form>
  )
}

const PagePostItem = ({ post, setPost }: { post: Post, setPost: React.Dispatch<Post> }) => {
  const navigate = useNavigate()

  const [modifying, setModifying] = useState<boolean>(false)
  const currentUser = useAppSelector((state) => state.auth.currentUser)
  
  const [deletePost] = useDeletePostMutation()
  const [errorMsg, setErrorMsg] = useState('')
  const submitDeletePost = async () => {
    await deletePost({ id: post.id }).unwrap().then(() => {
      navigate('/')
    }).catch((error) => {
      if (!error.data) {
        setErrorMsg('Failed to access the API')
      } else {
        setErrorMsg(error.data.msg)
      }
    })
  }
  
  const { data: authorData } = useGetUserByIdQuery({ id: post.authorId })
  return (
    <Paper shadow='lg' radius='lg' p='lg' withBorder>
      {authorData && (
        <Group position='apart'>
          <Badge
            pl={0}
            color='teal'
            leftSection={
              <ThemeIcon color='teal' size={24} radius='xl' mr={5}>
                <IconUser size={12} />
              </ThemeIcon>
            }
            component={Link}
            to={`/user/${authorData.username}`}
          >
            {`user/${authorData.username}`}
          </Badge>
          {(currentUser && (currentUser.id == post.authorId || currentUser.isAdmin)) && (
            <Menu>
              <Menu.Target>
                <ActionIcon color='teal' variant='light' radius='xl' size={24}>
                  <IconMenu2 size={12} />
                </ActionIcon>
              </Menu.Target>
              <Menu.Dropdown>
                <Menu.Label>Post Options</Menu.Label>
                { currentUser.id == post.authorId && (
                  modifying ? <Menu.Item icon={<IconPencil size={14} />} onClick={() => setModifying(false)}>Stop Modifying</Menu.Item>
                  : <Menu.Item icon={<IconPencil size={14} />} onClick={() => setModifying(true)}>Modify</Menu.Item>
                )}
                <Menu.Item color='red' icon={<IconTrash size={14} />} onClick={() => submitDeletePost()}>Delete</Menu.Item>
              </Menu.Dropdown>
            </Menu>
          )}
        </Group>
      )}

      {modifying ? <ModifyPostForm post={post} setModifying={setModifying} setPost={setPost} /> : (
        <Stack align='flex-start' mt={2} spacing={1}>
          <Text weight={600}>{post.title}</Text>
          <Text size='sm'>{post.content}</Text>
          <Text size='xs' color='dimmed' mt={3}>Created {post.createdAt}</Text>
        </Stack>
      )}

      {errorMsg && <Text color='red' align='center' size='xs' mt='md'>{'Error: ' + errorMsg}</Text>}
    </Paper>
  )
}

const PagePost = ({ post: initialPost }: { post: Post }) => {
  const [post, setPost] = useState<Post>(initialPost)
  return <PagePostItem post={post} setPost={setPost} />
}

export default PagePost