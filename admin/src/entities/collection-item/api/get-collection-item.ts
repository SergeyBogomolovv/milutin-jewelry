'use server'
import { fetcher } from '@/shared/lib/fetcher'
import { collectionItemSchema } from '../model/collection-item'

export async function getCollectionItem(id: string) {
  const res = await fetcher(`/collection-items/${id}`)
  const data = await res.json()
  return collectionItemSchema.parse(data)
}
