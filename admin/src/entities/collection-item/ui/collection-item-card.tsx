import { CollectionItem } from '../model/collection-item'
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/shared/ui/card'
import { CustomImage } from '@/shared/ui/image'
import { Button } from '@/shared/ui/button'
import DeleteButton from './delete-button'
import { Pencil } from 'lucide-react'
import { UpdateCollectionItemForm } from '@/features/collection-item-form/ui/update-collection-item'

export function CollectionItemCard({ item }: { item: CollectionItem }) {
  return (
    <Card className='flex flex-col'>
      <CardHeader>
        <CardTitle>{item.title ? item.title : 'Название отсутствует'}</CardTitle>
        <CardDescription>
          {item.description ? item.description : 'Описание отсутствует'}
        </CardDescription>
      </CardHeader>
      <CardContent className='grow flex'>
        <CustomImage
          className='w-full object-cover rounded-md aspect-auto grow'
          src={item.image_id}
          width={500}
          height={500}
          alt={item.title}
        />
      </CardContent>
      <CardFooter className='flex items-center gap-2'>
        <UpdateCollectionItemForm collectionItem={item}>
          <Button variant={'outline'}>
            <Pencil />
            Редактировать
          </Button>
        </UpdateCollectionItemForm>
        <DeleteButton id={item.id} />
      </CardFooter>
    </Card>
  )
}
