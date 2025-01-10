import { Collection, CollectionCard } from '@/entities/collection'

export function Collections({ collections }: { collections: Collection[] }) {
  return (
    <section>
      <h1 className='bg-neutral-700 h-40 lg:my-20 sm:my-16 my-12 text-5xl flex items-center justify-center tracking-wider'>
        Коллекции
      </h1>
      {/* TODO: add loading sceletons */}
      <div className='grid lg:grid-cols-3 sm:grid-cols-2 gap-10 container w-11/12 mx-auto'>
        {collections.map((collection) => (
          <CollectionCard key={collection.id} collection={collection} />
        ))}
      </div>
    </section>
  )
}
