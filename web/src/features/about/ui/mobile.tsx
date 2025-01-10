import { Button } from '@/shared/ui/button'
import { Separator } from '@/shared/ui/separator'
import { ExternalLinkIcon } from 'lucide-react'
import Image from 'next/image'
import mikhail from '@/assets/mikhail.jpg'

export default function AboutMobile() {
  return (
    <section className='w-11/12 mx-auto flex flex-col items-center gap-8 my-8'>
      <Separator />
      <h1 className='text-4xl'>Об авторе</h1>
      <Image priority src={mikhail} alt='Михал Милютин' className='rounded-lg w-9/12 sm:w-7/12' />
      <div className='flex flex-col gap-2 items-center'>
        <p className='text-center text-muted-foreground tracking-wider text-lg sm:w-9/12'>
          Художник, создатель драгоценностей. Произведения Михаила Милютина – синтез ювелирного
          мастерства и художественной фантазии.
        </p>
        <Button variant='link' className='font-bold text-lg'>
          Подробнее
          <ExternalLinkIcon />
        </Button>
      </div>
      <Separator />
    </section>
  )
}
