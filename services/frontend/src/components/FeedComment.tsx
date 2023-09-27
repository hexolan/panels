import { useState } from 'react'
import { Link } from 'react-router-dom'
import { useForm, hasLength } from '@mantine/form'
import { Paper, Group, Box, ThemeIcon, Text, ActionIcon, Menu, Textarea, Flex } from '@mantine/core'
import { IconMessage, IconMenu2, IconTrash, IconPencil, IconPencilCancel } from '@tabler/icons-react'

import { useAppSelector } from '../app/hooks'
import { useGetUserByIdQuery } from '../app/api/users'
import { useDeletePostCommentMutation, useUpdatePostCommentMutation } from '../app/api/comments'
import type { Comment } from '../app/types/common'
import type { UpdateCommentData } from '../app/types/comments'

const FeedCommentBase = ({ children, extraChildren }: { children: React.ReactNode, extraChildren?: React.ReactNode }) => (
  <Paper shadow='sm' radius='md' p='md' withBorder>
    <Group w='100%' position='apart'>
      <Group>
        <ThemeIcon color='teal' variant='light' size='xl'><IconMessage /></ThemeIcon>
        {children}
      </Group>
      {extraChildren}
    </Group>
  </Paper>
)

const StandardFeedComment = ({ comment, authorElement }: { comment: Comment, authorElement: React.ReactNode }) => (
  <FeedCommentBase>
    <Box>
      <Text size='sm'>{comment.message}</Text>
      {authorElement}
    </Box>
  </FeedCommentBase>
)

const ModifiableFeedComment = ({
  comment,
  authorElement,
  setSelf,
  isAuthor
}: {
  comment: Comment,
  authorElement: React.ReactNode,
  setSelf: React.Dispatch<Comment | undefined>,
  isAuthor: boolean
}) => {
  const [modifying, setModifying] = useState<boolean>(false)
  const [errorMsg, setErrorMsg] = useState('')
  const commentForm = useForm<UpdateCommentData>({
    initialValues: {
      message: comment.message,
    },
    validate: {
      message: hasLength({ min: 3, max: 512 }, 'Message must be between 3 and 512 characters'),
    }
  })

  const [updateComment, { isLoading }] = useUpdatePostCommentMutation()
  const submitUpdateComment = async (values: UpdateCommentData) => {
    await updateComment({
      id: comment.id,
      postId: comment.postId,
      data: values
    }).unwrap().then((commentInfo) => {
      setSelf(commentInfo)
      setModifying(false)
    }).catch((error) => {
      if (!error.data) {
        setErrorMsg('Failed to access the API')
      } else {
        setErrorMsg(error.data.msg)
      }
    })
  }
  
  const [deleteComment] = useDeletePostCommentMutation()
  const submitDeleteComment = async () => {
    await deleteComment({
      id: comment.id,
      postId: comment.postId
    }).unwrap().then(() => {
      setSelf(undefined)
    }).catch((error) => {
      if (!error.data) {
        setErrorMsg('Failed to access the API')
      } else {
        setErrorMsg(error.data.msg)
      }
    })
  }

  return (
    <FeedCommentBase
      extraChildren={
        <Menu>
          <Menu.Target>
            <ActionIcon color='teal' variant='light' radius='xl' size='xl'><IconMenu2 /></ActionIcon>
          </Menu.Target>
          <Menu.Dropdown>
            <Menu.Label>Comment Options</Menu.Label>
            {isAuthor && (modifying
              ? <Menu.Item icon={<IconPencilCancel size={14} />} onClick={() => setModifying(false)}>Stop Modifying</Menu.Item>
              : <Menu.Item icon={<IconPencil size={14} />} onClick={() => setModifying(true)}>Modify</Menu.Item>
            )}
            <Menu.Item color='red' icon={<IconTrash size={14} />} onClick={() => submitDeleteComment()}>Delete</Menu.Item>
          </Menu.Dropdown>
        </Menu>
      }
    >
      {modifying ? (
        <form onSubmit={commentForm.onSubmit(submitUpdateComment)}>
          <Flex>
            <Textarea size='xs' w='100%' radius='lg' variant='filled' error={errorMsg} {...commentForm.getInputProps('message')} />
            <ActionIcon type='submit' radius='lg' color='teal' variant='outline' size='xl' aria-label='Update Comment' disabled={isLoading}>
              <IconPencil />
            </ActionIcon>
          </Flex>
        </form>
      ) : (
        <Box>
          <Text size='sm'>{comment.message}</Text>
          {authorElement}
        </Box>
      )}
    </FeedCommentBase>
  )
}

const FeedCommentItem = ({ comment, setSelf }: { comment: Comment, setSelf: React.Dispatch<Comment | undefined> }) => {
  const currentUser = useAppSelector((state) => state.auth.currentUser)

  // fetching comment author info
  const { data, isLoading } = useGetUserByIdQuery({ id: comment.authorId })
  let authorElement = null
  if (isLoading) {
    authorElement = <Text color='dimmed' size='xs'>Loading Author Info...</Text>
  } else if (!data) {
    authorElement = <Text color='red' size='xs'>Failed to load Author Info</Text>
  } else {
    authorElement = <Text color='dimmed' size='xs' mt={3} component={Link} to={`/user/${data.username}`}>by user/{data.username}</Text>
  }

  if (currentUser && (currentUser.id == comment.authorId || currentUser.isAdmin)) {
    return <ModifiableFeedComment comment={comment} authorElement={authorElement} isAuthor={currentUser.id == comment.authorId} setSelf={setSelf} />
  } else {
    return <StandardFeedComment comment={comment} authorElement={authorElement} />
  }
}

const FeedComment = ({ comment: initialComment }: { comment: Comment }) => {
  const [comment, setComment] = useState<Comment | undefined>(initialComment)
  return comment ? <FeedCommentItem comment={comment} setSelf={setComment} /> : null
}

export default FeedComment