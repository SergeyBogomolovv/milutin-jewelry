import { Button } from '@/shared/ui/button'
import { Separator } from '@/shared/ui/separator'
import { ExternalLinkIcon } from 'lucide-react'
import Image from 'next/image'
import mikhail from '@/assets/mikhail.jpg'
import Link from 'next/link'

export function About() {
  return (
    <section className='md:container w-11/12 mx-auto lg:gap-20 md:gap-16 gap-8 lg:my-20 md:my-16 my-8 flex flex-col items-center'>
      <Separator />
      {/* Desktop */}
      <div className='hidden md:flex justify-center lg:gap-20 gap-16'>
        <Image priority src={mikhail} alt='Михал Милютин' className='rounded-lg' />
        <div className='flex flex-col justify-between items-center'>
          <h1 className='text-4xl'>Об авторе</h1>
          <div className='flex flex-col gap-2 items-center'>
            <p className='text-center text-muted-foreground tracking-wider text-xl'>
              Художник, создатель драгоценностей.
              <br /> Произведения Михаила Милютина – синтез ювелирного <br /> мастерства и
              художественной фантазии.
            </p>
            <Link href='/about'>
              <Button variant='link' className='font-bold text-lg'>
                Подробнее
                <ExternalLinkIcon />
              </Button>
            </Link>
          </div>
          <div />
        </div>
      </div>

      {/* Mobile */}
      <h1 className='text-4xl md:hidden'>Об авторе</h1>
      <Image
        priority
        src={mikhail}
        alt='Михал Милютин'
        className='rounded-lg w-9/12 sm:w-7/12 md:hidden'
      />
      <div className='flex flex-col gap-2 items-center md:hidden'>
        <p className='text-center text-muted-foreground tracking-wider text-lg sm:w-9/12'>
          Художник, создатель драгоценностей. Произведения Михаила Милютина – синтез ювелирного
          мастерства и художественной фантазии.
        </p>
        <Link href='/about'>
          <Button variant='link' className='font-bold text-lg'>
            Подробнее
            <ExternalLinkIcon />
          </Button>
        </Link>
      </div>
      <Separator />
    </section>
  )
}
