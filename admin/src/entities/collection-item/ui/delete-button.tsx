'use client'
import { Button } from '@/shared/ui/button'
import { toast } from 'sonner'
import { useState } from 'react'
import { deleteCollectionItem } from '../api/delete-collection-item'
import ConfirmDialog from '@/shared/ui/confirm-dialog'
import { Trash } from 'lucide-react'

export default function DeleteButton({ id }: { id: number }) {
  const [isLoading, setLoading] = useState(false)
  const deleteHandler = async () => {
    setLoading(true)
    const ok = await deleteCollectionItem(id)
    setLoading(false)
    if (!ok) {
      toast.error('Ошибка удаления украшения')
      return
    }
    toast.success('Украшение удалено')
  }
  return (
    <ConfirmDialog
      title='Подтвердите удаление'
      description={`Вы действительно хотите удалить это украшение?`}
      handleConfirm={deleteHandler}
    >
      <Button disabled={isLoading} variant={'destructive'}>
        <Trash />
        {isLoading ? 'Удаление...' : 'Удалить'}
      </Button>
    </ConfirmDialog>
  )
}
