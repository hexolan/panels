import { ReactNode, Suspense } from 'react'
import { AppShell } from '@mantine/core'
import { Outlet } from 'react-router-dom'

import AppNavbar from './AppNavbar'
import AppHeader from './AppHeader'
import LoadingBar from './LoadingBar'

interface AppLayoutProps {
  children?: ReactNode;
}

function AppLayout(props: AppLayoutProps) {
  return (
    <AppShell
      navbar={<AppNavbar />} 
      header={<AppHeader />}
      padding={0}
    >
      <Suspense fallback={<LoadingBar />}>
        {props?.children ? props.children : <Outlet /> }
      </Suspense>
    </AppShell>
  );
}

export default AppLayout