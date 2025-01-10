import { Carousel } from '@/features/carousel'
import { About } from '@/features/about'
import { Collections } from '@/features/collections'
import { getBanners } from '@/entities/banner'

export const dynamic = 'force-dynamic'

export default async function Home() {
  const { data } = await getBanners()

  return (
    <main>
      <Carousel banners={data} />
      <About />
      <Collections />
    </main>
  )
}
