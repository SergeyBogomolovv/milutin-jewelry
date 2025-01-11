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
import { Image } from '@/shared/ui/image'
import Link from 'next/link'
import DeleteButton from './delete-button'
import { Pencil } from 'lucide-react'
import { UpdateCollectionForm } from '@/features/collection-form'
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from '@/shared/ui/tooltip'

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
      <TooltipProvider>
        <Tooltip>
          <TooltipTrigger asChild>
            <CardContent className='grow flex'>
              <Link href={`/${collection.id}`} className='grow flex'>
                <Image
                  className='w-full object-cover rounded-md aspect-auto grow'
                  src={collection.image_id}
                  width={500}
                  height={500}
                  alt={collection.title}
                />
              </Link>
            </CardContent>
          </TooltipTrigger>
          <TooltipContent>
            <p>Перейти на страницу коллекции</p>
          </TooltipContent>
        </Tooltip>
      </TooltipProvider>
      <CardFooter className='flex items-center gap-2'>
        <UpdateCollectionForm collection={collection}>
          <Button variant={'outline'}>
            <Pencil />
            Редактировать
          </Button>
        </UpdateCollectionForm>
        <DeleteButton id={collection.id} />
      </CardFooter>
    </Card>
  )
}
