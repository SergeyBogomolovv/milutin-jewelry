'use server'
import { fetchWithAuth } from '@/shared/lib/fetcher'
import { revalidateTag } from 'next/cache'
import { UpdateItemFields } from '../model/update-item.schema'

export const updateItem = async (fields: UpdateItemFields, id: number) => {
  const formData = new FormData()
  formData.append('title', fields.title || '')
  formData.append('description', fields.description || '')
  if (fields.image) formData.append('image', fields.image)

  await fetchWithAuth(`/items/update/${id}`, { method: 'PUT', body: formData })

  revalidateTag('items')
}
