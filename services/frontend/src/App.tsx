import { lazy } from 'react'
import { MantineProvider } from '@mantine/core'
import { RouterProvider, createBrowserRouter } from 'react-router-dom'

import AppLayout from './components/AppLayout'
import LoadingBar from './components/LoadingBar'
import ErrorPage from './pages/Error'

const Homepage = lazy(() => import('./pages/Home'))

const SignInPage = lazy(() => import('./pages/SignIn'))
const SignUpPage = lazy(() => import('./pages/SignUp'))

const UserLayout = lazy(() => import('./components/UserLayout'))
const UserPage = lazy(() => import('./pages/User'))
const UserAboutPage = lazy(() => import('./pages/UserAbout'))
const UserSettingsPage = lazy(() => import('./pages/UserSettings'))

const ExplorePanelsPage = lazy(() => import('./pages/ExplorePanels'))
const NewPanelPage = lazy(() => import('./pages/NewPanel'))

const PanelLayout = lazy(() => import('./components/PanelLayout'))
const PanelPage = lazy(() => import('./pages/Panel'))
const PanelSettingsPage = lazy(() => import('./pages/PanelSettings'))
const PanelPostPage = lazy(() => import('./pages/PanelPost'))
const NewPanelPostPage = lazy(() => import('./pages/NewPanelPost'))

const router = createBrowserRouter([
  {
    element: <AppLayout />,
    errorElement: <AppLayout><ErrorPage /></AppLayout>,
    children: [
      {
        index: true,
        element: <Homepage />,
      },
      {
        path: '/signin',
        element: <SignInPage />,
      },
      {
        path: '/signup',
        element: <SignUpPage />,
      },
      {
        path: '/user/:username',
        element: <UserLayout />,
        children: [
          {
            index: true,
            element: <UserPage />,
          },
          {
            path: '/user/:username/about',
            element: <UserAboutPage />,
          },
          {
            path: '/user/:username/settings',
            element: <UserSettingsPage />,
          },
        ],
      },
      {
        path: '/panels',
        children: [
          {
            index: true,
            element: <ExplorePanelsPage />,
          },
          {
            path: '/panels/new',
            element: <NewPanelPage />,
          },
        ]
      },
      {
        path: '/panel/:panelName',
        element: <PanelLayout />,
        children: [
          {
            index: true,
            element: <PanelPage />,
          },
          {
            path: '/panel/:panelName/settings',
            element: <PanelSettingsPage />,
          },
          {
            path: '/panel/:panelName/post/:postId',
            element: <PanelPostPage />,
          },
          {
            path: '/panel/:panelName/posts/new',
            element: <NewPanelPostPage />,
          }
        ],
      },
    ]
  }
])

function App() {
  return (
    <MantineProvider withGlobalStyles withNormalizeCSS>
      <RouterProvider router={router} fallbackElement={<LoadingBar />} />
    </MantineProvider>
  );
}

export default App