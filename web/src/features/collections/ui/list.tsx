import { CollectionCard, getCollections } from '@/entities/collection'
import { use } from 'react'

export function List() {
  const { data, success } = use(getCollections())
  return (
    <div className='grid lg:grid-cols-3 sm:grid-cols-2 gap-10 container w-11/12 mx-auto'>
      {success &&
        data?.map((collection) => <CollectionCard key={collection.id} collection={collection} />)}
    </div>
  )
}
