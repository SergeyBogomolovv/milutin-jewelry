'use server'
import { fetcher } from '@/shared/lib/fetcher'
import { collectionSchema } from '../model/schema'

export async function getCollection(id: string) {
  try {
    const res = await fetcher(`/collections/${id}`, { next: { tags: ['collections', id] } })
    const json = await res.json()
    const { data, success } = collectionSchema.safeParse(json)
    return { data, success }
  } catch (error) {
    return { success: false, data: undefined }
  }
}
