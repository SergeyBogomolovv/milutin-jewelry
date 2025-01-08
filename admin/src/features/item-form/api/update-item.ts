'use server'
import { fetcher } from '@/shared/lib/fetcher'
import { revalidateTag } from 'next/cache'
import { UpdateItemFields } from '../model/update-item.schema'

export const updateItem = async (fields: UpdateItemFields, collectionId: string) => {
  try {
    const formData = new FormData()
    formData.append('title', fields.title || '')
    formData.append('description', fields.description || '')
    if (fields.image) formData.append('image', fields.image)
    formData.append('collection_id', collectionId)
    const res = await fetcher(`/items/update/${collectionId}`, {
      method: 'PUT',
      body: formData,
    })
    if (!res.ok) {
      return false
    }
    revalidateTag('items')
    return true
  } catch (error) {
    return false
  }
}
