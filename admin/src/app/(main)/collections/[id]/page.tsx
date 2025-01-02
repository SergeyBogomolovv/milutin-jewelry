import { ItemCard, getItemsByCollection } from '@/entities/item'

export default async function CollectionPage({ params }: { params: Promise<{ id: string }> }) {
  const { id } = await params

  const { data: items, success } = await getItemsByCollection(id)
  return (
    <main className='flex flex-col gap-4 px-6 py-2'>
      {!success && (
        <h1 className='text-2xl font-semibold text-red-500'>Ошибка загрузки украшений.</h1>
      )}

      {items && (
        <>
          <h2 className='font-bold text-2xl'>
            {items.length > 0
              ? 'Украшения в коллекции:'
              : 'Вы пока не добавили ни одного украшения в коллекцию.'}
          </h2>
          <section className='grid md:grid-cols-2 lg:grid-cols-3 gap-4'>
            {items.map((item) => (
              <ItemCard key={item.id} item={item} />
            ))}
          </section>
        </>
      )}
    </main>
  )
}
