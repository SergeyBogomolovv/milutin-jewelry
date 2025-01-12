import { getCollection } from '@/entities/collection'
import { notFound } from 'next/navigation'
import type { Metadata } from 'next'
import { Separator } from '@/shared/ui/separator'
import { PreviewGroup } from '@/shared/ui/image'
import { getItems } from '@/entities/item'
import { ItemCard } from '@/entities/item'
import Link from 'next/link'
import { FaHome } from 'react-icons/fa'

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

export const dynamic = 'force-dynamic'

export default async function CollectionPage({ params }: { params: Promise<{ id: string }> }) {
  try {
    const id = (await params).id
    const collection = await getCollection(id)
    const items = await getItems(id)

    return (
      <main className='grow xl:w-10/12 lg:w-11/12 sm:w-10/12 w-11/12 mx-auto flex flex-col md:gap-10 gap-6 items-center md:py-10 py-6'>
        <h1 className='sm:text-5xl text-4xl text-center tracking-wide'>{collection.title}</h1>
        {collection.description && (
          <p className='sm:text-xl md:w-10/12 lg:w-3/4 text-lg text-center tracking-wide'>
            {collection.description}
          </p>
        )}
        <Separator />
        <div className='grid lg:grid-cols-3 sm:grid-cols-2 gap-8'>
          <PreviewGroup>
            {items.map((item) => (
              <ItemCard key={item.id} item={item} />
            ))}
          </PreviewGroup>
        </div>
        <Link
          href='/'
          className='tracking-wider text-xl flex items-center gap-2 border-b-2 hover:border-white border-transparent hover:font-bold mt-4'
        >
          <FaHome className='size-5' /> На главную
        </Link>
      </main>
    )
  } catch (error) {
    notFound()
  }
}
