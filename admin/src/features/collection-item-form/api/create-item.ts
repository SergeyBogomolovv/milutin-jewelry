'use server'
import { fetcher } from '@/shared/lib/fetcher'
import { NewItemFields } from '../model/new-item.schema'

export const createCollectionItem = async (fields: NewItemFields, collectionId: string) => {
  try {
    const formData = new FormData()
    formData.append('title', fields.title)
    if (fields.description) formData.append('description', fields.description)
    formData.append('image', fields.image)
    formData.append('collection_id', collectionId)
    const res = await fetcher('/collection-items/create', { method: 'POST', body: formData })
    if (!res.ok) {
      return false
    }
    return true
  } catch (error) {
    return false
  }
}
