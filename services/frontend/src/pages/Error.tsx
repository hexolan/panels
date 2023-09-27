import { Link, useRouteError, isRouteErrorResponse } from 'react-router-dom'
import { Title, Text, Button, Center, Container, Group, rem } from '@mantine/core'

const ErrorPage = () => {
  const error = useRouteError()

  let title = 'Uh, oh!'
  let subTitle = 'Something went wrong.'
  if (isRouteErrorResponse(error)) {
    title = `Error ${error.status}`
    subTitle = error.statusText
  } else if (error instanceof Error) {
    subTitle = error.message
  }

  return (
    <Center h='100%'>
      <Container>
        <Title
          align='center'
          weight={800}
          sx={(theme) => ({
            fontSize: rem(38),
            [theme.fn.smallerThan('sm')]: {
              fontSize: rem(32),
            },
          })}
        >
          {title}
        </Title>
        <Text size='lg' color='dimmed' maw={rem(250)} align='center' my='xl'>
          {subTitle}
        </Text>

        <Group position='center'>
          <Button component={Link} to='/' variant='subtle' size='md'>
            Back to Home
          </Button>
        </Group>
      </Container>
    </Center>
  )
}

export default ErrorPage