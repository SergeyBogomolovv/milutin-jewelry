import { getCollection } from '@/entities/collection'
import { ItemForm } from '@/features/item-form'
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
  try {
    const { id } = await params
    const collection = await getCollection(id)

    return (
      <div className='flex flex-col w-full'>
        <header className='p-4 items-center flex pt-4'>
          <SidebarTrigger />
          <div className='flex items-center justify-between w-full'>
            <h2 className='font-bold text-lg'>{collection.title}</h2>
            <ItemForm id={id}>
              <Button variant={'outline'}>
                <Plus />
                Добавить украшение
              </Button>
            </ItemForm>
          </div>
        </header>
        {children}
      </div>
    )
  } catch (error) {
    return (
      <div className='flex flex-col w-full'>
        <header className='p-4 items-center flex pt-4'>
          <SidebarTrigger />
          <div className='flex items-center justify-between w-full'>
            <h2 className='font-bold text-lg'>Коллекция не найдена</h2>
          </div>
        </header>
        {children}
      </div>
    )
  }
}
