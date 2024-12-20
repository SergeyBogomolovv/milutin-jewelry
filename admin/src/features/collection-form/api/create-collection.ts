'use server'
import { NewCollectionFields } from '../model/schema'
import { revalidateTag } from 'next/cache'
import { fetcher } from '@/shared/lib/fetcher'

export const createCollection = async (fields: NewCollectionFields) => {
  try {
    const formData = new FormData()
    formData.append('title', fields.title)
    if (fields.description) formData.append('description', fields.description)
    formData.append('image', fields.image)
    const res = await fetcher('/collections/create', { method: 'POST', body: formData })
    if (!res.ok) {
      return false
    }
    revalidateTag('collections')
    return true
  } catch (error) {
    return false
  }
}
