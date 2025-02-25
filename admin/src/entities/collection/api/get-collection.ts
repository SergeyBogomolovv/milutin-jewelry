'use server'
import { fetcher } from '@/shared/lib/fetcher'
import { collectionSchema } from '../model/collection'

export async function getCollection(id: string) {
  const res = await fetcher(`/collections/${id}`, { next: { tags: ['collections'] } })
  const data = await res.json()
  return collectionSchema.parse(data)
}
