'use client'

import { Button } from '@/shared/ui/button'

export default function ErrorPage({ reset }: { reset: () => void }) {
  return (
    <main className='flex flex-col items-center justify-center grow w-full gap-6'>
      <h1 className='font-bold text-4xl text-red-500'>Ошибка загрузки коллекций</h1>
      <Button size='lg' onClick={() => reset()}>
        Попробовать ещё раз
      </Button>
    </main>
  )
}
