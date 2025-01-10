'use server'
import { NewCollectionFields } from '../model/schema'
import { revalidateTag } from 'next/cache'
import { fetchWithAuth } from '@/shared/lib/fetcher'

export async function createCollection(fields: NewCollectionFields) {
  const formData = new FormData()
  formData.append('image', fields.image)
  formData.append('title', fields.title)
  if (fields.description) formData.append('description', fields.description)

  await fetchWithAuth('/collections/create', { method: 'POST', body: formData })

  revalidateTag('collections')
}
