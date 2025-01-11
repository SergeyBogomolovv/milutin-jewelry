'use client'
import { Button } from '@/shared/ui/button'

export default function Error({ reset }: { reset: () => void }) {
  return (
    <main className='flex flex-col items-center justify-center grow w-full min-h-screen gap-6'>
      <h1 className='font-bold text-4xl text-red-500'>Что-то пошло не так!</h1>
      <Button size='lg' onClick={() => reset()}>
        Попробовать ещё раз
      </Button>
    </main>
  )
}
