import { Button } from '@/shared/ui/button'
import { SidebarTrigger } from '@/shared/ui/sidebar'
import { Plus } from 'lucide-react'

export default function CollectionItemsLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  return (
    <div className='flex flex-col w-full'>
      <header className='p-4 items-center flex pt-4'>
        <SidebarTrigger />
        <div className='flex items-center justify-between w-full'>
          <h2 className='font-bold text-lg'>Украшения</h2>
          <Button variant={'outline'}>
            <Plus />
            Добавить украшение
          </Button>
        </div>
      </header>
      {children}
    </div>
  )
}
