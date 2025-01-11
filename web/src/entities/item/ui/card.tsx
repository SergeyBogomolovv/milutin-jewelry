import { Image } from '@/shared/ui/image'
import { Item } from '../model/schema'
import Line from '@/assets/line.svg'

export function ItemCard({ item }: { item: Item }) {
  return (
    <div className='flex flex-col items-center'>
      <Image
        src={item.image_id}
        width={'100%'}
        alt={item.title || 'Название отсутствует'}
        className='object-cover rounded-lg w-full h-full aspect-auto'
      />
      {(item.description || item.title) && (
        <>
          <Line className='w-full text-muted-foreground mt-6 mb-2' />
          <div className='text-center space-y-2'>
            {item.title && <p className='text-lg font-bold tracking-widest'>{item.title}</p>}
            {item.description && <p className='tracking-widest'>{item.description}</p>}
          </div>
        </>
      )}
    </div>
  )
}
