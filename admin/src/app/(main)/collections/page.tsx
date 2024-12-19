import { CollectionCard, getCollections } from '@/entities/collection'

export default async function CollectionsPage() {
  const collections = await getCollections()
  return (
    <main className='grid md:grid-cols-2 lg:grid-cols-3 gap-4 px-6 py-2'>
      {collections.map((collection) => (
        <CollectionCard key={collection.id} collection={collection} />
      ))}
    </main>
  )
}
