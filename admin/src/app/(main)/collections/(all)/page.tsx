import { CollectionCard, getCollections } from '@/entities/collection'

export default async function CollectionsPage() {
  const { data: collections, success } = await getCollections()

  return (
    <main className='grid md:grid-cols-2 lg:grid-cols-3 gap-4 px-6 py-2'>
      {!success && (
        <h1 className='text-3xl font-semibold text-red-500'>Ошибка загрузки коллекций.</h1>
      )}
      {collections &&
        collections.map((collection) => (
          <CollectionCard key={collection.id} collection={collection} />
        ))}
    </main>
  )
}
