import { fetcher } from '@/shared/lib/fetcher'
import { collectionSchema } from '../model/schema'

export async function getCollection(id: string) {
  const res = await fetcher(`/collections/${id}`, {
    cache: 'force-cache',
    next: { revalidate: 300 },
  })
  const json = await res.json()
  return collectionSchema.parse(json)
}
