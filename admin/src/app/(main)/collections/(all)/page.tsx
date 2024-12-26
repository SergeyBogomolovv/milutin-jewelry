import { CollectionCard, getCollections } from '@/entities/collection'

export default async function CollectionsPage() {
  const { data: collections, success } = await getCollections()

  return (
    <main className='grid md:grid-cols-2 lg:grid-cols-3 gap-4 px-6 py-2'>
      {!success && (
        <h1 className='text-3xl font-semibold text-red-500 col-span-3'>
          Ошибка загрузки коллекций.
        </h1>
      )}
      {collections && (
        <>
          {collections.length === 0 && (
            <h1 className='text-2xl font-semibold col-span-3'>
              Вы пока не создали ни одну коллекцию.
            </h1>
          )}
          {collections.map((collection) => (
            <CollectionCard key={collection.id} collection={collection} />
          ))}
        </>
      )}
    </main>
  )
}
