import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/shared/ui/card'
import { Banner } from '../model/banner'
import { Image } from '@/shared/ui/image'
import DeleteButton from './delete-button'
import { Collection } from '@/entities/collection'

export function BannerCard({ banner, collections }: { banner: Banner; collections: Collection[] }) {
  let title = 'Без ссылки на коллекцию'
  collections.forEach((collection) => {
    if (collection.id === banner.collection_id) {
      title = collection.title
    }
  })

  return (
    <Card className='flex flex-col'>
      <CardHeader>
        <CardTitle>{title}</CardTitle>
      </CardHeader>
      <CardContent className='flex flex-col gap-2 grow'>
        <CardDescription>Изображение:</CardDescription>
        <Image
          className='w-full object-cover rounded-md aspect-auto'
          src={banner.image_id}
          width={500}
          height={500}
          alt={'Баннер десктоп'}
        />
        <CardDescription className='mt-2'>Изображение для мобильных устройств:</CardDescription>
        <Image
          className='w-full object-cover rounded-md aspect-auto'
          src={banner.mobile_image_id}
          width={500}
          height={500}
          alt={'Мобильный баннер'}
        />
      </CardContent>
      <CardFooter className='flex items-center gap-2'>
        <DeleteButton id={banner.id} />
      </CardFooter>
    </Card>
  )
}
