import { useState } from 'react'
import { useNavigate, useOutletContext } from 'react-router-dom'
import { Paper, Text, Button } from '@mantine/core'

import { setUnauthed } from '../app/features/auth'
import { useDeleteUserByIdMutation } from '../app/api/users'
import { useAppSelector, useAppDispatch } from '../app/hooks'
import type { UserContext } from '../components/UserLayout'

function UserSettingsPage() {
  const navigate = useNavigate()
  const dispatch = useAppDispatch()
  const [errorMsg, setErrorMsg] = useState('')
  const { user } = useOutletContext<UserContext>()
  
  const currentUser = useAppSelector((state) => state.auth.currentUser)
  if (user && (!currentUser || (currentUser.id != user.id && !currentUser.isAdmin))) {
    throw Error('You do not have permission to view that page')
  }

  const [deleteUser, { isLoading }] = useDeleteUserByIdMutation()
  const submitDeleteAccount = async () => {
    await deleteUser({id: user.id}).unwrap().then(() => {
      dispatch(setUnauthed())
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
      <Button color='red' onClick={() => submitDeleteAccount()} disabled={isLoading}>Delete Account</Button>
      { errorMsg && <Text color='red'>{'Error: ' + errorMsg}</Text> }
    </Paper>
  )
}

export default UserSettingsPage