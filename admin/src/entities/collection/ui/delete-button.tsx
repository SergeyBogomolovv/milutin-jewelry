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
    const ok = await deleteCollection(id)
    setLoading(false)
    if (!ok) {
      toast.error('Ошибка удаления коллекции')
      return
    }
    toast.success('Коллекция удалена')
  }
  return (
    <ConfirmDialog
      title='Подтвердите удаление'
      description={`Вы действительно хотите удалить коллекцию? При удалении коллекции будут так же удалены все ее украшения.`}
      handleConfirm={deleteHandler}
    >
      <Button disabled={isLoading} variant={'destructive'}>
        <Trash />
        {isLoading ? 'Удаление...' : 'Удалить'}
      </Button>
    </ConfirmDialog>
  )
}
