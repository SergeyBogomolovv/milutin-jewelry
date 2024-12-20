'use server'
import { fetcher } from '@/shared/lib/fetcher'
import { z } from 'zod'
import { collectionItemSchema } from '../model/collection-item'

export async function getCollectionItemsByCollection(id: string) {
  const res = await fetcher(`/collection-items/collection/${id}`, {
    next: { tags: ['collection-items', id] },
    cache: 'force-cache',
  })
  const data = await res.json()
  return z.array(collectionItemSchema).parse(data)
}
