'use client'

import { Button } from '@/shared/ui/button'
import { deleteCollection } from '../api/delete-collection'
import { toast } from 'sonner'
import { useState } from 'react'
import ConfirmDialog from '@/shared/ui/confirm-dialog'
import { Trash } from 'lucide-react'

export default function DeleteButton({ id }: { id: number }) {
  const [isLoading, setLoading] = useState(false)

  const deleteHandler = async () => {
    setLoading(true)
    const result = await deleteCollection(id)
    setLoading(false)
    if (!result.success) {
      toast.error(`Ошибка удаления коллекции: ${result.error || 'Неизвестная ошибка'}`)
      return
    }
    toast.success('Коллекция успешно удалена')
  }

  return (
    <ConfirmDialog
      title='Подтвердите удаление'
      description='Вы действительно хотите удалить коллекцию? При удалении коллекции будут также удалены все ее украшения.'
      handleConfirm={deleteHandler}
    >
      <Button disabled={isLoading} variant='destructive'>
        <Trash />
        {isLoading ? 'Удаление...' : 'Удалить'}
      </Button>
    </ConfirmDialog>
  )
}
