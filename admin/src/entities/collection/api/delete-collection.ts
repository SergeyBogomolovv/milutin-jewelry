'use server'
import { fetchWithAuth } from '@/shared/lib/fetcher'
import { revalidateTag } from 'next/cache'

export async function deleteCollection(id: number) {
  await fetchWithAuth(`/collections/delete/${id}`, { method: 'DELETE' })
  revalidateTag('collections')
}
