import { AppHeader } from '@/features/header'
import { AppSidebar } from '@/features/sidebar'
import { SidebarProvider } from '@/shared/ui/sidebar'

export default function Layout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  return (
    <SidebarProvider>
      <AppSidebar />
      <div className='flex flex-col w-full'>
        <AppHeader />
        {children}
      </div>
    </SidebarProvider>
  )
}
