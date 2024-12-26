import { getCollection } from '@/entities/collection'
import { CollectionItemCard, getCollectionItemsByCollection } from '@/entities/collection-item'
import { CollectionInfo } from '@/features/collection-info'

export default async function CollectionPage({ params }: { params: Promise<{ id: string }> }) {
  const { id } = await params

  const [{ data: collection, success: collectionSuccess }, { data: items, success: itemsSuccess }] =
    await Promise.all([getCollection(id), getCollectionItemsByCollection(id)])

  const hasItems = items && items.length > 0

  return (
    <main className='flex flex-col gap-4 px-6 py-2'>
      <section>
        {!collectionSuccess ? (
          <h1 className='text-3xl font-semibold text-red-500'>Ошибка загрузки коллекции.</h1>
        ) : (
          collection && <CollectionInfo collection={collection} />
        )}
      </section>

      <h2 className='font-bold text-2xl'>Украшения:</h2>

      {!itemsSuccess ? (
        <h1 className='text-2xl font-semibold text-red-500'>Ошибка загрузки украшений.</h1>
      ) : (
        <section className='grid md:grid-cols-2 lg:grid-cols-3 gap-4'>
          {!hasItems ? (
            <h2 className='font-semibold text-xl'>Украшений в коллекции нет</h2>
          ) : (
            items.map((item) => <CollectionItemCard key={item.id} item={item} />)
          )}
        </section>
      )}
    </main>
  )
}
