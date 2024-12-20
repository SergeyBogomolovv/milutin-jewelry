'use server'
import { fetcher } from '@/shared/lib/fetcher'
import { revalidateTag } from 'next/cache'

export async function deleteCollection(id: number) {
  try {
    const res = await fetcher(`/collections/delete/${id}`, { method: 'DELETE' })
    if (!res.ok) {
      return false
    }
    revalidateTag('collections')
    return true
  } catch (error) {
    return false
  }
}
