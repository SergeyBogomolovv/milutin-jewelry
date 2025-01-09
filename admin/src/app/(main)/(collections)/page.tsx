import { CollectionCard, getCollections } from '@/entities/collection'

export default async function CollectionsPage() {
  const { data: collections, success } = await getCollections()

  if (!success) {
    return (
      <main className='flex flex-col items-center justify-center grow w-full'>
        <h1 className='font-bold text-4xl text-red-500'>Ошибка загрузки коллекций</h1>
      </main>
    )
  }

  if (collections?.length === 0) {
    return (
      <main className='flex flex-col items-center justify-center grow w-full'>
        <h1 className='font-bold text-4xl'>Вы пока не создали ни одной коллекции</h1>
      </main>
    )
  }

  return (
    <main className='flex flex-col gap-4 px-6 py-2'>
      {collections && (
        <section className='grid md:grid-cols-2 lg:grid-cols-3 gap-4'>
          {collections.map((collection) => (
            <CollectionCard key={collection.id} collection={collection} />
          ))}
        </section>
      )}
    </main>
  )
}
