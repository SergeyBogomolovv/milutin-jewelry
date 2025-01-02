'use server'

import { fetcher } from '@/shared/lib/fetcher'
import { revalidateTag } from 'next/cache'

interface DeleteItemResponse {
  success: boolean
  error?: string
}

export async function deleteItem(id: number): Promise<DeleteItemResponse> {
  try {
    const res = await fetcher(`/items/delete/${id}`, { method: 'DELETE' })

    if (!res.ok) {
      const errorMessage = await res.text()
      return { success: false, error: errorMessage || 'Ошибка при удалении украшения' }
    }

    revalidateTag('items')
    return { success: true }
  } catch (error) {
    return { success: false, error: (error as Error).message || 'Произошла ошибка при удалении' }
  }
}
