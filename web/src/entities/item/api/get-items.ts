import { fetcher } from '@/shared/lib/fetcher'
import { z } from 'zod'
import { itemSchema } from '../model/schema'

export async function getItems(collectionId: string) {
  const res = await fetcher(`/items/collection/${collectionId}`, {
    cache: 'force-cache',
    next: { revalidate: 300 },
  })
  const json = await res.json()
  return z.array(itemSchema).parse(json)
}
