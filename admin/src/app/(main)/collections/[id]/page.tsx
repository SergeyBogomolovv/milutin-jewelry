import { CollectionCard, getCollection } from '@/entities/collection'

export default async function CollectionPage({ params }: { params: Promise<{ id: string }> }) {
  const { id } = await params
  const collection = await getCollection(id)
  return (
    <main>
      <CollectionCard collection={collection} />
    </main>
  )
}
