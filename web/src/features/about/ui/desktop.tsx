import { Button } from '@/shared/ui/button'
import { Separator } from '@/shared/ui/separator'
import { ExternalLinkIcon } from 'lucide-react'
import Image from 'next/image'
import mikhail from '@/assets/mikhail.jpg'

export default function AboutDesktop() {
  return (
    <section className='container mx-auto lg:space-y-20 space-y-16 lg:my-20 my-16'>
      <Separator />
      <div className='flex justify-center lg:gap-20 gap-16'>
        <Image src={mikhail} alt='Михал Милютин' className='rounded-lg' />
        <div className='flex flex-col justify-between items-center'>
          <h1 className='text-4xl'>Об авторе</h1>
          <div className='flex flex-col gap-2 items-center'>
            <p className='text-center text-muted-foreground tracking-wider text-xl'>
              Художник, создатель драгоценностей.
              <br /> Произведения Михаила Милютина – синтез ювелирного <br /> мастерства и
              художественной фантазии.
            </p>
            <Button variant='link' className='font-bold text-lg'>
              Подробнее
              <ExternalLinkIcon />
            </Button>
          </div>
          <div />
        </div>
      </div>
      <Separator />
    </section>
  )
}
