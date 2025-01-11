import { CollectionCard, getCollections } from '@/entities/collection'
import { CollectionForm } from '@/features/collection-form'
import { Button } from '@/shared/ui/button'
import { Plus } from 'lucide-react'

export const dynamic = 'force-dynamic'

export default async function CollectionsPage() {
  const collections = await getCollections()

  if (!collections.length) {
    return (
      <main className='flex flex-col items-center justify-center grow w-full gap-6'>
        <h1 className='font-bold text-4xl'>Вы пока не создали ни одной коллекции</h1>
        <CollectionForm>
          <Button variant={'outline'} size='lg'>
            <Plus />
            Создать коллекцию
          </Button>
        </CollectionForm>
      </main>
    )
  }

  return (
    <main className='flex flex-col gap-4 px-6 py-2'>
      <section className='grid md:grid-cols-2 lg:grid-cols-3 gap-4'>
        {collections.map((collection) => (
          <CollectionCard key={collection.id} collection={collection} />
        ))}
      </section>
    </main>
  )
}
