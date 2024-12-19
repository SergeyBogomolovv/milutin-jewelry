'use server'
import { fetcher } from '@/shared/api/fetcher'
import { collectionSchema } from '../model/collection'

export async function getCollection(id: string) {
  const res = await fetcher(`/collections/${id}`, { next: { tags: ['collections', id] } })
  const data = await res.json()
  return collectionSchema.parse(data)
}
