'use server'
import { fetchWithAuth } from '@/shared/lib/fetcher'
import { revalidateCacheTag } from '@/shared/lib/revalidate'

export async function deleteCollection(id: number) {
  await fetchWithAuth(`/collections/delete/${id}`, { method: 'DELETE' })
  revalidateCacheTag('collections')
}
