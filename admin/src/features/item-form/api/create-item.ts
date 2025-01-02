'use server'
import { fetcher } from '@/shared/lib/fetcher'
import { NewItemFields } from '../model/new-item.schema'
import { revalidateTag } from 'next/cache'

export const createItem = async (fields: NewItemFields, collectionId: string) => {
  try {
    const formData = new FormData()
    if (fields.title) formData.append('title', fields.title)
    if (fields.description) formData.append('description', fields.description)
    formData.append('image', fields.image)
    formData.append('collection_id', collectionId)
    const res = await fetcher('/items/create', { method: 'POST', body: formData })
    if (!res.ok) {
      return false
    }
    revalidateTag('items')
    return true
  } catch (error) {
    return false
  }
}
