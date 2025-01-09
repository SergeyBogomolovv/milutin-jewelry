'use server'
import { fetcher } from '@/shared/lib/fetcher'
import { collectionSchema } from '../model/collection'

export async function getCollection(id: string) {
  try {
    const res = await fetcher(`/collections/${id}`, {
      next: { tags: ['collections', id] },
      cache: 'force-cache',
    })
    const data = await res.json()
    return collectionSchema.safeParse(data)
  } catch (error) {
    return { success: false, data: undefined }
  }
}
