'use client'
import { Button } from '@/shared/ui/button'
import { deleteCollection } from '../api/delete-collection'
import { toast } from 'sonner'

export default function DeleteButton({ id }: { id: number }) {
  return (
    <Button
      onClick={async () => {
        const ok = await deleteCollection(id)
        if (!ok) {
          toast.error('Ошибка удаления коллекции')
          return
        }
        toast.success('Коллекция удалена')
      }}
      variant={'destructive'}
    >
      Удалить
    </Button>
  )
}
