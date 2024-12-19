'use server'
import { fetcher } from '@/shared/api/fetcher'
import { NewCollectionFields } from '../model/schema'
import { revalidateTag } from 'next/cache'

export const createCollection = async (fields: NewCollectionFields) => {
  try {
    const formData = new FormData()
    formData.append('title', fields.title)
    if (fields.description) formData.append('description', fields.description)
    formData.append('image', fields.image)
    const res = await fetcher('/collections/create', {
      method: 'POST',
      body: formData,
    })
    if (!res.ok) {
      return false
    }
    revalidateTag('collections')
    return true
  } catch (error) {
    return false
  }
}
