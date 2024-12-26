'use server'

import { fetcher } from '@/shared/lib/fetcher'
import { revalidateTag } from 'next/cache'

interface DeleteCollectionItemResponse {
  success: boolean
  error?: string
}

export async function deleteCollectionItem(id: number): Promise<DeleteCollectionItemResponse> {
  try {
    const res = await fetcher(`/collection-items/delete/${id}`, { method: 'DELETE' })

    if (!res.ok) {
      const errorMessage = await res.text()
      return { success: false, error: errorMessage || 'Ошибка при удалении украшения' }
    }

    revalidateTag('collection-items')
    return { success: true }
  } catch (error) {
    return { success: false, error: (error as Error).message || 'Произошла ошибка при удалении' }
  }
}
