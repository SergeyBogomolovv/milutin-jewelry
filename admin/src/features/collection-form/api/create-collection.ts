'use server'

import { NewCollectionFields } from '../model/schema'
import { revalidateTag } from 'next/cache'
import { fetcher } from '@/shared/lib/fetcher'

interface CreateCollectionResponse {
  success: boolean
  error?: string
}

export const createCollection = async (
  fields: NewCollectionFields,
): Promise<CreateCollectionResponse> => {
  try {
    const formData = new FormData()
    formData.append('title', fields.title)
    if (fields.description) formData.append('description', fields.description)
    formData.append('image', fields.image)

    const res = await fetcher('/collections/create', { method: 'POST', body: formData })

    if (!res.ok) {
      const errorMessage = await res.text()
      return { success: false, error: errorMessage || 'Неизвестная ошибка' }
    }

    revalidateTag('collections')
    return { success: true }
  } catch (error) {
    return { success: false, error: (error as Error).message || 'Произошла ошибка' }
  }
}
