import { Item } from '../model/collection-item'
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
import { UpdateItemForm } from '@/features/item-form'

export function ItemCard({ item }: { item: Item }) {
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
          alt={item.title || 'Название отсутствует'}
        />
      </CardContent>
      <CardFooter className='flex items-center gap-2'>
        <UpdateItemForm item={item}>
          <Button variant={'outline'}>
            <Pencil />
            Редактировать
          </Button>
        </UpdateItemForm>
        <DeleteButton id={item.id} />
      </CardFooter>
    </Card>
  )
}
