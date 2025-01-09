import { getCollections } from '@/entities/collection'
import { BannerForm } from '@/features/banner-form'
import { Button } from '@/shared/ui/button'
import { SidebarTrigger } from '@/shared/ui/sidebar'
import { Plus } from 'lucide-react'

export default async function BannersLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  const { data, success } = await getCollections()
  return (
    <div className='flex flex-col w-full'>
      <header className='p-4 items-center flex pt-4'>
        <SidebarTrigger />
        <div className='flex items-center justify-between w-full'>
          <h2 className='font-bold text-lg'>Баннеры</h2>
          {!success && (
            <h2 className='font-bold text-lg text-red-500'>Ошибка загрузки коллекций.</h2>
          )}
          {data && (
            <BannerForm collections={data}>
              <Button variant={'outline'}>
                <Plus />
                Добавить баннер
              </Button>
            </BannerForm>
          )}
        </div>
      </header>
      {children}
    </div>
  )
}
