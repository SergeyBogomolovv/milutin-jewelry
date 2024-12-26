'use server'

import { fetcher } from '@/shared/lib/fetcher'
import { revalidateTag } from 'next/cache'

interface DeleteCollectionResponse {
  success: boolean
  error?: string
}

export async function deleteCollection(id: number): Promise<DeleteCollectionResponse> {
  try {
    const res = await fetcher(`/collections/delete/${id}`, { method: 'DELETE' })

    if (!res.ok) {
      const errorMessage = await res.text()
      return { success: false, error: errorMessage || 'Ошибка при удалении коллекции' }
    }

    revalidateTag('collections')
    return { success: true }
  } catch (error) {
    return { success: false, error: (error as Error).message || 'Произошла ошибка при удалении' }
  }
}
