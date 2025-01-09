import { BannerCard, getBanners } from '@/entities/banner'

export default async function BannersPage() {
  const { data, success } = await getBanners()

  if (!success) {
    return (
      <main className='flex flex-col items-center justify-center grow w-full'>
        <h1 className='font-bold text-4xl text-red-500'>Ошибка загрузки баннеров</h1>
      </main>
    )
  }

  if (data?.length === 0) {
    return (
      <main className='flex flex-col items-center justify-center grow w-full'>
        <h1 className='font-bold text-4xl'>Вы пока не добавили ни одного баннера</h1>
      </main>
    )
  }

  return (
    <main className='flex flex-col gap-4 px-6 py-2'>
      {data && (
        <section className='grid md:grid-cols-2 lg:grid-cols-3 gap-4'>
          {data.map((banner) => (
            <BannerCard key={banner.id} banner={banner} />
          ))}
        </section>
      )}
    </main>
  )
}
