'use client'
import { Button } from '@/shared/ui/button'
import { toast } from 'sonner'
import { useState } from 'react'
import { deleteBanner } from '../api/delete-banner'
import { Trash } from 'lucide-react'
import ConfirmDialog from '@/shared/ui/confirm-dialog'

export default function DeleteButton({ id }: { id: number }) {
  const [isLoading, setLoading] = useState(false)

  const deleteHandler = async () => {
    try {
      setLoading(true)
      await deleteBanner(id)
      setLoading(false)
      toast.success('Баннер удален')
    } catch (error) {
      setLoading(false)
      toast.error('Произошла ошибка удаления баннера')
    }
  }

  return (
    <ConfirmDialog
      title='Подтвердите удаление'
      description='Вы действительно хотите удалить этот баннер?'
      handleConfirm={deleteHandler}
    >
      <Button disabled={isLoading} variant='destructive'>
        <Trash />
        {isLoading ? 'Удаление...' : 'Удалить'}
      </Button>
    </ConfirmDialog>
  )
}
