'use server'
import { fetchWithAuth } from '@/shared/lib/fetcher'
import { updateTag } from 'next/cache'

export async function deleteItem(id: number) {
  await fetchWithAuth(`/items/delete/${id}`, { method: 'DELETE' })
  updateTag('items')
}
