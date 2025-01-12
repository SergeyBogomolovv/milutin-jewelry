import type { Metadata } from 'next'
import { Separator } from '@/shared/ui/separator'
import { AboutInfo } from '@/sections/about/info'
import { Awards } from '@/sections/about/awards'

export const metadata: Metadata = {
  title: 'Об Авторе',
  description: 'История и достижения Михаила Милютина',
  keywords: ['Достижения Михаила Милютина', 'История Михаила Милютина', 'Михаил Милютин'],
}

export default function AboutPage() {
  return (
    <main className='grow xl:w-9/12 lg:w-11/12 sm:w-9/12 w-11/12 mx-auto flex flex-col md:gap-10 gap-6 items-center md:py-10 py-6'>
      <h1 className='text-5xl text-center tracking-wide'>Об авторе</h1>
      <AboutInfo />
      <Separator />
      <h2 className='text-4xl text-center tracking-wide'>Достижения</h2>
      <Awards />
    </main>
  )
}
