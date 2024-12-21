import { Button } from '@/shared/ui/button'
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/shared/ui/card'
import { Collection } from '../model/collection'
import { CustomImage } from '@/shared/ui/image'
import Link from 'next/link'
import DeleteButton from './delete-button'
import { Pencil } from 'lucide-react'

interface Props {
  collection: Collection
}

export function CollectionCard({ collection }: Props) {
  return (
    <Card className='flex flex-col'>
      <CardHeader>
        <CardTitle>{collection.title}</CardTitle>
        <CardDescription>
          {collection.description ? collection.description : 'Описание отсутствует'}
        </CardDescription>
      </CardHeader>
      <CardContent className='grow flex'>
        <CustomImage
          className='w-full object-cover rounded-md aspect-auto grow'
          src={collection.image_id}
          width={500}
          height={500}
          alt={collection.title}
        />
      </CardContent>
      <CardFooter className='flex items-center gap-2'>
        <Button asChild variant={'outline'}>
          <Link href={`/collections/${collection.id}`}>
            <Pencil />
            Управление
          </Link>
        </Button>
        <DeleteButton id={collection.id} />
      </CardFooter>
    </Card>
  )
}
