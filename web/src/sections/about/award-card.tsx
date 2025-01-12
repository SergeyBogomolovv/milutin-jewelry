import { Separator } from '@/shared/ui/separator'
import { Image } from 'antd'

interface Props {
  src: string
  description: string
}

export function AwardCard({ src, description }: Props) {
  return (
    <div className='flex flex-col gap-4 items-center text-center'>
      <Image src={src} alt='Фотография' />
      <Separator />
      <p className='text-lg font-bold'>{description}</p>
    </div>
  )
}
