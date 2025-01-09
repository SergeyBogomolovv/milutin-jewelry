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
    setLoading(true)
    const result = await deleteBanner(id)
    setLoading(false)
    if (!result.success) {
      toast.error(`Ошибка удаления баннера: ${result.error || 'Неизвестная ошибка'}`)
      return
    }
    toast.success('Баннер удален')
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
