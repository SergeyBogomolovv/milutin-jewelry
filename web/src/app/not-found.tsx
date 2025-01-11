'use client'
import { Button } from '@/shared/ui/button'
import { useRouter } from 'next/navigation'

export default function NotFoundPage() {
  const router = useRouter()
  return (
    <main className='flex flex-col items-center justify-center grow w-full gap-6'>
      <h1 className='font-bold xl:text-5xl md:text-4xl text-3xl tracking-widest text-center'>
        Страница не найдена
      </h1>
      <Button size='lg' className='font-bold text-lg tracking-widest' onClick={() => router.back()}>
        Вернуться назад
      </Button>
    </main>
  )
}
