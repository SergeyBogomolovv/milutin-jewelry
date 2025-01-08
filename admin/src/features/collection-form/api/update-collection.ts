'use server'
import { fetcher } from '@/shared/lib/fetcher'
import { UpdateCollectionFields } from '../model/update.schema'
import { revalidateTag } from 'next/cache'

interface UpdateCollectionResponse {
  success: boolean
  error?: string
}

export const updateCollection = async (
  data: UpdateCollectionFields,
  id: number,
): Promise<UpdateCollectionResponse> => {
  try {
    const formData = new FormData()
    if (data.title) formData.append('title', data.title)
    formData.append('description', data.description || '')
    if (data.image) formData.append('image', data.image)

    const res = await fetcher(`/collections/update/${id}`, {
      method: 'PUT',
      body: formData,
    })

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
