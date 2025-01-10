'use server'
import { fetchWithAuth } from '@/shared/lib/fetcher'
import { NewItemFields } from '../model/new-item.schema'
import { revalidateTag } from 'next/cache'

export const createItem = async (fields: NewItemFields, collectionId: string) => {
  const formData = new FormData()

  if (fields.title) formData.append('title', fields.title)
  if (fields.description) formData.append('description', fields.description)
  formData.append('image', fields.image)
  formData.append('collection_id', collectionId)

  await fetchWithAuth('/items/create', { method: 'POST', body: formData })

  revalidateTag('items')
}
