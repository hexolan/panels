import { useState } from 'react'

import { Link, useNavigate } from 'react-router-dom'
import { useForm, hasLength, matchesField } from '@mantine/form'
import { Center, Container, Paper, Title, Text, Anchor, TextInput, PasswordInput, Stack, Button } from '@mantine/core'

import { useAppSelector } from '../app/hooks'
import { useRegisterUserMutation } from '../app/api/users'

type RegistrationFormValues = {
  username: string;
  password: string;
  confPassword: string;
}

const SignUpPage = () => {
  const navigate = useNavigate()
  
  // Ensure the user isn't authenticated already
  const currentUser = useAppSelector((state) => state.auth.currentUser)
  if (currentUser) {
    throw new Error('You are already authenticated.')
  }

  const [errorMsg, setErrorMsg] = useState('')
  const registrationForm = useForm<RegistrationFormValues>({
    initialValues: {
      username: '',
      password: '',
      confPassword: '',
    },
    validate: {
      username: hasLength({ min: 3, max: 32 }, 'Username must be between 3 and 32 characters'),
      password: hasLength({ min: 8 }, 'Password must have a minimum of 8 characters'),
      confPassword: matchesField('password', 'Confirmation password does not match'),
    }
  })

  const [registerUser, { isLoading }] = useRegisterUserMutation()
  const submitRegistrationForm = async (values: RegistrationFormValues) => {
    await registerUser({
      username: values.username,
      password: values.password
    }).unwrap().then(() => {
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
        <Title align='center' weight={900}>Sign Up</Title>
        <Text color='dimmed' size='sm' align='center' mt={5}>
          Already have an account?{' '}
          <Anchor size='sm' component={Link} to='/signin'>Sign in</Anchor> instead.
        </Text>

        <Paper withBorder shadow='md' radius='md' p={30} mt={30}>
          <form onSubmit={registrationForm.onSubmit(submitRegistrationForm)}>
            <Stack spacing='md'>
              <TextInput 
                label='Username'
                placeholder='Your username'
                {...registrationForm.getInputProps('username')}
              />

              <PasswordInput 
                label='Password'
                placeholder='Your password'
                {...registrationForm.getInputProps('password')}
              />

              <PasswordInput
                label='Confirm Password'
                placeholder='Confirm password'
                {...registrationForm.getInputProps('confPassword')}
              />

              { errorMsg && <Text color='red' align='center'>{'Error: ' + errorMsg}</Text> }

              <Button type='submit' color='teal' disabled={isLoading} fullWidth>Register</Button>
            </Stack>
          </form>
        </Paper>
      </Container>
    </Center>
  )
}

export default SignUpPage