import { ItemCard, getItemsByCollection } from '@/entities/item'

export default async function CollectionPage({ params }: { params: Promise<{ id: string }> }) {
  const { id } = await params

  const { data: items, success } = await getItemsByCollection(id)

  if (!success) {
    return (
      <main className='flex flex-col items-center justify-center grow w-full'>
        <h1 className='font-bold text-4xl text-red-500'>Ошибка загрузки украшений</h1>
      </main>
    )
  }
  if (items?.length === 0) {
    return (
      <main className='flex flex-col items-center justify-center grow w-full'>
        <h1 className='font-bold text-4xl'>Вы пока не добавили ни одного украшения в коллекцию</h1>
      </main>
    )
  }
  return (
    <main className='flex flex-col gap-4 px-6 py-2'>
      {items && (
        <section className='grid md:grid-cols-2 lg:grid-cols-3 gap-4'>
          {items.map((item) => (
            <ItemCard key={item.id} item={item} />
          ))}
        </section>
      )}
    </main>
  )
}
