'use server'
import { fetchWithAuth } from '@/shared/lib/fetcher'
import { UpdateCollectionFields } from '../model/update.schema'
import { revalidateTag } from 'next/cache'

export async function updateCollection(data: UpdateCollectionFields, id: number) {
  const formData = new FormData()
  if (data.title) formData.append('title', data.title)
  formData.append('description', data.description || '')
  if (data.image) formData.append('image', data.image)

  await fetchWithAuth(`/collections/update/${id}`, { method: 'PUT', body: formData })

  revalidateTag('collections')
}
