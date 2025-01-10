import { Carousel } from '@/features/carousel'
import { About } from '@/features/about'
import { Collections } from '@/features/collections'
import { getBanners } from '@/entities/banner'
import { getCollections } from '@/entities/collection'

export const revalidate = 60

export default async function Home() {
  const banners = await getBanners()
  const collections = await getCollections()

  return (
    <main>
      <Carousel banners={banners} />
      <About />
      <Collections collections={collections} />
    </main>
  )
}
