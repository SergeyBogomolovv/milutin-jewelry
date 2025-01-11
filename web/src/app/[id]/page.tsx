import { getCollection, getCollections } from '@/entities/collection'
import { notFound } from 'next/navigation'
import type { Metadata } from 'next'

type Props = {
  params: Promise<{ id: string }>
  searchParams: Promise<{ [key: string]: string | string[] | undefined }>
}

export async function generateMetadata({ params }: Props): Promise<Metadata> {
  try {
    const id = (await params).id
    const collection = await getCollection(id)
    return { title: `${collection.title} | Milutin Jewellery` }
  } catch (error) {
    return { title: 'Страница не найдена' }
  }
}

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
    return <main className='grow'>{collection.title}</main>
  } catch (error) {
    notFound()
  }
}
