'use client'
import { Button } from '@/shared/ui/button'
import { toast } from 'sonner'
import { useState } from 'react'
import { deleteCollectionItem } from '../api/delete-collection-item'

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
    <Button disabled={isLoading} onClick={deleteHandler} variant={'destructive'}>
      {isLoading ? 'Удаление...' : 'Удалить'}
    </Button>
  )
}
