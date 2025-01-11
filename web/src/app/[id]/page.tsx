import { getCollection, getCollections } from '@/entities/collection'
import { notFound } from 'next/navigation'

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
  try {
    const id = (await params).id
    const collection = await getCollection(id)
    return <div>{collection.title}</div>
  } catch (error) {
    notFound()
  }
}
