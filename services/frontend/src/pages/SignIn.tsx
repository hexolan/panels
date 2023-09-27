import { useState } from 'react'

import { Link, useNavigate } from 'react-router-dom'
import { useForm, hasLength } from '@mantine/form'
import { Center, Container, Paper, Title, Text, Anchor, TextInput, PasswordInput, Stack, Button } from '@mantine/core'

import { useAppSelector } from '../app/hooks'
import { useLoginMutation } from '../app/api/auth'
import type { LoginRequest } from '../app/types/auth'

function SignInPage() {
  const navigate = useNavigate()

  // Ensure the user isn't authenticated already
  const currentUser = useAppSelector((state) => state.auth.currentUser)
  if (currentUser) {
    throw new Error('You are already authenticated.')
  }

  const [errorMsg, setErrorMsg] = useState('')
  const loginForm = useForm<LoginRequest>({
    initialValues: {
      username: '',
      password: '',
    },
    validate: {
      username: hasLength({ min: 3, max: 32 }, 'Invalid username'),
      password: hasLength({ min: 8 }, 'Invalid password'),
    }
  })

  const [login, { isLoading }] = useLoginMutation()
  const submitLoginForm = async (values: LoginRequest) => {
    // Attempt to authenticate the user.
    await login(values).unwrap().then(() => {
      // Redirect to homepage.
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
    <Center h='95%'>
      <Container>
        <Title align='center' weight={900}>Sign In</Title>
        <Text color='dimmed' size='sm' align='center' mt={5}>
          Do not have an account yet?{' '}
          <Anchor size='sm' component={Link} to='/signup'>Sign up</Anchor> instead.
        </Text>


        <Paper withBorder shadow='md' radius='md' p={30} mt={30}>
          <form onSubmit={loginForm.onSubmit(submitLoginForm)}>
            <Stack spacing='md'>
              <TextInput 
                label='Username'
                placeholder='Your username'
                {...loginForm.getInputProps('username')}
              />

              <PasswordInput 
                label='Password' 
                placeholder='Your password' 
                {...loginForm.getInputProps('password')}
              />

              { errorMsg && <Text color='red' align='center'>{'Error: ' + errorMsg}</Text> }

              <Button type='submit' color='teal' disabled={isLoading} fullWidth>Login</Button>
            </Stack>
          </form>
        </Paper>
      </Container>
    </Center>
  )
}

export default SignInPage