'use server'
import { fetcher } from '@/shared/lib/fetcher'
import { NewBannerFields } from '../model/schema'
import { revalidateTag } from 'next/cache'

interface CreateBannerResponse {
  success: boolean
  error?: string
}

export const createBanner = async (fields: NewBannerFields): Promise<CreateBannerResponse> => {
  try {
    const formData = new FormData()
    if (fields.collection_id) formData.append('collection_id', fields.collection_id)
    formData.append('description', fields.image)
    formData.append('description', fields.mobile_image)

    const res = await fetcher('/banners/create', { method: 'POST', body: formData })

    if (!res.ok) {
      const errorMessage = await res.text()
      return { success: false, error: errorMessage || 'Неизвестная ошибка' }
    }

    revalidateTag('banners')
    return { success: true }
  } catch (error) {
    return { success: false, error: (error as Error).message || 'Произошла ошибка' }
  }
}
