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

interface Props {
  collection: Collection
}

export function CollectionCard({ collection }: Props) {
  return (
    <Card>
      <CardHeader>
        <CardTitle>{collection.title}</CardTitle>
        <CardDescription>
          Lorem ipsum dolor, sit amet consectetur adipisicing elit. Id consectetur a suscipit quod
          impedit architecto nihil molestias, explicabo dolorum, vel ipsam dignissimos temporibus
          dolor tempora corporis assumenda quidem necessitatibus laudantium.
        </CardDescription>
      </CardHeader>
      <CardContent>
        <CustomImage
          className='w-full object-cover rounded-md'
          src={collection.image_id}
          width={500}
          height={500}
          alt={collection.title}
        />
      </CardContent>
      <CardFooter className='flex items-center gap-2'>
        <Button variant={'outline'}>Перейти</Button>
        <Button variant={'destructive'}>Удалить</Button>
      </CardFooter>
    </Card>
  )
}
