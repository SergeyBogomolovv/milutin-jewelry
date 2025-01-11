'use client'
import { Button } from '@/shared/ui/button'

export default function Error({ reset }: { reset: () => void }) {
  return (
    <main className='flex flex-col items-center justify-center grow w-full gap-6'>
      <h1 className='font-bold xl:text-5xl md:text-4xl sm:text-3xl text-2xl tracking-widest text-center'>
        Произошла непредвиденная ошибка!
      </h1>
      <Button size='lg' className='font-bold text-lg tracking-widest' onClick={() => reset()}>
        Попробовать ещё раз
      </Button>
    </main>
  )
}
