'use client'
import { SidebarTrigger } from '@/shared/ui/sidebar'
import { usePathname } from 'next/navigation'

export function AppHeader() {
  const pathname = usePathname()
  return (
    <header className='p-2 items-center flex pt-4'>
      <SidebarTrigger />
      <h2 className='font-bold text-lg'>{getHeader(pathname)}</h2>
    </header>
  )
}

function getHeader(pathname: string): string {
  if (pathname.startsWith('/collections')) {
    return 'Коллекции'
  }
  if (pathname.startsWith('/collection-items')) {
    return 'Украшения'
  }
  if (pathname.startsWith('/posts')) {
    return 'Статьи'
  }
  return 'Milutin jewelry'
}
