import { Suspense } from 'react'
import { List } from './list'

export function Collections() {
  return (
    <section>
      <h1 className='bg-neutral-700 h-40 lg:my-20 sm:my-16 my-12 text-5xl flex items-center justify-center tracking-wider'>
        Коллекции
      </h1>
      {/* TODO: add loading sceletons */}
      <Suspense fallback={<span>Загрузка...</span>}>
        <List />
      </Suspense>
    </section>
  )
}
