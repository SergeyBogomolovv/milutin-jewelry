'use client'
import { Button } from '@/shared/ui/button'
import { deleteCollection } from '../api/delete-collection'
import { toast } from 'sonner'
import { useState } from 'react'

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
    <Button disabled={isLoading} onClick={deleteHandler} variant={'destructive'}>
      {isLoading ? 'Удаление...' : 'Удалить'}
    </Button>
  )
}
