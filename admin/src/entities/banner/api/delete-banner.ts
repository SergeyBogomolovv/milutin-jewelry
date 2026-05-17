'use server'
import { fetchWithAuth } from '@/shared/lib/fetcher'
import { updateTag } from 'next/cache'

export async function deleteBanner(id: number) {
  await fetchWithAuth(`/banners/delete/${id}`, { method: 'DELETE' })
  updateTag('banners')
}
