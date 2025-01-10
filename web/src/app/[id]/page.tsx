import { getCollection, getCollections } from '@/entities/collection'

export const revalidate = 60

export const dynamicParams = true

export async function generateStaticParams() {
  try {
    const collections = await getCollections()
    return collections.map((collection) => ({ id: String(collection.id) }))
  } catch (error) {
    return []
  }
}

export default async function CollectionPage({ params }: { params: Promise<{ id: string }> }) {
  const id = (await params).id
  const collection = await getCollection(id)
  return <div>{collection.title}</div>
}
