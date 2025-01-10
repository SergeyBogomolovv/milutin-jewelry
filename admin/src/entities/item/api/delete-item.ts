'use server'
import { fetchWithAuth } from '@/shared/lib/fetcher'
import { revalidateTag } from 'next/cache'

export async function deleteItem(id: number) {
  await fetchWithAuth(`/items/delete/${id}`, { method: 'DELETE' })
  revalidateTag('items')
}
