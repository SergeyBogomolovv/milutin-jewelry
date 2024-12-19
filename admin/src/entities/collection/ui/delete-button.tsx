'use client'
import { Button } from '@/shared/ui/button'
import { deleteCollection } from '../api/delete-collection'

export default function DeleteButton({ id }: { id: number }) {
  return (
    <Button onClick={() => deleteCollection(id)} variant={'destructive'}>
      Удалить
    </Button>
  )
}
