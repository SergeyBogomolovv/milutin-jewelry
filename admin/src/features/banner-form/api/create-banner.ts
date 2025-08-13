'use server'
import { fetchWithAuth } from '@/shared/lib/fetcher'
import { NewBannerFields } from '../model/schema'
import { revalidateTag } from 'next/cache'

export async function createBanner(fields: NewBannerFields) {
  const formData = new FormData()

  if (fields.collection_id) formData.append('collection_id', fields.collection_id)
  formData.append('image', fields.image)
  formData.append('mobile_image', fields.mobile_image)

  const res = await fetchWithAuth('/banners/create', { method: 'POST', body: formData })

  if (!res.ok) {
    throw new Error('Failed to create banner')
  }

  revalidateTag('banners')
}
