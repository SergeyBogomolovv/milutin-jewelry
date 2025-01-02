import { CollectionForm } from '@/features/collection-form'
import { Button } from '@/shared/ui/button'
import { SidebarTrigger } from '@/shared/ui/sidebar'
import { Plus } from 'lucide-react'

export default function CollectionsLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  return (
    <div className='flex flex-col w-full'>
      <header className='p-4 items-center flex pt-4'>
        <SidebarTrigger />
        <div className='flex items-center justify-between w-full'>
          <h2 className='font-bold text-lg'>Коллекции</h2>
          <CollectionForm>
            <Button variant={'outline'}>
              <Plus />
              Создать коллекцию
            </Button>
          </CollectionForm>
        </div>
      </header>
      {children}
    </div>
  )
}
