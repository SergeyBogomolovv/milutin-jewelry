import { getCollection } from '@/entities/collection'
import { CollectionItemForm } from '@/features/collection-item-form'
import { Button } from '@/shared/ui/button'
import { SidebarTrigger } from '@/shared/ui/sidebar'
import { Plus } from 'lucide-react'

export default async function CollectionLayout({
  children,
  params,
}: Readonly<{
  children: React.ReactNode
  params: Promise<{ id: string }>
}>) {
  const { id } = await params
  //TODO: handle error
  const collection = await getCollection(id)

  return (
    <div className='flex flex-col w-full'>
      <header className='p-4 items-center flex pt-4'>
        <SidebarTrigger />
        <div className='flex items-center justify-between w-full'>
          <h2 className='font-bold text-lg'>{collection.title}</h2>
          <CollectionItemForm id={id}>
            <Button variant={'outline'}>
              <Plus />
              Добавить украшение
            </Button>
          </CollectionItemForm>
        </div>
      </header>
      {children}
    </div>
  )
}
