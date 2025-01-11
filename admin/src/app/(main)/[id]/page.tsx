import { ItemCard, getItemsByCollection } from '@/entities/item'
import { ItemForm } from '@/features/item-form'
import { Button } from '@/shared/ui/button'
import { Plus } from 'lucide-react'

export const dynamic = 'force-dynamic'

export default async function CollectionPage({ params }: { params: Promise<{ id: string }> }) {
  const { id } = await params
  const items = await getItemsByCollection(id)

  if (!items.length) {
    return (
      <main className='flex flex-col items-center justify-center grow w-full gap-6'>
        <h1 className='font-bold text-4xl'>Вы пока не добавили ни одного украшения в коллекцию</h1>
        <ItemForm id={id}>
          <Button variant={'outline'} size='lg'>
            <Plus />
            Добавить украшение
          </Button>
        </ItemForm>
      </main>
    )
  }
  return (
    <main className='flex flex-col gap-4 px-6 py-2'>
      <section className='grid md:grid-cols-2 lg:grid-cols-3 gap-4'>
        {items.map((item) => (
          <ItemCard key={item.id} item={item} />
        ))}
      </section>
    </main>
  )
}
