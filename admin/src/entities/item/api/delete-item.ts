'use server'
import { fetchWithAuth } from '@/shared/lib/fetcher'
import { revalidateCacheTag } from '@/shared/lib/revalidate'

export async function deleteItem(id: number) {
  await fetchWithAuth(`/items/delete/${id}`, { method: 'DELETE' })
  revalidateCacheTag('items')
}
