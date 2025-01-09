import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/shared/ui/card'
import { Banner } from '../model/banner'
import { CustomImage } from '@/shared/ui/image'
import DeleteButton from './delete-button'

export function BannerCard({ banner }: { banner: Banner }) {
  return (
    <Card className='flex flex-col'>
      <CardHeader>
        {/* TODO: Add collection */}
        <CardTitle>Без ссылки на коллекцию</CardTitle>
      </CardHeader>
      <CardContent className='flex flex-col gap-2'>
        <CardDescription>Изображение:</CardDescription>
        <CustomImage
          className='w-full object-cover rounded-md aspect-auto grow'
          src={banner.image_id}
          width={500}
          height={500}
          alt={'Баннер десктоп'}
        />
        <CardDescription className='mt-2'>Изображение для мобильных устройств:</CardDescription>
        <CustomImage
          className='w-full object-cover rounded-md aspect-auto grow'
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
