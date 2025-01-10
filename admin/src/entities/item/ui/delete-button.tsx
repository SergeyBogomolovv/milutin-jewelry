'use client'
import { Button } from '@/shared/ui/button'
import { toast } from 'sonner'
import { useState } from 'react'
import { deleteItem } from '../api/delete-item'
import ConfirmDialog from '@/shared/ui/confirm-dialog'
import { Trash } from 'lucide-react'

export default function DeleteButton({ id }: { id: number }) {
  const [isLoading, setLoading] = useState(false)

  const deleteHandler = async () => {
    try {
      setLoading(true)
      await deleteItem(id)
      setLoading(false)
      toast.success('Украшение удалено')
    } catch (error) {
      setLoading(false)
      toast.error('Произошла ошибка удаления украшения')
    }
  }

  return (
    <ConfirmDialog
      title='Подтвердите удаление'
      description='Вы действительно хотите удалить это украшение?'
      handleConfirm={deleteHandler}
    >
      <Button disabled={isLoading} variant='destructive'>
        <Trash />
        {isLoading ? 'Удаление...' : 'Удалить'}
      </Button>
    </ConfirmDialog>
  )
}
