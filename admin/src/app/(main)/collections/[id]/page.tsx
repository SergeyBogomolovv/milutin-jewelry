import { getCollection } from '@/entities/collection'
import { CollectionItemCard, getCollectionItemsByCollection } from '@/entities/collection-item'
import { CollectionInfo } from '@/features/collection-info'

export default async function CollectionPage({ params }: { params: Promise<{ id: string }> }) {
  const { id } = await params
  const collection = await getCollection(id)
  const items = await getCollectionItemsByCollection(id)

  return (
    <main className='flex flex-col gap-4 px-6 py-2'>
      <section>
        <CollectionInfo collection={collection} />
      </section>
      <h2 className='font-bold text-2xl'>Украшения:</h2>
      {items.length === 0 && <h2 className='font-semibold text-xl'>Украшений в коллекции нет</h2>}
      <section className='grid md:grid-cols-2 lg:grid-cols-3 gap-4'>
        {items.map((item) => (
          <CollectionItemCard key={item.id} item={item} />
        ))}
      </section>
    </main>
  )
}
