import { BannerCard, getBanners } from '@/entities/banner'
import { getCollections } from '@/entities/collection'

export const revalidate = 60

export default async function BannersPage() {
  const banners = await getBanners()

  if (banners.length === 0) {
    return (
      <main className='flex flex-col items-center justify-center grow w-full'>
        <h1 className='font-bold text-4xl'>Вы пока не добавили ни одного баннера</h1>
      </main>
    )
  }

  const collections = await getCollections()

  return (
    <main className='flex flex-col gap-4 px-6 py-2'>
      <section className='grid md:grid-cols-2 lg:grid-cols-3 gap-4'>
        {banners.map((banner) => (
          <BannerCard key={banner.id} banner={banner} collections={collections} />
        ))}
      </section>
    </main>
  )
}
