import { fetcher } from '@/shared/lib/fetcher'
import { collectionSchema } from '../model/schema'

export async function getCollection(id: string) {
  const res = await fetcher(`/collections/${id}`)
  const json = await res.json()
  return collectionSchema.parse(json)
}
