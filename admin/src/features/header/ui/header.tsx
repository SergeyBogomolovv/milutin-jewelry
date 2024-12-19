'use client'
import { SidebarTrigger } from '@/shared/ui/sidebar'
import { usePathname } from 'next/navigation'
import { CollectionForm } from '@/features/collection-form'
import { Button } from '@/shared/ui/button'
import { Plus } from 'lucide-react'

export function AppHeader() {
  const pathname = usePathname()
  return (
    <header className='p-4 items-center flex pt-4'>
      <SidebarTrigger />
      <div className='flex items-center justify-between w-full'>
        <h2 className='font-bold text-lg'>{getHeader(pathname)}</h2>
        {pathname.startsWith('/collections') && (
          <CollectionForm>
            <Button variant={'outline'}>
              <Plus />
              Создать коллекцию
            </Button>
          </CollectionForm>
        )}
      </div>
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
