import type { Collection } from '../model/schema'
import Link from 'next/link'
import Line from '@/assets/line.svg'
import S3Image from '@/shared/ui/s3-image'

export function CollectionCard({ collection }: { collection: Collection }) {
  return (
    <Link href={`/${collection.id}`} className='flex flex-col items-center gap-5'>
      <S3Image
        alt={collection.title}
        src={collection.image_id}
        width={500}
        height={500}
        className='rounded-lg aspect-auto object-cover grow'
      />
      <Line className='w-full text-muted-foreground' />
      <p className='text-xl font-bold tracking-widest'>{collection.title}</p>
    </Link>
  )
}
