import { BannerCard, getBanners } from '@/entities/banner'
import { getCollections } from '@/entities/collection'
import { BannerForm } from '@/features/banner-form'
import { Button } from '@/shared/ui/button'
import { Plus } from 'lucide-react'

export const dynamic = 'force-dynamic'

export default async function BannersPage() {
  const banners = await getBanners()
  const collections = await getCollections()

  if (!banners.length) {
    return (
      <main className='flex flex-col items-center justify-center grow w-full gap-6'>
        <h1 className='font-bold text-4xl'>Вы пока не добавили ни одного баннера</h1>
        <BannerForm collections={collections}>
          <Button variant={'outline'} size='lg'>
            <Plus />
            Добавить баннер
          </Button>
        </BannerForm>
      </main>
    )
  }

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
