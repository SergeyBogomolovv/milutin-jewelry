import { Carousel } from '@/features/carousel'
import { About } from '@/features/about'
import { Collections } from '@/features/collections'

export default function Home() {
  return (
    <main>
      <Carousel />
      <About />
      <Collections />
    </main>
  )
}
