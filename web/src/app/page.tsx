import { Carousel } from '@/sections/main/carousel'
import { About } from '@/sections/main/about'
import { Collections } from '@/sections/main/collections'
import { getBanners } from '@/entities/banner'
import { getCollections } from '@/entities/collection'
import { Metadata } from 'next'

export const metadata: Metadata = {
  title: 'Главная',
  description:
    'Михаил Милютин — Художник, создатель драгоценностей. Произведения Михаила Милютина – синтез ювелирного мастерства и художественной фантазии.',
  keywords: ['Коллекции Михаила Милютина', 'Контакты Михаила Милютина'],
}

export const dynamic = 'force-dynamic'

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
