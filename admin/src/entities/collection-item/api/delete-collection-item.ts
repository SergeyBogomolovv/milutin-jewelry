'use server'
import { fetcher } from '@/shared/lib/fetcher'
import { revalidateTag } from 'next/cache'

export async function deleteCollectionItem(id: number) {
  try {
    const res = await fetcher(`/collection-items/delete/${id}`, { method: 'DELETE' })
    if (!res.ok) {
      return false
    }
    revalidateTag('collection-items')
    return true
  } catch (error) {
    return false
  }
}
