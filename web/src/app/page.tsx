import { Carousel } from '@/features/carousel'
import { About } from '@/features/about'
import { Collections } from '@/features/collections'
import { getBanners } from '@/entities/banner'
import { getCollections } from '@/entities/collection'
import { Metadata } from 'next'

export const metadata: Metadata = {
  title: 'Главная | Milutin Jewellery',
}

export const revalidate = 60

export default async function Home() {
  const banners = await getBanners()
  const collections = await getCollections()

  return (
    <main className='grow'>
      <Carousel banners={banners} />
      <About />
      <Collections collections={collections} />
    </main>
  )
}
