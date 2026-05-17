'use server'
import { fetchWithAuth } from '@/shared/lib/fetcher'
import { updateTag } from 'next/cache'

export async function deleteCollection(id: number) {
  await fetchWithAuth(`/collections/delete/${id}`, { method: 'DELETE' })
  updateTag('collections')
}
